package day08

import (
	"errors"
	"math"
)

// A Program is a set of instructions and the state of the virtual processor.
type Program struct {
	// Instructions is the list of instructions the program should execute.
	Instructions []Instruction
	// Accumulator is the current value of the virtual processor's accumulator.
	Accumulator int
	// ProgramCounter indicates the slice index of the next instruction to be
	// executed.
	ProgramCounter int
	// Terminated is true when the program has exited normally.
	Terminated bool
	// SafetyCounter is a special counter that can be used to detect infinite
	// loops.
	SafetyCounter int
}

// ParseProgram reads a series of instruction lines and creates a new Program.
func ParseProgram(lines []string) (Program, error) {
	var p Program
	for _, line := range lines {
		i, err := ParseInstruction(line)
		if err != nil {
			return Program{}, err
		}
		p.Instructions = append(p.Instructions, i)
	}
	return p, nil
}

// ExecuteNext executes the next instruction upon the virtual processor.
func (p *Program) ExecuteNext() error {
	// Obtain a reference to the next instruction.
	instr := &p.Instructions[p.ProgramCounter]

	switch instr.OpCode {
	case Jump:
		p.ProgramCounter += instr.Parameter
	case Accumulator:
		p.Accumulator += instr.Parameter
	case NoOp:
		break
	default:
		return errors.New("unknown op code")
	}

	if instr.OpCode != Jump {
		p.ProgramCounter++
	}

	instr.ExecutedOnce = true

	if p.ProgramCounter == len(p.Instructions) {
		p.Terminated = true
	}

	return nil
}

// Reset sets the processor's internal status registers back to the initial
// state so the program can be run again.
func (p *Program) Reset() {
	p.Accumulator = 0
	p.ProgramCounter = 0
	p.Terminated = false
	p.SafetyCounter = 0

	for i := range p.Instructions {
		p.Instructions[i].ExecutedOnce = false
	}
}

// RunUntilLoop will execute the program until either any given instruction is
// executed more than one time or or the program terminates. The value of the
// accumulator when this occurs is returned.
func (p *Program) RunUntilLoop() (int, error) {
	p.Reset()
	for !p.Instructions[p.ProgramCounter].ExecutedOnce && !p.Terminated {
		err := p.ExecuteNext()
		if err != nil {
			return 0, err
		}
	}
	return p.Accumulator, nil
}

// RunWithSafetyCounter will execute the program until either the program
// is considered to be in an infinite loop (safety counter exceeded) or the
// program terminates. The value of the accumulator when either occurs is
// returned.
//
// Interested parties should check p.Terminated after running the program to
// see if the program exited normally.
func (p *Program) RunWithSafetyCounter() (int, error) {
	// The safetyThreshold is an arbitrary value. Once the program has
	// executed this many instructions, it shall consider the program to be
	// stuck in an infinite loop and halt execution.
	//
	// Of course, this probably does not hold up very well in the real world.
	// However, for a toy program like this it is very unlikely that the program
	// will execute the total number of instructions squared.
	safetyThreshold := int(math.Pow(float64(len(p.Instructions)), 2))

	p.Reset()
	for !p.Terminated && p.SafetyCounter < safetyThreshold {
		err := p.ExecuteNext()
		if err != nil {
			return 0, err
		}
		p.SafetyCounter++
	}
	return p.Accumulator, nil
}

// FixOneBugInfiniteLoop can be used to fix programs with exactly one flipped
// noop/jump instruction. The safety counter run method will be used to
// stop infinitely looping programs and attempt the next fix.
//
// If a fix is found, the value of the accumulator is returned and the program
// instructions are left in the fixed state.
// If no fix is available, an error is returned and the program instructions are
// returned to their original state.
func (p *Program) FixOneBugInfiniteLoop() (int, error) {
	// We are told there is exactly one bug in the program in a jump or noop
	// instruction. We shall loop through all instructions and try to flip these
	// until the bug is removed.
	for i := range p.Instructions {
		// Save the original instruction, so we can revert.
		original := p.Instructions[i]
		instr := &p.Instructions[i]

		// Flip instruction and proceed if it is a jump or noop.
		if instr.OpCode == Jump {
			instr.OpCode = NoOp
		} else if instr.OpCode == NoOp {
			instr.OpCode = Jump
		} else {
			continue
		}

		// Run the program with the safety counter enabled.
		acc, err := p.RunWithSafetyCounter()
		if err != nil {
			return 0, err
		}

		// If the program terminated successfully, then no infinite loop has
		// occurred. The program is fixed. Return the accumulator value.
		if p.Terminated {
			return acc, nil
		}

		// Revert the program back to its original state.
		p.Instructions[i] = original
	}

	// No fix could be found using this algorithm.
	return 0, errors.New("no fix found")
}
