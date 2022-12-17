package projectile

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
	"math"
)

type Bullet struct {
	body  *cp.Body
	shape *cp.Shape
	size  float64
}

func NewBullet(x, y, heading, speed, size float64) *Bullet {
	pos := cp.Vector{
		X: x,
		Y: y,
	}

	body := cp.NewBody(0, 0)
	body.SetPosition(pos)
	body.SetAngle(heading)

	dx := math.Cos(heading) * speed
	dy := math.Sin(heading) * speed
	body.SetVelocity(dx, dy)

	shape := cp.NewCircle(body, size/2, pos)
	shape.SetMass(1.0)

	return &Bullet{
		body:  body,
		shape: shape,
		size:  size,
	}
}

func (b *Bullet) Update() error {
	return nil
}

func (b *Bullet) Draw(screen *ebiten.Image) {
	pos := b.body.Position()
	ebitenutil.DrawCircle(screen, pos.X, pos.Y, b.size/2, colornames.Aqua)
}

func (b *Bullet) AddToSpace(space *cp.Space) {
	space.AddBody(b.body)
	space.AddShape(b.shape)
}
