# Maintaining

This contains personal notes on how I maintain the project on my machine. It is highly subjective and probably more troublesome that standard setup so please adjust this to suit your own preferences!

## Setup

### macOS

On macOS, I use Homebrew mostly on packages that have deep or many dependencies such as `pkg-config`. I prefer to build SDL2 and its related packages myself as they are mostly independent and install it to my own home directory so I don't clutter system directories unnecessarily.

#### Initial directory setup

1. Create directories for containing CLI tools and library source codes.
```
mkdir -p $HOME/Downloads/{CLI,Libraries}
```

2. Create a directory for containing the `go-sdl2` project. In my case, I create a `Projects` directory at my home directory with another directory called `veandco` inside it as it is the entity associated with the `go-sdl2` project.
```
mkdir -p $HOME/Projects/veandco
```

#### Install Go

1. Go to CLI directory.
```
cd $HOME/Downloads/CLI
```

2. Download compressed version of the Go toolchain at https://golang.org/dl.
```
curl -O https://golang.org/dl/go1.16.darwin-amd64.tar.gz # Intel
# or
curl -O https://golang.org/dl/go1.16.darwin-arm64.tar.gz # M1
```

3. Extract the compressed Go toolchain file.
```
rm -r go # if Go has been installed before
tar xf go1.16.darwin-amd64.tar.gz # Intel
# or
tar xf go1.16.darwin-arm64.tar.gz # M1
```

4. Set up PATH in ZSH startup file by adding the following content to `$HOME/.zshrc` so the shell can find Go executables.
```
export GOROOT="$HOME/Downloads/CLI/go"
export GOPATH="$HOME/.go"
export PATH="$GOPATH/bin:$GOROOT/bin:$PATH"
```

5. Open another Terminal tab or window and test that it works.
```
go version
```

6. I also needed to allow Go linker tool to be executed by going to `$GOROOT/pkg/tool/darwin_amd64` for Intel or `$GOROOT/pkg/tool/darwin_arm64` for M1 in the **Finder** app and double-click `link` which would be prohibited to run by macOS the first time. After that, I need to go to **System Preferences** -> **Securiy and Privacy** and allow the program to run at the bottom.

#### Install Homebrew

Go to https://brew.sh and run the installation command shown on the website.

#### Install pkg-config

1. Install it via Homebrew.
```
brew install pkgconfig
```

2. Set up `PKG_CONFIG_PATH` in ZSH startup file at `$HOME/.zshrc` with the following content.
```
export PKG_CONFIG_PATH="$HOME/.local/lib/pkgconfig:$PKG_CONFIG_PATH"
```

#### Install SDL2

1. Go to the libraries directory 
```
cd $HOME/Downloads/Libraries
```

2. Download the source code for SDL2 and optionally its related packages. I also included some dependencies such as mpg123 for MP3 playback support in SDL2_mixer.
```
curl -O https://libsdl.org/releases/SDL2-2.0.14.tar.gz
curl -O https://www.libsdl.org/projects/SDL_mixer/release/SDL2_mixer-2.0.4.tar.gz
curl -O https://www.libsdl.org/projects/SDL_image/release/SDL2_image-2.0.5.tar.gz
curl -O https://www.libsdl.org/projects/SDL_ttf/release/SDL2_ttf-2.0.15.tar.gz
curl -O http://www.ferzkopp.net/Software/SDL2_gfx/SDL2_gfx-1.0.4.tar.gz

# for SDL2_mixer
curl -O https://www.mpg123.de/download/mpg123-1.26.4.tar.bz2 # MP3 support

# for SDL2_ttf
curl -O -L https://download.savannah.gnu.org/releases/freetype/freetype-2.10.4.tar.gz
```

3. Set up `LDFLAGS` and `C_INCLUDE_PATH` in `$HOME/.zshrc` so `configure` can find libraries and header files.
```
export LDFLAGS="-L$HOME/.local/lib $LDFLAGS"
export C_INCLUDE_PATH="$HOME/.local/include:$C_INCLUDE_PATH"
```

4. Extract and build each package by following the template below. Please build the dependencies such as `mpg123` first!
```
tar xf [PACKAGE].[EXTENSION]
mkdir -p [PACKAGE]/build
cd [PACKAGE]/build
../configure --prefix=$HOME/.local
make
make install
cd -
```

For example, for SDL2-2.0.14.tar.gz it would be like this:
```
tar xf SDL2-2.0.14.tar.gz
mkdir -p SDL2-2.0.14/build
cd SDL2-2.0.14/build
../configure --prefix=$HOME/.local
make
make install
cd -
```

5. Check that pkg-config can find them.
```
pkg-config --cflags --libs sdl2 sdl2_mixer sdl2_image sdl2_ttf sdl2_gfx 
```

#### Set up go-sdl2

1. Go to the project parent directory.
```
cd $HOME/Projects/veandco
```

2. Git clone the project from Github. I usually use the SSH protocol instead of the HTTPS protocol. I also use `--recursive` flag as it will also initialize the submodules such as `.go-sdl2-examples`.
```
git clone --recursive git@github.com:veandco/go-sdl2
```

3. Change to the project directory.
```
cd go-sdl2
```

3. Test that everything works by running an example.
```
cd .go-sdl2-examples/examples/render
go run render.go
```

## Development

### Adding new bindings

We try to keep `go-sdl2` buildable with the oldest SDL2 version. In case where a new function is introduced in the actual SDL2 library, we create a dummy function in the related `.go` file. For example, let's see the following function:

```c
// $HOME/.local/include/SDL2/SDL_sensor.h
extern DECLSPEC void SDLCALL SDL_LockSensors(void);
```

The function was introduced in SDL2 2.0.14 so if we try to build `go-sdl2` using older SDL2 such as SDL2 2.0.0, the build would fail. To keep it working, we need to create fail-safe measures in the corresponding `sdl/sensor.go` file.

Almost every `.go` file has the following at the top the file.

```go
package sdl

/*
#include "sdl_wrapper.h"

...

*/
```

In the `...` area is where we implement the fail-safe measures.

1. Create an #ifdef guard to check if our installed SDL2 version is older than the version in which the function is introduced. In this case it is 2.0.14, so this is how it would like:

```go
package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,14))

...

#endif
*/
```

2. Then in the `...` area, we implement a warning section if it doesn't have it already. The warnings are not printed by default when building the `go-sdl2` package but it can enabled by setting `CGO_CPPFLAGS=-DWARN_OUTDATED` environment variable.

```go
package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,14))

#if defined(WARN_OUTDATED)
#pragma message("SDL_LockSensors is not supported before SDL 2.0.14")
#endif

#endif
*/
```

3. Below the warning block, we define our fail-safe functions like so:

```go
package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,14))

#if defined(WARN_OUTDATED)
#pragma message("SDL_LockSensors is not supported before SDL 2.0.14")
#endif

static void SDL_LockSensors(void)
{
        // do nothing
}

#endif
*/
```

4. If we have more functions that need fail-safe, it would look like this:

```go
package sdl
  
/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,14))

#if defined(WARN_OUTDATED)
#pragma message("SDL_LockSensors is not supported before SDL 2.0.14")
#pragma message("SDL_UnlockSensors is not supported before SDL 2.0.14")
#endif

static void SDL_LockSensors(void)
{
        // do nothing
}

static void SDL_UnlockSensors(void)
{
        // do nothing
}

#endif
```