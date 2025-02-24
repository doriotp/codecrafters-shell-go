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
	for {
		fmt.Fprint(os.Stdout, "$ ")
		// Wait for user input
		userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		cmds := strings.Split(userInput, " ")
		if err != nil {
			fmt.Printf("failed to read user input due to error %s", err)
			return
		}

		if userInput == "exit 0" {
			os.Exit(0)
		} else if strings.HasPrefix(userInput, "echo") {
			tmp := userInput[4:]
			fmt.Println(strings.TrimSpace(tmp))
		} else if strings.HasPrefix(userInput, "type") {
			tmp := strings.TrimSpace(userInput[4:])
			handleTypeCommand(tmp)
		} else if cmds[0]=="pwd"{
			mydir, err := os.Getwd() 
			if err!=nil{
				fmt.Println(err)
			}

			fmt.Println(mydir)
		} else if cmds[0]=="cd"{
			err := os.Chdir(cmds[1])
			if err!=nil{
				fmt.Printf("cd: %s: No such file or directory\n", cmds[1])
			}
		} else {
			command := exec.Command(cmds[0], cmds[1:]...)
			command.Stdout = os.Stdout
			command.Stderr = os.Stderr
			err := command.Run()
			if err != nil {
				fmt.Printf("%s: command not found\n", cmds[0])
			}
		}

	}
	// a better approach here would be
	// fmt.Printf("%s: command not found\n", strings.TrimSpace(userInput))
}

func checkIfExecInPath(e string) (string, error) {
	path, err := exec.LookPath(e)
	return path, err
}

func handleTypeCommand(tmp string) {
	if tmp == "echo" || tmp == "exit" || tmp == "type" || tmp=="pwd" {
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
