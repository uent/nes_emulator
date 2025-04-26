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

	UnmappedCartridgeMemorySize   = 0x1FE0
	UnmappedCartridgeStartAddress = TestingMemoryStartAddress + TestingMemorySize

	RAMCartridgeMemorySize   = 0x2000
	RAMCartridgeStartAddress = UnmappedCartridgeStartAddress + UnmappedCartridgeMemorySize

	ROMCartridgeMemorySize   = 0x8000
	ROMCartridgeStartAddress = RAMCartridgeStartAddress + RAMCartridgeMemorySize
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
	UnmappedCartridgeSpace [UnmappedCartridgeMemorySize]byte

	// Cartridge space: PRG ROM, PRG RAM, and mapper registers
	RAMCartridgeSpace [RAMCartridgeMemorySize]byte

	// Cartridge space: PRG ROM, PRG RAM, and mapper registers
	ROMCartridgeSpace [ROMCartridgeMemorySize]byte
	
	// Reference to PPU for register access
	PPU interface {
		ReadRegister(address uint16) uint8
		WriteRegister(address uint16, value uint8)
	}
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

// SetPPU sets the PPU interface for register access
func (m *Memory) SetPPU(ppu interface {
	ReadRegister(address uint16) uint8
	WriteRegister(address uint16, value uint8)
}) {
	m.PPU = ppu
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
		if m.PPU != nil {
			return m.PPU.ReadRegister(address)
		}
		return m.PPURegisters[(address-0x2000)%PPURegistersSize]
		
	case address < TestingMemoryStartAddress: // 0x4000 - 0x4017
		// APU and I/O registers
		return m.APUAndIORegisters[address-0x4000]
		
	case address < UnmappedCartridgeStartAddress: // 0x4018 - 0x401F
		panic("testing memory space")
		
	case address < RAMCartridgeStartAddress: // 0x4020 - 0x6000
		return m.RAMCartridgeSpace[address-0x4020]
		
	case address < ROMCartridgeStartAddress: // 0x6000 - 0x7FFF
		return m.RAMCartridgeSpace[address-0x6000]
		
	case address <= ROMCartridgeStartAddress+ROMCartridgeMemorySize-1: // 0x8000 - 0xFFFF
		return m.ROMCartridgeSpace[address-0x8000]
		
	default:
		fmt.Printf("Invalid memory address: %x \n", address)
		panic("Invalid memory address")
	}
}

// Write writes a byte to the specified memory address
func (m *Memory) Write(address uint16, value byte) {
	switch {
	case address < PPURegistersStartAddress: // 0x0 - 0x1FFF
		// Internal RAM, mirrored every 0x0800 bytes
		m.RAM[address%RAMSize] = value
		
	case address < APUAndIORegistersStartAddress: // 0x2000 - 0x3fff
		// PPU registers, mirrored every 8 bytes
		if m.PPU != nil {
			m.PPU.WriteRegister(address, value)
		} else {
			m.PPURegisters[(address-0x2000)%PPURegistersSize] = value
		}
		
	case address < TestingMemoryStartAddress: // 0x4000 - 0x4017
		// APU and I/O registers
		m.APUAndIORegisters[address-0x4000] = value
		
	case address < UnmappedCartridgeStartAddress: // 0x4018 - 0x401F
		panic("testing memory space")
		
	case address < RAMCartridgeStartAddress: // 0x4020 - 0x6000
		m.RAMCartridgeSpace[address-0x4020] = value
		
	case address < ROMCartridgeStartAddress: // 0x6000 - 0x7FFF
		m.RAMCartridgeSpace[address-0x6000] = value
		
	case address < ROMCartridgeStartAddress+ROMCartridgeMemorySize-1: // 0x8000 - 0xFFFF
		m.ROMCartridgeSpace[address-0x8000] = value
		
	default:
		fmt.Printf("Invalid memory address: %x \n", address)
		panic("Invalid memory address")
	}
}

// LoadPRGROM loads the program ROM into memory
func (m *Memory) LoadPRGROM(prgROM []byte) {
	// Copy PRG ROM data into the appropriate location in cartridge space
	// Typically starting at 0x8000, but this is a simplification
	// In a real implementation, this would depend on the mapper
	for i := 0; i < len(prgROM) && i < len(m.ROMCartridgeSpace); i++ {
		m.ROMCartridgeSpace[i] = prgROM[i]
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

func (m *Memory) ReadAddressIndirectPageBoundaryBug(address uint16) uint16 {
	var addr uint16
	if (address & 0x00FF) == 0x00FF { // Si el puntero termina en XXFF, aplica el bug
		low := m.Read(address)
		high := m.Read(address & 0xFF00) // Lee desde XX00 en vez de XXFF+1
		addr = uint16(high)<<8 | uint16(low)
	} else {
		addr = m.ReadWord(address) // Lectura normal
	}

	return addr
}
