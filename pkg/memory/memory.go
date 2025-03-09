// Package memory implements the NES memory system
package memory

import "fmt"

// Memory represents the memory system of the NES
type Memory struct {
	// RAM (2KB internal RAM, mirrored 4 times)
	RAM [0x0800]byte

	// PPU registers (8 bytes, mirrored throughout $2000-$3FFF)
	PPURegisters [8]byte

	// APU and I/O registers
	APUAndIORegisters [0x0020]byte

	// Cartridge space: PRG ROM, PRG RAM, and mapper registers
	CartridgeSpace [0xBFE0]byte
}

// New creates a new Memory instance
func New() *Memory {
	// Create a new memory instance
	m := &Memory{}
	
	// Initialize RAM (addresses 0x0000 to 0x07FF) with zeros
	for i := 0; i < 0x0800; i++ {
		m.RAM[i] = 0
	}
	
	return m
}

// Reset initializes the memory to its power-on state
func (m *Memory) Reset() {
	// Clear RAM
	for i := range m.RAM {
		m.RAM[i] = 0
	}

	// Clear PPU registers
	for i := range m.PPURegisters {
		m.PPURegisters[i] = 0
	}

	// Clear APU and I/O registers
	for i := range m.APUAndIORegisters {
		m.APUAndIORegisters[i] = 0
	}

	// We don't clear CartridgeSpace as it should be loaded from ROM
}

// Read returns a byte from the specified memory address
func (m *Memory) Read(address uint16) byte {
	fmt.Println("Read from address:", address)
	switch {
	case address < 0x2000:
		// Internal RAM, mirrored every 0x0800 bytes
		return m.RAM[address%0x0800]
	case address < 0x4000:
		// PPU registers, mirrored every 8 bytes
		return m.PPURegisters[(address-0x2000)%8]
	case address < 0x4020:
		// APU and I/O registers
		return m.APUAndIORegisters[address-0x4000]
	default:
		// Cartridge space: PRG ROM, PRG RAM, and mapper registers
		return m.CartridgeSpace[address-0x4020]
	}
}

// Write writes a byte to the specified memory address
func (m *Memory) Write(address uint16, value byte) {
	switch {
	case address < 0x2000:
		// Internal RAM, mirrored every 0x0800 bytes
		m.RAM[address%0x0800] = value
	case address < 0x4000:
		// PPU registers, mirrored every 8 bytes
		m.PPURegisters[(address-0x2000)%8] = value
	case address < 0x4020:
		// APU and I/O registers
		m.APUAndIORegisters[address-0x4000] = value
	default:
		// Cartridge space: PRG ROM, PRG RAM, and mapper registers
		m.CartridgeSpace[address-0x4020] = value
	}
}

// LoadPRGROM loads the program ROM into memory
func (m *Memory) LoadPRGROM(prgROM []byte) {
	// Copy PRG ROM data into the appropriate location in cartridge space
	// Typically starting at 0x8000, but this is a simplification
	// In a real implementation, this would depend on the mapper
	for i := 0; i < len(prgROM) && i < len(m.CartridgeSpace); i++ {
		m.CartridgeSpace[i] = prgROM[i]
	}
}

// ReadWord reads a 16-bit word from the specified memory address
// NES is little-endian, so the first byte is the low byte
func (m *Memory) ReadWord(address uint16) uint16 {
	fmt.Println("ReadWord from address:", address)
	low := uint16(m.Read(address))
	high := uint16(m.Read(address + 1))
	return (high << 8) | low
}

// WriteWord writes a 16-bit word to the specified memory address
func (m *Memory) WriteWord(address uint16, value uint16) {
	m.Write(address, byte(value&0xFF))
	m.Write(address+1, byte(value>>8))
}
