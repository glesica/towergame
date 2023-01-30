package state

import (
	"github.com/jakecoffman/cp"
	"math"
	"sync/atomic"
)

// TODO: Consider using this as a value type instead of a pointer

var nextEnemyID atomic.Uint64

type Enemy struct {
	id uint64

	// HeadingAngle is the tower's current heading as an angle,
	// in the interval [0, 2PI) where 0 is pointing east (right).
	HeadingAngle float64

	// HeadingVector is the tower's current heading, generally the
	// direction the gun is pointing or, for enemies, the direction
	// of movement. This should be a unit vector.
	HeadingVector cp.Vector

	// Health is the amount of damage the enemy can take before
	// it is removed from the game (killed).
	Health int64

	// Position is the location of the tower in 2D space.
	Position cp.Vector

	// MovementSpeed is the rate at which the entity moves, in units
	// per second.
	// TODO: Should this be on the updater instead?
	MovementSpeed float64
}

func NewEnemy() *Enemy {
	id := nextTowerID.Add(1)
	s := &Enemy{}
	s.id = id

	s.Rotate(0)
	s.Position = cp.Vector{}

	return s
}

func (e *Enemy) Dead() bool {
	return e.Health <= 0
}

// Rotate causes the heading angle and heading vector to update to
// reflect a rotation by the given amount, in radians.
func (e *Enemy) Rotate(angleDelta float64) {
	e.HeadingAngle = math.Mod(e.HeadingAngle+angleDelta, 2*math.Pi)
	if e.HeadingAngle < 0 {
		e.HeadingAngle = 2*math.Pi + e.HeadingAngle
	}
	e.HeadingVector = cp.ForAngle(e.HeadingAngle)
}

// Toward returns a unit vector pointing from the receiver state,
// based on its position, to the given position.
func (e *Enemy) Toward(position cp.Vector) cp.Vector {
	return cp.Vector{
		X: position.X - e.Position.X,
		Y: position.Y - e.Position.Y,
	}.Normalize()
}

func (e *Enemy) ID() uint64 {
	return e.id
}

func (e *Enemy) Pos() cp.Vector {
	return e.Position
}
