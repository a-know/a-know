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
	expected := fmt.Sprint("@a_know")

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

func TestRunSubCommandHomepage(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("a-know homepage", " ")

	cli.Run(args)
	expected := fmt.Sprint("https://a-know.me/")

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestRunSubCommandHomepageWithAdminFlag(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("a-know homepage --admin", " ")

	cli.Run(args)
	expected := fmt.Sprint("https://a-know.me/?admin=true")

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestRunSubCommandGithub(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("a-know github", " ")

	cli.Run(args)
	expected := fmt.Sprint("https://github.com/a-know")

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestRunSubCommandPresentation(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("a-know presentation", " ")

	cli.Run(args)
	expected := fmt.Sprint("http://www.slideshare.net/aknow3373")

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestRunSubCommandPhoto(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("a-know photo", " ")

	cli.Run(args)
	expected := fmt.Sprint("http://photos.a-know.me/")

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestRunSubCommandHatena(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("a-know hatena", " ")

	cli.Run(args)
	expected := fmt.Sprint("http://profile.hatena.ne.jp/a-know/")

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}

func TestRunSubCommandAsk(t *testing.T) {
	outStream, errStream := new(bytes.Buffer), new(bytes.Buffer)
	cli := &CLI{outStream: outStream, errStream: errStream}
	args := strings.Split("a-know ask", " ")

	cli.Run(args)
	expected := fmt.Sprint("http://ask.fm/a_know")

	if !strings.Contains(outStream.String(), expected) {
		t.Errorf("expected %q to eq %q", outStream.String(), expected)
	}
}
