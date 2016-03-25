SDL2 binding for Go [![Build Status](https://travis-ci.org/veandco/go-sdl2.svg?branch=master)](https://travis-ci.org/veandco/go-sdl2)
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
_Note: Ubuntu 14.04 currently has broken header file in the SDL2 package that disables people from compiling against it. It will be needed to either patch the header file or install SDL2 from source._

On __Fedora 20 and above__, type:  
`yum install SDL2{,_mixer,_image,_ttf}-devel`

On __Arch Linux__, type:  
`pacman -S sdl2{,_mixer,_image,_ttf}`

On __Mac OS X__, install SDL2 via [Homebrew](http://brew.sh) like so:
`brew install sdl2{,_image,_ttf,_mixer}`

On __Windows__, install SDL2 via [Msys2](https://msys2.github.io) like so:
`pacman -S mingw-w64-x86_64-gcc mingw-w64-x86_64-SDL2{,_mixer,_image,_ttf}`

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
```go
package main

import "github.com/veandco/go-sdl2/sdl"

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	rect := sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()

	sdl.Delay(1000)
	sdl.Quit()
}
```



For more complete examples, see inside the _examples_ folder. Run any of the .go files with `go run`.

FAQ
===
__Why does my program crash randomly?__  
Putting `runtime.LockOSThread()` at the start of your main() usually solves the problem. We will reimplement parts of go-sdl2 in different way so this doesn't happen in the future.

__Why can't SDL_mixer seem to play MP3 audio file?__  
Your installed SDL_mixer probably doesn't support MP3 file. You will need to compile smpeg and SDL_mixer from source with the MP3 option enabled. You can find smpeg in the `external` directory of SDL_mixer. Refer to issue [#148](https://github.com/veandco/go-sdl2/issues/148) for instructions.

__Does go-sdl2 support compiling on mobile platforms like Android and iOS?__  
Not yet

__Will there be Go port of SDL2?__  
There's some work on it but no plan to open-source it yet simply because it's a mess right now :)

__How do I contribute?__  
Generally by forking the repository, and then sending pull requests. But unfortunately this is a go project, and the absolute import statements make forking a bit more complicated. [Here](http://blog.campoy.cat/2014/03/github-and-go-forking-pull-requests-and.html) are some instructions, how you can work with that. Generally pull requests are very welcome.

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
* [hschendel](https://github.com/hschendel)
* [Bastien Dejean](https://github.com/baskerville)
* [Pirmin Fix](https://github.com/PirminFix)
* [Robert Lillack](https://github.com/roblillack)
* [tfogal](https://github.com/tfogal)
* [Philipp Meinen](https://github.com/PhiCode)
* [Thomas McGrew](https://github.com/mcgrew)
* [Geoff Catlin](https://github.com/gcatlin)
* [Jason Alan Palmer](https://github.com/jalan)
* [Seuk Won Kang](https://github.com/kasworld)
* [Brandon Mulcahy](https://github.com/jangler)

_if anyone is missing, let me know!_

License
=======
Go-SDL2 is BSD 3-clause licensed.
