package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

func main() {

	cmd := exec.Command("echo", "Hello World!")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Println("Output:", string(output))

	cmd = exec.Command("grep", "foo")
	// Set input for the command
	cmd.Stdin = strings.NewReader("foo\nbar\nbaz\nfoo bar izo")
	output, err = cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Println("Grep Output:\n", string(output))

	cmd = exec.Command("sleep", "3")
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	time.Sleep(2 * time.Second) // Simulating some work with sleep

	// Process ID
	fmt.Println("Process ID:", cmd.Process.Pid)

	// Killing the process after 2 seconds of sleep
	err = cmd.Process.Kill()
	if err != nil {
		fmt.Println("Error killing command:", err)
		return
	}
	fmt.Println("Process killed")

	// Waiting: comment this out to see the difference
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for command:", err)
		return
	}
	fmt.Println("Command finished successfully")
}
