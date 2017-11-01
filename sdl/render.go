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
import "reflect"
import "unsafe"

const (
	RENDERER_SOFTWARE      = C.SDL_RENDERER_SOFTWARE
	RENDERER_ACCELERATED   = C.SDL_RENDERER_ACCELERATED
	RENDERER_PRESENTVSYNC  = C.SDL_RENDERER_PRESENTVSYNC
	RENDERER_TARGETTEXTURE = C.SDL_RENDERER_TARGETTEXTURE

	TEXTUREACCESS_STATIC    = C.SDL_TEXTUREACCESS_STATIC
	TEXTUREACCESS_STREAMING = C.SDL_TEXTUREACCESS_STREAMING
	TEXTUREACCESS_TARGET    = C.SDL_TEXTUREACCESS_TARGET

	TEXTUREMODULATE_NONE  = C.SDL_TEXTUREMODULATE_NONE
	TEXTUREMODULATE_COLOR = C.SDL_TEXTUREMODULATE_COLOR
	TEXTUREMODULATE_ALPHA = C.SDL_TEXTUREMODULATE_ALPHA
)

const (
	FLIP_NONE       RendererFlip = C.SDL_FLIP_NONE
	FLIP_HORIZONTAL              = C.SDL_FLIP_HORIZONTAL
	FLIP_VERTICAL                = C.SDL_FLIP_VERTICAL
)

// RendererInfo (https://wiki.libsdl.org/SDL_RendererInfo)
type RendererInfo struct {
	Name string
	RendererInfoData
}

type cRendererInfo struct {
	Name *C.char
	RendererInfoData
}

type RendererInfoData struct {
	Flags             uint32
	NumTextureFormats uint32
	TextureFormats    [16]int32
	MaxTextureWidth   int32
	MaxTextureHeight  int32
}

func (info *RendererInfo) cptr() *C.SDL_RendererInfo {
	return (*C.SDL_RendererInfo)(unsafe.Pointer(info))
}

func (info *cRendererInfo) cptr() *C.SDL_RendererInfo {
	return (*C.SDL_RendererInfo)(unsafe.Pointer(info))
}

// RendererFlip (https://wiki.libsdl.org/SDL_RendererFlip)
type RendererFlip uint32
type cRendererFlip C.SDL_RendererFlip

func (flip RendererFlip) c() C.SDL_RendererFlip {
	return C.SDL_RendererFlip(flip)
}

// GetNumRenderDrivers (https://wiki.libsdl.org/SDL_GetNumRenderDrivers)
func GetNumRenderDrivers() int {
	return int(C.SDL_GetNumRenderDrivers())
}

// GetRenderDriverInfo (https://wiki.libsdl.org/SDL_GetRenderDriverInfo)
func GetRenderDriverInfo(index int, info *RendererInfo) int {
	var cinfo cRendererInfo
	ret := int(C.SDL_GetRenderDriverInfo(C.int(index), cinfo.cptr()))

	info.RendererInfoData = cinfo.RendererInfoData
	// No need to free, it's done by DestroyRenderer
	info.Name = C.GoString(cinfo.Name)

	return ret
}

// CreateWindowAndRenderer (https://wiki.libsdl.org/SDL_CreateWindowAndRenderer)
func CreateWindowAndRenderer(w, h int, flags uint32) (*Window, *Renderer, error) {
	var window *C.SDL_Window
	var renderer *C.SDL_Renderer
	ret := C.SDL_CreateWindowAndRenderer(C.int(w), C.int(h), C.Uint32(flags), &window, &renderer)
	if ret == -1 {
		return nil, nil, GetError()
	}
	return (*Window)(unsafe.Pointer(window)), (*Renderer)(unsafe.Pointer(renderer)), nil
}

// CreateRenderer (https://wiki.libsdl.org/SDL_CreateRenderer)
func CreateRenderer(window *Window, index int, flags uint32) (*Renderer, error) {
	_renderer := C.SDL_CreateRenderer(window.cptr(), C.int(index), C.Uint32(flags))
	if _renderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(unsafe.Pointer(_renderer)), nil
}

// CreateSoftwareRenderer (https://wiki.libsdl.org/SDL_CreateSoftwareRenderer)
func CreateSoftwareRenderer(surface *Surface) (*Renderer, error) {
	_renderer := C.SDL_CreateSoftwareRenderer(surface.cptr())
	if _renderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(unsafe.Pointer(_renderer)), nil
}

// Window (https://wiki.libsdl.org/SDL_GetRenderer)
func (window *Window) GetRenderer() (*Renderer, error) {
	_renderer := C.SDL_GetRenderer(window.cptr())
	if _renderer == nil {
		return nil, GetError()
	}
	return (*Renderer)(unsafe.Pointer(_renderer)), nil
}

