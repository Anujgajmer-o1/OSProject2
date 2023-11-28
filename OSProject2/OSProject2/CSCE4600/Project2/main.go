package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"github.com/Anujgajmer-o1/OSProject2/Project2/builtins"
)

func main() {
	exit := make(chan struct{}, 2) // buffer this so there's no deadlock.
	runLoop(os.Stdin, os.Stdout, os.Stderr, exit)
}

func runLoop(r io.Reader, w, errW io.Writer, exit chan struct{}) {
	var (
		input    string
		err      error
		readLoop = bufio.NewReader(r)
	)
	for {
		select {
		case <-exit:
			_, _ = fmt.Fprintln(w, "exiting gracefully...")
			return
		default:
			if err := printPrompt(w); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if input, err = readLoop.ReadString('\n'); err != nil {
				_, _ = fmt.Fprintln(errW, err)
				continue
			}
			if err = handleInput(w, input, exit); err != nil {
				_, _ = fmt.Fprintln(errW, err)
			}
		}
	}
}

func printPrompt(w io.Writer) error {
	// Get current user.
	// Don't prematurely memoize this because it might change due to `su`?
	u, err := user.Current()
	if err != nil {
		return err
	}
	// Get current working directory.
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	package main

	import (
		"fmt"
		"os"
		"strings"
		"github.com/Anujgajmer-o1/OSProject2/Project2/builtin"
	)
	
	func main() {
		args := os.Args
	
		if len(args) < 2 {
			fmt.Println("Usage: myshell <command> [args...]")
			os.Exit(1)
		}
	
		command := args[1]
		commandArgs := args[1:]
	
		switch command {
		case "echo":
			builtins.Echo(commandArgs)
		case "pwd":
			builtins.Pwd()
		case "mkdir":
			builtins.Mkdir(commandArgs)
		case "cp":
			builtins.Cp(commandArgs)
		case "rm":
			builtins.Rm(commandArgs)
		default:
			fmt.Println("Unknown command:", command)
		}
	}
	
	// /home/User [Username] $
	_, err = fmt.Fprintf(w, "%v [%v] $ ", wd, u.Username)

	return err
}

func handleInput(w io.Writer, input string, exit chan<- struct{}) error {
	// Remove trailing spaces.
	input = strings.TrimSpace(input)

	// Split the input separate the command name and the command arguments.
	args := strings.Split(input, " ")
	name, args := args[0], args[1:]

	// Check for built-in commands.
	// New builtin commands should be added here. Eventually this should be refactored to its own func.
	switch name {
	case "cd":
		return builtins.ChangeDirectory(args...)
	case "env":
		return builtins.EnvironmentVariables(w, args...)
	case "exit":
		exit <- struct{}{}
		return nil
	}

	return executeCommand(name, args...)
}
switch name {
case "cd":
    return builtins.ChangeDirectory(args...)
case "env":
    return builtins.EnvironmentVariables(w, args...)
case "exit":
    exit <- struct{}{}
    return nil
case "listdir":
    return builtins.ListDir(args...)
case "makefile":
    return builtins.MakeFile(args...)
case "countlines":
    return builtins.CountLines(args...)
case "renamefile":
    return builtins.RenameFile(args...)
case "fileinfo":
    return builtins.FileInfo(args...)
default:
    fmt.Println("Unknown command:", name)
    return nil
}


func executeCommand(name string, arg ...string) error {
	// Otherwise prep the command
	cmd := exec.Command(name, arg...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}
