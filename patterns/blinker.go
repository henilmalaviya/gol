package patterns

import (
	gr "github.com/henilmalaviya/gol/grid"
)

// Blinker returns a grid initialized with the blinker pattern.
func Blinker() *gr.Grid {
	g := gr.NewGrid()
	// Vertical blinker pattern
	g.SetCell(1, 0)
	g.SetCell(1, 1)
	g.SetCell(1, 2)
	return g
}

// BlinkerHorizontal returns a grid initialized with the horizontal blinker pattern.
func BlinkerHorizontal() *gr.Grid {
	g := gr.NewGrid()
	// Horizontal blinker pattern
	g.SetCell(0, 1)
	g.SetCell(1, 1)
	g.SetCell(2, 1)
	return g
}
