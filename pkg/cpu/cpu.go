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
	SP uint8  // Stack pointer
	PC uint16 // Program counter

	// Memory interface
	Memory interface {
		Read(address uint16) byte
		Write(address uint16, value byte)
		ReadWord(address uint16) uint16
		WriteWord(address uint16, value uint16)
		ReadAddressIndirectPageBoundaryBug(address uint16) uint16
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
	ReadAddressIndirectPageBoundaryBug(address uint16) uint16
}) {
	c.Memory = memory
}

func (c *CPU) MovePC(offset uint16) {
	c.PC = c.PC + offset
}

// Reset resets the CPU to its initial state
func (c *CPU) Reset() {
	c.A = 0
	c.X = 0
	c.Y = 0
	c.P = 0x24  //TODO: check
	c.SP = 0xFD // TODO: check
	// Read reset vector at 0xFFFC and 0xFFFD
	fmt.Printf("Reset vector: %04X\n", c.Memory.ReadWord(0xFFFC))
	//c.PC = 0xFFFC
	//resetAddress := c.Memory.ReadWord(0xFFFC)
	c.PC = c.Memory.ReadWord(0xFFFC)
}

// GetInstruction returns instruction information for the given opcode
func (c *CPU) GetInstruction(opcode byte) Instruction {
	return GetInstruction(opcode)
}

// Step executes a single CPU instruction
func (c *CPU) Step() (uint8, error) {
	// Read opcode
	var cycles uint8
	fmt.Printf("PC: %02X\n", c.PC)
	opcode := c.Memory.Read(c.PC)

	instruction := GetInstruction(opcode)
	fmt.Printf("Executing opcode: %02X (%s), PC: %02X\n", opcode, instruction.Mnemonic, c.PC)

	// Get the execution function for the instruction
	executeFunc := GetInstructionFunc(opcode)
	if executeFunc != nil {
		cycles = executeFunc(c)
	} else {
		return 0, fmt.Errorf("missing method for instruction opcode: %02X", opcode)
	}

	// Add sleep for debugging/visualization purposes
	time.Sleep(50 * time.Millisecond)

	//c.PC++
	//c.PC = c.PC + 2

	return cycles, nil // Return cycles used and no error
}
