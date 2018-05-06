package anim

import (
	"fmt"

	"github.com/bit101/blg"
)

// RenderCallback signature for the function that will be called each frame.
type RenderCallback func(*blg.Surface, float64)

// Animation is a animated gif maker.
type Animation struct {
	Width      float64
	Height     float64
	FrameCount int
}

// NewAnimation creates a new animated gif project.
func NewAnimation(width, height float64, frameCount int) *Animation {
	return &Animation{
		Width:      width,
		Height:     height,
		FrameCount: frameCount,
	}
}

// Render renders the gif
func (p *Animation) Render(framesDir string, prefix string, renderCallback RenderCallback) {
	surface := blg.NewSurface(p.Width, p.Height)
	for i := 0; i < p.FrameCount; i++ {
		renderCallback(surface, float64(i)/float64(p.FrameCount))
		filename := fmt.Sprintf("%s/%s_%0.4d.png", framesDir, prefix, i)
		surface.WriteToPNG(filename)
	}
}

// SetSize sets the size
func (p *Animation) SetSize(w float64, h float64) {
	p.Width = w
	p.Height = h
}
