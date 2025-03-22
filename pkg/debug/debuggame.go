package debug

import (
	"fmt"
	"log"

	"github.com/example/my-golang-project/pkg/nes"
	"github.com/hajimehoshi/ebiten/v2"
)

// DebugGame implements ebiten.Game for the NES debugging interface
type DebugGame struct {
	cpuDebugger   *CPUDebugger
	nes           *nes.NES
	isFirstRender bool
}

// InitScreen initializes the screen rendering settings
func (g *DebugGame) InitScreen() {
	// Initialize screen-related components
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("NES Debugger")

	g.cpuDebugger.PrintDisassembly()
}

// NewDebugGame creates a new debugging game instance
func NewDebugGame(nes *nes.NES) *DebugGame {
	game := &DebugGame{
		cpuDebugger:   NewCPUDebugger(nes),
		nes:           nes,
		isFirstRender: false,
	}
	game.InitScreen() // Initialize screen settings
	return game
}

// Update updates the game state
func (g *DebugGame) Update() error {
	return g.cpuDebugger.Update()
}

// Draw draws the game screen
func (g *DebugGame) Draw(screen *ebiten.Image) {
	if !g.isFirstRender {
		g.OnScreenRenderInit()
		g.isFirstRender = true
	}
	g.cpuDebugger.Draw(screen)
}

// OnScreenRenderInit runs when the screen is first rendered
func (g *DebugGame) OnScreenRenderInit() {
	// Add any initialization code that needs to run when screen rendering starts
	fmt.Println("Screen rendering initialized")
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
