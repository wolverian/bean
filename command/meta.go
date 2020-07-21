package command

import (
	"fmt"
	"os"
)

// Meta is the common options and features for all commands.
type Meta struct{}

// FatalError prints err to os.Stderr and returns a Unix exit code indicating an error.
func (m *Meta) FatalError(err error) int {
	fmt.Fprintf(os.Stderr, "error: %s\n", err)
	return 1
}
