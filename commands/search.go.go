package commands

import (
	"bean/brew"
	"bean/file"
	"fmt"
	"github.com/mitchellh/cli"
	"os"
	"strings"
)

type SearchCommand struct {
	fs []brew.Formula
}

func (sc SearchCommand) Help() string {
	return strings.TrimSpace(`
Usage: bean search <string>

	Searches formulae names for the given string.
`)
}

func (sc SearchCommand) Run(args []string) int {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "error: not enough arguments")
		return 1
	}
	for _, f := range sc.fs {
		if strings.Contains(f.Name, args[0]) {
			fmt.Println(f.Name)
		}
	}
	return 0
}

func (sc SearchCommand) Synopsis() string {
	return "search formulae"
}

func SearchFactory() (cli.Command, error) {
	fs, err := file.ReadFormulae()
	if err != nil {
		return nil, err
	}
	return SearchCommand{
		fs: fs,
	}, nil
}
