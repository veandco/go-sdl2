package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"

const (
	RENDERER_SOFTWARE      = 0x00000001
	RENDERER_ACCELERATED   = 0x00000002
	RENDERER_PRESENTVSYNC  = 0x00000004
	RENDERER_TARGETTEXTURE = 0x00000008

	TEXTUREACCESS_STATIC    = 0x00000001
	TEXTUREACCESS_STREAMING = 0x00000002
	TEXTUREACCESS_TARGET    = 0x00000003

	TEXTUREMODULATE_NONE  = 0x00000000
	TEXTUREMODULATE_COLOR = 0x00000001
	TEXTUREMODULATE_ALPHA = 0x00000002

	SDL_FLIP_NONE       = 0x00000000
	SDL_FLIP_HORIZONTAL = 0x00000001
	SDL_FLIP_VERTICAL   = 0x00000002
)

type RendererInfo struct {
	Name              string
	Flags             uint32
	NumTextureFormats uint32
	TextureFormats    [16]int32
	MaxTextureWidth   int
	MaxTextureHeight  int
}

func GetNumRenderDrivers() int {
	return (int)(C.SDL_GetNumRenderDrivers())
}

func GetRenderDriverInfo(index int, info *RendererInfo) int {
	_index := (C.int)(index)
	_info := (*C.SDL_RendererInfo)(unsafe.Pointer(info))
	return (int)(C.SDL_GetRenderDriverInfo(_index, _info))
}

func CreateWindowAndRenderer(w, h int, flags uint32) (*Window, *Renderer) {
	var window *C.SDL_Window
	var renderer *C.SDL_Renderer
	C.SDL_CreateWindowAndRenderer(C.int(w), C.int(h), C.Uint32(flags), &window, &renderer)
	return (*Window)(unsafe.Pointer(window)), (*Renderer)(unsafe.Pointer(renderer))
}

func CreateRenderer(window *Window, index int, flags uint32) *Renderer {
	_window := (*C.SDL_Window)(unsafe.Pointer(window))
	_index := (C.int)(index)
	_flags := (C.Uint32)(flags)
	return (*Renderer)(unsafe.Pointer(C.SDL_CreateRenderer(_window, _index, _flags)))
}

func CreateSoftwareRenderer(surface *Surface) *Renderer {
	_surface := (*C.SDL_Surface)(unsafe.Pointer(surface))
	return (*Renderer)(unsafe.Pointer(C.SDL_CreateSoftwareRenderer(_surface)))
}

func (window *Window) GetRenderer() *Renderer {
	_window := (*C.SDL_Window)(unsafe.Pointer(window))
	return (*Renderer)(unsafe.Pointer(C.SDL_GetRenderer(_window)))
}

func (renderer *Renderer) GetRendererInfo(info *RendererInfo) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_info := (*C.SDL_RendererInfo)(unsafe.Pointer(info))
	return (int)(C.SDL_GetRendererInfo(_renderer, _info))
}

func (renderer *Renderer) GetRendererOutputSize() (w, h int, status int) {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_w := (*C.int)(unsafe.Pointer(&w))
	_h := (*C.int)(unsafe.Pointer(&h))
	status = (int)(C.SDL_GetRendererOutputSize(_renderer, _w, _h))
	return
}

func CreateTexture(renderer *Renderer, format uint32, access int, w int, h int) *Texture {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_format := (C.Uint32)(format)
	_access := (C.int)(access)
	_w := (C.int)(w)
	_h := (C.int)(h)
	return (*Texture)(unsafe.Pointer(C.SDL_CreateTexture(_renderer, _format, _access, _w, _h)))
}

func (renderer *Renderer) CreateTextureFromSurface(surface *Surface) *Texture {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_surface := (*C.SDL_Surface)(unsafe.Pointer(surface))
	return (*Texture)(unsafe.Pointer(C.SDL_CreateTextureFromSurface(_renderer, _surface)))
}

func QueryTexture(texture *Texture, format *uint32, access *int, w *int, h *int) int {
	_texture := (*C.SDL_Texture)(unsafe.Pointer(texture))
	_format := (*C.Uint32)(unsafe.Pointer(access))
	_access := (*C.int)(unsafe.Pointer(access))
	_w := (*C.int)(unsafe.Pointer(w))
	_h := (*C.int)(unsafe.Pointer(h))
	return (int)(C.SDL_QueryTexture(_texture, _format, _access, _w, _h))
}

func (texture *Texture) SetColorMod(r uint8, g uint8, b uint8) int {
	_texture := (*C.SDL_Texture)(unsafe.Pointer(texture))
	_r := (C.Uint8)(r)
	_g := (C.Uint8)(g)
	_b := (C.Uint8)(b)
	return (int)(C.SDL_SetTextureColorMod(_texture, _r, _g, _b))
}

func (texture *Texture) SetAlphaMod(alpha uint8) int {
	_texture := (*C.SDL_Texture)(unsafe.Pointer(texture))
	_alpha := (C.Uint8)(alpha)
	return (int)(C.SDL_SetTextureAlphaMod(_texture, _alpha))
}

