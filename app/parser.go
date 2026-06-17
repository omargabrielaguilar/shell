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

		if escaping {
			current.WriteByte(ch)
			escaping = false
			continue
		}

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
