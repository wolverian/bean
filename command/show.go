package command

import (
	"fmt"
	"strings"

	"github.com/wolverian/bean/file"
)

type Show struct{ Meta }

func (s *Show) Help() string {
	return strings.TrimSpace(`
Usage: bean show <name>

	Shows information for a formula.
`)
}

func (s *Show) Run(args []string) int {
	if len(args) != 1 {
		return s.FatalError(fmt.Errorf("not enough arguments"))
	}
	fs, err := file.ReadFormulae()
	if err != nil {
		return s.FatalError(err)
	}
	for _, f := range fs {
		if strings.EqualFold(f.Name, args[0]) {
			fmt.Println(f)
			return 0
		}
	}
	return s.FatalError(fmt.Errorf("formula not found"))
}

func (s *Show) Synopsis() string {
	return "show formula"
}
