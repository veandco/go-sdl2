// author: Jacky Boen

package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
	"os"
)

var winTitle string = "Go-SDL2 Render"
var winWidth, winHeight int = 800, 600

func main() {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var points []sdl.Point
	var rect sdl.Rect
	var rects []sdl.Rect

	window = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
					winWidth, winHeight, sdl.WINDOW_SHOWN)
	if window == nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", sdl.GetError());
		os.Exit(1);
	}

	renderer = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if renderer == nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", sdl.GetError());
		os.Exit(2);
	}

	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.DrawPoint(150, 300)

	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.DrawLine(0, 0, 200, 200)

	points = []sdl.Point { {0, 0}, {100, 300}, {100, 300}, {200, 0} }
	renderer.SetDrawColor(255, 255, 0, 255)
	renderer.DrawLines(points)

	rect = sdl.Rect { 300, 0, 200, 200 }
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawRect(&rect)

	rects = []sdl.Rect { {400, 400, 100, 100}, {550, 350, 200, 200} }
	renderer.SetDrawColor(0, 255, 255, 255)
	renderer.DrawRects(rects)

	rect = sdl.Rect { 250, 250, 200, 200 }
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.FillRect(&rect)

	rects = []sdl.Rect { {500, 300, 100, 100}, {200, 300, 200, 200} }
	renderer.SetDrawColor(255, 0, 255, 255)
	renderer.FillRects(rects)

	renderer.Present()

	sdl.Delay(2000)

	renderer.Destroy()
	window.Destroy()
}
