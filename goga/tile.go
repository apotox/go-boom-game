package goga

import "github.com/hajimehoshi/ebiten/v2"

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

func GetSpriteByKind(kind int) *Sprite {
	switch kind {
	case 0:
		return NewSprite(GetResource("tiles"), 1, 0, 32, &DefaultImageCords{
			i: 2,
			j: 6,
		}, nil)
	case 1:
		return NewSprite(GetResource("tiles"), 1, 0, 32, &DefaultImageCords{
			i: 4,
			j: 9,
		}, nil)
	case 2:
		return NewSprite(GetResource("tiles"), 1, 0, 32, &DefaultImageCords{0, 0}, nil)
	default:
		return NewSprite(GetResource("tiles"), 1, 0, 32, &DefaultImageCords{0, 0}, nil)
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
