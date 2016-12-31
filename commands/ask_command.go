package commands

import (
	"fmt"
	"io"
)

// AskCommand is CLI struct for `ask` sub-command
type AskCommand struct {
	OutStream, ErrStream io.Writer
}

// Synopsis is short-message for this command
func (c *AskCommand) Synopsis() string {
	return "Display a-know's ask.fm page URL"
}

// Help is content of help for this command
func (c *AskCommand) Help() string {
	return "Usage: a-know ask"
}

// Run is main method of this command
func (c *AskCommand) Run(args []string) int {
	fmt.Fprintln(c.OutStream, "http://ask.fm/a_know")
	return 0
}
