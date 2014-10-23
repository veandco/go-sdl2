package sdl

// #include <SDL2/SDL.h>
import "C"
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

	FLIP_NONE       = C.SDL_FLIP_NONE
	FLIP_HORIZONTAL = C.SDL_FLIP_HORIZONTAL
	FLIP_VERTICAL   = C.SDL_FLIP_VERTICAL
)

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
	MaxTextureWidth   int
	MaxTextureHeight  int
}

type RendererFlip uint

func (info *RendererInfo) cptr() *C.SDL_RendererInfo {
	return (*C.SDL_RendererInfo)(unsafe.Pointer(info))
}

func (info *cRendererInfo) cptr() *C.SDL_RendererInfo {
	return (*C.SDL_RendererInfo)(unsafe.Pointer(info))
}

func (flip RendererFlip) c() C.SDL_RendererFlip {
	return C.SDL_RendererFlip(flip)
}

func GetNumRenderDrivers() int {
	return int(C.SDL_GetNumRenderDrivers())
}

func GetRenderDriverInfo(index int, info *RendererInfo) int {
	var cinfo cRendererInfo
	ret := int(C.SDL_GetRenderDriverInfo(C.int(index), cinfo.cptr()))

	info.RendererInfoData = cinfo.RendererInfoData
	// No need to free, it's done by DestroyRenderer
	info.Name = C.GoString(cinfo.Name)

	return ret
}

func CreateWindowAndRenderer(w, h int, flags uint32) (*Window, *Renderer) {
	var window *C.SDL_Window
	var renderer *C.SDL_Renderer
	C.SDL_CreateWindowAndRenderer(C.int(w), C.int(h), C.Uint32(flags), &window, &renderer)
	return (*Window)(unsafe.Pointer(window)), (*Renderer)(unsafe.Pointer(renderer))
}

func CreateRenderer(window *Window, index int, flags uint32) *Renderer {
	return (*Renderer)(unsafe.Pointer(C.SDL_CreateRenderer(window.cptr(), C.int(index), C.Uint32(flags))))
}

func CreateSoftwareRenderer(surface *Surface) *Renderer {
	return (*Renderer)(unsafe.Pointer(C.SDL_CreateSoftwareRenderer(surface.cptr())))
}

func (window *Window) GetRenderer() *Renderer {
	return (*Renderer)(unsafe.Pointer(C.SDL_GetRenderer(window.cptr())))
}

func (renderer *Renderer) GetRendererInfo(info *RendererInfo) int {
	var cinfo cRendererInfo
	ret := int(C.SDL_GetRendererInfo(renderer.cptr(), cinfo.cptr()))

	info.RendererInfoData = cinfo.RendererInfoData
	// No need to free, it's done by DestroyRenderer
	info.Name = C.GoString(cinfo.Name)

	return ret
}

func (renderer *Renderer) GetRendererOutputSize() (w, h int, status int) {
	_w := (*C.int)(unsafe.Pointer(&w))
	_h := (*C.int)(unsafe.Pointer(&h))
	status = int(C.SDL_GetRendererOutputSize(renderer.cptr(), _w, _h))
	return
}

func (renderer *Renderer) CreateTexture(format uint32, access int, w int, h int) *Texture {
	_format := C.Uint32(format)
	_access := C.int(access)
	_w := C.int(w)
	_h := C.int(h)
	return (*Texture)(unsafe.Pointer(C.SDL_CreateTexture(renderer.cptr(), _format, _access, _w, _h)))
}

func (renderer *Renderer) CreateTextureFromSurface(surface *Surface) *Texture {
	return (*Texture)(unsafe.Pointer(C.SDL_CreateTextureFromSurface(renderer.cptr(), surface.cptr())))
}

func (texture *Texture) Query(format *uint32, access *int, w *int, h *int) int {
	_format := (*C.Uint32)(unsafe.Pointer(access))
	_access := (*C.int)(unsafe.Pointer(access))
	_w := (*C.int)(unsafe.Pointer(w))
	_h := (*C.int)(unsafe.Pointer(h))
	return int(C.SDL_QueryTexture(texture.cptr(), _format, _access, _w, _h))
}

