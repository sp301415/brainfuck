package brainfuck

// InstructionType represents the type of operation in the brainfuck language.
type InstructionType int

const (
	NONE = InstructionType(iota)
	ChangePtr
	ChangeMem
	Scan
	Print
	JumpIf
	Jump
)

// Instruction represents a single instruction in the brainfuck language.
type Instruction struct {
	Type    InstructionType
	Operand int
}

// Program is a executable slice of instructions.
type Program []Instruction
