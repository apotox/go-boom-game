package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type BombEffect struct {
	pos            Position
	sprite         *Sprite
	playerPosition *Position
}

func (b *BombEffect) Update(g *Game) error {

	if b.sprite == nil || b.playerPosition == nil {
		return nil
	}

	b.sprite.Animate()

	currentPlayerX := b.playerPosition.X / tileSize
	currentPlayerY := b.playerPosition.Y / tileSize

	if currentPlayerX == b.pos.X && currentPlayerY == b.pos.Y {
		g.player1.Die()
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("BombEffect.Update:", r)
		}
	}()

	return nil
}

func (b *BombEffect) Draw(boardImage *ebiten.Image) error {
	if boardImage == nil || b.sprite == nil {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.pos.X*tileSize), float64(b.pos.Y*tileSize))
	boardImage.DrawImage(b.sprite.current, op)

	return nil
}

func NewBombEffect(pos Position, playerPosition *Position) *BombEffect {
	return &BombEffect{
		pos:            pos,
		playerPosition: playerPosition,
		sprite:         NewSprite(GetResource(ResourceNameBomb), 8, 0, 32, nil, nil, true),
	}
}
