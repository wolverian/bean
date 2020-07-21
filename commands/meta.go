package commands

import (
	"fmt"
	"os"
)

type Meta struct{}

func (m *Meta) FatalError(err error) int {
	fmt.Fprintf(os.Stderr, "error: %s", err)
	return 1
}
