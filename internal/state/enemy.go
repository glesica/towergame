package state

import (
	"github.com/jakecoffman/cp"
	"math"
)

type Enemy struct {
	// HeadingAngle is the tower's current heading as an angle,
	// in the interval [0, 2PI) where 0 is pointing east (right).
	HeadingAngle float64

	// HeadingVector is the tower's current heading, generally the
	// direction the gun is pointing or, for enemies, the direction
	// of movement. This should be a unit vector.
	HeadingVector cp.Vector

	// Position is the location of the tower in 2D space.
	Position cp.Vector

	// MovementSpeed is the rate at which the entity moves, in units
	// per second.
	// TODO: Should this be on the updater instead?
	MovementSpeed float64
}

func NewEnemy() *Enemy {
	s := &Enemy{}

	s.Rotate(0)
	s.Position = cp.Vector{}

	return s
}

// Rotate causes the heading angle and heading vector to update to
// reflect a rotation by the given amount, in radians.
func (s *Enemy) Rotate(angleDelta float64) {
	s.HeadingAngle = math.Mod(s.HeadingAngle+angleDelta, 2*math.Pi)
	if s.HeadingAngle < 0 {
		s.HeadingAngle = 2*math.Pi + s.HeadingAngle
	}
	s.HeadingVector = cp.ForAngle(s.HeadingAngle)
}

// Toward returns a unit vector pointing from the receiver state,
// based on its position, to the given position.
func (s *Enemy) Toward(position cp.Vector) cp.Vector {
	return cp.Vector{
		X: position.X - s.Position.X,
		Y: position.Y - s.Position.Y,
	}.Normalize()
}
