package logger

import (
	"fmt"
	"log"
	"os"
)

var (
	//_TRACE   *log.Logger
	_info    *log.Logger
	_debug   *log.Logger
	_warning *log.Logger
	_error   *log.Logger
	_logFile *os.File
)

// Init initializes logger
func Init(logPath string) error {
	fmt.Println("Logger init logPath:", logPath)
	var err error
	if _logFile, err = os.Create(logPath); err != nil {
		return fmt.Errorf("Fail to open log file %v", err)
	}

	log.SetOutput(_logFile)

	_info = log.New(_logFile, "[INFO]: ", log.LUTC|log.Ldate|log.Ltime|log.Lshortfile)
	_debug = log.New(_logFile, "[DEBUG]: ", log.LUTC|log.Ldate|log.Ltime|log.Lshortfile)
	_warning = log.New(_logFile, "[WARN]: ", log.LUTC|log.Ldate|log.Ltime|log.Lshortfile)
	_error = log.New(_logFile, "[ERROR]: ", log.LUTC|log.Ldate|log.Ltime|log.Lshortfile)

	return nil
}

// Close is meant to be on programs shutdown to properly close log file used
func Close() {
	_logFile.Close()
}

func Debugf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	_debug.Output(2, s)
}

func Infof(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	_info.Output(2, s)
}

func Warnf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	_warning.Output(2, s)
}

func Errorf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	_error.Output(2, s)
}

func Debug(s string) {
	_debug.Output(2, s)
}

func Info(s string) {
	_info.Output(2, s)
}

func Warn(s string) {
	_warning.Output(2, s)
}

func Error(s string) {
	_error.Output(2, s)
}
