package core

import (
	"github.com/glesica/towergame/internal/tower"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Game struct {
	towers []*tower.State
}

func NewGame() *Game {
	game := &Game{}
	return game
}

func (g *Game) AddTower(t *tower.State) {
	g.towers = append(g.towers, t)
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Floralwhite)

	for _, t := range g.towers {
		s := tower.Simple{}
		s.Draw(t, screen)
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	return 1024, 768
}
