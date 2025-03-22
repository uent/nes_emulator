// Package cpu implements the NES CPU (6502) instruction functions
package cpu

// Immediate: value is in c.PC + 1
func Immediate(c *CPU) uint8 {
	return c.Memory.Read(c.PC + 1)
}

// Absolute: value in the memory direction found in c.PC + 1 (2 bytes)
func AbsoluteMemoryDirection(c *CPU) uint16 {
	return c.Memory.ReadWord(c.PC + 1)
}

// AbsoluteX: value in the memory direction found in c.PC + 1 (2 bytes) + c.X
func AbsoluteXMemoryDirection(c *CPU) uint16 {
	baseAddress := c.Memory.ReadWord(c.PC + 1)
	return baseAddress + uint16(c.X)
}

// indirect: the real memory direction value is in the memory direction found in c.PC + 1 (2 bytes)
func Indirect(c *CPU) uint16 {
	indirect_memory_address := c.Memory.ReadWord(AbsoluteMemoryDirection(c))

	address := c.Memory.ReadAddressIndirectPageBoundaryBug(indirect_memory_address)
	return address
}

func IndirectY(c *CPU) (uint16, bool) {
	zeroPageAddr := uint16(c.Memory.Read(c.PC + 1)) // ✅ Dirección en Zero Page
	baseAddr := c.Memory.ReadWord(zeroPageAddr)     // ✅ Leer puntero de 2 bytes

	effectiveAddr := baseAddr + uint16(c.Y)                        // ✅ Sumar Y al puntero
	pageCrossed := (baseAddr & 0xFF00) != (effectiveAddr & 0xFF00) // ✅ Detectar cruce de página

	return effectiveAddr, pageCrossed
}

func ZeroPage(c *CPU) uint16 {
	return ZeroPageMemoryDirection(c) // ✅ Debe devolver la dirección, no el valor.
}

func ZeroPageMemoryDirection(c *CPU) uint16 {
	return uint16(Immediate(c))
}

func ZeroPageX(c *CPU) uint16 {
	address := (uint16(Immediate(c)) + uint16(c.X)) & 0xFF // ✅ Asegurar direccionamiento Zero Page
	return address
}

// return the casted value and a boolean indicating if there was an overflow
func CastUint16ToUint8(value uint16) (uint8, bool) {
	if value > 0xFF {
		return uint8(value & 0xFF), true
	}
	return uint8(value), false
}
