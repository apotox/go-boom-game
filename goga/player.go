package goga

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Position struct {
	X int
	Y int
}

type Task struct {
	taskType string
	initData map[string]interface{}
	action   func(g *Game, initData map[string]interface{}) error
}

type Player struct {
	pos           *Position
	tasks         chan Task
	power         int
	life          int
	movingSprite  *Sprite
	idleSprite    *Sprite
	rotation      float64
	direction     Dir
	nextDirection Dir
	nextTile      *Tile
	speed         int
}

func NewPlayer() *Player {
	return &Player{
		pos:           &Position{X: tileSize * 2, Y: tileSize * 2},
		tasks:         make(chan Task, 1),
		power:         1,
		life:          1,
		movingSprite:  NewSprite(GetResource("runner"), 8, 1, 32, nil),
		idleSprite:    NewSprite(GetResource("runner"), 5, 0, 32, nil),
		speed:         1,
		direction:     DirDown,
		nextDirection: DirDown,
	}
}

func (p *Player) CurrentImage() *ebiten.Image {

	return p.movingSprite.current
}

func (p *Player) Move(g *Game) {

	_, up, left, down, right, _ := GetTiles(p.GetTilePos(), g)

	if p.nextTile == nil {

		switch p.nextDirection {
		case DirUp:
			if up != nil && up.Walkable() {
				p.nextTile = up
			}
		case DirDown:
			if down != nil && down.Walkable() {
				p.nextTile = down
			}
		case DirLeft:
			if left != nil && left.Walkable() {
				p.nextTile = left
			}
		case DirRight:
			if right != nil && right.Walkable() {
				p.nextTile = right
			}
		}

		if p.nextTile != nil {
			p.direction = p.nextDirection
		} else {
			p.nextDirection = p.direction
		}

	} else {
		dx := p.nextTile.pos.X - p.pos.X
		dy := p.nextTile.pos.Y - p.pos.Y

		if abs(dx) > 0 {
			p.pos.X += (dx / abs(dx)) * p.speed
		} else if abs(dy) > 0 {
			p.pos.Y += (dy / abs(dy)) * p.speed
		} else {
			p.pos.Y = p.nextTile.pos.Y
			p.pos.X = p.nextTile.pos.X
			p.nextTile = nil

		}
	}

}

func (p *Player) GetTilePos() Position {

	return Position{X: (p.pos.X) / tileSize, Y: (p.pos.Y) / tileSize}
}

func (p *Player) Update(game *Game) error {

	if dir, ok := game.input.Dir(); ok {
		p.nextDirection = dir
	}

	if action, ok := game.input.Action(); ok {
		p.AddTask(action)
	}

	p.Move(game)

	p.RunTasks(game)

	p.movingSprite.Animate()

	return nil
}

func (p *Player) AddTask(action Action) {
	p.tasks <- Task{
		taskType: "action",
		initData: map[string]interface{}{
			"action": action,
		},
		action: func(g *Game, initData map[string]interface{}) error {
			action := initData["action"].(Action)
			fmt.Printf("executed action: %d\n", action)
			g.bombs = append(g.bombs, NewBomb(p.pos.X, p.pos.Y))
			return nil
		},
	}
}

func (p *Player) RunTasks(game *Game) error {
	select {
	case task := <-p.tasks:
		task.action(game, task.initData)
	default:
		// nothing to do
	}

	return nil
}

func (p *Player) Draw(boardImage *ebiten.Image) error {
	if boardImage == nil {
		return nil
	}

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.pos.X), float64(p.pos.Y))
	// p.movingSprite.current.Fill(playerColor)
	boardImage.DrawImage(p.movingSprite.current, op)
	return nil
}
