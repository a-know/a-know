package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunTwitter(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &TwitterCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("a-know twitter", " ")

	if status := target.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}
}
