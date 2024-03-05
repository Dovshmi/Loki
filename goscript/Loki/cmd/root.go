/*
Copyright © 2024 Dovshmi
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

var (
	a string
	o string
	d string
	l string
	f string
)

var (
	t bool
	p bool
)

var languageMap = map[string]struct {
	extension string
	shebang   string
	comments  string // Single-line and multi-line comment symbols
}{
	"bash":       {"sh", "#!/bin/bash", "#"},
	"python":     {"py", "#!/usr/bin/env python", "#"},
	"perl":       {"pl", "#!/usr/bin/perl", "#"},
	"ruby":       {"rb", "#!/usr/bin/env ruby", "#"},
	"php":        {"php", "#!/usr/bin/env php", "//"},
	"javascript": {"js", "#!/usr/bin/env node", "//"},
	"nim":        {"nim", "", "#"},
	"nimlang":    {"nim", "", "#"},
	"c":          {"c", "", "//"},
	"c++":        {"cpp", "", "//"},
	"java":       {"java", "", "//"},
	"golang":     {"go", "", "//"},
	"go":         {"go", "", "//"},
	"html":       {"html", "", ""},
	"css":        {"css", "", ""},
	"markdown":   {"md", "", ""},
	"rust":       {"rs", "", "//"},
	"kotlin":     {"kt", "#!/usr/bin/env kotlin", "//"},
	"typescript": {"ts", "", "//"},
	"scala":      {"scala", "#!/usr/bin/env scala", "//"},
	"lua":        {"lua", "", "--"},
	"r":          {"r", "", "#"},
	"vhdl":       {"vhd", "", "--"},
	"verilog":    {"v", "", "//"},
	"lisp":       {"lisp", "", ";"},
	"scheme":     {"scm", "", ";"},
	"haskell":    {"hs", "", "--"},
	"prolog":     {"pl", "", "%"},
	"julia":      {"jl", "", "#"},
}

func handleLanguage(lang string) (string, string, string) {
	if val, ok := languageMap[lang]; ok {
		return val.extension, val.shebang, val.comments
	}
	return "false", "", ""
}

func extolan(ex string) string {
	for key, val := range languageMap {
		if val.extension == ex {
			return key
		}
	}
	return "false"
}

func filewrite(output string, shabeng string, comment string, description string, asciis ...[]string) {
	if _, err := os.Stat(output); os.IsNotExist(err) {
		// Create the file
		file, err := os.Create(output)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer file.Close()
		// content := fmt.Sprintf("%s\n%s\n%s\n%s\n", output, shabeng, comment, description)
		_, err = file.WriteString(shabeng + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
		for i, ascii := range asciis {
			for _, line := range ascii {
				fmt.Fprintf(file, "%s%s\n", comment, line)
			}
			if i == len(asciis)-1 {
				file.WriteString(comment + "\n" + comment + " Description: " + description)
			}
		}

		fmt.Println("Content written to the file successfully.")
		err = os.Chmod(output, 0755)
		if err != nil {
			fmt.Println("Error making the file executable:", err)
		} else {
			fmt.Println("File is now executable.")
		}
	} else {
		fmt.Println("File already exists. Not writing to it.")
	}
}

func printfig(shabeng string, comment string, description string, asciis ...[]string) {
	fmt.Println(shabeng)
	for i, ascii := range asciis {
		for _, line := range ascii {
			fmt.Println(comment, line)
		}
		if i == len(asciis)-1 {
			fmt.Println(comment)
			fmt.Println(comment, "Description: ", description)
		}
	}
}

func figlet(cmd *cobra.Command, author string, output string, description string, lang string, font string, times bool, prints bool) {
	name := strings.Split(output, ".")
	if len(name) > 1 && lang == "" {
		lang = extolan(name[1])
	} else if len(name) > 1 && lang != "" {
		cmd.Help()
		os.Exit(2)
	} else if len(name) == 1 && lang == "" {
		cmd.Help()
		os.Exit(1)
	}
	realan := extolan(lang)
	if realan != "false" {
		lang = realan
	}
	ex, shabeng, comment := handleLanguage(lang)
	if ex == "false" {
		cmd.Help()
		os.Exit(3)
	}
	output = name[0] + "." + ex
	asciio := figure.NewFigure(name[0], font, true)
	asciia := figure.NewFigure("By ."+author, font, true)
	formattedTime := time.Now().UTC().Format("02 Jan 2006")
	asciit := figure.NewFigure(formattedTime, font, true)

	if times {
		filewrite(output, shabeng, comment, description, asciio.Slicify(), asciia.Slicify(), asciit.Slicify())
	} else {
		filewrite(output, shabeng, comment, description, asciio.Slicify(), asciia.Slicify())
	}
	if prints {
		if times {
			printfig(shabeng, comment, description, asciio.Slicify(), asciia.Slicify(), asciit.Slicify())
		} else {
			printfig(shabeng, comment, description, asciio.Slicify(), asciia.Slicify())
		}
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "loki",
	Short: "Loki... A Custom Script Header Generator",
	Args:  cobra.MatchAll(cobra.OnlyValidArgs),
	Long: `Loki... A Custom Script Header Generator
Designed to create custom script headers with Figlet ASCII art.
Generating headers that include script name, author, description, date and Figlet ASCII art.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Error:", r)
				cmd.Help()
				return
			}
		}()
		figlet(cmd, a, o, d, l, f, t, p)
		fmt.Println(a, o, d, l, f, t, p)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.Cobra.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&a, "author", "a", "", "Specify the author name.")
	rootCmd.Flags().StringVarP(&o, "output", "o", "", "Specify the output file name.")
	rootCmd.Flags().StringVarP(&d, "description", "d", "", "Specify the header description.")
	rootCmd.Flags().StringVarP(&l, "language", "l", "", "Specify the script language.")
	rootCmd.Flags().StringVarP(&f, "font", "f", "", "Specify the Figlet font.")
	rootCmd.Flags().BoolVarP(&t, "time", "t", false, "Add the current date to the header.")
	rootCmd.Flags().BoolVarP(&p, "print", "p", false, "Print the header.")
	rootCmd.MarkFlagRequired("author")
	rootCmd.MarkFlagRequired("output")
	// make the help menu not ordered
	rootCmd.Flags().SortFlags = false

	// change the usage shown in the help command
	rootCmd.SetUsageTemplate(`Usage:
    {{.UseLine}}

Flags:
{{.LocalFlags.FlagUsages | trimRightSpace}}{{if .HasAvailableInheritedFlags}}

Global Flags:
{{.InheritedFlags.FlagUsages | trimRightSpace}}{{end}}

Examples:
  loki -a "John Doe" -d "A sample script" -o script.sh
      Generates a script header with author name "John Doe", description "A sample script", and saves it to "script.sh".

  loki -a "Jane Smith" -d "Python script" -o script.py -l python -f standard -t
      Generates a script header with author name "Jane Smith", description "Python script", current date, and saves it to "script.py".

  loki -a "Anonymous" -d "Bash script" -o script.bash -l bash -f mini -p
      Generates a script header with author name "Anonymous", description "Bash script", using mini font, and prints it to console.`)
}
