package grid

import "sort"

// Clone returns a deep copy of the grid.
func (g *Grid) Clone() *Grid {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	newGrid := NewGrid()
	for cell := range g.cells {
		newGrid.cells[cell] = true
	}
	return newGrid
}

// Diff compares the current grid with another grid and returns the born and died cells.
func (g *Grid) Diff(other *Grid) ([]Cell, []Cell) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	other.mutex.RLock()
	defer other.mutex.RUnlock()

	var bornCells, diedCells []Cell

	// Find cells that were born (exist in other but not in g).
	for cell := range other.cells {
		if !g.cells[cell] {
			bornCells = append(bornCells, cell)
		}
	}

	// Find cells that died (exist in g but not in other).
	for cell := range g.cells {
		if !other.cells[cell] {
			diedCells = append(diedCells, cell)
		}
	}

	return bornCells, diedCells
}

// SortCells sorts a slice of cells in place.
// Sorting is done first by Y coordinate, then by X coordinate.
func SortCells(cells []Cell) {
	sort.Slice(cells, func(i, j int) bool {
		if cells[i].Y != cells[j].Y {
			return cells[i].Y < cells[j].Y
		}
		return cells[i].X < cells[j].X
	})
}

// Compare checks if the grid is identical to another grid.
func (g1 *Grid) Compare(g2 *Grid) bool {
	g1.mutex.RLock()
	defer g1.mutex.RUnlock()
	g2.mutex.RLock()
	defer g2.mutex.RUnlock()

	if len(g1.cells) != len(g2.cells) {
		return false
	}

	for cell := range g1.cells {
		if !g2.cells[cell] {
			return false
		}
	}

	return true
}
