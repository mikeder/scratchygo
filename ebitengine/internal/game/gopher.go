package game

import (
	"bytes"
	"image"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/unitoftime/ecs"
)

const (
	gopherSpeed float64 = 8.0
)

var (
	gopherImage *ebiten.Image
)

func init() {
	gopherPng, _, err := image.Decode(bytes.NewReader(Gopher_png))
	if err != nil {
		log.Fatal(err)
	}
	gopherImage = ebiten.NewImageFromImage(gopherPng)
}

type Gopher struct {
	id       ecs.Id
	image    *ebiten.Image
	pos      Vec2
	velocity Vec2
}

func NewGopher(id ecs.Id, pos Vec2) Gopher {
	return Gopher{
		id:       id,
		image:    gopherImage,
		pos:      pos,
		velocity: Vec2{0, 0},
	}
}

func MoveGopher(input *input, world *ecs.World) {
	q := ecs.Query1[Gopher](world)

	q.MapId(func(id ecs.Id, a *Gopher) {
		if input.up {
			a.pos.Y -= 1 * gopherSpeed
		}
		if input.down {
			a.pos.Y += 1 * gopherSpeed
		}
		if input.left {
			a.pos.X -= 1 * gopherSpeed
		}
		if input.right {
			a.pos.X += 1 * gopherSpeed
		}
	})
}

func DrawGopher(screen *ebiten.Image, op *ebiten.DrawImageOptions, world *ecs.World) {
	q := ecs.Query1[Gopher](world)

	q.MapId(func(id ecs.Id, g *Gopher) {
		op.GeoM.Reset()
		op.GeoM.Scale(0.5, 0.5)
		op.GeoM.Translate(float64(g.pos.X), float64(g.pos.Y))
		screen.DrawImage(g.image, op)
	})
}