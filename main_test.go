package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestRunSubCommandTwitter(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("a-know twitter", " ")

	cli.Run(args)
	expected := fmt.Sprint("a-know's Twitter account : @a_know")

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestRunSubCommandBlog(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("a-know blog", " ")

	cli.Run(args)
	expected := fmt.Sprint("http://blog.a-know.me")

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
