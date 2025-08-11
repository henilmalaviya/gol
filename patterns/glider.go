package patterns

import gr "github.com/henilmalaviya/gol/grid"

// Glider returns a grid initialized with the standard glider pattern at origin.
func Glider() *gr.Grid {
	g := gr.NewGrid()
	g.SetCell(1, 0)
	g.SetCell(2, 1)
	g.SetCell(0, 2)
	g.SetCell(1, 2)
	g.SetCell(2, 2)
	return g
}

// GliderGun returns a grid with a partial Gosper glider gun seed.
func GliderGun() *gr.Grid {
	g := gr.NewGrid()
	g.SetCell(0, 24)
	g.SetCell(1, 24)
	g.SetCell(2, 24)
	g.SetCell(0, 25)
	g.SetCell(1, 25)
	g.SetCell(2, 25)
	g.SetCell(10, 22)
	g.SetCell(10, 23)
	g.SetCell(10, 24)
	g.SetCell(11, 21)
	g.SetCell(12, 20)
	g.SetCell(13, 20)
	g.SetCell(14, 21)
	g.SetCell(15, 22)
	return g
}
