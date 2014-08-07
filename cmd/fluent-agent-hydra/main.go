package main

import (
	"flag"
	"fmt"
	"github.com/fujiwara/fluent-agent-hydra/hydra"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"
	"syscall"
)

var (
	trapSignals = []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT}
)

func main() {
	var (
		configFile  string
		help        bool
		fieldName   string
		monitorAddr string
	)
	flag.StringVar(&configFile, "c", "", "configuration file path")
	flag.BoolVar(&help, "h", false, "show help message")
	flag.BoolVar(&help, "help", false, "show help message")
	flag.StringVar(&fieldName, "f", hydra.DefaultFieldName, "fieldname of fluentd log message attribute (DEFAULT: message)")
	flag.StringVar(&monitorAddr, "monitor", "127.0.0.1:24223", "monitor HTTP server address")
	flag.Parse()

	if help {
		usage()
	}
	if pprofile := os.Getenv("PPROF"); pprofile != "" {
		f, err := os.Create(pprofile)
		if err != nil {
			log.Fatal("[error] Can't create profiling stat file.", err)
		}
		log.Println("[info] StartCPUProfile() stat file", f.Name())
		pprof.StartCPUProfile(f)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, trapSignals...)
	if configFile != "" {
		config, err := hydra.ReadConfig(configFile)
		if err != nil {
			log.Println("Can't load config", err)
			os.Exit(2)
		}
		runWithConfig(config)
	} else if args := flag.Args(); len(args) >= 3 {
		config := hydra.NewConfigByArgs(args, fieldName, monitorAddr)
		runWithConfig(config)
	} else {
		usage()
	}
	sig := <-done
	log.Println("[info] SIGNAL", sig, "exit.")
	pprof.StopCPUProfile()
	os.Exit(0)
}

func usage() {
	fmt.Println("Usage of fluent-agent-hydra")
	fmt.Println("")
	fmt.Println("  fluent-agent-hydra -c config.toml")
	fmt.Println("  fluent-agent-hydra [options] TAG TARGET_FILE PRIMARY_SERVER SECONDARY_SERVER")
	fmt.Println("")
	flag.PrintDefaults()
	os.Exit(1)
}

func runWithConfig(config *hydra.Config) {
	messageCh, monitorCh := hydra.NewChannel()

	// start monitor server
	_, err := hydra.MonitorProcess(config, monitorCh)
	if err != nil {
		log.Println("[error] Couldn't start monitor server.", err)
	}

	// start out_forward
	go hydra.OutForward(config.Servers, messageCh, monitorCh)

	// start in_tail
	for _, configLogfile := range config.Logs {
		go hydra.InTail(configLogfile, messageCh, monitorCh)
	}

	// start in_forward
	for _, configReceiver := range config.Receivers {
		hydra.InForward(configReceiver, messageCh, monitorCh)
	}
}
