// Package nes implements the main NES system
package nes

import (
	"github.com/example/my-golang-project/pkg/apu"
	"github.com/example/my-golang-project/pkg/cpu"
	"github.com/example/my-golang-project/pkg/memory"
	"github.com/example/my-golang-project/pkg/ppu"
)

// NES represents the Nintendo Entertainment System
type NES struct {
	CPU     *cpu.CPU
	PPU     *ppu.PPU
	APU     *apu.APU
	Memory  *memory.Memory
	running bool
}

// New creates a new NES instance with all components initialized
func New() *NES {
	mem := memory.New()

	cpuInstance := cpu.NewCPU()
	// Use the main memory system directly instead of a restricted view
	cpuInstance.SetMemory(mem)

	nes := &NES{
		CPU:     cpuInstance,
		PPU:     ppu.NewPPU(),
		APU:     apu.NewAPU(),
		Memory:  mem,
		running: false,
	}

	return nes
}

// Reset resets all components of the NES
func (n *NES) Reset() {
	n.Memory.Reset()
	n.CPU.Reset()
	n.PPU.Reset()
	n.APU.Reset()
	n.running = false
}

// LoadROM loads a ROM into the NES
func (n *NES) LoadROM(prgROM []byte) {
	n.Memory.LoadPRGROM(prgROM)
}

// Step advances the NES emulation by one step
func (n *NES) Step() {
	n.CPU.Step()

	// PPU runs at 3x the speed of CPU
	n.PPU.Step()
	n.PPU.Step()
	n.PPU.Step()

	// APU step
	n.APU.Step()
}

// Run starts the execution of the NES after reset and ROM loading
// It will continuously execute instructions until Stop is called
func (n *NES) Run() {
	n.running = true
	for n.running {
		n.Step()
	}
}

// RunFor executes the NES for a specified number of cycles
// Useful for testing or when precise control is needed
func (n *NES) RunFor(cycles int) {
	for i := 0; i < cycles; i++ {
		n.Step()
	}
}

// Stop halts the execution of the NES
func (n *NES) Stop() {
	n.running = false
}
