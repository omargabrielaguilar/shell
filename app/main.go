package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// list of command accept
	commands := []string{
		"exit",
		"echo",
		"type",
	}

	for {
		fmt.Print("$ ")

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "exit" {
			break
		}

		if strings.HasPrefix(command, "echo ") {
			fmt.Println(command[5:])
			continue
		}

		if strings.HasPrefix(command, "type ") {
			arg := strings.TrimSpace(command[5:])

			if slices.Contains(commands, arg) {
				fmt.Println(arg + " is a shell builtin")
				continue
			}

			path, err := exec.LookPath(arg)

			if err == nil {
				fmt.Println(arg + " is " + path)
			} else {
				fmt.Println(arg + ": not found")
			}

			continue
		}

		// ---------------------------
		// External commands
		// ---------------------------

		parts := strings.Fields(command)

		if len(parts) == 0 {
			continue
		}

		program := parts[0]
		// cambios permitir imprimir el programa, no la ruta completa
		if _, err := exec.LookPath(program); err != nil {
			fmt.Println(program + ": command not found")
			continue
		}

		cmd := exec.Command(program, parts[1:]...)

		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			fmt.Println(err)
		}
	}
}
