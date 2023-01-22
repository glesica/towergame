package enemy

import (
	"github.com/glesica/towergame/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawer func(state *state.Enemy, screen *ebiten.Image)

type Updater func(state *state.Enemy, world *state.World, dt float64)

type T interface {
	// Draw the enemy to the screen.
	Draw(state *state.Enemy, screen *ebiten.Image)

	// Update mutates the given state
	// in order to apply the change that should occur over the given
	// time delta, for this particular enemy. Since different enemies
	// have different properties and capabilities, each one will
	// respond differently given a particular state.
	Update(state *state.Enemy, world *state.World, dt float64)
}
