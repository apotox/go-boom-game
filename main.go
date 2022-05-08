package main

import (
	"log"

	"github.com/apotox/goga/goga"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	goga.LoadResources()
	game := goga.NewGame()

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(380, 580)
	ebiten.SetWindowTitle("Goga boom")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
