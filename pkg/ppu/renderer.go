// Package ppu implements the NES Picture Processing Unit emulation
package ppu

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Renderer handles rendering the PPU output to an Ebiten image
type Renderer struct {
	ppu         *PPU
	frameBuffer *ebiten.Image
	scale       float64
	frameCount  int
}

// NewRenderer creates a new PPU renderer
func NewRenderer(ppu *PPU) *Renderer {
	// Create a new image with the NES resolution
	img := ebiten.NewImage(256, 240)
	
	// Fill with a recognizable color to verify it's working
	img.Fill(color.RGBA{0, 0, 255, 255})
	
	fmt.Println("PPU Renderer initialized with a 256x240 frame buffer")
	
	return &Renderer{
		ppu:         ppu,
		frameBuffer: img,
		scale:       0.8, // Smaller scale factor to fit better
		frameCount:  0,
	}
}

// SetScale sets the rendering scale factor
func (r *Renderer) SetScale(scale float64) {
	r.scale = scale
}

// Update updates the renderer state
func (r *Renderer) Update() error {
	// Check if a new frame is ready
	if r.ppu.FrameComplete {
		r.updateFrameBuffer()
		r.frameCount++
		
		if r.frameCount % 60 == 0 {
			fmt.Printf("Renderer: Updated frame %d\n", r.frameCount)
		}
	}
	return nil
}

// updateFrameBuffer copies the PPU's front buffer to the Ebiten image
func (r *Renderer) updateFrameBuffer() {
	// Create a temporary RGBA image to hold the pixel data
	img := image.NewRGBA(image.Rect(0, 0, 256, 240))
	
	// Copy pixel data from PPU's front buffer to the RGBA image
	for y := 0; y < 240; y++ {
		for x := 0; x < 256; x++ {
			offset := (y*256 + x) * 4
			c := color.RGBA{
				r.ppu.frontBuffer[offset],
				r.ppu.frontBuffer[offset+1],
				r.ppu.frontBuffer[offset+2],
				r.ppu.frontBuffer[offset+3],
			}
			img.SetRGBA(x, y, c)
		}
	}
	
	// Update the Ebiten image
	r.frameBuffer.WritePixels(img.Pix)
}

// Draw draws the frame buffer to the screen
func (r *Renderer) Draw(screen *ebiten.Image) {
	// Set up drawing options
	op := &ebiten.DrawImageOptions{}
	
	// Apply scaling
	op.GeoM.Scale(r.scale, r.scale)
	
	// Draw the frame buffer to the screen
	screen.DrawImage(r.frameBuffer, op)
}

// GetFrameBuffer returns the current frame buffer image
func (r *Renderer) GetFrameBuffer() *ebiten.Image {
	return r.frameBuffer
}
