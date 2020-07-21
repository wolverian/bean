package commands

import (
	"fmt"
	"strings"

	"github.com/wolverian/bean/file"
)

type SearchCommand struct{ Meta }

func (sc *SearchCommand) Run(args []string) int {
	if len(args) != 1 {
		return sc.FatalError(fmt.Errorf("not enough arguments"))
	}
	fs, err := file.ReadFormulae()
	if err != nil {
		return sc.FatalError(err)
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
