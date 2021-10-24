package main

import (
	"flag"
	"os"
)

var (
	showVersion bool
)

func init() {
	flag.BoolVar(&showVersion, "v", false, "show version")
}

func main() {
	flag.Parse()

	if showVersion {
		printVersion()
		os.Exit(0)
	}
}
