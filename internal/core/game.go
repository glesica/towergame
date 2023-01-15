package core

import (
	"fmt"
	"github.com/glesica/towergame/internal/runtime"
	"github.com/glesica/towergame/internal/tower"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
)

type Game struct {
	space  *cp.Space
	towers []*tower.T
}

func NewGame() *Game {
	space := cp.NewSpace()
	space.Iterations = 1

	game := &Game{
		space: space,
	}

	return game
}

func (g *Game) AddTower(x, y float64) {
	t := tower.NewCircle(x, y, 25)
	t.Spacer.AddToSpace(g.space)
	g.towers = append(g.towers, t)
}

func (g *Game) RemoveTower(x, y float64) {}

func (g *Game) Update() error {
	dt := 1.0 / float64(ebiten.TPS())
	g.space.Step(dt)

	// Tower Programs

	for _, t := range g.towers {
		runtime.RunTower(t)

		err := t.Update(dt)
		if err != nil {
			return fmt.Errorf("tower update error: %w", err)
		}
	}

	// Input Handling

	clicked := inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
	if clicked {
		x, y := ebiten.CursorPosition()
		fmt.Printf("%d, %d\n", x, y)
		g.AddTower(float64(x), float64(y))
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
