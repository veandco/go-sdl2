// author: Jacky Boen

package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"

	"github.com/veandco/go-sdl2/sdl"
)

var winTitle string = "Go-SDL2 Render"
var winWidth, winHeight int = 800, 600

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer

	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
	renderer.Clear()
	defer renderer.Destroy()

	// Set a WaitGroup to wait until all pixels are drawn
	var wg sync.WaitGroup

	for y := 0; y < winHeight; y++ {
		for x := 0; x < winWidth; x++ {
			wg.Add(1)

			go func(x, y int) {
				// Do some fake processing before rendering
				r := byte(rand.Int())
				g := byte(rand.Int())
				b := byte(rand.Int())

				// Call the render function in the 'render' thread synchronously
				sdl.CallQueue <- func() {
					renderer.SetDrawColor(r, g, b, 255)
					renderer.DrawPoint(x, y)
					wg.Done()
				}
			}(x, y)
		}
	}

	// Wait until all pixels are drawn
	wg.Wait()

	// Show the pixels for a while
	renderer.Present()
	sdl.Delay(3000)

	return 0
}

func main() {
	os.Exit(run())
}
