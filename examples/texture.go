// author: Jacky Boen

package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

var winTitle string = "Go-SDL2 Texture"
var winWidth, winHeight int = 800, 600
var imageName string = "test.bmp"

func main() {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var image *sdl.Surface
	var texture *sdl.Texture
	var src, dst sdl.Rect
	var err error

	window, err = sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(1)
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}
	defer renderer.Destroy()

	image = sdl.LoadBMP(imageName)
	if image == nil {
		fmt.Fprintf(os.Stderr, "Failed to load BMP: %s", sdl.GetError())
		os.Exit(3)
	}
	defer image.Free()

	texture, err = renderer.CreateTextureFromSurface(image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(4)
	}
	defer texture.Destroy()

	src = sdl.Rect{0, 0, 512, 512}
	dst = sdl.Rect{100, 50, 512, 512}

	renderer.Clear()
	renderer.Copy(texture, &src, &dst)
	renderer.Present()

	sdl.Delay(2000)
}
