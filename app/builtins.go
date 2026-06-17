package main

import (
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

var builtins = []string{
	"exit",
	"echo",
	"type",
	"pwd",
	"cd",
}

func handleBuiltin(parts []string) bool {
	switch parts[0] {

	case "exit":
		os.Exit(0)

	case "echo":
		fmt.Println(strings.Join(parts[1:], " "))
		return true

	case "pwd":
		cwd, err := os.Getwd()
		if err == nil {
			fmt.Println(cwd)
		}
		return true

	case "cd":
		if len(parts) < 2 {
			return true
		}

		dir := parts[1]

		if dir == "~" {
			dir = os.Getenv("HOME")
		}

		if err := os.Chdir(dir); err != nil {
			fmt.Printf("cd: %s: No such file or directory\n", dir)
		}

		return true

	case "type":
		if len(parts) < 2 {
			return true
		}

		arg := parts[1]

		if slices.Contains(builtins, arg) {
			fmt.Println(arg + " is a shell builtin")
			return true
		}

		path, err := exec.LookPath(arg)

		if err == nil {
			fmt.Println(arg + " is " + path)
		} else {
			fmt.Println(arg + ": not found")
		}

		return true
	}

	return false
}
