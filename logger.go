package logger

import (
	"log"
	"strings"
)

type LogLevel int

const (
	Debug LogLevel = iota
	Info
	Warning
	Error
	Panic
)

func (l LogLevel) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warning:
		return "WARNING"
	case Error:
		return "ERROR"
	case Panic:
		return "Panic"
	default:
		return "UNKNOWN"
	}
}

var colorReset = "\033[0m"

var colorCodes = map[LogLevel]string{
	Debug:   "\033[36m", // cyan for Debug
	Info:    "\033[32m", // green for Info (unchanged)
	Warning: "\033[33m", // yellow for Warning (unchanged)
	Error:   "\033[31m", // red for Error (unchanged)
	Panic:   "\033[35m", // magenta for Panic
}

func Log(l LogLevel, msg ...string) {
	if l == Panic {
		log.Fatalf("%s[%s]%s %s\n", colorCodes[l], l.String(), colorReset, strings.Join(msg, " "))
	}
	log.Printf("%s[%s]%s %s\n", colorCodes[l], l.String(), colorReset, strings.Join(msg, " "))
}
