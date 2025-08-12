package grid

// Cell is a coordinate of a cell in the infinite Game of Life plane.
type Cell struct {
	X int
	Y int
}

func (c Cell) Inside(rect *Rectangle) bool {
	return c.X >= rect.MinX && c.X <= rect.MaxX &&
		c.Y >= rect.MinY && c.Y <= rect.MaxY
}

// NewCellFromCords returns a cell with the given coordinates.
func NewCellFromCords(x, y int) *Cell {
	return &Cell{X: x, Y: y}
}

// NewCell returns a zero-valued cell (0,0).
func NewCell() *Cell {
	return &Cell{}
}
