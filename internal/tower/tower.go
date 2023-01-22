package tower

import "github.com/hajimehoshi/ebiten/v2"

type Drawer func(state *State, screen *ebiten.Image)

type Updater func(state *State, inst *Instruction, dt float64)

type T interface {
	// Draw the tower to the screen.
	Draw(state *State, screen *ebiten.Image)

	// Update mutates the given state based on the given instruction
	// in order to apply the change that should occur over the given
	// time delta, for this particular tower. Since different towers
	// have different properties and capabilities, each tower will
	// respond differently given a particular state and instruction.
	Update(state *State, inst *Instruction, dt float64)
}
