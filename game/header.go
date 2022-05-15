package game

import (
	"fmt"
	"image/color"
	"time"

	"github.com/apotox/goga/mycolors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

var (
	HeaderHeight = 20
)

type Header struct {
	score         int
	level         int
	lives         int
	remainingTime int
	countDown     *time.Ticker
}

func NewHeader() *Header {

	return &Header{
		score:         0,
		level:         0,
		lives:         0,
		remainingTime: 110,
		countDown:     time.NewTicker(1 * time.Second),
	}
}

func (h *Header) Update(g *Game) {

	select {
	case <-h.countDown.C:
		h.remainingTime--
		if h.remainingTime == 0 {
			g.SetScreen(GameScreenGameOver)
		}
	default:
		// nothing
	}

	h.level = g.level
	h.lives = g.player1.GetFeatures().life - g.player1.dieds
}

func (h *Header) Draw(image *ebiten.Image) {
	if image == nil {
		return
	}
	image.Fill(mycolors.PrimaryColor)
	minutes := (h.remainingTime / 60)
	zero := ""

	seconds := (h.remainingTime % 60)
	if seconds < 10 {
		zero = "0"
	}
	text.Draw(image, fmt.Sprintf("LIFE: %d  LEVEL: %d TIME: %d:%s%d", h.lives, h.level, minutes, zero, seconds), GetFont("default-12"), tileSize, tileSize, color.White)

}
