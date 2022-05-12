package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type ParticleEnum int

const (
	ParticleDust ParticleEnum = iota
	ParticleSword
)

type Particle struct {
	pos             *Position
	sprite          ISprite
	name            ParticleEnum
	playerDirection Dir
	isPlayerMoving  bool
}

func NewParticle(name ParticleEnum, pos *Position) *Particle {
	return &Particle{
		name:   name,
		pos:    pos,
		sprite: NewAnimatedSprite(GetResource(ResourceNameDust), 3, 0, 12, nil, false),
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

	op.GeoM.Translate(float64(p.pos.X), float64(p.pos.Y-tileSize/2))

	boardImage.DrawImage(p.sprite.GetCurrent(), op)
	return nil
}
