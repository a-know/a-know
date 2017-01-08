package main

import (
	"flag"
	"fmt"
	"io"
)

// TwitterCommand is CLI struct for `twitter` sub-command
type TwitterCommand struct {
	OutStream, ErrStream io.Writer
}

// Synopsis is short-message for this command
func (c *TwitterCommand) Synopsis() string {
	return "Display a-know's Twitter account name"
}

// Help is content of help for this command
func (c *TwitterCommand) Help() string {
	return "Usage: a-know twitter"
}

// Run is main method of this command
func (c *TwitterCommand) Run(args []string) int {
	var url bool

	flags := flag.NewFlagSet("twitter", flag.ContinueOnError)
	flags.BoolVar(&url, "url", false, "Get account page URL")

	if err := flags.Parse(args); err != nil {
		return 1
	}

	output := "@a_know"

	if url {
		output = "https://twitter.com/a_know"
	}

	fmt.Fprintln(c.OutStream, output)
	return 0
}
