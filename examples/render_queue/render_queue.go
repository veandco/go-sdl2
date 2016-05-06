// author: Jacky Boen

package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

var winTitle string = "Go-SDL2 Render"
var winWidth, winHeight int = 800, 600

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var points []sdl.Point
	var rect sdl.Rect
	var rects []sdl.Rect

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

	go func() {
		println("goroutine: A")
	}()

	sdl.CallQueue <- func() {
		renderer.SetDrawColor(255, 255, 255, 255)
		renderer.DrawPoint(150, 300)
		println("queue: A")
	}

	go func() {
		println("goroutine: B")
	}()

	sdl.CallQueue <- func() {
		renderer.SetDrawColor(0, 0, 255, 255)
		renderer.DrawLine(0, 0, 200, 200)
		println("queue: B")
	}

	go func() {
		println("goroutine: C")
	}()

	sdl.CallQueue <- func() {
		points = []sdl.Point{{0, 0}, {100, 300}, {100, 300}, {200, 0}}
		renderer.SetDrawColor(255, 255, 0, 255)
		renderer.DrawLines(points)
		println("queue: C")
	}

	go func() {
		println("goroutine: D")
	}()

	sdl.CallQueue <- func() {
		rect = sdl.Rect{300, 0, 200, 200}
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.DrawRect(&rect)
		println("queue: D")
	}

	go func() {
		println("goroutine: E")
	}()

	sdl.CallQueue <- func() {
		rects = []sdl.Rect{{400, 400, 100, 100}, {550, 350, 200, 200}}
		renderer.SetDrawColor(0, 255, 255, 255)
		renderer.DrawRects(rects)
		println("queue: E")
	}

	go func() {
		println("goroutine: F")
	}()

	sdl.CallQueue <- func() {
		rect = sdl.Rect{250, 250, 200, 200}
		renderer.SetDrawColor(0, 255, 0, 255)
		renderer.FillRect(&rect)
		println("queue: F")
	}

	go func() {
		println("goroutine: G")
	}()

	sdl.CallQueue <- func() {
		rects = []sdl.Rect{{500, 300, 100, 100}, {200, 300, 200, 200}}
		renderer.SetDrawColor(255, 0, 255, 255)
		renderer.FillRects(rects)
		println("queue: G")
	}

	go func() {
		println("goroutine: H")
	}()

	sdl.CallQueue <- func() {
		renderer.Present()
		println("queue: H")
	}

	sdl.Delay(2000)

	return 0
}

func main() {
	os.Exit(run())
}
