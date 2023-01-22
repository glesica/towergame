package tower

import (
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

func (e *Enemy) Draw(state *State, screen *ebiten.Image) {
	for _, d := range e.Drawers {
		d(state, screen)
	}
}

func (e *Enemy) Update(state *State, inst *Instruction, dt float64) {
	for _, u := range e.Updaters {
		u(state, inst, dt)
	}
}

func DrawDot(state *State, screen *ebiten.Image) {
	pos := state.Position
	ebitenutil.DrawCircle(screen, pos.X, pos.Y, 5, colornames.Aqua)
}

func MoveToNearestTower(state *State, inst *Instruction, dt float64) {
	t := state.NearestTower()
	if t != nil {
		delta := state.Toward(t.Position).Mult(state.MovementSpeed * dt)
		state.Position = state.Position.Add(delta)
	}
}
