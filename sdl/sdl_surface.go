package sdl

/*
#include <SDL2/SDL_surface.h>
*/
import "C"
import "unsafe"
import "reflect"

const (
	SWSURFACE	= 0
	PREALLOC	= 0x00000001
	RLEACCEL	= 0x00000002
	DONTFREE	= 0x00000004
)

type Surface struct {
	Flags uint32
	Format *PixelFormat
	W int32
	H int32
	Pitch int
	pixels unsafe.Pointer // use Pixels() for access
	UserData unsafe.Pointer
	Locked int
	LockData unsafe.Pointer
	ClipRect Rect
	_map *[0]byte
	RefCount int
}

type blit C.SDL_blit

func (surface *Surface) MustLock() bool {
	return (surface.Flags & RLEACCEL) != 0
}

func CreateRGBSurface(flags uint32, width, height, depth int32,
		Rmask, Gmask, Bmask, Amask uint32) *Surface {
	_flags := (C.Uint32) (flags)
	_width := (C.int) (width)
	_height := (C.int) (height)
	_depth := (C.int) (depth)
	_Rmask := (C.Uint32) (Rmask)
	_Gmask := (C.Uint32) (Gmask)
	_Bmask := (C.Uint32) (Bmask)
	_Amask := (C.Uint32) (Amask)
	return (*Surface) (unsafe.Pointer(C.SDL_CreateRGBSurface(_flags, _width, _height,
				_depth, _Rmask, _Gmask, _Bmask, _Amask)))
}

func CreateRGBSurfaceFrom(pixels unsafe.Pointer, width, height, depth, pitch int32,
		Rmask, Gmask, Bmask, Amask uint32) *Surface {
	_width := (C.int) (width)
	_height := (C.int) (height)
	_depth := (C.int) (depth)
	_pitch := (C.int) (pitch)
	_Rmask := (C.Uint32) (Rmask)
	_Gmask := (C.Uint32) (Gmask)
	_Bmask := (C.Uint32) (Bmask)
	_Amask := (C.Uint32) (Amask)
	return (*Surface) (unsafe.Pointer(C.SDL_CreateRGBSurfaceFrom(pixels, _width, _height,
				_depth, _pitch, _Rmask, _Gmask, _Bmask, _Amask)))
}

func (surface *Surface) Free() {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	C.SDL_FreeSurface(_surface)
}

func (surface *Surface) SetPalette(palette *Palette) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_palette := (*C.SDL_Palette) (unsafe.Pointer(palette))
	return (int) (C.SDL_SetSurfacePalette(_surface, _palette))
}

func (surface *Surface) Lock() {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	C.SDL_LockSurface(_surface)
}

func (surface *Surface) Unlock() {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	C.SDL_UnlockSurface(_surface)
}

func LoadBMP_RW(src *RWops, freesrc int) *Surface {
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	_freesrc := (C.int) (freesrc)
	_surface := C.SDL_LoadBMP_RW(_src, _freesrc)
	return (*Surface) (unsafe.Pointer(_surface))
}

func LoadBMP(file string) *Surface {
	return (*Surface) (LoadBMP_RW(RWFromFile(file, "rb"), 1))
}

func (surface *Surface) SaveBMP_RW(dst *RWops, freedst int) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_dst := (*C.SDL_RWops) (unsafe.Pointer(dst))
	_freedst := (C.int) (freedst)
	return (int) (C.SDL_SaveBMP_RW(_surface, _dst, _freedst))
}

func (surface *Surface) SaveBMP(file string) int {
	return (int) (surface.SaveBMP_RW(RWFromFile(file, "wb"), 1))
}

func (surface *Surface) SetRLE(flag int) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_flag := (C.int) (flag)
	return (int) (C.SDL_SetSurfaceRLE(_surface, _flag))
}

func (surface *Surface) SetColorKey(flag int, key uint32) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_flag := (C.int) (flag)
	_key := (C.Uint32) (key)
	return (int) (C.SDL_SetColorKey(_surface, _flag, _key))
}

