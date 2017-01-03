package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunPresentation(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &PresentationCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("a-know presentation", " ")

	if status := target.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}
}
