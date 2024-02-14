package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	figure "github.com/common-nighthawk/go-figure"
)

var supportedFonts = map[string]bool{
	"small":     true,
	"script":    true,
	"mini":      true,
	"bubbles":   true,
	"jerusalim": true,
}

type model struct {
	ScriptName    string
	AuthorName    string
	Description   string
	ScriptLang    string
	Font          string
	Step          int
	ErrorMessage  string
	ScriptCreated bool
}

func (m model) Init() tea.Cmd {
	return tea.Batch(tea.EnterAltScreen, tea.ClearScreen)
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		case "enter":
			m.Step++
    case "backspace":
			switch m.Step {
			case 0:
				if len(m.ScriptName) > 0 {
					m.ScriptName = m.ScriptName[:len(m.ScriptName)-1]
				}
			case 1:
				if len(m.AuthorName) > 0 {
					m.AuthorName = m.AuthorName[:len(m.AuthorName)-1]
				}
			case 2:
				if len(m.Description) > 0 {
					m.Description = m.Description[:len(m.Description)-1]
				}
			case 3:
				if len(m.ScriptLang) > 0 {
					m.ScriptLang = m.ScriptLang[:len(m.ScriptLang)-1]
				}
			case 4:
				if len(m.Font) > 0 {
					m.Font = m.Font[:len(m.Font)-1]
				}
			} 
    case "left":
			switch m.Step {
			case 0:
				if len(m.ScriptName) > 0 {
					m.ScriptName = m.ScriptName[:len(m.ScriptName)-1]
				}
			case 1:
				if len(m.AuthorName) > 0 {
					m.AuthorName = m.AuthorName[:len(m.AuthorName)-1]
				}
			case 2:
				if len(m.Description) > 0 {
					m.Description = m.Description[:len(m.Description)-1]
				}
			case 3:
				if len(m.ScriptLang) > 0 {
					m.ScriptLang = m.ScriptLang[:len(m.ScriptLang)-1]
				}
			case 4:
				if len(m.Font) > 0 {
					m.Font = m.Font[:len(m.Font)-1]
				}
			}
    default:
			switch m.Step {
			case 0:
				m.ScriptName += msg.String()
			case 1:
				m.AuthorName += msg.String()
			case 2:
				m.Description += msg.String()
			case 3:
				m.ScriptLang += msg.String()
			case 4:
				m.Font += msg.String()
			}
		}
	}

	if m.Step == 5 {
		if !validateScriptLanguage(m.ScriptLang) {
			m.ErrorMessage = "Unsupported language. Please choose from: bash, python, perl, ruby, php, javascript, c, cpp, java, go"
			m.Step--
		} else if !validateFont(m.Font) {
			m.ErrorMessage = "Unsupported font. Please choose from: small, script, mini, bubbles, jerusalim"
			m.Step--
		} else {
			m.ScriptCreated = true
			createScriptHeader(m.ScriptName, m.AuthorName, m.Description, m.ScriptLang, m.Font)
		}
	}

	return m, nil
}

func (m model) View() string {
	var output string

	if m.ScriptCreated {
		return fmt.Sprintf("Custom script header created successfully!\n")
	}

	switch m.Step {
	case 0:
		output = "Enter your script name: " + m.ScriptName
	case 1:
		output = "Enter your name or alias: " + m.AuthorName
	case 2:
		output = "Enter script description: " + m.Description
	case 3:
		output = "Enter script language (bash, python, perl, ruby, php, javascript, c, cpp, java, go): " + m.ScriptLang
	case 4:
		output = "Enter font for ASCII art (small, script, mini, bubbles, jerusalim): " + m.Font
	default:
		output = m.ErrorMessage
	}

	return output
}

func main() {
	p := tea.NewProgram(model{})
	if err := p.Start(); err != nil {
		fmt.Printf("Error starting program: %v", err)
		os.Exit(1)
	}
}

func validateFont(font string) bool {
	return supportedFonts[strings.ToLower(font)]
}

func validateScriptLanguage(scriptLang string) bool {
	switch strings.ToLower(strings.TrimSpace(scriptLang)) {
	case "bash", "python", "perl", "ruby", "php", "javascript", "c", "cpp", "java", "go":
		return true
	default:
		return false
	}
}

func createScriptHeader(scriptName, authorName, description, scriptLang, font string) {
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
