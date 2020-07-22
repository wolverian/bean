package command

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Add struct{ Meta }

func (a *Add) Help() string {
	return strings.TrimSpace(`
Usage: bean add <formula>

	Installs a formula.
`)
}

func (a *Add) Run(args []string) int {
	name := args[0]
	cmd := exec.Command("brew", "install", name)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println(cmd)
	err := cmd.Run()
	if err != nil {
		return a.FatalError(err)
	}
	return 0
}

func (a *Add) Synopsis() string {
	return "install a formula"
}
