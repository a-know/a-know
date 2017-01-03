package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunGithub(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &GithubCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("a-know github", " ")

	if status := target.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}
}
