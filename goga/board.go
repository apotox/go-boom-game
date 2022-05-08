package goga

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Board struct {
	widthSize  int
	heightSize int
	tiles      []*Tile
}

func (b *Board) Size() (int, int) {
	return b.widthSize, b.heightSize
}

func NewBoard(widthSize int, heightSize int, tiles []*Tile) (*Board, error) {
	b := &Board{
		widthSize:  widthSize,
		heightSize: heightSize,
		tiles:      tiles,
	}
	return b, nil
}

func (b *Board) Update(g *Game) error {

	for y := 0; y < b.heightSize; y++ {
		for x := 0; x < b.widthSize; x++ {
			tile := b.tiles[x+y*b.widthSize]
			tile.Update(g)
		}
	}
	return nil

}

func (b *Board) Draw(boardImage *ebiten.Image) error {
	for y := 0; y < b.heightSize; y++ {
		for x := 0; x < b.widthSize; x++ {
			tile := b.tiles[x+y*b.widthSize]
			tile.Draw(boardImage)
		}
	}
	return nil

}
