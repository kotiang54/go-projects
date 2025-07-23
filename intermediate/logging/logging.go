package main

import (
	"log"
	"os"
)

func main() {

	log.Println("This is a log message")

	log.SetPrefix("INFO: ")
	log.Println("This is an info message")

	// Log Flags
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.Println("This is a message with date and time")

	// Custom loggers
	infoLogger.Println("This is an info message with custom logger")
	warnLogger.Println("This is a warning message with custom logger")
	errorLogger.Println("This is an error message with custom logger")
}

// Custom log output
var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
)
