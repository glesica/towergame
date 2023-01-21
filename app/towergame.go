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
	s := tower.NewState()
	s.Position = cp.Vector{X: 450, Y: 300}
	game.AddTower(s)

	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
