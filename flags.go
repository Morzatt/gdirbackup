package main

import (
	"fmt"
	"os"
)

var args = os.Args;
var source string;
var destination string;
var depth int = 0;

// Read reads the source path provided by the arguments, and prints its directory structure.
func Read(source string) error {
	content, err := os.ReadDir(source)
	if err != nil {
		return fmt.Errorf("an error has ocurred when reading the path: %e", err)
	}
	for _, c := range content {		
		depth++
		if c.IsDir() {
			fmt.Printf("%*s |──%s/ \n", depth*2, "", c.Name())
			Read(fmt.Sprintf("%s/%s", source, c.Name())); 
		} else {
			fmt.Printf("%*s |──%s\n", depth*2, "", c.Name())
		}
		depth--
	}
	return nil
}