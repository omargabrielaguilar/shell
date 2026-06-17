package main

import "strings"

func parseCommand(input string) []string {
	var args []string
	var current strings.Builder

	inSingleQuote := false
	inDoubleQuote := false
	escaping := false

	for i := 0; i < len(input); i++ {

		ch := input[i]

		// Escapes fuera de quotes
		if escaping {
			current.WriteByte(ch)
			escaping = false
			continue
		}

		// Escapes dentro de double quotes
		if inDoubleQuote && ch == '\\' {

			if i+1 < len(input) {

				next := input[i+1]

				// Para este stage solo:
				// \" -> "
				// \\ -> \
				if next == '"' || next == '\\' {
					current.WriteByte(next)
					i++
					continue
				}
			}

			// Para cualquier otro caracter
			// el backslash es literal
			current.WriteByte(ch)
			continue
		}

		// Escapes fuera de quotes
		if ch == '\\' &&
			!inSingleQuote &&
			!inDoubleQuote {

			escaping = true
			continue
		}

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
			current.WriteByte(ch)
		}
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	return args
}
