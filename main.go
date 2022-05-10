package main

import (
	"log"

	"github.com/apotox/goga/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game.LoadResources()
	g := game.NewGame()
	zoom := 2

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(12*game.GetTileSize()*zoom, 16*game.GetTileSize()*zoom)
	ebiten.SetWindowTitle("Goga boom")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
