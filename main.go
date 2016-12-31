package main

import (
	"fmt"
	"io"
	"os"

	"github.com/a-know/a-know/commands"

	"github.com/mitchellh/cli"
)

// CLI is struct for switch stdout / stderr
type CLI struct {
	outStream, errStream io.Writer
}

func main() {
	cli := &CLI{outStream: os.Stdout, errStream: os.Stderr}
	os.Exit(cli.Run(os.Args))
}

// Run method is select correct command and run
func (c *CLI) Run(args []string) int {
	cl := cli.NewCLI("a-know", "0.0.1")
	cl.Args = args[1:]
	cl.Commands = map[string]cli.CommandFactory{
		"twitter": func() (cli.Command, error) {
			return &commands.TwitterCommand{OutStream: c.outStream, ErrStream: c.errStream}, nil
		},
		"blog": func() (cli.Command, error) {
			return &commands.BlogCommand{OutStream: c.outStream, ErrStream: c.errStream}, nil
		},
		"homepage": func() (cli.Command, error) {
			return &commands.HomepageCommand{OutStream: c.outStream, ErrStream: c.errStream}, nil
		},
	}

	exitStatus, err := cl.Run()
	if err != nil {
		fmt.Fprintf(c.errStream, "Failed to execute: %s\n", err.Error())
	}

	return exitStatus
}
