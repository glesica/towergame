package entity

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/jakecoffman/cp"
)

type T interface {
	AddToSpace(space *cp.Space)
	RemoveFromSpace(space *cp.Space)
	Update(dt float64) error
	Draw(screen *ebiten.Image)
}
