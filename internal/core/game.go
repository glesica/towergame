package core

import (
	"github.com/glesica/towergame/internal/tower"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"math"
)

type Game struct {
	enemies []*tower.State
	towers  []*tower.State
}

func NewGame() *Game {
	game := &Game{}
	return game
}

func (g *Game) AddEnemy(e *tower.State) {
	g.enemies = append(g.enemies, e)
}

func (g *Game) AddTower(t *tower.State) {
	g.towers = append(g.towers, t)
}

func (g *Game) Update() error {
	dt := 1.0 / float64(ebiten.TPS())

	s := tower.Simple{
		AimSpeed: 0.5 * math.Pi,
	}

	for _, t := range g.towers {
		t.Enemies = g.enemies
		t.Towers = g.towers

		s.Update(t, &tower.Instruction{
			Aim: 1,
		}, dt)
	}

	for _, e := range g.enemies {
		e.Enemies = g.enemies
		e.Towers = g.towers

		tower.Basic.Update(e, nil, dt)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Floralwhite)

	s := tower.Simple{
		AimSpeed: 0.3,
	}

	for _, t := range g.towers {
		s.Draw(t, screen)
	}

	for _, e := range g.enemies {
		tower.Basic.Draw(e, screen)
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	return 1024, 768
}
