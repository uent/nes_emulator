// Package nes implements the NES system integration
package nes

import (
	"github.com/example/my-golang-project/pkg/ppu"
	"github.com/hajimehoshi/ebiten/v2"
)

// Game implements ebiten.Game for the NES emulator
type Game struct {
	nes      *NES
	renderer *ppu.Renderer
	paused   bool
}

// NewGame creates a new Game instance
func NewGame(nes *NES) *Game {
	return &Game{
		nes:      nes,
		renderer: ppu.NewRenderer(nes.PPU),
		paused:   false,
	}
}

// Update updates the game state
func (g *Game) Update() error {
	// Handle input
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		g.paused = !g.paused
	}
	
	// Update the emulator if not paused
	if !g.paused {
		// Run multiple steps per frame for better performance
		for i := 0; i < 1000; i++ {
			if err := g.nes.Step(); err != nil {
				return err
			}
			
			// Break if a frame is complete
			if g.nes.PPU.FrameComplete {
				break
			}
		}
	}
	
	// Update the renderer
	return g.renderer.Update()
}

// Draw draws the game screen
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the NES output
	g.renderer.Draw(screen)
}

// Layout takes the outside size (e.g., the window size) and returns the logical screen size
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return the NES screen size (256x240)
	return 256, 240
}

// StartGame initializes and starts the NES game
func StartGame(nes *NES) error {
	game := NewGame(nes)
	
	// Configure window
	ebiten.SetWindowSize(512, 480) // 256x240 scaled by 2
	ebiten.SetWindowTitle("NES Emulator")
	
	// Run the game
	return ebiten.RunGame(game)
}
