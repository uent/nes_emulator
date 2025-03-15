// Package cpu implements the NES CPU (6502) instruction functions
package cpu

// This file contains implementations for all 6502 CPU instructions

type CPUOperation func(*CPU) uint8

// Instructions for loading and storing

// Immediate: value is in c.PC + 1
func Immediate(c *CPU) uint8 {
	return c.Memory.Read(c.PC + 1)
}

// Absolute: value in the memory direction found in c.PC + 1 (2 bytes)
func AbsoluteMemoryDirection(c *CPU) uint16 {
	return c.PC + 1
}

// AbsoluteX: value in the memory direction found in c.PC + 1 (2 bytes) + c.X
func AbsoluteX(c *CPU) byte {
	return c.Memory.Read(AbsoluteMemoryDirection(c) + uint16(c.X))
}

// LDA - Load Accumulator
// not tested
func LDA(c *CPU, address uint16) {
	c.A = c.Memory.Read(address)
	c.setFlagZ(c.A)
	c.setFlagN(c.A)
}

func LDAImmediate(c *CPU) uint8 {
	c.A = Immediate(c)

	c.setFlagZ(c.A)
	c.setFlagN(c.A)

	c.MovePC(2)
	return 2 // cycles 2
}

func LDAAbsoluteX(c *CPU) uint8 {
	c.A = AbsoluteX(c)

	c.setFlagZ(c.A)
	c.setFlagN(c.A)

	c.MovePC(3)
	return 4 // cycles 4 (+1 if page is crossed)
}

// LDX - Load X Register
// not tested
func LDX(c *CPU, address uint16) {
	c.X = c.Memory.Read(address)
	c.setFlagZ(c.X)
	c.setFlagN(c.X)
}

// LDY - Load Y Register
// not tested
func LDY(c *CPU, address uint16) {
	c.Y = c.Memory.Read(address)
	c.setFlagZ(c.Y)
	c.setFlagN(c.Y)
}

// STA - Store Accumulator
// not tested
func STA(c *CPU, address uint16) {
	c.Memory.Write(address, c.A)
}

