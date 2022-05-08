package goga

import "github.com/hajimehoshi/ebiten/v2"

func ScaleImage(img *ebiten.Image) *ebiten.Image {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(float64(tileSize)/float64(img.Bounds().Size().X), float64(tileSize)/float64(img.Bounds().Size().Y))

	u := ebiten.NewImage(tileSize, tileSize)
	u.DrawImage(img, op)

	return u
}

func Filter(arr []Task, f func(Task) bool) []interface{} {
	var r []interface{}
	for _, v := range arr {
		if f(v) {
			r = append(r, v)
		}
	}
	return r
}
