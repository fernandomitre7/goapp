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

func init() {
	fmt.Println("Logger init")
	var err error
	if _logFile, err = os.Create("logs/cryptochecker.log"); err != nil {
		log.Fatalln("Fail to open log file")
	}

	log.SetOutput(_logFile)

	_info = log.New(_logFile, "[INFO]: ", log.LUTC|log.Ldate|log.Ltime|log.Lshortfile)
	_debug = log.New(_logFile, "[DEBUG]: ", log.LUTC|log.Ldate|log.Ltime|log.Lshortfile)
	_warning = log.New(_logFile, "[WARN]: ", log.LUTC|log.Ldate|log.Ltime|log.Lshortfile)
	_error = log.New(_logFile, "[ERROR]: ", log.LUTC|log.Ldate|log.Ltime|log.Lshortfile)
}

// Close is meant to be on programs shutdown to properly close log file used
func Close() {
	_logFile.Close()
}

func Debug(format string, v ...interface{}) {
	s := formatLog(format, v)
	_debug.Output(2, s)
}

func Info(format string, v ...interface{}) {
	s := formatLog(format, v)
	_info.Output(2, s)
}

func Warn(format string, v ...interface{}) {
	s := formatLog(format, v)
	_warning.Output(2, s)
}

func Error(format string, v ...interface{}) {
	s := formatLog(format, v)
	_error.Output(2, s)
}

func formatLog(format string, v ...interface{}) string {
	if len(v) > 0 {
		return fmt.Sprintf(format, v...)
	}
	return format
}
