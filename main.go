package main

import (
	"github.com/jackyb/go-sdl2/sdl"
	"fmt"
	"os"
)

var winTitle string = "Hello, Go-SDL2!"
var winWidth, winHeight int = 800, 600
var imageName string = "test.bmp"

func testTexture() {
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

func testRendering() {
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
	renderer.DrawLines(&points, 4)

	rect = sdl.Rect { 300, 0, 200, 200 }
	renderer.SetDrawColor(255, 0, 0, 255)
	renderer.DrawRect(&rect)

	rects = []sdl.Rect { {400, 400, 100, 100}, {550, 350, 200, 200} }
	renderer.SetDrawColor(0, 255, 255, 255)
	renderer.DrawRects(&rects, 2)

	rect = sdl.Rect { 250, 250, 200, 200 }
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.FillRect(&rect)

	rects = []sdl.Rect { {500, 300, 100, 100}, {200, 300, 200, 200} }
	renderer.SetDrawColor(255, 0, 255, 255)
	renderer.FillRects(&rects, 2)

	renderer.Present()
	sdl.Delay(2000)
	renderer.Destroy()
	window.Destroy()
}

func testEvents() {
	var window *sdl.Window
	var renderer *sdl.Renderer
	var event sdl.Event
	var running bool

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

	running = true
	for running {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.MouseMotionEvent:
				fmt.Printf("[%d ms] MouseMotion\tid:%d\tx:%d\ty:%d\txrel:%d\tyrel:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.XRel, t.YRel)
			case *sdl.MouseButtonEvent:
				fmt.Printf("[%d ms] MouseButton\tid:%d\tx:%d\ty:%d\tbutton:%d\tstate:%d\n", t.Timestamp, t.Which, t.X, t.Y, t.Button, t.State)
			case *sdl.MouseWheelEvent:
				fmt.Printf("[%d ms] MouseWheel\tid:%d\tx:%d\ty:%d\n", t.Timestamp, t.Which, t.X, t.Y)
			case *sdl.KeyUpEvent:
				fmt.Printf("[%d ms] Keyboard\tsym:%c\tmodifiers:%d\tstate:%d\trepeat:%d\n", t.Timestamp, t.Keysym.Sym, t.Keysym.Mod, t.State, t.Repeat)
			}
		}
	}

	renderer.Destroy()
	window.Destroy()
}

func main() {
	/*
	testEvents()
	testRendering()
	*/
	testTexture()
}
