package grid

import (
	"testing"
)

func TestGlobalObserver(t *testing.T) {
	grid := NewGrid()
	var event ObserverEvent
	observer := NewGlobalObserver(func(e ObserverEvent) {
		event = e
	})
	grid.AddObserver(observer)

	// Test SetCell event
	grid.SetCell(1, 1)
	if event == nil || event.Type() != SetCellEventType {
		t.Errorf("Expected SetCell event, got %v", event)
	}
	if event.(SetCellObserverEvent).Data() != (Cell{1, 1}) {
		t.Errorf("Expected cell {1, 1}, got %v", event.(SetCellObserverEvent).Data())
	}

	// Test ClearCell event
	grid.ClearCell(1, 1)
	if event == nil || event.Type() != ClearCellEventType {
		t.Errorf("Expected ClearCell event, got %v", event)
	}
	if event.(ClearCellObserverEvent).Data() != (Cell{1, 1}) {
		t.Errorf("Expected cell {1, 1}, got %v", event.(ClearCellObserverEvent).Data())
	}

	// Test Tick event
	grid.SetCell(0, 1)
	grid.SetCell(1, 1)
	grid.SetCell(2, 1)
	grid.Tick()
	if event == nil || event.Type() != TickEventType {
		t.Errorf("Expected Tick event, got %v", event)
	}
	born, died := event.(TickObserverEvent).Data()
	if len(born) != 2 || len(died) != 2 {
		t.Errorf("Expected 2 born and 2 died cells, got %d born and %d died", len(born), len(died))
	}

	// Test Step event
	grid.Step(1)
	if event == nil || event.Type() != StepEventType {
		t.Errorf("Expected Step event, got %v", event)
	}
	n, _, _ := event.(StepObserverEvent).Data()
	if n != 1 {
		t.Errorf("Expected step 1, got %d", n)
	}

	// Test ClearGrid event
	grid.Clear()
	if event == nil || event.Type() != ClearGridEventType {
		t.Errorf("Expected ClearGrid event, got %v", event)
	}
}

func TestRegionObserver(t *testing.T) {
	grid := NewGrid()
	region := NewRectangle(0, 0, 4, 4)
	var event ObserverEvent
	observer := NewRegionObserver(*region, func(e ObserverEvent) {
		event = e
	})
	grid.AddObserver(observer)

	// Test SetCell event inside region
	event = nil
	grid.SetCell(1, 1)
	if event == nil || event.Type() != SetCellEventType {
		t.Errorf("Expected SetCell event for cell inside region, got %v", event)
	}

	// Test SetCell event outside region
	event = nil
	grid.SetCell(5, 5)
	if event != nil {
		t.Errorf("Expected no event for cell outside region, got %v", event)
	}

	// Test Tick event
	event = nil
	grid.Clear()
	observer.SetRegion(*NewRectangle(-1, -1, 1, 1))
	grid.SetCell(0, 1)  // inside
	grid.SetCell(0, 0)  // inside
	grid.SetCell(0, -1) // inside

	grid.SetCell(6, 1)  // outside
	grid.SetCell(6, 0)  // outside
	grid.SetCell(6, -1) // outside

	grid.Tick()
	if event == nil || event.Type() != TickEventType {
		t.Errorf("Expected Tick event, got %v", event)
	}
	born, died := event.(TickObserverEvent).Data()
	if len(born) != 2 || len(died) != 2 {
		t.Errorf("Expected 2 born and 2 died cell in region, got %d born and %d died", len(born), len(died))
	}

	// Test ClearGrid event
	event = nil
	grid.Clear()
	if event == nil || event.Type() != ClearGridEventType {
		t.Errorf("Expected ClearGrid event, got %v", event)
	}
}

func TestAddRemoveObserver(t *testing.T) {
	grid := NewGrid()
	var event received = 0
	observer := NewGlobalObserver(func(e ObserverEvent) {
		event++
	})

	grid.AddObserver(observer)
	grid.SetCell(0, 0)
	if event != 1 {
		t.Errorf("Expected 1 event, got %d", event)
	}

	grid.RemoveObserver(observer)
	grid.SetCell(1, 1)
	if event != 1 {
		t.Errorf("Expected 1 event after remove, got %d", event)
	}
}

type received int
