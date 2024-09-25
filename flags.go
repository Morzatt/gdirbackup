package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
)

var args = os.Args;

// Path to the directory to scan and print from.
var source = flag.String("source", "", "path to the directory to scan and print from");
// Path to the destination to print in.
var destination = flag.String("destination", "", "path to the destination file to print in, if no destination is provided it will print out directly into the Standard Output")

// Indicates the subdirectories inside the directory to scan to ignore
var ignore = flag.String("ignore", "", "indicates the name of a subdirectory or file inside the directory to scan to ignore")

var depth int = 0;
var filler string;

// Read reads the source path provided by the arguments, and prints its directory structure.
func Read(source string) error {
	content, err := os.ReadDir(source)
	if err != nil {
		return err
	}

	if depth == 0 {
		Write(fmt.Sprintf("%s \n", path.Base(source)), *destination)
	}

	for _, c := range content {		
		if c.Name() == *ignore {
			return nil
		}

		depth++
		filler = strings.Repeat(fmt.Sprintf("|%*s", 3, ""), depth)
		if c.IsDir() {
			Write(fmt.Sprintf("%s|─%s/ \n", filler, c.Name()), *destination)
			Read(fmt.Sprintf("%s/%s", source, c.Name())); 
		} else {
			Write(fmt.Sprintf("%s|─%s\n", filler, c.Name()), *destination)
		}
		depth--
	}
	return nil
}

func Write(content string, destination string) {
	switch destination {
		case "": 
			if err := WriteToStdout(content); err != nil {
				panic(err)
			}
		default: 
			if err := WriteToFile(content, destination); err != nil {
				panic(err)
			}
	}
}

func WriteToFile(content, destination string) error {
	f, err := os.OpenFile(destination, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	_, err = w.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func WriteToStdout(content string) error {
	_, err := fmt.Print(content)
	if err != nil {
		return err
	}
	return nil
}

func ParseIgnoreFile() error {
	return nil
}