// Renderer (https://wiki.libsdl.org/SDL_GetRendererInfo)
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

// Renderer (https://wiki.libsdl.org/SDL_GetRendererOutputSize)
func (renderer *Renderer) GetOutputSize() (w, h int, err error) {
	_w := (*C.int)(unsafe.Pointer(&w))
	_h := (*C.int)(unsafe.Pointer(&h))
	_ret := C.SDL_GetRendererOutputSize(renderer.cptr(), _w, _h)
	if _ret < 0 {
		err = GetError()
	}
	return
}

// Renderer (https://wiki.libsdl.org/SDL_CreateTexture)
func (renderer *Renderer) CreateTexture(format uint32, access int, w int, h int) (*Texture, error) {
	_format := C.Uint32(format)
	_access := C.int(access)
	_w := C.int(w)
	_h := C.int(h)
	_texture := C.SDL_CreateTexture(renderer.cptr(), _format, _access, _w, _h)
	if _texture == nil {
		return nil, GetError()
	}
	return (*Texture)(unsafe.Pointer(_texture)), nil
}

// Renderer (https://wiki.libsdl.org/SDL_CreateTextureFromSurface)
func (renderer *Renderer) CreateTextureFromSurface(surface *Surface) (*Texture, error) {
	_texture := C.SDL_CreateTextureFromSurface(renderer.cptr(), surface.cptr())
	if _texture == nil {
		return nil, GetError()
	}
	return (*Texture)(unsafe.Pointer(_texture)), nil
}

// Texture (https://wiki.libsdl.org/SDL_QueryTexture)
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

