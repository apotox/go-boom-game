package game

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Header struct {
	score         int
	level         int
	lives         int
	remainingTime int
	countDown     *time.Ticker
}

func NewHeader(score int, level int, lives int) *Header {

	return &Header{
		score:         score,
		level:         level,
		lives:         lives,
		remainingTime: 180,
		countDown:     time.NewTicker(1 * time.Second),
	}
}

func (h *Header) Update(g *Game) {

}

func (h *Header) Draw(image *ebiten.Image) {
	if image == nil {
		return
	}

}
