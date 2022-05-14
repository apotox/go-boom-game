package game

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type BombEffect struct {
	pos    Position
	sprite ISprite
}

func (b *BombEffect) Update(g *Game) error {

	if b.sprite == nil {
		return nil
	}

	b.sprite.Animate()
	_, playerTilePosition := GetEntityTile(g, g.player1)

	if playerTilePosition.X == b.pos.X && playerTilePosition.Y == b.pos.Y {
		g.player1.Die(g)
	}

	tile := GetTileByBoardPosition(&b.pos, g)
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
	boardImage.DrawImage(b.sprite.GetCurrent(), op)

	return nil
}

func NewBombEffect(pos Position, playerPosition *Position) *BombEffect {
	return &BombEffect{
		pos:    pos,
		sprite: NewAnimatedSprite(GetResource(ResourceNameBoomOn), 6, 0, 52, nil, true),
	}
}
