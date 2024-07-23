package main

import (
	"log"

	"github.com/shariqali-dev/conway-sdl2/cmd/sdl"
)

func main() {
	defer sdl.Close()
	if err := sdl.Init(); err != nil {
		log.Fatal(error.Error(err))
	}

	game := sdl.NewGame()
	defer game.Close()
	if err := game.Init(); err != nil {
		log.Fatal(error.Error(err))
	}

	game.Tick()
}
