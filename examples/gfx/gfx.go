package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "SDL2 GFX"
var winWidth, winHeight int32 = 800, 600

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var vx, vy = make([]int16, 3), make([]int16, 3)
	var err error

	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize SDL: %s\n", err)
		return 1
	}
	defer sdl.Quit()

	if window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 2
	}
	defer window.Destroy()

	if renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 3 // don't use os.Exit(3); otherwise, previous deferred calls will never run
	}
	renderer.Clear()
	defer renderer.Destroy()

	vx[0] = int16(winWidth / 3)
	vy[0] = int16(winHeight / 3)
	vx[1] = int16(winWidth * 2 / 3)
	vy[1] = int16(winHeight / 3)
	vx[2] = int16(winWidth / 2)
	vy[2] = int16(winHeight * 2 / 3)
	gfx.FilledPolygonColor(renderer, vx, vy, sdl.Color{0, 0, 255, 255})

	gfx.CharacterColor(renderer, winWidth-16, 16, 'X', sdl.Color{255, 0, 0, 255})
	gfx.StringColor(renderer, 16, 16, "GFX Demo", sdl.Color{0, 255, 0, 255})

	renderer.Present()
	sdl.Delay(3000)

	return 0
}

func main() {
	os.Exit(run())
}
