package main

import (
	"log"

	"github.com/apotox/goga/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game.LoadResources()
	g := game.NewGame()

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(game.ScreenWidth*game.Zoom, game.ScreenHeight*game.Zoom)
	ebiten.SetWindowTitle("boom")
	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
