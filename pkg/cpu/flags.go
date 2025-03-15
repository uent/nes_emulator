// Package cpu implements the NES CPU (6502) flag operations

package cpu

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

func (c *CPU) setFlagZ(value uint8) {
	c.setFlag(FlagZ, value == 0)
}

func (c *CPU) setFlagN(value uint8) {
	c.setFlag(FlagN, value&0x80 != 0)
}

func (c *CPU) setFlagV(value uint8) {
	c.setFlag(FlagV, value&0x40 != 0)
}

func (c *CPU) setFlagI(value bool, delayed bool) {
	// TODO: Implement delayed funtionality
	c.setFlag(FlagI, value)
}

func (c *CPU) setFlagB(value bool) {
	// TODO: Implement delayed funtionality
	c.setFlag(FlagB, value)
}

// UpdateZN updates the Zero and Negative flags based on the given value
//func (c *CPU) UpdateZN(value uint8) {
//	c.setFlag(FlagZ, value == 0)
//	c.setFlag(FlagN, value&0x80 != 0)
//}
