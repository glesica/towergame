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

	t := tower.NewState()
	t.Position = cp.Vector{X: 450, Y: 300}
	game.AddTower(t)

	e := tower.NewState()
	e.Position = cp.Vector{X: 700, Y: 500}
	e.MovementSpeed = 100
	game.AddEnemy(e)

	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
