package sdl

/*
#include <SDL2/SDL.h>
*/
import "C"
import "unsafe"
import "reflect"

const (
	SWSURFACE = 0
	PREALLOC  = 0x00000001
	RLEACCEL  = 0x00000002
	DONTFREE  = 0x00000004
)

type Surface struct {
	Flags    uint32
	Format   *PixelFormat
	W        int32
	H        int32
	Pitch    int
	pixels   unsafe.Pointer // use Pixels() for access
	UserData unsafe.Pointer
	Locked   int
	LockData unsafe.Pointer
	ClipRect Rect
	_map     *[0]byte
	RefCount int
}

type blit C.SDL_blit

func (surface *Surface) cptr() *C.SDL_Surface {
	return (*C.SDL_Surface)(unsafe.Pointer(surface))
}

func (surface *Surface) MustLock() bool {
	return (surface.Flags & RLEACCEL) != 0
}

func CreateRGBSurface(flags uint32, width, height, depth int32, Rmask, Gmask, Bmask, Amask uint32) *Surface {
	return (*Surface)(unsafe.Pointer(C.SDL_CreateRGBSurface(C.Uint32(flags),
                                                            C.int(width),
                                                            C.int(height),
                                                            C.int(depth),
                                                            C.Uint32(Rmask),
                                                            C.Uint32(Gmask),
                                                            C.Uint32(Bmask),
                                                            C.Uint32(Amask))))
}

func CreateRGBSurfaceFrom(pixels unsafe.Pointer, width, height, depth, pitch int, Rmask, Gmask, Bmask, Amask uint32) *Surface {
	return (*Surface)(unsafe.Pointer(C.SDL_CreateRGBSurfaceFrom(pixels,
                                                                C.int(width),
                                                                C.int(height),
                                                                C.int(depth),
                                                                C.int(pitch),
                                                                C.Uint32(Rmask),
                                                                C.Uint32(Gmask),
                                                                C.Uint32(Bmask),
                                                                C.Uint32(Amask))))
}

func (surface *Surface) Free() {
	C.SDL_FreeSurface(surface.cptr())
}

func (surface *Surface) SetPalette(palette *Palette) int {
	return int(C.SDL_SetSurfacePalette(surface.cptr(), palette.cptr()))
}

func (surface *Surface) Lock() {
	C.SDL_LockSurface(surface.cptr())
}

func (surface *Surface) Unlock() {
	C.SDL_UnlockSurface(surface.cptr())
}

func LoadBMP_RW(src *RWops, freeSrc int) *Surface {
    _surface := C.SDL_LoadBMP_RW(src.cptr(), C.int(freeSrc))
	return (*Surface)(unsafe.Pointer(_surface))
}

func LoadBMP(file string) *Surface {
	return (*Surface)(LoadBMP_RW(RWFromFile(file, "rb"), 1))
}

func (surface *Surface) SaveBMP_RW(dst *RWops, freeDst int) int {
	return int(C.SDL_SaveBMP_RW(surface.cptr(), dst.cptr(), C.int(freeDst)))
}

func (surface *Surface) SaveBMP(file string) int {
	return int(surface.SaveBMP_RW(RWFromFile(file, "wb"), 1))
}

func (surface *Surface) SetRLE(flag int) int {
	return int(C.SDL_SetSurfaceRLE(surface.cptr(), C.int(flag)))
}

func (surface *Surface) SetColorKey(flag int, key uint32) int {
	return int(C.SDL_SetColorKey(surface.cptr(), C.int(flag), C.Uint32(key)))
}

func (surface *Surface) GetColorKey() (key uint32, status int) {
	_key := (*C.Uint32)(unsafe.Pointer(&key))
	status = int(C.SDL_GetColorKey(surface.cptr(), _key))
	return key, status
}

func (surface *Surface) SetColorMod(r, g, b uint8) int {
	return int(C.SDL_SetSurfaceColorMod(surface.cptr(), C.Uint8(r), C.Uint8(g), C.Uint8(b)))
}

