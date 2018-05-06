package blg

import (
	"github.com/bit101/blg/color"
	cairo "github.com/ungerik/go-cairo"
)

// Surface represents a drawing surface with added methods.
type Surface struct {
	Width  float64
	Height float64
	cairo.Surface
}

// NewSurface creates a new Surface.
func NewSurface(width float64, height float64) *Surface {
	return &Surface{
		width,
		height,
		*cairo.NewSurface(cairo.FORMAT_ARGB32, int(width), int(height)),
	}
}

// NewSurfaceFromPNG creates a new Surface.
func NewSurfaceFromPNG(filename string) (*Surface, cairo.Status) {
	surface, status := cairo.NewSurfaceFromPNG(filename)
	if status != cairo.STATUS_SUCCESS {
		return nil, status
	}
	return &Surface{
		float64(surface.GetWidth()),
		float64(surface.GetHeight()),
		*surface,
	}, status
}

// ClearRGB clears the surface to the given rgb color.
func (s *Surface) ClearRGB(r float64, g float64, b float64) {
	s.Save()
	// todo: set identity transform
	s.SetSourceRGB(r, g, b)
	s.Paint()
	s.Restore()
}

// ClearColor clears surface to given color.
func (s *Surface) ClearColor(color color.Color) {
	s.ClearRGB(color.R, color.G, color.B)
}

// SetSourceColor sets the source to the given color.
func (s *Surface) SetSourceColor(color color.Color) {
	s.SetSourceRGBA(color.R, color.G, color.B, color.A)
}

// GetPixel returns the r, g, b, a value at a given x, y location.
func (s *Surface) GetPixel(x int, y int) (byte, byte, byte, byte) {
	data := s.GetData()
	index := (y*s.GetWidth() + x) * 4
	return data[index+2], data[index+1], data[index], data[index+3]
}
