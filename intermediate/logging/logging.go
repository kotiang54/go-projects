package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

// main demonstrates various logging techniques in Go, including:
// - Using the standard log package for basic logging with custom prefixes and flags.
// - Creating custom loggers for different log levels (info, warning, error).
// - Logging messages to a file using a custom logger.
// - Utilizing a third-party logging library (logrus) for structured logging, setting log levels and formats, and logging messages with additional fields.
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

	// log to a file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer file.Close()

	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is a debug message logged to a file")

	// 3rd party logging library example - e.g., logrus
	logrusLogger := logrus.New()

	// Set log level
	logrusLogger.SetLevel(logrus.InfoLevel)

	// Set log format
	logrusLogger.SetFormatter(&logrus.JSONFormatter{})

	// Log messages with different levels
	logrusLogger.Info("This is an info message.")
	logrusLogger.Warn("This is a warning message.")
	logrusLogger.Error("This is an error message.")

	// Adding structured fields for logging:
	// - "username": Represents the name of the user performing the action.
	// - "method": Represents the HTTP method used in the request.
	logrusLogger.WithFields(logrus.Fields{
		"username": "John Doe", // User's name
		"method":   "GET",      // HTTP method
	}).Info("User logged in.")
}

// Custom log output
var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
)
