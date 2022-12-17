package tower

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/jakecoffman/cp"
	"golang.org/x/image/colornames"
	"math"
)

type Basic struct {
	x, y                 float64
	heading, goalHeading float64
	aimSpeed             float64
	radius               float64
	body                 *cp.Body
	shapes               []*cp.Shape
}

func NewBasic(x, y, heading float64) *Basic {
	pos := cp.Vector{
		X: x,
		Y: y,
	}

	body := cp.NewBody(0, 0)
	body.SetPosition(pos)
	body.SetAngle(heading)

	radius := 25.0
	shape := cp.NewCircle(body, radius, pos)

	return &Basic{
		x:           x,
		y:           y,
		heading:     heading,
		goalHeading: heading,
		aimSpeed:    1.0,
		radius:      radius,
		body:        body,
		shapes: []*cp.Shape{
			shape,
		},
	}
}

func (b *Basic) Aim(heading float64) float64 {
	b.goalHeading = heading

	return math.Abs(b.heading-heading) / b.aimSpeed
}

func (b *Basic) Heading() float64 {
	return b.heading
}

func (b *Basic) Fire() float64 {
	return 0
}

func (b *Basic) Update(dt float64) error {
	//hDelta := (b.heading - b.goalHeading) / b.aimSpeed * dt

	return nil
}

func (b *Basic) Draw(screen *ebiten.Image) {
	// Turret
	ebitenutil.DrawCircle(screen, b.x, b.y, b.radius, colornames.Darkmagenta)

	// Gun
	gr := b.radius / 4
	gx := math.Cos(b.heading)*(b.radius-gr) + b.x
	gy := math.Sin(b.heading)*(b.radius-gr) + b.y
	ebitenutil.DrawCircle(screen, gx, gy, gr, colornames.Fuchsia)
}

func (b *Basic) AddToSpace(space *cp.Space) {
	space.AddBody(b.body)
	for _, s := range b.shapes {
		space.AddShape(s)
	}
}

func (b *Basic) RemoveFromSpace(space *cp.Space) {
	space.RemoveBody(b.body)
	for _, s := range b.shapes {
		space.RemoveShape(s)
	}
}
