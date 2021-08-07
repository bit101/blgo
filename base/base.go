package base

import (
	"github.com/bit101/blgo"
	"github.com/bit101/blgo/anim"
	"github.com/bit101/blgo/util"
)

type RenderFunc func(surface *blgo.Surface, width, height, percent float64)
type SetupFunc func(surface *blgo.Surface, width, height float64)

type Sketch struct {
	width     float64
	height    float64
	surface   *blgo.Surface
	setup     SetupFunc
	render    RenderFunc
	framesDir string
}

func NewSketch(setup SetupFunc, render RenderFunc) *Sketch {
	return &Sketch{
		setup:     setup,
		render:    render,
		framesDir: "frames", // Must Exist!!!
	}
}

func (s *Sketch) SetSize(width, height float64) {
	s.width = width
	s.height = height
	s.surface = blgo.NewSurface(width, height)
}

func (s *Sketch) RenderImage(width, height, t float64) {
	outFileName := "out.png"
	s.SetSize(width, height)
	s.setup(s.surface, s.width, s.height)
	s.render(s.surface, s.width, s.height, t)
	s.surface.WriteToPNG(outFileName)
	util.ViewImage(outFileName)
}

func (s *Sketch) RenderGif(width, height float64, seconds, fps int) {
	outFileName := "out.gif"
	s.SetSize(width, height)
	s.renderAnim(seconds, fps)
	util.ConvertToGIF(s.framesDir, outFileName, float64(fps))
	util.ViewImage(outFileName)
}

func (s *Sketch) RenderVideo(width, height float64, seconds, fps int) {
	outFileName := "out.mp4"
	s.SetSize(width, height)
	s.renderAnim(seconds, fps)
	util.ConvertToYoutube(s.framesDir, outFileName, fps)
	util.VLC(outFileName)
}

func (s *Sketch) renderAnim(seconds, fps int) {
	s.setup(s.surface, s.width, s.height)
	animation := anim.NewAnimation(s.surface, fps*seconds, s.framesDir)
	animation.Render(s.internalRender)
}

func (s *Sketch) internalRender(percent float64) {
	s.render(s.surface, s.width, s.height, percent)
}
