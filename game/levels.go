package game

type Level struct {
	name  string
	index int
	tiles [][]int
}

var levels = []*Level{
	{
		name:  "level1",
		index: 0,
		tiles: [][]int{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			{1, 0, 1, 1, 1, 1, 1, 1, 0, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			{1, 1, 1, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			{1, 1, 0, 1, 1, 1, 1, 0, 1, 1, 1, 1},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			{1, 1, 1, 1, 0, 1, 1, 1, 1, 0, 1, 1},
			{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			{1, 1, 1, 1, 0, 1, 1, 1, 1, 1, 0, 1},
			{1, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 1},
			{1, 1, 0, 0, 1, 1, 1, 1, 1, 0, 1, 1},
			{1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 1},
			{1, 1, 0, 1, 1, 1, 1, 1, 0, 1, 1, 1},
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		},
	},
}

func GetLevelBoard(index int) *Board {

	level := levels[index]

	widthSize := len(level.tiles[0])
	heightSize := len(level.tiles)

	b := &Board{
		widthSize:  widthSize,
		heightSize: heightSize,
	}

	for y := 0; y < heightSize; y++ {
		for x := 0; x < widthSize; x++ {
			b.tiles = append(b.tiles, NewTile(x*tileSize, y*tileSize, level.tiles[y][x]))
		}
	}

	return b

}
