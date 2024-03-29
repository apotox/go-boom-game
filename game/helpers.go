package game

import (
	"math"
	"time"

	"github.com/apotox/go-boom-game/joystick"
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

func GetEntityTile(g *Game, entity Entity) (*Tile, *Position) {
	pos := entity.GetPosition()
	i := (pos.X + entity.GetSize()/2) / tileSize
	j := (pos.Y + entity.GetSize()/2) / tileSize
	return g.board.tiles[i+j*g.board.widthSize], &Position{X: i, Y: j}
}

func FilterTiles(arr []*Tile, f func(*Tile) bool) []*Tile {
	var r = []*Tile{}

	for _, v := range arr {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}

func Distance(a, b Position) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2))
}

func GetDxDyDir(a, b *Position) (dx int, dy int, dir joystick.Dir) {

	dx = a.X - b.X
	dy = a.Y - b.Y

	if joystick.Abs(dx) > 0 {
		if dx > 0 {
			dir = joystick.DirRight
		} else {
			dir = joystick.DirLeft
		}
	} else if joystick.Abs(dy) > 0 {

		if dy > 0 {
			dir = joystick.DirUp
		} else {
			dir = joystick.DirDown
		}
	}

	return

}

func SetTimeout(t int, f func()) {
	go func() {
		time.Sleep(time.Duration(t) * time.Millisecond)
		f()
	}()
}
