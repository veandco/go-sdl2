package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_gfx"
)

var winTitle string = "SDL2 GFX"
var winWidth, winHeight int = 800, 600

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var vx, vy = make([]int16, 3), make([]int16, 3)
	var wg sync.WaitGroup // Set a WaitGroup to wait until all pixels are drawn
	var err error

	if window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	if renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
	renderer.Clear()
	defer renderer.Destroy()

	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			wg.Add(1)

			go func(x, y int) {
				// Call the render function in the 'render' thread synchronously
				sdl.CallQueue <- func() {
					gfx.PixelColor(renderer, x, y, rand.Uint32())
					wg.Done()
				}
			}(x, y)
		}
	}

	// Wait until all pixels are drawn
	wg.Wait()

	vx[0] = int16(winWidth / 3)
	vy[0] = int16(winHeight / 3)
	vx[1] = int16(winWidth * 2 / 3)
	vy[1] = int16(winHeight / 3)
	vx[2] = int16(winWidth / 2)
	vy[2] = int16(winHeight * 2 / 3)
	gfx.FilledPolygonColor(renderer, vx, vy, 0xFF0000FF)

	gfx.CharacterColor(renderer, winWidth - 16, 16, 'X', 0xFFFF0000)
	gfx.StringColor(renderer, 16, 16, "GFX Demo", 0xFF00FF00)

	renderer.Present()
	sdl.Delay(3000)

	return 0
}

func main() {
	os.Exit(run())
}
