package main

import (
	"github.com/glesica/towergame/internal/core"
	"github.com/glesica/towergame/internal/tower"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

func main() {
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Tower Game")

	game := core.NewGame()
	game.AddTower(&tower.State{
		HeadingVector: cp.Vector{1, 0},
		Position:      cp.Vector{450, 300},
		TimeToFire:    0,
	})

	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
