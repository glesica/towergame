package tower

import "github.com/glesica/towergame/internal/entity"

type T interface {
	entity.T

	// Aim instructs the tower to aim in a given direction. The return
	// value is the number of seconds until the aiming operation will
	// be complete.
	Aim(heading float64) float64

	// Heading returns the direction the tower is currently aiming.
	Heading() float64

	// Fire instructs the tower to fire its weapon. The return value is
	// the number of seconds until firing will actually happen.
	Fire() float64
}
