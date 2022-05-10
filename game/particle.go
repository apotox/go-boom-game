package game

import (
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type ParticleEnum int

const (
	ParticleDust ParticleEnum = iota
	ParticleSword
)

type Particle struct {
	pos             *Position
	sprite          *Sprite
	name            ParticleEnum
	playerDirection Dir
	isPlayerMoving  bool
}

func NewParticle(name ParticleEnum, pos *Position) *Particle {
	return &Particle{
		name:   name,
		pos:    pos,
		sprite: NewSprite(GetResource(ResourceNameDust), 3, 0, 12, nil, nil, false),
	}
}

func (p *Particle) Update(g *Game) error {
	p.playerDirection = g.player1.direction
	p.isPlayerMoving = g.player1.nextTile != nil
	p.sprite.Animate()
	return nil
}

func (p *Particle) Draw(boardImage *ebiten.Image) error {
	if boardImage == nil || !p.isPlayerMoving {
		return nil
	}

	op := &ebiten.DrawImageOptions{}

	randomInt := rand.Intn(tileSize)
	randomIntY := rand.Intn(tileSize / 4)
	switch p.playerDirection {
	case DirRight:
		op.GeoM.Translate(float64(p.pos.X-randomInt), float64(randomIntY+p.pos.Y+tileSize/2))
	case DirLeft:
		op.GeoM.Translate(float64(p.pos.X+tileSize-randomInt), float64(randomIntY+p.pos.Y+tileSize/2))

	default:
		return nil
	}

	boardImage.DrawImage(p.sprite.current, op)
	return nil
}
