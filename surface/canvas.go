package surface

import (
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

// todo when i have a color struct
//     func ClearColor(color Color) {
//         s.ClearRgb(color.r, color.g, color.b)
//     }

//     func SetSourceColor(color Color) {
//         s.SetSourceRgba(color.r, color.g, color.b, color.a)
//     }
