package main

import (
	"bytes"
	"fmt"
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

	expected := fmt.Sprint("http://www.slideshare.net/aknow3373")
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestSynopsisPresentation(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &PresentationCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Display a-know's slideshare URL")
	if !strings.Contains(target.Synopsis(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestHelpPresentation(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &PresentationCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Usage: a-know presentation")
	if !strings.Contains(target.Help(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
