package sdl

/*
#include "sdl_wrapper.h"
*/
import "C"
import "unsafe"
import "reflect"

const (
	SWSURFACE = C.SDL_SWSURFACE
	PREALLOC  = C.SDL_PREALLOC
	RLEACCEL  = C.SDL_RLEACCEL
	DONTFREE  = C.SDL_DONTFREE
)

// Surface (https://wiki.libsdl.org/SDL_Surface)
type Surface struct {
	Flags    uint32
	Format   *PixelFormat
	W        int32
	H        int32
	Pitch    int32
	pixels   unsafe.Pointer // use Pixels() for access
	UserData unsafe.Pointer
	Locked   int32
	LockData unsafe.Pointer
	ClipRect Rect
	_        unsafe.Pointer // private c field 'map'
	RefCount int32
}
type cSurface C.SDL_Surface

func (surface *Surface) cptr() *C.SDL_Surface {
	return (*C.SDL_Surface)(unsafe.Pointer(surface))
}

// Surface (https://wiki.libsdl.org/SDL_MustLock)
func (surface *Surface) MustLock() bool {
	return (surface.Flags & RLEACCEL) != 0
}

// CreateRGBSurface (https://wiki.libsdl.org/SDL_CreateRGBSurface)
func CreateRGBSurface(flags uint32, width, height, depth int32, Rmask, Gmask, Bmask, Amask uint32) (*Surface, error) {
	surface := (*Surface)(unsafe.Pointer(C.SDL_CreateRGBSurface(
		C.Uint32(flags),
		C.int(width),
		C.int(height),
		C.int(depth),
		C.Uint32(Rmask),
		C.Uint32(Gmask),
		C.Uint32(Bmask),
		C.Uint32(Amask))))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// CreateRGBSurfaceFrom (https://wiki.libsdl.org/SDL_CreateRGBSurfaceFrom)
func CreateRGBSurfaceFrom(pixels unsafe.Pointer, width, height, depth, pitch int, Rmask, Gmask, Bmask, Amask uint32) (*Surface, error) {
	surface := (*Surface)(unsafe.Pointer(C.SDL_CreateRGBSurfaceFrom(
		pixels,
		C.int(width),
		C.int(height),
		C.int(depth),
		C.int(pitch),
		C.Uint32(Rmask),
		C.Uint32(Gmask),
		C.Uint32(Bmask),
		C.Uint32(Amask))))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// Surface (https://wiki.libsdl.org/SDL_FreeSurface)
func (surface *Surface) Free() {
	C.SDL_FreeSurface(surface.cptr())
}

// Surface (https://wiki.libsdl.org/SDL_SetSurfacePalette)
func (surface *Surface) SetPalette(palette *Palette) error {
	if C.SDL_SetSurfacePalette(surface.cptr(), palette.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_LockSurface)
func (surface *Surface) Lock() error {
	if C.SDL_LockSurface(surface.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_UnlockSurface)
func (surface *Surface) Unlock() {
	C.SDL_UnlockSurface(surface.cptr())
}

// LoadBMP_RW (https://wiki.libsdl.org/SDL_LoadBMP_RW)
func LoadBMP_RW(src *RWops, freeSrc int) (*Surface, error) {
	surface := (*Surface)(unsafe.Pointer(C.SDL_LoadBMP_RW(src.cptr(), C.int(freeSrc))))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// LoadBMP (https://wiki.libsdl.org/SDL_LoadBMP)
func LoadBMP(file string) (*Surface, error) {
	return LoadBMP_RW(RWFromFile(file, "rb"), 1)
}

// Surface (https://wiki.libsdl.org/SDL_SaveBMP_RW)
func (surface *Surface) SaveBMP_RW(dst *RWops, freeDst int) error {
	if C.SDL_SaveBMP_RW(surface.cptr(), dst.cptr(), C.int(freeDst)) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_SaveBMP)
func (surface *Surface) SaveBMP(file string) error {
	return surface.SaveBMP_RW(RWFromFile(file, "wb"), 1)
}

// Surface (https://wiki.libsdl.org/SDL_SetSurfaceRLE)
func (surface *Surface) SetRLE(flag int) error {
	if C.SDL_SetSurfaceRLE(surface.cptr(), C.int(flag)) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_SetColorKey)
func (surface *Surface) SetColorKey(flag int, key uint32) error {
	if C.SDL_SetColorKey(surface.cptr(), C.int(flag), C.Uint32(key)) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_GetColorKey)
func (surface *Surface) GetColorKey() (key uint32, err error) {
	_key := (*C.Uint32)(unsafe.Pointer(&key))
	if C.SDL_GetColorKey(surface.cptr(), _key) != 0 {
		return key, GetError()
	}
	return key, nil
}

// Surface (https://wiki.libsdl.org/SDL_SetSurfaceColorMod)
func (surface *Surface) SetColorMod(r, g, b uint8) error {
	if C.SDL_SetSurfaceColorMod(surface.cptr(), C.Uint8(r), C.Uint8(g), C.Uint8(b)) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_GetSurfaceColorMod)
func (surface *Surface) GetColorMod() (r, g, b uint8, err error) {
	_r := (*C.Uint8)(unsafe.Pointer(&r))
	_g := (*C.Uint8)(unsafe.Pointer(&g))
	_b := (*C.Uint8)(unsafe.Pointer(&b))
	if C.SDL_GetSurfaceColorMod(surface.cptr(), _r, _g, _b) != 0 {
		return r, g, b, GetError()
	}
	return r, g, b, nil
}

// Surface (https://wiki.libsdl.org/SDL_SetSurfaceAlphaMod)
func (surface *Surface) SetAlphaMod(alpha uint8) error {
	if C.SDL_SetSurfaceAlphaMod(surface.cptr(), C.Uint8(alpha)) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_GetSurfaceAlphaMod)
func (surface *Surface) GetAlphaMod() (alpha uint8, err error) {
	_alpha := (*C.Uint8)(unsafe.Pointer(&alpha))
	if C.SDL_GetSurfaceAlphaMod(surface.cptr(), _alpha) != 0 {
		return alpha, GetError()
	}
	return alpha, nil
}

// Surface (https://wiki.libsdl.org/SDL_SetSurfaceBlendMode)
func (surface *Surface) SetBlendMode(bm BlendMode) error {
	if C.SDL_SetSurfaceBlendMode(surface.cptr(), bm.c()) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_GetSurfaceBlendMode)
func (surface *Surface) GetBlendMode() (bm BlendMode, err error) {
	if C.SDL_GetSurfaceBlendMode(surface.cptr(), bm.cptr()) != 0 {
		return bm, GetError()
	}
	return bm, nil
}

// Surface (https://wiki.libsdl.org/SDL_SetClipRect)
func (surface *Surface) SetClipRect(rect *Rect) bool {
	return C.SDL_SetClipRect(surface.cptr(), rect.cptr()) > 0
}

// Surface (https://wiki.libsdl.org/SDL_GetClipRect)
func (surface *Surface) GetClipRect(rect *Rect) {
	C.SDL_GetClipRect(surface.cptr(), rect.cptr())
}

// Surface (https://wiki.libsdl.org/SDL_ConvertSurface)
func (surface *Surface) Convert(fmt *PixelFormat, flags uint32) (*Surface, error) {
	_surface := (*Surface)(unsafe.Pointer(C.SDL_ConvertSurface(surface.cptr(), fmt.cptr(), C.Uint32(flags))))
	if _surface == nil {
		return nil, GetError()
	}
	return _surface, nil
}

// Surface (https://wiki.libsdl.org/SDL_ConvertSurfaceFormat)
func (surface *Surface) ConvertFormat(pixelFormat uint32, flags uint32) (*Surface, error) {
	_surface := (*Surface)(unsafe.Pointer(C.SDL_ConvertSurfaceFormat(surface.cptr(), C.Uint32(pixelFormat), C.Uint32(flags))))
	if _surface == nil {
		return nil, GetError()
	}
	return _surface, nil
}

// ConvertPixels (https://wiki.libsdl.org/SDL_ConvertPixels)
func ConvertPixels(width, height int, srcFormat uint32, src unsafe.Pointer, srcPitch int,
	dstFormat uint32, dst unsafe.Pointer, dstPitch int) error {
	if C.SDL_ConvertPixels(C.int(width), C.int(height), C.Uint32(srcFormat), src, C.int(srcPitch), C.Uint32(dstFormat), dst, C.int(dstPitch)) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_FillRect)
func (surface *Surface) FillRect(rect *Rect, color uint32) error {
	if C.SDL_FillRect(surface.cptr(), rect.cptr(), C.Uint32(color)) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_FillRects)
func (surface *Surface) FillRects(rects []Rect, color uint32) error {
	if C.SDL_FillRects(surface.cptr(), rects[0].cptr(), C.int(len(rects)), C.Uint32(color)) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_BlitSurface)
func (src *Surface) Blit(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	if C.SDL_BlitSurface(src.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_BlitScaled)
func (src *Surface) BlitScaled(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	if C.SDL_BlitScaled(src.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_UpperBlit)
func (src *Surface) UpperBlit(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	if C.SDL_UpperBlit(src.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_LowerBlit)
func (src *Surface) LowerBlit(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	if C.SDL_LowerBlit(src.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_SoftStretch)
func (src *Surface) SoftStretch(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	if C.SDL_SoftStretch(src.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_UpperBlitScaled)
func (src *Surface) UpperBlitScaled(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	if C.SDL_UpperBlitScaled(src.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Surface (https://wiki.libsdl.org/SDL_LowerBlitScaled)
func (src *Surface) LowerBlitScaled(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	if C.SDL_LowerBlitScaled(src.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

func (surface *Surface) PixelNum() int {
	return int(surface.W * surface.H)
}

func (surface *Surface) BytesPerPixel() int {
	return int(surface.Format.BytesPerPixel)
}

func (surface *Surface) Pixels() []byte {
	var b []byte
	length := int(surface.W*surface.H) * int(surface.Format.BytesPerPixel)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(surface.pixels)
	return b
}

func (surface *Surface) Data() unsafe.Pointer {
	return surface.pixels
}
