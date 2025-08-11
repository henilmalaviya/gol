package grid

import (
	"encoding/binary"
	"hash/fnv"
	"sort"
)

// Hash returns a stable 64-bit FNV-1a hash of the set of live cells.
// The hash is deterministic across runs and independent of map iteration order.
func (g *Grid) Hash() uint64 {
	g.mutex.RLock()
	defer g.mutex.RUnlock()

	if len(g.cells) == 0 {
		return 0
	}

	cells := make([]Cell, 0, g.Population())
	for c := range g.cells {
		cells = append(cells, c)
	}

	sort.Slice(cells, func(i, j int) bool {
		if cells[i].Y == cells[j].Y {
			return cells[i].X < cells[j].X
		}
		return cells[i].Y < cells[j].Y
	})

	h := fnv.New64a()
	var buf [16]byte
	for _, c := range cells {
		binary.LittleEndian.PutUint64(buf[0:], uint64(int64(c.X)))
		binary.LittleEndian.PutUint64(buf[8:], uint64(int64(c.Y)))
		_, _ = h.Write(buf[:])
	}
	return h.Sum64()
}
