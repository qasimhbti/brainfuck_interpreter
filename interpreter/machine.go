package main

import (
	"io"

	"github.com/pkg/errors"
)

// Machine define the BrainFuck Machine
type Machine struct {
	code   string
	ip     int
	memory [30000]int
	dp     int
	input  io.Reader
	output io.Writer
	buf    []byte
}

// NewMachine represent the BrainFuck Machine
func NewMachine(code string, in io.Reader, out io.Writer) *Machine {
	return &Machine{
		code:   code,
		input:  in,
		output: out,
		buf:    make([]byte, 1),
	}
}

// Execute use to process the machine.
func (m *Machine) Execute() error {
	for m.ip < len(m.code) {
		ins := m.code[m.ip]

		switch ins {
		case '+':
			m.memory[m.dp]++
		case '-':
			m.memory[m.dp]--
		case '>':
			m.dp++
		case '<':
			m.dp--
		case ',':
			err := m.readChar()
			if err != nil {
				return err
			}
		case '.':
			err := m.putChar()
			if err != nil {
				return err
			}
		case '[':
			if m.memory[m.dp] == 0 {
				depth := 1
				for depth != 0 {
					m.ip++
					switch m.code[m.ip] {
					case '[':
						depth++
					case ']':
						depth--
					}
				}
			}
		case ']':
			if m.memory[m.dp] != 0 {
				depth := 1
				for depth != 0 {
					m.ip--
					switch m.code[m.ip] {
					case ']':
						depth++
					case '[':
						depth--
					}
				}
			}
		}
		m.ip++
	}
	return nil
}

func (m *Machine) readChar() error {
	n, err := m.input.Read(m.buf)
	if err != nil {
		return errors.WithMessage(err, "read char")
	}
	if n != 1 {
		return errors.New("wrong num bytes read")
	}

	m.memory[m.dp] = int(m.buf[0])
	return nil
}

func (m *Machine) putChar() error {
	m.buf[0] = byte(m.memory[m.dp])

	n, err := m.output.Write(m.buf)
	if err != nil {
		return errors.WithMessage(err, "put char")
	}
	if n != 1 {
		return errors.New("wrong num bytes written")
	}
	return nil
}
