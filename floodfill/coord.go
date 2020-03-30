package floodfill

type coord struct {
	index int
	x     int
	y     int
}

func newCoord(x, y, w int) coord {
	// we get width passed so we can precompute the data index
	index := (y*w + x) * 4
	return coord{index: index, x: x, y: y}
}