func STAAbsolute(c *CPU) uint8 {
	c.Memory.Write(AbsoluteMemoryDirection(c), c.A)
	c.MovePC(3)

	return 3 // cicles 3
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
func PHA(c *CPU) {
	c.pushStack(c.A)
}

func PHAImplied(c *CPU) uint8 {
	c.pushStack(c.A)
	c.MovePC(1)
	return 3 // cycles 3
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
	c.setFlagZ(c.A)
	c.setFlagN(c.A)
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
	c.setFlagZ(c.A)
	c.setFlagN(c.A)
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
	c.setFlagZ(c.A)
	c.setFlagN(c.A)
}

// Instructions for increments and decrements

// INC - Increment Memory
// not tested
func INC(c *CPU, address uint16) {
	value := c.Memory.Read(address) + 1
	c.Memory.Write(address, value)
	c.setFlagZ(value)
	c.setFlagN(value)
}

// INX - Increment X Register
// not tested
func INX(c *CPU) {
	c.X++
	c.setFlagZ(c.X)
	c.setFlagN(c.X)
}

// INY - Increment Y Register
// not tested
func INY(c *CPU) {
	c.Y++
	c.setFlagZ(c.Y)
	c.setFlagN(c.Y)
}

// DEC - Decrement Memory
// not tested
func DEC(c *CPU, address uint16) {
	value := c.Memory.Read(address) - 1
	c.Memory.Write(address, value)
	c.setFlagZ(value)
	c.setFlagN(value)
}

// DEX - Decrement X Register
// not tested
func DEX(c *CPU) {
	c.X--
	c.setFlagZ(c.X)
	c.setFlagN(c.X)
}

// DEY - Decrement Y Register
// not tested
func DEY(c *CPU) {
	c.Y--
	c.setFlagZ(c.Y)
	c.setFlagN(c.Y)
}

// Instructions for logical operations

// AND - Logical AND
// not tested
func AND(c *CPU, address uint16) {
	c.A &= c.Memory.Read(address)
	c.setFlagZ(c.A)
	c.setFlagN(c.A)
}

// ORA - Logical OR
// not tested
func ORA(c *CPU, address uint16) {
	c.A |= c.Memory.Read(address)
	c.setFlagZ(c.A)
	c.setFlagN(c.A)
}

// EOR - Exclusive OR
// not tested
func EOR(c *CPU, address uint16) {
	c.A ^= c.Memory.Read(address)
	c.setFlagZ(c.A)
	c.setFlagN(c.A)
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
	c.setFlagZ(value)
	c.setFlagN(value)
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
	c.setFlagZ(value)
	c.setFlagN(value)
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
	c.setFlagZ(value)
	c.setFlagN(value)
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
	c.setFlagZ(value)
	c.setFlagN(value)
}

// Instructions for comparisons

// CMP - Compare Accumulator
// not tested
func CMP(c *CPU, address uint16) {
	value := c.Memory.Read(address)
	result := c.A - value
	c.setFlag(FlagC, c.A >= value)
	c.setFlagZ(result)
	c.setFlagN(result)
}

// CPX - Compare X Register
// not tested
func CPX(c *CPU, address uint16) {
	value := c.Memory.Read(address)
	result := c.X - value
	c.setFlag(FlagC, c.X >= value)
	c.setFlagZ(result)
	c.setFlagN(result)
}

// CPY - Compare Y Register
// not tested
func CPY(c *CPU, address uint16) {
	value := c.Memory.Read(address)
	result := c.Y - value
	c.setFlag(FlagC, c.Y >= value)
	c.setFlagZ(result)
	c.setFlagN(result)
}

// Instructions for branches

// BCC - Branch if Carry Clear
// not tested
func BCC(c *CPU, offset int8) {
	if c.getFlag(FlagC) == 0 {
		c.MovePC(uint16(int32(offset)))
	}
}

// BCS - Branch if Carry Set
// not tested
func BCS(c *CPU, offset int8) {
	if c.getFlag(FlagC) == 1 {
		c.MovePC(uint16(int32(offset)))
	}
}

// BEQ - Branch if Equal (Z=1)
// not tested
func BEQ(c *CPU, offset int8) {
	if c.getFlag(FlagZ) == 1 {
		c.MovePC(uint16(int32(offset)))
	}
}

// BIT - Bit Test
func BITZero(c *CPU) uint8 {
	memory_value := c.Memory.Read(c.PC + 1)
	result := c.A & memory_value

	c.setFlagZ(result)
	c.setFlagN(memory_value)
	c.setFlagV(memory_value)

	c.MovePC(2)
	return 3 // cycles 3
}

// BNE - Branch if Not Equal (Z=0)
func BNE(c *CPU, offset int8) {
	if c.getFlag(FlagZ) == 0 {
		c.MovePC(uint16(int32(offset)))
	}
}

// BMI - Branch if Minus (N=1)
func BMI(c *CPU, offset int8) {
	if c.getFlag(FlagN) == 1 {
		c.MovePC(uint16(int32(offset)))
	}
}

// BPL - Branch if Plus (N=0)
// not tested
func BPL(c *CPU, offset int8) {
	if c.getFlag(FlagN) == 0 {
		c.MovePC(uint16(int32(offset)))
	}
}

// BVC - Branch if Overflow Clear
// not tested
func BVC(c *CPU, offset int8) {
	if c.getFlag(FlagV) == 0 {
		c.MovePC(uint16(int32(offset)))
	}
}

// BVS - Branch if Overflow Set
// not tested
func BVS(c *CPU, offset int8) {
	if c.getFlag(FlagV) == 1 {
		c.MovePC(uint16(int32(offset)))
	}
}

// Instructions for jumps and subroutines

// JMP - Jump
// not tested
func JMP(c *CPU, address uint16) {
	c.MovePC(address - c.PC)
}

// JSR - Jump to Subroutine
// not tested
func JSR(c *CPU, address uint16) {
	// Push return address (PC-1) to stack
	c.pushStackWord(c.PC - 1)
	c.MovePC(address - c.PC)
}

func JSRAbsolute(c *CPU) uint8 {
	address := AbsoluteMemoryDirection(c)
	// a real cpu left the PC in the position PC + 1
	c.pushStackWord(c.PC + 2 - 1) //TODO: check this
	c.PC = address

	return 6 // cycles 6
}

// RTS - Return from Subroutine
// not tested
func RTS(c *CPU) {
	pulled := c.pullStackWord() + 1
	c.MovePC(pulled - c.PC)
}

// RTI - Return from Interrupt
// not tested
func RTI(c *CPU) {
	c.P = (c.pullStack() & 0xEF) | 0x20 // Clear B flag, set bit 5
	pulled := c.pullStackWord()
	c.MovePC(pulled - c.PC)
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
	c.setFlagI(true, true)
}

func SEIImplied(c *CPU) uint8 {
	c.setFlagI(true, true)

	c.MovePC(1)
	return 2 // cycles 2
}

// Miscellaneous instructions

// BRK - Force Interrupt
// not tested
func BRK(c *CPU) uint8 {
	c.pushStackWord(c.PC + 2)
	c.pushStack(c.P | 0x30) // Set B flag when pushed
	c.setFlagI(true, false)
	c.setFlagB(true)
	c.PC = c.Memory.ReadWord(0xFFFE)

	return 0 // TODO: check
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
	c.setFlagZ(c.X)
	c.setFlagN(c.X)
}

// TAY - Transfer A to Y
// not tested
func TAY(c *CPU) {
	c.Y = c.A
	c.setFlagZ(c.Y)
	c.setFlagN(c.Y)
}

// TSX - Transfer Stack Pointer to X
// not tested
func TSX(c *CPU) {
	c.X = c.SP
	c.setFlagZ(c.X)
	c.setFlagN(c.X)
}

// TXA - Transfer X to A
// not tested
func TXA(c *CPU) {
	c.A = c.X
	c.setFlagZ(c.A)
	c.setFlagN(c.A)
}

// TXS - Transfer X to Stack Pointer
// not tested
func TXS(c *CPU) {
	c.SP = c.X
}

// TYA - Transfer Y to A
// not tested
func TYA(c *CPU) {
	c.A = c.Y
	c.setFlagZ(c.A)
	c.setFlagN(c.A)
}
