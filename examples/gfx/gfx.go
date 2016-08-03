package main

import (
	"fmt"
	//"math/rand"
	"os"
	"sync"
	//"time"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_gfx"
)

var winTitle string = "SDL2 GFX"
var winWidth, winHeight int = 800, 600

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var wg sync.WaitGroup // Set a WaitGroup to wait until all pixels are drawn
	var err error

	if window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, winWidth, winHeight, sdl.WINDOW_SHOWN); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	if renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_SOFTWARE); err != nil {
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
					gfx.PixelColor(renderer, x, y, 0xffff00ff)
					wg.Done()
				}
			}(x, y)
		}
	}

	// Wait until all pixels are drawn
	wg.Wait()
	renderer.Present()
	sdl.Delay(3000)

	return 0
}

func main() {
	os.Exit(run())
}
