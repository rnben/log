package log

import "fmt"

var l *Log

func init() {
	l = New()
}

func SetCallDepth(d int) {
	l.depth = d
}

func Info(args ...interface{}) {
	l.info.Output(l.depth, fmt.Sprintln(args...))
}

func Infof(format string, args ...interface{}) {
	l.info.Output(l.depth, fmt.Sprintf(format, args...))
}

func Warn(args ...interface{}) {
	l.warn.Output(l.depth, fmt.Sprintln(args...))
}

func Warnf(format string, args ...interface{}) {
	l.warn.Output(l.depth, fmt.Sprintf(format, args...))
}

func Error(args ...interface{}) {
	l.erro.Output(l.depth, fmt.Sprintln(args...))
}

func Errorf(format string, args ...interface{}) {
	l.erro.Output(l.depth, fmt.Sprintf(format, args...))
}

func Debug(args ...interface{}) {
	l.debu.Output(l.depth, fmt.Sprintln(args...))
}

func Debugf(format string, args ...interface{}) {
	l.debu.Output(l.depth, fmt.Sprintf(format, args...))
}

func Fatal(args ...interface{}) {
	l.fata.Output(l.depth, fmt.Sprintln(args...))
}

func Fatalf(format string, args ...interface{}) {
	l.fata.Output(l.depth, fmt.Sprintf(format, args...))
}
