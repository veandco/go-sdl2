package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// PixelFormat (https://wiki.libsdl.org/SDL_PixelFormat)
type PixelFormat struct {
	Format        uint32
	Palette       *Palette
	BitsPerPixel  uint8
	BytesPerPixel uint8
	_             [2]uint8 // padding
	Rmask         uint32
	Gmask         uint32
	Bmask         uint32
	Amask         uint32
	Rloss         uint8
	Gloss         uint8
	Bloss         uint8
	Aloss         uint8
	Rshift        uint8
	Gshift        uint8
	Bshift        uint8
	Ashift        uint8
	RefCount      int32
	Next          *PixelFormat
}
type cPixelFormat C.SDL_PixelFormat

// Palette (https://wiki.libsdl.org/SDL_Palette)
type Palette struct {
	Ncolors  int32
	Colors   *Color
	Version  uint32
	RefCount int32
}
type cPalette C.SDL_Palette

// Color (https://wiki.libsdl.org/SDL_Color)
type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (c Color) Uint32() uint32 {
	var v uint32
	v |= uint32(c.A) << 24
	v |= uint32(c.R) << 16
	v |= uint32(c.G) << 8
	v |= uint32(c.B)
	return v
}

// RGBA returns the alpha-premultiplied red, green, blue and alpha values
// for the color. Each value ranges within [0, 0xffff], but is represented
// by a uint32 so that multiplying by a blend factor up to 0xffff will not
// overflow.
//
// An alpha-premultiplied color component c has been scaled by alpha (a),
// so has valid values 0 <= c <= a.
func (c Color) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = uint32(c.A)
	a |= a << 8
	return
}

const (
	PIXELTYPE_UNKNOWN  = C.SDL_PIXELTYPE_UNKNOWN
	PIXELTYPE_INDEX1   = C.SDL_PIXELTYPE_INDEX1
	PIXELTYPE_INDEX4   = C.SDL_PIXELTYPE_INDEX4
	PIXELTYPE_INDEX8   = C.SDL_PIXELTYPE_INDEX8
	PIXELTYPE_PACKED8  = C.SDL_PIXELTYPE_PACKED8
	PIXELTYPE_PACKED16 = C.SDL_PIXELTYPE_PACKED16
	PIXELTYPE_PACKED32 = C.SDL_PIXELTYPE_PACKED32
	PIXELTYPE_ARRAYU8  = C.SDL_PIXELTYPE_ARRAYU8
	PIXELTYPE_ARRAYU16 = C.SDL_PIXELTYPE_ARRAYU16
	PIXELTYPE_ARRAYU32 = C.SDL_PIXELTYPE_ARRAYU32
	PIXELTYPE_ARRAYF16 = C.SDL_PIXELTYPE_ARRAYF16
	PIXELTYPE_ARRAYF32 = C.SDL_PIXELTYPE_ARRAYF32
)

/** Bitmap pixel order high bit -> low bit. */
const (
	BITMAPORDER_NONE = C.SDL_BITMAPORDER_NONE
	BITMAPORDER_4321 = C.SDL_BITMAPORDER_4321
	BITMAPORDER_1234 = C.SDL_BITMAPORDER_1234
)

/** Packed component order high bit -> low bit. */
const (
	PACKEDORDER_NONE = C.SDL_PACKEDORDER_NONE
	PACKEDORDER_XRGB = C.SDL_PACKEDORDER_XRGB
	PACKEDORDER_RGBX = C.SDL_PACKEDORDER_RGBX
	PACKEDORDER_ARGB = C.SDL_PACKEDORDER_ARGB
	PACKEDORDER_RGBA = C.SDL_PACKEDORDER_RGBA
	PACKEDORDER_XBGR = C.SDL_PACKEDORDER_XBGR
	PACKEDORDER_BGRX = C.SDL_PACKEDORDER_BGRX
	PACKEDORDER_ABGR = C.SDL_PACKEDORDER_ABGR
	PACKEDORDER_BGRA = C.SDL_PACKEDORDER_BGRA
)

