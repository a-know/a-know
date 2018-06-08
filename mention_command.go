package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
)

// MentionCommand is CLI struct for `mention` sub-command
type MentionCommand struct {
	OutStream, ErrStream io.Writer
}

// Synopsis is short-message for this command
func (c *MentionCommand) Synopsis() string {
	return "Display a-know's Mention account name"
}

// Help is content of help for this command
func (c *MentionCommand) Help() string {
	return "Usage: a-know mention <message>"
}

// Run is main method of this command
func (c *MentionCommand) Run(args []string) int {
	flags := flag.NewFlagSet("twitter", flag.ContinueOnError)

	if err := flags.Parse(args); err != nil {
		return 1
	}

	if len(args) < 1 {
		fmt.Fprintln(c.OutStream, "error : sending message is required.")
		return 1
	}

	name := "from a-know CLI"
	webhookURL := "https://hooks.slack.com/services/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"
	channel := "a-know-cli"
	message := args[0]

	jsonStr := `{"channel":"` + channel + `","username":"` + name + `","text":"` + message + `"}`

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Fprintln(c.OutStream, "error occurs sending message. %s : %v", resp.Body, err)
		return 1
	}

	fmt.Fprintln(c.OutStream, "success to send message : ", message)
	return 0
}
