package debug

import (
	"fmt"
	"log"

	"github.com/example/my-golang-project/pkg/nes"
	"github.com/hajimehoshi/ebiten/v2"
)

// DebugGame implements ebiten.Game for the NES debugging interface
type DebugGame struct {
	cpuDebugger *CPUDebugger
	nes         *nes.NES
}

// NewDebugGame creates a new debugging game instance
func NewDebugGame(nes *nes.NES) *DebugGame {
	return &DebugGame{
		cpuDebugger: NewCPUDebugger(nes),
		nes:         nes,
	}
}

// Update updates the game state
func (g *DebugGame) Update() error {
	return g.cpuDebugger.Update()
}

// Draw draws the game screen
func (g *DebugGame) Draw(screen *ebiten.Image) {
	g.cpuDebugger.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the logical screen size
func (g *DebugGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.cpuDebugger.Layout(outsideWidth, outsideHeight)
}

// StartDebugger initializes and starts the NES debugger UI
func StartDebugger(nes *nes.NES) error {
	if nes == nil {
		return fmt.Errorf("NES instance is nil")
	}
	
	game := NewDebugGame(nes)
	
	// Prepare window configuration
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("NES CPU Debugger")
	
	// Run the game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
		return err
	}
	
	return nil
}