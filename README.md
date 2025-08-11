# gol

A small, concurrent-safe core library for Conway's Game of Life in Go.

- Module: `github.com/henilmalaviya/gol`
- Core grid API subpackage: `github.com/henilmalaviya/gol/grid`
- Thread-safe reads/writes via RWMutex
- Infinite grid using a sparse set of live cells

## Install

```bash
go get github.com/henilmalaviya/gol
```

## Quick start

```go
package main

import (
	"fmt"
	gol "github.com/henilmalaviya/gol"
)

func main() {
	g := gol.NewGame()

    // advance one generation
    g.GetGrid().Tick()
}
```

## Testing

```bash
go test ./...
```
