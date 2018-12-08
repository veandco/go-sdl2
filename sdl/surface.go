package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,8))
typedef enum {
	SDL_YUV_CONVERSION_JPEG,
	SDL_YUV_CONVERSION_BT601,
	SDL_YUV_CONVERSION_BT709,
	SDL_YUV_CONVERSION_AUTOMATIC
} SDL_YUV_CONVERSION_MODE;


#if defined(WARN_OUTDATED)
#pragma message("SDL_SetYUVConversionMode is not supported before SDL 2.0.8")
#endif

void SDL_SetYUVConversionMode(SDL_YUV_CONVERSION_MODE mode)
{
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_GetYUVConversionMode is not supported before SDL 2.0.8")
#endif

SDL_YUV_CONVERSION_MODE SDL_GetYUVConversionMode(void)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_GetYUVConversionModeForResolution is not supported before SDL 2.0.8")
#endif

SDL_YUV_CONVERSION_MODE SDL_GetYUVConversionModeForResolution(int width, int height)
{
	return -1;
}
#endif

#if !(SDL_VERSION_ATLEAST(2,0,6))

#if defined(WARN_OUTDATED)
#pragma message("SDL_DuplicateSurface is not supported before SDL 2.0.6")
#endif

static inline SDL_Surface* SDL_DuplicateSurface(SDL_Surface *surface)
{
	return NULL;
}
#endif

#if !(SDL_VERSION_ATLEAST(2,0,5))

#if defined(WARN_OUTDATED)
#pragma message("SDL_CreateRGBSurfaceWithFormat is not supported before SDL 2.0.5")
#endif

static inline SDL_Surface* SDL_CreateRGBSurfaceWithFormat(Uint32 flags, int width, int height, int depth, Uint32 format)
{
	return NULL;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_CreateRGBSurfaceWithFormatFrom is not supported before SDL 2.0.5")
#endif

static inline SDL_Surface* SDL_CreateRGBSurfaceWithFormatFrom(void* pixels, int width, int height, int depth, int pitch, Uint32 format)
{
	return NULL;
}
#endif
*/
import "C"
import "unsafe"
import "reflect"
import "image"
import "image/color"

// Surface flags (internal use)
const (
	SWSURFACE = C.SDL_SWSURFACE // just here for compatibility
	PREALLOC  = C.SDL_PREALLOC  // surface uses preallocated memory
	RLEACCEL  = C.SDL_RLEACCEL  // surface is RLE encoded
	DONTFREE  = C.SDL_DONTFREE  // surface is referenced internally
)

type YUV_CONVERSION_MODE C.SDL_YUV_CONVERSION_MODE

// YUV Conversion Modes
const (
	YUV_CONVERSION_JPEG      YUV_CONVERSION_MODE = C.SDL_YUV_CONVERSION_JPEG      // Full range JPEG
	YUV_CONVERSION_BT601                         = C.SDL_YUV_CONVERSION_BT601     // BT.601 (the default)
	YUV_CONVERSION_BT709                         = C.SDL_YUV_CONVERSION_BT709     // BT.709
	YUV_CONVERSION_AUTOMATIC                     = C.SDL_YUV_CONVERSION_AUTOMATIC // BT.601 for SD content, BT.709 for HD content
)

// Surface contains a collection of pixels used in software blitting.
// (https://wiki.libsdl.org/SDL_Surface)
type Surface struct {
	flags    uint32         // (internal use)
	Format   *PixelFormat   // the format of the pixels stored in the surface (read-only) (https://wiki.libsdl.org/SDL_PixelFormat)
	W        int32          // the width in pixels (read-only)
	H        int32          // the height in pixels (read-only)
	Pitch    int32          // the length of a row of pixels in bytes (read-only)
	pixels   unsafe.Pointer // the pointer to the actual pixel data; use Pixels() for access
	UserData unsafe.Pointer // an arbitrary pointer you can set
	locked   int32          // used for surfaces that require locking (internal use)
	lockData unsafe.Pointer // used for surfaces that require locking (internal use)
	ClipRect Rect           // a Rect structure used to clip blits to the surface which can be set by SetClipRect() (read-only)
	_        unsafe.Pointer // map; info for fast blit mapping to other surfaces (internal use)
	RefCount int32          // reference count that can be incremented by the application
}
type cSurface C.SDL_Surface

func (surface *Surface) cptr() *C.SDL_Surface {
	return (*C.SDL_Surface)(unsafe.Pointer(surface))
}

// MustLock reports whether the surface must be locked for access.
// (https://wiki.libsdl.org/SDL_MUSTLOCK)
func (surface *Surface) MustLock() bool {
	return (surface.flags & RLEACCEL) != 0
}

// CreateRGBSurface allocates a new RGB surface.
// (https://wiki.libsdl.org/SDL_CreateRGBSurface)
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

// CreateRGBSurfaceFrom allocate a new RGB surface with existing pixel data.
// (https://wiki.libsdl.org/SDL_CreateRGBSurfaceFrom)
func CreateRGBSurfaceFrom(pixels unsafe.Pointer, width, height int32, depth, pitch int, Rmask, Gmask, Bmask, Amask uint32) (*Surface, error) {
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

// CreateRGBSurfaceWithFormat allocates an RGB surface.
// (https://wiki.libsdl.org/SDL_CreateRGBSurfaceWithFormat)
func CreateRGBSurfaceWithFormat(flags uint32, width, height, depth int32, format uint32) (*Surface, error) {
	surface := (*Surface)(unsafe.Pointer(C.SDL_CreateRGBSurfaceWithFormat(
		C.Uint32(flags),
		C.int(width),
		C.int(height),
		C.int(depth),
		C.Uint32(format))))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// CreateRGBSurfaceWithFormatFrom allocates an RGB surface from provided pixel data.
// (https://wiki.libsdl.org/SDL_CreateRGBSurfaceWithFormatFrom)
func CreateRGBSurfaceWithFormatFrom(pixels unsafe.Pointer, width, height, depth, pitch int32, format uint32) (*Surface, error) {
	surface := (*Surface)(unsafe.Pointer(C.SDL_CreateRGBSurfaceWithFormatFrom(
		pixels,
		C.int(width),
		C.int(height),
		C.int(depth),
		C.int(pitch),
		C.Uint32(format))))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// SetYUVConversionMode sets the YUV conversion mode
// TODO: (https://wiki.libsdl.org/SDL_SetYUVConversionMode)
func SetYUVConversionMode(mode YUV_CONVERSION_MODE) {
	_mode := C.SDL_YUV_CONVERSION_MODE(mode)
	C.SDL_SetYUVConversionMode(_mode)
}

// GetYUVConversionMode gets the YUV conversion mode
// TODO: (https://wiki.libsdl.org/SDL_GetYUVConversionMode)
func GetYUVConversionMode() YUV_CONVERSION_MODE {
	return YUV_CONVERSION_MODE(C.SDL_GetYUVConversionMode())
}

// GetYUVConversionModeForResolution gets the YUV conversion mode
// TODO: (https://wiki.libsdl.org/SDL_GetYUVConversionModeForResolution)
func GetYUVConversionModeForResolution(width, height int) YUV_CONVERSION_MODE {
	_width := C.int(width)
	_height := C.int(height)
	return YUV_CONVERSION_MODE(C.SDL_GetYUVConversionModeForResolution(_width, _height))
}

// Free frees the RGB surface.
// (https://wiki.libsdl.org/SDL_FreeSurface)
func (surface *Surface) Free() {
	C.SDL_FreeSurface(surface.cptr())
}

// SetPalette sets the palette used by the surface.
// (https://wiki.libsdl.org/SDL_SetSurfacePalette)
func (surface *Surface) SetPalette(palette *Palette) error {
	if C.SDL_SetSurfacePalette(surface.cptr(), palette.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Lock sets up the surface for directly accessing the pixels.
// (https://wiki.libsdl.org/SDL_LockSurface)
func (surface *Surface) Lock() error {
	if C.SDL_LockSurface(surface.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Unlock releases the surface after directly accessing the pixels.
// (https://wiki.libsdl.org/SDL_UnlockSurface)
func (surface *Surface) Unlock() {
	C.SDL_UnlockSurface(surface.cptr())
}

// LoadBMPRW loads a BMP image from a seekable SDL data stream (memory or file).
// (https://wiki.libsdl.org/SDL_LoadBMP_RW)
func LoadBMPRW(src *RWops, freeSrc bool) (*Surface, error) {
	surface := (*Surface)(unsafe.Pointer(C.SDL_LoadBMP_RW(src.cptr(), C.int(Btoi(freeSrc)))))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// LoadBMP loads a surface from a BMP file.
// (https://wiki.libsdl.org/SDL_LoadBMP)
func LoadBMP(file string) (*Surface, error) {
	return LoadBMPRW(RWFromFile(file, "rb"), true)
}

// SaveBMPRW save the surface to a seekable SDL data stream (memory or file) in BMP format.
// (https://wiki.libsdl.org/SDL_SaveBMP_RW)
func (surface *Surface) SaveBMPRW(dst *RWops, freeDst bool) error {
	if C.SDL_SaveBMP_RW(surface.cptr(), dst.cptr(), C.int(Btoi(freeDst))) != 0 {
		return GetError()
	}
	return nil
}

// SaveBMP saves the surface to a BMP file.
// (https://wiki.libsdl.org/SDL_SaveBMP)
func (surface *Surface) SaveBMP(file string) error {
	return surface.SaveBMPRW(RWFromFile(file, "wb"), true)
}

// SetRLE sets the RLE acceleration hint for the surface.
// (https://wiki.libsdl.org/SDL_SetSurfaceRLE)
func (surface *Surface) SetRLE(flag bool) error {
	if C.SDL_SetSurfaceRLE(surface.cptr(), C.int(Btoi(flag))) != 0 {
		return GetError()
	}
	return nil
}

// SetColorKey sets the color key (transparent pixel) in the surface.
// (https://wiki.libsdl.org/SDL_SetColorKey)
func (surface *Surface) SetColorKey(flag bool, key uint32) error {
	if C.SDL_SetColorKey(surface.cptr(), C.int(Btoi(flag)), C.Uint32(key)) != 0 {
		return GetError()
	}
	return nil
}

// GetColorKey retruns the color key (transparent pixel) for the surface.
// (https://wiki.libsdl.org/SDL_GetColorKey)
func (surface *Surface) GetColorKey() (key uint32, err error) {
	_key := (*C.Uint32)(unsafe.Pointer(&key))
	if C.SDL_GetColorKey(surface.cptr(), _key) != 0 {
		return key, GetError()
	}
	return key, nil
}

// SetColorMod sets an additional color value multiplied into blit operations.
// (https://wiki.libsdl.org/SDL_SetSurfaceColorMod)
func (surface *Surface) SetColorMod(r, g, b uint8) error {
	if C.SDL_SetSurfaceColorMod(surface.cptr(), C.Uint8(r), C.Uint8(g), C.Uint8(b)) != 0 {
		return GetError()
	}
	return nil
}

// GetColorMod returns the additional color value multiplied into blit operations.
// (https://wiki.libsdl.org/SDL_GetSurfaceColorMod)
func (surface *Surface) GetColorMod() (r, g, b uint8, err error) {
	_r := (*C.Uint8)(unsafe.Pointer(&r))
	_g := (*C.Uint8)(unsafe.Pointer(&g))
	_b := (*C.Uint8)(unsafe.Pointer(&b))
	if C.SDL_GetSurfaceColorMod(surface.cptr(), _r, _g, _b) != 0 {
		return r, g, b, GetError()
	}
	return r, g, b, nil
}

// SetAlphaMod sets an additional alpha value used in blit operations.
// (https://wiki.libsdl.org/SDL_SetSurfaceAlphaMod)
func (surface *Surface) SetAlphaMod(alpha uint8) error {
	if C.SDL_SetSurfaceAlphaMod(surface.cptr(), C.Uint8(alpha)) != 0 {
		return GetError()
	}
	return nil
}

// GetAlphaMod returns the additional alpha value used in blit operations.
// (https://wiki.libsdl.org/SDL_GetSurfaceAlphaMod)
func (surface *Surface) GetAlphaMod() (alpha uint8, err error) {
	_alpha := (*C.Uint8)(unsafe.Pointer(&alpha))
	if C.SDL_GetSurfaceAlphaMod(surface.cptr(), _alpha) != 0 {
		return alpha, GetError()
	}
	return alpha, nil
}

// SetBlendMode sets the blend mode used for blit operations.
// (https://wiki.libsdl.org/SDL_SetSurfaceBlendMode)
func (surface *Surface) SetBlendMode(bm BlendMode) error {
	if C.SDL_SetSurfaceBlendMode(surface.cptr(), bm.c()) != 0 {
		return GetError()
	}
	return nil
}

// GetBlendMode returns the blend mode used for blit operations.
// (https://wiki.libsdl.org/SDL_GetSurfaceBlendMode)
func (surface *Surface) GetBlendMode() (bm BlendMode, err error) {
	if C.SDL_GetSurfaceBlendMode(surface.cptr(), bm.cptr()) != 0 {
		return bm, GetError()
	}
	return bm, nil
}

// SetClipRect sets the clipping rectangle for the surface
// (https://wiki.libsdl.org/SDL_SetClipRect)
func (surface *Surface) SetClipRect(rect *Rect) bool {
	return C.SDL_SetClipRect(surface.cptr(), rect.cptr()) > 0
}

// GetClipRect returns the clipping rectangle for a surface.
// (https://wiki.libsdl.org/SDL_GetClipRect)
func (surface *Surface) GetClipRect(rect *Rect) {
	C.SDL_GetClipRect(surface.cptr(), rect.cptr())
}

// Convert copies the existing surface into a new one that is optimized for blitting to a surface of a specified pixel format.
// (https://wiki.libsdl.org/SDL_ConvertSurface)
func (surface *Surface) Convert(fmt *PixelFormat, flags uint32) (*Surface, error) {
	_surface := (*Surface)(unsafe.Pointer(C.SDL_ConvertSurface(surface.cptr(), fmt.cptr(), C.Uint32(flags))))
	if _surface == nil {
		return nil, GetError()
	}
	return _surface, nil
}

// ConvertFormat copies the existing surface to a new surface of the specified format.
// (https://wiki.libsdl.org/SDL_ConvertSurfaceFormat)
func (surface *Surface) ConvertFormat(pixelFormat uint32, flags uint32) (*Surface, error) {
	_surface := (*Surface)(unsafe.Pointer(C.SDL_ConvertSurfaceFormat(surface.cptr(), C.Uint32(pixelFormat), C.Uint32(flags))))
	if _surface == nil {
		return nil, GetError()
	}
	return _surface, nil
}

// ConvertPixels copies a block of pixels of one format to another format.
// (https://wiki.libsdl.org/SDL_ConvertPixels)
func ConvertPixels(width, height int32, srcFormat uint32, src unsafe.Pointer, srcPitch int,
	dstFormat uint32, dst unsafe.Pointer, dstPitch int) error {
	if C.SDL_ConvertPixels(C.int(width), C.int(height), C.Uint32(srcFormat), src, C.int(srcPitch), C.Uint32(dstFormat), dst, C.int(dstPitch)) != 0 {
		return GetError()
	}
	return nil
}

// FillRect performs a fast fill of a rectangle with a specific color.
// (https://wiki.libsdl.org/SDL_FillRect)
func (surface *Surface) FillRect(rect *Rect, color uint32) error {
	if C.SDL_FillRect(surface.cptr(), rect.cptr(), C.Uint32(color)) != 0 {
		return GetError()
	}
	return nil
}

// FillRects performs a fast fill of a set of rectangles with a specific color.
// (https://wiki.libsdl.org/SDL_FillRects)
func (surface *Surface) FillRects(rects []Rect, color uint32) error {
	if C.SDL_FillRects(surface.cptr(), rects[0].cptr(), C.int(len(rects)), C.Uint32(color)) != 0 {
		return GetError()
	}
	return nil
}

// Blit performs a fast surface copy to a destination surface.
// (https://wiki.libsdl.org/SDL_BlitSurface)
func (surface *Surface) Blit(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	if C.SDL_BlitSurface(surface.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// BlitScaled performs a scaled surface copy to a destination surface.
// (https://wiki.libsdl.org/SDL_BlitScaled)
func (surface *Surface) BlitScaled(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	if C.SDL_BlitScaled(surface.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// UpperBlit has been replaced by Blit().
// (https://wiki.libsdl.org/SDL_UpperBlit)
func (surface *Surface) UpperBlit(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	if C.SDL_UpperBlit(surface.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// LowerBlit performs low-level surface blitting only.
// (https://wiki.libsdl.org/SDL_LowerBlit)
func (surface *Surface) LowerBlit(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	if C.SDL_LowerBlit(surface.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// SoftStretch has been replaced by BlitScaled().
// (https://wiki.libsdl.org/SDL_SoftStretch)
func (surface *Surface) SoftStretch(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	if C.SDL_SoftStretch(surface.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// UpperBlitScaled has been replaced by BlitScaled().
// (https://wiki.libsdl.org/SDL_UpperBlitScaled)
func (surface *Surface) UpperBlitScaled(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	if C.SDL_UpperBlitScaled(surface.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// LowerBlitScaled performs low-level surface scaled blitting only.
// (https://wiki.libsdl.org/SDL_LowerBlitScaled)
func (surface *Surface) LowerBlitScaled(srcRect *Rect, dst *Surface, dstRect *Rect) error {
	if C.SDL_LowerBlitScaled(surface.cptr(), srcRect.cptr(), dst.cptr(), dstRect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// PixelNum returns the number of pixels stored in the surface.
func (surface *Surface) PixelNum() int {
	return int(surface.W * surface.H)
}

// BytesPerPixel return the number of significant bits in a pixel values of the surface.
func (surface *Surface) BytesPerPixel() int {
	return int(surface.Format.BytesPerPixel)
}

// Pixels returns the actual pixel data of the surface.
func (surface *Surface) Pixels() []byte {
	var b []byte
	length := int(surface.W*surface.H) * int(surface.Format.BytesPerPixel)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(surface.pixels)
	return b
}

// Data returns the pointer to the actual pixel data of the surface.
func (surface *Surface) Data() unsafe.Pointer {
	return surface.pixels
}

// Duplicate creates a new surface identical to the existing surface
func (surface *Surface) Duplicate() (newSurface *Surface, err error) {
	_newSurface := C.SDL_DuplicateSurface(surface.cptr())
	if _newSurface == nil {
		err = GetError()
		return
	}

	newSurface = (*Surface)(unsafe.Pointer(_newSurface))
	return
}

// ColorModel returns the color model used by this Surface.
func (surface *Surface) ColorModel() color.Model {
	switch surface.Format.Format {
	case PIXELFORMAT_ARGB8888, PIXELFORMAT_ABGR8888:
		return color.RGBAModel
	case PIXELFORMAT_RGB888:
		return color.RGBAModel
	default:
		panic("Not implemented yet")
	}
}

// Bounds return the bounds of this surface. Currently, it always starts at
// (0,0), but this is not guaranteed in the future so don't rely on it.
func (surface *Surface) Bounds() image.Rectangle {
	return image.Rect(0, 0, int(surface.W), int(surface.H))
}

// At returns the pixel color at (x, y)
func (surface *Surface) At(x, y int) color.Color {
	pix := surface.Pixels()
	i := int32(y)*surface.Pitch + int32(x)*int32(surface.Format.BytesPerPixel)
	switch surface.Format.Format {
	/*
		case PIXELFORMAT_ARGB8888:
			return color.RGBA{pix[i+3], pix[i], pix[i+1], pix[i+2]}
		case PIXELFORMAT_ABGR8888:
			return color.RGBA{pix[i], pix[i+3], pix[i+2], pix[i+1]}
	*/
	case PIXELFORMAT_RGB888:
		return color.RGBA{pix[i], pix[i+1], pix[i+2], 0xff}
	default:
		panic("Not implemented yet")
	}
}

// Set the color of the pixel at (x, y) using this surface's color model to
// convert c to the appropriate color. This method is required for the
// draw.Image interface. The surface may require locking before calling Set.
func (surface *Surface) Set(x, y int, c color.Color) {
	pix := surface.Pixels()
	i := int32(y)*surface.Pitch + int32(x)*int32(surface.Format.BytesPerPixel)
	switch surface.Format.Format {
	case PIXELFORMAT_ARGB8888:
		col := surface.ColorModel().Convert(c).(color.RGBA)
		pix[i+0] = col.R
		pix[i+1] = col.G
		pix[i+2] = col.B
		pix[i+3] = col.A
	case PIXELFORMAT_ABGR8888:
		col := surface.ColorModel().Convert(c).(color.RGBA)
		pix[i+3] = col.R
		pix[i+2] = col.G
		pix[i+1] = col.B
		pix[i+0] = col.A
	case PIXELFORMAT_RGB24, PIXELFORMAT_RGB888:
		col := surface.ColorModel().Convert(c).(color.RGBA)
		pix[i+0] = col.B
		pix[i+1] = col.G
		pix[i+2] = col.R
	case PIXELFORMAT_BGR24, PIXELFORMAT_BGR888:
		col := surface.ColorModel().Convert(c).(color.RGBA)
		pix[i+2] = col.R
		pix[i+1] = col.G
		pix[i+0] = col.B
	default:
		panic("Unknown pixel format!")
	}
}
