// Package cpu implements the NES CPU (6502) instruction functions
package cpu

import "fmt"

// This file contains implementations for all 6502 CPU instructions

type CPUOperation func(*CPU) uint8

// LDA - Load Accumulator
// not tested
func LDA(c *CPU, value uint8) {
	c.A = value
	c.setFlagZByValue(c.A)
	c.setFlagNByValue(c.A)
}

func LDAImmediate(c *CPU) uint8 {
	value := Immediate(c)

	LDA(c, value)

	c.MovePC(2)
	return 2 // cycles 2
}

func LDAAbsolute(c *CPU) uint8 {
	value := c.Memory.Read(AbsoluteMemoryDirection(c))

	LDA(c, value)

	c.MovePC(3)
	return 4 // cycles 4
}

func LDAAbsoluteX(c *CPU) uint8 {
	value := c.Memory.Read(AbsoluteXMemoryDirection(c))

	LDA(c, value)

	c.MovePC(3)
	return 4 // cycles 4 (+1 if page is crossed)
}

// LDX - Load X Register
// not tested
/* func LDX(c *CPU, address uint16) {
	c.X = c.Memory.Read(address)
	c.setFlagZByValue(c.X)
	c.setFlagNByValue(c.X)
} */

func LDXImmediate(c *CPU) uint8 {
	c.X = Immediate(c)

	c.setFlagZByValue(c.X)
	c.setFlagNByValue(c.X)

	c.MovePC(2)

	return 2 // cycles 2
}

// LDY - Load Y Register
// not tested
/* func LDY(c *CPU, address uint16) {
	c.Y = c.Memory.Read(address)
	c.setFlagZByValue(c.Y)
	c.setFlagNByValue(c.Y)
} */

func LDYImmediate(c *CPU) uint8 {
	c.Y = Immediate(c)
	c.setFlagZByValue(c.Y)
	c.setFlagNByValue(c.Y)

	c.MovePC(2)

	return 2 // cycles 2
}

// STA - Store Accumulator
func STA(c *CPU, address uint16) {
	c.Memory.Write(address, c.A)
}

func STAZeroPage(c *CPU) uint8 {
	address := ZeroPageMemoryDirection(c)
	STA(c, address)

	c.MovePC(2)

	return 3 // cicles 3
}

func STAAbsolute(c *CPU) uint8 {
	address := AbsoluteMemoryDirection(c)
	STA(c, address)

	c.MovePC(3)

	return 4 // cicles 4
}

func STAAbsoluteX(c *CPU) uint8 {
	address := AbsoluteXMemoryDirection(c)

	STA(c, address)

	c.MovePC(3)

	return 5 // cicles 5
}

func STAIndirectY(c *CPU) uint8 {
	address, _ := IndirectY(c)

	STA(c, address)

	c.MovePC(2)

	return 6 // cicles 6
}

// STX - Store X Register
// not tested
/* func STX(c *CPU, address uint16) {
	c.Memory.Write(address, c.X)
} */

func STXZeroPage(c *CPU) uint8 {
	address := uint16(ZeroPage(c))
	c.Memory.Write(address, c.X)

	c.MovePC(2)

	return 3 // 3 cycles
}

// STY - Store Y Register
// not tested
/* func STY(c *CPU, address uint16) {
	c.Memory.Write(address, c.Y)
} */

// Instructions for stack operations

// PHA - Push Accumulator
/* func PHA(c *CPU) {
	c.pushStack(c.A)
} */

func PHAImplied(c *CPU) uint8 {
	c.pushStack(c.A)
	c.MovePC(1)
	return 3 // cycles 3
}

// PHP - Push Processor Status
// not tested
/* func PHP(c *CPU) {
	c.pushStack(c.P | 0x10) // Set B flag when pushed
} */

// PLA - Pull Accumulator
func PLAImplied(c *CPU) uint8 {
	c.A = c.pullStack()
	c.setFlagZByValue(c.A)
	c.setFlagNByValue(c.A)

	c.MovePC(1)

	return 4 // 4 cycles
}

// PLP - Pull Processor Status
// not tested
/* func PLP(c *CPU) {
	c.P = (c.pullStack() & 0xEF) | 0x20 // Clear B flag, set bit 5
} */

