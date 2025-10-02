package main

import (
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
)

func main() {

	// **** Running a simple command and capturing its output ****
	cmd := exec.Command("printenv", "SHELL")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Println("Output:", string(output))

	// Another example with echo
	cmd = exec.Command("echo", "Hello World!")
	output, err = cmd.Output()
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

	// **** Starting a long-running process and managing it ****
	// Here we use "sleep" command to simulate a long-running process
	cmd = exec.Command("sleep", "3")
	err = cmd.Start()
	if err != nil {
		fmt.Println("Error starting command:", err)
		return
	}

	time.Sleep(2 * time.Second) // Simulating some work with sleep

	// Process ID
	fmt.Println("Process ID:", cmd.Process.Pid)

	// Waiting: comment this out to see the difference
	err = cmd.Wait()
	if err != nil {
		fmt.Println("Error waiting for command:", err)
		return
	}
	fmt.Println("Command finished successfully")

	// Killing the process after 2 seconds of sleep
	// Comment the Wait() above to see the kill in action
	// err = cmd.Process.Kill()
	// if err != nil {
	// 	fmt.Println("Error killing command:", err)
	// 	return
	// }
	// fmt.Println("Process killed")

	pr, pw := io.Pipe()
	cmd = exec.Command("grep", "foo")
	cmd.Stdin = pr

	go func() {
		defer pw.Close()
		pw.Write([]byte("food is good\nbar\nbaz\nfoo bar izo"))
	}()

	output, err = cmd.Output()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}
	fmt.Println("Grep Output with Pipe:\n", string(output))
}
