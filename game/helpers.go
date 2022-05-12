package game

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

func ScaleImage(img *ebiten.Image) *ebiten.Image {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(float64(tileSize)/float64(img.Bounds().Size().X), float64(tileSize)/float64(img.Bounds().Size().Y))

	u := ebiten.NewImage(tileSize, tileSize)
	u.DrawImage(img, op)

	return u
}

func Filter(arr []Task, f func(Task) bool) []interface{} {
	var r []interface{}
	for _, v := range arr {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func DegreesToRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180)
}

type Entity interface {
	GetPosition() *Position
	GetName() string
	GetSize() int
}

func GetEntityTile(g *Game, entity Entity) *Tile {

	pos := entity.GetPosition()

	i := (pos.X + entity.GetSize()/2) / tileSize
	j := (pos.Y + entity.GetSize()/2) / tileSize

	return g.board.tiles[i+j*g.board.widthSize]
}
