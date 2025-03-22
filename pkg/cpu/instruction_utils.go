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
// AbsoluteX: returns the absolute memory address with X offset by adding X register value to base address
// The base address is a 16-bit value stored at PC+1
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

// IndirectY: Implements indirect indexed addressing mode
// Returns the effective address and a bool indicating if a page boundary was crossed
// First gets a zero page address, then reads a 16-bit pointer from that address
// Finally adds Y register to the pointer to get the effective address
func IndirectY(c *CPU) (uint16, bool) {
	zeroPageAddr := uint16(c.Memory.Read(c.PC + 1)) // ✅ Dirección en Zero Page
	baseAddr := c.Memory.ReadWord(zeroPageAddr)     // ✅ Leer puntero de 2 bytes

	effectiveAddr := baseAddr + uint16(c.Y)                        // ✅ Sumar Y al puntero
	pageCrossed := (baseAddr & 0xFF00) != (effectiveAddr & 0xFF00) // ✅ Detectar cruce de página

	return effectiveAddr, pageCrossed
}

// ZeroPage: Returns a memory address in the zero page (first 256 bytes)
// The address is specified by a single byte following the opcode
func ZeroPage(c *CPU) uint16 {
	return ZeroPageMemoryDirection(c) // ✅ Debe devolver la dirección, no el valor.
}

// ZeroPageMemoryDirection: Helper function that returns a zero page memory address
// Reads a single byte following the opcode and returns it as a uint16 address
func ZeroPageMemoryDirection(c *CPU) uint16 {
	return uint16(Immediate(c))
}

// ZeroPageX: Returns a zero page address offset by X register
// Takes the byte following the opcode, adds X register, and wraps to stay in zero page
func ZeroPageX(c *CPU) uint16 {
	address := (uint16(Immediate(c)) + uint16(c.X)) & 0xFF // ✅ Asegurar direccionamiento Zero Page
	return address
}

// CastUint16ToUint8: Safely casts a uint16 to uint8
// Returns the casted value and a boolean indicating if there was an overflow
// This is used when arithmetic operations need to detect carry or overflow conditions
func CastUint16ToUint8(value uint16) (uint8, bool) {
	if value > 0xFF {
		return uint8(value & 0xFF), true
	}
	return uint8(value), false
}
