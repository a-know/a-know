package commands

import (
	"fmt"
	"io"
)

// PhotoCommand is CLI struct for `photo` sub-command
type PhotoCommand struct {
	OutStream, ErrStream io.Writer
}

// Synopsis is short-message for this command
func (c *PhotoCommand) Synopsis() string {
	return "Display a-know's photograph blog URL"
}

// Help is content of help for this command
func (c *PhotoCommand) Help() string {
	return "Usage: a-know photo"
}

// Run is main method of this command
func (c *PhotoCommand) Run(args []string) int {
	fmt.Fprintln(c.OutStream, "http://photos.a-know.me/")
	return 0
}
