package bullet

import (
	"github.com/glesica/towergame/internal/state"
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawer func(state *state.Bullet, screen *ebiten.Image)

type Updater func(state *state.Bullet, world *state.World, dt float64)

type T interface {
	Draw(state *state.Bullet, screen *ebiten.Image)

	Update(state *state.Bullet, world *state.World, dt float64)
}
