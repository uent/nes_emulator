// Package ppu implements the NES Picture Processing Unit emulation
package ppu

import (
	"image/color"
)

// PPU represents the Picture Processing Unit of the NES
type PPU struct {
	// PPU registers
	PPUCTRL   uint8 // $2000
	PPUMASK   uint8 // $2001
	PPUSTATUS uint8 // $2002
	OAMADDR   uint8 // $2003
	OAMDATA   uint8 // $2004
	PPUSCROLL uint8 // $2005
	PPUADDR   uint8 // $2006
	PPUDATA   uint8 // $2007
	
	// PPU memory
	VRAM      []uint8
	OAM       []uint8
	Palette   []uint8
	
	// Internal registers
	v         uint16 // Current VRAM address (15 bits)
	t         uint16 // Temporary VRAM address (15 bits)
	x         uint8  // Fine X scroll (3 bits)
	w         uint8  // First or second write toggle
	
	// Frame information
	Scanline  int
	Cycle     int
	FrameComplete bool
	
	// NMI handling
	nmiOccurred bool
	nmiOutput   bool
	nmiPrevious bool
	nmiDelay    uint8
	
	// Data buffer for PPUDATA reads
	readBuffer uint8
	
	// Rendering buffers
	frontBuffer []uint8 // RGBA buffer (256x240x4)
	backBuffer  []uint8 // RGBA buffer (256x240x4)
	
	// Reference to CPU for NMI triggering
	CPU interface {
		TriggerNMI()
	}
}

// NewPPU creates a new PPU instance
func NewPPU() *PPU {
	return &PPU{
		VRAM: make([]uint8, 0x4000),
		OAM:  make([]uint8, 256),
		Palette: make([]uint8, 32),
		frontBuffer: make([]uint8, 256*240*4), // RGBA buffer
		backBuffer: make([]uint8, 256*240*4),  // RGBA buffer
	}
}

// SetCPU sets the CPU reference for NMI triggering
func (p *PPU) SetCPU(cpu interface {
	TriggerNMI()
}) {
	p.CPU = cpu
}

// Reset resets the PPU to its initial state
func (p *PPU) Reset() {
	p.PPUCTRL = 0
	p.PPUMASK = 0
	p.PPUSTATUS = 0
	p.OAMADDR = 0
	p.PPUSCROLL = 0
	p.PPUADDR = 0
	p.PPUDATA = 0
	
	p.v = 0
	p.t = 0
	p.x = 0
	p.w = 0
	
	p.Scanline = 0
	p.Cycle = 0
	p.FrameComplete = false
	
	p.nmiOccurred = false
	p.nmiOutput = false
	p.nmiPrevious = false
	p.nmiDelay = 0
	
	p.readBuffer = 0
	
	// Initialize buffers with a recognizable pattern
	for y := 0; y < 240; y++ {
		for x := 0; x < 256; x++ {
			offset := (y*256 + x) * 4
			
			// Create a checkerboard pattern
			if ((x / 16) + (y / 16)) % 2 == 0 {
				// Blue squares
				p.frontBuffer[offset] = 0
				p.frontBuffer[offset+1] = 0
				p.frontBuffer[offset+2] = 255
				p.frontBuffer[offset+3] = 255
				
				p.backBuffer[offset] = 0
				p.backBuffer[offset+1] = 0
				p.backBuffer[offset+2] = 255
				p.backBuffer[offset+3] = 255
			} else {
				// White squares
				p.frontBuffer[offset] = 255
				p.frontBuffer[offset+1] = 255
				p.frontBuffer[offset+2] = 255
				p.frontBuffer[offset+3] = 255
				
				p.backBuffer[offset] = 255
				p.backBuffer[offset+1] = 255
				p.backBuffer[offset+2] = 255
				p.backBuffer[offset+3] = 255
			}
		}
	}
}

