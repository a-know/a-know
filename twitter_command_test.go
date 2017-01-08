package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRunTwitter(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &TwitterCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("", " ")

	if status := target.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := fmt.Sprint("@a_know")
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestRunTwitterWithUrlFlag(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &TwitterCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("--url", " ")

	if status := target.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := fmt.Sprint("https://twitter.com/a_know")
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestRunTwitterWithInvalidUrlFlag(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &TwitterCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("--urll", " ")

	if status := target.Run(args); status != ExitCodeError {
		t.Errorf("expected %d to eq %d", status, ExitCodeError)
	}
}

func TestSynopsisTwitter(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &TwitterCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Display a-know's Twitter account name")
	if !strings.Contains(target.Synopsis(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestHelpTwitter(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &TwitterCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Usage: a-know twitter")
	if !strings.Contains(target.Help(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
