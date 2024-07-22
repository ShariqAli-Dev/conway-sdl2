package sdl

import (
	"fmt"
	"math/rand"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	cellSize = 20
	rows     = windowHeight / cellSize
	columns  = windowWidth / cellSize
)

type game struct {
	window   *sdl.Window
	renderer *sdl.Renderer
	paused   bool
	cells    [columns][rows]bool // (x,y), (columns, rows)
}

func (g *game) randomizeCells() {
	for x := range g.cells {
		for y := range g.cells[x] {
			g.cells[x][y] = rand.Float32() > 0.7
		}
	}
}

func NewGame() *game {
	game := game{}
	game.randomizeCells()
	return &game
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
					case sdl.SCANCODE_P:
						if g.paused {
							fmt.Println("unpaused")
						} else {

							fmt.Println("paused")
						}
						g.paused = !g.paused

					}
				}
			}
		}
		g.Draw()
		sdl.Delay(uint32(1000 / 60))
	}
}

func (g *game) Draw() {
	g.renderer.SetDrawColor(255, 255, 255, 255) // white
	g.renderer.Clear()

	for x := 0; x < windowWidth; x += int(cellSize) {
		for y := 0; y < windowHeight; y += int(cellSize) {
			cell := sdl.Rect{
				X: int32(x),
				Y: int32(y),
				W: int32(cellSize),
				H: int32(cellSize),
			}

			if g.cells[x/cellSize][y/cellSize] {
				g.renderer.SetDrawColor(0, 0, 0, 255) // black
				g.renderer.FillRect(&cell)
			}

		}

	}

	g.renderer.Present()
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
	g.renderer.Destroy()
	g.renderer = nil

	g.window.Destroy()
	g.window = nil
}
