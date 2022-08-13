package brainfuck

import (
	"fmt"
)

// Execute executes a Program.
func Execute(program Program) {
	MEM := [30000]int{}
	ptr := 0

	for pc := 0; pc < len(program); pc++ {
		instruction := program[pc]
		switch instruction.Type {
		case ChangePtr:
			ptr += instruction.Operand
		case ChangeMem:
			MEM[ptr] += instruction.Operand
		case Scan:
			fmt.Scanf("%c", &MEM[ptr])
		case Print:
			fmt.Printf("%c", MEM[ptr])
		case JumpIf:
			if MEM[ptr] == 0 {
				pc = instruction.Operand
			}
		case Jump:
			pc = instruction.Operand
		}
	}
}
