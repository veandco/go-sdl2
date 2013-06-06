package main

import (
	"./sdl"
)

func main() {
	sdl.Init(sdl.INIT_VIDEO)
	window := sdl.CreateWindow("Hello World!", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
				800, 600, sdl.WINDOW_SHOWN | sdl.WINDOW_OPENGL)
	renderer := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	println(sdl.GetError())
	sdl.ClearError()

	println("Window title is", window.GetTitle())
	if sdl.IsScreenSaverEnabled() == true {
		println("Screensaver is enabled")
	} else {
		println("Screensaver is disabled")
	}

	renderer.SetDrawColor(255, 0, 0, 255);
	renderer.DrawLine(0, 0, 200, 200)
	renderer.Present()

	sdl.Delay(1000)
}
