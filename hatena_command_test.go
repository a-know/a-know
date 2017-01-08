package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRunHatena(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &HatenaCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("", " ")

	if status := target.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := fmt.Sprint("http://profile.hatena.ne.jp/a-know/")
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestSynopsisHatena(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &HatenaCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Display a-know's Hatena profile page URL")
	if !strings.Contains(target.Synopsis(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestHelpHatena(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &HatenaCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Usage: a-know hatena")
	if !strings.Contains(target.Help(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
