package grid

type ObserverEventType string

const (
	SetCellEventType   ObserverEventType = "set_cell"
	ClearCellEventType ObserverEventType = "clear_cell"
	TickEventType      ObserverEventType = "tick"
	ClearGridEventType ObserverEventType = "clear_grid"
	StepEventType      ObserverEventType = "step"
)

type ObserverEvent interface {
	Type() ObserverEventType
}

type Observer interface {
	Update(event ObserverEvent)
}

// ---

func (g *Grid) AddObserver(observer Observer) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if g.observers == nil {
		g.observers = make(map[Observer]struct{})
	}
	g.observers[observer] = struct{}{}
}

func (g *Grid) RemoveObserver(observer Observer) {
	g.mutex.Lock()
	defer g.mutex.Unlock()

	if g.observers == nil {
		return
	}

	for obs := range g.observers {
		if obs == observer {
			delete(g.observers, obs)
			break
		}
	}
}

func (g *Grid) notifyObservers(event ObserverEvent) {
	if g.observers == nil {
		return
	}

	for observer := range g.observers {
		observer.Update(event)
	}
}

// ---

type SetCellObserverEvent struct {
	cell Cell
}

func (e SetCellObserverEvent) Type() ObserverEventType {
	return SetCellEventType
}

func (e SetCellObserverEvent) Data() Cell {
	return e.cell
}

func newSetCellObserverEvent(cell Cell) SetCellObserverEvent {
	return SetCellObserverEvent{cell: cell}
}

// ---

type ClearCellObserverEvent struct {
	cell Cell
}

func (e ClearCellObserverEvent) Type() ObserverEventType {
	return ClearCellEventType
}

func (e ClearCellObserverEvent) Data() Cell {
	return e.cell
}

func newClearCellObserverEvent(cell Cell) ClearCellObserverEvent {
	return ClearCellObserverEvent{cell: cell}
}

// ---

type TickObserverEvent struct {
	bornCells []Cell
	diedCells []Cell
}

func (e TickObserverEvent) Type() ObserverEventType {
	return TickEventType
}

func (e TickObserverEvent) Data() ([]Cell, []Cell) {
	return e.bornCells, e.diedCells
}

func newTickObserverEvent(bornCells, diedCells []Cell) TickObserverEvent {
	return TickObserverEvent{
		bornCells: bornCells,
		diedCells: diedCells,
	}
}

// ---

type ClearGridObserverEvent struct{}

func (e ClearGridObserverEvent) Type() ObserverEventType {
	return ClearGridEventType
}

func newClearGridObserverEvent() ClearGridObserverEvent {
	return ClearGridObserverEvent{}
}

// ---

type StepObserverEvent struct {
	n         int
	bornCells []Cell
	diedCells []Cell
}

func (e StepObserverEvent) Type() ObserverEventType {
	return StepEventType
}

func (e StepObserverEvent) Data() (int, []Cell, []Cell) {
	return e.n, e.bornCells, e.diedCells
}

func newStepObserverEvent(n int, bornCells, diedCells []Cell) StepObserverEvent {
	return StepObserverEvent{
		n:         n,
		bornCells: bornCells,
		diedCells: diedCells,
	}
}

// ---

type GlobalObserver struct {
	updateFunc func(event ObserverEvent)
}

func (o *GlobalObserver) Update(event ObserverEvent) {
	o.updateFunc(event)
}

func NewGlobalObserver(updateFunc func(event ObserverEvent)) *GlobalObserver {
	return &GlobalObserver{
		updateFunc: updateFunc,
	}
}

// ---

type RegionObserver struct {
	region     Rectangle
	updateFunc func(event ObserverEvent)
}

func (o *RegionObserver) filteredCells(cells []Cell) []Cell {
	var filtered []Cell
	for _, cell := range cells {
		if cell.Inside(&o.region) {
			filtered = append(filtered, cell)
		}
	}
	return filtered
}

func (o *RegionObserver) Update(event ObserverEvent) {
	switch e := event.(type) {
	case SetCellObserverEvent:
		cell := e.Data()
		if !cell.Inside(&o.region) {
			return
		}
		o.updateFunc(e)
	case ClearCellObserverEvent:
		cell := e.Data()
		if !cell.Inside(&o.region) {
			return
		}
		o.updateFunc(e)
	case TickObserverEvent:
		bornCells, diedCells := e.Data()
		var filteredBorn, filteredDied []Cell = o.filteredCells(bornCells), o.filteredCells(diedCells)
		if len(filteredBorn) == 0 && len(filteredDied) == 0 {
			return
		}
		o.updateFunc(TickObserverEvent{
			bornCells: filteredBorn,
			diedCells: filteredDied,
		})
	case StepObserverEvent:
		n, bornCells, diedCells := e.Data()
		var filteredBorn, filteredDied []Cell = o.filteredCells(bornCells), o.filteredCells(diedCells)
		if len(filteredBorn) == 0 && len(filteredDied) == 0 {
			return
		}
		o.updateFunc(StepObserverEvent{
			n:         n,
			bornCells: filteredBorn,
			diedCells: filteredDied,
		})
	default:
		o.updateFunc(event)
	}
}

func (o *RegionObserver) SetRegion(region Rectangle) {
	o.region = region
}

func (o *RegionObserver) GetRegion() Rectangle {
	return o.region
}

func NewRegionObserver(region Rectangle, updateFunc func(event ObserverEvent)) *RegionObserver {
	return &RegionObserver{
		region:     region,
		updateFunc: updateFunc,
	}
}
