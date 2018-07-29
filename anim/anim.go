package anim

import (
	"fmt"

	"github.com/bit101/blgo"
)

// RenderCallback signature for the function that will be called each frame.
type RenderCallback func(float64)

// Animation is a animated gif maker.
type Animation struct {
	Surface    *blgo.Surface
	FrameCount int
	FramesDir  string
}

// NewAnimation creates a new animated gif project.
func NewAnimation(surface *blgo.Surface, frameCount int, framesDir string) *Animation {
	return &Animation{
		Surface:    surface,
		FrameCount: frameCount,
		FramesDir:  framesDir,
	}
}

// Render renders the gif
func (p *Animation) Render(renderCallback RenderCallback) {
	for i := 0; i < p.FrameCount; i++ {
		percent := float64(i) / float64(p.FrameCount)
		fmt.Printf("\r%f", percent)
		renderCallback(percent)
		filename := fmt.Sprintf("%s/frame_%0.4d.png", p.FramesDir, i)
		p.Surface.WriteToPNG(filename)
	}
	fmt.Println()
}
