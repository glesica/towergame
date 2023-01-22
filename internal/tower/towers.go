package tower

import (
	"github.com/glesica/towergame/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
	"math"
)

// TODO: Need a way to associate States and towers

var Basic = Composed{
	Drawers: []Drawer{
		DrawCircles,
	},
	Updaters: []Updater{
		UpdateAim,
	},
}

type Composed struct {
	Drawers  []Drawer
	Updaters []Updater
}

func (c *Composed) Draw(state *state.Tower, screen *ebiten.Image) {
	for _, d := range c.Drawers {
		d(state, screen)
	}
}

func (c *Composed) Update(state *state.Tower, world *state.World, inst *state.Instruction, dt float64) {
	for _, u := range c.Updaters {
		u(state, world, inst, dt)
	}
}

func DrawCircles(state *state.Tower, screen *ebiten.Image) {
	pos := state.Position
	ebitenutil.DrawCircle(screen, pos.X, pos.Y, 25, colornames.Aqua)
	ind := pos.Add(cp.Vector{X: state.AimVector.X, Y: -state.AimVector.Y}.Mult(20))
	ebitenutil.DrawCircle(screen, ind.X, ind.Y, 5, colornames.Fuchsia)
}

func UpdateAim(state *state.Tower, world *state.World, inst *state.Instruction, dt float64) {
	if inst.Aim < 0 {
		state.Rotate(-state.AimSpeed * dt)
	} else if inst.Aim > 0 {
		state.Rotate(state.AimSpeed * dt)
	}

	state.TimeToFire = math.Max(0, state.TimeToFire-dt)
}
