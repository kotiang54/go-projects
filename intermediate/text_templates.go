package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {
	// tmpl := template.New("example")
	/*
		tmpl, err := template.New("").Parse("Welcome, {{.name}}! How are you doing?\n")
		if err != nil {
			panic(err)
		}

		// Using template.Must() we do not need to check the error. It's done internally
		tmpl = template.Must(template.New("").Parse("Welcome, {{.name}}! How are you doing?\n"))

		// Define data for the welcome message template
		data := map[string]interface{}{
			"name": "John",
		}

		err = tmpl.Execute(os.Stdout, data)
		if err != nil {
			panic(err)
		}
	*/

	// Reading from the console
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	// Define name templates for different types
	templates := map[string]string{
		"welcome":      "Welcome, {{.name}}! We are glad you joined. ",
		"notification": "{{.name}}, you have a new notification: {{.notification}}",
		"error":        "Ooops! An error occured: {{.errorMessage}}",
	}

	// Parse and store templates
	parsedTemplates := make(map[string]*template.Template)
	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		// Show menu
		fmt.Println("\nMenu:")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notification")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Choose an option: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["welcome"]
			data = map[string]interface{}{"name": name}

		case "2":
			fmt.Println("Enter your notification message: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["notification"]
			data = map[string]interface{}{"name": name, "notification": notification}

		case "3":
			fmt.Println("Enter your error message: ")
			errMsg, _ := reader.ReadString('\n')
			errMsg = strings.TrimSpace(errMsg)
			tmpl = parsedTemplates["error"]
			data = map[string]interface{}{"errorMessage": errMsg}

		case "4":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid choice! please select a valid option")
			continue
		}

		// render and print the template to the console
		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error executing template", err)
		}
	}
}
