package grid

import (
	"testing"
)

func TestNewGridFromCells(t *testing.T) {
	cells := []Cell{{X: 0, Y: 0}, {X: 1, Y: 1}}
	g := NewGridFromCells(cells...)
	if g.Population() != 2 {
		t.Errorf("Expected population 2, got %d", g.Population())
	}
	if !g.IsAlive(0, 0) || !g.IsAlive(1, 1) {
		t.Errorf("Grid not initialized correctly from cells")
	}
}

func TestNewGridFromMap(t *testing.T) {
	cellMap := map[Cell]bool{
		{X: 0, Y: 0}: true,
		{X: 1, Y: 1}: true,
		{X: 2, Y: 2}: false,
	}
	g := NewGridFromMap(cellMap)
	if g.Population() != 2 {
		t.Errorf("Expected population 2, got %d", g.Population())
	}
	if !g.IsAlive(0, 0) || !g.IsAlive(1, 1) {
		t.Errorf("Grid not initialized correctly from map")
	}
	if g.IsAlive(2, 2) {
		t.Errorf("Cell with false value should not be alive")
	}
}

func TestNewGridFromXY(t *testing.T) {
	coords := [][2]int{{0, 0}, {1, 1}}
	g := NewGridFromXY(coords)
	if g.Population() != 2 {
		t.Errorf("Expected population 2, got %d", g.Population())
	}
	if !g.IsAlive(0, 0) || !g.IsAlive(1, 1) {
		t.Errorf("Grid not initialized correctly from XY coordinates")
	}
}

func TestNewGridFromBitmap(t *testing.T) {
	bitmap := [][]bool{
		{true, false},
		{false, true},
	}
	rect := Rectangle{X1: 1, Y1: 1}
	g := NewGridFromBitmap(rect, bitmap)
	if g.Population() != 2 {
		t.Errorf("Expected population 2, got %d", g.Population())
	}
	if !g.IsAlive(1, 1) || !g.IsAlive(2, 2) {
		t.Errorf("Grid not initialized correctly from bitmap")
	}
}

func TestNewGridFromStrings(t *testing.T) {
	lines := []string{
		"* ",
		" *",
	}
	g := NewGridFromStrings(lines)
	if g.Population() != 2 {
		t.Errorf("Expected population 2, got %d", g.Population())
	}
	if !g.IsAlive(0, 0) || !g.IsAlive(1, 1) {
		t.Errorf("Grid not initialized correctly from strings")
	}
}
