package game

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/apotox/go-boom-game/joystick"
	ui "github.com/apotox/go-boom-game/ui"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	Zoom         = 2
	ScreenWidth  = 12 * GetTileSize()
	ScreenHeight = 16 * GetTileSize()
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
	enemyTicker    *time.Ticker
	gameScreen     GameScreen
	UiComponents   map[GameScreen][]ui.Component

	reservedTiles []*Tile
	header        *Header
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
	return 12 * tileSize, 16*tileSize + 25
}

func NewGame() *Game {

	level := 0

	game := &Game{
		level:          level,
		player1:        NewPlayer(),
		input:          joystick.NewInput(),
		pickableTicker: time.NewTicker(time.Second * 10),
		enemyTicker:    time.NewTicker(time.Second * 3),
		board:          GetLevelBoard(level),
		gameScreen:     GameScreenGameOver,
		UiComponents:   make(map[GameScreen][]ui.Component),
		header:         NewHeader(),
	}

	// init screens
	for _, screen := range screens {
		if screen.Init != nil {
			screen.Init(game)
		}
	}

	return game
}

func (g *Game) AddEnemy() {

	rand.Seed(time.Now().UnixNano())

	if len(g.enemies) > 1 {
		return
	}
	emptyTiles := FilterTiles(g.board.tiles, func(t *Tile) bool {
		return t.Walkable()
	})

	tIndex := rand.Intn(len(emptyTiles) - 1)
	enemy := NewEnemy(&Position{
		X: emptyTiles[tIndex].pos.X,
		Y: emptyTiles[tIndex].pos.Y,
	})
	g.enemies = append(g.enemies, enemy)

}

func (g *Game) RemoveEnemy(enemy *Enemy) {
	for i, e := range g.enemies {
		if e == enemy {
			g.enemies = append(g.enemies[:i], g.enemies[i+1:]...)
		}
	}
}

func (g *Game) RemoveBomb(bomb *Bomb) {
	h := GetTileBoardPos(bomb.pos)
	p := GetTileByBoardPosition(&h, g)
	p.reserved = false

	for i, b := range g.bombs {
		if b == bomb {
			g.bombs = append(g.bombs[:i], g.bombs[i+1:]...)
		}
	}
}

func (g *Game) AddBomb(pos *Position, lifeTime int) {
	g.bombs = append(g.bombs, NewBomb(g, pos.X, pos.Y, lifeTime))
}

func (g *Game) AddPickable(kind PickableEnum) {

	if len(g.pickables) > 0 {
		return
	}
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
