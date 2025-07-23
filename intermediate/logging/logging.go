package main

import (
	"log"
	"os"

	"github.com/sirupsen/logrus"
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

	// log to a file
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer file.Close()

	debugLogger := log.New(file, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLogger.Println("This is a debug message logged to a file")

	// 3rd party logging library example - e.g., logrus
	log := logrus.New()

	// Set log level
	log.SetLevel(logrus.InfoLevel)

	// Set log format
	log.SetFormatter(&logrus.JSONFormatter{})

	// Log messages with different levels
	log.Info("This is an info message.")
	log.Warn("This is a warning message.")
	log.Error("This is an error message.")

	log.WithFields(logrus.Fields{
		"username": "John Doe",
		"method":   "GET",
	}).Info("User logged in.")

}

// Custom log output
var (
	infoLogger  = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warnLogger  = log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime|log.Llongfile)
)
