// Package img is a simple library to load images of various formats as SDL surfaces.
package img

//#include <stdlib.h>
//#include "sdl_image_wrapper.h"
import "C"
import "unsafe"
import "errors"
import "github.com/veandco/go-sdl2/sdl"

// Flags which may be passed to img.Init() to load support of image formats, can be bitwise OR'd together.
const (
	INIT_JPG  = 0x00000001 // JPG
	INIT_PNG  = 0x00000002 // PNG
	INIT_TIF  = 0x00000004 // TIF
	INIT_WEBP = 0x00000008 // WebP
)

// LinkedVersion returns the version of the dynamically linked SDL_image library.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_7.html)
func LinkedVersion() *sdl.Version {
	return (*sdl.Version)(unsafe.Pointer(C.IMG_Linked_Version()))
}

// Init loads dynamic libraries and prepares them for use. Flags should be one or more flags from IMG_InitFlags OR'd together. It returns the flags successfully initialized, or 0 on failure.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_8.html)
func Init(flags int) error {
	_flags := (C.int)(flags)
	ret := int(C.IMG_Init(_flags))
	if ret == 0 {
		return GetError()
	}
	return nil
}

// Quit unloads libraries loaded with img.Init().
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_9.html)
func Quit() {
	C.IMG_Quit()
}

// GetError returns the last error that occurred, or an empty string if there hasn't been an error message set since the last call to sdl.ClearError().
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_45.html)
func GetError() error {
	e := C.IMG_GetError()
	if e == nil {
		return nil
	}
	return errors.New(C.GoString(e))
}

// LoadTypedRW loads an image from an SDL data source. The 'type' may be one of: "BMP", "GIF", "PNG", etc. If the image format supports a transparent pixel, SDL will set the colorkey for the surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_13.html)
func LoadTypedRW(src *sdl.RWops, freesrc bool, type_ string) (*sdl.Surface, error) {
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

// Load loads a file for use as an image in a new surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_11.html)
func Load(file string) (*sdl.Surface, error) {
	_file := C.CString(file)
	defer C.free(unsafe.Pointer(_file))
	_surface := C.IMG_Load(_file)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadRW loads an image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_12.html)
func LoadRW(src *sdl.RWops, freesrc bool) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := (C.int)(sdl.Btoi(freesrc))
	_surface := C.IMG_Load_RW(_src, _freesrc)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadTexture loads an image directly into a render texture.
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

// LoadTextureRW loads an image from an SDL data source directly into a render texture.
func LoadTextureRW(renderer *sdl.Renderer, src *sdl.RWops, freesrc bool) (*sdl.Texture, error) {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := (C.int)(sdl.Btoi(freesrc))
	_surface := C.IMG_LoadTexture_RW(_renderer, _src, _freesrc)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Texture)(unsafe.Pointer(_surface)), nil
}

// IsICO reports whether ICO format is supported and image data is readable as an ICO.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_31.html)
func IsICO(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isICO(_src)) > 0
}

// IsCUR reports whether CUR format is supported and image data is readable as a CUR.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_30.html)
func IsCUR(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isCUR(_src)) > 0
}

// IsBMP reports whether BMP format is supported and image data is readable as a BMP.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_32.html)
func IsBMP(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isBMP(_src)) > 0
}

// IsGIF reports whether GIF format is supported and image data is readable as a GIF.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_37.html)
func IsGIF(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isGIF(_src)) > 0
}

// IsJPG reports whether JPG format is supported and image data is readable as a JPG.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_38.html)
func IsJPG(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isJPG(_src)) > 0
}

// IsLBM reports whether LBM format is supported and image data is readable as an LBM.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_41.html)
func IsLBM(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isLBM(_src)) > 0
}

// IsPCX reports whether PCX format is supported and image data is readable as a PCX.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_36.html)
func IsPCX(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isPCX(_src)) > 0
}

// IsPNG reports whether PNG format is supported and image data is readable as a PNG.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_40.html)
func IsPNG(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isPNG(_src)) > 0
}

// IsPNM reports whether PNM format is supported and image data is readable as a PNM.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_33.html)
func IsPNM(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isPNM(_src)) > 0
}

// IsTIF reports whether TIF format is supported and image data is readable as a TIF.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_39.html)
func IsTIF(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isTIF(_src)) > 0
}

// IsXCF reports whether XCF format is supported and image data is readable as an XCF.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_35.html)
func IsXCF(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isXCF(_src)) > 0
}

// IsXPM reports whether XPM format is supported and image data is readable as an XPM.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_34.html)
func IsXPM(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isXPM(_src)) > 0
}

// IsXV reports whether XV format is supported and image data is readable as an XV thumbnail.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_42.html)
func IsXV(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isXV(_src)) > 0
}

// IsWEBP reports whether WEBP format is supported and image data is readable as a WEBP.
func IsWEBP(src *sdl.RWops) bool {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return int(C.IMG_isWEBP(_src)) > 0
}

// LoadICORW loads an ICO image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_15.html)
func LoadICORW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadICO_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadCURRW loads a CUR image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_14.html)
func LoadCURRW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadCUR_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadBMPRW loads a BMP image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_16.html)
func LoadBMPRW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadBMP_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadGIFRW loads a GIF image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_21.html)
func LoadGIFRW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadGIF_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadJPGRW loads a JPG image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_22.html)
func LoadJPGRW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadJPG_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadLBMRW loads an LBM image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_26.html)
func LoadLBMRW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadLBM_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadPCXRW loads a PCX image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_20.html)
func LoadPCXRW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadPCX_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadPNGRW loads a PNG image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_24.html)
func LoadPNGRW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadPNG_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadPNMRW loads a PNM image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_17.html)
func LoadPNMRW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadPNM_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadTGARW loads a TGA image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_25.html)
func LoadTGARW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadTGA_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadTIFRW loads a TIF image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_23.html)
func LoadTIFRW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadTIF_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadXCFRW loads an XCF image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_19.html)
func LoadXCFRW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadXCF_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadXPMRW loads an XPM image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_18.html)
func LoadXPMRW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadXPM_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadXVRW loads an XV thumbnail image from an SDL data source for use as a surface.
// (http://www.libsdl.org/projects/SDL_image/docs/SDL_image_27.html)
func LoadXVRW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadXV_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// LoadWEBPRW loads a WEBP image from an SDL data source for use as a surface.
func LoadWEBPRW(src *sdl.RWops) (*sdl.Surface, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_surface := C.IMG_LoadWEBP_RW(_src)
	if _surface == nil {
		return nil, GetError()
	}
	return (*sdl.Surface)(unsafe.Pointer(_surface)), nil
}

// ReadXPMFromArray loads an XPM image from xpm data for use as a surface.
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

// SavePNG saves a surface as PNG file.
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

// SavePNGRW saves a surface to an SDL data source.
func SavePNGRW(surface *sdl.Surface, dst *sdl.RWops, freedst int) error {
	_surface := (*C.SDL_Surface)(unsafe.Pointer(surface))
	_dst := (*C.SDL_RWops)(unsafe.Pointer(dst))
	_freedst := (C.int)(freedst)
	_ret := C.IMG_SavePNG_RW(_surface, _dst, _freedst)
	if _ret < 0 {
		return GetError()
	}
	return nil
}
