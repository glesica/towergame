package enemy

import (
	"github.com/glesica/towergame/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

var Basic = Enemy{
	Drawers: []Drawer{
		DrawDot,
	},
	Updaters: []Updater{
		MoveToNearestTower,
	},
}

type Enemy struct {
	Drawers  []Drawer
	Updaters []Updater
}

func (e *Enemy) Draw(state *state.Enemy, screen *ebiten.Image) {
	for _, d := range e.Drawers {
		d(state, screen)
	}
}

func (e *Enemy) Update(state *state.Enemy, world *state.World, dt float64) {
	for _, u := range e.Updaters {
		u(state, world, dt)
	}
}

func DrawDot(state *state.Enemy, screen *ebiten.Image) {
	pos := state.Position
	ebitenutil.DrawCircle(screen, pos.X, pos.Y, 5, colornames.Aqua)
}

func MoveToNearestTower(state *state.Enemy, world *state.World, dt float64) {
	t := world.NearestTower(state.Position)
	if t != nil {
		delta := state.Toward(t.Position).Mult(state.MovementSpeed * dt)
		state.Position = state.Position.Add(delta)
	}
}
