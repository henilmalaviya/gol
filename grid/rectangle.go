package grid

import "math"

// Rectangle is an inclusive axis-aligned bounding box on the integer grid.
type Rectangle struct {
	X1 int
	Y1 int
	X2 int
	Y2 int
}

// NewRectangle constructs a rectangle (inclusive bounds).
func NewRectangle(x1, y1, x2, y2 int) *Rectangle {
	return &Rectangle{X1: x1, Y1: y1, X2: x2, Y2: y2}
}

// Width returns the width of the rectangle.
func (r *Rectangle) Width() int {
	return int(math.Abs(float64(r.X2-r.X1))) + 1
}

// Height returns the height of the rectangle.
func (r *Rectangle) Height() int {
	return int(math.Abs(float64(r.Y2-r.Y1))) + 1
}

func (r *Rectangle) ToArray() [4]int {
	return [4]int{r.X1, r.Y1, r.X2, r.Y2}
}

func (r *Rectangle) Min() (int, int) {
	return int(math.Min(float64(r.X1), float64(r.X2))), int(math.Min(float64(r.Y1), float64(r.Y2)))
}

func (r *Rectangle) Max() (int, int) {
	return int(math.Max(float64(r.X1), float64(r.X2))), int(math.Max(float64(r.Y1), float64(r.Y2)))
}

// account for inverted cords
func (r *Rectangle) PointInside(x, y int) bool {
	minX, minY := r.Min()
	maxX, maxY := r.Max()
	return x >= minX && x <= maxX && y >= minY && y <= maxY
}

func (r *Rectangle) Normalized() *Rectangle {
	nr := NewRectangle(r.X1, r.Y1, r.X2, r.Y2)
	minX, minY := nr.Min()
	maxX, maxY := nr.Max()
	nr.X1, nr.Y1, nr.X2, nr.Y2 = minX, minY, maxX, maxY
	return nr
}

func (r *Rectangle) ToNestedArray() [2][2]int {
	return [2][2]int{
		{r.X1, r.Y1},
		{r.X2, r.Y2},
	}
}
