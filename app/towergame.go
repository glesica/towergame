package main

import (
	"github.com/glesica/towergame/internal/core"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Tower Game")

	game := core.NewGame()

	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
