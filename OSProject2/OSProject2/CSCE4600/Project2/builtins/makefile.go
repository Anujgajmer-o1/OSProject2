// builtins/makefile.go
package builtins

import (
	"os"
)

// MakeFile is a shell builtin to create an empty file with the specified name.
func MakeFile(args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: makefile <filename>")
	}

	filename := args[0]
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return nil
}
