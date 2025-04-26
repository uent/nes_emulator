// Package nes implements the NES system integration
package nes

import (
	"github.com/example/my-golang-project/pkg/cpu"
	"github.com/example/my-golang-project/pkg/memory"
	"github.com/example/my-golang-project/pkg/ppu"
)

// NES represents the Nintendo Entertainment System
type NES struct {
	CPU    *cpu.CPU
	PPU    *ppu.PPU
	Memory *memory.Memory

	// System state
	Running bool
	Cycles  uint64
}

// New creates a new NES instance
func New() *NES {
	nes := &NES{
		CPU:     cpu.NewCPU(),
		PPU:     ppu.NewPPU(),
		Memory:  memory.New(),
		Running: false,
		Cycles:  0,
	}

	// Connect components
	nes.CPU.SetMemory(nes.Memory)
	nes.PPU.SetCPU(nes.CPU)
	nes.Memory.SetPPU(nes.PPU)

	return nes
}

// Reset resets the NES to its initial state
func (n *NES) Reset() {
	n.Memory.Reset()
	n.PPU.Reset()
	n.CPU.Reset()
	n.Cycles = 0
}

// LoadROM loads a ROM file into memory
func (n *NES) LoadROM(romData []byte) error {
	// Load PRG-ROM data directly
	n.Memory.LoadPRGROM(romData)

	return nil
}

// Step advances the NES emulation by one CPU instruction
func (n *NES) Step() error {
	// Execute one CPU instruction
	cpuCycles, err := n.CPU.Step()
	if err != nil {
		return err
	}

	// For each CPU cycle, the PPU runs 3 cycles
	for i := uint8(0); i < cpuCycles*3; i++ {
		n.PPU.Step()
	}

	// Update total cycles
	n.Cycles += uint64(cpuCycles)

	return nil
}

// Run runs the NES emulation until stopped
func (n *NES) Run() error {
	n.Running = true

	for n.Running {
		err := n.Step()
		if err != nil {
			return err
		}
	}

	return nil
}

// Stop stops the NES emulation
func (n *NES) Stop() {
	n.Running = false
}
