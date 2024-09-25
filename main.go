package main

import (
	"flag"
	"fmt"
)

func main() {
	if len(args) < 2 {
		fmt.Println("usage: gdirbackub [source_dirname] > [destination] [options]")
		panic("not enough arguments provided")
	}

	flag.Parse()
	err := Read(*source)
	if err != nil {
		panic(err)
	}
}