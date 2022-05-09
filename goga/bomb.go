package goga

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type BombState string

const (
	BombStateIdle      BombState = "idle"
	BombStateExploding BombState = "exploding"
)

type Bomb struct {
	index   int
	pos     *Position
	sprites map[BombState]*Sprite
	state   BombState
	timer   *time.Timer
	effects []*BombEffect
	radius  int
}

func NewBomb(index, x, y, lifeTime int) *Bomb {

	sprites := make(map[BombState]*Sprite)
	sprites[BombStateIdle] = NewSprite(GetResource(ResourceNameTiles), 1, 0, 32, &DefaultImageCords{
		i: 1,
		j: 9,
	}, nil)
	sprites[BombStateExploding] = NewSprite(GetResource(ResourceNameBomb), 8, 0, 32, nil, nil)
	return &Bomb{
		pos:     &Position{X: x, Y: y},
		sprites: sprites,
		timer:   time.NewTimer(time.Duration(lifeTime) * time.Second),
		index:   index,
		state:   BombStateIdle,
		effects: make([]*BombEffect, 0),
		radius:  3,
	}
}

func (b *Bomb) Update(g *Game) error {

	select {
	case <-b.timer.C:

		if b.state == BombStateIdle {
			b.state = BombStateExploding
			b.MakeBombEffects(g)
		} else {
			//g.RemoveBomb(b.index)
		}

	default:
	}

	b.sprites[b.state].Animate()

	for _, e := range b.effects {
		e.sprite.Animate()
	}

	return nil
}

func (b *Bomb) Draw(boardImage *ebiten.Image) error {
	if boardImage == nil {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.pos.X), float64(b.pos.Y))
	boardImage.DrawImage(b.sprites[b.state].current, op)

	for _, e := range b.effects {
		e.Draw(boardImage)
	}

	return nil
}

func GetVectorTiles(vet *Position, center Position, radius int, g *Game) []Position {

	postions := []Position{}

	for i := 0; i <= radius; i++ {
		tX := center.X + vet.X*i
		tY := center.Y + vet.Y*i
		tile := GetTileByPosition(&Position{X: tX, Y: tY}, g)
		if tile != nil && tile.kind == TileKindEmpty {
			postions = append(postions, Position{X: tX, Y: tY})
		} else {
			break
		}
	}

	return postions
}

func (b *Bomb) MakeBombEffects(g *Game) bool {
	if b.state != BombStateExploding {
		return false
	}

	center := GetTilePos(b.pos)
	// 4 vectors

	v1 := GetVectorTiles(&Position{X: -1, Y: 0}, center, b.radius, g)
	v2 := GetVectorTiles(&Position{X: 0, Y: 1}, center, b.radius, g)
	v3 := GetVectorTiles(&Position{X: 1, Y: 0}, center, b.radius, g)
	v4 := GetVectorTiles(&Position{X: 0, Y: -1}, center, b.radius, g)

	vectors := []Position{}
	vectors = append(vectors, v1...)
	vectors = append(vectors, v2...)
	vectors = append(vectors, v3...)
	vectors = append(vectors, v4...)

	for _, t := range vectors {

		b.effects = append(b.effects, &BombEffect{
			pos:    &Position{X: t.X * tileSize, Y: t.Y * tileSize},
			sprite: NewSprite(GetResource(ResourceNameBomb), 8, 0, 32, nil, nil),
			timer:  time.NewTimer(time.Duration(1) * time.Second),
		})
	}

	return false

}
