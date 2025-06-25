package main

import (
	"os"
	"text/template"
)

func main() {
	// tmpl := template.New("example")
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
}
