package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type bLogger struct {
	*log.Logger
}

// Log define new log struct.
type Log struct {
	command string
	message string
	level   string
}

// Color macro for bash
var (
	errCode   = "\033[38;5;196m"
	infoCode  = "\033[38;5;266m"
	warnCode  = "\033[38;5;214m"
	flagsCode = "\033[38;5;8m"
	nameCode  = "\033[38;5;157m"
	msgCode   = "\033[38;5;243m"
	cCode     = "\033[38;5;6m"
	wCode     = "\033[38;5;2m"
	eCode     = "\033[38;5;7m"
	iCode     = "\033[38;5;7m"
	resetCode = "\033[0m"
	bl        = bLogger{log.New(os.Stdout, nameCode+"go-streamer"+flagsCode, 0)}
)

// Command set type of command.
func (l Log) Command(command, code string) Log {
	switch code {
	case "c":
		l.command = cCode + command + resetCode + " "
	case "i":
		l.command = iCode + command + resetCode + " "
	case "w":
		l.command = wCode + command + resetCode + " "
	case "e":
		l.command = eCode + command + resetCode + " "
	}
	return l
}

// Message sets logging message.
func (l Log) Message(message ...string) Log {
	l.message = strings.Join(message, " ")
	return l
}

// Error level.
func Error() Log {
	return Log{level: "error"}
}

// Info level.
func Info() Log {
	return Log{level: "info"}
}

// Warn level.
func Warn() Log {
	return Log{level: "warn"}
}

// Log construct log message and display.
func (l Log) Log() {
	out := "[" + time.Now().Format("15:04:05") + "]"
	if l.command != "" {
		out += l.command
	}
	switch l.level {
	case "error":
		fmt.Println("enter-------")
		out += errCode
	case "info":
		out += infoCode
	case "warn":
		out += warnCode
	}

	out += l.message + resetCode
	bl.Print(out)
}

// FormattedMessage format logging message.
func FormattedMessage(message string) string {
	return msgCode + "[" + message + "]"
}