/** Array component order low byte -> high byte. */
const (
	ARRAYORDER_NONE = C.SDL_ARRAYORDER_NONE
	ARRAYORDER_RGB  = C.SDL_ARRAYORDER_RGB
	ARRAYORDER_RGBA = C.SDL_ARRAYORDER_RGBA
	ARRAYORDER_ARGB = C.SDL_ARRAYORDER_ARGB
	ARRAYORDER_BGR  = C.SDL_ARRAYORDER_BGR
	ARRAYORDER_BGRA = C.SDL_ARRAYORDER_BGRA
	ARRAYORDER_ABGR = C.SDL_ARRAYORDER_ABGR
)

/** Packed component layout. */
const (
	PACKEDLAYOUT_NONE    = C.SDL_PACKEDLAYOUT_NONE
	PACKEDLAYOUT_332     = C.SDL_PACKEDLAYOUT_332
	PACKEDLAYOUT_4444    = C.SDL_PACKEDLAYOUT_4444
	PACKEDLAYOUT_1555    = C.SDL_PACKEDLAYOUT_1555
	PACKEDLAYOUT_5551    = C.SDL_PACKEDLAYOUT_5551
	PACKEDLAYOUT_565     = C.SDL_PACKEDLAYOUT_565
	PACKEDLAYOUT_8888    = C.SDL_PACKEDLAYOUT_8888
	PACKEDLAYOUT_2101010 = C.SDL_PACKEDLAYOUT_2101010
	PACKEDLAYOUT_1010102 = C.SDL_PACKEDLAYOUT_1010102
)

const (
	PIXELFORMAT_UNKNOWN     = C.SDL_PIXELFORMAT_UNKNOWN
	PIXELFORMAT_INDEX1LSB   = C.SDL_PIXELFORMAT_INDEX1LSB
	PIXELFORMAT_INDEX1MSB   = C.SDL_PIXELFORMAT_INDEX1MSB
	PIXELFORMAT_INDEX4LSB   = C.SDL_PIXELFORMAT_INDEX4LSB
	PIXELFORMAT_INDEX4MSB   = C.SDL_PIXELFORMAT_INDEX4MSB
	PIXELFORMAT_INDEX8      = C.SDL_PIXELFORMAT_INDEX8
	PIXELFORMAT_RGB332      = C.SDL_PIXELFORMAT_RGB332
	PIXELFORMAT_RGB444      = C.SDL_PIXELFORMAT_RGB444
	PIXELFORMAT_RGB555      = C.SDL_PIXELFORMAT_RGB555
	PIXELFORMAT_BGR555      = C.SDL_PIXELFORMAT_BGR555
	PIXELFORMAT_ARGB4444    = C.SDL_PIXELFORMAT_ARGB4444
	PIXELFORMAT_RGBA4444    = C.SDL_PIXELFORMAT_RGBA4444
	PIXELFORMAT_ABGR4444    = C.SDL_PIXELFORMAT_ABGR4444
	PIXELFORMAT_BGRA4444    = C.SDL_PIXELFORMAT_BGRA4444
	PIXELFORMAT_ARGB1555    = C.SDL_PIXELFORMAT_ARGB1555
	PIXELFORMAT_RGBA5551    = C.SDL_PIXELFORMAT_RGBA5551
	PIXELFORMAT_ABGR1555    = C.SDL_PIXELFORMAT_ABGR1555
	PIXELFORMAT_BGRA5551    = C.SDL_PIXELFORMAT_BGRA5551
	PIXELFORMAT_RGB565      = C.SDL_PIXELFORMAT_RGB565
	PIXELFORMAT_BGR565      = C.SDL_PIXELFORMAT_BGR565
	PIXELFORMAT_RGB24       = C.SDL_PIXELFORMAT_RGB24
	PIXELFORMAT_BGR24       = C.SDL_PIXELFORMAT_BGR24
	PIXELFORMAT_RGB888      = C.SDL_PIXELFORMAT_RGB888
	PIXELFORMAT_RGBX8888    = C.SDL_PIXELFORMAT_RGBX8888
	PIXELFORMAT_BGR888      = C.SDL_PIXELFORMAT_BGR888
	PIXELFORMAT_BGRX8888    = C.SDL_PIXELFORMAT_BGRX8888
	PIXELFORMAT_ARGB8888    = C.SDL_PIXELFORMAT_ARGB8888
	PIXELFORMAT_RGBA8888    = C.SDL_PIXELFORMAT_RGBA8888
	PIXELFORMAT_ABGR8888    = C.SDL_PIXELFORMAT_ABGR8888
	PIXELFORMAT_BGRA8888    = C.SDL_PIXELFORMAT_BGRA8888
	PIXELFORMAT_ARGB2101010 = C.SDL_PIXELFORMAT_ARGB2101010
	PIXELFORMAT_YV12        = C.SDL_PIXELFORMAT_YV12
	PIXELFORMAT_IYUV        = C.SDL_PIXELFORMAT_IYUV
	PIXELFORMAT_YUY2        = C.SDL_PIXELFORMAT_YUY2
	PIXELFORMAT_UYVY        = C.SDL_PIXELFORMAT_UYVY
	PIXELFORMAT_YVYU        = C.SDL_PIXELFORMAT_YVYU
)

