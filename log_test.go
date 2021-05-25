package log

import (
	"testing"
)

func Test_Log(t *testing.T) {
	Info("this is info log")
	Infof("%s", "this is infof log")
	Debug("this is debug log")
	Debugf("%s", "this is debugf log")
	Error("this error log")
	Errorf("%s", "this is errorf log")
	Warn("this is warn log")
	Warnf("%s", "this is warnf log")
	Fatal("this is fatal log")
	Fatalf("%s", "this is fatalf log")
}
