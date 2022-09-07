package main

import (
	"brainfuck"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Fprintln(os.Stderr, "Error: No input files!")
		os.Exit(1)
	} else if len(os.Args) > 2 {
		fmt.Fprintln(os.Stderr, "Error: Too many arguments!")
		os.Exit(1)
	}

	code, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	program, err := brainfuck.Compile(string(code))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		os.Exit(1)
	}

	brainfuck.Execute(program)
}