func (surface *Surface) GetColorKey(key *uint32) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_key := (*C.Uint32) (unsafe.Pointer(key))
	return (int) (C.SDL_GetColorKey(_surface, _key))
}

func (surface *Surface) SetColorMod(r, g, b uint8) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_r := (C.Uint8) (r)
	_g := (C.Uint8) (g)
	_b := (C.Uint8) (b)
	return (int) (C.SDL_SetSurfaceColorMod(_surface, _r, _g, _b))
}

func (surface *Surface) GetColorMod(r, g, b *uint8) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_r := (*C.Uint8) (unsafe.Pointer(r))
	_g := (*C.Uint8) (unsafe.Pointer(g))
	_b := (*C.Uint8) (unsafe.Pointer(b))
	return (int) (C.SDL_GetSurfaceColorMod(_surface, _r, _g, _b))
}

func (surface *Surface) SetAlphaMod(alpha uint8) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_alpha := (C.Uint8) (alpha)
	return (int) (C.SDL_SetSurfaceAlphaMod(_surface, _alpha))
}

func (surface *Surface) GetAlphaMod(alpha *uint8) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_alpha := (*C.Uint8) (unsafe.Pointer(alpha))
	return (int) (C.SDL_GetSurfaceAlphaMod(_surface, _alpha))
}

func (surface *Surface) SetBlendMode(blendMode BlendMode) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_blendMode := (C.SDL_BlendMode) (blendMode)
	return (int) (C.SDL_SetSurfaceBlendMode(_surface, _blendMode))
}

func (surface *Surface) GetBlendMode(blendMode *BlendMode) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_blendMode := (*C.SDL_BlendMode) (unsafe.Pointer(blendMode))
	return (int) (C.SDL_GetSurfaceBlendMode(_surface, _blendMode))
}

func (surface *Surface) SetClipRect(rect *Rect) bool {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_rect := (*C.SDL_Rect) (unsafe.Pointer(rect))
	return C.SDL_SetClipRect(_surface, _rect) > 0
}

func (surface *Surface) GetClipRect(rect *Rect) {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_rect := (*C.SDL_Rect) (unsafe.Pointer(rect))
	C.SDL_GetClipRect(_surface, _rect)
}

func (surface *Surface) Convert(fmt *PixelFormat, flags uint32) *Surface {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_fmt := (*C.SDL_PixelFormat) (unsafe.Pointer(fmt))
	_flags := (C.Uint32) (flags)
	_surface = C.SDL_ConvertSurface(_surface, _fmt, _flags)
	return (*Surface) (unsafe.Pointer(_surface))
}

func (surface *Surface) ConvertFormat(pixel_format uint32, flags uint32) *Surface {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_pixel_format := (C.Uint32) (pixel_format)
	_flags := (C.Uint32) (flags)
	return (*Surface) (unsafe.Pointer(C.SDL_ConvertSurfaceFormat(_surface, _pixel_format, _flags)))
}

func ConvertPixels(width, height int, src_format uint32, src unsafe.Pointer, src_pitch int,
		dst_format uint32, dst unsafe.Pointer, dst_pitch int) int {
	_width := (C.int) (width)
	_height := (C.int) (height)
	_src_format := (C.Uint32) (src_format)
	_src_pitch := (C.int) (src_pitch)
	_dst_format := (C.Uint32) (dst_format)
	_dst_pitch := (C.int) (dst_pitch)
	return (int) (C.SDL_ConvertPixels(_width, _height, _src_format, src, _src_pitch, _dst_format, dst, _dst_pitch))
}

func (surface *Surface) FillRect(rect *Rect, color uint32) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_rect := (*C.SDL_Rect) (unsafe.Pointer(rect))
	_color := (C.Uint32) (color)
	return (int) (C.SDL_FillRect(_surface, _rect, _color))
}

func (surface *Surface) FillRects(rects *Rect, count int, color uint32) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_rects := (*C.SDL_Rect) (unsafe.Pointer(rects))
	_count := (C.int) (count)
	_color := (C.Uint32) (color)
	return (int) (C.SDL_FillRects(_surface, _rects, _count, _color))
}

