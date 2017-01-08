package main

import (
	"flag"
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
	var admin bool

	flags := flag.NewFlagSet("homepage", flag.ContinueOnError)
	flags.BoolVar(&admin, "admin", false, "Get URL for admin")

	if len(args) > 2 {
		if err := flags.Parse(args[2:]); err != nil {
			return 1
		}
	}

	homepageURL := "https://a-know.me/"

	if admin {
		homepageURL = homepageURL + "?admin=true"
	}

	fmt.Fprintln(c.OutStream, homepageURL)
	return 0
}
