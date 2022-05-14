package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/apotox/goga/joystick"
	ui "github.com/apotox/goga/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	level          int
	board          *Board
	player1        *Player
	bombs          []*Bomb
	enemies        []*Enemy
	pickables      []*Pickable
	boardImage     *ebiten.Image
	input          *joystick.Input
	pickableTicker *time.Ticker
	gameScreen     GameScreen
	UiComponents   map[GameScreen][]ui.Component
}

func (g *Game) SetScreen(s GameScreen) {
	g.gameScreen = s
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	screens[g.gameScreen].Update(g)
	return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	screens[g.gameScreen].Draw(g, screen)
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 12 * tileSize, 16 * tileSize
}

func NewGame() *Game {

	game := &Game{
		level:          0,
		player1:        NewPlayer(),
		input:          joystick.NewInput(),
		pickableTicker: time.NewTicker(time.Second * 10),
		board:          GetLevelBoard(0),
		gameScreen:     GameScreenPlay,
		UiComponents:   make(map[GameScreen][]ui.Component),
	}

	// init screens
	for _, screen := range screens {
		if screen.Init != nil {
			screen.Init(game)
		}
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

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %+v\n", r)
		}
	}()
}

func (g *Game) AddBomb(pos *Position, lifeTime int) {
	index := len(g.bombs)
	g.bombs = append(g.bombs, NewBomb(index, pos.X, pos.Y, lifeTime))
}

func (g *Game) AddPickable(kind PickableEnum) {

	tIndex := rand.Intn(len(g.board.tiles) - 1)

	if tIndex < 0 {
		tIndex = 0 // weird bug
	}

	for g.board.tiles[tIndex].kind != TileKindEmpty {
		tIndex = rand.Intn(len(g.board.tiles) - 1)
	}

	if g.board.tiles[tIndex] != nil {
		g.pickables = append(g.pickables, NewPickable(kind, g.board.tiles[tIndex].pos))
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Panic: %+v\n", r)
		}
	}()
}
