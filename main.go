package main

import (
	"fmt"
)

func main() {
	if len(args) < 2 {
		fmt.Println("usage: gdirbackub [source_dirname] [destination] [options]")
		panic("not enough arguments provided")
	}

	source = args[1]
	destination = args[0]

	err := Read(source)
	if err != nil {
		panic(err)
	}
}