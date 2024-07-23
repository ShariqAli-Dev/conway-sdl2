package sdl

import "github.com/veandco/go-sdl2/sdl"

func (g *game) Draw() {
	g.renderer.SetDrawColor(255, 255, 255, 255) // white
	g.renderer.Clear()

	g.drawGrid()
	if g.paused {
		g.renderer.Copy(g.unpauseTextTexture, nil, &g.unpauseTextRectangle)
	} else {
		g.renderer.Copy(g.pauseTextTexture, nil, &g.pauseTextRectangle)
	}
	g.renderer.Copy(g.spaceTextTexture, nil, &g.spaceTextRectangle)
	g.renderer.Copy(g.mouseTextTexture, nil, &g.mouseTextRectangle)

	g.renderer.Present()
}

func (g *game) drawGrid() {
	for xPixel := 0; xPixel < windowWidth; xPixel += int(cellSize) {
		for yPixel := 0; yPixel < windowHeight; yPixel += int(cellSize) {
			cell := sdl.Rect{
				X: int32(xPixel),
				Y: int32(yPixel),
				W: int32(cellSize),
				H: int32(cellSize),
			}

			if g.cells[xPixel/cellSize][yPixel/cellSize] {
				g.renderer.SetDrawColor(0, 0, 0, 255) // black
				g.renderer.FillRect(&cell)
			}

		}

	}
}
