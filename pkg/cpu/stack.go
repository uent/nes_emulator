// Package cpu implements the NES CPU (6502) stack operations
package cpu

// pushStack pushes a byte onto the stack and decrements the stack pointer
func (c *CPU) pushStack(value uint8) {
	c.Memory.Write(0x0100+uint16(c.SP), value)
	c.SP--
}

// pullStack increments the stack pointer and pulls a byte from the stack
func (c *CPU) pullStack() uint8 {
	c.SP++
	return c.Memory.Read(0x0100 + uint16(c.SP))
}

// pushStackWord pushes a 16-bit word onto the stack by pushing the high byte first, then the low byte
func (c *CPU) pushStackWord(value uint16) {
	hi := uint8(value >> 8)
	lo := uint8(value & 0xFF)
	c.pushStack(hi)
	c.pushStack(lo)
}

// pullStackWord pulls a 16-bit word from the stack by first pulling the low byte, then the high byte
func (c *CPU) pullStackWord() uint16 {
	lo := uint16(c.pullStack())
	hi := uint16(c.pullStack())
	return (hi << 8) | lo
}
