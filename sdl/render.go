package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,1))
#pragma message("SDL_UpdateYUVTexture is not supported before SDL 2.0.1")
static inline int SDL_UpdateYUVTexture(SDL_Texture* texture, const SDL_Rect* rect, const Uint8* Yplane, int Ypitch, const Uint8* Uplane, int Upitch, const Uint8* Vplane, int Vpitch)
{
	return -1;
}
#endif
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
func (texture *Texture) Query() (uint32, int, int32, int32, error) {
	var format C.Uint32
	var access C.int
	var width C.int
	var height C.int
	ret := C.SDL_QueryTexture(texture.cptr(), &format, &access, &width, &height)
	if ret < 0 {
		return 0, 0, 0, 0, GetError()
	}
	return uint32(format), int(access), int32(width), int32(height), nil
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
func (texture *Texture) Update(rect *Rect, pixels []byte, pitch int) error {
	return errorFromInt(int(
		C.SDL_UpdateTexture(
			texture.cptr(),
			rect.cptr(),
			unsafe.Pointer(&pixels[0]),
			C.int(pitch))))
}

// UpdateYUV updates a rectangle within a planar YV12 or IYUV texture with new pixel data.
// (https://wiki.libsdl.org/SDL_UpdateYUVTexture)
func (texture *Texture) UpdateYUV(rect *Rect, yPlane []byte, yPitch int, uPlane []byte, uPitch int, vPlane []byte, vPitch int) error {
	return errorFromInt(int(
		C.SDL_UpdateYUVTexture(
			texture.cptr(),
			rect.cptr(),
			(*C.Uint8)(unsafe.Pointer(&yPlane[0])),
			C.int(yPitch),
			(*C.Uint8)(unsafe.Pointer(&uPlane[0])),
			C.int(uPitch),
			(*C.Uint8)(unsafe.Pointer(&vPlane[0])),
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
	return errorFromInt(int(
		C.SDL_RenderCopy(
			renderer.cptr(),
			texture.cptr(),
			src.cptr(),
			dst.cptr())))
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
