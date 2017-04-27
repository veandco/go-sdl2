package img

//#cgo linux freebsd darwin pkg-config: sdl2
//#cgo linux freebsd darwin LDFLAGS: -lSDL2_image
//#cgo windows LDFLAGS: -lSDL2 -lSDL2_image
//#include <stdlib.h>
//#include "sdl_image_wrapper.h"
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

// LinkedVersion
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_7.html)
func LinkedVersion() *sdl.Version {
	return (*sdl.Version)(unsafe.Pointer(C.IMG_Linked_Version()))
}

// Init (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_8.html)
func Init(flags int) int {
	_flags := (C.int)(flags)
	return int(C.IMG_Init(_flags))
}

// Quit (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_9.html)
func Quit() {
	C.IMG_Quit()
}

// GetError (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_45.html)
func GetError() error {
	e := C.IMG_GetError()
	if e == nil {
		return nil
	}
	return errors.New(C.GoString(e))
}

// LoadTyped_RW
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_13.html)
func LoadTyped_RW(src *sdl.RWops, freesrc bool, type_ string) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := (C.int)(sdl.Btoi(freesrc))
	_type := C.CString(type_)
	defer C.free(unsafe.Pointer(_type))
	_surface := C.IMG_LoadTyped_RW(_src, _freesrc, _type)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// Load (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_11.html)
func Load(file string) (*sdl.Surface, error) {
	_file := C.CString(file)
	defer C.free(unsafe.Pointer(_file))
	_surface := C.IMG_Load(_file)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// Load_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_12.html)
func Load_RW(src *sdl.RWops, freesrc bool) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := (C.int)(sdl.Btoi(freesrc))
	_surface := C.IMG_Load_RW(_src, _freesrc)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func LoadTexture(renderer *sdl.Renderer, file string) (*sdl.Texture, error) {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_file := C.CString(file)
	defer C.free(unsafe.Pointer(_file))
	_surface := C.IMG_LoadTexture(_renderer, _file)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Texture)(unsafe.Pointer(_surface)), nil
}

func LoadTexture_RW(renderer *sdl.Renderer, src *sdl.RWops, freesrc bool) (*sdl.Texture, error) {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := (C.int)(sdl.Btoi(freesrc))
	_surface := C.IMG_LoadTexture_RW(_renderer, _src, _freesrc)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Texture)(unsafe.Pointer(_surface)), nil
}

// IsICO (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_31.html)
func IsICO(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isICO(_src)) > 0
}

// IsCUR (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_30.html)
func IsCUR(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isCUR(_src)) > 0
}

// IsBMP (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_32.html)
func IsBMP(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isBMP(_src)) > 0
}

// IsGIF (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_37.html)
func IsGIF(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isGIF(_src)) > 0
}

// IsJPG (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_38.html)
func IsJPG(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isJPG(_src)) > 0
}

// IsLBM (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_41.html)
func IsLBM(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isLBM(_src)) > 0
}

// IsPCX (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_36.html)
func IsPCX(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isPCX(_src)) > 0
}

// IsPNG (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_40.html)
func IsPNG(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isPNG(_src)) > 0
}

// IsPNM (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_33.html)
func IsPNM(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isPNM(_src)) > 0
}

// IsTIF (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_39.html)
func IsTIF(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isTIF(_src)) > 0
}

// IsXCF (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_35.html)
func IsXCF(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isXCF(_src)) > 0
}

// IsXPM (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_34.html)
func IsXPM(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isXPM(_src)) > 0
}

// IsXV (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_42.html)
func IsXV(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isXV(_src)) > 0
}

func IsWEBP(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isWEBP(_src)) > 0
}

// LoadICO_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_15.html)
func LoadICO_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadICO_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadCUR_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_14.html)
func LoadCUR_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadCUR_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadBMP_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_16.html)
func LoadBMP_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadBMP_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadGIF_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_21.html)
func LoadGIF_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadGIF_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadJPG_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_22.html)
func LoadJPG_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadJPG_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadLBM_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_26.html)
func LoadLBM_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadLBM_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadPCX_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_20.html)
func LoadPCX_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadPCX_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadPNG_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_24.html)
func LoadPNG_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadPNG_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadPNM_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_17.html)
func LoadPNM_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadPNM_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadTGA_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_25.html)
func LoadTGA_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadTGA_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadTIF_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_23.html)
func LoadTIF_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadTIF_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadXCF_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_19.html)
func LoadXCF_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadXCF_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadXPM_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_18.html)
func LoadXPM_RW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadXPM_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadXV_RW (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_27.html)
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

// ReadXPMFromArray
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_28.html)
func ReadXPMFromArray(xpm string) (*sdl.Surface, error) {
	_xpm := C.CString(xpm)
	defer C.free(unsafe.Pointer(_xpm))
	_surface := C.IMG_ReadXPMFromArray(&_xpm)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

func SavePNG(surface *sdl.Surface, file string) error {
	_surface := (*C.SDL_Surface)(unsafe.Pointer(surface))
	_file := C.CString(file)
	defer C.free(unsafe.Pointer(_file))
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
