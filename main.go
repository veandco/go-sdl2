package main

import (
	"./sdl"
)

func main() {
	sdl.Init(sdl.INIT_VIDEO)
	window := sdl.CreateWindow("Hello World!", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
				800, 600, sdl.WINDOW_SHOWN)
	println(sdl.GetError())
	sdl.Error(sdl.UNSUPPORTED)
	println(sdl.GetError())
	sdl.ClearError()

	println("Window title is", window.GetTitle())
	if sdl.IsScreenSaverEnabled() == true {
		println("Screensaver is enabled")
	} else {
		println("Screensaver is disabled")
	}

}