// ReadRegister reads from a PPU register
func (p *PPU) ReadRegister(address uint16) uint8 {
	// Map the address to 0-7 range (PPU registers)
	reg := address % 8
	
	switch reg {
	case 0x0: // PPUCTRL ($2000) - Write only
		return 0
		
	case 0x1: // PPUMASK ($2001) - Write only
		return 0
		
	case 0x2: // PPUSTATUS ($2002)
		// Reading PPUSTATUS has side effects
		result := p.PPUSTATUS
		// Clear VBlank flag
		p.PPUSTATUS &= 0x7F
		// Reset address latch
		p.w = 0
		return result
		
	case 0x3: // OAMADDR ($2003) - Write only
		return 0
		
	case 0x4: // OAMDATA ($2004)
		return p.OAM[p.OAMADDR]
		
	case 0x5: // PPUSCROLL ($2005) - Write only
		return 0
		
	case 0x6: // PPUADDR ($2006) - Write only
		return 0
		
	case 0x7: // PPUDATA ($2007)
		// Reading from PPUDATA has special behavior
		value := p.readBuffer
		
		// Get the data at the current VRAM address
		p.readBuffer = p.readPPUData()
		
		// Special case for palette memory
		if p.v >= 0x3F00 && p.v <= 0x3FFF {
			value = p.readBuffer
			p.readBuffer = p.readPPUData()
		}
		
		// Auto-increment address
		p.incrementVRAMAddress()
		
		return value
	}
	
	return 0
}

// WriteRegister writes to a PPU register
func (p *PPU) WriteRegister(address uint16, value uint8) {
	// Map the address to 0-7 range (PPU registers)
	reg := address % 8
	
	switch reg {
	case 0x0: // PPUCTRL ($2000)
		p.PPUCTRL = value
		// t: ...BA.. ........ = d: ......BA
		p.t = (p.t & 0xF3FF) | ((uint16(value) & 0x03) << 10)
		// Update NMI output
		p.nmiOutput = (value & 0x80) != 0
		// If NMI occurs and output is enabled, trigger NMI
		if p.nmiOutput && p.nmiOccurred && p.CPU != nil {
			p.CPU.TriggerNMI()
		}
		
	case 0x1: // PPUMASK ($2001)
		p.PPUMASK = value
		
	case 0x2: // PPUSTATUS ($2002) - Read only
		// Writing to PPUSTATUS has no effect
		
	case 0x3: // OAMADDR ($2003)
		p.OAMADDR = value
		
	case 0x4: // OAMDATA ($2004)
		p.OAM[p.OAMADDR] = value
		p.OAMADDR++
		
	case 0x5: // PPUSCROLL ($2005)
		if p.w == 0 {
			// First write (X scroll)
			p.x = value & 0x07
			p.t = (p.t & 0xFFE0) | (uint16(value) >> 3)
			p.w = 1
		} else {
			// Second write (Y scroll)
			p.t = (p.t & 0x8FFF) | ((uint16(value) & 0x07) << 12)
			p.t = (p.t & 0xFC1F) | ((uint16(value) & 0xF8) << 2)
			p.w = 0
		}
		
	case 0x6: // PPUADDR ($2006)
		if p.w == 0 {
			// First write (high byte)
			p.t = (p.t & 0x80FF) | ((uint16(value) & 0x3F) << 8)
			p.w = 1
		} else {
			// Second write (low byte)
			p.t = (p.t & 0xFF00) | uint16(value)
			p.v = p.t
			p.w = 0
		}
		
	case 0x7: // PPUDATA ($2007)
		p.writePPUData(value)
		// Auto-increment address
		p.incrementVRAMAddress()
	}
}

// readPPUData reads data from the current VRAM address
func (p *PPU) readPPUData() uint8 {
	address := p.v & 0x3FFF
	
	// Handle different memory regions
	if address < 0x2000 {
		// Pattern tables
		return p.VRAM[address]
	} else if address < 0x3F00 {
		// Nametables (with mirroring)
		return p.VRAM[address]
	} else {
		// Palette RAM
		paletteAddress := address & 0x1F
		// Handle palette mirroring
		if paletteAddress >= 16 && paletteAddress%4 == 0 {
			paletteAddress -= 16
		}
		return p.Palette[paletteAddress]
	}
}

