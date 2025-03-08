// Package ppu implements the NES Picture Processing Unit emulation
package ppu

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
	
	// Internal registers
	v         uint16 // Current VRAM address (15 bits)
	t         uint16 // Temporary VRAM address (15 bits)
	x         uint8  // Fine X scroll (3 bits)
	w         uint8  // First or second write toggle
	
	// Frame information
	Scanline  int
	Cycle     int
	FrameComplete bool
}

// NewPPU creates a new PPU instance
func NewPPU() *PPU {
	return &PPU{
		VRAM: make([]uint8, 0x4000),
		OAM:  make([]uint8, 256),
	}
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
}

// Step advances the PPU by one cycle
func (p *PPU) Step() {
	// TODO: Implement PPU cycle emulation
	
	// Update cycle and scanline counters
	p.Cycle++
	if p.Cycle > 340 {
		p.Cycle = 0
		p.Scanline++
		
		if p.Scanline > 261 {
			p.Scanline = 0
			p.FrameComplete = true
		}
	}
}