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

On __Windows__, you can either install SDL2 via [Msys2](https://msys2.github.io) like so:
`pacman -S mingw-w64-x86_64-gcc mingw-w64-x86_64-SDL2{,_mixer,_image,_ttf}`

or use [MinGW](http://mingw.org) and download pre-compiled development libraries from http://libsdl.org/download-2.0.php. In this case, you need to put files inside `i686-w64-mingw32` or `x86_64-w64-mingw32` into MinGW's `mingw32` directory e.g. `C:\MinGW\mingw32`. Assuming you have setup MinGW correctly (install base system so it has `gcc` and friends, set PATH environment variables), you should be able to build go-sdl2. If you have 64-bit system, you can still build by modifying certain environment flags before building such as `CGO_ENABLED=1 GOARCH=386 go build`.


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

UPDATE: Recent update added a call queue system where you can put thread-sensitive code and have it called synchronously on the same OS thread. See the `render_queue` or `render_goroutines` examples to see how it works.

__Why can't SDL_mixer seem to play MP3 audio file?__  
Your installed SDL_mixer probably doesn't support MP3 file. You will need to compile smpeg and SDL_mixer from source with the MP3 option enabled. You can find smpeg in the `external` directory of SDL_mixer. Refer to issue [#148](https://github.com/veandco/go-sdl2/issues/148) for instructions.

__Does go-sdl2 support compiling on mobile platforms like Android and iOS?__  
For Android, see https://github.com/gen2brain/go-sdl2-android-example.

There is currently no support for iOS yet.

__How do I contribute?__  
You can contribute by a lot of ways from improving README, fixing typos, coding style, specific bugs, performance optimizations. However, it is preferred that you break up your commits to single logical change using `git add -p` so it is easier to review the patch. The larger the change, the more necessary it is for the commit to be broken up to tiny little pieces. If your change is large but consistent throughout (e.g. fixing a specific coding style that happens on almost every file), that can be counted as single logical change.

You can generally start by forking the repository, and then sending pull requests. But unfortunately this is a Go project, and the absolute import statements make forking a bit more complicated. [Here](http://blog.campoy.cat/2014/03/github-and-go-forking-pull-requests-and.html) are some instructions, how you can work with that. Generally pull requests are very welcome.

Last but not least, we're starting to use commit messages that looks like this: `sdl: fixed some typos in render.go` or `examples: render_goroutine: fixed a dereferenced nil pointer` where it starts with folder hierarchy. It's not something strictly required but we would prefer it to be followed.

__Will there be Go port of SDL2?__  
Due to the way Go is going with C interopability, this might not be possible or even make sense. Another way that we're currently thinking is to create a Rust port instead as it has nicer compatibility with C.

License
=======
Go-SDL2 is BSD 3-clause licensed.

Contributors
============
Here's the list of contributors with their respective Github URLs.
* [Jacky Boen](https://github.com/jackyb)
* [HardWareGuy](https://github.com/HardWareGuy)
* [akovaski](https://github.com/akovaski)
* [Jeromy Johnson](https://github.com/whyrusleeping)
* [Cai Lei](https://github.com/ccll)
* [Arne Döring](https://github.com/krux02)
* [Marcus von Appen](https://github.com/marcusva)
* [Tom Murray](https://github.com/TomMurray)
* [Ian Davis](https://github.com/iand)
* [hschendel](https://github.com/hschendel)
* [Ingo Oeser](https://github.com/nightlyone)
* [nlordell](https://github.com/nlordell)
* [Ben Davies](https://github.com/JalfResi)
* [Bastien Dejean](https://github.com/baskerville)
* [Pirmin Tapken](https://github.com/PirminTapken)
* [Robert Lillack](https://github.com/roblillack)
* [Marcell Jusztin](https://github.com/morcmarc)
* [Stan Schwertly](https://github.com/Stantheman)
* [Michael Vetter](https://github.com/jubalh)
* [Tom Fogal](https://github.com/tfogal)
* [Philipp Meinen](https://github.com/PhiCode)
* [Thomas McGrew](https://github.com/mcgrew)
* [Geoff Catlin](https://github.com/gcatlin)
* [Schobers](https://github.com/Schobers)
* [Jan Tuitman](https://github.com/jantuitman)
* [Nick Powell](https://github.com/THUNDERGROOVE)
* [Steven R. Wiley](https://github.com/srwiley)
* [Franco Lazzarino](https://github.com/flazz)
* [Jason Alan Palmer](https://github.com/jalan)
* [Seuk Won Kang](https://github.com/kasworld)
* [Brandon Mulcahy](https://github.com/jangler)
* [Tim Anema](https://github.com/tanema)
* [Tyler Compton](https://github.com/velovix)
* [Nicolas Hess](https://github.com/n0dev)
* [Stephen Noonan](https://github.com/tycho)
* [Guilherme Freitas Nemeth](https://github.com/glhrmfrts)
* [Charney Kaye](https://github.com/charneykaye)
* [Lars Scheme](https://github.com/gonutz)
* [Emil Laine](https://github.com/emlai)
* [Andreas T. Jonsson](https://github.com/andreas-jonsson)
* [Milan Nikolic](https://github.com/gen2brain)
* [Mike Gerow](https://github.com/gerow)
* [Lilis Iskandar](https://github.com/veeableful)

_if anyone is missing, let me know!.. or you can add yourself in :)_
