package day08

import (
	"errors"
	"fmt"
)

// An OpCode specifies which instruction to execute.
type OpCode string

const (
	// Jump instruction adds its parameter to the program counter.
	Jump OpCode = "jmp"
	// Accumulator instruction adds its parameter to the accumulator.
	Accumulator OpCode = "acc"
	// NoOp instruction does nothing. Its parameter is read, but ignored.
	NoOp OpCode = "nop"
)

// An Instruction is a single operation that the program will take.
type Instruction struct {
	OpCode       OpCode
	Parameter    int
	ExecutedOnce bool
}

// ParseInstruction reads a program line string and converts to an instruction.
func ParseInstruction(line string) (Instruction, error) {
	var i Instruction
	n, err := fmt.Sscanf(line, "%3s %d", &i.OpCode, &i.Parameter)
	if err != nil {
		return Instruction{}, err
	} else if n != 2 {
		return Instruction{}, errors.New("scan did not find two parts")
	}
	return i, nil
}
