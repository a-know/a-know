package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

// MentionCommand is CLI struct for `mention` sub-command
type MentionCommand struct {
	OutStream, ErrStream io.Writer
}

// Synopsis is short-message for this command
func (c *MentionCommand) Synopsis() string {
	return "You can send a message to a-know's slack channel."
}

// Help is content of help for this command
func (c *MentionCommand) Help() string {
	return "Usage: a-know mention <message>"
}

// Run is main method of this command
func (c *MentionCommand) Run(args []string) int {
	flags := flag.NewFlagSet("mention", flag.ContinueOnError)

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
	replyToURL := os.Getenv("A_KNOW_REPLY_TO_WEBHOOK")
	replyToChannel := os.Getenv("A_KNOW_REPLY_TO_CHANNEL")

	if replyToURL != "" {
		r := regexp.MustCompile(`hooks.slack.com/services`)
		if !r.MatchString(replyToURL) {
			fmt.Fprintln(c.OutStream, "error : Only slack's incoming Webhook URL can be specified as a environment variable `A_KNOW_REPLY_TO` .")
			return 1
		}

		if replyToChannel == "" {
			fmt.Fprintln(c.OutStream, "error : If you need reply from a-know, A_KNOW_REPLY_TO_WEBHOOK and A_KNOW_REPLY_TO_CHANNEL are both required.")
			return 1
		}

		message = fmt.Sprintf("%s\n(Reply to : %s @ %s )", message, replyToURL, replyToChannel)
	}

	jsonStr := `{"channel":"` + channel + `","username":"` + name + `","text":"` + message + `"}`

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer([]byte(jsonStr)))
	if err != nil {
		fmt.Fprintln(c.OutStream, "error occurs sending message. %s : %v", resp.Body, err)
		return 1
	}

	fmt.Fprintln(c.OutStream, "success to send message : ", message)
	return 0
}
