package grid

import (
	"testing"
)

func TestNewGrid(t *testing.T) {
	g := NewGrid()
	if g.cells == nil {
		t.Errorf("NewGrid() should create a grid with a non-nil map, but it was nil")
	}
	if len(g.cells) != 0 {
		t.Errorf("NewGrid() should create an empty grid, but it has %d cells", len(g.cells))
	}
}

func TestSetCell(t *testing.T) {
	g := NewGrid()
	g.SetCell(1, 2)
	if !g.IsAlive(1, 2) {
		t.Errorf("SetCell(1, 2) should make the cell alive, but it's not")
	}
	if g.Population() != 1 {
		t.Errorf("Population should be 1 after setting one cell, but it's %d", g.Population())
	}
}

func TestClearCell(t *testing.T) {
	g := NewGridFromCells(Cell{X: 1, Y: 2})
	g.ClearCell(1, 2)
	if g.IsAlive(1, 2) {
		t.Errorf("ClearCell(1, 2) should make the cell dead, but it's alive")
	}
	if g.Population() != 0 {
		t.Errorf("Population should be 0 after clearing the only cell, but it's %d", g.Population())
	}
}

func TestIsAlive(t *testing.T) {
	g := NewGridFromCells(Cell{X: 1, Y: 2})
	if !g.IsAlive(1, 2) {
		t.Errorf("IsAlive(1, 2) should be true for a live cell, but it's false")
	}
	if g.IsAlive(0, 0) {
		t.Errorf("IsAlive(0, 0) should be false for a dead cell, but it's true")
	}
}

func TestPopulation(t *testing.T) {
	g := NewGridFromCells(Cell{X: 0, Y: 0}, Cell{X: 1, Y: 1})
	if g.Population() != 2 {
		t.Errorf("Population should be 2, but it's %d", g.Population())
	}
}

func TestClear(t *testing.T) {
	g := NewGridFromCells(Cell{X: 0, Y: 0}, Cell{X: 1, Y: 1})
	g.Clear()
	if g.Population() != 0 {
		t.Errorf("Population should be 0 after clearing the grid, but it's %d", g.Population())
	}
	if g.IsAlive(0, 0) {
		t.Errorf("Cell (0,0) should be dead after clearing the grid, but it's alive")
	}
}
