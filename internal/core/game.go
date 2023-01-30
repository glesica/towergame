package core

import (
	"github.com/glesica/towergame/internal/bullet"
	"github.com/glesica/towergame/internal/enemy"
	"github.com/glesica/towergame/internal/state"
	"github.com/glesica/towergame/internal/tower"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/colornames"
)

type Game struct {
	world *state.World
}

func NewGame(w *state.World) *Game {
	game := &Game{
		world: w,
	}
	return game
}

func (g *Game) Draw(screen *ebiten.Image) {
	img := ebiten.NewImageFromImage(screen)

	g.world.Towers.Apply(func(t *state.Tower) {
		tower.Basic.Draw(t, img)
	})

	g.world.Enemies.Apply(func(e *state.Enemy) {
		enemy.Basic.Draw(e, img)
	})

	g.world.Bullets.Apply(func(b *state.Bullet) {
		bullet.Basic.Draw(b, img)
	})

	screen.Fill(colornames.Floralwhite)

	ops := &ebiten.DrawImageOptions{}

	// Mirror over the Y-axis to make directions work properly
	ops.GeoM.Scale(1, -1)
	ops.GeoM.Translate(0, 768)

	screen.DrawImage(img, ops)
}

func (g *Game) Update() error {
	dt := 1.0 / float64(ebiten.TPS())

	g.updateTowers(dt)
	g.updateBullets(dt)
	g.updateEnemies(dt)

	g.world.Resolve()

	return nil
}

func (g *Game) Layout(width, height int) (int, int) {
	return 1024, 768
}

func (g *Game) updateBullets(dt float64) {
	g.world.Bullets.Apply(func(b *state.Bullet) {
		bullet.Basic.Update(b, g.world, dt)
	})
}

func (g *Game) updateEnemies(dt float64) {
	g.world.Enemies.Apply(func(e *state.Enemy) {
		enemy.Basic.Update(e, g.world, dt)
	})
}

func (g *Game) updateTowers(dt float64) {
	g.world.Towers.Apply(func(t *state.Tower) {
		tower.Basic.Update(t, g.world, &state.Instruction{
			Aim:  1,
			Fire: true,
		}, dt)
	})
}