// writePPUData writes data to the current VRAM address
func (p *PPU) writePPUData(value uint8) {
	address := p.v & 0x3FFF
	
	// Handle different memory regions
	if address < 0x2000 {
		// Pattern tables (usually ROM, but allow writing for testing)
		p.VRAM[address] = value
	} else if address < 0x3F00 {
		// Nametables (with mirroring)
		p.VRAM[address] = value
	} else {
		// Palette RAM
		paletteAddress := address & 0x1F
		// Handle palette mirroring
		if paletteAddress >= 16 && paletteAddress%4 == 0 {
			paletteAddress -= 16
		}
		p.Palette[paletteAddress] = value
	}
}

// incrementVRAMAddress increments the VRAM address based on PPUCTRL
func (p *PPU) incrementVRAMAddress() {
	// Increment by 32 if bit 2 of PPUCTRL is set, otherwise by 1
	if (p.PPUCTRL & 0x04) != 0 {
		p.v += 32
	} else {
		p.v += 1
	}
	p.v &= 0x3FFF // Keep address within 14-bit range
}

// GetPixel returns the color of a pixel at the specified coordinates
func (p *PPU) GetPixel(x, y int) color.RGBA {
	if x < 0 || x >= 256 || y < 0 || y >= 240 {
		return color.RGBA{0, 0, 0, 255}
	}
	
	offset := (y*256 + x) * 4
	return color.RGBA{
		p.frontBuffer[offset],
		p.frontBuffer[offset+1],
		p.frontBuffer[offset+2],
		p.frontBuffer[offset+3],
	}
}

// calculatePixelColor determines the color for the current pixel
func (p *PPU) calculatePixelColor() uint8 {
	// Simplified implementation - just for testing
	// In a real implementation, this would involve checking background and sprite priorities
	
	// If rendering is disabled, return the background color
	if p.PPUMASK&0x08 == 0 && p.PPUMASK&0x10 == 0 {
		return p.Palette[0] & 0x3F
	}
	
	// For now, just return a test pattern
	x := p.Cycle - 1
	y := p.Scanline
	
	// Create a more visible test pattern
	patternX := (x / 16) % 2
	patternY := (y / 16) % 2
	
	if (patternX + patternY) % 2 == 0 {
		return 0x21 // Blue
	} else {
		return 0x30 // White
	}
}

// SwapBuffers swaps the front and back buffers
func (p *PPU) SwapBuffers() {
	p.frontBuffer, p.backBuffer = p.backBuffer, p.frontBuffer
}

// Step advances the PPU by one cycle
func (p *PPU) Step() {
	// Pre-render scanline (-1 or 261)
	if p.Scanline == 261 {
		// Clear VBlank flag at dot 1 of pre-render scanline
		if p.Cycle == 1 {
			p.PPUSTATUS &= 0x7F // Clear bit 7 (VBlank)
			p.nmiOccurred = false
		}
	}
	
	// Visible scanlines (0-239)
	if p.Scanline >= 0 && p.Scanline < 240 {
		// Visible pixel
		if p.Cycle >= 1 && p.Cycle <= 256 {
			x := p.Cycle - 1
			y := p.Scanline
			
			// Calculate the color for this pixel
			colorIndex := p.calculatePixelColor()
			color := GetColor(colorIndex)
			
			// Set the pixel in the back buffer
			offset := (y*256 + x) * 4
			p.backBuffer[offset] = color.R
			p.backBuffer[offset+1] = color.G
			p.backBuffer[offset+2] = color.B
			p.backBuffer[offset+3] = 255 // Full opacity
		}
	}
	
	// VBlank scanlines (241-260)
	if p.Scanline == 241 && p.Cycle == 1 {
		// Set VBlank flag
		p.PPUSTATUS |= 0x80 // Set bit 7
		p.nmiOccurred = true
		
		// Generate NMI if enabled
		if p.nmiOutput && p.CPU != nil {
			p.CPU.TriggerNMI()
		}
		
		// Swap front and back buffers
		p.SwapBuffers()
		p.FrameComplete = true
	}
	
	// Update cycle and scanline counters
	p.Cycle++
	if p.Cycle > 340 {
		p.Cycle = 0
		p.Scanline++
		
		if p.Scanline > 261 {
			p.Scanline = 0
			p.FrameComplete = false
		}
	}
}