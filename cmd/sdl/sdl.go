package sdl

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func Init() error {
	var err error

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return fmt.Errorf("error initializing sdl2: %v", err)
	}

	if err = ttf.Init(); err != nil {
		return fmt.Errorf("error initializing ttf: %v", err)
	}

	return err

}

func Close() {
	ttf.Quit()
	sdl.Quit()
}
