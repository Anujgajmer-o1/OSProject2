// builtins/listdir.go
package builtins

import (
	"fmt"
	"io/ioutil"
)

// ListDir is a shell builtin to list files and directories in the current working directory.
func ListDir(args ...string) error {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	return nil
}
