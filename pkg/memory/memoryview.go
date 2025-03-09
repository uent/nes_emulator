// Package memory implements the NES memory system
package memory

// MemoryView provides a restricted view into the main memory
type MemoryView struct {
	baseMemory *Memory
	startAddr  uint16
	endAddr    uint16
}

// NewMemoryView creates a new view into memory with restricted access range
func NewMemoryView(baseMemory *Memory, startAddr uint16, endAddr uint16) *MemoryView {
	return &MemoryView{
		baseMemory: baseMemory,
		startAddr:  startAddr,
		endAddr:    endAddr,
	}
}

// Read returns a byte from the specified memory address if within allowed range
func (mv *MemoryView) Read(address uint16) byte {
	if address >= mv.startAddr && address <= mv.endAddr {
		return mv.baseMemory.Read(address)
	}
	// Out of bounds access, return 0
	return 0
}

// Write writes a byte to the specified memory address if within allowed range
func (mv *MemoryView) Write(address uint16, value byte) {
	if address >= mv.startAddr && address <= mv.endAddr {
		mv.baseMemory.Write(address, value)
	}
	// Silently ignore out of bounds writes
}

// ReadWord reads two consecutive bytes and returns them as a 16-bit word
func (mv *MemoryView) ReadWord(address uint16) uint16 {
	if address >= mv.startAddr && address+1 <= mv.endAddr {
		return mv.baseMemory.ReadWord(address)
	}
	// Out of bounds access, return 0
	return 0
}

// WriteWord writes a 16-bit word as two consecutive bytes
func (mv *MemoryView) WriteWord(address uint16, value uint16) {
	if address >= mv.startAddr && address+1 <= mv.endAddr {
		mv.baseMemory.WriteWord(address, value)
	}
	// Silently ignore out of bounds writes
}