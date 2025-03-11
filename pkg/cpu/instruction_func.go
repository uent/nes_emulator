// Package cpu implements the NES CPU (6502) instruction functions
package cpu

// This file contains implementations for all 6502 CPU instructions

// Instructions for loading and storing

// LDA - Load Accumulator
// not tested
func LDA(c *CPU, address uint16) {
	c.A = c.Memory.Read(address)
	c.UpdateZN(c.A)
}

// LDX - Load X Register
// not tested
func LDX(c *CPU, address uint16) {
	c.X = c.Memory.Read(address)
	c.UpdateZN(c.X)
}

// LDY - Load Y Register
// not tested
func LDY(c *CPU, address uint16) {
	c.Y = c.Memory.Read(address)
	c.UpdateZN(c.Y)
}

// STA - Store Accumulator
// not tested
func STA(c *CPU, address uint16) {
	c.Memory.Write(address, c.A)
}

// STX - Store X Register
// not tested
func STX(c *CPU, address uint16) {
	c.Memory.Write(address, c.X)
}

// STY - Store Y Register
// not tested
func STY(c *CPU, address uint16) {
	c.Memory.Write(address, c.Y)
}

// Instructions for stack operations

// PHA - Push Accumulator
// not tested
func PHA(c *CPU) {
	c.pushStack(c.A)
}

// PHP - Push Processor Status
// not tested
func PHP(c *CPU) {
	c.pushStack(c.P | 0x10) // Set B flag when pushed
}

// PLA - Pull Accumulator
// not tested
func PLA(c *CPU) {
	c.A = c.pullStack()
	c.UpdateZN(c.A)
}

// PLP - Pull Processor Status
// not tested
func PLP(c *CPU) {
	c.P = (c.pullStack() & 0xEF) | 0x20 // Clear B flag, set bit 5
}

// Instructions for arithmetic operations

// ADC - Add with Carry
// not tested
func ADC(c *CPU, address uint16) {
	value := c.Memory.Read(address)
	result := uint16(c.A) + uint16(value) + uint16(c.getFlag(FlagC))

	// Set carry flag
	c.setFlag(FlagC, result > 0xFF)

	// Set overflow flag
	overflow := (c.A^value)&0x80 == 0 && (c.A^uint8(result))&0x80 != 0
	c.setFlag(FlagV, overflow)

	// Set the accumulator
	c.A = uint8(result)
	c.UpdateZN(c.A)
}

// SBC - Subtract with Carry
// not tested
func SBC(c *CPU, address uint16) {
	value := c.Memory.Read(address)
	result := uint16(c.A) - uint16(value) - uint16(1-c.getFlag(FlagC))

	// Set carry flag (note: inverted logic compared to ADC)
	c.setFlag(FlagC, result < 0x100)

	// Set overflow flag
	overflow := (c.A^value)&0x80 != 0 && (c.A^uint8(result))&0x80 != 0
	c.setFlag(FlagV, overflow)

	// Set the accumulator
	c.A = uint8(result)
	c.UpdateZN(c.A)
}

// Instructions for increments and decrements

// INC - Increment Memory
// not tested
func INC(c *CPU, address uint16) {
	value := c.Memory.Read(address) + 1
	c.Memory.Write(address, value)
	c.UpdateZN(value)
}

// INX - Increment X Register
// not tested
func INX(c *CPU) {
	c.X++
	c.UpdateZN(c.X)
}

// INY - Increment Y Register
// not tested
func INY(c *CPU) {
	c.Y++
	c.UpdateZN(c.Y)
}

// DEC - Decrement Memory
// not tested
func DEC(c *CPU, address uint16) {
	value := c.Memory.Read(address) - 1
	c.Memory.Write(address, value)
	c.UpdateZN(value)
}

// DEX - Decrement X Register
// not tested
func DEX(c *CPU) {
	c.X--
	c.UpdateZN(c.X)
}

// DEY - Decrement Y Register
// not tested
func DEY(c *CPU) {
	c.Y--
	c.UpdateZN(c.Y)
}

// Instructions for logical operations

// AND - Logical AND
// not tested
func AND(c *CPU, address uint16) {
	c.A &= c.Memory.Read(address)
	c.UpdateZN(c.A)
}

// ORA - Logical OR
// not tested
func ORA(c *CPU, address uint16) {
	c.A |= c.Memory.Read(address)
	c.UpdateZN(c.A)
}

// EOR - Exclusive OR
// not tested
func EOR(c *CPU, address uint16) {
	c.A ^= c.Memory.Read(address)
	c.UpdateZN(c.A)
}

// Instructions for shifts and rotates

