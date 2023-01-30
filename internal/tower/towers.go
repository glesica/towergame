package tower

import (
	"github.com/glesica/towergame/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
	"math"
)

// TODO: Need a way to associate States and towers

var Basic = &Composed{
	Drawers: []Drawer{
		Circles,
	},
	Updaters: []Updater{
		Rotate(0.5 * math.Pi),
		Fire(0.25),
	},
}

type Composed struct {
	Drawers  []Drawer
	Updaters []Updater
}

func (c *Composed) Draw(s *state.Tower, screen *ebiten.Image) {
	for _, d := range c.Drawers {
		d(s, screen)
	}
}

func (c *Composed) Update(s *state.Tower, w *state.World, inst *state.Instruction, dt float64) {
	for _, u := range c.Updaters {
		u(s, w, inst, dt)
	}
}

// -------
// Drawers
// -------

func Circles(s *state.Tower, screen *ebiten.Image) {
	pos := s.Position
	ebitenutil.DrawCircle(screen, pos.X, pos.Y, 25, colornames.Aqua)
	ind := pos.Add(s.AimVector.Mult(20))
	ebitenutil.DrawCircle(screen, ind.X, ind.Y, 5, colornames.Fuchsia)
}

// --------
// Updaters
// --------

func Fire(fireDelay float64) Updater {
	return func(s *state.Tower, w *state.World, inst *state.Instruction, dt float64) {
		s.TimeToFire = math.Max(0, s.TimeToFire-dt)
		if inst.Fire && s.TimeToFire <= 0 {
			s.TimeToFire = fireDelay
			b := state.NewBullet()
			b.Position = s.Position
			// TODO: Figure out the bullet speed from the tower
			b.Velocity = s.AimVector.Normalize().Mult(30)
			// TODO: Figure out the amount of damage from the tower
			b.Damage = 10
			w.BulletQueue = append(w.BulletQueue, b)
		}
	}
}

func Rotate(aimSpeed float64) Updater {
	return func(s *state.Tower, w *state.World, inst *state.Instruction, dt float64) {
		if inst.Aim < 0 {
			s.Rotate(-aimSpeed * dt)
		} else if inst.Aim > 0 {
			s.Rotate(aimSpeed * dt)
		}
	}
}
