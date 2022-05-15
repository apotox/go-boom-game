package game

import (
	"image/color"
	"strconv"
	"time"

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
		remainingTime: 180,
		countDown:     time.NewTicker(1 * time.Second),
	}
}

func (h *Header) Update(g *Game) {

	h.level = g.level
	h.lives = g.player1.GetFeatures().life - g.player1.dieds
}

func (h *Header) Draw(image *ebiten.Image) {
	if image == nil {
		return
	}

	text.Draw(image, "Life: "+strconv.Itoa(h.lives), GetFont("default-12"), tileSize, tileSize, color.White)

}
