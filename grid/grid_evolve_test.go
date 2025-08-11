package grid

import (
	"reflect"
	"testing"
)

func TestGetNeighborCells(t *testing.T) {
	g := NewGrid()
	neighbors := g.getNeighborCells(0, 0)
	expected := [8]Cell{
		{X: -1, Y: -1}, {X: 0, Y: -1}, {X: 1, Y: -1},
		{X: -1, Y: 0}, {X: 1, Y: 0},
		{X: -1, Y: 1}, {X: 0, Y: 1}, {X: 1, Y: 1},
	}
	if neighbors != expected {
		t.Errorf("getNeighborCells(0,0) failed. Expected %v, got %v", expected, neighbors)
	}
}

func TestComputeNextGrid(t *testing.T) {
	// Still life: block
	block := NewGridFromCells(Cell{X: 0, Y: 0}, Cell{X: 1, Y: 0}, Cell{X: 0, Y: 1}, Cell{X: 1, Y: 1})
	born, died := block.ComputeNextGrid()
	if len(born) != 0 || len(died) != 0 {
		t.Errorf("Still life (block) should not change. Born: %d, Died: %d", len(born), len(died))
	}

	// Oscillator: blinker
	blinker := NewGridFromCells(Cell{X: 1, Y: 0}, Cell{X: 1, Y: 1}, Cell{X: 1, Y: 2})
	born, died = blinker.ComputeNextGrid()

	expectedBorn := []Cell{{X: 0, Y: 1}, {X: 2, Y: 1}}
	expectedDied := []Cell{{X: 1, Y: 0}, {X: 1, Y: 2}}

	// Sort slices to ensure comparison is order-independent
	SortCells(born)
	SortCells(died)
	SortCells(expectedBorn)
	SortCells(expectedDied)

	if !reflect.DeepEqual(born, expectedBorn) {
		t.Errorf("Blinker born cells incorrect. Expected %v, got %v", expectedBorn, born)
	}
	if !reflect.DeepEqual(died, expectedDied) {
		t.Errorf("Blinker died cells incorrect. Expected %v, got %v", expectedDied, died)
	}
}

func TestStep(t *testing.T) {
	// Blinker pattern
	blinker := NewGridFromCells(Cell{X: 1, Y: 0}, Cell{X: 1, Y: 1}, Cell{X: 1, Y: 2})
	g := blinker.Clone()
	g.Step(2) // After 2 steps, it should be back to original state
	if !g.Compare(blinker) {
		t.Errorf("Step(2) for blinker failed. Expected %v, got %v", blinker, g)
	}
}

func TestTick(t *testing.T) {
	// Blinker pattern
	blinker := NewGridFromCells(Cell{X: 1, Y: 0}, Cell{X: 1, Y: 1}, Cell{X: 1, Y: 2})
	g := blinker.Clone()
	g.Tick()
	expected := NewGridFromCells(Cell{X: 0, Y: 1}, Cell{X: 1, Y: 1}, Cell{X: 2, Y: 1})
	if !g.Compare(expected) {
		t.Errorf("Tick() for blinker failed. Expected %v, got %v", expected, g)
	}
	g.Tick() // Second tick should return to original state
	if !g.Compare(blinker) {
		t.Errorf("Second Tick() for blinker failed. Expected %v, got %v", blinker, g)
	}
}
