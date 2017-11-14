package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// Texture contains an efficient, driver-specific representation of pixel data.
// (https://wiki.libsdl.org/SDL_Texture)
type Texture C.SDL_Texture

// Renderer contains a rendering state.
// (https://wiki.libsdl.org/SDL_Renderer)
type Renderer C.SDL_Renderer

func (t *Texture) cptr() *C.SDL_Texture {
	return (*C.SDL_Texture)(unsafe.Pointer(t))
}

func (r *Renderer) cptr() *C.SDL_Renderer {
	return (*C.SDL_Renderer)(unsafe.Pointer(r))
}