func (surface *Surface) GetColorMod() (r, g, b uint8, status int) {
	_r := (*C.Uint8)(unsafe.Pointer(&r))
	_g := (*C.Uint8)(unsafe.Pointer(&g))
	_b := (*C.Uint8)(unsafe.Pointer(&b))
	status = int(C.SDL_GetSurfaceColorMod(surface.cptr(), _r, _g, _b))
	return r, g, b, status
}

func (surface *Surface) SetAlphaMod(alpha uint8) int {
	return int(C.SDL_SetSurfaceAlphaMod(surface.cptr(), C.Uint8(alpha)))
}

func (surface *Surface) GetAlphaMod() (alpha uint8, status int) {
	_alpha := (*C.Uint8)(unsafe.Pointer(&alpha))
	status = int(C.SDL_GetSurfaceAlphaMod(surface.cptr(), _alpha))
	return alpha, status
}

func (surface *Surface) SetBlendMode(bm BlendMode) int {
	return int(C.SDL_SetSurfaceBlendMode(surface.cptr(), bm.c()))
}

func (surface *Surface) GetBlendMode() (bm BlendMode, status int) {
	status = int(C.SDL_GetSurfaceBlendMode(surface.cptr(), bm.cptr()))
	return bm, status
}

func (surface *Surface) SetClipRect(rect *Rect) bool {
	return C.SDL_SetClipRect(surface.cptr(), rect.cptr()) > 0
}

func (surface *Surface) GetClipRect(rect *Rect) {
	C.SDL_GetClipRect(surface.cptr(), rect.cptr())
}

func (surface *Surface) Convert(fmt *PixelFormat, flags uint32) *Surface {
	_surface := C.SDL_ConvertSurface(surface.cptr(), fmt.cptr(), C.Uint32(flags))
	return (*Surface)(unsafe.Pointer(_surface))
}

func (surface *Surface) ConvertFormat(pixelFormat uint32, flags uint32) *Surface {
	return (*Surface)(unsafe.Pointer(C.SDL_ConvertSurfaceFormat(surface.cptr(), C.Uint32(pixelFormat), C.Uint32(flags))))
}

func ConvertPixels(width, height int, srcFormat uint32, src unsafe.Pointer, srcPitch int,
	dstFormat uint32, dst unsafe.Pointer, dstPitch int) int {
	return int(C.SDL_ConvertPixels(C.int(width), C.int(height), C.Uint32(srcFormat), src, C.int(srcPitch), C.Uint32(dstFormat), dst, C.int(dstPitch)))
}

func (surface *Surface) FillRect(rect *Rect, color uint32) int {
	return int(C.SDL_FillRect(surface.cptr(), rect.cptr(), C.Uint32(color)))
}

func (surface *Surface) FillRects(rects []Rect, color uint32) int {
	return int(C.SDL_FillRects(surface.cptr(), rects[0].cptr(), C.int(len(rects)), C.Uint32(color)))
}

func (src *Surface) Blit(srcRect *Rect, dst *Surface, dstRect *Rect) int {
	return int(C.SDL_BlitSurface(src.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()))
}

func (src *Surface) BlitScaled(srcRect *Rect, dst *Surface, dstRect *Rect) int {
	return int(C.SDL_BlitScaled(src.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()))
}

func (src *Surface) UpperBlit(srcRect *Rect, dst *Surface, dstRect *Rect) int {
	return int(C.SDL_UpperBlit(src.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()))
}

func (src *Surface) LowerBlit(srcRect *Rect, dst *Surface, dstRect *Rect) int {
	return int(C.SDL_LowerBlit(src.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()))
}

func (src *Surface) SoftStretch(srcRect *Rect, dst *Surface, dstRect *Rect) int {
	return int(C.SDL_SoftStretch(src.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()))
}

func (src *Surface) UpperBlitScaled(srcRect *Rect, dst *Surface, dstRect *Rect) int {
	return int(C.SDL_UpperBlitScaled(src.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()))
}

func (src *Surface) LowerBlitScaled(srcRect *Rect, dst *Surface, dstRect *Rect) int {
	return int(C.SDL_LowerBlitScaled(src.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()))
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
