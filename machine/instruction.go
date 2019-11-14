package main

// InsType - instruction type
type InsType byte

const (
	Plus          InsType = '+'
	Minus         InsType = '-'
	Right         InsType = '>'
	Left          InsType = '<'
	PutChar       InsType = '.'
	ReadChar      InsType = ','
	JumpIfZero    InsType = '['
	JumpIfNotZero InsType = ']'
)

// Instruction - represent instruction with argument of consecutive
// instruction (+ , -, >, <, , .) and matching in case of [, ].
type Instruction struct {
	Type     InsType
	Argument int
}

type customInstruction struct {
	Type      byte
	Operation InsType
}

var customInstructions []customInstruction

func addCustomOperation(ins []byte, operation InsType) {
	cusIns := customInstruction{
		Type:      ins[0],
		Operation: operation,
	}

	customInstructions = append(customInstructions, cusIns)
}
