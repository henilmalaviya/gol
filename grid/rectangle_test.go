package grid

import (
	"testing"
)

func TestRectangle_Width(t *testing.T) {
	tests := []struct {
		name string
		rect Rectangle
		want int
	}{
		{
			name: "standard rectangle",
			rect: *NewRectangle(1, 1, 5, 5),
			want: 5,
		},
		{
			name: "rectangle with negative coordinates",
			rect: *NewRectangle(-5, -5, -1, -1),
			want: 5,
		},
		{
			name: "zero-width rectangle",
			rect: *NewRectangle(1, 1, 1, 5),
			want: 1,
		},
		{
			name: "inverted rectangle",
			rect: *NewRectangle(5, 1, 1, 5),
			want: 5,
		},
		{
			name: "point",
			rect: *NewRectangle(0, 0, 0, 0),
			want: 1,
		},
		{
			name: "center block",
			rect: *NewRectangle(-1, -1, 1, 1),
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rect.Width(); got != tt.want {
				t.Errorf("Rectangle.Width() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRectangle_Height(t *testing.T) {
	tests := []struct {
		name string
		rect Rectangle
		want int
	}{
		{
			name: "standard rectangle",
			rect: *NewRectangle(1, 1, 5, 5),
			want: 5,
		},
		{
			name: "rectangle with negative coordinates",
			rect: *NewRectangle(-5, -5, -1, -1),
			want: 5,
		},
		{
			name: "zero-height rectangle",
			rect: *NewRectangle(1, 1, 5, 1),
			want: 1,
		},
		{
			name: "inverted rectangle",
			rect: *NewRectangle(1, 5, 5, 1),
			want: 5,
		},
		{
			name: "point",
			rect: *NewRectangle(0, 0, 0, 0),
			want: 1,
		},
		{
			name: "center block",
			rect: *NewRectangle(-1, -1, 1, 1),
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rect.Height(); got != tt.want {
				t.Errorf("Rectangle.Height() = %v, want %v", got, tt.want)
			}
		})
	}
}
