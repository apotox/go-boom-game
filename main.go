package main

import (
	"log"

	"github.com/apotox/goga/game"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ScreenWidth  = 12 * game.GetTileSize()
	ScreenHeight = 16 * game.GetTileSize()
)

func main() {
	game.LoadResources()
	g := game.NewGame()
	zoom := 2

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(ScreenWidth*zoom, ScreenHeight*zoom)
	ebiten.SetWindowTitle("Goga boom")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
