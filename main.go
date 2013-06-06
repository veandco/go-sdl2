package main

import (
	"./sdl"
)

func main() {
	sdl.Init(sdl.INIT_VIDEO)
	window := sdl.CreateWindow("Hello World!", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
				800, 600, sdl.WINDOW_SHOWN | sdl.WINDOW_OPENGL)
	renderer := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	
	renderer.SetDrawColor(255, 255, 255, 255)
	renderer.DrawPoint(150, 300)

	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.DrawLine(0, 0, 200, 200)

	points := []sdl.Point { {0, 0}, {100, 300}, {100, 300}, {200, 0} }
	renderer.SetDrawColor(255, 255, 0, 255)
	renderer.DrawLines(&points, 4)

	rect := sdl.Rect { 300, 0, 200, 200 }
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawRect(&rect)

	rects := []sdl.Rect { {400, 400, 100, 100}, {550, 350, 200, 200} }
	renderer.SetDrawColor(0, 255, 255, 255)
	renderer.DrawRects(&rects, 2)

	rect = sdl.Rect { 250, 250, 200, 200 }
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.FillRect(&rect)

	rects = []sdl.Rect { {500, 300, 100, 100}, {200, 300, 200, 200} }
	renderer.SetDrawColor(255, 0, 255, 255)
	renderer.FillRects(&rects, 2)
	
	renderer.Present()

	sdl.Delay(1000)
	sdl.Quit()
}
