package sdl

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func (g *game) LoadMedia() error {
	var err error
	font, err := ttf.OpenFont("./assets/fonts/freesansbold.ttf", fontSize)
	if err != nil {
		return fmt.Errorf("error opening font: %v", err)
	}
	defer font.Close()

	// pause text
	pauseFontSurface, err := font.RenderUTF8Blended("[P] to pause", g.fontColor)
	if err != nil {
		return fmt.Errorf("error creating text surface: %v", err)
	}
	defer pauseFontSurface.Free()
	if g.pauseTextTexture, err = g.renderer.CreateTextureFromSurface(pauseFontSurface); err != nil {
		return fmt.Errorf("could not create texture: %v", err)
	}
	g.pauseTextRectangle = sdl.Rect{
		W: pauseFontSurface.W,
		H: pauseFontSurface.H,
	}

	// unpause text
	unpauseFontSurface, err := font.RenderUTF8Blended("[P] to unpause", g.fontColor)
	if err != nil {
		return fmt.Errorf("error creating text surface: %v", err)
	}
	defer unpauseFontSurface.Free()
	if g.unpauseTextTexture, err = g.renderer.CreateTextureFromSurface(unpauseFontSurface); err != nil {
		return fmt.Errorf("could not create texture: %v", err)
	}
	g.unpauseTextRectangle = sdl.Rect{
		W: unpauseFontSurface.W,
		H: unpauseFontSurface.H,
	}

	// unpause text
	spaceFontSurface, err := font.RenderUTF8Blended("[Space] to clear", g.fontColor)
	if err != nil {
		return fmt.Errorf("error creating text surface: %v", err)
	}
	defer spaceFontSurface.Free()
	if g.spaceTextTexture, err = g.renderer.CreateTextureFromSurface(spaceFontSurface); err != nil {
		return fmt.Errorf("could not create texture: %v", err)
	}
	g.spaceTextRectangle = sdl.Rect{
		W: spaceFontSurface.W,
		H: spaceFontSurface.H,
		Y: unpauseFontSurface.H,
	}

	// mouse text
	mouseFontSurface, err := font.RenderUTF8Blended("[Click] or [Drag] mouse to toggle cell", g.fontColor)
	if err != nil {
		return fmt.Errorf("error creating text surface: %v", err)
	}
	defer mouseFontSurface.Free()
	if g.mouseTextTexture, err = g.renderer.CreateTextureFromSurface(mouseFontSurface); err != nil {
		return fmt.Errorf("could not create texture: %v", err)
	}
	g.mouseTextRectangle = sdl.Rect{
		W: mouseFontSurface.W,
		H: mouseFontSurface.H,
		Y: unpauseFontSurface.H + spaceFontSurface.H,
	}

	return err
}
