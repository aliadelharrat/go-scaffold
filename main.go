package main

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var fileContent = `package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, %FOLDER_NAME%!")
}`

var modfileContent = `module github.com/aliadelharrat/%FOLDER_NAME%

go 1.24.3
`

var fileName = "main.go"

func main() {
	// Ensure folderName argument is provided
	if len(os.Args) < 2 {
		check(errors.New("Project name argument missing"))
	}
	folderName := os.Args[1]
	folderName = strings.ToLower(folderName)

	// Ensure folderName is valid
	isValid := isValidString(folderName)
	if !isValid {
		check(errors.New("invalid name: must start with a letter, contain only letters, numbers, or underscores, and no whitespace"))
	}

	// Retrieve current working directory
	currentDir, err := os.Getwd()
	check(err)

	folderPath := filepath.Join(currentDir, folderName)
	filePath := filepath.Join(folderPath, fileName)

	// Create project folder
	err = os.MkdirAll(folderPath, os.ModePerm)
	check(err)

	// Create main.go file
	f, err := os.Create(filePath)
	check(err)
	defer f.Close()

	// Subtitute %FOLDER_NAME% with provided project name
	fileContent = strings.ReplaceAll(
		fileContent,
		"%FOLDER_NAME%",
		strings.Title(folderName),
	)

	// Write content to main.go
	err = os.WriteFile(filePath, []byte(fileContent), 0644)
	check(err)

	// Create go.mod file
	modFilePath := filepath.Join(folderPath, "go.mod")
	f, err = os.Create(modFilePath)
	check(err)
	defer f.Close()

	// Write content to go.mod
	modfileContent = strings.ReplaceAll(
		modfileContent,
		"%FOLDER_NAME%",
		strings.ToLower(folderName),
	)
	err = os.WriteFile(modFilePath, []byte(modfileContent), 0644)
	check(err)
}

// Handle and log error
func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func isValidString(s string) bool {
	// Check for whitespace
	if strings.ContainsAny(s, " \t\n") {
		return false
	}
	// Check if string starts with a letter and contains only letters, numbers, underscores
	matched, _ := regexp.MatchString(`^[a-zA-Z][a-zA-Z0-9_-]*$`, s)
	return matched
}
