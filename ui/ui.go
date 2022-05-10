package ui

import "github.com/hajimehoshi/ebiten/v2"

type Button struct {
	X, Y   int
	Action func()
	Text   string
	Image  *ebiten.Image
}

func (b *Button) Click() {
	b.Action()
}

func (b *Button) Draw(screen *ebiten.Image) {

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(b.X), float64(b.Y))
	screen.DrawImage(b.Image, op)
}
