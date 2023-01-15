package tower

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/colornames"
)

// TODO: Need a way to associate States and towers
// TODO: Might be able to have just one tower implementation with customization

type Simple struct{}

func (c *Simple) Draw(state *State, screen *ebiten.Image) {
	pos := state.Position
	ebitenutil.DrawCircle(screen, pos.X, pos.Y, 25, colornames.Aqua)
	ind := pos.Add(state.Heading.Mult(20))
	ebitenutil.DrawCircle(screen, ind.X, ind.Y, 5, colornames.Fuchsia)
}

func (c *Simple) Update(state *State, inst *Instruction, dt float64) {
	//TODO implement me
	panic("implement me")
}
