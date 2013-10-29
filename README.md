SDL2 binding for Go
===================
go-sdl2 is a straightforward Go wrapper for SDL2.

Requirements
============
* [SDL2](http://libsdl.org/download-2.0.php)
* [SDL2_mixer (optional)](http://www.libsdl.org/projects/SDL_mixer/)
* [SDL2_image (optional)](http://www.libsdl.org/projects/SDL_image/)

On _Fedora/Linux_, you can type this command to install SDL2 and SDL2_image
development files (unfortunately, no SDL2_mixer yet):  
`yum install SDL2-devel SDL2_image-devel`

Installation
============
To get just the SDL2 binding, type:  
`go get -v github.com/jackyb/go-sdl2/sdl`

If you wish to get SDL2_mixer and SDL2_image bindings as well, type:  
`go get -v github.com/jackyb/go-sdl2/sdl_mixer github.com/jackyb/go-sdl2/sdl_image`

If you use Bash, you can type this to get the whole package:  
`go get -v github.com/jackyb/go-sdl2/sdl{,_mixer,_image}`

Documentation
=============
For now, take a look at http://godoc.org/github.com/jackyb/go-sdl2/sdl.

Contributors
============
* [Jacky Boen](https://github.com/jackyb)
* [HardWareGuy](https://github.com/HardWareGuy)

License
=======
Go-SDL2 is BSD 3-clause licensed.
