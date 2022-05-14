package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	pos      *Position
	boardPos *Position
	sprite   ISprite
	tasks    []Task
	kind     int
	oldKind  int
}

const (
	tileSize      = 14
	TileKindEmpty = 0
	TileKindWall  = 1
	TileKindWood  = 2
)

func GetTileSize() int {
	return tileSize
}

func GetSpriteByKind(kind int) ISprite {
	switch kind {
	case 0:
		return NewSingleSprite(GetResource(ResourceNameFloor), &Position{
			X: 0,
			Y: 0,
		}, 32, true)
	case 1:
		return NewSingleSprite(GetResource(ResourceNameWall), &Position{
			X: 0,
			Y: 0,
		}, 32, true)
	case 2:
		return NewSingleSprite(GetResource(ResourceNameDoor), &Position{
			X: 0,
			Y: 0,
		}, 32, true)
	default:
		return NewSingleSprite(GetResource(ResourceNameWall), &Position{
			X: 0,
			Y: 0,
		}, 32, true)
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
	if t.kind != t.oldKind {
		t.sprite = GetSpriteByKind(t.kind)
		t.oldKind = t.kind
	}

	t.sprite.Animate()
	return nil
}

func (t *Tile) Walkable() bool {

	return t.kind == 0
}

func (t *Tile) Draw(boardImage *ebiten.Image) error {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(t.pos.X), float64(t.pos.Y))
	if t.sprite.GetCurrent() != nil {
		boardImage.DrawImage(t.sprite.GetCurrent(), op)
	}

	return nil
}

func GetTileByBoardPosition(pos *Position, g *Game) *Tile {
	tIndex := pos.X + (pos.Y)*g.board.widthSize
	if tIndex < 0 || tIndex >= len(g.board.tiles) {
		return nil
	}
	return g.board.tiles[tIndex]
}

func GetTileBoardPos(pos *Position) Position {

	return Position{X: (pos.X + tileSize/2) / tileSize, Y: (pos.Y + tileSize/2) / tileSize}
}

func GetSurroundedTiles(tilePos Position, g *Game) (center, up, left, down, right *Tile, tiles []*Tile, err error) {
	tiles = make([]*Tile, 0)
	center = g.board.tiles[tilePos.Y*g.board.widthSize+tilePos.X]

	if tilePos.Y > 0 {
		up = g.board.tiles[(tilePos.Y-1)*g.board.widthSize+tilePos.X]
		tiles = append(tiles, up)
	}
	if tilePos.X > 0 {
		left = g.board.tiles[tilePos.Y*g.board.widthSize+tilePos.X-1]
		tiles = append(tiles, left)
	}
	if tilePos.Y < g.board.heightSize-1 {
		down = g.board.tiles[(tilePos.Y+1)*g.board.widthSize+tilePos.X]
		tiles = append(tiles, down)
	}
	if tilePos.X < g.board.widthSize-1 {
		right = g.board.tiles[tilePos.Y*g.board.widthSize+tilePos.X+1]
		tiles = append(tiles, right)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %+v\n", r)
			err = fmt.Errorf("%+v", r)
		}
	}()

	return
}
