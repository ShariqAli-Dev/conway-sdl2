package sdl

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth  = 800
	windowHeight = 500
	windowTitle  = "Conway's Game of Life"
)

func Init() error {
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("error initializing sdl2: %v", err)
	}
	return err

}

func Close() {
	sdl.Quit()
}
