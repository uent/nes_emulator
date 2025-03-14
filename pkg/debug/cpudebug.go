// Package debug provides debugging tools for the NES emulator
package debug

import (
	"fmt"
	"image/color"

	//"strconv"

	"github.com/example/my-golang-project/pkg/cpu"
	"github.com/example/my-golang-project/pkg/nes"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 800
	screenHeight = 600
	padding      = 20
	lineHeight   = 20
)

// CPUDebugger represents a debugger for CPU state visualization
type CPUDebugger struct {
	nes         *nes.NES
	paused      bool
	stepMode    bool
	nextStep    bool
	disassembly []string
	cycleCount  int
}

// NewCPUDebugger creates a new CPU debugger UI
func NewCPUDebugger(nes *nes.NES) *CPUDebugger {
	return &CPUDebugger{
		nes:         nes,
		paused:      true,
		stepMode:    true,
		nextStep:    false,
		disassembly: make([]string, 0),
		cycleCount:  0,
	}
}

// Update updates the debugger state
func (d *CPUDebugger) Update() error {
	// Space key toggles pause
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		d.paused = !d.paused
		d.stepMode = d.paused
	}

	// S key steps forward when in step mode
	if d.stepMode && inpututil.IsKeyJustPressed(ebiten.KeyS) {
		d.nextStep = true
	}

	// Step the NES if running or in step mode with next step requested
	if (!d.paused || (d.stepMode && d.nextStep)) && d.nes != nil {
		d.nes.Step()
		d.cycleCount++
		d.nextStep = false

		// Capture current instruction for disassembly
		cpu := d.nes.CPU
		opcode := d.nes.Memory.Read(cpu.PC)
		instruction := cpu.GetInstruction(opcode)

		// Add to disassembly log (keeping only last 20 instructions)
		disasm := fmt.Sprintf("%04X: %02X %s", cpu.PC, opcode, instruction.Mnemonic)
		d.disassembly = append(d.disassembly, disasm)
		if len(d.disassembly) > 20 {
			d.disassembly = d.disassembly[len(d.disassembly)-20:]
		}
	}

	return nil
}

