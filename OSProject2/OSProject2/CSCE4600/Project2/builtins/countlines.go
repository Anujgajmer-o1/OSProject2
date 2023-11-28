// builtins/countlines.go
package builtins

import (
	"bufio"
	"fmt"
	"os"
)

// CountLines is a shell builtin to count the number of lines in a given file.
func CountLines(args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("Usage: countlines <filename>")
	}

	filename := args[0]
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}

	fmt.Printf("Number of lines in %s: %d\n", filename, lineCount)
	return nil
}
