package state

import (
	"github.com/jakecoffman/cp"
	"math"
)

type Tower struct {
	// AimAngle is the tower's current heading as an angle,
	// in the interval [0, 2PI) where 0 is pointing east (right).
	AimAngle float64

	// AimVector is the tower's current heading, generally the
	// direction the gun is pointing or, for enemies, the direction
	// of movement. This should be a unit vector.
	AimVector cp.Vector

	// Position is the location of the tower in 2D space.
	Position cp.Vector

	// TimeToFire is the amount of time, in seconds, until the tower
	// is ready to fire its gun.
	TimeToFire float64
}

func NewTower() *Tower {
	s := &Tower{}

	s.Rotate(0)
	s.Position = cp.Vector{}

	return s
}

// Rotate causes the heading angle and heading vector to update to
// reflect a rotation by the given amount, in radians.
func (w *Tower) Rotate(angleDelta float64) {
	w.AimAngle = math.Mod(w.AimAngle+angleDelta, 2*math.Pi)
	if w.AimAngle < 0 {
		w.AimAngle = 2*math.Pi + w.AimAngle
	}
	w.AimVector = cp.ForAngle(w.AimAngle)
}

// Toward returns a unit vector pointing from the receiver state,
// based on its position, to the given position.
func (w *Tower) Toward(position cp.Vector) cp.Vector {
	return cp.Vector{
		X: position.X - w.Position.X,
		Y: position.Y - w.Position.Y,
	}.Normalize()
}
