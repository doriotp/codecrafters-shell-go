package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
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

		userInput = strings.TrimSpace(userInput)

		if userInput == "exit 0" {
			break
		} else if strings.HasPrefix(userInput, "echo") {
			tmp := userInput[4:]
			fmt.Println(strings.TrimSpace(tmp))
		} else if strings.HasPrefix(userInput, "type") {
			tmp := strings.TrimSpace(userInput[4:])
			handleTypeCommand(tmp)
		} else {
			fmt.Printf("%s: command not found\n", userInput)
		}

		fmt.Fprint(os.Stdout, "$ ")

	}
	// a better approach here would be
	// fmt.Printf("%s: command not found\n", strings.TrimSpace(userInput))
}

func checkIfExecInPath(e string) (string, error) {
	path, err := exec.LookPath(e)
	return path, err
}

func handleTypeCommand(tmp string) {
	if tmp == "echo" || tmp == "exit" || tmp == "type" {
		fmt.Printf("%s is a shell builtin\n", tmp)
	} else {
		path, err := checkIfExecInPath(tmp)
		if err != nil {
			fmt.Printf("%s: not found\n", tmp)
		} else {
			fmt.Printf("%s is %s\n", tmp, path)
		}
	}
}
