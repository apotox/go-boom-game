package game

import (
	"fmt"

	"github.com/apotox/goga/joystick"
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

type PlayerState string

const (
	PlayerStateIdle PlayerState = "idle"
	PlayerStateWalk PlayerState = "walk"
	PlayerStateDie  PlayerState = "die"
)

type Player struct {
	pos           *Position
	tasks         chan Task
	power         int
	sprites       map[PlayerState]ISprite
	particles     []*Particle
	state         PlayerState
	rotation      float64
	direction     joystick.Dir
	oldDirection  joystick.Dir
	nextDirection joystick.Dir
	nextTile      *Tile
}

func NewPlayer() *Player {

	sprites := make(map[PlayerState]ISprite)
	sprites[PlayerStateWalk] = NewAnimatedSprite(GetResource(ResourceNameLizardRun), 4, 0, 32, &Offsets{
		offsetX: 0,
		offsetY: 12,
	}, true)
	sprites[PlayerStateIdle] = NewAnimatedSprite(GetResource(ResourceNameLizardIdle), 3, 0, 32, &Offsets{
		offsetX: 0,
		offsetY: 4,
	}, true)

	particles := make([]*Particle, 0)

	p := &Player{
		pos:           &Position{X: tileSize * 2, Y: tileSize * 2},
		tasks:         make(chan Task, 2),
		power:         1,
		sprites:       sprites,
		direction:     joystick.DirDown,
		nextDirection: joystick.DirDown,
		state:         PlayerStateIdle,
		particles:     particles,
	}

	p.particles = append(particles, NewParticle(ParticleDust, p.pos))

	return p
}

func (p *Player) CurrentImage() *ebiten.Image {

	return p.sprites[p.state].GetCurrent()
}

func (p *Player) GetFeatures() Features {
	return allFeatures[p.power]
}

func (p *Player) UpgradePlayer() error {
	p.power++
	return nil
}

func (p *Player) GetNextTile(g *Game, direction joystick.Dir) *Tile {
	_, up, left, down, right, _, _ := GetSurroundedTiles(GetTileBoardPos(p.pos), g)

	switch direction {
	case joystick.DirUp:
		if up != nil && up.Walkable() {
			return up
		}
	case joystick.DirDown:
		if down != nil && down.Walkable() {
			return down
		}
	case joystick.DirLeft:
		if left != nil && left.Walkable() {
			return left
		}
	case joystick.DirRight:
		if right != nil && right.Walkable() {
			return right
		}
	}
	return nil
}

func (p *Player) Move(g *Game) {

	if p.nextTile == nil {
		if p.nextDirection != p.direction {
			n := p.GetNextTile(g, p.nextDirection)
			if n != nil {
				p.direction = p.nextDirection
			}
		}
		p.nextTile = p.GetNextTile(g, p.direction)
		// still nil?
		if p.nextTile == nil {
			p.state = PlayerStateIdle
		} else {
			p.state = PlayerStateWalk
		}

	} else {

		dx := p.nextTile.pos.X - p.pos.X
		dy := p.nextTile.pos.Y - p.pos.Y

		if joystick.Abs(dx) > 0 {
			p.pos.X += (dx / joystick.Abs(dx)) * p.GetFeatures().speed
		} else if joystick.Abs(dy) > 0 {
			p.pos.Y += (dy / joystick.Abs(dy)) * p.GetFeatures().speed
		} else {
			p.pos.Y = p.nextTile.pos.Y
			p.pos.X = p.nextTile.pos.X
			p.nextTile = nil

		}
	}

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
	p.Animate(game)

	//draw particles
	// for _, particle := range p.particles {
	// 	particle.Update(game)
	// }

	return nil
}

func (p *Player) Animate(game *Game) error {

	p.sprites[p.state].Animate()

	return nil
}

func (p *Player) GetPosition() *Position {
	return p.pos
}
func (p *Player) GetName() string {
	return "player"
}
func (p *Player) GetSize() int {
	return tileSize
}

func (p *Player) AddTask(action joystick.Action) {
	p.tasks <- Task{
		taskType: "action",
		initData: map[string]interface{}{
			"action": action,
		},
		action: func(g *Game, initData map[string]interface{}) error {
			// action := initData["action"].(Action)

			if len(g.bombs) < p.GetFeatures().maxItems {
				g.AddBomb(p.pos, 3)
			}

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

	//flip image
	if p.direction == joystick.DirLeft {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(tileSize, 0)
	}

	op.GeoM.Translate(float64(p.pos.X), float64(p.pos.Y))

	boardImage.DrawImage(p.sprites[p.state].GetCurrent(), op)

	//draw particles
	// for _, particle := range p.particles {
	// 	particle.Draw(boardImage)
	// }
	return nil
}

func (p *Player) Die(g *Game) error {
	//p.state = PlayerStateDie
	fmt.Print("player died")
	g.SetScreen(GameScreenGameOver)
	return nil
}
