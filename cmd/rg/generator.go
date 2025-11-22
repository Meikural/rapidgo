package main

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

//go:embed templates/*
var templates embed.FS

func createNewProject(appName string) error {
	// Validate app name
	if appName == "" {
		return fmt.Errorf("app name cannot be empty")
	}

	// Check if directory already exists
	if _, err := os.Stat(appName); err == nil {
		return fmt.Errorf("directory '%s' already exists", appName)
	}

	fmt.Printf("Creating RapidGo project: %s\n", appName)

	// Create directory structure
	if err := createDirectories(appName); err != nil {
		return err
	}
	fmt.Println("✓ Created folders")

	// Generate files
	if err := generateFiles(appName); err != nil {
		return err
	}
	fmt.Println("✓ Generated files")

	// Print next steps
	fmt.Printf("\nNext steps:\n")
	fmt.Printf("  cd %s\n", appName)
	fmt.Printf("  go mod tidy\n")
	fmt.Printf("  go run main.go\n")
	fmt.Printf("\nVisit http://localhost:8080\n")

	return nil
}

func createDirectories(appName string) error {
	dirs := []string{
		appName,
		filepath.Join(appName, "app"),
		filepath.Join(appName, "app", "logic"),
		filepath.Join(appName, "app", "models"),
		filepath.Join(appName, "app", "services"),
		filepath.Join(appName, "config"),
		filepath.Join(appName, "db"),
		filepath.Join(appName, "db", "migrations"),
		filepath.Join(appName, "public"),
	}

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory %s: %w", dir, err)
		}
	}

	return nil
}

func generateFiles(appName string) error {
	files := map[string]string{
		filepath.Join(appName, "main.go"):                  "templates/main.go.tmpl",
		filepath.Join(appName, "app", "logic", "hello.go"): "templates/hello.go.tmpl",
		filepath.Join(appName, ".env"):                     "templates/.env.tmpl",
		filepath.Join(appName, "go.mod"):                   "templates/go.mod.tmpl",
		filepath.Join(appName, "README.md"):                "templates/README.md.tmpl",
	}

	for filePath, templatePath := range files {
		content, err := templates.ReadFile(templatePath)
		if err != nil {
			return fmt.Errorf("failed to read template %s: %w", templatePath, err)
		}

		// Replace placeholders
		fileContent := string(content)
		fileContent = strings.ReplaceAll(fileContent, "[APPNAME]", appName)

		if err := writeFile(filePath, fileContent); err != nil {
			return err
		}
	}

	return nil
}

func writeFile(path, content string) error {
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file %s: %w", path, err)
	}
	return nil
}