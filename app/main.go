package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt"
var _ = fmt.Print

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$ ")
		command, _ := reader.ReadString('\n')

		// slice mode
		commands := []string{
			"exit",
			"echo",
			"type",
		}

		// 1. if i want to exit the shell
		command = strings.TrimSpace(command)
		if command == "exit" {
			break
		}

		// after trim the string
		builtinToCheck := command[5:]

		// 2.  if i want echo to my command, method hasPrefix helps to read before the argument passed
		if strings.HasPrefix(command, "echo ") {
			fmt.Println(builtinToCheck)
		} else if strings.HasPrefix(command, "type ") {
			// 3. Work with typeee
			found := false
			for _, cmd := range commands {
				if builtinToCheck == cmd {
					found = true
				} else {
					found = false
				}

				if found == true {
					fmt.Println(command + " is a shell builtin")
				} else {
					fmt.Println(command + ": command not found")
				}
			}
		} else {
			fmt.Println(command + ": command not found")
		}

	}
}
