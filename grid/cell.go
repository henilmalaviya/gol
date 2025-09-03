package grid

// Cell is a coordinate of a cell in the infinite Game of Life plane.
type Cell struct {
	X int
	Y int
}

func (c Cell) Inside(rect *Rectangle) bool {
	return rect.PointInside(c.X, c.Y)
}

func (c Cell) GetNeighbors() [8]Cell {
	return [8]Cell{
		{X: c.X - 1, Y: c.Y - 1},
		{X: c.X, Y: c.Y - 1},
		{X: c.X + 1, Y: c.Y - 1},
		{X: c.X - 1, Y: c.Y},
		{X: c.X + 1, Y: c.Y},
		{X: c.X - 1, Y: c.Y + 1},
		{X: c.X, Y: c.Y + 1},
		{X: c.X + 1, Y: c.Y + 1},
	}
}

// NewCellFromCords returns a cell with the given coordinates.
func NewCellFromCords(x, y int) *Cell {
	return &Cell{X: x, Y: y}
}

// NewCell returns a zero-valued cell (0,0).
func NewCell() *Cell {
	return &Cell{}
}