func (src *Surface) Blit(srcrect *Rect, dst *Surface, dstrect *Rect) int {
	_src := (*C.SDL_Surface) (unsafe.Pointer(src))
	_srcrect := (*C.SDL_Rect) (unsafe.Pointer(srcrect))
	_dst := (*C.SDL_Surface) (unsafe.Pointer(dst))
	_dstrect := (*C.SDL_Rect) (unsafe.Pointer(dstrect))
	return (int) (C.SDL_BlitSurface(_src, _srcrect, _dst, _dstrect))
}

func (src *Surface) BlitScaled(srcrect *Rect, dst *Surface, dstrect *Rect) int {
	_src := (*C.SDL_Surface) (unsafe.Pointer(src))
	_srcrect := (*C.SDL_Rect) (unsafe.Pointer(srcrect))
	_dst := (*C.SDL_Surface) (unsafe.Pointer(dst))
	_dstrect := (*C.SDL_Rect) (unsafe.Pointer(dstrect))
	return (int) (C.SDL_BlitScaled(_src, _srcrect, _dst, _dstrect))
}

func (src *Surface) UpperBlit(srcrect *Rect, dst *Surface, dstrect *Rect) int {
	_src := (*C.SDL_Surface) (unsafe.Pointer(src))
	_srcrect := (*C.SDL_Rect) (unsafe.Pointer(srcrect))
	_dst := (*C.SDL_Surface) (unsafe.Pointer(dst))
	_dstrect := (*C.SDL_Rect) (unsafe.Pointer(dstrect))
	return (int) (C.SDL_UpperBlit(_src, _srcrect, _dst, _dstrect))
}

func (src *Surface) LowerBlit(srcrect *Rect, dst *Surface, dstrect *Rect) int {
	_src := (*C.SDL_Surface) (unsafe.Pointer(src))
	_srcrect := (*C.SDL_Rect) (unsafe.Pointer(srcrect))
	_dst := (*C.SDL_Surface) (unsafe.Pointer(dst))
	_dstrect := (*C.SDL_Rect) (unsafe.Pointer(dstrect))
	return (int) (C.SDL_LowerBlit(_src, _srcrect, _dst, _dstrect))
}

func (src *Surface) SoftStretch(srcrect *Rect, dst *Surface, dstrect *Rect) int {
	_src := (*C.SDL_Surface) (unsafe.Pointer(src))
	_srcrect := (*C.SDL_Rect) (unsafe.Pointer(srcrect))
	_dst := (*C.SDL_Surface) (unsafe.Pointer(dst))
	_dstrect := (*C.SDL_Rect) (unsafe.Pointer(dstrect))
	return (int) (C.SDL_SoftStretch(_src, _srcrect, _dst, _dstrect))
}

func (src *Surface) UpperBlitScaled(srcrect *Rect, dst *Surface, dstrect *Rect) int {
	_src := (*C.SDL_Surface) (unsafe.Pointer(src))
	_srcrect := (*C.SDL_Rect) (unsafe.Pointer(srcrect))
	_dst := (*C.SDL_Surface) (unsafe.Pointer(dst))
	_dstrect := (*C.SDL_Rect) (unsafe.Pointer(dstrect))
	return (int) (C.SDL_UpperBlitScaled(_src, _srcrect, _dst, _dstrect))
}

func (src *Surface) LowerBlitScaled(srcrect *Rect, dst *Surface, dstrect *Rect) int {
	_src := (*C.SDL_Surface) (unsafe.Pointer(src))
	_srcrect := (*C.SDL_Rect) (unsafe.Pointer(srcrect))
	_dst := (*C.SDL_Surface) (unsafe.Pointer(dst))
	_dstrect := (*C.SDL_Rect) (unsafe.Pointer(dstrect))
	return (int) (C.SDL_LowerBlitScaled(_src, _srcrect, _dst, _dstrect))
}

func (surface *Surface) Pixels() []byte {
	var b []byte
	length := int(surface.W * surface.H) * int(surface.Format.BytesPerPixel)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(surface.pixels)
	return b
}
