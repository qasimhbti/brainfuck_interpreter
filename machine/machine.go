package main

import "io"

import "github.com/pkg/errors"

// Machine - represent brainfuck machine
type Machine struct {
	code []*Instruction
	ip   int

	memory [30000]int
	dp     int

	input  io.Reader
	output io.Writer

	readBuf []byte
}

// NewMachine - give new BrainFuck Machine
func NewMachine(instructions []*Instruction, in io.Reader, out io.Writer) *Machine {
	return &Machine{
		code:    instructions,
		input:   in,
		output:  out,
		readBuf: make([]byte, 1),
	}
}

// Execute - execute machine
func (m *Machine) Execute() error {
	for m.ip < len(m.code) {
		ins := m.code[m.ip]

		switch ins.Type {
		case Plus:
			m.memory[m.dp] += ins.Argument
		case Minus:
			m.memory[m.dp] -= ins.Argument
		case Right:
			m.dp += ins.Argument
		case Left:
			m.dp -= ins.Argument
		case PutChar:
			for i := 0; i < ins.Argument; i++ {
				err := m.putChar()
				if err != nil {
					return err
				}
			}
		case ReadChar:
			for i := 0; i < ins.Argument; i++ {
				err := m.readChar()
				if err != nil {
					return err
				}
			}
		case JumpIfZero:
			if m.memory[m.dp] == 0 {
				m.ip = ins.Argument
				continue
			}
		case JumpIfNotZero:
			if m.memory[m.dp] != 0 {
				m.ip = ins.Argument
				continue
			}
		// Custom Operation
		default:
			for _, v := range customInstructions {
				if ins.Type == InsType(v.Type) {
					m.dp += ins.Argument * 2
				}
			}

		}

		m.ip++
	}
	return nil
}

func (m *Machine) readChar() error {
	n, err := m.input.Read(m.readBuf)
	if err != nil {
		return errors.WithMessage(err, "readBuf")
	}
	if n != 1 {
		return errors.New("wrong num bytes read")
	}

	m.memory[m.dp] = int(m.readBuf[0])
	return nil
}

func (m *Machine) putChar() error {
	m.readBuf[0] = byte(m.memory[m.dp])

	n, err := m.output.Write(m.readBuf)
	if err != nil {
		return errors.WithMessage(err, "writeBuf")
	}
	if n != 1 {
		return errors.New("wrong num bytes written")
	}
	return nil
}
