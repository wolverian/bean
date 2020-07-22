package command

import (
	"fmt"
	"strings"

	"github.com/wolverian/bean/file"
)

// Search is a cli.Command implementation that takes an argument and looks if any available package matches that argument.
type Search struct{ Meta }

func (s *Search) Run(args []string) int {
	if len(args) != 1 {
		return s.FatalError(fmt.Errorf("not enough arguments"))
	}
	fs, err := file.FindAll(args[0], strings.Contains)
	if err != nil {
		return s.FatalError(fmt.Errorf("no such formula"))
	}
	for _, f := range fs {
		fmt.Println(f.Name)
	}
	return 0
}

func (s *Search) Help() string {
	return strings.TrimSpace(`
Usage: bean search <string>

	Searches formulae names for the given string.
`)
}

func (s *Search) Synopsis() string {
	return "search formulae"
}
