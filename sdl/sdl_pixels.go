package sdl

// #include <SDL2/SDL.h>
import "C"

type PixelFormat struct {
	Format uint32
	Palette *Palette
	BitsPerPixels uint8
	BytesPerPixel uint8
	Padding [2]uint8
	Rmask uint32
	Gmask uint32
	Bmask uint32
	Amask uint32
	Rloss uint8
	Gloss uint8
	Bloss uint8
	Aloss uint8
	Rshift uint8
	Gshift uint8
	Bshift uint8
	Ashift uint8
	RefCount int
	next *PixelFormat
}

type Palette struct {
	Ncolors int
	Colors *Color
	Version uint32
	RefCount int
}

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func (c Color) Uint32() uint32 {
	var v uint32
	v |= uint32(c.R) << 16
	v |= uint32(c.G) << 8
	v |= uint32(c.B)
	return v
}

const (
    PIXELTYPE_UNKNOWN = iota
    PIXELTYPE_INDEX1
    PIXELTYPE_INDEX4
    PIXELTYPE_INDEX8
    PIXELTYPE_PACKED8
    PIXELTYPE_PACKED16
    PIXELTYPE_PACKED32
    PIXELTYPE_ARRAYU8
    PIXELTYPE_ARRAYU16
    PIXELTYPE_ARRAYU32
    PIXELTYPE_ARRAYF16
    PIXELTYPE_ARRAYF32
)

/** Bitmap pixel order high bit -> low bit. */
const (
    BITMAPORDER_NONE = iota
    BITMAPORDER_4321
    BITMAPORDER_1234
)

/** Packed component order high bit -> low bit. */
const (
    PACKEDORDER_NONE = iota
    PACKEDORDER_XRGB
    PACKEDORDER_RGBX
    PACKEDORDER_ARGB
    PACKEDORDER_RGBA
    PACKEDORDER_XBGR
    PACKEDORDER_BGRX
    PACKEDORDER_ABGR
    PACKEDORDER_BGRA
)

/** Array component order low byte -> high byte. */
const (
    ARRAYORDER_NONE = iota
    ARRAYORDER_RGB
    ARRAYORDER_RGBA
    ARRAYORDER_ARGB
    ARRAYORDER_BGR
    ARRAYORDER_BGRA
    ARRAYORDER_ABGR
)

/** Packed component layout. */
const (
    PACKEDLAYOUT_NONE = iota
    PACKEDLAYOUT_332
    PACKEDLAYOUT_4444
    PACKEDLAYOUT_1555
    PACKEDLAYOUT_5551
    PACKEDLAYOUT_565
    PACKEDLAYOUT_8888
    PACKEDLAYOUT_2101010
    PACKEDLAYOUT_1010102
)

const (
    PIXELFORMAT_UNKNOWN = iota
    PIXELFORMAT_ARGB8888 = ((1 << 28) | ((PIXELTYPE_PACKED32) << 24) | ((PACKEDORDER_ARGB) << 20) | ((PACKEDLAYOUT_8888) << 16) | ((32) << 8) | ((4) << 0))
)
