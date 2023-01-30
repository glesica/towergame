package state

type World struct {
	// TODO: Replace the slices with R-trees or something else fast

	// Bullets holds a reference to the container of bullets or other
	// projectiles flying around the game map.
	Bullets *Index[*Bullet]

	// BulletQueue represents bullets that have been fired during the
	// current time step and have, therefore, not been added to the
	// index yet to avoid inconsistency.
	BulletQueue []*Bullet

	// Enemies holds a reference to the current slice of enemy states,
	// it contains all existing enemies.
	Enemies *Index[*Enemy]

	// Towers holds a reference to the current slice of tower states,
	// it contains all existing towers.
	Towers *Index[*Tower]
}

func NewWorld() *World {
	s := &World{
		Bullets: NewIndex[*Bullet](),
		Enemies: NewIndex[*Enemy](),
		Towers:  NewIndex[*Tower](),
	}
	return s
}

func (w *World) Resolve() {
	// TODO: Remove dead enemies
	// TODO: Remove expired bullets
	for _, b := range w.BulletQueue {
		w.Bullets.Add(b)
	}
	w.BulletQueue = nil
}
