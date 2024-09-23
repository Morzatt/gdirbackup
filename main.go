package main

import (
	"os"
	"fmt"
)

func main() {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "not enough arguments provided")
		fmt.Println("usage: gdirbackub [source_dirname] [destination] [options]")
		os.Exit(65)
	}
	source = args[1]
	destination = args[0]

	Read(source)
}