package main

type Redirect struct {
	Enabled bool
	File    string
}

func extractRedirect(parts []string) ([]string, Redirect) {
	var redirect Redirect

	for i := 0; i < len(parts); i++ {
		if parts[i] == ">" || parts[i] == "1>" {

			if i+1 < len(parts) {
				redirect.Enabled = true
				redirect.File = parts[i+1]
			}

			return parts[:i], redirect
		}
	}

	return parts, redirect
}
