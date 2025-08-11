package patterns

import (
	"testing"

	"github.com/henilmalaviya/gol/grid"
)

func TestGlider(t *testing.T) {
	g := Glider()
	expected := grid.NewGridFromCells(
		grid.Cell{X: 1, Y: 0},
		grid.Cell{X: 2, Y: 1},
		grid.Cell{X: 0, Y: 2},
		grid.Cell{X: 1, Y: 2},
		grid.Cell{X: 2, Y: 2},
	)
	if !g.Compare(expected) {
		t.Errorf("Glider() = %v, want %v", g, expected)
	}
}

func TestGliderGun(t *testing.T) {
	g := GliderGun()
	expected := grid.NewGridFromCells(
		grid.Cell{X: 0, Y: 24},
		grid.Cell{X: 1, Y: 24},
		grid.Cell{X: 2, Y: 24},
		grid.Cell{X: 0, Y: 25},
		grid.Cell{X: 1, Y: 25},
		grid.Cell{X: 2, Y: 25},
		grid.Cell{X: 10, Y: 22},
		grid.Cell{X: 10, Y: 23},
		grid.Cell{X: 10, Y: 24},
		grid.Cell{X: 11, Y: 21},
		grid.Cell{X: 12, Y: 20},
		grid.Cell{X: 13, Y: 20},
		grid.Cell{X: 14, Y: 21},
		grid.Cell{X: 15, Y: 22},
	)
	if !g.Compare(expected) {
		t.Errorf("GliderGun() = %v, want %v", g, expected)
	}
}
