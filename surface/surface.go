package surface

import (
	"github.com/bit101/bitlibgo/color"
	cairo "github.com/ungerik/go-cairo"
)

// BitSurface represents a drawing surface with added methods.
type BitSurface struct {
	Width  float64
	Height float64
	cairo.Surface
}

// NewBitSurface creates a new BitSurface.
func NewBitSurface(width float64, height float64) *BitSurface {
	return &BitSurface{
		width,
		height,
		*cairo.NewSurface(cairo.FORMAT_ARGB32, int(width), int(height)),
	}
}

// ClearRGB clears the surface to the given rgb color.
func (s *BitSurface) ClearRGB(r float64, g float64, b float64) {
	s.Save()
	// todo: set identity transform
	s.SetSourceRGB(r, g, b)
	s.Paint()
	s.Restore()
}

// ClearColor clears surface to given color.
func (s *BitSurface) ClearColor(color color.Color) {
	s.ClearRGB(color.R, color.G, color.B)
}

// SetSourceColor sets the source to the given color.
func (s *BitSurface) SetSourceColor(color color.Color) {
	s.SetSourceRGBA(color.R, color.G, color.B, color.A)
}
