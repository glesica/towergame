package enemy

import (
	"github.com/glesica/towergame/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
)

var Basic = &Composed{
	Drawers: []Drawer{
		Dot,
	},
	Updaters: []Updater{
		MoveToNearestTower(10),
	},
}

type Composed struct {
	Drawers  []Drawer
	Updaters []Updater
}

func (c *Composed) Draw(state *state.Enemy, screen *ebiten.Image) {
	for _, d := range c.Drawers {
		d(state, screen)
	}
}

func (c *Composed) Update(state *state.Enemy, world *state.World, dt float64) {
	for _, u := range c.Updaters {
		u(state, world, dt)
	}
}

// -------
// Drawers
// -------

func Dot(state *state.Enemy, screen *ebiten.Image) {
	pos := state.Position
	ebitenutil.DrawCircle(screen, pos.X, pos.Y, 5, colornames.Aqua)
}

// --------
// Updaters
// --------

func MoveToNearestTower(speed float64) Updater {
	return func(s *state.Enemy, w *state.World, dt float64) {
		t := w.Towers.Nearest(s.Position)
		if t != nil {
			delta := s.Toward(t.Position).Mult(speed * dt)
			s.Position = s.Position.Add(delta)
		}
	}
}

func MoveToPoint(speed float64, point cp.Vector) Updater {
	return func(s *state.Enemy, w *state.World, dt float64) {
		delta := s.Toward(point).Mult(speed * dt)
		s.Position = s.Position.Add(delta)
	}
}
