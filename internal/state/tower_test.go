package state

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestTower_Rotate(t *testing.T) {
	for _, c := range []struct {
		start, delta, end float64
	}{
		{0, math.Pi, math.Pi},
		{0, -math.Pi, math.Pi},
		{0.5 * math.Pi, math.Pi, 1.5 * math.Pi},
		{0.5 * math.Pi, -math.Pi, 1.5 * math.Pi},
	} {
		s := NewTower()
		s.AimAngle = c.start
		s.Rotate(c.delta)
		assert.InDelta(t, c.end, s.AimAngle, 0.01, c)
	}

	for _, c := range []struct {
		angle, x, y float64
	}{
		{0, 1, 0},
		{0.5 * math.Pi, 0, 1},
		{math.Pi, -1, 0},
		{1.5 * math.Pi, 0, -1},
	} {
		s := NewTower()
		s.AimAngle = c.angle
		s.Rotate(0)
		assert.InDelta(t, c.x, s.AimVector.X, 0.01, c)
		assert.InDelta(t, c.y, s.AimVector.Y, 0.01, c)
	}
}
