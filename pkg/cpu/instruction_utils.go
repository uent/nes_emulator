// Package cpu implements the NES CPU (6502) instruction functions
package cpu

// Immediate: value is in c.PC + 1
func Immediate(c *CPU) uint8 {
	return c.Memory.Read(c.PC + 1)
}

// Absolute: value in the memory direction found in c.PC + 1 (2 bytes)
func AbsoluteMemoryDirection(c *CPU) uint16 {
	return c.PC + 1
}

// AbsoluteX: value in the memory direction found in c.PC + 1 (2 bytes) + c.X
func AbsoluteX(c *CPU) uint16 {
	return AbsoluteMemoryDirection(c) + uint16(c.X)
}

// indirect: the real memory direction value is in the memory direction found in c.PC + 1 (2 bytes)
func Indirect(c *CPU) uint16 {
	indirect_memory_address := c.Memory.ReadWord(AbsoluteMemoryDirection(c))

	address := c.Memory.ReadAddressIndirectPageBoundaryBug(indirect_memory_address)
	return address
}

func ZeroPage(c *CPU) byte {
	return c.Memory.Read(uint16(Immediate(c)))
}

func ZeroPageX(c *CPU) byte {
	address := (c.X + Immediate(c)) & 0xFF
	return c.Memory.Read(uint16(address))
}

// return the casted value and a boolean indicating if there was an overflow
func CastUint16ToUint8(value uint16) (uint8, bool) {
	if value > 0xFF {
		return uint8(value & 0xFF), true
	}
	return uint8(value), false
}
