package grid

import (
	"reflect"
	"testing"
)

func TestGridBitmap(t *testing.T) {

	grid := NewGrid()
	grid.SetCell(0, 1)
	grid.SetCell(0, 0)
	grid.SetCell(0, -1)

	rect := NewRectangle(-1, 1, 1, -1)

	gotBitmap := grid.Bitmap(*rect)

	expectedBitMap := [][]bool{{false, true, false}, {false, true, false}, {false, true, false}}

	if !reflect.DeepEqual(gotBitmap, expectedBitMap) {
		t.Errorf("Bitmap() = %v, want %v", gotBitmap, expectedBitMap)
	}

	rect2 := NewRectangle(-1, -1, 1, 1)
	gotBitmap = grid.Bitmap(*rect2)

	if !reflect.DeepEqual(gotBitmap, expectedBitMap) {
		t.Errorf("Bitmap() = %v, want %v", gotBitmap, expectedBitMap)
	}

}
