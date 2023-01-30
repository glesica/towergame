package main

import (
	"github.com/glesica/towergame/internal/core"
	"github.com/glesica/towergame/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

func main() {
	ebiten.SetWindowSize(1024, 768)
	ebiten.SetWindowTitle("Tower Game")

	w := state.NewWorld()

	t := state.NewTower()
	t.Position = cp.Vector{X: 450, Y: 300}
	w.Towers.Add(t)

	e := state.NewEnemy()
	e.Position = cp.Vector{X: 700, Y: 500}
	e.MovementSpeed = 100
	w.Enemies.Add(e)

	game := core.NewGame(w)

	err := ebiten.RunGame(game)
	if err != nil {
		panic(err)
	}
}
