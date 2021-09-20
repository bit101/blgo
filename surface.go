package blgo

import (
	"github.com/bit101/blgo/color"
	"github.com/bit101/blgo/geom"
	cairo "github.com/bit101/go-cairo"
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
		*cairo.NewSurface(cairo.FormatARGB32, int(width), int(height)),
	}
}

// NewSurfaceFromPNG creates a new Surface.
func NewSurfaceFromPNG(filename string) (*Surface, cairo.Status) {
	surface, status := cairo.NewSurfaceFromPNG(filename)
	if status != cairo.StatusSuccess {
		return nil, status
	}
	return &Surface{
		float64(surface.GetWidth()),
		float64(surface.GetHeight()),
		*surface,
	}, status
}

// NewSVGSurface creates a new surface for creating an SVG image. Finish with surface.Finish()
func NewSVGSurface(filename string, width, height float64) *Surface {
	return &Surface{
		width,
		height,
		*cairo.NewSVGSurface(filename, width, height, cairo.SVGVersion12),
	}
}

// ClearRGB clears the surface to the given rgb color.
func (s *Surface) ClearRGB(r float64, g float64, b float64) {
	s.Save()
	// todo: set identity transform
	s.SetSourceRGB(r, g, b)
	s.Paint()
	s.Restore()
}

// ClearRGBA clears the surface to the given rgba color.
func (s *Surface) ClearRGBA(r, g, b, a float64) {
	s.Save()
	// todo: set identity transform
	s.SetSourceRGBA(r, g, b, a)
	s.Paint()
	s.Restore()
}

// ClearColor clears surface to given color.
func (s *Surface) ClearColor(color color.Color) {
	s.ClearRGB(color.R, color.G, color.B)
}

// ClearWhite clears surface to white.
func (s *Surface) ClearWhite() {
	s.ClearRGB(1, 1, 1)
}

// ClearBlack clears surface to white.
func (s *Surface) ClearBlack() {
	s.ClearRGB(0, 0, 0)
}

// ClearGrey clears surface to white.
func (s *Surface) ClearGrey(g float64) {
	s.ClearRGB(g, g, g)
}

// SetSourceColor sets the source to the given color.
func (s *Surface) SetSourceColor(color color.Color) {
	s.SetSourceRGBA(color.R, color.G, color.B, color.A)
}

// SetSourceHSV sets the source to a color created with the given hue, saturation and value.
func (s *Surface) SetSourceHSV(hue, sat, val float64) {
	s.SetSourceColor(color.HSV(hue, sat, val))
}

func (s *Surface) SetSourceBlack() {
	s.SetSourceRGB(0, 0, 0)
}

func (s *Surface) SetSourceWhite() {
	s.SetSourceRGB(1, 1, 1)
}

func (s *Surface) SetSourceGray(gray float64) {
	s.SetSourceRGB(gray, gray, gray)
}

// GetPixel returns the b, g, r, a value at a given x, y location.
func (s *Surface) GetPixel(x int, y int) (byte, byte, byte, byte) {
	data := s.GetData()
	index := (y*s.GetWidth() + x) * 4
	return data[index+2], data[index+1], data[index], data[index+3]
}

func (s *Surface) Center() {
	s.Translate(s.Width/2, s.Height/2)
}

// GetBounds returns a rectangle defined by the size of the surface.
func (s *Surface) GetBounds() *geom.Rectangle {
	return geom.NewRectangle(0, 0, s.Width, s.Height)
}

// GetAspectRatio returns the aspect ratio of the surface (width / height).
func (s *Surface) GetAspectRatio() float64 {
	return s.Width / s.Height
}

// FillText draws text
func (s *Surface) FillText(text string, x, y float64) {
	s.Save()
	s.Translate(x, y)
	s.ShowText(text)
	s.Fill()
	s.Restore()
}

// TraverseFunc defines a callback function for traversing a surface.
type TraverseFunc func(x, y float64)

// Traverse calls a callback function for every x, y "pixel" on a surface
func (s *Surface) Traverse(callback TraverseFunc) {
	for x := 0.0; x < s.Width; x++ {
		for y := 0.0; y < s.Height; y++ {
			callback(x, y)
		}
	}
}
