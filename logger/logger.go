package logger

import (
	"log"
	"os"
	"strings"
	"time"
)

var (
	errorC   = "\033[38;5;196m"
	infoC    = "\033[38;5;266m"
	warnC    = "\033[38;5;214m"
	flagsC   = "\033[38;5;8m"
	nameC    = "\033[38;5;157m"
	messageC = "\033[38;5;243m"
	commandC = "\033[38;5;248m"
	resetC   = "\033[0m"
	vl       = vLogger{log.New(os.Stdout, nameC+"v-comp"+flagsC, 0)}
)

type vLogger struct {
	*log.Logger
}

// Log struct
type Log struct {
	Cmd string
	Msg string
	Lvl string
}

// Command set type of command.
func (l Log) Command(c string) Log {
	l.Cmd = c
	return l
}

// Message sets logging message.
func (l Log) Message(m ...string) Log {
	l.Msg = strings.Join(m, " ")
	return l
}

// Error level.
func Error() Log {
	return Log{Lvl: "error"}
}

// Warn level.
func Warn() Log {
	return Log{Lvl: "warn"}
}

// Info level.
func Info() Log {
	return Log{Lvl: "info"}
}

// Log construct log message and display.
func (l Log) Log() {

	out := "[" + time.Now().Format("15:04:05") + "]"
	if l.Cmd != "" {
		out += commandC + l.Cmd + resetC + ""
	}

	// log level
	switch l.Lvl {
	case "error":
		out += errorC
	case "warn":
		out += warnC
	case "info":
		out += infoC
	}

	// final msg
	out += l.Msg + resetC
	vl.Printf(out)
}

// FormattedMessage format logging message.
func FormattedMessage(msg string) string {
	return commandC + "'" + messageC + msg + commandC + "'"
}
