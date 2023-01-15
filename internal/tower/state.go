package tower

import (
	"github.com/jakecoffman/cp"
	"math"
)

type State struct {
	// HeadingAngle is the tower's current heading as an angle,
	// in the interval [0, 2PI) where 0 is pointing east (right).
	HeadingAngle float64

	// The tower's current heading, generally the direction the gun
	// is pointing. This should be a unit vector.
	HeadingVector cp.Vector

	// Position is the location of the tower in 2D space.
	Position cp.Vector

	// The amount of time, in milliseconds, until the tower is ready
	// to fire its gun.
	TimeToFire uint
}

// Rotate causes the heading angle and heading vector to update to
// reflect a rotation by the given amount, in radians.
func (s *State) Rotate(angleDelta float64) {
	// TODO: This won't work for negative angles, maybe do [-PI, PI)?
	s.HeadingAngle = math.Mod(s.HeadingAngle+angleDelta, 2*math.Pi)
	s.HeadingVector = cp.Vector{
		X: math.Cos(s.HeadingAngle),
		Y: math.Sin(s.HeadingAngle),
	}
}