func PLPImplied(c *CPU) uint8 {
	value := c.pullStack()

	c.setFlagC((value >> 0 & 1) == 1)   // bit 0
	c.setFlagZ((value >> 1 & 1) == 1)   // bit 1
	c.setFlagI((value>>2&1) == 1, true) //bit 2, TODO: delayed 1 instrucion
	c.setFlagD((value >> 3 & 1) == 1)   // bit 3
	c.setFlagV((value >> 6 & 1) == 1)   // bit 6
	c.setFlagN((value >> 7 & 1) == 1)   // bit 7

	c.MovePC(1)

	return 4 // cycles 4
}

// Instructions for arithmetic operations

// ADC - Add with Carry
// not tested
/* func ADC(c *CPU, address uint16) {
	value := c.Memory.Read(address)
	result := uint16(c.A) + uint16(value) + uint16(c.GetFlag(FlagC))

	// Set carry flag
	c.setFlag(FlagC, result > 0xFF)

	// Set overflow flag
	overflow := (c.A^value)&0x80 == 0 && (c.A^uint8(result))&0x80 != 0
	c.setFlag(FlagV, overflow)

	// Set the accumulator
	c.A = uint8(result)
	c.setFlagZByValue(c.A)
	c.setFlagNByValue(c.A)
} */

func ADCZeroPageX(c *CPU) uint8 {
	memoryValue := ZeroPageX(c)

	value := uint16(c.A) + uint16(memoryValue) + uint16(c.GetFlagC())
	cast_value, overflow := CastUint16ToUint8(value)

	if overflow {
		c.setFlagC(true)
	}

	overflowFlag := ((cast_value ^ c.A) & (cast_value ^ uint8(memoryValue)) & 0x80) != 0 // TODO: check logic

	c.setFlag(FlagV, overflowFlag)
	c.setFlagZByValue(cast_value)
	c.setFlagNByValue(cast_value)

	c.A = cast_value

	c.MovePC(2)
	return 4 // cycles 4
}

// SBC - Subtract with Carry
// not tested
/* func SBC(c *CPU, address uint16) {
	value := c.Memory.Read(address)
	result := uint16(c.A) - uint16(value) - uint16(1-c.GetFlag(FlagC))

	// Set carry flag (note: inverted logic compared to ADC)
	c.setFlag(FlagC, result < 0x100)

	// Set overflow flag
	overflow := (c.A^value)&0x80 != 0 && (c.A^uint8(result))&0x80 != 0
	c.setFlag(FlagV, overflow)

	// Set the accumulator
	c.A = uint8(result)
	c.setFlagZByValue(c.A)
	c.setFlagNByValue(c.A)
} */

// Instructions for increments and decrements

// INC - Increment Memory
// not tested
/* func INC(c *CPU, address uint16) {
	value := c.Memory.Read(address) + 1
	c.Memory.Write(address, value)
	c.setFlagZByValue(value)
	c.setFlagNByValue(value)
} */

// INX - Increment X Register
func INXImplied(c *CPU) uint8 {
	c.X++

	c.setFlagZByValue(c.X)
	c.setFlagNByValue(c.X)

	c.MovePC(1)

	return 2 // 2 cycles
}

// INY - Increment Y Register
// not tested
/* func INY(c *CPU) {
	c.Y++
	c.setFlagZByValue(c.Y)
	c.setFlagNByValue(c.Y)
} */

// DEC - Decrement Memory
// not tested
/* func DEC(c *CPU, address uint16) {
	value := c.Memory.Read(address) - 1
	c.Memory.Write(address, value)
	c.setFlagZByValue(value)
	c.setFlagNByValue(value)
} */

// DEX - Decrement X Register
// not tested
/* func DEX(c *CPU) {
	c.X--
	c.setFlagZByValue(c.X)
	c.setFlagNByValue(c.X)
} */

func DEXImplied(c *CPU) uint8 {
	c.X--
	c.setFlagZByValue(c.X)
	c.setFlagNByValue(c.X)

	c.MovePC(1)

	return 2 // 2 cycles
}

// DEY - Decrement Y Register
// not tested
/* func DEY(c *CPU) {
	c.Y--
	c.setFlagZByValue(c.Y)
	c.setFlagNByValue(c.Y)
} */

func DEYImplied(c *CPU) uint8 {
	c.Y--
	c.setFlagZByValue(c.Y)
	c.setFlagNByValue(c.Y)

	c.MovePC(1)

	return 2 // 2 cycles
}

