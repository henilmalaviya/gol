package grid

// NewGrid creates an empty grid.
func NewGrid() *Grid {
	return &Grid{
		cells: make(map[Cell]bool),
	}
}

// NewGridFromCells constructs a grid with the provided live cells.
// Usage: NewGridFromCells(Cell{X:1,Y:2}, Cell{X:3,Y:4})
func NewGridFromCells(cells ...Cell) *Grid {
	g := NewGrid()
	for _, c := range cells {
		g.cells[c] = true
	}
	return g
}

// NewGridFromMap constructs a grid from a map[Cell]bool (true means alive).
func NewGridFromMap(alive map[Cell]bool) *Grid {
	g := NewGrid()
	for c, v := range alive {
		if v {
			g.cells[c] = true
		}
	}
	return g
}

// NewGridFromXY constructs a grid from coordinate pairs.
// Example: NewGridFromXY([][2]int{{1,2},{3,4}})
func NewGridFromXY(coords [][2]int) *Grid {
	g := NewGrid()
	for _, p := range coords {
		g.cells[Cell{X: p[0], Y: p[1]}] = true
	}
	return g
}

// NewGridFromBitmap creates a grid from a 2D boolean matrix positioned at rect.MinX, rect.MinY.
// bitmap[y][x] corresponds to the cell at (rect.MinX + x, rect.MinY + y).
func NewGridFromBitmap(rect Rectangle, bitmap [][]bool) *Grid {
	g := NewGrid()
	for y := 0; y < len(bitmap); y++ {
		row := bitmap[y]
		for x := 0; x < len(row); x++ {
			if row[x] {
				g.cells[Cell{X: rect.MinX + x, Y: rect.MinY + y}] = true
			}
		}
	}
	return g
}

// NewGridFromStrings creates a grid from ASCII art lines at origin (0,0).
// Any of the characters in aliveChars will be treated as a live cell. If none are provided,
// defaults to one of: '*', '#', 'X', 'O', '1'.
func NewGridFromStrings(lines []string, aliveChars ...rune) *Grid {
	g := NewGrid()
	alive := map[rune]struct{}{}
	if len(aliveChars) == 0 {
		for _, r := range []rune{'*', '#', 'X', 'O', '1'} {
			alive[r] = struct{}{}
		}
	} else {
		for _, r := range aliveChars {
			alive[r] = struct{}{}
		}
	}

	for y, line := range lines {
		for x, r := range line {
			if _, ok := alive[r]; ok {
				g.cells[Cell{X: x, Y: y}] = true
			}
		}
	}
	return g
}