var (
	PIXELFORMAT_RGBA32 int
	PIXELFORMAT_ARGB32 int
	PIXELFORMAT_BGRA32 int
	PIXELFORMAT_ABGR32 int
)

const (
	ALPHA_OPAQUE      = C.SDL_ALPHA_OPAQUE
	ALPHA_TRANSPARENT = C.SDL_ALPHA_TRANSPARENT
)

func init() {
	if BYTEORDER == BIG_ENDIAN {
		PIXELFORMAT_RGBA32 = PIXELFORMAT_RGBA8888
		PIXELFORMAT_ARGB32 = PIXELFORMAT_ARGB8888
		PIXELFORMAT_BGRA32 = PIXELFORMAT_BGRA8888
		PIXELFORMAT_ABGR32 = PIXELFORMAT_ABGR8888
	} else {
		PIXELFORMAT_RGBA32 = PIXELFORMAT_ABGR8888
		PIXELFORMAT_ARGB32 = PIXELFORMAT_BGRA8888
		PIXELFORMAT_BGRA32 = PIXELFORMAT_ARGB8888
		PIXELFORMAT_ABGR32 = PIXELFORMAT_RGBA8888
	}
}

func (fmt *PixelFormat) cptr() *C.SDL_PixelFormat {
	return (*C.SDL_PixelFormat)(unsafe.Pointer(fmt))
}

func (p *Palette) cptr() *C.SDL_Palette {
	return (*C.SDL_Palette)(unsafe.Pointer(p))
}

/*
 * the following code is modified version of the code from bitbucket.org/dooots/go-sdl2
 */

// GetPixelFormatName gets the human readable name of a pixel format
// GetPixelFormatName (https://wiki.libsdl.org/SDL_GetPixelFormatName)
func GetPixelFormatName(format uint) string {
	return C.GoString(C.SDL_GetPixelFormatName(C.Uint32(format)))
}

// PixelFormatEnumToMasks converts format into a bpp and RGBA masks.
// PixelFormatEnumToMasks (https://wiki.libsdl.org/SDL_PixelFormatEnumToMasks)
func PixelFormatEnumToMasks(format uint) (bpp int, rmask, gmask, bmask, amask uint32, err error) {
	result := C.SDL_PixelFormatEnumToMasks(C.Uint32(format), (*C.int)(unsafe.Pointer(&bpp)),
		(*C.Uint32)(&rmask), (*C.Uint32)(&gmask), (*C.Uint32)(&bmask),
		(*C.Uint32)(&amask))
	if result == C.SDL_FALSE {
		err = GetError()
	}
	return
}

// MasksTouint converts a bpp and RGBA masks to a uint.
// MasksToPixelFormatEnum (https://wiki.libsdl.org/SDL_MasksToPixelFormatEnum)
func MasksToPixelFormatEnum(bpp int, rmask, gmask, bmask, amask uint32) uint {
	return uint(C.SDL_MasksToPixelFormatEnum(C.int(bpp), C.Uint32(rmask), C.Uint32(gmask),
		C.Uint32(bmask), C.Uint32(amask)))
}

