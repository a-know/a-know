package main

import (
	"fmt"
	"io"
)

// HatenaCommand is CLI struct for `hatena` sub-command
type HatenaCommand struct {
	OutStream, ErrStream io.Writer
}

// Synopsis is short-message for this command
func (c *HatenaCommand) Synopsis() string {
	return "Display a-know's Hatena profile page URL"
}

// Help is content of help for this command
func (c *HatenaCommand) Help() string {
	return "Usage: a-know hatena"
}

// Run is main method of this command
func (c *HatenaCommand) Run(args []string) int {
	fmt.Fprintln(c.OutStream, "http://profile.hatena.ne.jp/a-know/")
	return 0
}
