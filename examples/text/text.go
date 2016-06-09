package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
)

var winTitle string = "Text"
var winWidth, winHeight int = 800, 600

func run() int {
	var window *sdl.Window
	var font *ttf.Font
	var surface *sdl.Surface
	var solid *sdl.Surface
	var err error

	sdl.Init(sdl.INIT_VIDEO)

	if err := ttf.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize TTF: %s\n", err)
		return 1
	}

	if window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 2
	}
	defer window.Destroy()

	if font, err = ttf.OpenFont("../../assets/test.ttf", 32); err != nil {
		fmt.Fprint(os.Stderr, "Failed to open font: %s\n", err)
		return 4
	}
	defer font.Close()

	if solid, err = font.RenderUTF8_Solid("Hello, World!", sdl.Color{255, 0, 0, 255}); err != nil {
		fmt.Fprint(os.Stderr, "Failed to render text: %s\n", err)
		return 5
	}
	defer solid.Free()

	if surface, err = window.GetSurface(); err != nil {
		fmt.Fprint(os.Stderr, "Failed to get window surface: %s\n", err)
		return 6
	}

	if err = solid.Blit(nil, surface, nil); err != nil {
		fmt.Fprint(os.Stderr, "Failed to put text on window surface: %s\n", err)
		return 7
	}

	// Show the pixels for a while
	window.UpdateSurface()
	sdl.Delay(3000)

	return 0
}

func main() {
	os.Exit(run())
}