// AllocFormat creates a PixelFormat structure from a uint.
// AllocFormat (https://wiki.libsdl.org/SDL_AllocFormat)
func AllocFormat(format uint) (*PixelFormat, error) {
	r := (*PixelFormat)(unsafe.Pointer(C.SDL_AllocFormat(C.Uint32(format))))
	if r == nil {
		return nil, GetError()
	}
	return r, nil
}

// Free frees a PixelFormat created by AllocFormat.
func (format *PixelFormat) Free() {
	C.SDL_FreeFormat((*C.SDL_PixelFormat)(unsafe.Pointer(format)))
}

// AllocPalette create a palette structure with the specified number of color
// entries.
// AllocPalette (https://wiki.libsdl.org/SDL_AllocPalette)
func AllocPalette(ncolors int) (*Palette, error) {
	r := (*Palette)(unsafe.Pointer(C.SDL_AllocPalette(C.int(ncolors))))
	if r == nil {
		return nil, GetError()
	}
	return r, nil
}

// SetPalette sets the palette for format.
// PixelFormat (https://wiki.libsdl.org/SDL_SetPixelFormatPalette)
func (format *PixelFormat) SetPalette(palette *Palette) error {
	r := C.SDL_SetPixelFormatPalette((*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		(*C.SDL_Palette)(unsafe.Pointer(palette)))
	if r != 0 {
		return GetError()
	}
	return nil
}

// SetColors sets a range of colors in a palette.
// Palette (https://wiki.libsdl.org/SDL_SetPaletteColors)
func (palette *Palette) SetColors(colors []Color) error {
	var ptr *C.SDL_Color
	if len(colors) > 0 {
		ptr = (*C.SDL_Color)(unsafe.Pointer(&colors[0]))
	}

	r := C.SDL_SetPaletteColors((*C.SDL_Palette)(unsafe.Pointer(palette)),
		ptr, 0, C.int(len(colors)))
	if r != 0 {
		return GetError()
	}
	return nil
}

// Free frees a palette created with AllocPalette.
// Palette (https://wiki.libsdl.org/SDL_FreePalette)
func (palette *Palette) Free() {
	C.SDL_FreePalette((*C.SDL_Palette)(unsafe.Pointer(palette)))
}

// MapRGB maps an RGB triple to an opaque pixel value for a given pixel format.
// MapRGB (https://wiki.libsdl.org/SDL_MapRGB)
func MapRGB(format *PixelFormat, r, g, b uint8) uint32 {
	return uint32(C.SDL_MapRGB((*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		C.Uint8(r), C.Uint8(g), C.Uint8(b)))
}

// MapRGBA maps an RGBA quadruple to a pixel value for a given pixel format.
// MapRGBA (https://wiki.libsdl.org/SDL_MapRGBA)
func MapRGBA(format *PixelFormat, r, g, b, a uint8) uint32 {
	return uint32(C.SDL_MapRGBA((*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
}

// GetRGB gets the RGB components from a pixel of the specified format.
// GetRGB (https://wiki.libsdl.org/SDL_GetRGB)
func GetRGB(pixel uint32, format *PixelFormat) (r, g, b uint8) {
	C.SDL_GetRGB(C.Uint32(pixel), (*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		(*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b))
	return
}

// GetRGBA gets the RGBA components from a pixel of the specified format.
// GetRGBA (https://wiki.libsdl.org/SDL_GetRGBA)
func GetRGBA(pixel uint32, format *PixelFormat) (r, g, b, a uint8) {
	C.SDL_GetRGBA(C.Uint32(pixel), (*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		(*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b), (*C.Uint8)(&a))
	return
}

// CalculateGammaRamp calculates a 256 entry gamma ramp for a gamma value.
// CalculateGammaRamp (https://wiki.libsdl.org/SDL_CalculateGammaRamp)
func CalculateGammaRamp(gamma float32, ramp *[256]uint16) {
	C.SDL_CalculateGammaRamp(C.float(gamma), (*C.Uint16)(unsafe.Pointer(&ramp[0])))
}
