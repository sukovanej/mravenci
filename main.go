package main

import (
	"github.com/sukovanej/mravenci/src"
)

func main() {
	game := src.NewGame()
	controller := src.NewController(game)

	controller.Run()
}
