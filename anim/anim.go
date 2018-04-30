package anim

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"

	"github.com/bit101/bitlibgo"
)

// RenderCallback signature for the function that will be called each frame.
type RenderCallback func(*bitlibgo.BitSurface, float64)

// MakeGif makes a gif.
func MakeGif(
	surface *bitlibgo.BitSurface,
	frameCount int,
	fps int,
	tempDir string,
	outputFile string,
	renderFunc RenderCallback,
) {
	framesSpec := fmt.Sprintf("%s/*.png", tempDir)
	makeFrames(surface, frameCount, tempDir, renderFunc)
	convertGif(fps, framesSpec, outputFile)
}

// requires imagemagick to be installed.
func convertGif(fps int, input string, output string) {
	delay := 100.0 / float64(fps)
	cmd := exec.Command("convert", "-delay", strconv.FormatFloat(delay, 'E', -1, 32), input, output)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("saved!")

}

func makeFrames(surface *bitlibgo.BitSurface, numFrames int, framesPath string, renderFunc RenderCallback) {
	for i := 0; i < numFrames; i++ {
		renderFunc(surface, float64(i)/float64(numFrames))
		filename := fmt.Sprintf("%s/anim_%0.3d.png", framesPath, i)
		surface.WriteToPNG(filename)
	}
}
