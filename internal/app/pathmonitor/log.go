package pathmonitor

import (
	"log"
	"os"
)

var (
	Trace   *io.Writer
	Info    *io.Writer
	Warning *io.Writer
	Error   *io.Writer
)

func LogInit() {
	if Trace == nil {
		traceHandle := os.Stdout
		infoHandle := os.Stdout
		warningHandle := os.Stdout
		errorHandle := os.Stderr

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
}
