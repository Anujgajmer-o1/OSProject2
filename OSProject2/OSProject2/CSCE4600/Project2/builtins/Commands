// builtins/renamefile.go
package builtins

import (
	"fmt"
	"os"
)

// RenameFile is a shell builtin to rename or move a file to a different location.
func RenameFile(args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("Usage: renamefile <oldname> <newname>")
	}

	oldName := args[0]
	newName := args[1]

	err := os.Rename(oldName, newName)
	if err != nil {
		return err
	}

	return nil
}
