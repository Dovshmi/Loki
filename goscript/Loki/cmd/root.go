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

var t bool

var languageMap = map[string]struct {
	extension string
	shebang   string
	comments  string // Single-line and multi-line comment symbols
}{
	"bash":       {"sh", "#!/bin/bash", "#"},
	"sh":         {"sh", "#!/bin/bash", "#"},
	"python":     {"py", "#!/usr/bin/env python", "#"},
	"py":         {"py", "#!/usr/bin/env python", "#"},
	"perl":       {"pl", "#!/usr/bin/perl", "#"},
	"pl":         {"pl", "#!/usr/bin/perl", "#"},
	"ruby":       {"rb", "#!/usr/bin/env ruby", "#"},
	"rb":         {"rb", "#!/usr/bin/env ruby", "#"},
	"php":        {"php", "#!/usr/bin/env php", "//"},
	"javascript": {"js", "#!/usr/bin/env node", "//"},
	"js":         {"js", "#!/usr/bin/env node", "//"},
	"nim":        {"nim", "", "#"},
	"nimlang":    {"nim", "", "#"},
	"c":          {"c", "", "//"},
	"cpp":        {"cpp", "", "//"},
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

func getcap(lang string) string {
	if val, ok := languageMap[lang]; ok {
		return val.comments
	}
	return "false"
}

func getShebang(lang string) string {
	if val, ok := languageMap[lang]; ok {
		return val.shebang
	}
	return "false"
}

func filgetwithcap(ascii []string, lang string) {
	// lines := ascii.Slicify()
	for _, line := range ascii {
		fmt.Println(getcap(lang), line)
	}
	// fmt.Println(ascii.String(), a, o, d, l, f, t)
	// fmt.Println(handleLanguage(l))
}

func extolan(ex string) string {
	for key, val := range languageMap {
		if val.extension == ex {
			return key
		}
	}
	return "false"
}

func figlet(cmd *cobra.Command, author string, output string, description string, lang string, font string, times bool) {
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
	// Need to add a check for false language , and add to exists 1,2
	asciio := figure.NewFigure(name[0], font, true)
	asciia := figure.NewFigure("By ."+author, font, true)
	fmt.Println(getShebang(lang))
	filgetwithcap(asciio.Slicify(), lang)
	filgetwithcap(asciia.Slicify(), lang)
	if times {
		formattedTime := time.Now().UTC().Format("02 Jan 2006")
		asciit := figure.NewFigure(formattedTime, font, true)
		filgetwithcap(asciit.Slicify(), lang)
	}
	fmt.Println(getcap(lang))
	fmt.Println(getcap(lang), "Description:", description)
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "loki",
	Short: "A brief description of your application",
	Args:  cobra.MatchAll(cobra.OnlyValidArgs),
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
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
		figlet(cmd, a, o, d, l, f, t)
		// ascii := figure.NewFigure(o, f, true)
		//lines := ascii.Slicify()
		//for _, line := range lines {
		//	fmt.Println(getcap(l), line)
		//}
		//fmt.Println()
		fmt.Println(a, o, d, l, f, t)
		// fmt.Println(handleLanguage(l))
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
	rootCmd.Flags().BoolVarP(&t, "time", "t", false, "Add the current date to the header.")
	rootCmd.Flags().StringVarP(&a, "author", "a", "", "Specify the author name.")
	rootCmd.Flags().StringVarP(&o, "output", "o", "", "Specify the output file name.")
	rootCmd.Flags().StringVarP(&d, "description", "d", "", "Specify the header description.")
	rootCmd.Flags().StringVarP(&l, "language", "l", "", "Specify the script language.")
	rootCmd.Flags().StringVarP(&f, "font", "f", "", "Specify the Figlet font.")
	rootCmd.MarkFlagRequired("author")
	rootCmd.MarkFlagRequired("output")
}