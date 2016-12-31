package commands

import (
	"fmt"
	"io"
)

// HomepageCommand is CLI struct for `homepage` sub-command
type HomepageCommand struct {
	OutStream, ErrStream io.Writer
}

// Synopsis is short-message for this command
func (c *HomepageCommand) Synopsis() string {
	return "Display a-know's homepage (portfolio page) URL"
}

// Help is content of help for this command
func (c *HomepageCommand) Help() string {
	return "Usage: a-know homepage"
}

// Run is main method of this command
func (c *HomepageCommand) Run(args []string) int {
	fmt.Fprintln(c.OutStream, "https://a-know.me/")
	return 0
}
