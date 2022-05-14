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
	ResourceNameChortIdle    ResourceName = "chort_idle"
	ResourceNameChostRun     ResourceName = "chort_run"
	ResourceNameDoor         ResourceName = "door"
	ResourceNameWall         ResourceName = "wall"
	ResourceNameWallFountain ResourceName = "wall_fountain"
	ResourceNameLizardIdle   ResourceName = "lizard_idle"
	ResourceNameLizardRun    ResourceName = "lizard_run"
	ResourceNameChestAnim    ResourceName = "chest_anim"
	ResourceNameFlaskBlue    ResourceName = "flask_blue"
	ResourceNameFloor        ResourceName = "floor"
	ResourceNameReplay       ResourceName = "replay"
	ResourceNameBoomIdle     ResourceName = "boom_idle"
	ResourceNameBoomOn       ResourceName = "boom_on"
)

var _resources = map[ResourceName]*ebiten.Image{}

func LoadResources() error {

	_resources[ResourceNameChortIdle] = getImage(images.Chort_idle_png)
	_resources[ResourceNameChostRun] = getImage(images.Chort_run_png)
	_resources[ResourceNameDoor] = getImage(images.Door_png)
	_resources[ResourceNameWall] = getImage(images.Wall_png)
	_resources[ResourceNameWallFountain] = getImage(images.Wall_fountain_png)
	_resources[ResourceNameLizardIdle] = getImage(images.Lizard_idle_png)
	_resources[ResourceNameLizardRun] = getImage(images.Lizard_run_png)
	_resources[ResourceNameChestAnim] = getImage(images.Chest_anim_png)
	_resources[ResourceNameFlaskBlue] = getImage(images.Flask_blue_png)
	_resources[ResourceNameFloor] = getImage(images.Floor_png)
	_resources[ResourceNameReplay] = getImage(images.Replay_png)
	_resources[ResourceNameBoomIdle] = getImage(images.Boom_idle_png)
	_resources[ResourceNameBoomOn] = getImage(images.Boom_on_png)
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
