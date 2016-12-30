package main

import (
	"bytes"
	"strings"
	"testing"
)

var outStream = new(bytes.Buffer)
var errStream = new(bytes.Buffer)
var target = &TwitterCommand{outStream: outStream, errStream: errStream}

func TestRun(t *testing.T) {
	args := strings.Split("a-know twitter", " ")

	if status := target.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}
}
