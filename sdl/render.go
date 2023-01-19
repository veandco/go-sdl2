package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,1))

#if defined(WARN_OUTDATED)
#pragma message("SDL_UpdateYUVTexture is not supported before SDL 2.0.1")
#endif

static inline int SDL_UpdateYUVTexture(SDL_Texture* texture, const SDL_Rect* rect, const Uint8* Yplane, int Ypitch, const Uint8* Uplane, int Upitch, const Uint8* Vplane, int Vpitch)
{
	return -1;
}
#endif

#if !(SDL_VERSION_ATLEAST(2,0,4))
#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderIsClipEnabled is not supported before SDL 2.0.4")
#endif
static inline SDL_bool SDLCALL SDL_RenderIsClipEnabled(SDL_Renderer * renderer)
{
	return SDL_FALSE;
}
#endif

#if !(SDL_VERSION_ATLEAST(2,0,5))

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderSetIntegerScale is not supported before SDL 2.0.5")
#endif

static inline int SDL_RenderSetIntegerScale(SDL_Renderer* renderer, SDL_bool enable)
{
	SDL_Unsupported();
	return -1;
}

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderGetIntegerScale is not supported before SDL 2.0.5")
#endif

static inline SDL_bool SDL_RenderGetIntegerScale(SDL_Renderer* renderer)
{
	SDL_Unsupported();
	return -1;
}
#endif

#if !(SDL_VERSION_ATLEAST(2,0,8))


#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderGetMetalLayer is not supported before SDL 2.0.8")
#endif

