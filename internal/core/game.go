package core

import (
	"github.com/glesica/towergame/internal/enemy"
	"github.com/glesica/towergame/internal/state"
	"github.com/glesica/towergame/internal/tower"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
	"math"
)

type Game struct {
	enemies []*state.Enemy
	towers  []*state.Tower
}

func NewGame() *Game {
	game := &Game{}
	return game
}

func (g *Game) AddEnemy(e *state.Enemy) {
	g.enemies = append(g.enemies, e)
}

func (g *Game) AddTower(t *state.Tower) {
	g.towers = append(g.towers, t)
}

func (g *Game) Update() error {
	dt := 1.0 / float64(ebiten.TPS())

	s := tower.Simple{
		AimSpeed: 0.5 * math.Pi,
	}

	w := &state.World{
		Enemies: g.enemies,
		Towers:  g.towers,
	}

	for _, t := range g.towers {
		s.Update(t, w, &state.Instruction{
			Aim: 1,
		}, dt)
	}

	for _, e := range g.enemies {
		enemy.Basic.Update(e, w, dt)
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
		enemy.Basic.Draw(e, screen)
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	return 1024, 768
}
