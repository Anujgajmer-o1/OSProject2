// builtins/fileinfo.go
package builtins

import (
	"fmt"
	"os"
)

// FileInfo is a shell builtin to display detailed information about a file.
func FileInfo(args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: fileinfo <filename>")
	}

	filename := args[0]
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return err
	}

	fmt.Println("File Information:")
	fmt.Println("Name:", fileInfo.Name())
	fmt.Println("Size:", fileInfo.Size(), "bytes")
	fmt.Println("Mode:", fileInfo.Mode())
	fmt.Println("Last Modified:", fileInfo.ModTime())

	return nil
}
