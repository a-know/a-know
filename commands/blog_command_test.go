package commands

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunBlog(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &BlogCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("a-know blog", " ")

	if status := target.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}
}
