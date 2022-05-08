package goga

type Enemy struct {
	pos   *Position
	tasks []Task
	power int
	life  int
}

func (e *Enemy) Update() error {
	return nil
}

func NewEnemy(pos *Position) *Enemy {
	return &Enemy{
		pos:   pos,
		tasks: []Task{},
		power: 1,
		life:  1,
	}
}
