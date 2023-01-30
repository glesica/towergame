package bullet

import (
	"github.com/glesica/towergame/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	"image/color"
)

var Basic = &Composed{
	Drawers: []Drawer{
		DrawDot(3, colornames.Blueviolet),
	},
	Updaters: []Updater{
		ConstantSpeed(),
		FixedRange(100),
	},
}

type Composed struct {
	Drawers  []Drawer
	Updaters []Updater
}

func (c *Composed) Draw(s *state.Bullet, img *ebiten.Image) {
	for _, d := range c.Drawers {
		d(s, img)
	}
}

func (c *Composed) Update(s *state.Bullet, w *state.World, dt float64) {
	for _, u := range c.Updaters {
		u(s, w, dt)
	}
}

func DrawDot(r float64, c color.Color) Drawer {
	return func(state *state.Bullet, screen *ebiten.Image) {
		pos := state.Position
		ebitenutil.DrawCircle(screen, pos.X, pos.Y, r, c)
	}
}

func ConstantSpeed() Updater {
	return func(s *state.Bullet, w *state.World, dt float64) {
		delta := s.Velocity.Mult(dt)
		s.Position = s.Position.Add(delta)
		s.Traveled += delta.Length()
	}
}

func FixedRange(r float64) Updater {
	return func(s *state.Bullet, w *state.World, dt float64) {
		if s.Traveled >= r {
			s.Expired = true
		}
	}
}
