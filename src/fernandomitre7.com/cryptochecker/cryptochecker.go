package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fernandomitre7.com/cryptochecker/config"
	"fernandomitre7.com/cryptochecker/logger"
)

var (
	configFile = flag.String("config", "/home/fernando/FERNANDO/cryptochecker/config/cryptochecker.conf.json", "Configuration File")
	// Used for shutdown operations
	killC      chan os.Signal
	interruptC chan os.Signal
)

func init() {
	killC = make(chan os.Signal, 1)
	interruptC = make(chan os.Signal, 1)
}

func main() {
	// Load Configurations
	_conf, err := config.LoadConfiguration(*configFile)
	if err != nil {
		panic(err)
	}
	logger.Debug("Crypto Checker started on port %v", _conf.Port)

	time.Sleep(time.Second * 1)

	signal.Notify(killC, syscall.SIGTERM)
	signal.Notify(interruptC, os.Interrupt)
	// go start()

	// This select will prevent program to stop until is forcelly shutdown
	select {
	case <-interruptC:
		logger.Info("INTERRUPT signal received! shutdown initiated...")
		//stop()
	case <-killC:
		logger.Info("SIGTERM signal received! shutdown initiated ...")
		//stop()
	}
	fmt.Println("\nStopping...")
	logger.Info("Crypto Checker stopped")
	logger.Close()
}
