package state

import (
	"github.com/jakecoffman/cp"
)

// The basic interface here will be replaced with an r-tree or similar
// data structure with efficient spatial lookups. This implementation
// is just meant to facilitate discovery on the rest of the game.

type Entity interface {
	ID() uint64
	Pos() cp.Vector
}

type Index[E Entity] struct {
	elements map[uint64]E
}

func NewIndex[E Entity]() *Index[E] {
	return &Index[E]{
		elements: make(map[uint64]E),
	}
}

func (c *Index[E]) Size() int {
	return len(c.elements)
}

func (c *Index[E]) Add(e E) {
	c.elements[e.ID()] = e
}

func (c *Index[E]) Apply(f func(E)) {
	for _, e := range c.elements {
		f(e)
	}
}

func (c *Index[E]) Near(p cp.Vector, d float64) []E {
	var hits []E
	for _, e := range c.elements {
		if e.Pos().Near(p, d) {
			hits = append(hits, e)
		}
	}

	return hits
}

func (c *Index[E]) Nearest(p cp.Vector) E {
	// TODO: Make this handle empty index properly
	var hit uint64
	var d float64
	for i, e := range c.elements {
		if hit == 0 {
			hit = i
			d = e.Pos().DistanceSq(p)
		}

		ed := e.Pos().Distance(p)
		if ed < d {
			hit = i
			d = ed
		}
	}

	return c.elements[hit]
}

func (c *Index[E]) Remove(t E) {
	delete(c.elements, t.ID())
}
