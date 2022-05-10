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

type Sprite struct {
	source    *ebiten.Image
	images    []*ebiten.Image
	line      int
	length    int
	current   *ebiten.Image
	counter   float32
	alpha     float32
	tileWidth int
	scale     bool
}

func (s *Sprite) Animate() error {

	if s.length == 1 {
		return nil
	}
	tileIndex := int(s.counter) % s.length
	s.SetCurrent(s.images[tileIndex])

	if s.counter += s.alpha; s.counter > float32(s.length) {
		s.counter = 0
	}

	return nil
}

func (s *Sprite) SetCurrent(i *ebiten.Image) {

	s.current = i
}

func (s *Sprite) Reset() error {

	s.counter = 0

	return nil
}

func NewSprite(img *ebiten.Image, length int, line int, tileWidth int, defaultImageCord *DefaultImageCords, offset *Offsets, scale bool) *Sprite {

	if offset == nil {
		offset = &Offsets{0, 0}
	}
	s := &Sprite{
		alpha:     0.1,
		line:      line,
		source:    img,
		length:    length,
		tileWidth: tileWidth,
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

	if defaultImageCord != nil {
		i := img.SubImage(image.Rect(offset.offsetX+defaultImageCord.i*s.tileWidth, offset.offsetY+defaultImageCord.j*s.tileWidth, offset.offsetX+defaultImageCord.i*s.tileWidth+s.tileWidth, offset.offsetY+(defaultImageCord.j+1)*s.tileWidth)).(*ebiten.Image)
		s.SetCurrent(ScaleImage(i))

		if scale {
			s.SetCurrent(ScaleImage(i))
		} else {
			s.SetCurrent(i)
		}
	} else {
		s.SetCurrent(s.images[0])
	}

	return s
}
