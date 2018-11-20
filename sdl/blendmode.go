package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,6))


#if defined(WARN_OUTDATED)
#pragma message("SDL_BLENDMODE_INVALID is not supported before SDL 2.0.6")
#endif

#define SDL_BLENDMODE_INVALID (0x7FFFFFFF)


#if defined(WARN_OUTDATED)
#pragma message("SDL_BlendOperation is not supported before SDL 2.0.6")
#endif

typedef enum
{
    SDL_BLENDOPERATION_ADD              = 0x1,
    SDL_BLENDOPERATION_SUBTRACT         = 0x2,
    SDL_BLENDOPERATION_REV_SUBTRACT     = 0x3,
    SDL_BLENDOPERATION_MINIMUM          = 0x4,
    SDL_BLENDOPERATION_MAXIMUM          = 0x5
} SDL_BlendOperation;


#if defined(WARN_OUTDATED)
#pragma message("SDL_BlendFactor is not supported before SDL 2.0.6")
#endif

typedef enum
{
    SDL_BLENDFACTOR_ZERO                = 0x1,
    SDL_BLENDFACTOR_ONE                 = 0x2,
    SDL_BLENDFACTOR_SRC_COLOR           = 0x3,
    SDL_BLENDFACTOR_ONE_MINUS_SRC_COLOR = 0x4,
    SDL_BLENDFACTOR_SRC_ALPHA           = 0x5,
    SDL_BLENDFACTOR_ONE_MINUS_SRC_ALPHA = 0x6,
    SDL_BLENDFACTOR_DST_COLOR           = 0x7,
    SDL_BLENDFACTOR_ONE_MINUS_DST_COLOR = 0x8,
    SDL_BLENDFACTOR_DST_ALPHA           = 0x9,
    SDL_BLENDFACTOR_ONE_MINUS_DST_ALPHA = 0xA

} SDL_BlendFactor;


#if defined(WARN_OUTDATED)
#pragma message("SDL_ComposeCustomBlendMode is not supported before SDL 2.0.6")
#endif

SDL_BlendMode SDLCALL SDL_ComposeCustomBlendMode(SDL_BlendFactor srcColorFactor, SDL_BlendFactor dstColorFactor, SDL_BlendOperation colorOperation, SDL_BlendFactor srcAlphaFactor, SDL_BlendFactor dstAlphaFactor, SDL_BlendOperation alphaOperation)
{
	return SDL_BLENDMODE_NONE;
}
#endif
*/
import "C"
import "unsafe"

// BlendMode is an enumeration of blend modes used in Render.Copy() and drawing operations.
// (https://wiki.libsdl.org/SDL_BlendMode)
type BlendMode uint32

const (
	BLENDMODE_NONE    = C.SDL_BLENDMODE_NONE  // no blending
	BLENDMODE_BLEND   = C.SDL_BLENDMODE_BLEND // alpha blending
	BLENDMODE_ADD     = C.SDL_BLENDMODE_ADD   // additive blending
	BLENDMODE_MOD     = C.SDL_BLENDMODE_MOD   // color modulate
	BLENDMODE_INVALID = C.SDL_BLENDMODE_INVALID
)

func (bm BlendMode) c() C.SDL_BlendMode {
	return C.SDL_BlendMode(C.Uint32(bm))
}

func (bm *BlendMode) cptr() *C.SDL_BlendMode {
	return (*C.SDL_BlendMode)(unsafe.Pointer(bm))
}

// BlendOperation is an enumeration of blend operations used when creating a custom blend mode with ComposeCustomBlendMode().
// (https://wiki.libsdl.org/SDL_BlendOperation)
type BlendOperation C.SDL_BlendOperation

const (
	BLENDOPERATION_ADD          = C.SDL_BLENDOPERATION_ADD
	BLENDOPERATION_SUBTRACT     = C.SDL_BLENDOPERATION_SUBTRACT
	BLENDOPERATION_REV_SUBTRACT = C.SDL_BLENDOPERATION_REV_SUBTRACT
	BLENDOPERATION_MINIMUM      = C.SDL_BLENDOPERATION_MINIMUM
	BLENDOPERATION_MAXIMUM      = C.SDL_BLENDOPERATION_MAXIMUM
)

// BlendFactor is an enumeration of blend factors used when creating a custom blend mode with ComposeCustomBlendMode().
// (https://wiki.libsdl.org/SDL_BlendFactor)
type BlendFactor C.SDL_BlendFactor

const (
	BLENDFACTOR_ZERO                = C.SDL_BLENDFACTOR_ZERO                // 0, 0, 0, 0
	BLENDFACTOR_ONE                 = C.SDL_BLENDFACTOR_ONE                 // 1, 1, 1, 1
	BLENDFACTOR_SRC_COLOR           = C.SDL_BLENDFACTOR_SRC_COLOR           // srcR, srcG, srcB, srcA
	BLENDFACTOR_ONE_MINUS_SRC_COLOR = C.SDL_BLENDFACTOR_ONE_MINUS_SRC_COLOR // 1-srcR, 1-srcG, 1-srcB, 1-srcA
	BLENDFACTOR_SRC_ALPHA           = C.SDL_BLENDFACTOR_SRC_ALPHA           // srcA, srcA, srcA, srcA
	BLENDFACTOR_ONE_MINUS_SRC_ALPHA = C.SDL_BLENDFACTOR_ONE_MINUS_SRC_ALPHA // 1-srcA, 1-srcA, 1-srcA, 1-srcA
	BLENDFACTOR_DST_COLOR           = C.SDL_BLENDFACTOR_DST_COLOR           // dstR, dstG, dstB, dstA
	BLENDFACTOR_ONE_MINUS_DST_COLOR = C.SDL_BLENDFACTOR_ONE_MINUS_DST_COLOR // 1-dstR, 1-dstG, 1-dstB, 1-dstA
	BLENDFACTOR_DST_ALPHA           = C.SDL_BLENDFACTOR_DST_ALPHA           // dstA, dstA, dstA, dstA
	BLENDFACTOR_ONE_MINUS_DST_ALPHA = C.SDL_BLENDFACTOR_ONE_MINUS_DST_ALPHA // 1-dstA, 1-dstA, 1-dstA, 1-dstA
)

// ComposeCustomBlendMode creates a custom blend mode, which may or may not be supported by a given renderer
// The result of the blend mode operation will be:
//     dstRGB = dstRGB * dstColorFactor colorOperation srcRGB * srcColorFactor
// and
//     dstA = dstA * dstAlphaFactor alphaOperation srcA * srcAlphaFactor
// (https://wiki.libsdl.org/SDL_ComposeCustomBlendMode)
func ComposeCustomBlendMode(srcColorFactor, dstColorFactor BlendFactor, colorOperation BlendOperation, srcAlphaFactor, dstAlphaFactor BlendFactor, alphaOperation BlendOperation) BlendMode {
	_srcColorFactor := C.SDL_BlendFactor(srcColorFactor)
	_dstColorFactor := C.SDL_BlendFactor(dstColorFactor)
	_colorOperation := C.SDL_BlendOperation(colorOperation)
	_srcAlphaFactor := C.SDL_BlendFactor(srcAlphaFactor)
	_dstAlphaFactor := C.SDL_BlendFactor(dstAlphaFactor)
	_alphaOperation := C.SDL_BlendOperation(alphaOperation)
	return BlendMode(C.SDL_ComposeCustomBlendMode(_srcColorFactor, _dstColorFactor, _colorOperation, _srcAlphaFactor, _dstAlphaFactor, _alphaOperation))
}
