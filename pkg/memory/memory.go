// Package memory implements the NES memory system
package memory

import "fmt"

const (
	// RAMSize represents the size of the NES's internal RAM in bytes
	RAMSize         = 0x0800
	RAMStartAddress = 0x0000

	// Ram cloned 3 times

	PPURegistersSize         = 0x0008
	PPURegistersStartAddress = RAMStartAddress + 4*RAMSize // RAM is cloned 3 times

	// PPU cloned 1023 times

	APUAndIORegistersSize         = 0x0018
	APUAndIORegistersStartAddress = PPURegistersStartAddress + 1024*PPURegistersSize // PPU cloned 1023 times

	TestingMemorySize         = 0x0008
	TestingMemoryStartAddress = APUAndIORegistersStartAddress + APUAndIORegistersSize

	ROMMemorySize   = 0x1FE0 + 0x2000 + 0x8000 // rom: unmapped + ram + rom
	ROMStartAddress = TestingMemoryStartAddress + TestingMemorySize
)

// Memory represents the memory system of the NES
type Memory struct {
	// RAM (2KB internal RAM, mirrored 4 times)
	RAM [RAMSize]byte // 2024

	// PPU registers (8 bytes, mirrored throughout $2000-$3FFF)
	PPURegisters [PPURegistersSize]byte

	// APU and I/O registers
	APUAndIORegisters [APUAndIORegistersSize]byte

	// Cartridge space: PRG ROM, PRG RAM, and mapper registers
	//CartridgeSpace [0xBFE0]byte
	CartridgeSpace [ROMMemorySize]byte
}

// New creates a new Memory instance
func New() *Memory {
	// Create a new memory instance
	m := &Memory{}

	// Initialize RAM (addresses 0x0000 to 0x07FF) with zeros
	for i := 0; i < 0x0800; i++ {
		m.RAM[i] = 0
	}
	fmt.Println("CartridgeSpace size:", (0xBFE0 + 0x2000 + 0x8000 + 0xFFFF))

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

	switch {
	case address < PPURegistersStartAddress: // 0x0 - 0x1FFF
		// Internal RAM, mirrored every 0x0800 bytes
		return m.RAM[address%RAMSize]
	case address < APUAndIORegistersStartAddress: // 0x2000 - 0x3fff
		// PPU registers, mirrored every 8 bytes
		return m.PPURegisters[address%PPURegistersSize]
	case address < TestingMemoryStartAddress: // 0x4000 - 0x4017
		// APU and I/O registers
		return m.APUAndIORegisters[address-0x4000]
	case address < ROMStartAddress: // 0x4018 - 0x401F
		panic("testing memory space")
	case address <= (ROMStartAddress + ROMMemorySize - 1): // 0x4020 - 0xFFFF
		// Cartridge space: PRG ROM, PRG RAM, and mapper registers
		return m.CartridgeSpace[address-0x4020]
	default:
		fmt.Printf("Invalid memory address: %x \n", address)
		panic("Invalid memory address")
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
	low := uint16(m.Read(address))
	high := uint16(m.Read(address + 1))
	return (high << 8) | low
}

// WriteWord writes a 16-bit word to the specified memory address
func (m *Memory) WriteWord(address uint16, value uint16) {
	m.Write(address, byte(value&0xFF))
	m.Write(address+1, byte(value>>8))
}
