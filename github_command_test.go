package main

import (
	"bytes"
	"fmt"
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

	expected := fmt.Sprint("https://github.com/a-know")
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestSynopsisGithub(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &GithubCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Display a-know's GitHub Account page")
	if !strings.Contains(target.Synopsis(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestHelpGithub(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &GithubCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Usage: a-know github")
	if !strings.Contains(target.Help(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