func (texture *Texture) SetColorMod(r uint8, g uint8, b uint8) int {
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	return int(C.SDL_SetTextureColorMod(texture.cptr(), _r, _g, _b))
}

func (texture *Texture) SetAlphaMod(alpha uint8) int {
	return int(C.SDL_SetTextureAlphaMod(texture.cptr(), C.Uint8(alpha)))
}

func (texture *Texture) GetAlphaMod() (alpha uint8, status int) {
	_alpha := (*C.Uint8)(unsafe.Pointer(&alpha))
	status = int(C.SDL_GetTextureAlphaMod(texture.cptr(), _alpha))
	return alpha, status
}


func (texture *Texture) SetBlendMode(bm BlendMode) int {
	return int(C.SDL_SetTextureBlendMode(texture.cptr(), bm.c()))
}

func (texture *Texture) GetBlendMode() (bm BlendMode, status int) {
	status = int(C.SDL_GetTextureBlendMode(texture.cptr(), bm.cptr()))
	return
}

func (texture *Texture) Update(rect *Rect, pixels unsafe.Pointer, pitch int) int {
	_pitch := C.int(pitch)
	return int(C.SDL_UpdateTexture(texture.cptr(), rect.cptr(), pixels, _pitch))
}

func (texture *Texture) Lock(rect *Rect, pixels unsafe.Pointer, pitch *int) int {
	_pitch := (*C.int)(unsafe.Pointer(pitch))
	return int(C.SDL_LockTexture(texture.cptr(), rect.cptr(), &pixels, _pitch))
}

func (texture *Texture) Unlock() {
	C.SDL_UnlockTexture(texture.cptr())
}

func (renderer *Renderer) RenderTargetSupported() bool {
	return C.SDL_RenderTargetSupported(renderer.cptr()) != 0
}

func (renderer *Renderer) SetRenderTarget(texture *Texture) int {
	return int(C.SDL_SetRenderTarget(renderer.cptr(), texture.cptr()))
}

func (renderer *Renderer) GetRenderTarget() *Texture {
	return (*Texture)(unsafe.Pointer(C.SDL_GetRenderTarget(renderer.cptr())))
}

func (renderer *Renderer) SetLogicalSize(w int, h int) int {
	return int(C.SDL_RenderSetLogicalSize(renderer.cptr(), C.int(w), C.int(h)))
}

func (renderer *Renderer) SetViewport(rect *Rect) int {
	return int(C.SDL_RenderSetViewport(renderer.cptr(), rect.cptr()))
}

func (renderer *Renderer) GetViewport(rect *Rect) {
	C.SDL_RenderGetViewport(renderer.cptr(), rect.cptr())
}

func (renderer *Renderer) SetClipRect(rect *Rect) int {
	return int(C.SDL_RenderSetClipRect(renderer.cptr(), rect.cptr()))
}

func (renderer *Renderer) GetClipRect(rect *Rect) {
	C.SDL_RenderGetClipRect(renderer.cptr(), rect.cptr())
}

func (renderer *Renderer) SetScale(scaleX, scaleY float32) int {
	_scaleX := C.float(scaleX)
	_scaleY := C.float(scaleY)
	return int(C.SDL_RenderSetScale(renderer.cptr(), _scaleX, _scaleY))
}

func (renderer *Renderer) GetScale() (scaleX, scaleY float32) {
	_scaleX := (*C.float)(unsafe.Pointer(&scaleX))
	_scaleY := (*C.float)(unsafe.Pointer(&scaleY))
	C.SDL_RenderGetScale(renderer.cptr(), _scaleX, _scaleY)
	return
}

func (renderer *Renderer) SetDrawColor(r, g, b, a uint8) int {
	_r := C.Uint8(r)
	_g := C.Uint8(g)
	_b := C.Uint8(b)
	_a := C.Uint8(a)
	return int(C.SDL_SetRenderDrawColor(renderer.cptr(), _r, _g, _b, _a))
}

