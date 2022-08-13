package brainfuck

import (
	"errors"
)

// Compile compiles the given brainfuck string to Program.
func Compile(code string) (Program, error) {
	program := make(Program, 0, len(code))
	jumpStack := []int{}

	for _, word := range code {
		pc := len(program) - 1
		switch word {
		case '>':
			if pc >= 0 && program[pc].Type == ChangePtr {
				program[pc].Operand++
			} else {
				program = append(program, Instruction{ChangePtr, 1})
			}
		case '<':
			if pc >= 0 && program[pc].Type == ChangePtr {
				program[pc].Operand--
			} else {
				program = append(program, Instruction{ChangePtr, -1})
			}
		case '+':
			if pc >= 0 && program[pc].Type == ChangeMem {
				program[pc].Operand++
			} else {
				program = append(program, Instruction{ChangeMem, 1})
			}
		case '-':
			if pc >= 0 && program[pc].Type == ChangeMem {
				program[pc].Operand--
			} else {
				program = append(program, Instruction{ChangeMem, -1})
			}
		case '.':
			program = append(program, Instruction{Print, 0})
		case ',':
			program = append(program, Instruction{Scan, 0})
		case '[':
			jumpStack = append(jumpStack, pc+1)
			program = append(program, Instruction{JumpIf, 0})
		case ']':
			if len(jumpStack) == 0 {
				return nil, errors.New("mismatched brackets")
			}
			matchingPC := jumpStack[len(jumpStack)-1]
			jumpStack = jumpStack[:len(jumpStack)-1]

			program[matchingPC].Operand = pc + 1
			program = append(program, Instruction{Jump, matchingPC - 1})
		}
	}

	return program, nil
}
