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

		if handleBuiltin(parts) {
			continue
		}

		program := parts[0]

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
