// Package cpu implements the NES CPU (6502) emulation
package cpu

import (
	"fmt"
	"time"
)

// CPU represents the 6502 processor of the NES
type CPU struct {
	// CPU registers
	A  uint8  // Accumulator
	X  uint8  // X index register
	Y  uint8  // Y index register
	P  uint8  // Processor status
	S  uint8  // Stack pointer
	PC uint16 // Program counter

	// Memory interface
	Memory interface {
		Read(address uint16) byte
		Write(address uint16, value byte)
		ReadWord(address uint16) uint16
		WriteWord(address uint16, value uint16)
	}
}

// NewCPU creates a new CPU instance
func NewCPU() *CPU {
	return &CPU{}
}

// SetMemory sets the memory interface for the CPU
func (c *CPU) SetMemory(memory interface {
	Read(address uint16) byte
	Write(address uint16, value byte)
	ReadWord(address uint16) uint16
	WriteWord(address uint16, value uint16)
}) {
	c.Memory = memory
}

// Reset resets the CPU to its initial state
func (c *CPU) Reset() {
	c.A = 0
	c.X = 0
	c.Y = 0
	c.P = 0
	c.S = 0xFD
	// Read reset vector at 0xFFFC and 0xFFFD
	fmt.Printf("Reset vector: %04X\n", c.Memory.ReadWord(0xFFFC))
	c.PC = c.Memory.ReadWord(0xFFFC)
}

// Step executes a single CPU instruction
func (c *CPU) Step() int {
	// Read opcode
	fmt.Printf("PC: %02X\n", c.PC)
	opcode := c.Memory.Read(c.PC)
	c.PC++

	instruction := GetInstruction(opcode)
	fmt.Printf("Executing opcode: %02X (%s), PC: %02X\n", opcode, instruction.Mnemonic, c.PC)

	// Get the execution function for the instruction
	executeFunc := GetInstructionFunc(opcode)
	if executeFunc != nil {
		executeFunc(c)
	} else {
		panic(fmt.Sprintf("Missing method for instruction opcode: %02X", opcode))
		//fmt.Printf("Warning: No execution function for opcode %02X\n", opcode)
	}

	// Add sleep for debugging/visualization purposes
	time.Sleep(1 * time.Second)

	return instruction.Cycles // Return cycles used
}
