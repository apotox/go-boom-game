//go:generate file2byteslice -package=images -input=./images/door.png -output=./images/door.go -var=Door_png
//go:generate file2byteslice -package=images -input=./images/lizard_idle.png -output=./images/lizard_idle.go -var=Lizard_idle_png
//go:generate file2byteslice -package=images -input=./images/lizard_run.png -output=./images/lizard_run.go -var=Lizard_run_png
//go:generate file2byteslice -package=images -input=./images/chort_run.png -output=./images/chort_run.go -var=Chort_run_png
//go:generate file2byteslice -package=images -input=./images/chort_idle.png -output=./images/chort_idle.go -var=Chort_idle_png
//go:generate file2byteslice -package=images -input=./images/chest_anim.png -output=./images/chest_anim.go -var=Chest_anim_png
//go:generate file2byteslice -package=images -input=./images/wall_fountain.png -output=./images/wall_fountain.go -var=Wall_fountain_png
//go:generate file2byteslice -package=images -input=./images/orc_idle.png -output=./images/orc_idle.go -var=Orc_idle_png
//go:generate file2byteslice -package=images -input=./images/orc_run.png -output=./images/orc_run.go -var=Orc_run_png
//go:generate file2byteslice -package=images -input=./images/wall.png -output=./images/wall.go -var=Wall_png
//go:generate file2byteslice -package=images -input=./images/flask_blue.png -output=./images/flask_blue.go -var=Flask_blue_png
//go:generate file2byteslice -package=images -input=./images/floor.png -output=./images/floor.go -var=Floor_png
//go:generate file2byteslice -package=images -input=./images/boom_on.png -output=./images/boom_on.go -var=Boom_on_png
//go:generate file2byteslice -package=images -input=./images/boom_idle.png -output=./images/boom_idle.go -var=Boom_idle_png
//go:generate file2byteslice -package=fonts -input=./fonts/PixelAEBold.ttf -output=./fonts/PixelAEBold.go -var=PixelAEBold_ttf

//go:generate gofmt -s -w .

package resources

import (
	// Dummy imports for go.mod for some Go files with 'ignore' tags. For example, `go mod tidy` does not
	// recognize Go files with 'ignore' build tag.
	//
	// Note that this affects only importing this package, but not 'file2byteslice' commands in //go:generate.
	_ "github.com/hajimehoshi/file2byteslice"
)
