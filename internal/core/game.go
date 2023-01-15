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
	dt := 1.0 / float64(ebiten.TPS())

	for _, t := range g.towers {
		s := tower.Simple{
			AimSpeed: 0.3,
		}
		s.Update(t, &tower.Instruction{
			Aim: 1,
		}, dt)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Floralwhite)

	for _, t := range g.towers {
		s := tower.Simple{
			AimSpeed: 0.3,
		}
		s.Draw(t, screen)
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	return 1024, 768
}
