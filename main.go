package main

import (
	"bean/commands"
	"fmt"
	"github.com/mitchellh/cli"
	"os"
)

func main() {
	c := cli.NewCLI("bean", "0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"search": func() (cli.Command, error) {
			return &commands.SearchCommand{}, nil
		},
	}
	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
	}
	os.Exit(exitStatus)
}
