package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"slices"
	"strings"
)

func parseCommand(input string) []string {
	var args []string
	var current strings.Builder

	inSingleQuote := false
	inDoubleQuote := false

	for _, ch := range input {
		switch {

		case ch == '\'' && !inDoubleQuote:
			inSingleQuote = !inSingleQuote

		case ch == '"' && !inSingleQuote:
			inDoubleQuote = !inDoubleQuote

		case (ch == ' ' || ch == '\t') &&
			!inSingleQuote &&
			!inDoubleQuote:

			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}

		default:
			current.WriteRune(ch)
		}
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	return args
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	// list of command accept
	commands := []string{
		"exit",
		"echo",
		"type",
		"pwd", // ya que se va a usar como commando aceptado para las pruebas de impresion de rutas
		"cd",
	}

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

		// exit
		if parts[0] == "exit" {
			break
		}

		// echo
		if parts[0] == "echo" {
			fmt.Println(strings.Join(parts[1:], " "))
			continue
		}

		// pwd
		if parts[0] == "pwd" {
			cwd, err := os.Getwd()
			if err != nil {
				fmt.Println(err)
				continue
			}

			fmt.Println(cwd)
			continue
		}

		// cd
		if parts[0] == "cd" {
			if len(parts) < 2 {
				continue
			}

			dir := parts[1]

			if dir == "~" {
				dir = os.Getenv("HOME")
			}

			if err := os.Chdir(dir); err != nil {
				fmt.Printf("cd: %s: No such file or directory\n", dir)
			}

			continue
		}

		// type
		if parts[0] == "type" {
			if len(parts) < 2 {
				continue
			}

			arg := parts[1]

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

		// External commands
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
