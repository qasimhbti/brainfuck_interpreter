package main

// Compiler - use to create set of instructions
type Compiler struct {
	code       string
	codeLength int
	position   int

	instructions []*Instruction
}

// NewCompiler - represent a compiler
func NewCompiler(code string) *Compiler {
	return &Compiler{
		code:         code,
		codeLength:   len(code),
		instructions: []*Instruction{},
	}
}

// Compile - create set of instructions
func (c *Compiler) Compile() []*Instruction {
	loopStack := []int{}

	for c.position < c.codeLength {
		current := c.code[c.position]

		switch current {
		case '[':
			insPos := c.calcArg(JumpIfZero, 0)
			loopStack = append(loopStack, insPos)
		case ']':
			// Pop position of last JumpIfZero ("[") instruction off stack
			openInstruction := loopStack[len(loopStack)-1]
			loopStack = loopStack[:len(loopStack)-1]
			// Emit the new JumpIfNotZero ("]") instruction, with correct position as argument
			closeInstructionPos := c.calcArg(JumpIfNotZero, openInstruction)
			// Patch the old JumpIfZero ("[") instruction with new position
			c.instructions[openInstruction].Argument = closeInstructionPos

		case '+':
			c.compileInstruction('+', Plus)
		case '-':
			c.compileInstruction('-', Minus)
		case '<':
			c.compileInstruction('<', Left)
		case '>':
			c.compileInstruction('>', Right)
		case '.':
			c.compileInstruction('.', PutChar)
		case ',':
			c.compileInstruction(',', ReadChar)
		default:
			//check custom operation
			for _, v := range customInstructions {
				if current == v.Type {
					c.compileInstruction(current, v.Operation)
				}
			}
		}

		c.position++
	}

	return c.instructions
}

func (c *Compiler) compileInstruction(char byte, insType InsType) {
	count := 1

	for c.position < c.codeLength-1 && c.code[c.position+1] == char {
		count++
		c.position++
	}

	c.calcArg(insType, count)
}

func (c *Compiler) calcArg(insType InsType, arg int) int {
	ins := &Instruction{Type: insType, Argument: arg}
	c.instructions = append(c.instructions, ins)
	return len(c.instructions) - 1
}
