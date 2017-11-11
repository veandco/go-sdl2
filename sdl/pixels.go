package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// PixelFormat contains pixel format information.
// (https://wiki.libsdl.org/SDL_PixelFormat)
type PixelFormat struct {
	Format        uint32       // one of the PIXELFORMAT values (https://wiki.libsdl.org/SDL_PixelFormatEnum)
	Palette       *Palette     // palette structure associated with this pixel format, or nil if the format doesn't have a palette (https://wiki.libsdl.org/SDL_Palette)
	BitsPerPixel  uint8        // the number of significant bits in a pixel value, eg: 8, 15, 16, 24, 32
	BytesPerPixel uint8        // the number of bytes required to hold a pixel value, eg: 1, 2, 3, 4
	_             [2]uint8     // padding
	Rmask         uint32       // a mask representing the location of the red component of the pixel
	Gmask         uint32       // a mask representing the location of the green component of the pixel
	Bmask         uint32       // a mask representing the location of the blue component of the pixel
	Amask         uint32       // a mask representing the location of the alpha component of the pixel or 0 if the pixel format doesn't have any alpha information
	rLoss         uint8        // (internal use)
	gLoss         uint8        // (internal use)
	bLoss         uint8        // (internal use)
	aLoss         uint8        // (internal use)
	rShift        uint8        // (internal use)
	gShift        uint8        // (internal use)
	bShift        uint8        // (internal use)
	aShift        uint8        // (internal use)
	refCount      int32        // (internal use)
	next          *PixelFormat // (internal use)
}
type cPixelFormat C.SDL_PixelFormat

// Palette contains palette information.
// (https://wiki.libsdl.org/SDL_Palette)
type Palette struct {
	Ncolors  int32  // the number of colors in the palette
	Colors   *Color // an array of Color structures representing the palette (https://wiki.libsdl.org/SDL_Color)
	version  uint32 // incrementally tracks changes to the palette (internal use)
	refCount int32  // reference count (internal use)
}
type cPalette C.SDL_Palette

// Color represents a color.
// (https://wiki.libsdl.org/SDL_Color)
type Color struct {
	R uint8 // the red component in the range 0-255
	G uint8 // the green component in the range 0-255
	B uint8 // the blue component in the range 0-255
	A uint8 // the alpha component in the range 0-255
}

// Uint32 return uint32 representation of RGBA color.
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

// Pixel types.
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

// Bitmap pixel order high bit -> low bit.
const (
	BITMAPORDER_NONE = C.SDL_BITMAPORDER_NONE
	BITMAPORDER_4321 = C.SDL_BITMAPORDER_4321
	BITMAPORDER_1234 = C.SDL_BITMAPORDER_1234
)

// Packed component order high bit -> low bit.
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

// Array component order low byte -> high byte.
const (
	ARRAYORDER_NONE = C.SDL_ARRAYORDER_NONE
	ARRAYORDER_RGB  = C.SDL_ARRAYORDER_RGB
	ARRAYORDER_RGBA = C.SDL_ARRAYORDER_RGBA
	ARRAYORDER_ARGB = C.SDL_ARRAYORDER_ARGB
	ARRAYORDER_BGR  = C.SDL_ARRAYORDER_BGR
	ARRAYORDER_BGRA = C.SDL_ARRAYORDER_BGRA
	ARRAYORDER_ABGR = C.SDL_ARRAYORDER_ABGR
)

// Packed component layout.
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

// Pixel format values.
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

// Pixel format variables.
var (
	PIXELFORMAT_RGBA32 int
	PIXELFORMAT_ARGB32 int
	PIXELFORMAT_BGRA32 int
	PIXELFORMAT_ABGR32 int
)

// These define alpha as the opacity of a surface.
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

func (format *PixelFormat) cptr() *C.SDL_PixelFormat {
	return (*C.SDL_PixelFormat)(unsafe.Pointer(format))
}

func (palette *Palette) cptr() *C.SDL_Palette {
	return (*C.SDL_Palette)(unsafe.Pointer(palette))
}

/*
 * the following code is modified version of the code from bitbucket.org/dooots/go-sdl2
 */

// GetPixelFormatName returns the human readable name of a pixel format.
// (https://wiki.libsdl.org/SDL_GetPixelFormatName)
func GetPixelFormatName(format uint) string {
	return C.GoString(C.SDL_GetPixelFormatName(C.Uint32(format)))
}

