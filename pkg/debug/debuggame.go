package debug

import (
	"fmt"
	"image/color"
	"log"

	"github.com/example/my-golang-project/pkg/nes"
	"github.com/example/my-golang-project/pkg/ppu"
	"github.com/hajimehoshi/ebiten/v2"
)

// DebugGame implements ebiten.Game for the NES debugging interface
type DebugGame struct {
	cpuDebugger   *CPUDebugger
	nes           *nes.NES
	renderer      *ppu.Renderer
	isFirstRender bool
	frameCount    int
}

// InitScreen initializes the screen rendering settings
func (g *DebugGame) InitScreen() {
	// Initialize screen-related components
	ebiten.SetWindowSize(1300, 600)
	ebiten.SetWindowTitle("NES Debugger with Graphics")

	g.cpuDebugger.PrintDisassembly()
}

// NewDebugGame creates a new debugging game instance
func NewDebugGame(nes *nes.NES) *DebugGame {
	game := &DebugGame{
		cpuDebugger:   NewCPUDebugger(nes),
		nes:           nes,
		renderer:      ppu.NewRenderer(nes.PPU),
		isFirstRender: false,
		frameCount:    0,
	}
	game.InitScreen() // Initialize screen settings
	return game
}

// Update updates the game state
func (g *DebugGame) Update() error {
	// Update CPU debugger
	err := g.cpuDebugger.Update()
	if err != nil {
		return err
	}
	
	// Run multiple steps per frame for better performance when not paused
	if !g.cpuDebugger.paused {
		for i := 0; i < 100; i++ {
			if err := g.nes.Step(); err != nil {
				return err
			}
			
			// Break if a frame is complete
			if g.nes.PPU.FrameComplete {
				break
			}
		}
	}
	
	// Update renderer
	g.frameCount++
	if g.frameCount % 10 == 0 {
		fmt.Printf("Frame %d, PPU cycle: %d, scanline: %d\n", 
			g.frameCount, g.nes.PPU.Cycle, g.nes.PPU.Scanline)
	}
	
	return g.renderer.Update()
}

// Draw draws the game screen
func (g *DebugGame) Draw(screen *ebiten.Image) {
	if !g.isFirstRender {
		g.OnScreenRenderInit()
		g.isFirstRender = true
	}
	
	// Fill background
	screen.Fill(color.RGBA{40, 40, 40, 255})
	
	// Draw a vertical separator line
	for y := 0; y < 600; y++ {
		screen.Set(650, y, color.RGBA{100, 100, 100, 255})
	}
	
	// Draw CPU debugger on the left side
	g.cpuDebugger.Draw(screen)
	
	// Draw PPU output on the right side
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(800, 50) // Position the PPU output much further to the right
	screen.DrawImage(g.renderer.GetFrameBuffer(), op)
}

// OnScreenRenderInit runs when the screen is first rendered
func (g *DebugGame) OnScreenRenderInit() {
	// Add any initialization code that needs to run when screen rendering starts
	fmt.Println("Screen rendering initialized with both CPU debugger and PPU output")
}

// Layout takes the outside size (e.g., the window size) and returns the logical screen size
func (g *DebugGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	// Return a larger logical screen size to accommodate both debugger and game output
	return 1300, 600
}

// StartDebugger initializes and starts the NES debugger UI
func StartDebugger(nes *nes.NES) error {
	if nes == nil {
		return fmt.Errorf("NES instance is nil")
	}

	game := NewDebugGame(nes)

	// Prepare window configuration
	ebiten.SetWindowSize(1300, 600)
	ebiten.SetWindowTitle("NES CPU Debugger with Graphics")

	// Run the game
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
