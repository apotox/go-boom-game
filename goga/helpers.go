package goga

import (
	"fmt"

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

func GetTiles(tilePos Position, g *Game) (center, up, left, down, right *Tile, err error) {

	center = g.board.tiles[tilePos.Y*g.board.widthSize+tilePos.X]

	if tilePos.Y > 0 {
		up = g.board.tiles[(tilePos.Y-1)*g.board.widthSize+tilePos.X]

	}
	if tilePos.X > 0 {
		left = g.board.tiles[tilePos.Y*g.board.widthSize+tilePos.X-1]
	}
	if tilePos.Y < g.board.heightSize-1 {
		down = g.board.tiles[(tilePos.Y+1)*g.board.widthSize+tilePos.X]
	}
	if tilePos.X < g.board.widthSize-1 {
		right = g.board.tiles[tilePos.Y*g.board.widthSize+tilePos.X+1]
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %+v\n", r)
			err = fmt.Errorf("%+v", r)
		}
	}()

	return
}
