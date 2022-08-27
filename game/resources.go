package game

import (
	"bytes"
	"fmt"
	"image"
	_ "image/png"
	"log"

	fonts "github.com/apotox/go-boom-game/resources/fonts"
	"github.com/apotox/go-boom-game/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
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
var _fonts = map[string]font.Face{}

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

	tt, err := opentype.Parse(fonts.PixelAEBold_ttf)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 62
	PixelAEBold14, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    14,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	_fonts["default"] = PixelAEBold14

	PixelAEBold12, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    12,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})

	_fonts["default-12"] = PixelAEBold12

	return nil
}

func GetFont(name string) font.Face {

	if len(_resources) == 0 {
		LoadResources()
	}

	if _, ok := _fonts[name]; ok {
		return _fonts[name]
	}

	panic("resource not found")
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
