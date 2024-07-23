package sdl

import (
	"fmt"
	"math/rand/v2"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth  = 1280
	windowHeight = 720
	windowTitle  = "Conway's Game of Life"
	cellSize     = 20
	rowsY        = windowHeight / cellSize
	columnsX     = windowWidth / cellSize
	fps          = 15
	fontSize     = 30
)

type game struct {
	window    *sdl.Window
	renderer  *sdl.Renderer
	fontColor sdl.Color

	pauseTextRectangle sdl.Rect
	pauseTextTexture   *sdl.Texture

	unpauseTextRectangle sdl.Rect
	unpauseTextTexture   *sdl.Texture

	spaceTextRectangle sdl.Rect
	spaceTextTexture   *sdl.Texture

	mouseTextRectangle sdl.Rect
	mouseTextTexture   *sdl.Texture

	paused    bool
	cells     [columnsX][rowsY]bool
	lastCellX int32
	lastCellY int32
	dragging  bool
}

func NewGame() *game {
	game := game{
		fontColor: sdl.Color{R: 255, B: 0, G: 0},
	}
	// randomize cells
	for x := range game.cells {
		for y := range game.cells[x] {
			game.cells[x][y] = rand.Float32() > 0.7
		}
	}
	return &game
}

func (g *game) Init() error {
	var err error

	g.window, err = sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, windowWidth, windowHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("error creating window: %v", err)
	}

	g.renderer, err = sdl.CreateRenderer(g.window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return fmt.Errorf("error creating renderer: %v", err)
	}

	return err
}

func (g *game) Close() {
	g.mouseTextTexture.Destroy()
	g.mouseTextTexture = nil

	g.spaceTextTexture.Destroy()
	g.spaceTextTexture = nil

	g.unpauseTextTexture.Destroy()
	g.unpauseTextTexture = nil

	g.pauseTextTexture.Destroy()
	g.pauseTextTexture = nil

	g.renderer.Destroy()
	g.renderer = nil

	g.window.Destroy()
	g.window = nil
}

func (g *game) Tick() {
	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch e := event.(type) {
			case *sdl.QuitEvent:
				return
			case *sdl.KeyboardEvent:
				if e.Type == sdl.KEYDOWN {
					switch e.Keysym.Scancode {
					case sdl.SCANCODE_ESCAPE:
						return
					case sdl.SCANCODE_SPACE:
						for col := range columnsX {
							for row := range rowsY {
								g.cells[col][row] = false
							}
						}
					case sdl.SCANCODE_P:
						g.paused = !g.paused
					}
				}
			case *sdl.MouseButtonEvent:
				if e.Type == sdl.MOUSEBUTTONDOWN {
					if e.Button == sdl.BUTTON_LEFT {
						cellX := e.X / cellSize
						cellY := e.Y / cellSize
						g.cells[cellX][cellY] = !g.cells[cellX][cellY]
						g.dragging = true
						g.lastCellX = cellX
						g.lastCellY = cellY
					}
				} else if e.Type == sdl.MOUSEBUTTONUP {
					g.dragging = false
					g.lastCellX = -1
					g.lastCellY = -1
				}
			case *sdl.MouseMotionEvent:
				if g.dragging {
					cellX := e.X / cellSize
					cellY := e.Y / cellSize

					if cellX != g.lastCellX || cellY != g.lastCellY {
						g.cells[cellX][cellY] = !g.cells[cellX][cellY]
						g.lastCellX = cellX
						g.lastCellY = cellY
					}
				}
			}
		}

		if !g.paused {
			g.update()
		}
		g.Draw()
		sdl.Delay(uint32(1000 / fps))

	}
}
