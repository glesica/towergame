package state

import (
	"github.com/jakecoffman/cp"
	"math"
	"sync/atomic"
)

// TODO: Consider using this as a value type instead of a pointer

var nextTowerID atomic.Uint64

type Tower struct {
	id uint64

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
	id := nextTowerID.Add(1)
	s := &Tower{
		id: id,
	}

	s.Rotate(0)
	s.Position = cp.Vector{}

	return s
}

func (t *Tower) ID() uint64 {
	return t.id
}

// Rotate causes the heading angle and heading vector to update to
// reflect a rotation by the given amount, in radians.
func (t *Tower) Rotate(angleDelta float64) {
	t.AimAngle = math.Mod(t.AimAngle+angleDelta, 2*math.Pi)
	if t.AimAngle < 0 {
		t.AimAngle = 2*math.Pi + t.AimAngle
	}
	t.AimVector = cp.ForAngle(t.AimAngle)
}

// Toward returns a unit vector pointing from the receiver state,
// based on its position, to the given position.
func (t *Tower) Toward(position cp.Vector) cp.Vector {
	return cp.Vector{
		X: position.X - t.Position.X,
		Y: position.Y - t.Position.Y,
	}.Normalize()
}

func (t *Tower) Pos() cp.Vector {
	return t.Position
}
