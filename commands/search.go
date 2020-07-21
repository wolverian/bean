package commands

import (
	"bean/file"
	"fmt"
	"os"
	"strings"
)

type SearchCommand struct{}

func (sc *SearchCommand) Run(args []string) int {
	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "error: not enough arguments")
		return 1
	}
	fs, err := file.ReadFormulae()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		return 1
	}
	for _, f := range fs {
		if strings.Contains(f.Name, args[0]) {
			fmt.Println(f.Name)
		}
	}
	return 0
}

func (sc *SearchCommand) Help() string {
	return strings.TrimSpace(`
Usage: bean search <string>

	Searches formulae names for the given string.
`)
}

func (sc *SearchCommand) Synopsis() string {
	return "search formulae"
}
