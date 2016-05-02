// DO NOT EDIT!
// Code generated by ffjson <https://github.com/pquerna/ffjson>
// source: hydra.go
// DO NOT EDIT!

package hydra

import (
	"bytes"
	"encoding/json"
	"fmt"
	fflib "github.com/pquerna/ffjson/fflib/v1"
)

func (mj *Context) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Context) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{"MessageCh":`)
	/* Falling back. type=chan *fluent.FluentRecordSet kind=chan */
	err = buf.Encode(mj.MessageCh)
	if err != nil {
		return err
	}
	buf.WriteString(`,"MonitorCh":`)
	/* Falling back. type=chan hydra.Stat kind=chan */
	err = buf.Encode(mj.MonitorCh)
	if err != nil {
		return err
	}
	buf.WriteString(`,"ControlCh":`)
	/* Falling back. type=chan interface {} kind=chan */
	err = buf.Encode(mj.ControlCh)
	if err != nil {
		return err
	}
	/* Struct fall back. type=sync.WaitGroup kind=struct */
	buf.WriteString(`,"InputProcess":`)
	err = buf.Encode(&mj.InputProcess)
	if err != nil {
		return err
	}
	/* Struct fall back. type=sync.WaitGroup kind=struct */
	buf.WriteString(`,"OutputProcess":`)
	err = buf.Encode(&mj.OutputProcess)
	if err != nil {
		return err
	}
	/* Struct fall back. type=sync.WaitGroup kind=struct */
	buf.WriteString(`,"StartProcess":`)
	err = buf.Encode(&mj.StartProcess)
	if err != nil {
		return err
	}
	buf.WriteByte('}')
	return nil
}

const (
	ffj_t_Contextbase = iota
	ffj_t_Contextno_such_key

	ffj_t_Context_MessageCh

	ffj_t_Context_MonitorCh

	ffj_t_Context_ControlCh

	ffj_t_Context_InputProcess

	ffj_t_Context_OutputProcess

	ffj_t_Context_StartProcess
)

var ffj_key_Context_MessageCh = []byte("MessageCh")

var ffj_key_Context_MonitorCh = []byte("MonitorCh")

var ffj_key_Context_ControlCh = []byte("ControlCh")

var ffj_key_Context_InputProcess = []byte("InputProcess")

var ffj_key_Context_OutputProcess = []byte("OutputProcess")

var ffj_key_Context_StartProcess = []byte("StartProcess")

func (uj *Context) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Context) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Contextbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_Contextno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				case 'C':

					if bytes.Equal(ffj_key_Context_ControlCh, kn) {
						currentKey = ffj_t_Context_ControlCh
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'I':

					if bytes.Equal(ffj_key_Context_InputProcess, kn) {
						currentKey = ffj_t_Context_InputProcess
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'M':

					if bytes.Equal(ffj_key_Context_MessageCh, kn) {
						currentKey = ffj_t_Context_MessageCh
						state = fflib.FFParse_want_colon
						goto mainparse

					} else if bytes.Equal(ffj_key_Context_MonitorCh, kn) {
						currentKey = ffj_t_Context_MonitorCh
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'O':

					if bytes.Equal(ffj_key_Context_OutputProcess, kn) {
						currentKey = ffj_t_Context_OutputProcess
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				case 'S':

					if bytes.Equal(ffj_key_Context_StartProcess, kn) {
						currentKey = ffj_t_Context_StartProcess
						state = fflib.FFParse_want_colon
						goto mainparse
					}

				}

				if fflib.EqualFoldRight(ffj_key_Context_StartProcess, kn) {
					currentKey = ffj_t_Context_StartProcess
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Context_OutputProcess, kn) {
					currentKey = ffj_t_Context_OutputProcess
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Context_InputProcess, kn) {
					currentKey = ffj_t_Context_InputProcess
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Context_ControlCh, kn) {
					currentKey = ffj_t_Context_ControlCh
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.SimpleLetterEqualFold(ffj_key_Context_MonitorCh, kn) {
					currentKey = ffj_t_Context_MonitorCh
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				if fflib.EqualFoldRight(ffj_key_Context_MessageCh, kn) {
					currentKey = ffj_t_Context_MessageCh
					state = fflib.FFParse_want_colon
					goto mainparse
				}

				currentKey = ffj_t_Contextno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_Context_MessageCh:
					goto handle_MessageCh

				case ffj_t_Context_MonitorCh:
					goto handle_MonitorCh

				case ffj_t_Context_ControlCh:
					goto handle_ControlCh

				case ffj_t_Context_InputProcess:
					goto handle_InputProcess

				case ffj_t_Context_OutputProcess:
					goto handle_OutputProcess

				case ffj_t_Context_StartProcess:
					goto handle_StartProcess

				case ffj_t_Contextno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

handle_MessageCh:

	/* handler: uj.MessageCh type=chan *fluent.FluentRecordSet kind=chan quoted=false*/

	{
		/* Falling back. type=chan *fluent.FluentRecordSet kind=chan */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.MessageCh)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_MonitorCh:

	/* handler: uj.MonitorCh type=chan hydra.Stat kind=chan quoted=false*/

	{
		/* Falling back. type=chan hydra.Stat kind=chan */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.MonitorCh)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_ControlCh:

	/* handler: uj.ControlCh type=chan interface {} kind=chan quoted=false*/

	{
		/* Falling back. type=chan interface {} kind=chan */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.ControlCh)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_InputProcess:

	/* handler: uj.InputProcess type=sync.WaitGroup kind=struct quoted=false*/

	{
		/* Falling back. type=sync.WaitGroup kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.InputProcess)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_OutputProcess:

	/* handler: uj.OutputProcess type=sync.WaitGroup kind=struct quoted=false*/

	{
		/* Falling back. type=sync.WaitGroup kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.OutputProcess)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

handle_StartProcess:

	/* handler: uj.StartProcess type=sync.WaitGroup kind=struct quoted=false*/

	{
		/* Falling back. type=sync.WaitGroup kind=struct */
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = json.Unmarshal(tbuf, &uj.StartProcess)
		if err != nil {
			return fs.WrapErr(err)
		}
	}

	state = fflib.FFParse_after_value
	goto mainparse

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}

func (mj *Signal) MarshalJSON() ([]byte, error) {
	var buf fflib.Buffer
	if mj == nil {
		buf.WriteString("null")
		return buf.Bytes(), nil
	}
	err := mj.MarshalJSONBuf(&buf)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
func (mj *Signal) MarshalJSONBuf(buf fflib.EncodingBuffer) error {
	if mj == nil {
		buf.WriteString("null")
		return nil
	}
	var err error
	var obj []byte
	_ = obj
	_ = err
	buf.WriteString(`{}`)
	return nil
}

const (
	ffj_t_Signalbase = iota
	ffj_t_Signalno_such_key
)

func (uj *Signal) UnmarshalJSON(input []byte) error {
	fs := fflib.NewFFLexer(input)
	return uj.UnmarshalJSONFFLexer(fs, fflib.FFParse_map_start)
}

func (uj *Signal) UnmarshalJSONFFLexer(fs *fflib.FFLexer, state fflib.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_Signalbase
	_ = currentKey
	tok := fflib.FFTok_init
	wantedTok := fflib.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == fflib.FFTok_error {
			goto tokerror
		}

		switch state {

		case fflib.FFParse_map_start:
			if tok != fflib.FFTok_left_bracket {
				wantedTok = fflib.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_key
			continue

		case fflib.FFParse_after_value:
			if tok == fflib.FFTok_comma {
				state = fflib.FFParse_want_key
			} else if tok == fflib.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = fflib.FFTok_comma
				goto wrongtokenerror
			}

		case fflib.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == fflib.FFTok_right_bracket {
				goto done
			}
			if tok != fflib.FFTok_string {
				wantedTok = fflib.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()
			if len(kn) <= 0 {
				// "" case. hrm.
				currentKey = ffj_t_Signalno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			} else {
				switch kn[0] {

				}

				currentKey = ffj_t_Signalno_such_key
				state = fflib.FFParse_want_colon
				goto mainparse
			}

		case fflib.FFParse_want_colon:
			if tok != fflib.FFTok_colon {
				wantedTok = fflib.FFTok_colon
				goto wrongtokenerror
			}
			state = fflib.FFParse_want_value
			continue
		case fflib.FFParse_want_value:

			if tok == fflib.FFTok_left_brace || tok == fflib.FFTok_left_bracket || tok == fflib.FFTok_integer || tok == fflib.FFTok_double || tok == fflib.FFTok_string || tok == fflib.FFTok_bool || tok == fflib.FFTok_null {
				switch currentKey {

				case ffj_t_Signalno_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = fflib.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.WrapErr(fs.BigError)
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}
