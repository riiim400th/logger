package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
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

func Log(l LogLevel, v ...any) {
	var msgArgs []string
	for _, arg := range v {
		msgArgs = append(msgArgs, fmt.Sprint(arg))
	}
	inlogMsg := strings.Join(msgArgs, " ")
	logMsg := fmt.Sprintf("%s[%s]%s %s %s\n", colorCodes[l], l.String(), colorReset, time.Now().Format("2006-01-02 - 15:04:05"), inlogMsg)
	if l == Panic {
		fmt.Print(logMsg)
		os.Exit(1)
	}
	fmt.Print(logMsg)
}

func P(v ...any) {
	Log(Info, v...)
}

// func logmessage(msg ...any)string{
// 	return
// }
