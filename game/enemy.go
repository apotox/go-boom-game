package game

import (
	"github.com/apotox/goga/joystick"
	"github.com/hajimehoshi/ebiten/v2"
)

type EnemyFamily string

const (
	Chort EnemyFamily = "chort"
	Orc   EnemyFamily = "orc"
)

type EnemyState string

const (
	EnemyStateIdle   EnemyState = "idle"
	EnemyStateAttack EnemyState = "walk"
)

type PrevEntry struct {
	tile      *Tile
	direction joystick.Dir
}
type Enemy struct {
	pos           *Position
	tasks         []Task
	power         int
	life          int
	state         EnemyState
	nextTile      *Tile
	prevDir       joystick.Dir
	speed         int
	sprites       map[EnemyState]ISprite
	family        EnemyFamily
	direction     joystick.Dir
	nextDirection joystick.Dir
}

func (e *Enemy) GetFeatures() Features {
	return allFeatures[e.power]
}

func (e *Enemy) AllowedTile(t *Tile) bool {
	if !t.Walkable() {
		return false
	}

	dx := t.pos.X - e.pos.X
	dy := t.pos.Y - e.pos.Y

	if joystick.Abs(dx) > 0 {
		if dx > 0 {
			e.nextDirection = joystick.DirRight
		} else {
			e.nextDirection = joystick.DirLeft
		}
	} else if joystick.Abs(dy) > 0 {

		if dy > 0 {
			e.nextDirection = joystick.DirUp
		} else {
			e.nextDirection = joystick.DirDown
		}
	}

	return e.nextDirection.Oposite() != e.direction
}

func (e *Enemy) GetNearTile(g *Game) *Tile {

	enemyPos := GetTileBoardPos(e.pos)
	playerPos := GetTileBoardPos(g.player1.pos)

	_, _, _, _, _, tiles, _ := GetSurroundedTiles(enemyPos, g)

	var nearTile *Tile = nil

	for _, t := range tiles {
		if t != nil && e.AllowedTile(t) {

			if nearTile == nil {
				nearTile = t
				continue
			}
			tb := GetTileBoardPos(t.pos)
			nt := GetTileBoardPos(nearTile.pos)

			if Distance(tb, playerPos) < Distance(nt, playerPos) {
				nearTile = t
			}
		}
	}

	if nearTile == nil {
		e.direction = e.direction.Oposite()
	}

	return nearTile

}

func (e *Enemy) Move(g *Game) {
	if e.nextTile == nil {

		t := e.GetNearTile(g)

		if t == nil {
			e.state = EnemyStateIdle
		} else {
			e.state = EnemyStateAttack
			e.nextTile = t
		}

	} else {

		dx := e.nextTile.pos.X - e.pos.X
		dy := e.nextTile.pos.Y - e.pos.Y

		if joystick.Abs(dx) > 0 {

			if dx > 0 {
				e.direction = joystick.DirRight
			} else {
				e.direction = joystick.DirLeft
			}

			e.pos.X += (dx / joystick.Abs(dx)) * e.GetFeatures().speed
		} else if joystick.Abs(dy) > 0 {

			if dy > 0 {
				e.direction = joystick.DirUp
			} else {
				e.direction = joystick.DirDown
			}

			e.pos.Y += (dy / joystick.Abs(dy)) * e.GetFeatures().speed

		} else {
			e.pos.Y = e.nextTile.pos.Y
			e.pos.X = e.nextTile.pos.X
			e.nextTile = nil

		}
	}
}

func (e *Enemy) Update(g *Game) error {

	e.sprites[e.state].Animate()

	//if e.state == EnemyStateAttack {
	e.Move(g)
	//}

	return nil
}

func (e *Enemy) Draw(boardImage *ebiten.Image) error {
	if boardImage == nil {
		return nil
	}
	op := &ebiten.DrawImageOptions{}

	//flip image
	if e.direction == joystick.DirLeft {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(tileSize, 0)
	}

	op.GeoM.Translate(float64(e.pos.X), float64(e.pos.Y))

	boardImage.DrawImage(e.sprites[e.state].GetCurrent(), op)

	return nil
}

func (p *Enemy) GetPosition() *Position {
	return p.pos
}
func (p *Enemy) GetName() string {
	return "enemy"
}
func (p *Enemy) GetSize() int {
	return tileSize
}

func NewEnemy(pos *Position) *Enemy {
	e := &Enemy{
		pos:       pos,
		tasks:     []Task{},
		power:     1,
		life:      1,
		state:     EnemyStateIdle,
		speed:     1,
		sprites:   make(map[EnemyState]ISprite),
		direction: joystick.DirDown,
	}

	e.sprites[EnemyStateIdle] = NewAnimatedSprite(GetResource(ResourceNameChortIdle), 4, 0, 10, &Offsets{
		offsetX: 0,
		offsetY: 8,
	}, true)

	e.sprites[EnemyStateAttack] = NewAnimatedSprite(GetResource(ResourceNameChostRun), 3, 0, 10, &Offsets{
		offsetX: 0,
		offsetY: 8,
	}, true)

	return e
}