// PixelFormatEnumToMasks converts one of the enumerated pixel formats to a bpp value and RGBA masks.
// (https://wiki.libsdl.org/SDL_PixelFormatEnumToMasks)
func PixelFormatEnumToMasks(format uint) (bpp int, rmask, gmask, bmask, amask uint32, err error) {
	result := C.SDL_PixelFormatEnumToMasks(C.Uint32(format), (*C.int)(unsafe.Pointer(&bpp)),
		(*C.Uint32)(&rmask), (*C.Uint32)(&gmask), (*C.Uint32)(&bmask),
		(*C.Uint32)(&amask))
	if result == C.SDL_FALSE {
		err = GetError()
	}
	return
}

// MasksToPixelFormatEnum converts a bpp value and RGBA masks to an enumerated pixel format.
// (https://wiki.libsdl.org/SDL_MasksToPixelFormatEnum)
func MasksToPixelFormatEnum(bpp int, rmask, gmask, bmask, amask uint32) uint {
	return uint(C.SDL_MasksToPixelFormatEnum(C.int(bpp), C.Uint32(rmask), C.Uint32(gmask),
		C.Uint32(bmask), C.Uint32(amask)))
}

// AllocFormat creates a PixelFormat structure corresponding to a pixel format.
// (https://wiki.libsdl.org/SDL_AllocFormat)
func AllocFormat(format uint) (*PixelFormat, error) {
	r := (*PixelFormat)(unsafe.Pointer(C.SDL_AllocFormat(C.Uint32(format))))
	if r == nil {
		return nil, GetError()
	}
	return r, nil
}

// Free frees the PixelFormat structure allocated by AllocFormat().
// (https://wiki.libsdl.org/SDL_FreeFormat)
func (format *PixelFormat) Free() {
	C.SDL_FreeFormat((*C.SDL_PixelFormat)(unsafe.Pointer(format)))
}

// AllocPalette creates a palette structure with the specified number of color entries.
// (https://wiki.libsdl.org/SDL_AllocPalette)
func AllocPalette(ncolors int) (*Palette, error) {
	r := (*Palette)(unsafe.Pointer(C.SDL_AllocPalette(C.int(ncolors))))
	if r == nil {
		return nil, GetError()
	}
	return r, nil
}

// SetPalette sets the palette for the pixel format structure.
// (https://wiki.libsdl.org/SDL_SetPixelFormatPalette)
func (format *PixelFormat) SetPalette(palette *Palette) error {
	r := C.SDL_SetPixelFormatPalette((*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		(*C.SDL_Palette)(unsafe.Pointer(palette)))
	if r != 0 {
		return GetError()
	}
	return nil
}

// SetColors sets a range of colors in the palette.
// (https://wiki.libsdl.org/SDL_SetPaletteColors)
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

// Free frees the palette created with AllocPalette().
// (https://wiki.libsdl.org/SDL_FreePalette)
func (palette *Palette) Free() {
	C.SDL_FreePalette((*C.SDL_Palette)(unsafe.Pointer(palette)))
}

// MapRGB maps an RGB triple to an opaque pixel value for a given pixel format.
// (https://wiki.libsdl.org/SDL_MapRGB)
func MapRGB(format *PixelFormat, r, g, b uint8) uint32 {
	return uint32(C.SDL_MapRGB((*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		C.Uint8(r), C.Uint8(g), C.Uint8(b)))
}

// MapRGBA maps an RGBA quadruple to a pixel value for a given pixel format.
// (https://wiki.libsdl.org/SDL_MapRGBA)
func MapRGBA(format *PixelFormat, r, g, b, a uint8) uint32 {
	return uint32(C.SDL_MapRGBA((*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		C.Uint8(r), C.Uint8(g), C.Uint8(b), C.Uint8(a)))
}

// GetRGB returns RGB values from a pixel in the specified format.
// (https://wiki.libsdl.org/SDL_GetRGB)
func GetRGB(pixel uint32, format *PixelFormat) (r, g, b uint8) {
	C.SDL_GetRGB(C.Uint32(pixel), (*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		(*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b))
	return
}

// GetRGBA returns RGBA values from a pixel in the specified format.
// (https://wiki.libsdl.org/SDL_GetRGBA)
func GetRGBA(pixel uint32, format *PixelFormat) (r, g, b, a uint8) {
	C.SDL_GetRGBA(C.Uint32(pixel), (*C.SDL_PixelFormat)(unsafe.Pointer(format)),
		(*C.Uint8)(&r), (*C.Uint8)(&g), (*C.Uint8)(&b), (*C.Uint8)(&a))
	return
}

// CalculateGammaRamp calculates a 256 entry gamma ramp for a gamma value.
// (https://wiki.libsdl.org/SDL_CalculateGammaRamp)
func CalculateGammaRamp(gamma float32, ramp *[256]uint16) {
	C.SDL_CalculateGammaRamp(C.float(gamma), (*C.Uint16)(unsafe.Pointer(&ramp[0])))
}
