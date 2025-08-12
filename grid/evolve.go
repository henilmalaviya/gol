package grid

// getNeighborCells returns the 8 neighboring cells around (x, y).
func (g *Grid) getNeighborCells(x, y int) [8]Cell {
	return [8]Cell{
		{X: x - 1, Y: y - 1},
		{X: x, Y: y - 1},
		{X: x + 1, Y: y - 1},
		{X: x - 1, Y: y},
		{X: x + 1, Y: y},
		{X: x - 1, Y: y + 1},
		{X: x, Y: y + 1},
		{X: x + 1, Y: y + 1},
	}
}

// ComputeNextGrid computes the next generation and returns the cells that
// will be born and that will die, without mutating the current grid.
func (g *Grid) ComputeNextGrid() ([]Cell, []Cell) {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	var bornCells, diedCells []Cell
	candidates := make(map[Cell]struct{})

	for cell := range g.cells {
		neighbors := g.getNeighborCells(cell.X, cell.Y)

		candidates[cell] = struct{}{}
		for _, n := range neighbors {
			candidates[n] = struct{}{}
		}
	}

	for cell := range candidates {
		isCellAlive := g.cells[cell]
		aliveNeighborCount := 0

		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}

				if g.cells[Cell{X: cell.X + dx, Y: cell.Y + dy}] {
					aliveNeighborCount++
				}
			}
		}

		if !isCellAlive && aliveNeighborCount == 3 {
			bornCells = append(bornCells, cell)
		}

		if isCellAlive && !(aliveNeighborCount == 2 || aliveNeighborCount == 3) {
			diedCells = append(diedCells, cell)
		}

	}

	return bornCells, diedCells
}

// ComputeNextGridN computes the cells that would be born and die after n
// generations, without mutating the current grid.
func (g *Grid) ComputeNextGridN(n int) ([]Cell, []Cell) {
	// Create a temporary grid to simulate the changes without affecting the original grid.
	tempGrid := g.Clone()

	// Simulate n ticks on the temporary grid.
	for i := 0; i < n; i++ {
		tempGrid.Tick()
	}

	return g.Diff(tempGrid)
}

// Tick advances the grid by one generation and returns the born and died cells.
func (g *Grid) Tick() ([]Cell, []Cell) {
	bornCells, diedCells := g.ComputeNextGrid()

	g.mutex.Lock()
	defer g.mutex.Unlock()

	for _, c := range bornCells {
		g.cells[c] = true
	}

	for _, c := range diedCells {
		delete(g.cells, c)
	}

	g.notifyObservers(newTickObserverEvent(bornCells, diedCells))

	return bornCells, diedCells
}

// Step advances the grid by n generations and returns the aggregate born and died cells.
func (g *Grid) Step(n int) ([]Cell, []Cell) {
	bornCells, diedCells := g.ComputeNextGridN(n)

	g.mutex.Lock()
	defer g.mutex.Unlock()

	for _, c := range bornCells {
		g.cells[c] = true
	}

	for _, c := range diedCells {
		delete(g.cells, c)
	}

	g.notifyObservers(newStepObserverEvent(n, bornCells, diedCells))

	return bornCells, diedCells
}
