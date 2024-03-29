package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type BombState string

const (
	BombStateIdle      BombState = "idle"
	BombStateExploding BombState = "exploding"
)

var explodeDirections = []Position{
	{X: -1, Y: 0},
	{X: 0, Y: 1},
	{X: 1, Y: 0},
	{X: 0, Y: -1},
}

type Bomb struct {
	pos     *Position
	sprites map[BombState]ISprite
	state   BombState
	timer   *time.Timer
	effects []*BombEffect
	radius  int
}

func NewBomb(g *Game, x, y, lifeTime int) *Bomb {

	sprites := make(map[BombState]ISprite)
	sprites[BombStateIdle] = NewAnimatedSprite(GetResource(ResourceNameBoomIdle), 4, 0, 52, nil, true)
	sprites[BombStateExploding] = NewAnimatedSprite(GetResource(ResourceNameBoomOn), 6, 0, 52, nil, true)

	b := &Bomb{
		pos:     &Position{X: x, Y: y},
		sprites: sprites,
		timer:   time.NewTimer(time.Duration(lifeTime) * time.Second),
		state:   BombStateIdle,
		effects: make([]*BombEffect, 0),
		radius:  3,
	}
	h := GetTileBoardPos(b.pos)
	p := GetTileByBoardPosition(&h, g)
	p.reserved = true

	return b
}

func (b *Bomb) Update(g *Game) error {
	select {
	case <-b.timer.C:

		if b.state == BombStateIdle {
			b.state = BombStateExploding
			b.timer = time.NewTimer(time.Duration(1) * time.Second)
			b.MakeBombEffects(g)
		} else {
			g.RemoveBomb(b)
		}
	default:
	}

	b.sprites[b.state].Animate()

	for _, e := range b.effects {
		e.Update(g)
	}

	return nil
}

func (b *Bomb) Draw(boardImage *ebiten.Image) error {
	if boardImage == nil {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.pos.X), float64(b.pos.Y))

	boardImage.DrawImage(b.sprites[b.state].GetCurrent(), op)

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
		tile := GetTileByBoardPosition(&Position{X: tX, Y: tY}, g)
		if tile != nil {

			if tile.kind == TileKindEmpty {
				postions = append(postions, Position{X: tX, Y: tY})
			} else if tile.kind == TileKindWood {
				postions = append(postions, Position{X: tX, Y: tY})
				break
			} else if tile.kind == TileKindWall {
				break
			}

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

	center := GetTileBoardPos(b.pos)

	for _, ed := range explodeDirections {
		vt := GetVectorTiles(&ed, center, b.radius, g)

		for _, t := range vt {

			b.effects = append(b.effects, NewBombEffect(t, g.player1.pos))
		}
	}

	return false

}
