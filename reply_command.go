package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
)

// ReplyCommand is CLI struct for `reply` sub-command
type ReplyCommand struct {
	OutStream, ErrStream io.Writer
}

// Synopsis is short-message for this command
func (c *ReplyCommand) Synopsis() string {
	return "send a message to an arbitrary slack incoming webhook URL, slack channel."
}

// Help is content of help for this command
func (c *ReplyCommand) Help() string {
	return "Usage: a-know reply --msg <message> --url <reply webhook URL> --ch <reply channel name>"
}

// Run is main method of this command
func (c *ReplyCommand) Run(args []string) int {
	var (
		message        string
		replyToURL     string
		replyToChannel string
	)

	flags := flag.NewFlagSet("reply", flag.ContinueOnError)
	flags.StringVar(&message, "msg", "", "Reply message content")
	flags.StringVar(&replyToURL, "url", "", "Reply-to slack incoming webhook URL")
	flags.StringVar(&replyToChannel, "ch", "", "Reply-to slack channel name")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	if message == "" {
		fmt.Fprintln(c.OutStream, "error : message is empty.")
		return 1
	}

	if replyToURL == "" {
		fmt.Fprintln(c.OutStream, "error : Reply-To URL is empty.")
		return 1
	}

	if replyToChannel == "" {
		fmt.Fprintln(c.OutStream, "error : Reply-To channel name is empty.")
		return 1
	}

	name := "from a-know CLI"
	webhookURL := replyToURL
	message = fmt.Sprintf("%s (%s)", message, name)

	jsonStr := `{"channel":"` + replyToChannel + `","username":"` + name + `","text":"` + message + `"}`

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Fprintln(c.OutStream, "error occurs sending message. %s : %v", resp.Body, err)
		return 1
	}

	fmt.Fprintln(c.OutStream, "success to send message : ", message)
	return 0
}
