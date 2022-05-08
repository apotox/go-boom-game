package goga

import "github.com/hajimehoshi/ebiten/v2"

type Bomb struct {
	pos           *Position
	sprite        *Sprite
	explodeSprite *Sprite
	lifeTimeSec   int
}

func NewBomb(x, y int) *Bomb {
	return &Bomb{
		pos: &Position{X: x, Y: y},
		sprite: NewSprite(GetResource("tiles"), 1, 0, 32, &DefaultImageCords{
			i: 1,
			j: 9,
		}),
	}
}

func (b *Bomb) Update(g *Game) error {

	b.sprite.Animate()

	return nil
}

func (b *Bomb) Draw(boardImage *ebiten.Image) error {
	if boardImage == nil {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.pos.X), float64(b.pos.Y))
	// p.movingSprite.current.Fill(playerColor)
	boardImage.DrawImage(b.sprite.current, op)
	return nil
}
