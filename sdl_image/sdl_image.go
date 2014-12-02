package img

//#cgo linux freebsd darwin pkg-config: sdl2
//#cgo linux freebsd darwin LDFLAGS: -lSDL2_image
//#cgo windows LDFLAGS: -lSDL2 -lSDL2_image
//#include <stdlib.h>
//#include "sdl_image.h"
import "C"
import "unsafe"
import "errors"
import "github.com/veandco/go-sdl2/sdl"

const (
	INIT_JPG  = 0x00000001
	INIT_PNG  = 0x00000002
	INIT_TIF  = 0x00000004
	INIT_WEBP = 0x00000008
)

func LinkedVersion() *sdl.Version {
	return (*sdl.Version)(unsafe.Pointer(C.IMG_Linked_Version()))
}

func Init(flags int) int {
	_flags := (C.int)(flags)
	return int(C.IMG_Init(_flags))
}

func Quit() {
	C.IMG_Quit()
}

func GetError() error {
	e := C.IMG_GetError()
	if e == nil {
		return nil
	}
	return errors.New(C.GoString(e))
}

func LoadTyped_RW(src *sdl.RWops, freesrc int, type_ string) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := (C.int)(freesrc)
	_type := C.CString(type_)
	defer C.SDL_free(unsafe.Pointer(_type))
	_surface := C.IMG_LoadTyped_RW(_src, _freesrc, _type)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func Load(file string) (*sdl.Surface, error) {
	_file := C.CString(file)
	defer C.SDL_free(unsafe.Pointer(_file))
	_surface := C.IMG_Load(_file)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func Load_RW(src *sdl.RWops, freesrc int) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := (C.int)(freesrc)
	_surface := C.IMG_Load_RW(_src, _freesrc)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadTexture(renderer *sdl.Renderer, file string) (*sdl.Texture, error) {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_file := C.CString(file)
	defer C.SDL_free(unsafe.Pointer(_file))
	_surface := C.IMG_LoadTexture(_renderer, _file)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Texture)(unsafe.Pointer(_surface)), nil
}

func LoadTexture_RW(renderer *sdl.Renderer, src *sdl.RWops, freesrc int) (*sdl.Texture, error) {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := (C.int)(freesrc)
	_surface := C.IMG_LoadTexture_RW(_renderer, _src, _freesrc)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Texture)(unsafe.Pointer(_surface)), nil
}

func IsICO(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isICO(_src)) > 0
}

func IsCUR(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isCUR(_src)) > 0
}

func IsBMP(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isBMP(_src)) > 0
}

func IsGIF(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isGIF(_src)) > 0
}

func IsJPG(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isJPG(_src)) > 0
}

func IsLBM(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isLBM(_src)) > 0
}

func IsPCX(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isPCX(_src)) > 0
}

func IsPNG(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isPNG(_src)) > 0
}

func IsPNM(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isPNM(_src)) > 0
}

func IsTIF(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isTIF(_src)) > 0
}

func IsXCF(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isXCF(_src)) > 0
}

func IsXPM(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isXPM(_src)) > 0
}

func IsXV(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isXV(_src)) > 0
}

func IsWEBP(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isWEBP(_src)) > 0
}

func LoadICO_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadICO_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadCUR_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadCUR_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadBMP_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadBMP_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadGIF_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadGIF_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadJPG_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadJPG_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadLBM_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadLBM_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadPCX_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadPCX_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadPNG_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadPNG_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadPNM_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadPNM_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadTGA_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadTGA_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadTIF_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadTIF_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadXCF_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface :=C.IMG_LoadXCF_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadXPM_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadXPM_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadXV_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadXV_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadWEBP_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadWEBP_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func ReadXPMFromArray(xpm string) (*sdl.Surface, error) {
	_xpm := C.CString(xpm)
	C.SDL_free(unsafe.Pointer(_xpm))
	_surface := C.IMG_ReadXPMFromArray(&_xpm)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func SavePNG(surface *sdl.Surface, file string) error {
	_surface := (*C.SDL_Surface)(unsafe.Pointer(surface))
	_file := C.CString(file)
	C.SDL_free(unsafe.Pointer(_file))
	_ret := C.IMG_SavePNG(_surface, _file)
	if _ret < 0 {
		return GetError()
	}
	return nil
}

func SavePNG_RW(surface *sdl.Surface, dst *sdl.RWops, freedst int) error {
	_surface := (*C.SDL_Surface)(unsafe.Pointer(surface))
	_dst := (*C.SDL_RWops)(unsafe.Pointer(dst))
	_freedst := (C.int)(freedst)
	_ret := C.IMG_SavePNG_RW(_surface, _dst, _freedst)
	if _ret < 0 {
		return GetError()
	}
	return nil
}