// Instructions for logical operations

// AND - Logical AND
/* func AND(c *CPU, address uint16) {
	c.A &= c.Memory.Read(address)
	c.setFlagZByValue(c.A)
	c.setFlagNByValue(c.A)
} */

func ANDImmediate(c *CPU) uint8 {
	value := Immediate(c)

	c.A = c.A & value
	c.setFlagZByValue(c.A)
	c.setFlagNByValue(c.A)

	c.MovePC(2)

	return 2 // cycles 2
}

// ORA - Logical OR
// not tested
/* func ORA(c *CPU, address uint16) {
	c.A |= c.Memory.Read(address)
	c.setFlagZByValue(c.A)
	c.setFlagNByValue(c.A)
} */

// EOR - Exclusive OR
// not tested
/* func EOR(c *CPU, address uint16) {
	c.A ^= c.Memory.Read(address)
	c.setFlagZByValue(c.A)
	c.setFlagNByValue(c.A)
} */

// Instructions for shifts and rotates

// ASL - Arithmetic Shift Left
// not tested
/* func ASL(c *CPU, address uint16, accumulator bool) {
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
	c.setFlagZByValue(value)
	c.setFlagNByValue(value)
} */

// LSR - Logical Shift Right
// not tested
/* func LSR(c *CPU, address uint16, accumulator bool) {
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
	c.setFlagZByValue(value)
	c.setFlagNByValue(value)
} */

func LSRAccumulator(c *CPU) uint8 {
	off_bit := c.A & 0x01
	c.A = c.A >> 1

	c.setFlagC(off_bit == 1)
	c.setFlagZByValue(c.A)
	c.setFlagNByValue(0)

	c.MovePC(1)

	return 2 // cycles 2
}

// ROL - Rotate Left
// not tested
/* func ROL(c *CPU, address uint16, accumulator bool) {
	var value uint8
	if accumulator {
		value = c.A
		oldCarry := c.GetFlag(FlagC)
		c.setFlag(FlagC, (value&0x80) != 0)
		value = (value << 1) | oldCarry
		c.A = value
	} else {
		value = c.Memory.Read(address)
		oldCarry := c.GetFlag(FlagC)
		c.setFlag(FlagC, (value&0x80) != 0)
		value = (value << 1) | oldCarry
		c.Memory.Write(address, value)
	}
	c.setFlagZByValue(value)
	c.setFlagNByValue(value)
} */

// ROR - Rotate Right
// not tested
/* func ROR(c *CPU, address uint16, accumulator bool) {
	var value uint8
	if accumulator {
		value = c.A
		oldCarry := c.GetFlag(FlagC)
		c.setFlag(FlagC, (value&0x01) != 0)
		value = (value >> 1) | (oldCarry << 7)
		c.A = value
	} else {
		value = c.Memory.Read(address)
		oldCarry := c.GetFlag(FlagC)
		c.setFlag(FlagC, (value&0x01) != 0)
		value = (value >> 1) | (oldCarry << 7)
		c.Memory.Write(address, value)
	}
	c.setFlagZByValue(value)
	c.setFlagNByValue(value)
} */

// Instructions for comparisons

// CMP - Compare Accumulator
// not tested
/* func CMP(c *CPU, address uint16) {
	value := c.Memory.Read(address)
	result := c.A - value
	c.setFlag(FlagC, c.A >= value)
	c.setFlagZByValue(result)
	c.setFlagNByValue(result)
} */

// CPX - Compare X Register
// not tested
/* func CPX(c *CPU, address uint16) {
	value := c.Memory.Read(address)
	result := c.X - value
	c.setFlag(FlagC, c.X >= value)
	c.setFlagZByValue(result)
	c.setFlagNByValue(result)
} */

func CPXZeroPage(c *CPU) uint8 {
	value := ZeroPage(c)

	result := int16(c.X) - int16(value)

	c.setFlagC(result >= 0)
	c.setFlagZByValue(uint8(result & 0xFF))
	c.setFlagNByValue(uint8(result & 0xFF))

	c.MovePC(2)

	return 3 // cycles 3
}

// CPY - Compare Y Register
// not tested
/* func CPY(c *CPU, address uint16) {
	value := c.Memory.Read(address)
	result := c.Y - value
	c.setFlag(FlagC, c.Y >= value)
	c.setFlagZByValue(result)
	c.setFlagNByValue(result)
} */

