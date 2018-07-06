package hydra

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"sync"
	"syscall"
	"time"

	"github.com/fujiwara/fluent-agent-hydra/fluent"
)

const (
	OpenRetryInterval = 1 * time.Second
	SEEK_TAIL         = int64(-1)
	SEEK_HEAD         = int64(0)
	DEBUG             = false
)

var (
	ReadBufferSize = 64 * 1024
)

var m *sync.Mutex

type File struct {
	*os.File
	Path           string
	Tag            string
	Position       int64
	readBuf        []byte
	contBuf        []byte
	lastStat       os.FileInfo
	FieldName      string
	FileStat       *FileStat
	Format         FileFormat
	RecordModifier *RecordModifier
	Regexp         *Regexp
	Inode          uint64
	Pf             *os.File
}

func (f *File) ReadPosfile() (uint64, int64) {
	var Inode uint64
	var Position int64
	m.Lock()
	fmt.Fscanf(f.Pf, "%d\t%d", &Inode, &Position)
	m.Unlock()
	return Inode, Position
}

func (f *File) WritePosfile(inode uint64, position int64) error {
	m.Lock()
	f.Pf.Seek(int64(0), os.SEEK_SET)
	_, err := fmt.Fprintf(f.Pf, "%d\t%d", inode, position)
	m.Unlock()
	return err
}

func openFile(path string, startPos int64) (*File, error) {
	m = new(sync.Mutex)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	pf, err := os.OpenFile(path+".pos", os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return nil, err
	}
	var st syscall.Stat_t
	err = syscall.Stat(path, &st)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	stat, err := f.Stat()
	if err != nil {
		return nil, err
	}

	file := &File{
		f,
		path,
		"",
		startPos,
		make([]byte, ReadBufferSize),
		make([]byte, 0),
		stat,
		"",
		&FileStat{},
		FormatNone,
		nil,
		nil,
		st.Ino,
		pf,
	}
	pfIno, pfPos := file.ReadPosfile()
	if pfIno != st.Ino {
		file.WritePosfile(st.Ino, int64(0))
		startPos = int64(0)
	} else if startPos != pfPos {
		startPos = pfPos
	}

	if startPos == SEEK_TAIL {
		// seek to end of file
		size := file.lastStat.Size()
		pos, _ := file.Seek(size, os.SEEK_SET)
		file.Position = pos
	} else {
		pos, _ := file.Seek(startPos, os.SEEK_SET)
		file.Position = pos
	}
	log.Println("[info]", file.Path, "Seeked To", file.Position)
	log.Println("[info]", file.Path, "Inode", file.Inode)
	return file, nil
}

func (f *File) restrict() error {
	var err error
	f.lastStat, err = f.Stat()
	if err != nil {
		log.Println("[error]", f.Path, "stat failed", err)
		return err
	}
	if size := f.lastStat.Size(); size < f.Position {
	  pos, _ := f.Seek(size, os.SEEK_SET)
		f.Position = pos
		log.Println("[info]", f.Path, "was truncated. Seeked to", pos)
	}
	return nil
}

func (f *File) tailAndSend(messageCh chan *fluent.FluentRecordSet, monitorCh chan Stat) error {
	for {
		n, err := io.ReadAtLeast(f, f.readBuf, 1)
		if n == 0 || err == io.EOF {
			return err
		} else if err != nil {
			return err
		}
		f.Position += int64(n)
		sendBuf := make([]byte, 0)
		if f.readBuf[n-1] == '\n' {
			// f.readBuf is just terminated by '\n'
			if len(f.contBuf) > 0 {
				sendBuf = append(sendBuf, f.contBuf...)
				f.contBuf = make([]byte, 0)
			}
			sendBuf = append(sendBuf, f.readBuf[0:n-1]...)
		} else {
			blockLen := bytes.LastIndex(f.readBuf[0:n], LineSeparator)
			if blockLen == -1 {
				// whole of f.readBuf is continuous line
				f.contBuf = append(f.contBuf, f.readBuf[0:n]...)
				continue
			} else {
				// bottom line of f.readBuf is continuous line
				if len(f.contBuf) > 0 {
					sendBuf = append(sendBuf, f.contBuf...)
				}
				sendBuf = append(sendBuf, f.readBuf[0:blockLen]...)
				f.contBuf = make([]byte, n-blockLen-1)
				copy(f.contBuf, f.readBuf[blockLen+1:n])
			}
		}
		messageCh <- NewFluentRecordSet(f.Tag, f.FieldName, f.Format, f.RecordModifier, f.Regexp, sendBuf)
		monitorCh <- f.UpdateStat()
	}
}

func (f *File) UpdateStat() *FileStat {
	f.WritePosfile(f.Inode, f.Position)
	f.FileStat.File = f.Path
	f.FileStat.Position = f.Position
	f.FileStat.Tag = f.Tag
	return f.FileStat
}
