package game

type Features struct {
	power    int
	speed    int
	life     int
	maxItems int
}

type AllFeatures = map[int]Features

var allFeatures = AllFeatures{
	1: Features{
		power:    1,
		speed:    1,
		life:     1,
		maxItems: 10,
	},
	2: Features{
		power:    2,
		speed:    2,
		life:     2,
		maxItems: 2,
	},
}
