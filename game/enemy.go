package game

import "github.com/hajimehoshi/ebiten/v2"

type EnemyFamily string

const (
	Mikky EnemyFamily = "mikky"
	Tikky EnemyFamily = "tikky"
	Gikky EnemyFamily = "gikky"
)

type EnemyState string

const (
	EnemyStateIdle EnemyState = "idle"
	EnemyStateWalk EnemyState = "walk"
	EnemyStateDie  EnemyState = "die"
)

type Enemy struct {
	pos      *Position
	tasks    []Task
	power    int
	life     int
	state    EnemyState
	nextTile *Tile
	speed    int
	sprites  map[EnemyState]*Sprite
	family   EnemyFamily
}

func (e *Enemy) Update(g *Game) error {

	if e.nextTile == nil {

	}

	return nil
}

func (e *Enemy) Draw(boardImage *ebiten.Image) error {

	return nil
}

func NewEnemy(pos *Position) *Enemy {
	return &Enemy{
		pos:   pos,
		tasks: []Task{},
		power: 1,
		life:  1,
		state: EnemyStateIdle,
		speed: 1,
	}
}