// ASL - Arithmetic Shift Left
// not tested
func ASL(c *CPU, address uint16, accumulator bool) {
	var value uint8
	if accumulator {
		value = c.A
		c.setFlag(FlagC, (value&0x80) != 0)
		value <<= 1
		c.A = value
	} else {
		value = c.Memory.Read(address)
		c.setFlag(FlagC, (value&0x80) != 0)
		value <<= 1
		c.Memory.Write(address, value)
	}
	c.UpdateZN(value)
}

// LSR - Logical Shift Right
// not tested
func LSR(c *CPU, address uint16, accumulator bool) {
	var value uint8
	if accumulator {
		value = c.A
		c.setFlag(FlagC, (value&0x01) != 0)
		value >>= 1
		c.A = value
	} else {
		value = c.Memory.Read(address)
		c.setFlag(FlagC, (value&0x01) != 0)
		value >>= 1
		c.Memory.Write(address, value)
	}
	c.UpdateZN(value)
}

// ROL - Rotate Left
// not tested
func ROL(c *CPU, address uint16, accumulator bool) {
	var value uint8
	if accumulator {
		value = c.A
		oldCarry := c.getFlag(FlagC)
		c.setFlag(FlagC, (value&0x80) != 0)
		value = (value << 1) | oldCarry
		c.A = value
	} else {
		value = c.Memory.Read(address)
		oldCarry := c.getFlag(FlagC)
		c.setFlag(FlagC, (value&0x80) != 0)
		value = (value << 1) | oldCarry
		c.Memory.Write(address, value)
	}
	c.UpdateZN(value)
}

// ROR - Rotate Right
// not tested
func ROR(c *CPU, address uint16, accumulator bool) {
	var value uint8
	if accumulator {
		value = c.A
		oldCarry := c.getFlag(FlagC)
		c.setFlag(FlagC, (value&0x01) != 0)
		value = (value >> 1) | (oldCarry << 7)
		c.A = value
	} else {
		value = c.Memory.Read(address)
		oldCarry := c.getFlag(FlagC)
		c.setFlag(FlagC, (value&0x01) != 0)
		value = (value >> 1) | (oldCarry << 7)
		c.Memory.Write(address, value)
	}
	c.UpdateZN(value)
}

// Instructions for comparisons

// CMP - Compare Accumulator
// not tested
func CMP(c *CPU, address uint16) {
	value := c.Memory.Read(address)
	result := c.A - value
	c.setFlag(FlagC, c.A >= value)
	c.UpdateZN(result)
}

// CPX - Compare X Register
// not tested
func CPX(c *CPU, address uint16) {
	value := c.Memory.Read(address)
	result := c.X - value
	c.setFlag(FlagC, c.X >= value)
	c.UpdateZN(result)
}

// CPY - Compare Y Register
// not tested
func CPY(c *CPU, address uint16) {
	value := c.Memory.Read(address)
	result := c.Y - value
	c.setFlag(FlagC, c.Y >= value)
	c.UpdateZN(result)
}

// Instructions for branches

// BCC - Branch if Carry Clear
// not tested
func BCC(c *CPU, offset int8) {
	if c.getFlag(FlagC) == 0 {
		c.PC = uint16(int32(c.PC) + int32(offset))
	}
}

// BCS - Branch if Carry Set
// not tested
func BCS(c *CPU, offset int8) {
	if c.getFlag(FlagC) == 1 {
		c.PC = uint16(int32(c.PC) + int32(offset))
	}
}

// BEQ - Branch if Equal (Z=1)
// not tested
func BEQ(c *CPU, offset int8) {
	if c.getFlag(FlagZ) == 1 {
		c.PC = uint16(int32(c.PC) + int32(offset))
	}
}

// BNE - Branch if Not Equal (Z=0)
// not tested
func BNE(c *CPU, offset int8) {
	if c.getFlag(FlagZ) == 0 {
		c.PC = uint16(int32(c.PC) + int32(offset))
	}
}

// BMI - Branch if Minus (N=1)
// not tested
func BMI(c *CPU, offset int8) {
	if c.getFlag(FlagN) == 1 {
		c.PC = uint16(int32(c.PC) + int32(offset))
	}
}

// BPL - Branch if Plus (N=0)
// not tested
func BPL(c *CPU, offset int8) {
	if c.getFlag(FlagN) == 0 {
		c.PC = uint16(int32(c.PC) + int32(offset))
	}
}

// BVC - Branch if Overflow Clear
// not tested
func BVC(c *CPU, offset int8) {
	if c.getFlag(FlagV) == 0 {
		c.PC = uint16(int32(c.PC) + int32(offset))
	}
}

// BVS - Branch if Overflow Set
// not tested
func BVS(c *CPU, offset int8) {
	if c.getFlag(FlagV) == 1 {
		c.PC = uint16(int32(c.PC) + int32(offset))
	}
}

// Instructions for jumps and subroutines

