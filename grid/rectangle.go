package grid

// Rectangle is an inclusive axis-aligned bounding box on the integer grid.
type Rectangle struct {
	MinX int
	MinY int
	MaxX int
	MaxY int
}

// NewRectangle constructs a rectangle (inclusive bounds).
func NewRectangle(minX, minY, maxX, maxY int) *Rectangle {
	return &Rectangle{minX, minY, maxX, maxY}
}
