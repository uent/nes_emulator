// Package cpu implements the NES CPU (6502) emulation
package cpu

import "fmt"

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
	fmt.Println(c.Memory)
	c.PC = c.Memory.ReadWord(0xFFFC)
}

// Step executes a single CPU instruction
func (c *CPU) Step() int {
	// Read opcode
	opcode := c.Memory.Read(c.PC)
	c.PC++

	// TODO: Implement instruction decoding and execution
	// Placeholder usage to avoid unused variable error
	_ = opcode

	return 0 // Return cycles used
}
