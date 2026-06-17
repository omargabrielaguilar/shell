package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Print("$ ")

		command, _ := reader.ReadString('\n')
		command = strings.TrimSpace(command)

		if command == "" {
			continue
		}

		parts := parseCommand(command)

		if len(parts) == 0 {
			continue
		}

		parts, redirect := extractRedirect(parts)

		if len(parts) == 0 {
			continue
		}

		if handleBuiltin(parts, redirect) {
			continue
		}

		program := parts[0]

		if _, err := exec.LookPath(program); err != nil {
			fmt.Println(program + ": command not found")
			continue
		}

		cmd := exec.Command(program, parts[1:]...)

		cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr

		if redirect.Enabled {

			file, err := os.Create(redirect.File)
			if err != nil {
				fmt.Println(err)
				continue
			}

			defer file.Close()

			cmd.Stdout = file

		} else {
			cmd.Stdout = os.Stdout
		}

		_ = cmd.Run()
	}
}
