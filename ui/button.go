package ui

import (
	"image"
	"image/color"

	"github.com/apotox/go-boom-game/joystick"
	mycolors "github.com/apotox/go-boom-game/mycolors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

var (
	emptyImage = ebiten.NewImage(3, 3)

	// emptySubImage is an internal sub image of emptyImage.
	// Use emptySubImage at DrawTriangles instead of emptyImage in order to avoid bleeding edges.
	emptySubImage = emptyImage.SubImage(image.Rect(1, 1, 2, 2)).(*ebiten.Image)
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

	font      font.Face
	flooding  int
	_flooding float32
}

var mplusNormalFont font.Face

func NewButton(x, y, width int, text string, font font.Face, image *ebiten.Image, action func()) *Button {
	emptyImage.Fill(color.White)
	b := &Button{
		X:        x - width/2,
		Y:        y - width/2,
		Width:    width,
		Text:     text,
		Image:    image,
		font:     font,
		flooding: 0,
	}

	b.Action = func() {
		action()
	}

	return b
}

func (b *Button) Update(input *joystick.Input) {
	// up and down animation
	input.Update()
	if b._flooding += 0.1; b._flooding > 1 {
		b.flooding = (b.flooding + 1) % 2
		b._flooding = 0
	}

	if ok := input.Tap(); ok {
		b.Action()
	}
}

func (b *Button) Draw(screen *ebiten.Image) {

	if b.Image == nil {
		return
	}

	ebitenutil.DrawRect(screen, float64(b.X), float64(b.Y+b.flooding), float64(b.Width), float64(22), color.White)
	ebitenutil.DrawRect(screen, float64(b.X)+2, float64(b.Y+b.flooding+1), float64(b.Width), float64(22), mycolors.PrimaryColor)
	text.Draw(screen, b.Text, b.font, b.X+11, b.Y+b.flooding+16, color.White)

}
