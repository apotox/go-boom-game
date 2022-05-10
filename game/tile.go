package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	pos      *Position
	boardPos *Position
	sprite   *Sprite
	tasks    []Task
	kind     int
}

const (
	tileSize      = 14
	TileKindEmpty = 0
)

func GetTileSize() int {
	return tileSize
}

func GetSpriteByKind(kind int) *Sprite {
	switch kind {
	case 0:
		return NewSprite(GetResource(ResourceNameTiles), 1, 0, 32, &DefaultImageCords{
			i: 2,
			j: 6,
		}, nil, true)
	case 1:
		return NewSprite(GetResource(ResourceNameTiles), 1, 0, 32, &DefaultImageCords{
			i: 4,
			j: 9,
		}, nil, true)
	case 2:
		return NewSprite(GetResource(ResourceNameTiles), 1, 0, 32, &DefaultImageCords{0, 0}, nil, true)
	default:
		return NewSprite(GetResource(ResourceNameTiles), 1, 0, 32, &DefaultImageCords{0, 0}, nil, true)
	}
}

func NewTile(x, y int, kind int) *Tile {

	return &Tile{
		pos:    &Position{X: x, Y: y},
		sprite: GetSpriteByKind(kind),
		kind:   kind,
		tasks:  []Task{},
	}
}

func (t *Tile) Update(g *Game) error {
	t.sprite.Animate()
	return nil
}

func (t *Tile) Walkable() bool {

	return t.kind == 0
}

func (t *Tile) Draw(boardImage *ebiten.Image) error {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.pos.X), float64(t.pos.Y))
	if t.sprite.current != nil {
		boardImage.DrawImage(t.sprite.current, op)
	}

	return nil
}

func GetTileByPosition(pos *Position, g *Game) *Tile {
	tIndex := pos.X + (pos.Y)*g.board.widthSize
	if tIndex < 0 || tIndex >= len(g.board.tiles) {
		return nil
	}
	return g.board.tiles[tIndex]
}

func GetTilePos(pos *Position) Position {

	return Position{X: (pos.X) / tileSize, Y: (pos.Y) / tileSize}
}

func GetSurroundedTiles(tilePos Position, g *Game) (center, up, left, down, right *Tile, err error) {

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
