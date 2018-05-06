package anim

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"

	"github.com/bit101/blg"
)

// RenderCallback signature for the function that will be called each frame.
type RenderCallback func(*blg.BitSurface, float64)

// Animation is a animated gif maker.
type Animation struct {
	Width    float64
	Height   float64
	Frames   int
	FPS      int
	Filename string
}

// NewAnimation creates a new animated gif project.
func NewAnimation(filename string) *Animation {
	return &Animation{
		Width:    200,
		Height:   200,
		Frames:   60,
		FPS:      30,
		Filename: filename,
	}
}

// Render renders the gif
func (p *Animation) Render(renderCallback RenderCallback) {

	framesDir, err := ioutil.TempDir(".", "frames_")
	if err != nil {
		log.Fatal(err)
	}

	surface := blg.NewBitSurface(p.Width, p.Height)
	makeFrames(surface, p.Frames, framesDir, renderCallback)

	framesSpec := fmt.Sprintf("%s/*.png", framesDir)
	convertGif(p.FPS, framesSpec, p.Filename)

	os.RemoveAll(framesDir)
}

// SetSize sets the size
func (p *Animation) SetSize(w float64, h float64) {
	p.Width = w
	p.Height = h
}

func makeFrames(surface *blg.BitSurface, numFrames int, framesDir string, renderFunc RenderCallback) {
	for i := 0; i < numFrames; i++ {
		renderFunc(surface, float64(i)/float64(numFrames))
		filename := fmt.Sprintf("%s/anim_%0.3d.png", framesDir, i)
		surface.WriteToPNG(filename)
	}
}

// requires imagemagick to be installed.
func convertGif(fps int, input string, output string) {
	delay := 100.0 / float64(fps)
	cmd := exec.Command("convert", "-delay", strconv.FormatFloat(delay, 'E', -1, 32), input, output)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Gif saved as %s\n", output)

}
