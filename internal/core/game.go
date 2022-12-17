package core

import (
	"fmt"
	"github.com/glesica/towergame/internal/runtime"
	"github.com/glesica/towergame/internal/tower"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
)

type Game struct {
	space  *cp.Space
	towers []tower.T
}

func NewGame() *Game {
	space := cp.NewSpace()
	space.Iterations = 1

	game := &Game{
		space: space,
	}

	t0 := tower.NewBasic(250, 300, 0)
	game.AddTower(t0)

	t1 := tower.NewBasic(525, 300, 90)
	game.AddTower(t1)

	//bullet := projectile.NewBullet(300, 350, 0, 10, 10)
	//game.AddTower(bullet)

	return game
}

func (g *Game) Update() error {
	dt := 1.0 / float64(ebiten.TPS())

	g.space.Step(dt)

	for _, t := range g.towers {
		runtime.RunTower(t)

		err := t.Update(dt)
		if err != nil {
			return fmt.Errorf("tower update error: %w", err)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(colornames.Floralwhite)

	for _, t := range g.towers {
		t.Draw(screen)
	}
}

func (g *Game) Layout(width, height int) (int, int) {
	return 1024, 768
}

func (g *Game) AddTower(e tower.T) {
	e.AddToSpace(g.space)
	g.towers = append(g.towers, e)
}