func (texture *Texture) SetBlendMode(blendMode uint32) int {
	_texture := (*C.SDL_Texture)(unsafe.Pointer(texture))
	_blendMode := (C.SDL_BlendMode)(C.Uint32(blendMode))
	return (int)(C.SDL_SetTextureBlendMode(_texture, _blendMode))
}

func (texture *Texture) GetBlendMode() (blendMode uint32, status int) {
	_texture := (*C.SDL_Texture)(unsafe.Pointer(texture))
	_blendMode := (*C.SDL_BlendMode)(unsafe.Pointer(&blendMode))
	status = (int)(C.SDL_GetTextureBlendMode(_texture, _blendMode))
	return
}

func (texture *Texture) Update(rect *Rect, pixels unsafe.Pointer, pitch int) int {
	_texture := (*C.SDL_Texture)(unsafe.Pointer(texture))
	_rect := (*C.SDL_Rect)(unsafe.Pointer(rect))
	_pitch := (C.int)(pitch)
	return (int)(C.SDL_UpdateTexture(_texture, _rect, pixels, _pitch))
}

func (texture *Texture) Lock(rect *Rect, pixels unsafe.Pointer, pitch *int) int {
	_texture := (*C.SDL_Texture)(unsafe.Pointer(texture))
	_rect := (*C.SDL_Rect)(unsafe.Pointer(rect))
	_pitch := (*C.int)(unsafe.Pointer(pitch))
	return (int)(C.SDL_LockTexture(_texture, _rect, &pixels, _pitch))
}

func (texture *Texture) Unlock() {
	_texture := (*C.SDL_Texture)(unsafe.Pointer(texture))
	(C.SDL_UnlockTexture(_texture))
}

func (renderer *Renderer) RenderTargetSupported() bool {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	return C.SDL_RenderTargetSupported(_renderer) != 0
}

func (renderer *Renderer) SetRenderTarget(texture *Texture) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_texture := (*C.SDL_Texture)(unsafe.Pointer(texture))
	return (int)(C.SDL_SetRenderTarget(_renderer, _texture))
}

func (renderer *Renderer) GetRenderTarget() *Texture {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	return (*Texture)(unsafe.Pointer(C.SDL_GetRenderTarget(_renderer)))
}

func (renderer *Renderer) SetLogicalSize(w int, h int) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_w := (C.int)(w)
	_h := (C.int)(h)
	return (int)(C.SDL_RenderSetLogicalSize(_renderer, _w, _h))
}

func (renderer *Renderer) SetViewport(rect *Rect) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_rect := (*C.SDL_Rect)(unsafe.Pointer(rect))
	return (int)(C.SDL_RenderSetViewport(_renderer, _rect))
}

func (renderer *Renderer) GetViewport(rect *Rect) {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_rect := (*C.SDL_Rect)(unsafe.Pointer(rect))
	C.SDL_RenderGetViewport(_renderer, _rect)
}

func (renderer *Renderer) SetClipRect(rect *Rect) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_rect := (*C.SDL_Rect)(unsafe.Pointer(rect))
	return (int)(C.SDL_RenderSetClipRect(_renderer, _rect))
}

func (renderer *Renderer) GetClipRect(rect *Rect) {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_rect := (*C.SDL_Rect)(unsafe.Pointer(rect))
	C.SDL_RenderGetClipRect(_renderer, _rect)
}

func (renderer *Renderer) SetScale(scaleX, scaleY float32) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_scaleX := (C.float)(scaleX)
	_scaleY := (C.float)(scaleY)
	return (int)(C.SDL_RenderSetScale(_renderer, _scaleX, _scaleY))
}

func (renderer *Renderer) GetScale() (scaleX, scaleY float32) {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_scaleX := (*C.float)(unsafe.Pointer(&scaleX))
	_scaleY := (*C.float)(unsafe.Pointer(&scaleY))
	C.SDL_RenderGetScale(_renderer, _scaleX, _scaleY)
	return
}

func (renderer *Renderer) SetDrawColor(r, g, b, a uint8) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_r := (C.Uint8)(r)
	_g := (C.Uint8)(g)
	_b := (C.Uint8)(b)
	_a := (C.Uint8)(a)
	return (int)(C.SDL_SetRenderDrawColor(_renderer, _r, _g, _b, _a))
}

func (renderer *Renderer) GetDrawColor() (r, g, b, a uint8, status int) {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_r := (*C.Uint8)(unsafe.Pointer(&r))
	_g := (*C.Uint8)(unsafe.Pointer(&g))
	_b := (*C.Uint8)(unsafe.Pointer(&b))
	_a := (*C.Uint8)(unsafe.Pointer(&a))
	status = (int)(C.SDL_GetRenderDrawColor(_renderer, _r, _g, _b, _a))
	return
}

func (renderer *Renderer) SetDrawBlendMode(blendMode uint32) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_blendMode := (C.SDL_BlendMode)(blendMode)
	return (int)(C.SDL_SetRenderDrawBlendMode(_renderer, _blendMode))
}

