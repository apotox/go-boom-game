package game

import (
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

type ScreenOptions struct {
	Draw   func(*Game, *ebiten.Image) error
	Update func(*Game)
	Init   func(*Game)
}

var screens = map[GameScreen]ScreenOptions{
	GameScreenPlay: {
		Draw: func(g *Game, screen *ebiten.Image) error {

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
		Update: func(g *Game) {
			g.input.Update()

			if g.board != nil {
				g.board.Update(g)
			}

			if g.player1 != nil {

				g.player1.Update(g)
			}

			if len(g.bombs) > 0 {
				for _, bomb := range g.bombs {
					bomb.Update(g)
				}
			}

			if len(g.pickables) > 0 {
				for _, pickable := range g.pickables {
					pickable.Update(g)
				}
			}
		},
	},
	GameScreenGameOver: {
		Draw: func(g *Game, screen *ebiten.Image) error {
			screen.Fill(backgroundColor)

			for _, c := range g.UiComponents[GameScreenGameOver] {
				c.Draw(screen)
			}
			return nil
		},
		Update: func(g *Game) {
			for _, c := range g.UiComponents[GameScreenGameOver] {
				c.Update()
			}
		},
		Init: func(g *Game) {
			btn := ui.NewButton(50, 50, 100, func() {
				g.SetScreen(GameScreenPlay)
			}, "Play", GetResource(ResourceNameReplay))

			g.UiComponents[GameScreenGameOver] = append(g.UiComponents[GameScreenGameOver], btn)
		},
	},
}
