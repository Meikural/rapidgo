package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "new":
		if len(os.Args) < 3 {
			fmt.Println("Error: app name required")
			fmt.Println("Usage: rg new <appname>")
			os.Exit(1)
		}
		appName := os.Args[2]
		if err := createNewProject(appName); err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
	case "help", "-h", "--help":
		printUsage()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println(`
RapidGo CLI - Project Generator

Usage:
  rg new <appname>        Create a new RapidGo project
  rg help                 Show this help message

Example:
  rg new myapp
  cd myapp
  go mod tidy
  go run main.go
`)
}
