package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// Texture (https://wiki.libsdl.org/SDL_Texture)
type Texture C.SDL_Texture
type Renderer C.SDL_Renderer

func (t *Texture) cptr() *C.SDL_Texture {
	return (*C.SDL_Texture)(unsafe.Pointer(t))
}

func (r *Renderer) cptr() *C.SDL_Renderer {
	return (*C.SDL_Renderer)(unsafe.Pointer(r))
}
