package goga

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"

	"github.com/apotox/goga/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
)

func getImage(b []byte) *ebiten.Image {

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		fmt.Printf("failed to decode image: %s", err)

		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

var _resources = map[string]*ebiten.Image{}

func LoadResources() error {

	_resources["player"] = getImage(images.Player_png)
	_resources["runner"] = getImage(images.Runner_png)
	_resources["tiles"] = getImage(images.Tiles_png)
	return nil
}

func GetResource(name string) *ebiten.Image {

	if len(_resources) == 0 {
		LoadResources()
	}

	if _, ok := _resources[name]; ok {
		return _resources[name]
	}

	panic("resource not found")
}
