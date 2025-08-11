package gol

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()
	if game.grid == nil {
		t.Errorf("NewGame() should create a grid, but it was nil")
	}
}
