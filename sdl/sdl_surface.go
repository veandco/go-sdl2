package sdl

/*
#include <SDL2/SDL_surface.h>
*/
import "C"
import "unsafe"

type Surface struct {
	Flags uint32
	Format PixelFormat
	W int
	H int
	Pitch int
	Pixels unsafe.Pointer
	UserData unsafe.Pointer
	Locked int
	LockData unsafe.Pointer
	ClipRect Rect
	_map *[0]byte
	RefCount int
}

type blit C.SDL_blit

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
	return (*Surface) (unsafe.Pointer(C.SDL_LoadBMP_RW(_src, _freesrc)))
}

func LoadBMP(file string) *Surface {
	_file := (C.CString) (file)
	return (*Surface) (unsafe.Pointer(C.SDL_LoadBMP_RW(C.SDL_RWFromFile(_file, C.CString("rb")), 1)))
}

func (surface *Surface) SaveBMP_RW(dst *RWops, freedst int) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_dst := (*C.SDL_RWops) (unsafe.Pointer(dst))
	_freedst := (C.int) (freedst)
	return (int) (C.SDL_SaveBMP_RW(_surface, _dst, _freedst))
}

func (surface *Surface) SaveBMP(file string) int {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	_file := (C.CString) (file)
	return (int) (C.SDL_SaveBMP_RW(_surface, C.SDL_RWFromFile(_file, C.CString("rb")), 1))
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
