package commands

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunHomepage(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &HomepageCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("a-know homepage", " ")

	if status := target.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}
}