// Instructions for branches

// BCC - Branch if Carry Clear
// not tested
/* func BCC(c *CPU, offset int8) {
	if c.GetFlag(FlagC) == 0 {
		c.MovePC(uint16(int32(offset)))
	}
} */

// BCS - Branch if Carry Set
// not tested
/* func BCS(c *CPU, offset int8) {
	if c.GetFlag(FlagC) == 1 {
		c.MovePC(uint16(int32(offset)))
	}
} */

// BEQ - Branch if Equal (Z=1)
// not tested
/* func BEQ(c *CPU, offset int8) {
	if c.GetFlag(FlagZ) == 1 {
		c.MovePC(uint16(int32(offset)))
	}
} */

func BEQRelative(c *CPU) uint8 {
	offSet := int8(c.Memory.Read(c.PC + 1)) // Leer el offset como int8
	cycles := uint8(2)                      // Siempre consume al menos 2 ciclos

	if c.GetFlagZ() == 1 { // Si Z == 1, se ejecuta el salto
		oldPC := c.PC + 2 // La dirección de la siguiente instrucción
		fmt.Println("offSet", offSet)
		// Calcular la nueva dirección
		c.PC = uint16(int32(oldPC) + int32(offSet))
		cycles++ // Un ciclo adicional por el salto

		// Si el salto cruza una página, agregar un ciclo extra
		if (oldPC & 0xFF00) != (c.PC & 0xFF00) {
			cycles++
		}
	} else {
		c.PC += 2 // Si no salta, solo avanza el PC 2 bytes
	}

	return cycles
}

// BIT - Bit Test
func BITZero(c *CPU) uint8 {
	memory_value := c.Memory.Read(c.PC + 1)
	result := c.A & memory_value

	c.setFlagZByValue(result)
	c.setFlagNByValue(memory_value)
	c.setFlagVByValue(memory_value)

	c.MovePC(2)
	return 3 // cycles 3
}

// BNE - Branch if Not Equal (Z=0)
func BNERelative(c *CPU) uint8 {
	offSet := int16(int8(c.Memory.Read(c.PC + 1))) // Leer el offset como int8
	cycles := uint8(2)                             // Siempre consume al menos 2 ciclos

	if c.GetFlagZ() == 0 { // Si Z == 0, se ejecuta el salto
		oldPC := c.PC + 2 // La dirección de la siguiente instrucción

		// Calcular la nueva dirección
		c.PC = uint16(int16(oldPC) + offSet)
		cycles++ // Un ciclo adicional por el salto

		// Si el salto cruza una página, agregar un ciclo extra
		if (oldPC & 0xFF00) != (c.PC & 0xFF00) {
			cycles++
		}
	} else {
		c.PC += 2 // Si no salta, solo avanza el PC 2 bytes
	}

	return cycles
}

// BMI - Branch if Minus (N=1)
/* func BMI(c *CPU, offset int8) {
	if c.GetFlag(FlagN) == 1 {
		c.MovePC(uint16(int32(offset)))
	}
} */

// BPL - Branch if Plus (N=0)
// not tested
/* func BPL(c *CPU, offset int8) {
	if c.GetFlag(FlagN) == 0 {
		c.MovePC(uint16(int32(offset)))
	}
} */

func BPLRelative(c *CPU) uint8 {
	offSet := int8(c.Memory.Read(c.PC + 1))
	cycles := uint8(2)

	if c.GetFlagN() == 0 {
		oldPC := c.PC + 2
		c.PC = uint16(int32(oldPC) + int32(offSet))
		cycles++

		// ✅ Si el salto cruza una página, agregar un ciclo extra
		if (oldPC & 0xFF00) != (c.PC & 0xFF00) {
			cycles++
		}
	} else {
		c.MovePC(2)
	}

	return cycles
}

// BVC - Branch if Overflow Clear
// not tested
/* func BVC(c *CPU, offset int8) {
	if c.GetFlag(FlagV) == 0 {
		c.MovePC(uint16(int32(offset)))
	}
} */

// BVS - Branch if Overflow Set
// not tested
/* func BVS(c *CPU, offset int8) {
	if c.GetFlag(FlagV) == 1 {
		c.MovePC(uint16(int32(offset)))
	}
} */

