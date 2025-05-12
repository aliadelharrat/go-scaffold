Go Project Creator
A Go CLI tool to create a new Go project with a valid name (starts with a letter, no whitespace, only letters/numbers/underscores), generating main.go and go.mod files.
Prerequisites

Go 1.24.3 or later installed
Git installed

Installation

Clone the repository:
git clone https://github.com/aliadelharrat/your-repo-name.git
cd your-repo-name


Build the project:
go build -o go-project-creator


Add the executable to your environment's PATH:

Linux/macOS:export PATH=$PATH:$(pwd)

Add the above line to ~/.bashrc or ~/.zshrc for persistence.
Windows:
Move go-project-creator.exe to a directory like C:\Go\bin.
Add C:\Go\bin to your system PATH:
Search for "Environment Variables" in Windows.
Edit the "Path" variable under System or User variables.
Add the directory path and save.







Usage
Run the tool with a project name:
go-project-creator myproject

This creates a myproject folder with:

main.go: Basic "Hello" program
go.mod: Module configuration

Notes

Project name must:
Start with a letter
Contain only letters, numbers, or underscores
Have no whitespace


Example valid names: myProject, hello_world, project123

License
MIT License