func (renderer *Renderer) GetDrawBlendMode(blendMode *uint32) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_blendMode := (*C.SDL_BlendMode)(unsafe.Pointer(blendMode))
	return (int)(C.SDL_GetRenderDrawBlendMode(_renderer, _blendMode))
}

func (renderer *Renderer) Clear() int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	return (int)(C.SDL_RenderClear(_renderer))
}

func (renderer *Renderer) DrawPoint(x, y int) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_x := (C.int)(x)
	_y := (C.int)(y)
	return (int)(C.SDL_RenderDrawPoint(_renderer, _x, _y))
}

func (renderer *Renderer) DrawPoints(points []Point) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_points := (*C.SDL_Point)(unsafe.Pointer(&points[0]))
	_count := (C.int)(len(points))
	return (int)(C.SDL_RenderDrawPoints(_renderer, _points, _count))
}

func (renderer *Renderer) DrawLine(x1, y1, x2, y2 int) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_x1 := (C.int)(x1)
	_y1 := (C.int)(y1)
	_x2 := (C.int)(x2)
	_y2 := (C.int)(y2)
	return (int)(C.SDL_RenderDrawLine(_renderer, _x1, _y1, _x2, _y2))
}

func (renderer *Renderer) DrawLines(points []Point) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_points := (*C.SDL_Point)(unsafe.Pointer(&points[0]))
	_count := (C.int)(len(points))
	return (int)(C.SDL_RenderDrawLines(_renderer, _points, _count))
}

func (renderer *Renderer) DrawRect(rect *Rect) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_rect := (*C.SDL_Rect)(unsafe.Pointer(rect))
	return (int)(C.SDL_RenderDrawRect(_renderer, _rect))
}

func (renderer *Renderer) DrawRects(rects []Rect) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_rects := (*C.SDL_Rect)(unsafe.Pointer(&rects[0]))
	_count := (C.int)(len(rects))
	return (int)(C.SDL_RenderDrawRects(_renderer, _rects, _count))
}

func (renderer *Renderer) FillRect(rect *Rect) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_rect := (*C.SDL_Rect)(unsafe.Pointer(rect))
	return (int)(C.SDL_RenderFillRect(_renderer, _rect))
}

func (renderer *Renderer) FillRects(rects []Rect) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_rects := (*C.SDL_Rect)(unsafe.Pointer(&rects[0]))
	_count := (C.int)(len(rects))
	return (int)(C.SDL_RenderFillRects(_renderer, _rects, _count))
}

func (renderer *Renderer) Copy(texture *Texture, srcrect, dstrect *Rect) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_texture := (*C.SDL_Texture)(unsafe.Pointer(texture))
	_srcrect := (*C.SDL_Rect)(unsafe.Pointer(srcrect))
	_dstrect := (*C.SDL_Rect)(unsafe.Pointer(dstrect))
	return (int)(C.SDL_RenderCopy(_renderer, _texture, _srcrect, _dstrect))
}

func (renderer *Renderer) CopyEx(texture *Texture, srcrect, dstrect *Rect, angle float64, center *Point, flip uint32) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_texture := (*C.SDL_Texture)(unsafe.Pointer(texture))
	_srcrect := (*C.SDL_Rect)(unsafe.Pointer(srcrect))
	_dstrect := (*C.SDL_Rect)(unsafe.Pointer(dstrect))
	_angle := (C.double)(angle)
	_center := (*C.SDL_Point)(unsafe.Pointer(center))
	_flip := (C.SDL_RendererFlip)(flip)
	return (int)(C.SDL_RenderCopyEx(_renderer, _texture, _srcrect, _dstrect, _angle, _center, _flip))
}

func (renderer *Renderer) ReadPixels(rect *Rect, format uint32, pixels unsafe.Pointer, pitch int) int {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	_rect := (*C.SDL_Rect)(unsafe.Pointer(rect))
	_format := (C.Uint32)(format)
	_pitch := (C.int)(pitch)
	return (int)(C.SDL_RenderReadPixels(_renderer, _rect, _format, pixels, _pitch))
}

func (renderer *Renderer) Present() {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	C.SDL_RenderPresent(_renderer)
}

func (texture *Texture) Destroy() {
	_texture := (*C.SDL_Texture)(unsafe.Pointer(texture))
	C.SDL_DestroyTexture(_texture)
}

func (renderer *Renderer) Destroy() {
	_renderer := (*C.SDL_Renderer)(unsafe.Pointer(renderer))
	C.SDL_DestroyRenderer(_renderer)
}

func (texture *Texture) GL_BindTexture(texw, texh *float32) int {
	_texture := (*C.SDL_Texture)(unsafe.Pointer(texture))
	_texw := (*C.float)(unsafe.Pointer(texw))
	_texh := (*C.float)(unsafe.Pointer(texh))
	return (int)(C.SDL_GL_BindTexture(_texture, _texw, _texh))
}

func (texture *Texture) GL_UnbindTexture() int {
	_texture := (*C.SDL_Texture)(unsafe.Pointer(texture))
	return (int)(C.SDL_GL_UnbindTexture(_texture))
}
