package sdl

import (
	"testing"
)

func TestLog(t *testing.T) {
	Log("log_test: TestLog: %s", "this is a message")
}

func TestLogSetOutputFunction(t *testing.T) {
	LogSetOutputFunction(func(d interface{}, c int, pri LogPriority, message string) {
		println("CUSTOM:", message)
	}, nil)
	Log("log_test: TestLogSetOutputFunction: %s", "this is a message")
}
