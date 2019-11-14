package main

import (
	"io/ioutil"
	"log"
	"os"

	brainfuckinterpreter "github.com/brainfuck_interpreter"

	"github.com/brainfuck_interpreter/utils"
	"github.com/pkg/errors"
)

func main() {
	err := run()
	if err != nil {
		handleError(err)
	}
}

func run() error {
	utils.InitLog()
	utils.LogStart(brainfuckinterpreter.Version, "development")

	// Add the Custom operation
	var operation InsType
	addCustomOperation([]byte("?"), operation)

	// Read the file
	fileName := os.Args[1]
	code, err := ioutil.ReadFile(fileName)
	if err != nil {
		return errors.WithMessage(err, "reading file")
	}

	// Create set of instructions
	compiler := NewCompiler(string(code))
	instructions := compiler.Compile()

	// Prepare and execute the BrainFuck Machine
	m := NewMachine(instructions, os.Stdin, os.Stdout)
	err = m.Execute()
	if err != nil {
		return errors.WithMessage(err, "Execute")
	}
	return nil
}

func handleError(err error) {
	log.Fatalf("%s", err)
}
