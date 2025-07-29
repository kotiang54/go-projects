package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

func readFromReader(r io.Reader) {
	buf := make([]byte, 1024)
	n, err := r.Read(buf)
	if err != nil {
		log.Fatalln("Error reading from reader:", err)
		return
	}
	fmt.Println(string(buf[:n]))
}

func writeToWriter(w io.Writer, data string) {
	_, err := w.Write([]byte(data))
	if err != nil {
		log.Fatalln("Error writing to writer:", err)
	}
	fmt.Printf("Wrote %d bytes: %s\n", len(data), data)
}

func closeResource(c io.Closer) {
	if err := c.Close(); err != nil {
		log.Fatalln("Error closing resource:", err)
	}
}

func bufferExample() {
	var buf bytes.Buffer // allocate memory on stack
	buf.WriteString("Hello, buffer io package!\n")
	fmt.Println(buf.String())
}

// multiReaderExample demonstrates how to use io.MultiReader
// to concatenate multiple readers into a single stream.
func multiReaderExample() {
	r1 := strings.NewReader("First part of the data.\n")
	r2 := strings.NewReader("Second part of the data.\n")
	multiReader := io.MultiReader(r1, r2)
	buf := new(bytes.Buffer) // allocate memory on a heap

	_, err := buf.ReadFrom(multiReader)
	if err != nil {
		log.Fatalln("Error reading from multi-reader:", err)
	}
	fmt.Println(buf.String())
}

// pipeExample demonstrates the use of io.Pipe to create an in-memory pipe
// for concurrent data transfer between a writer and a reader. It launches
// a goroutine to write data to the pipe, reads the data from the pipe into a buffer,
func pipeExample() {
	pr, pw := io.Pipe() // create a pipe
	go func() {
		pw.Write([]byte("Data sent through the pipe.\n"))
		pw.Close()
	}()

	buf := new(bytes.Buffer)
	buf.ReadFrom(pr)
	fmt.Println("Data received from pipe:", buf.String())
	closeResource(pr)
	closeResource(pw)
	fmt.Println("Pipe closed successfully.")
}

func main() {

	fmt.Println("=== Read from Reader ===")
	readFromReader(strings.NewReader("Hello, io package!\nThis is a test."))

	fmt.Println("=== Write to Writer ===")
	var writer bytes.Buffer
	writeToWriter(&writer, "Writing to a bytes.Buffer using io.Writer interface.\n")
	fmt.Println(writer.String())

	fmt.Println("=== Buffer Example ===")
	bufferExample()

	fmt.Println("=== MultiReader Example ===")
	multiReaderExample()

	fmt.Println("=== Closing Resources ===")
	fmt.Println("No resources to close for bytes.Buffer.")

	fmt.Println("=== Pipe Example ===")
	pipeExample()
}
