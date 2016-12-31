package commands

import (
	"fmt"
	"io"
)

// BlogCommand is CLI struct for `blog` sub-command
type BlogCommand struct {
	OutStream, ErrStream io.Writer
}

// Synopsis is short-message for this command
func (c *BlogCommand) Synopsis() string {
	return "Display a-know's blog URL"
}

// Help is content of help for this command
func (c *BlogCommand) Help() string {
	return "Usage: a-know blog"
}

// Run is main method of this command
func (c *BlogCommand) Run(args []string) int {
	fmt.Fprintln(c.OutStream, "http://blog.a-know.me")
	return 0
}
