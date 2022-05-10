package game

import (
	"fmt"

	ui "github.com/apotox/goga/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameScreen int

const (
	GameScreenPlay GameScreen = iota
	GameScreenGameOver
	GameScreenPause
	GameScreenStart
)

var screens = map[GameScreen]func(*Game, *ebiten.Image) error{
	GameScreenStart: func(g *Game, screen *ebiten.Image) error {

		if g.boardImage == nil {
			g.boardImage = ebiten.NewImage(g.board.widthSize*tileSize, g.board.heightSize*tileSize)
		}
		g.boardImage.Fill(backgroundColor)

		if g.board != nil {
			g.board.Draw(g.boardImage)
		}

		if g.player1 != nil {

			g.player1.Draw(g.boardImage)
		}

		if len(g.bombs) > 0 {
			for _, bomb := range g.bombs {
				bomb.Draw(g.boardImage)
			}
		}

		if len(g.pickables) > 0 {
			for _, pickable := range g.pickables {
				pickable.Draw(g.boardImage)
			}
		}

		select {
		case <-g.pickableTicker.C:
			g.AddPickable(PickableEnumPower)
		default:
			// nothing
		}

		screen.DrawImage(g.boardImage, &ebiten.DrawImageOptions{})

		return nil
	},
	GameScreenGameOver: func(g *Game, screen *ebiten.Image) error {
		g.boardImage.Fill(backgroundColor)

		button := &ui.Button{
			X:      g.board.widthSize*tileSize/2 - tileSize/2,
			Y:      g.board.heightSize*tileSize/2 - tileSize/2,
			Action: func() { g.gameScreen = GameScreenStart },
		}

		fmt.Printf("Game over!\n %v", button)

		return nil
	},
}
