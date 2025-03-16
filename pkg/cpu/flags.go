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

// GetFlag returns the value of a specific status flag
func (c *CPU) GetFlag(flag uint8) uint8 {
	if c.P&(1<<flag) != 0 {
		return 1
	}
	return 0
}

// GetFlagC returns the Carry flag value
func (c *CPU) GetFlagC() uint8 {
	return c.GetFlag(FlagC)
}

// GetFlagZ returns the Zero flag value
func (c *CPU) GetFlagZ() uint8 {
	return c.GetFlag(FlagZ)
}

// GetFlagI returns the Interrupt Disable flag value
func (c *CPU) GetFlagI() uint8 {
	return c.GetFlag(FlagI)
}

// GetFlagD returns the Decimal Mode flag value
func (c *CPU) GetFlagD() uint8 {
	return c.GetFlag(FlagD)
}

// GetFlagB returns the Break Command flag value
func (c *CPU) GetFlagB() uint8 {
	return c.GetFlag(FlagB)
}

// GetFlag5 returns the unused flag value (always 1)
func (c *CPU) GetFlag5() uint8 {
	return c.GetFlag(Flag5)
}

// GetFlagV returns the Overflow flag value
func (c *CPU) GetFlagV() uint8 {
	return c.GetFlag(FlagV)
}

// GetFlagN returns the Negative flag value
func (c *CPU) GetFlagN() uint8 {
	return c.GetFlag(FlagN)
}

// setFlag sets or clears a specific status flag
func (c *CPU) setFlag(flag uint8, value bool) {
	if value {
		c.P |= 1 << flag
	} else {
		c.P &= ^(1 << flag)
	}
}

func (c *CPU) setFlagZ(value bool) {
	c.setFlag(FlagZ, value)
}

func (c *CPU) setFlagZByValue(value uint8) {
	c.setFlagZ(value == 0)
}

func (c *CPU) setFlagC(value bool) {
	c.setFlag(FlagC, value)
}

func (c *CPU) setFlagD(value bool) {
	c.setFlag(FlagD, value)
}

func (c *CPU) setFlagN(value bool) {
	c.setFlag(FlagN, value)
}

func (c *CPU) setFlagNByValue(value uint8) {
	c.setFlagN(value&0x80 != 0)
}

func (c *CPU) setFlagV(value bool) {
	c.setFlag(FlagV, value)
}

func (c *CPU) setFlagVByValue(value uint8) {
	c.setFlagV(value&0x40 != 0)
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