func (renderer *Renderer) GetDrawColor() (r, g, b, a uint8, status int) {
	_r := (*C.Uint8)(unsafe.Pointer(&r))
	_g := (*C.Uint8)(unsafe.Pointer(&g))
	_b := (*C.Uint8)(unsafe.Pointer(&b))
	_a := (*C.Uint8)(unsafe.Pointer(&a))
	status = int(C.SDL_GetRenderDrawColor(renderer.cptr(), _r, _g, _b, _a))
	return
}

func (renderer *Renderer) SetDrawBlendMode(bm BlendMode) int {
	return int(C.SDL_SetRenderDrawBlendMode(renderer.cptr(), bm.c()))
}

func (renderer *Renderer) GetDrawBlendMode(bm *BlendMode) int {
	return int(C.SDL_GetRenderDrawBlendMode(renderer.cptr(), bm.cptr()))
}

func (renderer *Renderer) Clear() int {
	return int(C.SDL_RenderClear(renderer.cptr()))
}

func (renderer *Renderer) DrawPoint(x, y int) int {
	return int(C.SDL_RenderDrawPoint(renderer.cptr(), C.int(x), C.int(y)))
}

func (renderer *Renderer) DrawPoints(points []Point) int {
	return int(C.SDL_RenderDrawPoints(renderer.cptr(), points[0].cptr(), C.int(len(points))))
}

func (renderer *Renderer) DrawLine(x1, y1, x2, y2 int) int {
	_x1 := C.int(x1)
	_y1 := C.int(y1)
	_x2 := C.int(x2)
	_y2 := C.int(y2)
	return int(C.SDL_RenderDrawLine(renderer.cptr(), _x1, _y1, _x2, _y2))
}

func (renderer *Renderer) DrawLines(points []Point) int {
	return int(C.SDL_RenderDrawLines(renderer.cptr(), points[0].cptr(), C.int(len(points))))
}

func (renderer *Renderer) DrawRect(rect *Rect) int {
	return int(C.SDL_RenderDrawRect(renderer.cptr(), rect.cptr()))
}

func (renderer *Renderer) DrawRects(rects []Rect) int {
	return int(C.SDL_RenderDrawRects(renderer.cptr(), rects[0].cptr(), C.int(len(rects))))
}

func (renderer *Renderer) FillRect(rect *Rect) int {
	return int(C.SDL_RenderFillRect(renderer.cptr(), rect.cptr()))
}

func (renderer *Renderer) FillRects(rects []Rect) int {
	return int(C.SDL_RenderFillRects(renderer.cptr(), rects[0].cptr(), C.int(len(rects))))
}

func (renderer *Renderer) Copy(texture *Texture, src, dst *Rect) int {
	return int(C.SDL_RenderCopy(renderer.cptr(), texture.cptr(), src.cptr(), dst.cptr()))
}

func (renderer *Renderer) CopyEx(texture *Texture, src, dst *Rect, angle float64, center *Point, flip RendererFlip) int {
	return int(C.SDL_RenderCopyEx(renderer.cptr(), texture.cptr(), src.cptr(), dst.cptr(), C.double(angle), center.cptr(), flip.c()))
}

func (renderer *Renderer) ReadPixels(rect *Rect, format uint32, pixels unsafe.Pointer, pitch int) int {
	return int(C.SDL_RenderReadPixels(renderer.cptr(), rect.cptr(), C.Uint32(format), pixels, C.int(pitch)))
}

func (renderer *Renderer) Present() {
	C.SDL_RenderPresent(renderer.cptr())
}

func (texture *Texture) Destroy() {
	C.SDL_DestroyTexture(texture.cptr())
}

func (renderer *Renderer) Destroy() {
	C.SDL_DestroyRenderer(renderer.cptr())
}

func (texture *Texture) GL_BindTexture(texw, texh *float32) int {
	_texw := (*C.float)(unsafe.Pointer(texw))
	_texh := (*C.float)(unsafe.Pointer(texh))
	return int(C.SDL_GL_BindTexture(texture.cptr(), _texw, _texh))
}

func (texture *Texture) GL_UnbindTexture() int {
	return int(C.SDL_GL_UnbindTexture(texture.cptr()))
}
