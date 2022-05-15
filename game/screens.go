package game

import (
	"image/color"

	"github.com/apotox/goga/mycolors"
	ui "github.com/apotox/goga/ui"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
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
			g.boardImage.Fill(mycolors.BackgroundColor)

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

			if len(g.enemies) > 0 {
				for _, enemy := range g.enemies {
					enemy.Draw(g.boardImage)
				}
			}

			select {
			case <-g.pickableTicker.C:
				g.AddPickable(PickableEnumPower)
			default:
				// nothing
			}

			select {
			case <-g.enemyTicker.C:
				g.AddEnemy()
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

			if len(g.enemies) > 0 {
				for _, enemy := range g.enemies {
					enemy.Update(g)
				}
			}
		},
	},
	GameScreenGameOver: {
		Draw: func(g *Game, screen *ebiten.Image) error {
			screen.Fill(mycolors.BackgroundColor)

			for _, c := range g.UiComponents[GameScreenGameOver] {
				c.Draw(screen)
			}

			text.Draw(screen, "GOGA GAME", GetFont("default"), ScreenWidth/2-36, ScreenHeight-4, color.White)

			return nil
		},
		Update: func(g *Game) {
			for _, c := range g.UiComponents[GameScreenGameOver] {
				c.Update(g.input)
			}
		},
		Init: func(g *Game) {

			btnPlay := ui.NewButton(ScreenWidth/2, ScreenHeight/2, 40, "Play", GetFont("default"), GetResource(ResourceNameReplay), func() {
				g.SetScreen(GameScreenPlay)
			})

			g.UiComponents[GameScreenGameOver] = append(g.UiComponents[GameScreenGameOver], btnPlay)
		},
	},
}
