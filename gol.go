package gol

import "github.com/henilmalaviya/gol/grid"

// Game is a thin wrapper around a Grid to represent a Life simulation instance.
type Game struct {
	grid *grid.Grid
}

// SetGrid sets the simulation's grid.
func (g *Game) SetGrid(grid *grid.Grid) {
	g.grid = grid
}

// GetGrid returns the simulation's grid.
func (g *Game) GetGrid() *grid.Grid {
	return g.grid
}

// NewGame creates a new game with an empty grid.
func NewGame() *Game {
	return &Game{
		grid: grid.NewGrid(),
	}
}