// Draw renders the debugger UI
func (d *CPUDebugger) Draw(screen *ebiten.Image) {
	// Fill background
	screen.Fill(color.RGBA{40, 40, 40, 255})

	face := basicfont.Face7x13

	if d.nes == nil || d.nes.CPU == nil {
		text.Draw(screen, "NES not initialized", face, padding, padding+lineHeight, color.White)
		return
	}

	cpu := d.nes.CPU

	// Draw status information
	statusText := "PAUSED"
	if !d.paused {
		statusText = "RUNNING"
	}
	if d.stepMode {
		statusText += " (STEP MODE - Press S to step)"
	}
	text.Draw(screen, statusText, face, padding, padding, color.RGBA{200, 200, 100, 255})

	// Draw CPU registers
	y := padding + 2*lineHeight
	text.Draw(screen, fmt.Sprintf("Cycle: %d", d.cycleCount), face, padding, y, color.White)
	y += lineHeight

	text.Draw(screen, fmt.Sprintf("A: $%02X (%d)", cpu.A, cpu.A), face, padding, y, color.White)
	y += lineHeight

	text.Draw(screen, fmt.Sprintf("X: $%02X (%d)", cpu.X, cpu.X), face, padding, y, color.White)
	y += lineHeight

	text.Draw(screen, fmt.Sprintf("Y: $%02X (%d)", cpu.Y, cpu.Y), face, padding, y, color.White)
	y += lineHeight

	text.Draw(screen, fmt.Sprintf("PC: $%04X", cpu.PC), face, padding, y, color.RGBA{100, 255, 100, 255})
	y += lineHeight

	text.Draw(screen, fmt.Sprintf("SP: $%02X", cpu.SP), face, padding, y, color.White)
	y += lineHeight

	// Draw CPU flags
	y += lineHeight
	text.Draw(screen, "Flags:", face, padding, y, color.RGBA{255, 200, 100, 255})
	y += lineHeight

	flagsStr := d.formatFlags(cpu)
	text.Draw(screen, flagsStr, face, padding, y, color.White)
	y += lineHeight

	// Show detailed flags
	y += lineHeight
	text.Draw(screen, fmt.Sprintf("N: %d | V: %d | -: %d | B: %d | D: %d | I: %d | Z: %d | C: %d",
		cpu.GetFlagN(), cpu.GetFlagV(), cpu.GetFlag5(),
		cpu.GetFlagB(), cpu.GetFlagD(), cpu.GetFlagI(),
		cpu.GetFlagZ(), cpu.GetFlagC()),
		face, padding, y, color.White)

	// Current and next instructions
	y += 2 * lineHeight
	text.Draw(screen, "Disassembly:", face, padding, y, color.RGBA{100, 200, 255, 255})

	// Draw disassembly
	y += lineHeight
	for i, line := range d.disassembly {
		var textColor color.Color = color.White
		// Highlight current instruction
		if i == len(d.disassembly)-1 {
			textColor = color.RGBA{255, 255, 0, 255}
		}
		text.Draw(screen, line, face, padding, y, textColor)
		y += lineHeight
	}

	// Draw memory view
	y = padding + 2*lineHeight
	x := screenWidth / 2

	text.Draw(screen, "Memory (Zero Page):", face, x, y, color.RGBA{100, 200, 255, 255})
	y += lineHeight

	// Show  first 64 bytes arround of the PC in 4 columns

	start_position := int(d.nes.CPU.PC / 64)

	cols := 4
	rows := 16

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			addr := uint16(row*cols + col + start_position*64)
			val := d.nes.Memory.Read(addr)
			var textColor color.RGBA = color.RGBA{255, 255, 255, 255}
			pcValue := d.nes.CPU.PC
			if addr == uint16(pcValue) {
				textColor = color.RGBA{100, 200, 255, 255} // Highlight for memory location pointed to by value at PC
			}
			text.Draw(screen, fmt.Sprintf("$%02X:%02X ", addr, val), face, x+(col*80), y, textColor)
		}
		y += lineHeight
	}

	// Draw stack view
	y += lineHeight
	text.Draw(screen, "Stack:", face, x, y, color.RGBA{100, 200, 255, 255})
	y += lineHeight

	stackBase := uint16(0x0100)
	for i := 0; i < 8; i++ {
		addr := stackBase + uint16(cpu.SP) + uint16(i+1)
		if addr <= 0x01FF {
			val := d.nes.Memory.Read(addr)
			text.Draw(screen, fmt.Sprintf("$%04X: %02X", addr, val), face, x, y, color.White)
			y += lineHeight
		}
	}

	// Draw debug controls help
	ebitenutil.DebugPrintAt(screen, "Controls: SPACE to toggle pause, S to step", padding, screenHeight-padding)
}

// formatFlags formats the processor status flags in a readable way
func (d *CPUDebugger) formatFlags(cpu *cpu.CPU) string {
	flagsStr := "Status: "

	// NV-BDIZC
	if cpu.GetFlagN() == 1 {
		flagsStr += "N"
	} else {
		flagsStr += "n"
	}

	if cpu.GetFlagV() == 1 {
		flagsStr += "V"
	} else {
		flagsStr += "v"
	}

	flagsStr += "-" // Unused bit always 1

	if cpu.GetFlagB() == 1 {
		flagsStr += "B"
	} else {
		flagsStr += "b"
	}

	if cpu.GetFlagD() == 1 {
		flagsStr += "D"
	} else {
		flagsStr += "d"
	}

	if cpu.GetFlagI() == 1 {
		flagsStr += "I"
	} else {
		flagsStr += "i"
	}

	if cpu.GetFlagZ() == 1 {
		flagsStr += "Z"
	} else {
		flagsStr += "z"
	}

	if cpu.GetFlagC() == 1 {
		flagsStr += "C"
	} else {
		flagsStr += "c"
	}

	return flagsStr
}

// Layout implements the ebiten.Game interface
func (d *CPUDebugger) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
