// Package cpu implements the NES CPU (6502) emulation
package cpu

import (
	"fmt"
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

	// Interrupt flags
	nmiPending bool // NMI interrupt pending
	irqPending bool // IRQ interrupt pending

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
	c.P = 0x24  // Standard initial value (I flag set)
	c.SP = 0xFD // Standard initial value
	
	// Clear interrupt flags
	c.nmiPending = false
	c.irqPending = false
	
	// Read reset vector at 0xFFFC and 0xFFFD
	fmt.Printf("Reset vector: %04X\n", c.Memory.ReadWord(0xFFFC))
	c.PC = c.Memory.ReadWord(0xFFFC)
}

// TriggerNMI triggers a non-maskable interrupt
func (c *CPU) TriggerNMI() {
	c.nmiPending = true
}

// TriggerIRQ triggers an interrupt request if the interrupt disable flag is not set
func (c *CPU) TriggerIRQ() {
	// Only set the IRQ pending flag if the interrupt disable flag is not set
	if c.GetFlag(FlagI) == 0 {
		c.irqPending = true
	}
}

// handleInterrupts processes any pending interrupts
func (c *CPU) handleInterrupts() uint8 {
	if c.nmiPending {
		c.nmiPending = false
		return c.handleNMI()
	} else if c.irqPending {
		c.irqPending = false
		return c.handleIRQ()
	}
	return 0
}

// handleNMI processes a non-maskable interrupt
func (c *CPU) handleNMI() uint8 {
	// Push PC and status to stack
	c.pushStackWord(c.PC)
	c.pushStack(c.P & 0xEF) // Push P with B flag cleared
	
	// Set interrupt disable flag
	c.setFlag(FlagI, true)
	
	// Load PC from NMI vector
	c.PC = c.Memory.ReadWord(0xFFFA)
	
	return 7 // NMI takes 7 cycles
}

// handleIRQ processes an interrupt request
func (c *CPU) handleIRQ() uint8 {
	// Push PC and status to stack
	c.pushStackWord(c.PC)
	c.pushStack(c.P & 0xEF) // Push P with B flag cleared
	
	// Set interrupt disable flag
	c.setFlag(FlagI, true)
	
	// Load PC from IRQ vector
	c.PC = c.Memory.ReadWord(0xFFFE)
	
	return 7 // IRQ takes 7 cycles
}

// GetInstruction returns instruction information for the given opcode
func (c *CPU) GetInstruction(opcode byte) Instruction {
	return GetInstruction(opcode)
}

// Step executes a single CPU instruction
func (c *CPU) Step() (uint8, error) {
	// Check for interrupts first
	if c.nmiPending || c.irqPending {
		return c.handleInterrupts(), nil
	}
	
	// Read opcode
	var cycles uint8
	opcode := c.Memory.Read(c.PC)

	// Get the execution function for the instruction
	executeFunc := GetInstructionFunc(opcode)
	if executeFunc != nil {
		cycles = executeFunc(c)
	} else {
		return 0, fmt.Errorf("missing method for instruction opcode: %02X", opcode)
	}

	return cycles, nil // Return cycles used and no error
}
