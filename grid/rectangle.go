package grid

import "math"

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

// Width returns the width of the rectangle.
func (r *Rectangle) Width() int {
	return int(math.Abs(float64(r.MaxX-r.MinX))) + 1
}

// Height returns the height of the rectangle.
func (r *Rectangle) Height() int {
	return int(math.Abs(float64(r.MaxY-r.MinY))) + 1
}
