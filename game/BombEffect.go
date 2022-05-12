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

	currentPlayerTile := GetEntityTile(g, g.player1)

	if currentPlayerTile.pos.X == b.pos.X && currentPlayerTile.pos.Y == b.pos.Y {
		g.player1.Die()
	}

	tile := GetTileByPosition(&b.pos, g)
	if tile != nil && tile.kind == TileKindWood {
		tile.kind = TileKindEmpty
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