static inline void * SDL_RenderGetMetalLayer(SDL_Renderer *renderer)
{
	return NULL;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderGetMetalCommandEncoder is not supported before SDL 2.0.8")
#endif

static inline void * SDL_RenderGetMetalCommandEncoder(SDL_Renderer *renderer)
{
	return NULL;
}

#endif

#if !(SDL_VERSION_ATLEAST(2,0,10))

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderDrawPointF is not supported before SDL 2.0.10")
#endif

static inline int SDL_RenderDrawPointF(SDL_Renderer * renderer, float x, float y)
{
	return SDL_RenderDrawPoint(renderer, (int) x, (int) y);
}

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderDrawPointsF is not supported before SDL 2.0.10")
#endif

static inline int SDL_RenderDrawPointsF(SDL_Renderer * renderer, const SDL_FPoint * points, int count)
{
	return SDL_RenderDrawPoints(renderer, (const SDL_Point *) points, count);
}

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderDrawLineF is not supported before SDL 2.0.10")
#endif

static inline int SDL_RenderDrawLineF(SDL_Renderer * renderer, float x1, float y1, float x2, float y2)
{
	return SDL_RenderDrawLine(renderer, (int) x1, (int) y1, (int) x2, (int) y2);
}

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderDrawLinesF is not supported before SDL 2.0.10")
#endif

static inline int SDL_RenderDrawLinesF(SDL_Renderer * renderer, const SDL_FPoint * points, int count)
{
	return SDL_RenderDrawLines(renderer, (const SDL_Point *) points, count);
}

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderDrawRectF is not supported before SDL 2.0.10")
#endif

static inline int SDL_RenderDrawRectF(SDL_Renderer * renderer, const SDL_FRect * rect)
{
	return SDL_RenderDrawRect(renderer, (const SDL_Rect *) rect);
}

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderDrawRectsF is not supported before SDL 2.0.10")
#endif

static inline int SDL_RenderDrawRectsF(SDL_Renderer * renderer, const SDL_FRect *rects, int count)
{
	return SDL_RenderDrawRects(renderer, (const SDL_Rect *) rects, count);
}

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderFillRectF is not supported before SDL 2.0.10")
#endif

static inline int SDL_RenderFillRectF(SDL_Renderer * renderer, const SDL_FRect * rect)
{
	return SDL_RenderFillRect(renderer, (const SDL_Rect *) rect);
}

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderFillRectsF is not supported before SDL 2.0.10")
#endif

static inline int SDL_RenderFillRectsF(SDL_Renderer * renderer, const SDL_FRect * rects, int count)
{
	return SDL_RenderFillRects(renderer, (const SDL_Rect *) rects, count);
}

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderCopyF is not supported before SDL 2.0.10")
#endif

static inline int SDL_RenderCopyF(SDL_Renderer * renderer, SDL_Texture * texture, const SDL_Rect * srcrect, const SDL_FRect * dstrect)
{
	return SDL_RenderCopy(renderer, texture, srcrect, (const SDL_Rect *) dstrect);
}

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderCopyExF is not supported before SDL 2.0.10")
#endif

static inline int SDL_RenderCopyExF(SDL_Renderer * renderer, SDL_Texture * texture, const SDL_Rect * srcrect, const SDL_FRect * dstrect, const double angle, const SDL_FPoint * center, const SDL_RendererFlip flip)
{
	return SDL_RenderCopyEx(renderer, texture, srcrect, (const SDL_Rect *) dstrect, angle, (const SDL_Point *) center, flip);
}

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderFlush is not supported before SDL 2.0.10")
#endif

static inline int SDL_RenderFlush(SDL_Renderer * renderer)
{
	return 0;
}

#endif

#if !(SDL_VERSION_ATLEAST(2,0,12))

typedef enum
{
    SDL_ScaleModeNearest,
    SDL_ScaleModeLinear,
    SDL_ScaleModeBest
} SDL_ScaleMode;

#if defined(WARN_OUTDATED)
#pragma message("SDL_SetTextureScaleMode is not supported before SDL 2.0.12")
#pragma message("SDL_GetTextureScaleMode is not supported before SDL 2.0.12")
#pragma message("SDL_LockTextureToSurface is not supported before SDL 2.0.12")
#endif

static int SDL_SetTextureScaleMode(SDL_Texture * texture, SDL_ScaleMode scaleMode)
{
	return -1;
}

static int SDLCALL SDL_GetTextureScaleMode(SDL_Texture * texture, SDL_ScaleMode *scaleMode)
{
	return -1;
}

static int SDL_LockTextureToSurface(SDL_Texture *texture, const SDL_Rect *rect, SDL_Surface **surface)
{
	return -1;
}
#endif


#if !(SDL_VERSION_ATLEAST(2,0,16))

#if defined(WARN_OUTDATED)
#pragma message("SDL_UpdateNVTexture is not supported before SDL 2.0.16")
#endif

static int SDL_UpdateNVTexture(SDL_Texture * texture, const SDL_Rect * rect, const Uint8 *Yplane, int Ypitch, const Uint8 *UVplane, int UVpitch)
{
	return -1;
}

#endif


#if !(SDL_VERSION_ATLEAST(2,0,18))

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderGeometry is not supported before SDL 2.0.18")
#pragma message("SDL_RenderGeometryRaw is not supported before SDL 2.0.18")
#pragma message("SDL_SetTextureUserData is not supported before SDL 2.0.18")
#pragma message("SDL_GetTextureUserData is not supported before SDL 2.0.18")
#pragma message("SDL_RenderWindowToLogical is not supported before SDL 2.0.18")
#pragma message("SDL_RenderLogicalToWindow is not supported before SDL 2.0.18")
#pragma message("SDL_RenderSetVSync is not supported before SDL 2.0.18")
#endif

// Vertex structure
typedef struct SDL_Vertex
{
    SDL_FPoint position;  // Vertex position, in SDL_Renderer coordinates
    SDL_Color  color;     // Vertex color
    SDL_FPoint tex_coord; // Normalized texture coordinates, if needed
} SDL_Vertex;

static int SDL_RenderGeometry(SDL_Renderer *renderer, SDL_Texture *texture, const SDL_Vertex *vertices, int num_vertices, const int *indices, int num_indices)
{
	return -1;
}

static int SDL_RenderGeometryRaw(SDL_Renderer *renderer, SDL_Texture *texture, const float *xy, int xy_stride, const SDL_Color *color, int color_stride, const float *uv, int uv_stride, int num_vertices, const void *indices, int num_indices, int size_indices)
{
	return -1;
}

static int SDL_SetTextureUserData(SDL_Texture * texture, void *userdata)
{
	return -1;
}

static void * SDLCALL SDL_GetTextureUserData(SDL_Texture * texture)
{
	return NULL;
}

static void SDL_RenderWindowToLogical(SDL_Renderer * renderer, int windowX, int windowY, float *logicalX, float *logicalY)
{
}

static void SDL_RenderLogicalToWindow(SDL_Renderer * renderer, float logicalX, float logicalY, int *windowX, int *windowY)
{
}

static int SDL_RenderSetVSync(SDL_Renderer* renderer, int vsync)
{
	return -1;
}

#endif

#if SDL_COMPILEDVERSION == SDL_VERSIONNUM(2,0,18)
static inline int RenderGeometryRaw(SDL_Renderer *renderer, SDL_Texture *texture, const float *xy, int xy_stride, const SDL_Color *color, int color_stride, const float *uv, int uv_stride, int num_vertices, const void *indices, int num_indices, int size_indices)
{
	return SDL_RenderGeometryRaw(renderer, texture, xy, xy_stride, (int*) color, color_stride, uv, uv_stride, num_vertices, indices, num_indices, size_indices);
}
#else
static inline int RenderGeometryRaw(SDL_Renderer *renderer, SDL_Texture *texture, const float *xy, int xy_stride, const SDL_Color *color, int color_stride, const float *uv, int uv_stride, int num_vertices, const void *indices, int num_indices, int size_indices)
{
	return SDL_RenderGeometryRaw(renderer, texture, xy, xy_stride, color, color_stride, uv, uv_stride, num_vertices, indices, num_indices, size_indices);
}
#endif

#if !(SDL_VERSION_ATLEAST(2,0,22))

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderGetWindow is not supported before SDL 2.0.22")
#endif

static inline SDL_Window * SDLCALL SDL_RenderGetWindow(SDL_Renderer *renderer)
{
	return NULL;
}

#endif

// WORKAROUND: This prevents audio from seemingly going corrupt when drawing outside the screen bounding box?
// It does that by allocating SDL_Rect in the C context instead of Go context.
static inline int RenderCopy(SDL_Renderer *renderer, SDL_Texture *texture, SDL_Rect *src, int dst_x, int dst_y, int dst_w, int dst_h)
{
	SDL_Rect dst = {dst_x, dst_y, dst_w, dst_h};
	return SDL_RenderCopy(renderer, texture, src, &dst);
}
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// An enumeration of flags used when creating a rendering context.
// (https://wiki.libsdl.org/SDL_RendererFlags)
const (
	RENDERER_SOFTWARE      = C.SDL_RENDERER_SOFTWARE      // the renderer is a software fallback
	RENDERER_ACCELERATED   = C.SDL_RENDERER_ACCELERATED   // the renderer uses hardware acceleration
	RENDERER_PRESENTVSYNC  = C.SDL_RENDERER_PRESENTVSYNC  // present is synchronized with the refresh rate
	RENDERER_TARGETTEXTURE = C.SDL_RENDERER_TARGETTEXTURE // the renderer supports rendering to texture
)

type ScaleMode C.SDL_ScaleMode

// The scaling mode for a texture.
const (
	ScaleModeNearest ScaleMode = C.SDL_ScaleModeNearest // nearest pixel sampling
	ScaleModeLinear            = C.SDL_ScaleModeLinear  // linear filtering
	ScaleModeBest              = C.SDL_ScaleModeBest    // anisotropic filtering
)

// An enumeration of texture access patterns..
// (https://wiki.libsdl.org/SDL_TextureAccess)
const (
	TEXTUREACCESS_STATIC    = C.SDL_TEXTUREACCESS_STATIC    // changes rarely, not lockable
	TEXTUREACCESS_STREAMING = C.SDL_TEXTUREACCESS_STREAMING // changes frequently, lockable
	TEXTUREACCESS_TARGET    = C.SDL_TEXTUREACCESS_TARGET    // can be used as a render target
)

// An enumeration of the texture channel modulation used in Renderer.Copy().
// (https://wiki.libsdl.org/SDL_TextureModulate)
const (
	TEXTUREMODULATE_NONE  = C.SDL_TEXTUREMODULATE_NONE  // no modulation
	TEXTUREMODULATE_COLOR = C.SDL_TEXTUREMODULATE_COLOR // srcC = srcC * color
	TEXTUREMODULATE_ALPHA = C.SDL_TEXTUREMODULATE_ALPHA // srcA = srcA * alpha
)

// An enumeration of flags that can be used in the flip parameter for Renderer.CopyEx().
// (https://wiki.libsdl.org/SDL_RendererFlip)
const (
	FLIP_NONE       RendererFlip = C.SDL_FLIP_NONE       // do not flip
	FLIP_HORIZONTAL              = C.SDL_FLIP_HORIZONTAL // flip horizontally
	FLIP_VERTICAL                = C.SDL_FLIP_VERTICAL   // flip vertically
)

// RendererInfo contains information on the capabilities of a render driver or the current render context.
// (https://wiki.libsdl.org/SDL_RendererInfo)
type RendererInfo struct {
	Name string // the name of the renderer
	RendererInfoData
}

type cRendererInfo struct {
	Name *C.char
	RendererInfoData
}

// RendererInfoData contains information on the capabilities of a render driver or the current render context.
// (https://wiki.libsdl.org/SDL_RendererInfo)
type RendererInfoData struct {
	Flags             uint32    // a mask of supported renderer flags
	NumTextureFormats uint32    // the number of available texture formats
	TextureFormats    [16]int32 // the available texture formats
	MaxTextureWidth   int32     // the maximum texture width
	MaxTextureHeight  int32     // the maximum texture height
}

func (info *RendererInfo) cptr() *C.SDL_RendererInfo {
	return (*C.SDL_RendererInfo)(unsafe.Pointer(info))
}

func (info *cRendererInfo) cptr() *C.SDL_RendererInfo {
	return (*C.SDL_RendererInfo)(unsafe.Pointer(info))
}

// RendererFlip is an enumeration of flags that can be used in the flip parameter for Renderer.CopyEx().
// (https://wiki.libsdl.org/SDL_RendererFlip)
type RendererFlip uint32
type cRendererFlip C.SDL_RendererFlip

// Vertex structure
type Vertex struct {
	Position FPoint // Vertex position, in SDL_Renderer coordinates
	Color    Color  // Vertex color
	TexCoord FPoint // Normalized texture coordinates, if needed
}

func (flip RendererFlip) c() C.SDL_RendererFlip {
	return C.SDL_RendererFlip(flip)
}

// GetNumRenderDrivers returns the number of 2D rendering drivers available for the current display.
// (https://wiki.libsdl.org/SDL_GetNumRenderDrivers)
func GetNumRenderDrivers() (int, error) {
	i := int(C.SDL_GetNumRenderDrivers())
	return i, errorFromInt(i)
}

// GetRenderDriverInfo returns information about a specific 2D rendering driver for the current display.
// (https://wiki.libsdl.org/SDL_GetRenderDriverInfo)
func GetRenderDriverInfo(index int, info *RendererInfo) (int, error) {
	var cinfo cRendererInfo
	ret := int(C.SDL_GetRenderDriverInfo(C.int(index), cinfo.cptr()))
	if ret < 0 {
		return ret, GetError()
	}
	info.RendererInfoData = cinfo.RendererInfoData
	// No need to free, it's done by DestroyRenderer
	info.Name = C.GoString(cinfo.Name)
	return ret, nil
}

// CreateWindowAndRenderer returns a new window and default renderer.
// (https://wiki.libsdl.org/SDL_CreateWindowAndRenderer)
func CreateWindowAndRenderer(w, h int32, flags uint32) (*Window, *Renderer, error) {
	var window *C.SDL_Window
	var renderer *C.SDL_Renderer
	ret := C.SDL_CreateWindowAndRenderer(
		C.int(w),
		C.int(h),
		C.Uint32(flags),
		&window,
		&renderer)
	if ret == -1 {
		return nil, nil, GetError()
	}
	return (*Window)(unsafe.Pointer(window)), (*Renderer)(unsafe.Pointer(renderer)), nil
}

// CreateRenderer returns a new 2D rendering context for a window.
// (https://wiki.libsdl.org/SDL_CreateRenderer)
func CreateRenderer(window *Window, index int, flags uint32) (*Renderer, error) {
	renderer := C.SDL_CreateRenderer(window.cptr(), C.int(index), C.Uint32(flags))
	if renderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(unsafe.Pointer(renderer)), nil
}

// CreateSoftwareRenderer returns a new 2D software rendering context for a surface.
// (https://wiki.libsdl.org/SDL_CreateSoftwareRenderer)
func CreateSoftwareRenderer(surface *Surface) (*Renderer, error) {
	renderer := C.SDL_CreateSoftwareRenderer(surface.cptr())
	if renderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(unsafe.Pointer(renderer)), nil
}

// GetRenderer returns the renderer associated with a window.
// (https://wiki.libsdl.org/SDL_GetRenderer)
func (window *Window) GetRenderer() (*Renderer, error) {
	renderer := C.SDL_GetRenderer(window.cptr())
	if renderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(unsafe.Pointer(renderer)), nil
}

// GetInfo returns information about a rendering context.
// (https://wiki.libsdl.org/SDL_GetRendererInfo)
func (renderer *Renderer) GetInfo() (RendererInfo, error) {
	var cinfo cRendererInfo
	var info RendererInfo
	ret := int(C.SDL_GetRendererInfo(renderer.cptr(), cinfo.cptr()))
	if ret < 0 {
		return info, GetError()
	}
	info.RendererInfoData = cinfo.RendererInfoData
	// No need to free, it's done by DestroyRenderer
	info.Name = C.GoString(cinfo.Name)
	return info, nil
}

// GetOutputSize returns the output size in pixels of a rendering context.
// (https://wiki.libsdl.org/SDL_GetRendererOutputSize)
func (renderer *Renderer) GetOutputSize() (w, h int32, err error) {
	ret := C.SDL_GetRendererOutputSize(
		renderer.cptr(),
		(*C.int)(unsafe.Pointer(&w)),
		(*C.int)(unsafe.Pointer(&h)))
	err = errorFromInt(int(ret))
	return
}

// CreateTexture returns a new texture for a rendering context.
// (https://wiki.libsdl.org/SDL_CreateTexture)
func (renderer *Renderer) CreateTexture(format uint32, access int, w, h int32) (*Texture, error) {
	texture := C.SDL_CreateTexture(
		renderer.cptr(),
		C.Uint32(format),
		C.int(access),
		C.int(w),
		C.int(h))
	if texture == nil {
		return nil, GetError()
	}
	return (*Texture)(unsafe.Pointer(texture)), nil
}

// CreateTextureFromSurface returns a new texture from an existing surface.
// (https://wiki.libsdl.org/SDL_CreateTextureFromSurface)
func (renderer *Renderer) CreateTextureFromSurface(surface *Surface) (*Texture, error) {
	texture := C.SDL_CreateTextureFromSurface(renderer.cptr(), surface.cptr())
	if texture == nil {
		return nil, GetError()
	}
	return (*Texture)(unsafe.Pointer(texture)), nil
}

// Query returns the attributes of a texture.
// (https://wiki.libsdl.org/SDL_QueryTexture)
func (texture *Texture) Query() (format uint32, access int, width int32, height int32, err error) {
	var _format C.Uint32
	var _access C.int
	var _width C.int
	var _height C.int

	ret := int(C.SDL_QueryTexture(texture.cptr(), &_format, &_access, &_width, &_height))

	format = uint32(_format)
	access = int(_access)
	width = int32(_width)
	height = int32(_height)
	err = errorFromInt(ret)

	return
}

// SetColorMod sets an additional color value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL_SetTextureColorMod)
func (texture *Texture) SetColorMod(r uint8, g uint8, b uint8) error {
	return errorFromInt(int(
		C.SDL_SetTextureColorMod(
			texture.cptr(),
			C.Uint8(r),
			C.Uint8(g),
			C.Uint8(b))))
}

// SetAlphaMod sets an additional alpha value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL_SetTextureAlphaMod)
func (texture *Texture) SetAlphaMod(alpha uint8) error {
	return errorFromInt(int(
		C.SDL_SetTextureAlphaMod(texture.cptr(), C.Uint8(alpha))))
}

// GetAlphaMod returns the additional alpha value multiplied into render copy operations.
// (https://wiki.libsdl.org/SDL_GetTextureAlphaMod)
func (texture *Texture) GetAlphaMod() (alpha uint8, err error) {
	ret := C.SDL_GetTextureAlphaMod(texture.cptr(), (*C.Uint8)(unsafe.Pointer(&alpha)))
	return alpha, errorFromInt(int(ret))
}

// SetBlendMode sets the blend mode for a texture, used by Renderer.Copy().
// (https://wiki.libsdl.org/SDL_SetTextureBlendMode)
func (texture *Texture) SetBlendMode(bm BlendMode) error {
	return errorFromInt(int(
		C.SDL_SetTextureBlendMode(texture.cptr(), bm.c())))
}

// GetBlendMode returns the blend mode used for texture copy operations.
// (https://wiki.libsdl.org/SDL_GetTextureBlendMode)
func (texture *Texture) GetBlendMode() (bm BlendMode, err error) {
	ret := C.SDL_GetTextureBlendMode(texture.cptr(), bm.cptr())
	return bm, errorFromInt(int(ret))
}

// Update updates the given texture rectangle with new pixel data.
// (https://wiki.libsdl.org/SDL_UpdateTexture)
func (texture *Texture) Update(rect *Rect, pixels unsafe.Pointer, pitch int) error {
	if pixels == nil {
		return nil
	}
	return errorFromInt(int(
		C.SDL_UpdateTexture(
			texture.cptr(),
			rect.cptr(),
			pixels,
			C.int(pitch))))
}

// UpdateRGBA updates the given texture rectangle with new uint32 pixel data.
// (https://wiki.libsdl.org/SDL_UpdateTexture)
func (texture *Texture) UpdateRGBA(rect *Rect, pixels []uint32, pitch int) error {
	if pixels == nil {
		return nil
	}
	return errorFromInt(int(
		C.SDL_UpdateTexture(
			texture.cptr(),
			rect.cptr(),
			unsafe.Pointer(&pixels[0]),
			C.int(4*pitch)))) // 4 bytes in one uint32
}

// UpdateYUV updates a rectangle within a planar YV12 or IYUV texture with new pixel data.
// (https://wiki.libsdl.org/SDL_UpdateYUVTexture)
func (texture *Texture) UpdateYUV(rect *Rect, yPlane []byte, yPitch int, uPlane []byte, uPitch int, vPlane []byte, vPitch int) error {
	var yPlanePtr, uPlanePtr, vPlanePtr *byte
	if yPlane != nil {
		yPlanePtr = &yPlane[0]
	}
	if uPlane != nil {
		uPlanePtr = &uPlane[0]
	}
	if vPlane != nil {
		vPlanePtr = &vPlane[0]
	}
	return errorFromInt(int(
		C.SDL_UpdateYUVTexture(
			texture.cptr(),
			rect.cptr(),
			(*C.Uint8)(unsafe.Pointer(yPlanePtr)),
			C.int(yPitch),
			(*C.Uint8)(unsafe.Pointer(uPlanePtr)),
			C.int(uPitch),
			(*C.Uint8)(unsafe.Pointer(vPlanePtr)),
			C.int(vPitch))))
}

// Lock locks a portion of the texture for write-only pixel access.
// (https://wiki.libsdl.org/SDL_LockTexture)
func (texture *Texture) Lock(rect *Rect) ([]byte, int, error) {
	var _pitch C.int
	var _pixels unsafe.Pointer
	var b []byte
	var length int

	ret := C.SDL_LockTexture(texture.cptr(), rect.cptr(), &_pixels, &_pitch)
	if ret < 0 {
		return b, int(_pitch), GetError()
	}

	_, _, w, h, err := texture.Query()
	if err != nil {
		return b, int(_pitch), GetError()
	}

	pitch := int32(_pitch)
	if rect != nil {
		bytesPerPixel := pitch / w
		length = int(bytesPerPixel * (w*rect.H - rect.X - (w - rect.X - rect.W)))
	} else {
		length = int(pitch * h)
	}
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(_pixels)

	return b, int(pitch), nil
}

// Unlock unlocks a texture, uploading the changes to video memory, if needed.
// (https://wiki.libsdl.org/SDL_UnlockTexture)
func (texture *Texture) Unlock() {
	C.SDL_UnlockTexture(texture.cptr())
}

// RenderTargetSupported reports whether a window supports the use of render targets.
// (https://wiki.libsdl.org/SDL_RenderTargetSupported)
func (renderer *Renderer) RenderTargetSupported() bool {
	return C.SDL_RenderTargetSupported(renderer.cptr()) != 0
}

// SetRenderTarget sets a texture as the current rendering target.
// (https://wiki.libsdl.org/SDL_SetRenderTarget)
func (renderer *Renderer) SetRenderTarget(texture *Texture) error {
	return errorFromInt(int(
		C.SDL_SetRenderTarget(renderer.cptr(), texture.cptr())))

}

// GetRenderTarget returns the current render target.
// (https://wiki.libsdl.org/SDL_GetRenderTarget)
func (renderer *Renderer) GetRenderTarget() *Texture {
	return (*Texture)(unsafe.Pointer(C.SDL_GetRenderTarget(renderer.cptr())))
}

// SetLogicalSize sets a device independent resolution for rendering.
// (https://wiki.libsdl.org/SDL_RenderSetLogicalSize)
func (renderer *Renderer) SetLogicalSize(w, h int32) error {
	return errorFromInt(int(
		C.SDL_RenderSetLogicalSize(renderer.cptr(), C.int(w), C.int(h))))
}

// GetLogicalSize returns device independent resolution for rendering.
// (https://wiki.libsdl.org/SDL_RenderGetLogicalSize)
func (renderer *Renderer) GetLogicalSize() (w, h int32) {
	C.SDL_RenderGetLogicalSize(
		renderer.cptr(),
		(*C.int)(unsafe.Pointer(&w)),
		(*C.int)(unsafe.Pointer(&h)))
	return
}

// SetViewport sets the drawing area for rendering on the current target.
// (https://wiki.libsdl.org/SDL_RenderSetViewport)
func (renderer *Renderer) SetViewport(rect *Rect) error {
	return errorFromInt(int(
		C.SDL_RenderSetViewport(renderer.cptr(), rect.cptr())))
}

// GetViewport returns the drawing area for the current target.
// (https://wiki.libsdl.org/SDL_RenderGetViewport)
func (renderer *Renderer) GetViewport() (rect Rect) {
	C.SDL_RenderGetViewport(renderer.cptr(), rect.cptr())
	return
}

// IsClipEnabled returns whether clipping is enabled on the given renderer.
// (https://wiki.libsdl.org/SDL_RenderIsClipEnabled)
func (renderer *Renderer) IsClipEnabled() bool {
	return C.SDL_RenderIsClipEnabled(renderer.cptr()) == C.SDL_TRUE
}

// SetClipRect sets the clip rectangle for rendering on the specified target.
// (https://wiki.libsdl.org/SDL_RenderSetClipRect)
func (renderer *Renderer) SetClipRect(rect *Rect) error {
	return errorFromInt(int(
		C.SDL_RenderSetClipRect(renderer.cptr(), rect.cptr())))
}

// GetClipRect returns the clip rectangle for the current target.
// (https://wiki.libsdl.org/SDL_RenderGetClipRect)
func (renderer *Renderer) GetClipRect() (rect Rect) {
	C.SDL_RenderGetClipRect(renderer.cptr(), rect.cptr())
	return
}

// SetScale sets the drawing scale for rendering on the current target.
// (https://wiki.libsdl.org/SDL_RenderSetScale)
func (renderer *Renderer) SetScale(scaleX, scaleY float32) error {
	return errorFromInt(int(
		C.SDL_RenderSetScale(
			renderer.cptr(),
			C.float(scaleX),
			C.float(scaleY))))
}

// GetScale returns the drawing scale for the current target.
// (https://wiki.libsdl.org/SDL_RenderGetScale)
func (renderer *Renderer) GetScale() (scaleX, scaleY float32) {
	C.SDL_RenderGetScale(
		renderer.cptr(),
		(*C.float)(unsafe.Pointer(&scaleX)),
		(*C.float)(unsafe.Pointer(&scaleY)))
	return
}

// SetIntegerScale sets whether to force integer scales for
// resolution-independent rendering.
//
// This function restricts the logical viewport to integer values - that is,
// when a resolution is between two multiples of a logical size, the viewport
// size is rounded down to the lower multiple.
//
// (https://wiki.libsdl.org/SDL_RenderSetIntegerScale)
func (renderer *Renderer) SetIntegerScale(v bool) error {
	var cv C.SDL_bool = C.SDL_FALSE
	if v {
		cv = C.SDL_TRUE
	}

	return errorFromInt(int(C.SDL_RenderSetIntegerScale(renderer.cptr(), cv)))
}

// GetIntegerScale reports whether integer scales are forced for
// resolution-independent rendering.
//
// (https://wiki.libsdl.org/SDL_RenderGetIntegerScale)
func (renderer *Renderer) GetIntegerScale() (bool, error) {
	ClearError()
	if C.SDL_RenderGetIntegerScale(renderer.cptr()) == C.SDL_TRUE {
		return true, nil
	}
	return false, GetError()
}

// SetDrawColor sets the color used for drawing operations (Rect, Line and Clear).
// (https://wiki.libsdl.org/SDL_SetRenderDrawColor)
func (renderer *Renderer) SetDrawColor(r, g, b, a uint8) error {
	return errorFromInt(int(
		C.SDL_SetRenderDrawColor(
			renderer.cptr(),
			C.Uint8(r),
			C.Uint8(g),
			C.Uint8(b),
			C.Uint8(a))))
}

// SetDrawColorArray is a custom variant of SetDrawColor.
func (renderer *Renderer) SetDrawColorArray(bs ...uint8) error {
	_bs := []C.Uint8{0, 0, 0, 255}
	for i := 0; i < len(_bs) && i < len(bs); i++ {
		_bs[i] = C.Uint8(bs[i])
	}
	return errorFromInt(int(
		C.SDL_SetRenderDrawColor(
			renderer.cptr(),
			_bs[0],
			_bs[1],
			_bs[2],
			_bs[3])))
}

// GetDrawColor returns the color used for drawing operations (Rect, Line and Clear).
// (https://wiki.libsdl.org/SDL_GetRenderDrawColor)
func (renderer *Renderer) GetDrawColor() (r, g, b, a uint8, err error) {
	ret := C.SDL_GetRenderDrawColor(
		renderer.cptr(),
		(*C.Uint8)(unsafe.Pointer(&r)),
		(*C.Uint8)(unsafe.Pointer(&g)),
		(*C.Uint8)(unsafe.Pointer(&b)),
		(*C.Uint8)(unsafe.Pointer(&a)))
	return r, g, b, a, errorFromInt(int(ret))
}

// SetDrawBlendMode sets the blend mode used for drawing operations (Fill and Line).
// (https://wiki.libsdl.org/SDL_SetRenderDrawBlendMode)
func (renderer *Renderer) SetDrawBlendMode(bm BlendMode) error {
	return errorFromInt(int(
		C.SDL_SetRenderDrawBlendMode(renderer.cptr(), bm.c())))
}

// GetDrawBlendMode returns the blend mode used for drawing operations.
// (https://wiki.libsdl.org/SDL_GetRenderDrawBlendMode)
func (renderer *Renderer) GetDrawBlendMode(bm *BlendMode) error {
	return errorFromInt(int(
		C.SDL_GetRenderDrawBlendMode(renderer.cptr(), bm.cptr())))
}

// Clear clears the current rendering target with the drawing color.
// (https://wiki.libsdl.org/SDL_RenderClear)
func (renderer *Renderer) Clear() error {
	return errorFromInt(int(
		C.SDL_RenderClear(renderer.cptr())))
}

// DrawPoint draws a point on the current rendering target.
// (https://wiki.libsdl.org/SDL_RenderDrawPoint)
func (renderer *Renderer) DrawPoint(x, y int32) error {
	return errorFromInt(int(
		C.SDL_RenderDrawPoint(renderer.cptr(), C.int(x), C.int(y))))
}

// DrawPoints draws multiple points on the current rendering target.
// (https://wiki.libsdl.org/SDL_RenderDrawPoints)
func (renderer *Renderer) DrawPoints(points []Point) error {
	return errorFromInt(int(
		C.SDL_RenderDrawPoints(
			renderer.cptr(),
			points[0].cptr(),
			C.int(len(points)))))
}

// DrawLine draws a line on the current rendering target.
// (https://wiki.libsdl.org/SDL_RenderDrawLine)
func (renderer *Renderer) DrawLine(x1, y1, x2, y2 int32) error {
	return errorFromInt(int(
		C.SDL_RenderDrawLine(
			renderer.cptr(),
			C.int(x1),
			C.int(y1),
			C.int(x2),
			C.int(y2))))
}

// DrawLines draws a series of connected lines on the current rendering target.
// (https://wiki.libsdl.org/SDL_RenderDrawLines)
func (renderer *Renderer) DrawLines(points []Point) error {
	return errorFromInt(int(
		C.SDL_RenderDrawLines(
			renderer.cptr(),
			points[0].cptr(),
			C.int(len(points)))))
}

// DrawRect draws a rectangle on the current rendering target.
// (https://wiki.libsdl.org/SDL_RenderDrawRect)
func (renderer *Renderer) DrawRect(rect *Rect) error {
	return errorFromInt(int(
		C.SDL_RenderDrawRect(renderer.cptr(), rect.cptr())))
}

// DrawRects draws some number of rectangles on the current rendering target.
// (https://wiki.libsdl.org/SDL_RenderDrawRects)
func (renderer *Renderer) DrawRects(rects []Rect) error {
	return errorFromInt(int(
		C.SDL_RenderDrawRects(
			renderer.cptr(),
			rects[0].cptr(),
			C.int(len(rects)))))
}

// FillRect fills a rectangle on the current rendering target with the drawing color.
// (https://wiki.libsdl.org/SDL_RenderFillRect)
func (renderer *Renderer) FillRect(rect *Rect) error {
	return errorFromInt(int(
		C.SDL_RenderFillRect(renderer.cptr(), rect.cptr())))
}

// FillRects fills some number of rectangles on the current rendering target with the drawing color.
// (https://wiki.libsdl.org/SDL_RenderFillRects)
func (renderer *Renderer) FillRects(rects []Rect) error {
	return errorFromInt(int(
		C.SDL_RenderFillRects(
			renderer.cptr(),
			rects[0].cptr(),
			C.int(len(rects)))))
}

// Copy copies a portion of the texture to the current rendering target.
// (https://wiki.libsdl.org/SDL_RenderCopy)
func (renderer *Renderer) Copy(texture *Texture, src, dst *Rect) error {
	if dst == nil {
		return errorFromInt(int(
			C.SDL_RenderCopy(
				renderer.cptr(),
				texture.cptr(),
				src.cptr(),
				dst.cptr())))
	}
	return errorFromInt(int(
		C.RenderCopy(
			renderer.cptr(),
			texture.cptr(),
			src.cptr(),
			C.int(dst.X), C.int(dst.Y), C.int(dst.W), C.int(dst.H))))
}

// CopyEx copies a portion of the texture to the current rendering target, optionally rotating it by angle around the given center and also flipping it top-bottom and/or left-right.
// (https://wiki.libsdl.org/SDL_RenderCopyEx)
func (renderer *Renderer) CopyEx(texture *Texture, src, dst *Rect, angle float64, center *Point, flip RendererFlip) error {
	return errorFromInt(int(
		C.SDL_RenderCopyEx(
			renderer.cptr(),
			texture.cptr(),
			src.cptr(),
			dst.cptr(),
			C.double(angle),
			center.cptr(),
			flip.c())))
}

// DrawPointF draws a point on the current rendering target.
// TODO: (https://wiki.libsdl.org/SDL_RenderDrawPointF)
func (renderer *Renderer) DrawPointF(x, y float32) error {
	return errorFromInt(int(
		C.SDL_RenderDrawPointF(renderer.cptr(), C.float(x), C.float(y))))
}

// DrawPointsF draws multiple points on the current rendering target.
// TODO: (https://wiki.libsdl.org/SDL_RenderDrawPointsF)
func (renderer *Renderer) DrawPointsF(points []FPoint) error {
	return errorFromInt(int(
		C.SDL_RenderDrawPointsF(
			renderer.cptr(),
			points[0].cptr(),
			C.int(len(points)))))
}

// DrawLineF draws a line on the current rendering target.
// TODO: (https://wiki.libsdl.org/SDL_RenderDrawLineF)
func (renderer *Renderer) DrawLineF(x1, y1, x2, y2 float32) error {
	return errorFromInt(int(
		C.SDL_RenderDrawLineF(
			renderer.cptr(),
			C.float(x1),
			C.float(y1),
			C.float(x2),
			C.float(y2))))
}

// DrawLinesF draws a series of connected lines on the current rendering target.
// TODO: (https://wiki.libsdl.org/SDL_RenderDrawLinesF)
func (renderer *Renderer) DrawLinesF(points []FPoint) error {
	return errorFromInt(int(
		C.SDL_RenderDrawLinesF(
			renderer.cptr(),
			points[0].cptr(),
			C.int(len(points)))))
}

// DrawRectF draws a rectangle on the current rendering target.
// TODO: (https://wiki.libsdl.org/SDL_RenderDrawRectF)
func (renderer *Renderer) DrawRectF(rect *FRect) error {
	return errorFromInt(int(
		C.SDL_RenderDrawRectF(renderer.cptr(), rect.cptr())))
}

// DrawRectsF draws some number of rectangles on the current rendering target.
// TODO: (https://wiki.libsdl.org/SDL_RenderDrawRectsF)
func (renderer *Renderer) DrawRectsF(rects []FRect) error {
	return errorFromInt(int(
		C.SDL_RenderDrawRectsF(
			renderer.cptr(),
			rects[0].cptr(),
			C.int(len(rects)))))
}

// FillRectF fills a rectangle on the current rendering target with the drawing color.
// TODO: (https://wiki.libsdl.org/SDL_RenderFillRectF)
func (renderer *Renderer) FillRectF(rect *FRect) error {
	return errorFromInt(int(
		C.SDL_RenderFillRectF(renderer.cptr(), rect.cptr())))
}

// FillRectsF fills some number of rectangles on the current rendering target with the drawing color.
// TODO: (https://wiki.libsdl.org/SDL_RenderFillRectsF)
func (renderer *Renderer) FillRectsF(rects []FRect) error {
	return errorFromInt(int(
		C.SDL_RenderFillRectsF(
			renderer.cptr(),
			rects[0].cptr(),
			C.int(len(rects)))))
}

// CopyF copies a portion of the texture to the current rendering target.
// TODO: (https://wiki.libsdl.org/SDL_RenderCopyF)
func (renderer *Renderer) CopyF(texture *Texture, src *Rect, dst *FRect) error {
	return errorFromInt(int(
		C.SDL_RenderCopyF(
			renderer.cptr(),
			texture.cptr(),
			src.cptr(),
			dst.cptr())))
}

// CopyExF copies a portion of the texture to the current rendering target, optionally rotating it by angle around the given center and also flipping it top-bottom and/or left-right.
// TODO: (https://wiki.libsdl.org/SDL_RenderCopyExF)
func (renderer *Renderer) CopyExF(texture *Texture, src *Rect, dst *FRect, angle float64, center *FPoint, flip RendererFlip) error {
	return errorFromInt(int(
		C.SDL_RenderCopyExF(
			renderer.cptr(),
			texture.cptr(),
			src.cptr(),
			dst.cptr(),
			C.double(angle),
			center.cptr(),
			flip.c())))
}

// ReadPixels reads pixels from the current rendering target.
// (https://wiki.libsdl.org/SDL_RenderReadPixels)
func (renderer *Renderer) ReadPixels(rect *Rect, format uint32, pixels unsafe.Pointer, pitch int) error {
	return errorFromInt(int(
		C.SDL_RenderReadPixels(
			renderer.cptr(),
			rect.cptr(),
			C.Uint32(format),
			pixels,
			C.int(pitch))))
}

// Present updates the screen with any rendering performed since the previous call.
// (https://wiki.libsdl.org/SDL_RenderPresent)
func (renderer *Renderer) Present() {
	C.SDL_RenderPresent(renderer.cptr())
}

// Destroy destroys the specified texture.
// (https://wiki.libsdl.org/SDL_DestroyTexture)
func (texture *Texture) Destroy() error {
	lastErr := GetError()
	ClearError()
	C.SDL_DestroyTexture(texture.cptr())
	err := GetError()
	if err != nil {
		return err
	}
	SetError(lastErr)
	return nil
}

// Destroy destroys the rendering context for a window and free associated textures.
// (https://wiki.libsdl.org/SDL_DestroyRenderer)
func (renderer *Renderer) Destroy() error {
	lastErr := GetError()
	ClearError()
	C.SDL_DestroyRenderer(renderer.cptr())
	err := GetError()
	if err != nil {
		return err
	}
	SetError(lastErr)
	return nil
}

// Flush forces the rendering context to flush any pending commands to the underlying rendering API.
// TODO: (https://wiki.libsdl.org/SDL_RenderFlush)
func (renderer *Renderer) Flush() error {
	return errorFromInt(int(C.SDL_RenderFlush(renderer.cptr())))
}

// GLBind binds an OpenGL/ES/ES2 texture to the current context for use with OpenGL instructions when rendering OpenGL primitives directly.
// (https://wiki.libsdl.org/SDL_GL_BindTexture)
func (texture *Texture) GLBind(texw, texh *float32) error {
	return errorFromInt(int(
		C.SDL_GL_BindTexture(
			texture.cptr(),
			(*C.float)(unsafe.Pointer(texw)),
			(*C.float)(unsafe.Pointer(texh)))))
}

// GLUnbind unbinds an OpenGL/ES/ES2 texture from the current context.
// (https://wiki.libsdl.org/SDL_GL_UnbindTexture)
func (texture *Texture) GLUnbind() error {
	return errorFromInt(int(
		C.SDL_GL_UnbindTexture(texture.cptr())))
}

// GetMetalLayer gets the CAMetalLayer associated with the given Metal renderer
// (https://wiki.libsdl.org/SDL_RenderGetMetalLayer)
func (renderer *Renderer) GetMetalLayer() (layer unsafe.Pointer, err error) {
	layer = C.SDL_RenderGetMetalLayer(renderer.cptr())
	if layer == nil {
		err = GetError()
	}
	return
}

// GetMetalCommandEncoder gets the Metal command encoder for the current frame
// (https://wiki.libsdl.org/SDL_RenderGetMetalCommandEncoder)
func (renderer *Renderer) GetMetalCommandEncoder() (encoder unsafe.Pointer, err error) {
	encoder = C.SDL_RenderGetMetalCommandEncoder(renderer.cptr())
	if encoder == nil {
		err = GetError()
	}
	return
}

// UpdateNV updates a rectangle within a planar NV12 or NV21 texture with new pixels.
// (https://wiki.libsdl.org/SDL_UpdateNVTexture)
func (texture *Texture) UpdateNV(rect *Rect, yPlane []byte, yPitch int, uvPlane []byte, uvPitch int) error {
	var yPlanePtr, uvPlanePtr *byte
	if yPlane != nil {
		yPlanePtr = &yPlane[0]
	}
	if uvPlane != nil {
		uvPlanePtr = &uvPlane[0]
	}
	return errorFromInt(int(
		C.SDL_UpdateNVTexture(
			texture.cptr(),
			rect.cptr(),
			(*C.Uint8)(unsafe.Pointer(yPlanePtr)),
			C.int(yPitch),
			(*C.Uint8)(unsafe.Pointer(uvPlanePtr)),
			C.int(uvPitch))))
}

// RenderGeometry renders a list of triangles, optionally using a texture and
// indices into the vertex array Color and alpha modulation is done per vertex
// (SDL_SetTextureColorMod and SDL_SetTextureAlphaMod are ignored).
// (https://wiki.libsdl.org/SDL_RenderGeometry)
func (renderer *Renderer) RenderGeometry(texture *Texture, vertices []Vertex, indices []int32) (err error) {
	_texture := texture.cptr()
	_vertices := (*C.SDL_Vertex)(unsafe.Pointer(&vertices[0]))
	_num_vertices := C.int(len(vertices))
	var _indices *C.int
	_num_indices := C.int(0)
	if indices != nil {
		_indices = (*C.int)(unsafe.Pointer(&indices[0]))
		_num_indices = C.int(len(indices))
	}
	err = errorFromInt(int(C.SDL_RenderGeometry(renderer.cptr(), _texture, _vertices, _num_vertices, _indices, _num_indices)))
	return
}

// RenderGeomtryRaw renders a list of triangles, optionally using a texture and
// indices into the vertex arrays Color and alpha modulation is done per vertex
// (SDL_SetTextureColorMod and SDL_SetTextureAlphaMod are ignored).
// (https://wiki.libsdl.org/SDL_RenderGeometryRaw)
func (renderer *Renderer) RenderGeometryRaw(texture *Texture, xy *float32, xy_stride int, color *Color, color_stride int, uv *float32, uv_stride int, num_vertices int, indices unsafe.Pointer, num_indices int, size_indices int) (err error) {
	_texture := texture.cptr()
	_xy := (*C.float)(xy)
	_xy_stride := C.int(xy_stride)
	_color := (*C.SDL_Color)(unsafe.Pointer(color))
	_color_stride := C.int(color_stride)
	_uv := (*C.float)(uv)
	_uv_stride := C.int(uv_stride)
	_num_vertices := C.int(num_vertices)
	_num_indices := C.int(num_indices)
	_size_indices := C.int(size_indices)
	_indices := indices

	err = errorFromInt(int(C.RenderGeometryRaw(renderer.cptr(), _texture, _xy, _xy_stride, _color, _color_stride, _uv, _uv_stride, _num_vertices, _indices, _num_indices, _size_indices)))
	return
}

// SetTextureUserData associates a user-specified pointer with a texture.
// (https://wiki.libsdl.org/SDL_SetTextureUserData)
func (texture *Texture) SetTextureUserData(userdata unsafe.Pointer) (err error) {
	err = errorFromInt(int(C.SDL_SetTextureUserData(texture.cptr(), userdata)))
	return
}

// GetTextureUserData gets the user-specified pointer associated with a texture.
// (https://wiki.libsdl.org/SDL_GetTextureUserData)
func (texture *Texture) GetTextureUserData() (userdata unsafe.Pointer) {
	userdata = C.SDL_GetTextureUserData(texture.cptr())
	return
}

// RenderWindowToLogical gets logical coordinates of point in renderer when given real coordinates of
// point in window.
//
// Logical coordinates will differ from real coordinates when render is scaled
// and logical renderer size set
//
// (https://wiki.libsdl.org/SDL_RenderWindowToLogical)
func (renderer *Renderer) RenderWindowToLogical(windowX, windowY int) (logicalX, logicalY float32) {
	_windowX := C.int(windowX)
	_windowY := C.int(windowY)
	_logicalX := C.float(0)
	_logicalY := C.float(0)
	C.SDL_RenderWindowToLogical(renderer.cptr(), _windowX, _windowY, &_logicalX, &_logicalY)
	logicalX = float32(_logicalX)
	logicalY = float32(_logicalY)
	return
}

// RenderLogicalToWindow gets real coordinates of point in window when given logical coordinates of point in renderer.
// Logical coordinates will differ from real coordinates when render is scaled and logical renderer size set.
// (https://wiki.libsdl.org/SDL_RenderLogicalToWindow)
func (renderer *Renderer) RenderLogicalToWindow(logicalX, logicalY float32) (windowX, windowY int) {
	_logicalX := C.float(logicalX)
	_logicalY := C.float(logicalY)
	_windowX := C.int(0)
	_windowY := C.int(0)
	C.SDL_RenderLogicalToWindow(renderer.cptr(), _logicalX, _logicalY, &_windowX, &_windowY)
	windowX = int(_windowX)
	windowY = int(_windowY)
	return
}

// RenderSetVSync toggles VSync of the given renderer.
// (https://wiki.libsdl.org/SDL_RenderSetVSync)
func (renderer *Renderer) RenderSetVSync(vsync bool) (err error) {
	_vsync := C.int(Btoi(vsync))
	err = errorFromInt(int(C.SDL_RenderSetVSync(renderer.cptr(), _vsync)))
	return
}

// GetWindow gets the window associated with a renderer.
// (https://wiki.libsdl.org/SDL_RenderGetWindow)
func (renderer *Renderer) GetWindow() (window *Window, err error) {
	window = (*Window)(unsafe.Pointer(C.SDL_RenderGetWindow(renderer.cptr())))
	return
}
