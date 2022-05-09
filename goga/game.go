package goga

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	level      int
	board      *Board
	player1    *Player
	bombs      []*Bomb
	enemies    []*Enemy
	boardImage *ebiten.Image
	input      *Input
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {

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

	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(screen, "Hello, World!")

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

	screen.DrawImage(g.boardImage, &ebiten.DrawImageOptions{})
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 200, 200
}

func NewGame() *Game {

	board := GetLevelBoard(0)

	game := &Game{
		level:   1,
		board:   board,
		player1: NewPlayer(),
		input:   NewInput(),
	}

	return game
}

func (g *Game) AddEnemy(board *Board, pos *Position) {

	enemy := NewEnemy(pos)
	g.enemies = append(g.enemies, enemy)

}

func (g *Game) RemoveBomb(index int) {

	for i, bomb := range g.bombs {
		if bomb.index == index {
			g.bombs = append(g.bombs[:i], g.bombs[i+1:]...)
			break
		}
	}
}

func (g *Game) AddBomb(pos *Position, lifeTime int) {
	index := len(g.bombs)
	g.bombs = append(g.bombs, NewBomb(index, pos.X, pos.Y, lifeTime))
}
