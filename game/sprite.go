package game

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type DefaultImageCords struct {
	i int
	j int
}

type Offsets struct {
	offsetX int
	offsetY int
}

type ISprite interface {
	Animate() error
	SetCurrent(*ebiten.Image)
	GetCurrent() *ebiten.Image
	IsAnimation() bool
}

type Sprite struct {
	source    *ebiten.Image
	images    []*ebiten.Image
	line      int
	length    int
	current   *ebiten.Image
	counter   float32
	alpha     float32
	tileWidth int
	//single
	pos *Position
}

func (s *Sprite) Animate() error {

	if !s.IsAnimation() {
		return nil
	}
	tileIndex := int(s.counter) % s.length
	s.SetCurrent(s.images[tileIndex])

	if s.counter += s.alpha; s.counter > float32(s.length) {
		s.counter = 0
	}

	return nil
}

func (s *Sprite) IsAnimation() bool {
	return s.length > 1
}

func (s *Sprite) SetCurrent(i *ebiten.Image) {
	s.current = i
}

func (s *Sprite) GetCurrent() *ebiten.Image {
	return s.current
}

func (s *Sprite) Reset() error {

	s.counter = 0

	return nil
}

func NewAnimatedSprite(img *ebiten.Image, length int, line int, tileWidth int, offset *Offsets, scale bool) ISprite {

	if offset == nil {
		offset = &Offsets{0, 0}
	}
	s := &Sprite{
		alpha:     0.1,
		line:      line,
		source:    img,
		length:    length,
		tileWidth: img.Bounds().Dx() / length,
		images:    make([]*ebiten.Image, length),
	}

	for tileIndex := 0; tileIndex < length; tileIndex++ {
		i := s.source.SubImage(image.Rect(offset.offsetX+tileIndex*s.tileWidth, offset.offsetY+s.line*s.tileWidth, offset.offsetX+s.tileWidth*(1+tileIndex), offset.offsetY+(s.line+1)*s.tileWidth)).(*ebiten.Image)
		if scale {
			s.images[tileIndex] = ScaleImage(i)
		} else {
			s.images[tileIndex] = i
		}

	}

	s.SetCurrent(s.images[0])

	return s
}

func NewSingleSprite(img *ebiten.Image, pos *Position, tileWidth int, scale bool) ISprite {

	rect := image.Rect(
		pos.X*tileWidth,
		pos.Y*tileWidth,
		tileWidth*(1+pos.X),
		tileWidth*(1+pos.Y),
	)

	image := img.SubImage(rect).(*ebiten.Image)

	if scale {
		image = ScaleImage(image)
	}

	sprite := &Sprite{
		images: []*ebiten.Image{
			image,
		},
		length:    1,
		pos:       pos,
		tileWidth: tileWidth,
		source:    img,
		current:   image,
	}

	return sprite
}
