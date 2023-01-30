package state

import (
	"github.com/jakecoffman/cp"
	"sync/atomic"
)

// TODO: Consider using this as a value type instead of a pointer

var nextBulletID atomic.Uint64

type Bullet struct {
	// Unique ID to facilitate lookups.
	id uint64

	// TODO: This might move to an updater parameter
	Damage float64

	// Expired indicates whether the bullet has exceeded
	// its maximum range or should otherwise be removed
	// from the game.
	Expired bool

	Position cp.Vector

	// Traveled is the total distance the bullet has
	// moved since it was fired.
	Traveled float64

	Velocity cp.Vector
}

func NewBullet() *Bullet {
	id := nextBulletID.Add(1)
	b := &Bullet{
		id: id,
	}
	return b
}

func (b *Bullet) ID() uint64 {
	return b.id
}

func (b *Bullet) Pos() cp.Vector {
	return b.Position
}
