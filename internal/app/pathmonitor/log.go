package pathmonitor

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"log/syslog"
	"os"
	"strings"
)

const (
	TAG_PATHMONITOR = "pathmonitor"
	PREFIX_FILE     = "file://"
)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func LogInit(logInfo Log) {
	var traceHandle, infoHandle, warningHandle, errorHandle io.Writer
	var err error
	o := logInfo.Output

	// setup outputs
	switch {
	case o == "syslog":
		traceHandle, err = syslog.New(syslog.LOG_DEBUG, TAG_PATHMONITOR)
		infoHandle, err = syslog.New(syslog.LOG_INFO, TAG_PATHMONITOR)
		warningHandle, err = syslog.New(syslog.LOG_WARNING, TAG_PATHMONITOR)
		errorHandle, err = syslog.New(syslog.LOG_ERR, TAG_PATHMONITOR)
	case strings.HasPrefix(o, "file://"):
		fname := o[len(PREFIX_FILE):]
		f, err := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Failed opening logfile:", err)
		}
		traceHandle, infoHandle, warningHandle, errorHandle = f, f, f, f
	default:
		traceHandle = os.Stdout
		infoHandle = os.Stdout
		warningHandle = os.Stdout
		errorHandle = os.Stderr
	}

	_ = err

	// adjust log levels
	switch logInfo.Level {
	case "error":
		traceHandle = ioutil.Discard
		infoHandle = ioutil.Discard
		warningHandle = ioutil.Discard
	case "warning":
		traceHandle = ioutil.Discard
		infoHandle = ioutil.Discard
	case "info":
		traceHandle = ioutil.Discard
	}

	// prepare actual loggers
	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile|log.Lshortfile)
}
