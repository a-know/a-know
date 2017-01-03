package main

import (
	"bytes"
	"fmt"
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

	expected := fmt.Sprint("http://ask.fm/a_know")
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestSynopsisAsk(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &AskCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Display a-know's ask.fm page URL")
	if !strings.Contains(target.Synopsis(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestHelpAsk(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &AskCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Usage: a-know ask")
	if !strings.Contains(target.Help(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
