package patterns

import (
	"testing"

	"github.com/henilmalaviya/gol/grid"
)

func TestBlinker(t *testing.T) {
	g := Blinker()
	expected := grid.NewGridFromCells(
		grid.Cell{X: 1, Y: 0},
		grid.Cell{X: 1, Y: 1},
		grid.Cell{X: 1, Y: 2},
	)
	if !g.Compare(expected) {
		t.Errorf("Blinker() = %v, want %v", g, expected)
	}
}

func TestBlinkerHorizontal(t *testing.T) {
	g := BlinkerHorizontal()
	expected := grid.NewGridFromCells(
		grid.Cell{X: 0, Y: 1},
		grid.Cell{X: 1, Y: 1},
		grid.Cell{X: 2, Y: 1},
	)
	if !g.Compare(expected) {
		t.Errorf("BlinkerHorizontal() = %v, want %v", g, expected)
	}
}
