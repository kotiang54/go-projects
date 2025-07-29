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

func main() {

	fmt.Println("=== Read from Reader ===")
	readFromReader(strings.NewReader("Hello, io package!\nThis is a test."))

	fmt.Println("=== Write to Writer ===")
	var writer bytes.Buffer
	writeToWriter(&writer, "Writing to a bytes.Buffer using io.Writer interface.\n")

}
