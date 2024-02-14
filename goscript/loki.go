package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	figure "github.com/common-nighthawk/go-figure"
)

var supportedFonts = map[string]bool{
	"small":     true,
	"script":    true,
	"mini":      true,
	"bubbles":   true,
	"jerusalim": true,
}
// createScriptHeader creates a custom script header based on user input
func createScriptHeader() {
	fmt.Println("Creating Custom Script Header...")
	var scriptName, authorName, description, scriptLang, font string

	fmt.Print("Enter your script name: ")
	fmt.Scanln(&scriptName)
	fmt.Print("Enter your name or alias: ")
	fmt.Scanln(&authorName)
	fmt.Print("Enter script description: ")
	fmt.Scanln(&description)
	fmt.Print("Enter script language (bash, python, perl, ruby, php, javascript, c, cpp, java, go): ")
	fmt.Scanln(&scriptLang)
	fmt.Print("Enter font for ASCII art (small, script, mini, bubbles, jerusalim): ")
	fmt.Scanln(&font)

  
	// Check if the entered font is supported
	if !supportedFonts[strings.ToLower(font)] {
		fmt.Println("Unsupported font. Please choose from: small, script, mini, bubbles, jerusalim")
		return
	}
	var fileExtension, shebang string

	switch strings.ToLower(strings.TrimSpace(scriptLang)) {
	case "bash":
		fileExtension = "sh"
		shebang = "#!/bin/bash"
	case "python":
		fileExtension = "py"
		shebang = "#!/usr/bin/env python"
	case "perl":
		fileExtension = "pl"
		shebang = "#!/usr/bin/perl"
	case "ruby":
		fileExtension = "rb"
		shebang = "#!/usr/bin/env ruby"
	case "php":
		fileExtension = "php"
		shebang = "#!/usr/bin/env php"
	case "javascript":
		fileExtension = "js"
		shebang = "#!/usr/bin/env node"
	case "c":
		fileExtension = "c"
	case "cpp":
		fileExtension = "cpp"
	case "java":
		fileExtension = "java"
	case "go":
		fileExtension = "go"
	default:
		fmt.Println("Unsupported language. Exiting...")
		return
	}

	fileName := fmt.Sprintf("%s.%s", scriptName, fileExtension)
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Write shebang line
	if shebang != "" {
		fmt.Fprintf(file, "%s\n", shebang)
	}

	// Create ASCII art with selected font
	asciiScriptName := figure.NewFigure(scriptName, font, true)
	asciiAuthor := figure.NewFigure("By ."+authorName, font, true)
	dateStr := time.Now().Format("Mon 02 Jan 2006")
	asciiDate := figure.NewFigure(dateStr, font, true)

	// Write ASCII art to file (commented out)
	commentedArt := func(art string) string {
		lines := strings.Split(art, "\n")
		for i, line := range lines {
			lines[i] = "# " + line
		}
		return strings.Join(lines, "\n")
	}

	fmt.Fprintf(file, "%s\n", commentedArt(asciiScriptName.String()))
	fmt.Fprintf(file, "%s\n", commentedArt(asciiAuthor.String()))
	fmt.Fprintf(file, "%s\n", commentedArt(asciiDate.String()))
	fmt.Fprintf(file, "# Description: %s\n", description)

	// Set executable permissions based on OS
	switch runtime.GOOS {
	case "linux":
		if err := os.Chmod(fileName, 0755); err != nil {
			fmt.Printf("Error setting executable permission: %v\n", err)
		}
	case "darwin":
		if err := exec.Command("chmod", "+x", fileName).Run(); err != nil {
			fmt.Printf("Error setting executable permission: %v\n", err)
		}
	default:
		// Do nothing for Windows
	}

	fmt.Printf("Custom script header created successfully in %s!\n", fileName)
}

func main() {
	createScriptHeader()
}
