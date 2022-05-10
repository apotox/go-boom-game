//go:generate file2byteslice -package=images -input=./images/player.png -output=./images/player.go -var=Player_png
//go:generate file2byteslice -package=images -input=./images/runner.png -output=./images/runner.go -var=Runner_png
//go:generate file2byteslice -package=images -input=./images/tiles.png -output=./images/tiles.go -var=Tiles_png
//go:generate file2byteslice -package=images -input=./images/bomb.png -output=./images/bomb.go -var=Bomb_png
//go:generate file2byteslice -package=images -input=./images/dust.png -output=./images/dust.go -var=Dust_png
//go:generate gofmt -s -w .

package resources

import (
	// Dummy imports for go.mod for some Go files with 'ignore' tags. For example, `go mod tidy` does not
	// recognize Go files with 'ignore' build tag.
	//
	// Note that this affects only importing this package, but not 'file2byteslice' commands in //go:generate.
	_ "github.com/hajimehoshi/file2byteslice"
)