// JMP - Jump
// not tested
func JMP(c *CPU, address uint16) {
	c.PC = address
}

// JSR - Jump to Subroutine
// not tested
func JSR(c *CPU, address uint16) {
	// Push return address (PC-1) to stack
	c.pushStackWord(c.PC - 1)
	c.PC = address
}

// RTS - Return from Subroutine
// not tested
func RTS(c *CPU) {
	c.PC = c.pullStackWord() + 1
}

// RTI - Return from Interrupt
// not tested
func RTI(c *CPU) {
	c.P = (c.pullStack() & 0xEF) | 0x20 // Clear B flag, set bit 5
	c.PC = c.pullStackWord()
}

// Flag instructions

// CLC - Clear Carry Flag
// not tested
func CLC(c *CPU) {
	c.setFlag(FlagC, false)
}

// CLD - Clear Decimal Mode
// not tested
func CLD(c *CPU) {
	c.setFlag(FlagD, false)
}

// CLI - Clear Interrupt Disable
// not tested
func CLI(c *CPU) {
	c.setFlag(FlagI, false)
}

// CLV - Clear Overflow Flag
// not tested
func CLV(c *CPU) {
	c.setFlag(FlagV, false)
}

// SEC - Set Carry Flag
// not tested
func SEC(c *CPU) {
	c.setFlag(FlagC, true)
}

// SED - Set Decimal Flag
// not tested
func SED(c *CPU) {
	c.setFlag(FlagD, true)
}

// SEI - Set Interrupt Disable
// not tested
func SEI(c *CPU) {
	c.setFlag(FlagI, true)
}

// Miscellaneous instructions

// BRK - Force Interrupt
// not tested
func BRK(c *CPU) {
	c.pushStackWord(c.PC + 2)
	c.pushStack(c.P | 0x10) // Set B flag when pushed
	c.setFlag(FlagI, true)
	//c.PC = c.Memory.ReadWord(0xFFFE)
}

// NOP - No Operation
// not tested
func NOP(c *CPU) {
	// Do nothing
}

// Transfers

// TAX - Transfer A to X
// not tested
func TAX(c *CPU) {
	c.X = c.A
	c.UpdateZN(c.X)
}

// TAY - Transfer A to Y
// not tested
func TAY(c *CPU) {
	c.Y = c.A
	c.UpdateZN(c.Y)
}

// TSX - Transfer Stack Pointer to X
// not tested
func TSX(c *CPU) {
	c.X = c.S
	c.UpdateZN(c.X)
}

// TXA - Transfer X to A
// not tested
func TXA(c *CPU) {
	c.A = c.X
	c.UpdateZN(c.A)
}

// TXS - Transfer X to Stack Pointer
// not tested
func TXS(c *CPU) {
	c.S = c.X
}

// TYA - Transfer Y to A
// not tested
func TYA(c *CPU) {
	c.A = c.Y
	c.UpdateZN(c.A)
}

// Helper methods for the CPU

// Flag bit positions
const (
	FlagC uint8 = 0 // Carry
	FlagZ uint8 = 1 // Zero
	FlagI uint8 = 2 // Interrupt Disable
	FlagD uint8 = 3 // Decimal Mode
	FlagB uint8 = 4 // Break Command
	Flag5 uint8 = 5 // Unused (always 1)
	FlagV uint8 = 6 // Overflow
	FlagN uint8 = 7 // Negative
)

// getFlag returns the value of a specific status flag
func (c *CPU) getFlag(flag uint8) uint8 {
	if c.P&(1<<flag) != 0 {
		return 1
	}
	return 0
}

// setFlag sets or clears a specific status flag
func (c *CPU) setFlag(flag uint8, value bool) {
	if value {
		c.P |= 1 << flag
	} else {
		c.P &= ^(1 << flag)
	}
}

// UpdateZN updates the Zero and Negative flags based on the given value
func (c *CPU) UpdateZN(value uint8) {
	c.setFlag(FlagZ, value == 0)
	c.setFlag(FlagN, value&0x80 != 0)
}

// pushStack pushes a byte onto the stack
func (c *CPU) pushStack(value uint8) {
	c.Memory.Write(0x0100|uint16(c.S), value)
	c.S--
}

// pullStack pulls a byte from the stack
func (c *CPU) pullStack() uint8 {
	c.S++
	return c.Memory.Read(0x0100 | uint16(c.S))
}

// pushStackWord pushes a word onto the stack
func (c *CPU) pushStackWord(value uint16) {
	hi := uint8(value >> 8)
	lo := uint8(value & 0xFF)
	c.pushStack(hi)
	c.pushStack(lo)
}

// pullStackWord pulls a word from the stack
func (c *CPU) pullStackWord() uint16 {
	lo := uint16(c.pullStack())
	hi := uint16(c.pullStack())
	return (hi << 8) | lo
}
