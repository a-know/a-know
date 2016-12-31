package commands

import (
	"fmt"
	"io"
)

// GithubCommand is CLI struct for `github` sub-command
type GithubCommand struct {
	OutStream, ErrStream io.Writer
}

// Synopsis is short-message for this command
func (c *GithubCommand) Synopsis() string {
	return "Display a-know's GitHub Account page"
}

// Help is content of help for this command
func (c *GithubCommand) Help() string {
	return "Usage: a-know github"
}

// Run is main method of this command
func (c *GithubCommand) Run(args []string) int {
	fmt.Fprintln(c.OutStream, "https://github.com/a-know")
	return 0
}
