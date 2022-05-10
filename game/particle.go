package game

import "github.com/hajimehoshi/ebiten/v2"

type ParticleEnum int

const (
	ParticleDust ParticleEnum = iota
	ParticleSword
)

type Particle struct {
	pos    *Position
	sprite *Sprite
	name   ParticleEnum
}

func NewParticle(name ParticleEnum, pos *Position) *Particle {
	return &Particle{
		name:   name,
		pos:    pos,
		sprite: NewSprite(GetResource(ResourceNameDust), 3, 0, 12, nil, nil, false),
	}
}

func (p *Particle) Update(g *Game) error {
	p.sprite.Animate()
	return nil
}

func (p *Particle) Draw(boardImage *ebiten.Image) error {
	if boardImage == nil {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.pos.X), float64(p.pos.Y+tileSize/2))
	boardImage.DrawImage(p.sprite.current, op)
	return nil
}
