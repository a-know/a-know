package commands

import (
	"fmt"
	"io"
)

type TwitterCommand struct {
	OutStream, ErrStream io.Writer
}

func (c *TwitterCommand) Synopsis() string {
	return "Display a-know's Twitter account name"
}

func (c *TwitterCommand) Help() string {
	return "Usage: a-know twitter"
}

func (c *TwitterCommand) Run(args []string) int {
	fmt.Fprintln(c.OutStream, "a-know's Twitter account : @a_know")
	return 0
}
