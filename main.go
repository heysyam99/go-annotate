package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	path        string
	showVersion bool
	showHelp    bool
)

func init() {
	flag.StringVar(&path, "path", "./models", "custom path")
	flag.BoolVar(&showVersion, "v", false, "show version")
	flag.BoolVar(&showHelp, "h", false, "help")
	flag.Parse()
}

func main() {
	if showVersion {
		printVersion()
		os.Exit(0)
	}

	if showHelp {
		fmt.Println("Usage of go-annotate:\n\nIf no command is provided go-annotate will start the runner with the provided flags\n\nCommands:\n  init  creates a gowatch.yml file with default settings to the current directory\n\nFlags:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	annotate(path)
}
