package commands

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunAsk(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &AskCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("a-know ask", " ")

	if status := target.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}
}
