package grid

import "sync"

// Grid holds the set of live cells and provides concurrent-safe operations
// to evolve and query a Game of Life universe.
type Grid struct {
	cells map[Cell]bool

	mutex sync.RWMutex
}

// SetCell marks the cell at (x, y) as alive.
func (g *Grid) SetCell(x, y int) {
	c := NewCellFromCords(x, y)
	g.mutex.Lock()
	defer g.mutex.Unlock()
	if _, exists := g.cells[*c]; !exists {
		g.cells[*c] = true
	}
}

// ClearCell marks the cell at (x, y) as dead.
func (g *Grid) ClearCell(x, y int) {
	c := NewCellFromCords(x, y)
	g.mutex.Lock()
	defer g.mutex.Unlock()
	delete(g.cells, *c)
}

// IsAlive reports whether the cell at (x, y) is alive.
func (g *Grid) IsAlive(x, y int) bool {
	c := NewCellFromCords(x, y)
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	return g.cells[*c]
}

// Population returns the number of live cells in the grid.
func (g *Grid) Population() int {
	g.mutex.RLock()
	defer g.mutex.RUnlock()
	return len(g.cells)
}

func (g *Grid) Clear() {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	clear(g.cells)
}
