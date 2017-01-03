package main

import (
	"bytes"
	"fmt"
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

	expected := fmt.Sprint("http://blog.a-know.me")
	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestSynopsisBlog(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &BlogCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Display a-know's blog URL")
	if !strings.Contains(target.Synopsis(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestHelpBlog(t *testing.T) {
	outStream := new(bytes.Buffer)
	errStream := new(bytes.Buffer)
	target := &BlogCommand{OutStream: outStream, ErrStream: errStream}

	expected := fmt.Sprint("Usage: a-know blog")
	if !strings.Contains(target.Help(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
