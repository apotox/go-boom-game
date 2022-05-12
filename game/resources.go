package game

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

type ResourceName string

const (
	ResourceNamePlayer ResourceName = "player"
	ResourceNameTiles  ResourceName = "tiles"
	ResourceNameRunner ResourceName = "runner"
	ResourceNameBomb   ResourceName = "bomb"
	ResourceNameDust   ResourceName = "dust"
	ResourceNameWood   ResourceName = "wood"
	ResourceNameReplay ResourceName = "replay"
)

var _resources = map[ResourceName]*ebiten.Image{}

func LoadResources() error {

	_resources[ResourceNamePlayer] = getImage(images.Player_png)
	_resources[ResourceNameRunner] = getImage(images.Runner_png)
	_resources[ResourceNameTiles] = getImage(images.Tiles_png)
	_resources[ResourceNameBomb] = getImage(images.Bomb_png)
	_resources[ResourceNameDust] = getImage(images.Dust_png)
	_resources[ResourceNameWood] = getImage(images.Wood_png)
	_resources[ResourceNameReplay] = getImage(images.Replay_png)
	return nil
}

func GetResource(name ResourceName) *ebiten.Image {

	if len(_resources) == 0 {
		LoadResources()
	}

	if _, ok := _resources[name]; ok {
		return _resources[name]
	}

	panic("resource not found")
}
