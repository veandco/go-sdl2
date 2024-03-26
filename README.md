# SDL2 binding for Go  [![Build Status](https://github.com/veandco/go-sdl2/actions/workflows/test-build.yml/badge.svg)](https://github.com/veandco/go-sdl2/actions/workflows/test-build.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/veandco/go-sdl2)](https://goreportcard.com/report/github.com/veandco/go-sdl2) [![Reviewed by Hound](https://img.shields.io/badge/Reviewed_by-Hound-8E64B0.svg)](https://houndci.com) [![Financial Contributors on Open Collective](https://opencollective.com/veandco/all/badge.svg?label=financial+contributors)](https://opencollective.com/veandco)

`go-sdl2` is SDL2 wrapped for Go users. It enables interoperability between Go and the SDL2 library which is written in C. That means the original SDL2 installation is required for this to work. Note that the first build may take several minutes on machines that are not powerful such as Raspberry Pi.

NOTE: For the latest versions of SDL2, please use the `master` branch!

# Table of Contents
* [Documentation](#documentation)
* [Examples](#examples)
* [Requirements](#requirements)
* [Installation](#installation)
* [Cross-compiling](#cross-compiling)
* [FAQ](#faq)
* [License](#license)


# Documentation
* [GoDoc documentation for go-sdl2](https://godoc.org/github.com/veandco/go-sdl2)
* [Original SDL2 wiki](https://wiki.libsdl.org)

# Getting Started

If you haven't created a Go module for your program, you can do it by making a directory called, for example, `app` and running the following command inside it:
```
go mod init app
```

After that you can start writing code that uses `go-sdl2`, for example, like the following:
```go
// main.go
package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	sdl.Init(sdl.INIT_EVERYTHING)
}
```

Then to tell Go to fetch dependencies, you can run:
```
go mod tidy
```

and then the program can be built using:
```
go build
```

After that, you can execute your program:

```
./app
```

# Requirements
* [Go v1.13+](https://go.dev/dl/)
* [SDL2](https://github.com/libsdl-org/SDL/releases)
* [SDL2_image (optional)](https://github.com/libsdl-org/SDL_image/releases)
* [SDL2_mixer (optional)](https://github.com/libsdl-org/SDL_mixer/releases)
* [SDL2_ttf (optional)](https://github.com/libsdl-org/SDL_ttf/releases)
* [SDL2_gfx (optional)](http://www.ferzkopp.net/wordpress/2016/01/02/sdl_gfx-sdl2_gfx/)

On __Ubuntu 22.04 and above__, type:\
`apt install libsdl2{,-image,-mixer,-ttf,-gfx}-dev`

On __Fedora 36 and above__, type:\
`dnf install SDL2{,_image,_mixer,_ttf,_gfx}-devel`

On __Arch Linux__, type:\
`pacman -S sdl2{,_image,_mixer,_ttf,_gfx}`

On __Gentoo__, type:\
`emerge -av libsdl2 sdl2-{image,mixer,ttf,gfx}`

On __macOS__, install SDL2 via [Homebrew](http://brew.sh) like so:\
`brew install sdl2{,_image,_mixer,_ttf,_gfx} pkg-config`

On __Windows__,
1. Install mingw-w64 from [Mingw-builds](https://github.com/niXman/mingw-builds-binaries/releases). A 7z archive extractor software might be needed which can be downloaded [here](https://www.7-zip.org/download.html). In this example, we extract the content, which is `mingw64`, into `C:\`.
2. Download and install `SDL2-devel-[version]-mingw.zip` files from https://github.com/libsdl-org/SDL/releases.
    * Extract the SDL2 folder from the archive using a tool like [7zip](http://7-zip.org)
    * Inside the extracted SDL2 folder, copy the `i686-w64-mingw32` and/or `x86_64-w64-mingw32` into mingw64 folder e.g. `C:\mingw64`
3. Setup `Path` environment variable
    * Put mingw-w64 binaries location into system `Path` environment variable (e.g. `C:\mingw64\bin`)
4. Close and open terminal again so the new `Path` environment variable takes effect. Now we should be able to run `go build` inside the project directory.
5. Download and install SDL2 runtime libraries from https://github.com/libsdl-org/SDL/releases. Extract and copy the `.dll` file into the project directory. After that, the program should become runnable.
6. (Optional) You can repeat __Step 2__ for [SDL_image](https://github.com/libsdl-org/SDL_image/releases), [SDL_mixer](https://github.com/libsdl-org/SDL_mixer/releases), [SDL_ttf](https://github.com/libsdl-org/SDL_ttf/releases)


# Installation
To get the bindings, type:\
`go get -v github.com/veandco/go-sdl2/sdl`\
`go get -v github.com/veandco/go-sdl2/img`\
`go get -v github.com/veandco/go-sdl2/mix`\
`go get -v github.com/veandco/go-sdl2/ttf`\
`go get -v github.com/veandco/go-sdl2/gfx`

or type this if you use Bash terminal:\
`go get -v github.com/veandco/go-sdl2/{sdl,img,mix,ttf}`

Due to `go-sdl2` being under active development, a lot of breaking changes are going to happen during v0.x. With [versioning system](https://github.com/golang/proposal/blob/master/design/24301-versioned-go.md) coming to Go soon, we'll make use of semantic versioning to ensure stability in the future.

# Static compilation

Since v0.3.0, it is possible to build statically against included libraries in `.go-sdl2-libs`. To build statically, run:

`CGO_ENABLED=1 CC=gcc GOOS=linux GOARCH=amd64 go build -tags static -ldflags "-s -w"`

You can also cross-compile to another OS. For example, to Windows:

`CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows GOARCH=amd64 go build -tags static -ldflags "-s -w"`

On Windows, if you would like to hide the Command Prompt window when running the statically-compiled program, you could append `-H windowsgui` to the `-ldflags` value.

For the list of OS and architecture, you can see inside the [.go-sdl2-libs](https://github.com/veandco/go-sdl2/tree/master/.go-sdl2-libs) directory.

NOTE: If you're using the new Go Module system, you will need to refer to the master branch for now by running:

`go get -v github.com/veandco/go-sdl2/sdl@master`

Before building the program.

# Cross-compiling
### Linux to Windows
1. Install MinGW toolchain.
   * On **Arch Linux**, it's simply `pacman -S mingw-w64`.
2. Download the SDL2 development package for MinGW [here](http://libsdl.org/download-2.0.php) (and the others like *SDL_image*, *SDL_mixer*, etc.. [here](https://www.libsdl.org/projects/) if you use them).
3. Extract the SDL2 development package and copy the `x86_64-w64-mingw32` folder inside recursively to the system's MinGW `x86_64-w64-mingw32` folder. You may also do the same for the `i686-w64-mingw32` folder.
   * On **Arch Linux**, it's `cp -r x86_64-w64-mingw32 /usr`.
4. Now you can start cross-compiling your Go program by running `env CGO_ENABLED="1" CC="/usr/bin/x86_64-w64-mingw32-gcc" GOOS="windows" CGO_LDFLAGS="-lmingw32 -lSDL2" CGO_CFLAGS="-D_REENTRANT" go build -x main.go`. You can change some of the parameters if you'd like to. In this example, it should produce a `main.exe` executable file.
5. Before running the program, you need to put `SDL2.dll` from the [SDL2 runtime package](http://libsdl.org/download-2.0.php) (For others like *SDL_image*, *SDL_mixer*, etc.., look for them [here](https://www.libsdl.org/projects/)) for Windows in the same folder as your executable.
6. Now you should be able to run the program using Wine or Windows!

### macOS to Windows
1. Install [Homebrew](https://brew.sh)
2. Install MinGW through Homebrew via `brew install mingw-w64`
3. Download the SDL2 development package for MinGW [here](http://libsdl.org/download-2.0.php) (and the others like *SDL_image*, *SDL_mixer*, etc.. [here](https://www.libsdl.org/projects/) if you use them).
4. Extract the SDL2 development package and copy the `x86_64-w64-mingw` folder inside recursively to the system's MinGW `x86_64-w64-mingw32 folder`. You may also do the same for the `i686-w64-mingw32` folder. The path to MinGW may be slightly different but the command should look something like `cp -r x86_64-w64-mingw32 /usr/local/Cellar/mingw-w64/5.0.3/toolchain-x86_64`.
5. Now you can start cross-compiling your Go program by running `env CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc GOOS=windows CGO_LDFLAGS="-L/usr/local/Cellar/mingw-w64/5.0.3/toolchain-x86_64/x86_64-w64-mingw32/lib -lSDL2" CGO_CFLAGS="-I/usr/local/Cellar/mingw-w64/5.0.3/toolchain-x86_64/x86_64-w64-mingw32/include -D_REENTRANT" go build -x main.go`. You can change some of the parameters if you'd like to. In this example, it should produce a `main.exe` executable file.
6. Before running the program, you need to put `SDL2.dll` from the [SDL2 runtime package](http://libsdl.org/download-2.0.php) (For others like *SDL_image*, *SDL_mixer*, etc.., look for them [here](https://www.libsdl.org/projects/)) for Windows in the same folder as your executable.
7. Now you should be able to run the program using Wine or Windows!

### Linux to macOS
1. Install macOS toolchain via [osxcross](https://github.com/tpoechtrager/osxcross)
2. Run the following build command (replace the values in parentheses):
```
CGO_ENABLED=1 CC=[path-to-osxcross]/target/bin/[arch]-apple-darwin[version]-clang GOOS=darwin GOARCH=[arch] go build -tags static -ldflags "-s -w" -a
```

# Examples

NOTE: The following example is for the `master` branch. Please check the [README](https://github.com/veandco/go-sdl2/tree/v0.4.x?tab=readme-ov-file#examples) of `v0.4.x` for the stable version.

```go
package main

import "github.com/veandco/go-sdl2/sdl"

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}
	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 200, 200}
	colour := sdl.Color{R: 255, G: 0, B: 255, A: 255} // purple
	pixel := sdl.MapRGBA(surface.Format, colour.R, colour.G, colour.B, colour.A)
	surface.FillRect(&rect, pixel)
	window.UpdateSurface()

	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case sdl.QuitEvent: // NOTE: Please use `*sdl.QuitEvent` for `v0.4.x` (current version).
				println("Quit")
				running = false
				break
			}
		}

        sdl.Delay(33)
	}
}
```

There are two ways a game might be running: one that updates on user input using `sdl.WaitEvent()` and one that updates regardless of user input using `sdl.PollEvent()`. You can check the examples of those two loops [here](https://github.com/veandco/go-sdl2-examples/tree/master/examples/gameloop).

For more runnable examples, see https://github.com/veandco/go-sdl2-examples. You can run any of the `.go` files with `go run`.

# FAQ
__Why does the program not run on Windows?__
Try putting the [runtime libraries](http://libsdl.org/download-2.0.php) (e.g. `SDL2.dll` and friends) in the same folder as your program.

__Why does my program crash randomly or hang?__
Putting `runtime.LockOSThread()` at the start of your main() usually solves the problem (see [SDL2 FAQ](https://wiki.libsdl.org/FAQDevelopment) about multi-threading).

UPDATE: Recent update added a call queue system where you can put thread-sensitive code and have it called synchronously on the same OS thread. See the `render_queue` or `render_goroutines` examples from https://github.com/veandco/go-sdl2-examples to see how it works.

__Why can't SDL_mixer seem to play MP3 audio file?__
Your installed SDL_mixer probably doesn't support MP3 file.

On __macOS__, this is easy to correct. First remove the faulty mixer: `brew remove sdl2_mixer`, then reinstall it with the MP3 option: `brew install sdl2_mixer --with-flac --with-fluid-synth --with-libmikmod --with-libmodplug --with-smpeg2`. If necessary, check which options you can enable with `brew info sdl2_mixer`. You could also try installing sdl2\_mixer with mpg123 by running `brew install sdl2_mixer --with-mpg123`.

On __Other Operating Systems__, you will need to compile smpeg and SDL_mixer from source with the MP3 option enabled. You can find smpeg in the `external` directory of SDL_mixer. Refer to issue [#148](https://github.com/veandco/go-sdl2/issues/148) for instructions.

_Note that there seems to be a problem with SDL_mixer 2.0.2 so you can also try to revert back to 2.0.1 and see if it solves your problem_

__Does go-sdl2 support compiling on mobile platforms like Android and iOS?__
For Android, see https://github.com/veandco/go-sdl2-examples/tree/master/examples/android.

There is currently no support for iOS yet.

__Why does my window not immediately render after creation?__
It appears the rendering subsystem needs some time to be able to present the drawn pixels. This can be workaround by adding delay using `sdl.Delay()` or put the rendering code inside a draw loop.

## Contributors

### Code Contributors

This project exists thanks to all the people who contribute. [[Contribute](CONTRIBUTING.md)].
<a href="https://github.com/veandco/go-sdl2/graphs/contributors"><img src="https://opencollective.com/veandco/contributors.svg?width=890&button=false" /></a>

### Financial Contributors

Become a financial contributor and help us sustain our community. [[Contribute](https://opencollective.com/veandco/contribute)]

#### Individuals

<a href="https://opencollective.com/veandco"><img src="https://opencollective.com/veandco/individuals.svg?width=890"></a>

#### Organizations

Support this project with your organization. Your logo will show up here with a link to your website. [[Contribute](https://opencollective.com/veandco/contribute)]

<a href="https://opencollective.com/veandco/organization/0/website"><img src="https://opencollective.com/veandco/organization/0/avatar.svg"></a>
<a href="https://opencollective.com/veandco/organization/1/website"><img src="https://opencollective.com/veandco/organization/1/avatar.svg"></a>
<a href="https://opencollective.com/veandco/organization/2/website"><img src="https://opencollective.com/veandco/organization/2/avatar.svg"></a>
<a href="https://opencollective.com/veandco/organization/3/website"><img src="https://opencollective.com/veandco/organization/3/avatar.svg"></a>
<a href="https://opencollective.com/veandco/organization/4/website"><img src="https://opencollective.com/veandco/organization/4/avatar.svg"></a>
<a href="https://opencollective.com/veandco/organization/5/website"><img src="https://opencollective.com/veandco/organization/5/avatar.svg"></a>
<a href="https://opencollective.com/veandco/organization/6/website"><img src="https://opencollective.com/veandco/organization/6/avatar.svg"></a>
<a href="https://opencollective.com/veandco/organization/7/website"><img src="https://opencollective.com/veandco/organization/7/avatar.svg"></a>
<a href="https://opencollective.com/veandco/organization/8/website"><img src="https://opencollective.com/veandco/organization/8/avatar.svg"></a>
<a href="https://opencollective.com/veandco/organization/9/website"><img src="https://opencollective.com/veandco/organization/9/avatar.svg"></a>

# License
Go-SDL2 is BSD 3-clause licensed.
