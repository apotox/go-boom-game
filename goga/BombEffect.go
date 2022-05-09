package goga

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type BombEffect struct {
	pos    *Position
	sprite *Sprite
	timer  *time.Timer
}

func (b *BombEffect) Draw(boardImage *ebiten.Image) error {
	if boardImage == nil {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.pos.X), float64(b.pos.Y))
	boardImage.DrawImage(b.sprite.current, op)

	return nil
}
