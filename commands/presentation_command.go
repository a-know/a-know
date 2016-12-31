package commands

import (
	"fmt"
	"io"
)

// PresentationCommand is CLI struct for `presentation` sub-command
type PresentationCommand struct {
	OutStream, ErrStream io.Writer
}

// Synopsis is short-message for this command
func (c *PresentationCommand) Synopsis() string {
	return "Display a-know's slideshare URL"
}

// Help is content of help for this command
func (c *PresentationCommand) Help() string {
	return "Usage: a-know presentation"
}

// Run is main method of this command
func (c *PresentationCommand) Run(args []string) int {
	fmt.Fprintln(c.OutStream, "http://www.slideshare.net/aknow3373")
	return 0
}
