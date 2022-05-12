package game

import "github.com/hajimehoshi/ebiten/v2"

type PickableEnum int

const (
	PickableEnumPower PickableEnum = iota
	PickableEnumSpeed
	PickableEnumLife
	PickableEnumMaxItems
)

type Pickable struct {
	kind   PickableEnum
	pos    *Position
	sprite *Sprite
}

func NewPickable(kind PickableEnum, pos *Position) *Pickable {
	return &Pickable{
		kind:   kind,
		pos:    pos,
		sprite: NewSprite(GetResource(ResourceNameRunner), 8, 1, 32, nil, nil, true),
	}
}

func (p *Pickable) GetKind() PickableEnum {
	return p.kind
}

func (p *Pickable) Draw(boardImage *ebiten.Image) error {
	if boardImage == nil {
		return nil
	}
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(float64(p.pos.X), float64(p.pos.Y))

	boardImage.DrawImage(p.sprite.current, op)
	return nil
}

func (p *Pickable) GetPosition() *Position {
	return p.pos
}
func (p *Pickable) GetName() string {
	return "pickable"
}
func (p *Pickable) GetSize() int {
	return tileSize
}

func (p *Pickable) Update(g *Game) {

	p.sprite.Animate()
}
