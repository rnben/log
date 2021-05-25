package log

import (
	"fmt"
	"log"
	"os"
)



type Log struct {
	depth int
	info  *log.Logger
	erro  *log.Logger
	warn  *log.Logger
	debu  *log.Logger
	fata  *log.Logger
}

func (l *Log) SetCallDepth(d int) {
	l.depth = d
}

func (l *Log) Info(format string, s ...interface{}) {
	l.info.Output(l.depth, fmt.Sprintf(format, s...))
}

func (l *Log) Warn(format string, s ...interface{}) {
	l.warn.Output(l.depth, fmt.Sprintf(format, s...))
}

func (l *Log) Error(format string, s ...interface{}) {
	l.erro.Output(l.depth, fmt.Sprintf(format, s...))
}

func (l *Log) Debug(format string, s ...interface{}) {
	l.debu.Output(l.depth, fmt.Sprintf(format, s...))
}

func (l *Log) Fatal(format string, s ...interface{}) {
	l.fata.Output(l.depth, fmt.Sprintf(format, s...))
	os.Exit(1)
}

func New() *Log {
	l := &Log{
		depth: 2,
		info:  log.New(os.Stdout, green(info), log.Llongfile|log.Ltime|log.Ldate),
		erro:  log.New(os.Stdout, red(erro), log.Llongfile|log.Ltime|log.Ldate),
		warn:  log.New(os.Stdout, yellow(warn), log.Llongfile|log.Ltime|log.Ldate),
		debu:  log.New(os.Stdout, blue(debu), log.Llongfile|log.Ltime|log.Ldate),
		fata:  log.New(os.Stdout, magenta(fata), log.Llongfile|log.Ltime|log.Ldate),
	}

	return l
}
