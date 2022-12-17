package runtime

import "github.com/glesica/towergame/internal/tower"

// RunTower executes a program against a particular tower.
func RunTower(t tower.T) {
	h := t.Heading()
	t.Aim(h + 0.01)
}