// Texture (https://wiki.libsdl.org/SDL_SetTextureColorMod)
func (texture *Texture) SetColorMod(r uint8, g uint8, b uint8) error {
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_ret := C.SDL_SetTextureColorMod(texture.cptr(), _r, _g, _b)
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Texture (https://wiki.libsdl.org/SDL_SetTextureAlphaMod)
func (texture *Texture) SetAlphaMod(alpha uint8) error {
	_ret := C.SDL_SetTextureAlphaMod(texture.cptr(), C.Uint8(alpha))
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Texture (https://wiki.libsdl.org/SDL_GetTextureAlphaMod)
func (texture *Texture) GetAlphaMod() (alpha uint8, err error) {
	_alpha := (*C.Uint8)(unsafe.Pointer(&alpha))
	_ret := C.SDL_GetTextureAlphaMod(texture.cptr(), _alpha)
	if _ret < 0 {
		err = GetError()
		return alpha, err
	}
	return alpha, nil
}

// Texture (https://wiki.libsdl.org/SDL_SetTextureBlendMode)
func (texture *Texture) SetBlendMode(bm BlendMode) error {
	_ret := C.SDL_SetTextureBlendMode(texture.cptr(), bm.c())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Texture (https://wiki.libsdl.org/SDL_GetTextureBlendMode)
func (texture *Texture) GetBlendMode() (bm BlendMode, err error) {
	_ret := C.SDL_GetTextureBlendMode(texture.cptr(), bm.cptr())
	if _ret < 0 {
		err = GetError()
		return bm, err
	}
	return bm, nil
}

// Texture (https://wiki.libsdl.org/SDL_UpdateTexture)
func (texture *Texture) Update(rect *Rect, pixels []byte, pitch int) error {
	_pixels := unsafe.Pointer(&pixels[0])
	_pitch := C.int(pitch)
	_ret := C.SDL_UpdateTexture(texture.cptr(), rect.cptr(), _pixels, _pitch)
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Texture (https://wiki.libsdl.org/SDL_UpdateYUVTexture)
func (texture *Texture) UpdateYUV(rect *Rect, yPlane []byte, yPitch int, uPlane []byte, uPitch int, vPlane []byte, vPitch int) error {
	_yPlane := (*C.Uint8)(unsafe.Pointer(&yPlane[0]))
	_yPitch := C.int(yPitch)
	_uPlane := (*C.Uint8)(unsafe.Pointer(&uPlane[0]))
	_uPitch := C.int(uPitch)
	_vPlane := (*C.Uint8)(unsafe.Pointer(&vPlane[0]))
	_vPitch := C.int(vPitch)
	_ret := C.SDL_UpdateYUVTexture(texture.cptr(), rect.cptr(), _yPlane, _yPitch, _uPlane, _uPitch, _vPlane, _vPitch)
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Texture (https://wiki.libsdl.org/SDL_LockTexture)
func (texture *Texture) Lock(rect *Rect) ([]byte, int, error) {
	var _pitch C.int
	var _pixels unsafe.Pointer
	var b []byte
	var length int

	_ret := C.SDL_LockTexture(texture.cptr(), rect.cptr(), &_pixels, &_pitch)
	if _ret < 0 {
		return b, int(_pitch), GetError()
	}

	_, _, w, h, err := texture.Query()
	if err != nil {
		return b, int(_pitch), GetError()
	}

	pitch := int32(_pitch)
	if rect != nil {
		length = int((pitch / w) * rect.W * rect.H)
	} else {
		length = int(pitch * h)
	}
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sliceHeader.Cap = int(length)
	sliceHeader.Len = int(length)
	sliceHeader.Data = uintptr(_pixels)

	return b, int(pitch), nil
}

// Texture (https://wiki.libsdl.org/SDL_UnlockTexture)
func (texture *Texture) Unlock() {
	C.SDL_UnlockTexture(texture.cptr())
}

// Renderer (https://wiki.libsdl.org/SDL_RenderTargetSupported)
func (renderer *Renderer) RenderTargetSupported() bool {
	return C.SDL_RenderTargetSupported(renderer.cptr()) != 0
}

// Renderer (https://wiki.libsdl.org/SDL_SetRenderTarget)
func (renderer *Renderer) SetRenderTarget(texture *Texture) error {
	_ret := C.SDL_SetRenderTarget(renderer.cptr(), texture.cptr())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_GetRenderTarget)
func (renderer *Renderer) GetRenderTarget() *Texture {
	return (*Texture)(unsafe.Pointer(C.SDL_GetRenderTarget(renderer.cptr())))
}

// Renderer (https://wiki.libsdl.org/SDL_RenderSetLogicalSize)
func (renderer *Renderer) SetLogicalSize(w int, h int) error {
	_ret := C.SDL_RenderSetLogicalSize(renderer.cptr(), C.int(w), C.int(h))
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderGetLogicalSize)
func (renderer *Renderer) GetLogicalSize() (w, h int) {
	_w := (*C.int)(unsafe.Pointer(&w))
	_h := (*C.int)(unsafe.Pointer(&h))
	C.SDL_RenderGetLogicalSize(renderer.cptr(), _w, _h)
	return
}

// Renderer (https://wiki.libsdl.org/SDL_RenderSetViewport)
func (renderer *Renderer) SetViewport(rect *Rect) error {
	_ret := C.SDL_RenderSetViewport(renderer.cptr(), rect.cptr())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderGetViewport)
func (renderer *Renderer) GetViewport(rect *Rect) {
	C.SDL_RenderGetViewport(renderer.cptr(), rect.cptr())
}

// Renderer (https://wiki.libsdl.org/SDL_RenderSetClipRect)
func (renderer *Renderer) SetClipRect(rect *Rect) error {
	_ret := C.SDL_RenderSetClipRect(renderer.cptr(), rect.cptr())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderGetClipRect)
func (renderer *Renderer) GetClipRect(rect *Rect) {
	C.SDL_RenderGetClipRect(renderer.cptr(), rect.cptr())
}

// Renderer (https://wiki.libsdl.org/SDL_RenderSetScale)
func (renderer *Renderer) SetScale(scaleX, scaleY float32) error {
	_scaleX := C.float(scaleX)
	_scaleY := C.float(scaleY)
	_ret := C.SDL_RenderSetScale(renderer.cptr(), _scaleX, _scaleY)
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderGetScale)
func (renderer *Renderer) GetScale() (scaleX, scaleY float32) {
	_scaleX := (*C.float)(unsafe.Pointer(&scaleX))
	_scaleY := (*C.float)(unsafe.Pointer(&scaleY))
	C.SDL_RenderGetScale(renderer.cptr(), _scaleX, _scaleY)
	return
}

// Renderer (https://wiki.libsdl.org/SDL_SetRenderDrawColor)
func (renderer *Renderer) SetDrawColor(r, g, b, a uint8) error {
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	_ret := C.SDL_SetRenderDrawColor(renderer.cptr(), _r, _g, _b, _a)
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Custom variant of SetDrawColor
func (renderer *Renderer) SetDrawColorArray(bs ...uint8) error {
	_bs := []C.Uint8{0, 0, 0, 255}
	for i := 0; i < len(_bs) && i < len(bs); i++ {
		_bs[i] = C.Uint8(bs[i])
	}
	_ret := C.SDL_SetRenderDrawColor(renderer.cptr(), _bs[0], _bs[1], _bs[2], _bs[3])
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_GetRenderDrawColor)
func (renderer *Renderer) GetDrawColor() (r, g, b, a uint8, err error) {
	_r := (*C.Uint8)(unsafe.Pointer(&r))
	_g := (*C.Uint8)(unsafe.Pointer(&g))
	_b := (*C.Uint8)(unsafe.Pointer(&b))
	_a := (*C.Uint8)(unsafe.Pointer(&a))
	_ret := C.SDL_GetRenderDrawColor(renderer.cptr(), _r, _g, _b, _a)
	if _ret < 0 {
		err = GetError()
		return r, g, b, a, err
	}
	return r, g, b, a, nil
}

// Renderer (https://wiki.libsdl.org/SDL_SetRenderDrawBlendMode)
func (renderer *Renderer) SetDrawBlendMode(bm BlendMode) error {
	_ret := C.SDL_SetRenderDrawBlendMode(renderer.cptr(), bm.c())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_GetRenderDrawBlendMode)
func (renderer *Renderer) GetDrawBlendMode(bm *BlendMode) error {
	_ret := C.SDL_GetRenderDrawBlendMode(renderer.cptr(), bm.cptr())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderClear)
func (renderer *Renderer) Clear() error {
	_ret := C.SDL_RenderClear(renderer.cptr())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderDrawPoint)
func (renderer *Renderer) DrawPoint(x, y int) error {
	_ret := C.SDL_RenderDrawPoint(renderer.cptr(), C.int(x), C.int(y))
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderDrawPoints)
func (renderer *Renderer) DrawPoints(points []Point) error {
	_ret := C.SDL_RenderDrawPoints(renderer.cptr(), points[0].cptr(), C.int(len(points)))
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderDrawLine)
func (renderer *Renderer) DrawLine(x1, y1, x2, y2 int) error {
	_x1 := C.int(x1)
	_y1 := C.int(y1)
	_x2 := C.int(x2)
	_y2 := C.int(y2)
	_ret := C.SDL_RenderDrawLine(renderer.cptr(), _x1, _y1, _x2, _y2)
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderDrawLines)
func (renderer *Renderer) DrawLines(points []Point) error {
	_ret := C.SDL_RenderDrawLines(renderer.cptr(), points[0].cptr(), C.int(len(points)))
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderDrawRect)
func (renderer *Renderer) DrawRect(rect *Rect) error {
	_ret := C.SDL_RenderDrawRect(renderer.cptr(), rect.cptr())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderDrawRects)
func (renderer *Renderer) DrawRects(rects []Rect) error {
	_ret := C.SDL_RenderDrawRects(renderer.cptr(), rects[0].cptr(), C.int(len(rects)))
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderFillRect)
func (renderer *Renderer) FillRect(rect *Rect) error {
	_ret := C.SDL_RenderFillRect(renderer.cptr(), rect.cptr())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderFillRects)
func (renderer *Renderer) FillRects(rects []Rect) error {
	_ret := C.SDL_RenderFillRects(renderer.cptr(), rects[0].cptr(), C.int(len(rects)))
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderCopy)
func (renderer *Renderer) Copy(texture *Texture, src, dst *Rect) error {
	_ret := C.SDL_RenderCopy(renderer.cptr(), texture.cptr(), src.cptr(), dst.cptr())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderCopyEx)
func (renderer *Renderer) CopyEx(texture *Texture, src, dst *Rect, angle float64, center *Point, flip RendererFlip) error {
	_ret := C.SDL_RenderCopyEx(renderer.cptr(), texture.cptr(), src.cptr(), dst.cptr(), C.double(angle), center.cptr(), flip.c())
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderReadPixels)
func (renderer *Renderer) ReadPixels(rect *Rect, format uint32, pixels unsafe.Pointer, pitch int) error {
	_ret := C.SDL_RenderReadPixels(renderer.cptr(), rect.cptr(), C.Uint32(format), pixels, C.int(pitch))
	if _ret < 0 {
		return GetError()
	}
	return nil
}

// Renderer (https://wiki.libsdl.org/SDL_RenderPresent)
func (renderer *Renderer) Present() {
	C.SDL_RenderPresent(renderer.cptr())
}

// Texture (https://wiki.libsdl.org/SDL_DestroyTexture)
func (texture *Texture) Destroy() {
	C.SDL_DestroyTexture(texture.cptr())
}

// Renderer (https://wiki.libsdl.org/SDL_DestroyRenderer)
func (renderer *Renderer) Destroy() {
	C.SDL_DestroyRenderer(renderer.cptr())
}

func (texture *Texture) GL_BindTexture(texw, texh *float32) error {
	_texw := (*C.float)(unsafe.Pointer(texw))
	_texh := (*C.float)(unsafe.Pointer(texh))
	_ret := C.SDL_GL_BindTexture(texture.cptr(), _texw, _texh)
	if _ret < 0 {
		return GetError()
	}
	return nil
}

func (texture *Texture) GL_UnbindTexture() error {
	_ret := C.SDL_GL_UnbindTexture(texture.cptr())
	if _ret < 0 {
		return GetError()
	}
	return nil
}
