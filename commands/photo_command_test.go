package commands

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunPhoto(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &PhotoCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("a-know photo", " ")

	if status := target.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}
}
