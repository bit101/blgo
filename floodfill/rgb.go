package floodfill

type rgb struct {
	r int
	g int
	b int
}

func newRGB(r, g, b int) rgb {
	return rgb{
		r: r,
		g: g,
		b: b,
	}
}

func (p rgb) isEqual(q rgb, max int) bool {
	return abs(p.r, q.r) <= max && abs(p.g, q.g) <= max && abs(p.b, q.b) <= max
}

func abs(a, b int) int {
	diff := a - b
	if diff < 0 {
		return -diff
	}
	return diff
}

func getRGB(data []byte, c coord) rgb {
	// stored as b, g, r, a
	return newRGB(int(data[c.index+2]), int(data[c.index+1]), int(data[c.index]))
}

func setRGB(data []byte, c coord, replacementRGB rgb) {
	// stored as b, g, r, a
	data[c.index+2] = byte(replacementRGB.r)
	data[c.index+1] = byte(replacementRGB.g)
	data[c.index] = byte(replacementRGB.b)
}
