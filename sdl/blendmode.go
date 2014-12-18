package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// BlendMode (https://wiki.libsdl.org/SDL_BlendMode)
type BlendMode uint32

const (
	BLENDMODE_NONE  = C.SDL_BLENDMODE_NONE
	BLENDMODE_BLEND = C.SDL_BLENDMODE_BLEND
	BLENDMODE_ADD   = C.SDL_BLENDMODE_ADD
	BLENDMODE_MOD   = C.SDL_BLENDMODE_MOD
)

func (bm BlendMode) c() C.SDL_BlendMode {
	return C.SDL_BlendMode(C.Uint32(bm))
}

func (bm *BlendMode) cptr() *C.SDL_BlendMode {
	return (*C.SDL_BlendMode)(unsafe.Pointer(bm))
}
