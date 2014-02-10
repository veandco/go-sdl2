SDL2 binding for Go
===================
go-sdl2 is a straightforward Go wrapper for SDL2.  
  
Where makes sense, it'll use methods instead of functions (such as
window.UpdateSurface() or surface.Free() instead of
sdl.UpdateWindowSurface(window) or sdl.FreeSurface(surface)).

Requirements
============
* [SDL2](http://libsdl.org/download-2.0.php)
* [SDL2_mixer (optional)](http://www.libsdl.org/projects/SDL_mixer/)
* [SDL2_image (optional)](http://www.libsdl.org/projects/SDL_image/)

On __Fedora 20__, you can run this command to install all the development
files:  
`yum install SDL2{,_mixer,_image}-devel`

On __Arch Linux__, type:  
`pacman -S sdl2{,mixer,image}`

Installation
============
To get the bindings, type:  
`go get -v github.com/jackyb/go-sdl2/sdl`  
`go get -v github.com/jackyb/go-sdl2/sdl_mixer`  
`go get -v github.com/jackyb/go-sdl2/sdl_image`

or type this if you use Bash terminal:  
`go get -v github.com/jackyb/go-sdl2/sdl{,_mixer,_image}`

Example
=======
	package main

	import "github.com/jackyb/go-sdl2/sdl"

	func main() {
		window := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
				800, 600, sdl.WINDOW_SHOWN)
		surface := window.GetSurface()

		rect := sdl.Rect { 0, 0, 200, 200 }
		surface.FillRect(&rect, 0xffff0000)
		window.UpdateSurface()

		sdl.Delay(1000)
		window.Destroy()
	}


For more complete examples, see inside the _examples_ folder.

Documentation
=============
For now, take a look at http://godoc.org/github.com/jackyb/go-sdl2/sdl.

Contributors
============
* [Jacky Boen](https://github.com/jackyb)
* [HardWareGuy](https://github.com/HardWareGuy)
* [akovaski](https://github.com/akovaski)
* [whyrusleeping](https://github.com/whyrusleeping)
* [ccll](https://github.com/ccll)

License
=======
Go-SDL2 is BSD 3-clause licensed.
