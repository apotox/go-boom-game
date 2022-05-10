package game

type Killable interface {
	GetName() string
	Die() error
}

type EndLife struct {
	entities []Killable
}
