package state

import (
	"github.com/jakecoffman/cp"
)

type World struct {
	// TODO: Replace the slices with R-trees or something else fast

	// Enemies holds a reference to the current slice of enemy states,
	// it contains all existing enemies.
	Enemies []*Enemy

	// Towers holds a reference to the current slice of tower states,
	// it contains all existing towers.
	Towers []*Tower
}

func NewWorld() *World {
	s := &World{}
	return s
}

func (w *World) NearbyEnemies(position cp.Vector, radius float64) []*Enemy {
	return nil
}

func (w *World) NearbyTowers(position cp.Vector, radius float64) []*Tower {
	return nil
}

func (w *World) NearestEnemy(position cp.Vector) *Enemy {
	return nil
}

func (w *World) NearestTower(position cp.Vector) *Tower {
	if len(w.Towers) > 0 {
		// TODO: Make this actually do something
		return w.Towers[0]
	}

	return nil
}
