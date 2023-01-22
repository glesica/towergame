package tower

import (
	"github.com/jakecoffman/cp"
	"math"
)

type State struct {
	// HeadingAngle is the tower's current heading as an angle,
	// in the interval [0, 2PI) where 0 is pointing east (right).
	HeadingAngle float64

	// HeadingVector is the tower's current heading, generally the
	// direction the gun is pointing or, for enemies, the direction
	// of movement. This should be a unit vector.
	HeadingVector cp.Vector

	// Position is the location of the tower in 2D space.
	Position cp.Vector

	// TimeToFire is the amount of time, in seconds, until the tower
	// is ready to fire its gun.
	TimeToFire float64

	// MovementSpeed is the rate at which the entity moves, in units
	// per second.
	// TODO: Should this be on the updater instead?
	MovementSpeed float64

	// TODO: Replace the slices with R-trees or something else fast

	// Enemies holds a reference to the current slice of enemy states,
	// it contains all existing enemies.
	Enemies []*State

	// Towers holds a reference to the current slice of tower states,
	// it contains all existing towers.
	Towers []*State
}

func NewState() *State {
	s := &State{}

	s.Rotate(0)
	s.Position = cp.Vector{}

	return s
}

// Rotate causes the heading angle and heading vector to update to
// reflect a rotation by the given amount, in radians.
func (s *State) Rotate(angleDelta float64) {
	s.HeadingAngle = math.Mod(s.HeadingAngle+angleDelta, 2*math.Pi)
	if s.HeadingAngle < 0 {
		s.HeadingAngle = 2*math.Pi + s.HeadingAngle
	}
	s.HeadingVector = cp.ForAngle(s.HeadingAngle)
}

func (s *State) NearbyEnemies(radius float64) []*State {
	return nil
}

func (s *State) NearbyTowers(radius float64) []*State {
	return nil
}

func (s *State) NearestEnemy() *State {
	return nil
}

func (s *State) NearestTower() *State {
	if len(s.Towers) > 0 {
		// TODO: Make this actually do something
		return s.Towers[0]
	}

	return nil
}

// Toward returns a unit vector pointing from the receiver state,
// based on its position, to the given position.
func (s *State) Toward(position cp.Vector) cp.Vector {
	return cp.Vector{
		X: position.X - s.Position.X,
		Y: position.Y - s.Position.Y,
	}.Normalize()
}
