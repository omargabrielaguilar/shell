package main

import (
	"fmt"
	"io"
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

func handleBuiltin(parts []string, redirect Redirect) bool {
	var out io.Writer = os.Stdout
	if redirect.Enabled {
		file, err := os.Create(redirect.File)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
			return true
		}
		defer file.Close()
		out = file
	}

	switch parts[0] {
	case "exit":
		os.Exit(0)
	case "echo":
		fmt.Fprintln(out, strings.Join(parts[1:], " "))
		return true
	case "pwd":
		cwd, err := os.Getwd()
		if err == nil {
			fmt.Fprintln(out, cwd)
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
			// Los errores de cd van a stderr, NUNCA se redirigen
			fmt.Fprintf(os.Stderr, "cd: %s: No such file or directory\n", dir)
		}
		return true
	case "type":
		if len(parts) < 2 {
			return true
		}
		arg := parts[1]
		if slices.Contains(builtins, arg) {
			fmt.Fprintln(out, arg+" is a shell builtin")
			return true
		}
		path, err := exec.LookPath(arg)
		if err == nil {
			fmt.Fprintln(out, arg+" is "+path)
		} else {
			// ✅ Los errores van a stderr
			fmt.Fprintf(os.Stderr, "%s: not found\n", arg)
		}
		return true
	}
	return false
}
