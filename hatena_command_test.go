package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunHatena(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &HatenaCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("a-know hatena", " ")

	if status := target.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}
}
