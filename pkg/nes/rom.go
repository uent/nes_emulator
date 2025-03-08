package nes

import (
	"encoding/binary"
	"fmt"
	"os"
)

// NESHeader represents the header of an NES ROM file
type NESHeader struct {
	Magic      [4]byte // Should be "NES^Z" or [0x4E, 0x45, 0x53, 0x1A]
	PRGROMSize byte    // Size of PRG ROM in 16 KB units
	CHRROMSize byte    // Size of CHR ROM in 8 KB units
	Flags6     byte    // Mapper, mirroring, battery, trainer
	Flags7     byte    // Mapper, VS/Playchoice, NES 2.0
	Flags8     byte    // PRG-RAM size (rarely used)
	Flags9     byte    // TV system (rarely used)
	Flags10    byte    // TV system, PRG-RAM presence (rarely used)
	Reserved   [5]byte // Reserved, should be zero
}

// PRGROM represents the Program ROM data of an NES ROM file
type PRGROM struct {
	Size int64  // Size in bytes
	Data []byte // The actual PRG ROM data
}

// String returns a string representation of the PRG ROM data
func (p *PRGROM) String() string {
	return fmt.Sprintf(
		"PRG ROM Size: %d bytes\nFirst 16 bytes: %X",
		p.Size,
		p.Data[:min(16, len(p.Data))], // Show first 16 bytes or less if data is smaller
	)
}

// min returns the smaller of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ReadNESFile reads a .nes file and returns the extracted header and PRG ROM data
// It loads the PRG ROM data (16 x PRGROMSize KB)
func ReadNESFile(filePath string) (*NESHeader, *PRGROM, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	header := &NESHeader{}
	err = binary.Read(file, binary.LittleEndian, header)
	if err != nil {
		return nil, nil, err
	}

	// Verify magic number
	if string(header.Magic[:]) != "NES\x1A" {
		return nil, nil, fmt.Errorf("not a valid NES ROM file")
	}

	// Load PRG ROM data (16 x PRGROMSize KB)
	prgROMSize := int64(header.PRGROMSize) * 16 * 1024 // Convert to bytes

	// Check if there's a trainer (512 bytes) that we need to skip
	if header.Flags6&0x04 != 0 {
		_, err = file.Seek(512, 1) // Skip trainer (current position + 512 bytes)
		if err != nil {
			return nil, nil, fmt.Errorf("error skipping trainer: %v", err)
		}
	}

	// Allocate a buffer for PRG ROM data
	prgROMData := make([]byte, prgROMSize)

	// Read PRG ROM data
	bytesRead, err := file.Read(prgROMData)
	if err != nil {
		return nil, nil, fmt.Errorf("error reading PRG ROM data: %v", err)
	}

	if int64(bytesRead) != prgROMSize {
		return nil, nil, fmt.Errorf("unexpected PRG ROM size: got %d bytes, expected %d", bytesRead, prgROMSize)
	}

	// Create and populate the PRGROM struct
	prgROM := &PRGROM{
		Size: prgROMSize,
		Data: prgROMData,
	}

	return header, prgROM, nil
}

// String returns a string representation of the NES header
func (h *NESHeader) String() string {
	return fmt.Sprintf(
		"Magic: %s\nPRG ROM: %d x 16KB\nCHR ROM: %d x 8KB\nMapper: %d\nMirroring: %s\nBattery: %t\nTrainer: %t",
		string(h.Magic[:3]),
		h.PRGROMSize,
		h.CHRROMSize,
		(h.Flags7&0xF0)|(h.Flags6>>4),
		mirroringType(h.Flags6),
		h.Flags6&0x02 != 0,
		h.Flags6&0x04 != 0,
	)
}

// mirroringType returns the mirroring type as a string
func mirroringType(flags6 byte) string {
	if flags6&0x08 != 0 {
		return "Four-screen"
	}
	if flags6&0x01 != 0 {
		return "Vertical"
	}
	return "Horizontal"
}