package sdl

func (g *game) update() {
	nextCells := new([columnsX][rowsY]bool)

	for colX := range columnsX {
		for rowY := range rowsY {
			aliveNeighbors := g.cellGetAliveNeighbors(colX, rowY)

			if g.cells[colX][rowY] {
				if aliveNeighbors < 2 || aliveNeighbors > 3 {
					nextCells[colX][rowY] = false
				} else {
					nextCells[colX][rowY] = true

				}
			} else {
				if aliveNeighbors == 3 {
					nextCells[colX][rowY] = true
				}
			}
		}
	}

	g.cells = *nextCells
}

func (g *game) cellGetAliveNeighbors(xCol, yRow int) (aliveNeighbors int) {
	// virtual 3x3 grid, xCol and yRow is 0,0
	for x := -1; x < 2; x++ {
		for y := -1; y < 2; y++ {
			if y == 0 && x == 0 {
				continue
			}

			xNeighbor := x + xCol
			yNeighbor := y + yRow

			// check if the coordinate is in bounds
			xInBounds := xNeighbor >= 0 && xNeighbor < (windowWidth/cellSize)
			_ = xInBounds
			yInBounds := yNeighbor >= 0 && yNeighbor < (windowHeight/cellSize)

			if !xInBounds || !yInBounds {
				continue
			}

			if g.cells[xNeighbor][yNeighbor] {
				aliveNeighbors++
			}
		}
	}
	return aliveNeighbors
}
