package tower

import (
	"github.com/glesica/towergame/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawer func(state *state.Tower, screen *ebiten.Image)

type Updater func(state *state.Tower, world *state.World, inst *state.Instruction, dt float64)

type T interface {
	// Draw the tower to the screen.
	Draw(state *state.Tower, screen *ebiten.Image)

	// Update mutates the given state based on the given instruction
	// to account for changes that should happen over the given
	// time delta.
	Update(state *state.Tower, world *state.World, inst *state.Instruction, dt float64)
}
