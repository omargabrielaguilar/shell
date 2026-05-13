package main

import (
	"bufio"
	"fmt"
	"os"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	fmt.Print("$ ")
	command, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// TODO: Uncomment the code below to pass the first stage
	fmt.Println(command[:len(command)-1] + ": command not found")
}
