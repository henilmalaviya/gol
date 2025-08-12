package grid

// Translate shifts all live cells by (dx, dy).
// Positive dx moves right; positive dy moves down.
func (g *Grid) Translate(dx, dy int) {
	if dx == 0 && dy == 0 {
		return
	}
	g.mutex.Lock()
	defer g.mutex.Unlock()

	newCells := make(map[Cell]bool, len(g.cells))
	for cell := range g.cells {
		newCells[Cell{X: cell.X + dx, Y: cell.Y + dy}] = true
	}
	g.cells = newCells
}

// Bitmap returns a dense boolean matrix covering rect where true marks a live cell.
// bitmap[y][x] corresponds to the cell at (rect.MinX + x, rect.MinY + y).
func (g *Grid) Bitmap(rect Rectangle) [][]bool {
	width := rect.Width()
	height := rect.Height()

	if width <= 0 || height <= 0 {
		return [][]bool{}
	}

	bitmap := make([][]bool, height)
	for y := range height {
		bitmap[y] = make([]bool, width)
	}

	// Use Subgrid to filter cells within rect, avoiding duplicate bounds checks
	sub := g.Subgrid(rect)
	// No need to lock sub since it's isolated and not shared, but keep symmetry
	sub.mutex.RLock()
	for c := range sub.cells {
		bitmap[c.Y-rect.MinY][c.X-rect.MinX] = true
	}
	sub.mutex.RUnlock()

	return bitmap
}

// Subgrid returns a new grid containing only the live cells within rect.
// Coordinates are preserved (absolute positions).
func (g *Grid) Subgrid(rect Rectangle) *Grid {
	ng := NewGrid()
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	for c := range g.cells {
		if c.Inside(&rect) && g.cells[c] {
			ng.cells[c] = true
		}
	}
	return ng
}

// SubgridNormalized returns a new grid containing only the live cells within rect,
// translated so that (rect.MinX, rect.MinY) becomes (0, 0).
func (g *Grid) SubgridNormalized(rect Rectangle) *Grid {
	ng := g.Subgrid(rect)
	ng.Translate(-rect.MinX, -rect.MinY)
	return ng
}

// Bounds returns the minimal inclusive rectangle that contains all live cells.
// If the grid is empty, it returns (0,0,0,0).
func (g *Grid) Bounds() Rectangle {

	g.mutex.RLock()
	defer g.mutex.RUnlock()

	if len(g.cells) == 0 {
		return *NewRectangle(0, 0, 0, 0)
	}

	var minX, minY, maxX, maxY int
	first := true
	for c := range g.cells {
		if first {
			minX, maxX = c.X, c.X
			minY, maxY = c.Y, c.Y
			first = false
		} else {
			if c.X < minX {
				minX = c.X
			}
			if c.X > maxX {
				maxX = c.X
			}
			if c.Y < minY {
				minY = c.Y
			}
			if c.Y > maxY {
				maxY = c.Y
			}
		}
	}

	return *NewRectangle(minX, minY, maxX, maxY)
}
