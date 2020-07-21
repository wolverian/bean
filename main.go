package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"

	"github.com/wolverian/bean/command"
)

func main() {
	c := cli.NewCLI("bean", "0.1")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"search": func() (cli.Command, error) {
			return &command.SearchCommand{}, nil
		},
	}
	exitStatus, err := c.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
	}
	os.Exit(exitStatus)
}
