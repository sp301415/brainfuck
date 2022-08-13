package brainfuck

import (
	"bufio"
	"fmt"
	"os"
)

var stdout = bufio.NewWriterSize(os.Stdout, 8192)

// Execute executes a Program.
func Execute(code Program) {
	defer stdout.Flush()

	MEM := [30000]int{}
	ptr := 0

	for pc := 0; pc < len(code); pc++ {
		instruction := code[pc]
		switch instruction.Type {
		case ChangePtr:
			ptr += instruction.Operand
		case ChangeMem:
			MEM[ptr] += instruction.Operand
		case Scan:
			fmt.Scanf("%c", &MEM[ptr])
		case Print:
			fmt.Fprintf(stdout, "%c", MEM[ptr])
		case JumpIf:
			if MEM[ptr] == 0 {
				pc = instruction.Operand
			}
		case Jump:
			pc = instruction.Operand
		}
	}
}
