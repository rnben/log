package log

import (
	"testing"
)

func TestGetLogger(t *testing.T) {
	defer Close()

	Debug("debug")
	Info("info")
	Warn("warn")
	Error("error")
}
