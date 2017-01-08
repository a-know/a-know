package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRunHomepage(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &HomepageCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("", " ")

	if status := target.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := fmt.Sprint("https://a-know.me/")
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestRunHomepageWithAdminFlag(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &HomepageCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("--admin", " ")

	if status := target.Run(args); status != ExitCodeOK {
		t.Errorf("expected %d to eq %d", status, ExitCodeOK)
	}

	expected := fmt.Sprint("https://a-know.me/?admin=true\n")
	if !strings.EqualFold(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestRunHomepageWithInvalidAdminFlag(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &HomepageCommand{OutStream: outStream, ErrStream: errStream}

	args := strings.Split("--admi", " ")

	if status := target.Run(args); status != ExitCodeError {
		t.Errorf("expected %d to eq %d", status, ExitCodeError)
	}
}

func TestSynopsisHomepage(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &HomepageCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Display a-know's homepage (portfolio page) URL")
	if !strings.Contains(target.Synopsis(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestHelpHomepage(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &HomepageCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Usage: a-know homepage")
	if !strings.Contains(target.Help(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
