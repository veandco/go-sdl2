SDL2 binding for Go
===================
go-sdl2 is SDL2 wrapped for Go users. It enables interoperability between Go and the SDL2 library which is written in C. That means the original SDL2 installation is required for this to work.

Requirements
============
* [SDL2](http://libsdl.org/download-2.0.php)
* [SDL2_mixer (optional)](http://www.libsdl.org/projects/SDL_mixer/)
* [SDL2_image (optional)](http://www.libsdl.org/projects/SDL_image/)
* [SDL2_ttf (optional)](http://www.libsdl.org/projects/SDL_ttf/)

Below is some commands that can be used to install the required packages in
some Linux distributions. Some older versions of the distributions such as
Ubuntu 13.10 may also be used but it may miss an optional package such as
_libsdl2-ttf-dev_ on Ubuntu 13.10's case which is available in Ubuntu 14.04.

On __Ubuntu 14.04 and above__, type:  
`apt-get install libsdl2{,-mixer,-image,-ttf}-dev`

On __Fedora 20 and above__, type:  
`yum install SDL2{,_mixer,_image,_ttf}-devel`

On __Arch Linux__, type:  
`pacman -S sdl2{,_mixer,_image,_ttf}`

For Mac OS X, download the development library dmg files from libsdl.org and drag each framework into /Library/Frameworks

Installation
============
To get the bindings, type:  
`go get -v github.com/veandco/go-sdl2/sdl`  
`go get -v github.com/veandco/go-sdl2/sdl_mixer`  
`go get -v github.com/veandco/go-sdl2/sdl_image`  
`go get -v github.com/veandco/go-sdl2/sdl_ttf`

or type this if you use Bash terminal:  
`go get -v github.com/veandco/go-sdl2/sdl{,_mixer,_image,_ttf}`

__Note__: If you didn't use the previous commands or use 'go install', you will experience long
compilation time because Go doesn't keep the built binaries unless you install them.

Example
=======
	package main

	import "github.com/veandco/go-sdl2/sdl"

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
For now, take a look at http://godoc.org/github.com/veandco/go-sdl2/sdl. A full-featured website will be created once we hit a stable point.

Notes
=====
A standalone Go SDL2 library _is_ being considered (read: figured out). That means users should be able to just go get go-sdl2 and compile it without the original C library. That could mean faster build times, more 'idiomatic' Go code, and hopefully more people interested in using and contributing to go-sdl2!
 
Contributors
============
* [Jacky Boen](https://github.com/jackyb)
* [HardWareGuy](https://github.com/HardWareGuy)
* [akovaski](https://github.com/akovaski)
* [Jeromy Johnson](https://github.com/whyrusleeping)
* [Cai Lei](https://github.com/ccll)
* [krux02](https://github.com/krux02)
* [marcusva](https://github.com/marcusva)
* [Tom Murray](https://github.com/TomMurray)
* [Ian Davis](https://github.com/iand)

License
=======
Go-SDL2 is BSD 3-clause licensed.
