package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Uncomment this block to pass the first stage
	fmt.Fprint(os.Stdout, "$ ")

	for {
		// Wait for user input
		userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Printf("failed to read user input due to error %s", err)
			return
		}

		if userInput == "exit 0" {
			return
		} else if strings.HasPrefix(userInput, "echo") {
			tmp := userInput[4:]
			fmt.Println(strings.TrimSpace(tmp))
		} else {
			fmt.Printf("%s: command not found\n", strings.TrimSpace(userInput))
		}

		fmt.Fprint(os.Stdout, "$ ")

	}
	// a better approach here would be
	// fmt.Printf("%s: command not found\n", strings.TrimSpace(userInput))
}
