// sample.go

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Simulate a potential security issue: executing a shell command with user input
	userInput := os.Args[1]
	command := fmt.Sprintf("echo %s", userInput)

	cmd := exec.Command("sh", "-c", command)

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	// Print the command output
	fmt.Println("Command output:", string(output))
}
