package main

import (
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
	fmt.Fprintln(c.OutStream, "@a_know")
	return 0
}
