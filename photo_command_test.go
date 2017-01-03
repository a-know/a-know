package main

import (
	"bytes"
	"fmt"
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

	expected := fmt.Sprint("http://photos.a-know.me/")
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestSynopsisPhoto(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &PhotoCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Display a-know's photograph blog URL")
	if !strings.Contains(target.Synopsis(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestHelpPhoto(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &PhotoCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Usage: a-know photo")
	if !strings.Contains(target.Help(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
