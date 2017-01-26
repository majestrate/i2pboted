package log

import (
	"fmt"
	native "log"
)

type logLevel int

const (
	debug = logLevel(0)
	info  = logLevel(1)
	warn  = logLevel(2)
	error = logLevel(3)
	fatal = logLevel(4)
)

func (l logLevel) Int() int {
	return int(l)
}

func (l logLevel) Name() string {

	switch l {
	case debug:
		return "DBG"
	case info:
		return "NFO"
	case warn:
		return "WRN"
	case error:
		return "ERR"
	case fatal:
		return "FTL"
	default:
		return "???"
	}

}

func (l logLevel) Color() string {
	switch l {
	case debug:
		return "[37m"
	case info:
		return "[36m"
	case warn:
		return "[33;1m"
	case fatal:
	case error:
	default:
		return "[31;1m"
	}
	return ""
}

var level = debug

func accept(lvl logLevel) bool {
	return lvl.Int() >= level.Int()
}

func log(lvl logLevel, f string, args ...interface{}) {
	if accept(lvl) {
		m := fmt.Sprintf(f, args...)
		native.Printf("\x1b%s[%s] %s", lvl.Color(), lvl.Name(), m)
		if lvl == fatal {
			panic(m)
		}
	}
}

func Debug(msg string) {
	log(debug, msg)
}

func Debugf(f string, args ...interface{}) {
	log(debug, f, args...)
}

func Info(msg string) {
	log(info, msg)
}

func Infof(f string, args ...interface{}) {
	log(info, f, args...)
}

func Warn(msg string) {
	log(warn, msg)
}

func Warnf(f string, args ...interface{}) {
	log(warn, f, args...)
}

func Error(msg string) {
	log(error, msg)
}

func Errorf(f string, args ...interface{}) {
	log(error, f, args...)
}

func Fatal(msg string) {
	log(fatal, msg)
}

func Fatalf(f string, args ...interface{}) {
	log(fatal, f, args...)
}
