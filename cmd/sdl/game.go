package sdl

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type game struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

func NewGame() *game {
	return &game{}
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
						// g.randColor()
					case sdl.SCANCODE_P:
						// g.pauseMusic()
					}
				}
			}
		}

		// g.updateText()
		// g.updateSprite()

		g.renderer.Clear()

		// g.renderer.Copy(g.backgroundImage, nil, nil)
		// g.renderer.Copy(g.textImage, nil, &g.textRectangle)
		// g.renderer.Copy(g.spriteImage, nil, &g.spriteRectangle)

		g.renderer.Present()
		sdl.Delay(uint32(1000 / 60))

	}
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
