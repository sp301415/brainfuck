package main

import (
	"brainfuck"
	"bufio"
	"fmt"
	"os"
)

var stdin = bufio.NewReader(os.Stdin)

func main() {
	for {
		fmt.Print(">>> ")
		code, err := stdin.ReadString('\n')
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		program, err := brainfuck.Compile(code)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		brainfuck.Execute(program)
	}
}
