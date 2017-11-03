package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// BlendMode is an enumeration of blend modes used in Render.Copy() and drawing operations.
// (https://wiki.libsdl.org/SDL_BlendMode)
type BlendMode uint32

// BlendMode is an enumeration of blend modes used in Render.Copy() and drawing operations.
// (https://wiki.libsdl.org/SDL_BlendMode)
const (
	BLENDMODE_NONE  = C.SDL_BLENDMODE_NONE  // no blending
	BLENDMODE_BLEND = C.SDL_BLENDMODE_BLEND // alpha blending
	BLENDMODE_ADD   = C.SDL_BLENDMODE_ADD   // additive blending
	BLENDMODE_MOD   = C.SDL_BLENDMODE_MOD   // color modulate
)

func (bm BlendMode) c() C.SDL_BlendMode {
	return C.SDL_BlendMode(C.Uint32(bm))
}

func (bm *BlendMode) cptr() *C.SDL_BlendMode {
	return (*C.SDL_BlendMode)(unsafe.Pointer(bm))
}
