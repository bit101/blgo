package floodfill

// Surface is an interface for a surface used for performing flood fills
// Used to avoid import cycles
type Surface interface {
	GetHeight() int
	GetWidth() int
	SetData([]byte)
	GetData() []byte
}

// FloodFill fills a surface from the specified point with the specified color
func FloodFill(surface Surface, x, y, r, g, b float64, threshold float64) {
	// surface uses float values, but pixel data is all int values. convert.
	xi, yi := int(x), int(y)
	w, h := surface.GetWidth(), surface.GetHeight()

	// if pixel outside of image, return
	if xi < 0 || xi >= w || yi < 0 || yi >= h {
		return
	}

	// more int conversions
	ri, gi, bi, th := int(r*255), int(g*255), int(b*255), int(threshold*255)

	data := surface.GetData()
	nextCoord := newCoord(xi, yi, w)

	// color of target color and replacement color for fill
	targetRGB := getRGB(data, nextCoord)
	replacementRGB := newRGB(ri, gi, bi)

	// target color is already replacement color, return
	if targetRGB.isEqual(replacementRGB, th) {
		return
	}

	// change color of first pixel to replacement color, add coord to queue
	setRGB(data, nextCoord, replacementRGB)
	var queue []coord
	queue = append(queue, nextCoord)

	// helper func. if coord is in bounds and has target color, change to replacement color and add to queue
	checkNeighbor := func(x, y int) {
		neighbor := newCoord(x, y, w)
		if neighbor.x < w && neighbor.x >= 0 && neighbor.y < h && neighbor.y >= 0 &&
			getRGB(data, neighbor).isEqual(targetRGB, th) {

			setRGB(data, neighbor, replacementRGB)
			queue = append(queue, neighbor)
		}
	}

	// while there are items in the queue
	for len(queue) > 0 {
		// shift the first one off (it's already been colored
		nextCoord, queue = queue[0], queue[1:]

		// check four surrounding coords
		checkNeighbor(nextCoord.x+1, nextCoord.y)
		checkNeighbor(nextCoord.x-1, nextCoord.y)
		checkNeighbor(nextCoord.x, nextCoord.y+1)
		checkNeighbor(nextCoord.x, nextCoord.y-1)
	}

	// queue is empty. rewrite data
	surface.SetData(data)
}
