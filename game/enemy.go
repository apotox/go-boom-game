package game

import (
	"github.com/apotox/go-boom-game/joystick"
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
	EnemyStateDead   EnemyState = "dead"
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
	sprites       map[EnemyState]ISprite
	family        EnemyFamily
	direction     joystick.Dir
	nextDirection joystick.Dir
	step          float32
	stepLength    float32
}

func (e *Enemy) GetFeatures() Features {
	return allFeatures[e.power]
}

func (e *Enemy) AllowedTile(t *Tile) bool {
	if !t.Walkable() {
		return false
	}

	_, _, dir := GetDxDyDir(t.pos, e.pos)

	e.nextDirection = dir

	return e.nextDirection.Oposite() != e.direction
}

func (e *Enemy) GetNearTile(g *Game) *Tile {

	enemyPos := GetTileBoardPos(e.pos)
	playerPos := GetTileBoardPos(g.player1.pos)

	tiles, _ := GetSurroundedTilesArray(enemyPos, g)

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

		if t != nil {
			e.nextTile = t
		}

	} else {

		if e.step += e.stepLength; e.step > 1 {
			e.step = 0
		} else {
			return
		}

		dx, dy, dir := GetDxDyDir(e.nextTile.pos, e.pos)

		if joystick.Abs(dx) > 0 {
			e.direction = dir
			e.pos.X += (dx / joystick.Abs(dx)) * e.GetFeatures().speed
		} else if joystick.Abs(dy) > 0 {
			e.direction = dir
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

	if e.state == EnemyStateAttack {
		e.Move(g)
	}

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
		pos:        pos,
		tasks:      []Task{},
		power:      1,
		life:       1,
		state:      EnemyStateIdle,
		sprites:    make(map[EnemyState]ISprite),
		direction:  joystick.DirDown,
		stepLength: 0.4,
		step:       0,
	}

	e.sprites[EnemyStateIdle] = NewAnimatedSprite(GetResource(ResourceNameChortIdle), 4, 0, 10, &Offsets{
		offsetX: 0,
		offsetY: 8,
	}, true)

	e.sprites[EnemyStateAttack] = NewAnimatedSprite(GetResource(ResourceNameChostRun), 3, 0, 10, &Offsets{
		offsetX: 0,
		offsetY: 8,
	}, true)

	e.sprites[EnemyStateDead] = NewAnimatedSprite(GetResource(ResourceNameChortIdle), 4, 0, 10, &Offsets{
		offsetX: 0,
		offsetY: 8,
	}, true)

	SetTimeout(3000, func() {
		e.state = EnemyStateAttack
	})
	return e
}

func (e *Enemy) Die(g *Game) error {
	if e.state == EnemyStateDead {
		return nil
	}

	e.state = EnemyStateDead
	SetTimeout(2000, func() {
		g.RemoveEnemy(e)
	})
	return nil
}
