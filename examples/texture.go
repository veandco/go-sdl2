// author: Jacky Boen

package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
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

	image = sdl.LoadBMP(imageName)
	if image == nil {
		fmt.Fprintf(os.Stderr, "Failed to load BMP: %s", sdl.GetError())
		os.Exit(3);
	}

	texture = renderer.CreateTextureFromSurface(image)
	if texture == nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", sdl.GetError());
		os.Exit(4);
	}

	src = sdl.Rect { 0, 0, 512, 512 }
	dst = sdl.Rect { 100, 50, 512, 512 }

	renderer.Clear()
	renderer.Copy(texture, &src, &dst)
	renderer.Present()

	sdl.Delay(2000)

	image.Free()
	texture.Destroy()
	renderer.Destroy()
	window.Destroy()
}
