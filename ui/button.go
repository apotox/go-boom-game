package ui

import (
	"log"

	"github.com/apotox/goga/joystick"
	fonts "github.com/apotox/goga/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type Component interface {
	Draw(*ebiten.Image)
	Update(*joystick.Input)
}

type Button struct {
	X, Y, Width  int
	Action       func()
	Text         string
	Image        *ebiten.Image
	imagePressed *ebiten.Image
	clicked      bool
}

func (b *Button) Click() {
	b.Action()
}

func (b *Button) Update(input *joystick.Input) {
	// up and down animation

}

var mplusNormalFont font.Face

func NewButton(x, y, width int, action func(), text string, image *ebiten.Image) *Button {

	tt, err := opentype.Parse(fonts.MPlus1pRegular_ttf)

	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont, err = opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	b := &Button{
		X:     x,
		Y:     y,
		Width: width,
		Text:  text,
		Image: image,
	}

	b.Action = func() {
		action()
	}

	return b
}

func (b *Button) Draw(screen *ebiten.Image) {

	if b.Image == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.X), float64(b.Y))

	op.GeoM.Translate(float64(screen.Bounds().Size().X/4)-float64(b.Image.Bounds().Size().X/2), float64(screen.Bounds().Size().Y/4)-float64(b.Image.Bounds().Size().Y/2))
	screen.DrawImage(b.Image, op)
}