// Instructions for jumps and subroutines

// JMP - Jump
func JMP(c *CPU, address uint16) {
	c.PC = c.Memory.ReadWord(address)
}

func JMPIndirect(c *CPU) uint8 {
	address := Indirect(c)

	JMP(c, address)

	return 5 // 5 cyles
}

func JMPAbsolute(c *CPU) uint8 {
	address := AbsoluteMemoryDirection(c)

	JMP(c, address)

	return 3 // 3 cyles
}

// JSR - Jump to Subroutine
// not tested
/* func JSR(c *CPU, address uint16) {
	// Push return address (PC-1) to stack
	c.pushStackWord(c.PC - 1)
	c.MovePC(address - c.PC)
} */

func JSRAbsolute(c *CPU) uint8 {
	address := AbsoluteMemoryDirection(c)
	// a real cpu left the PC in the position PC + 1
	c.pushStackWord(c.PC + 2 - 1) //TODO: check this
	c.PC = address

	return 6 // cycles 6
}

// RTS - Return from Subroutine
// not tested
/* func RTS(c *CPU) {
	pulled := c.pullStackWord() + 1
	c.MovePC(pulled - c.PC)
} */

// RTI - Return from Interrupt
// not tested
/* func RTI(c *CPU) {
	c.P = (c.pullStack() & 0xEF) | 0x20 // Clear B flag, set bit 5
	pulled := c.pullStackWord()
	c.MovePC(pulled - c.PC)
} */

// Flag instructions

// CLC - Clear Carry Flag
// not tested
/* func CLC(c *CPU) {
	c.setFlag(FlagC, false)
} */

// CLD - Clear Decimal Mode
// not tested
/* func CLD(c *CPU) {
	c.setFlagD(false)
} */

func CLDImplied(c *CPU) uint8 {
	c.setFlagD(false)

	c.MovePC(1)

	return 2 // cycles 2
}

// CLI - Clear Interrupt Disable
// not tested
/* func CLI(c *CPU) {
	c.setFlag(FlagI, false)
} */

// CLV - Clear Overflow Flag
// not tested
/* func CLV(c *CPU) {
	c.setFlag(FlagV, false)
} */

// SEC - Set Carry Flag
// not tested
/* func SEC(c *CPU) {
	c.setFlag(FlagC, true)
} */

// SED - Set Decimal Flag
// not tested
/* func SED(c *CPU) {
	c.setFlag(FlagD, true)
} */

// SEI - Set Interrupt Disable
// not tested
/* func SEI(c *CPU) {
	c.setFlagI(true, true)
} */

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
/* func NOP(c *CPU) {
	// Do nothing
} */

func NOPImplied(c *CPU) uint8 {
	// Do nothing

	c.MovePC(1)

	return 2 // 2 cyles
}

// Transfers

// TAX - Transfer A to X
// not tested
/* func TAX(c *CPU) {
	c.X = c.A
	c.setFlagZByValue(c.X)
	c.setFlagNByValue(c.X)
} */

func TAXImpplied(c *CPU) uint8 {
	c.X = c.A

	c.setFlagZByValue(c.X)
	c.setFlagNByValue(c.X)

	c.MovePC(1)

	return 2 // cycles 2
}

// TAY - Transfer A to Y
// not tested
/* func TAY(c *CPU) {
	c.Y = c.A
	c.setFlagZByValue(c.Y)
	c.setFlagNByValue(c.Y)
} */

// TSX - Transfer Stack Pointer to X
// not tested
/* func TSX(c *CPU) {
	c.X = c.SP
	c.setFlagZByValue(c.X)
	c.setFlagNByValue(c.X)
} */

// TXA - Transfer X to A
// not tested
/* func TXA(c *CPU) {
	c.A = c.X
	c.setFlagZByValue(c.A)
	c.setFlagNByValue(c.A)
} */

// TXS - Transfer X to Stack Pointer
// not tested
/* func TXS(c *CPU) {
	c.SP = c.X
} */

func TXSImplied(c *CPU) uint8 {
	c.SP = c.X

	c.MovePC(1)
	return 2 // cycles 2
}

// TYA - Transfer Y to A
// not tested
/* func TYA(c *CPU) {
	c.A = c.Y
	c.setFlagZByValue(c.A)
	c.setFlagNByValue(c.A)
} */
