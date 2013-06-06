package sdl

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL.h>
import "C"
import "unsafe"

type GLContext unsafe.Pointer

/* SDL.h */
func Init(flags uint32) int {
	_flags := (C.Uint32) (flags)
	return (int) (C.SDL_Init(_flags))
}

func Quit() {
	C.SDL_Quit()
}

func InitSubSystem(flags uint32) int {
	_flags := (C.Uint32) (flags)
	return (int) (C.SDL_InitSubSystem(_flags))
}

func QuitSubSystem(flags uint32) {
	_flags := (C.Uint32) (flags)
	C.SDL_QuitSubSystem(_flags)
}

func WasInit(flags uint32) uint32 {
	_flags := (C.Uint32) (flags)
	return (uint32) (C.SDL_WasInit(_flags))
}

func GetVersion(v *Version) {
	version := (*C.SDL_version) (unsafe.Pointer(v))
	C.SDL_GetVersion(version)
}

func GetRevision() string {
	return (string) (C.GoString(C.SDL_GetRevision()))
}

func GetRevisionNumber() int {
	return (int) (C.SDL_GetRevisionNumber())
}

func GetPlatform() string {
	return (string) (C.GoString(C.SDL_GetPlatform()))
}

/* SDL_video.h */
func GetNumVideoDrivers() int {
	return (int) (C.SDL_GetNumVideoDrivers())
}

func GetVideoDriver(index int) string {
	_index := (C.int) (index)
	return (string) (C.GoString(C.SDL_GetVideoDriver(_index)))
}

func VideoInit(driverName string) int {
	return (int) (C.SDL_VideoInit(C.CString(driverName)))
}

func VideoQuit() {
	C.SDL_VideoQuit()
}

func GetCurrentVideoDriver() string {
	return (string) (C.GoString(C.SDL_GetCurrentVideoDriver()))
}

func GetNumDisplayModes(displayIndex int) int {
	_displayIndex := (C.int) (displayIndex)
	return (int) (C.SDL_GetNumDisplayModes(_displayIndex))
}

func GetDisplayMode(displayIndex int, modeIndex int, mode *DisplayMode) int {
	_displayIndex := (C.int) (displayIndex)
	_modeIndex := (C.int) (modeIndex)
	_mode := (*C.SDL_DisplayMode) (unsafe.Pointer(mode))
	return (int) (C.SDL_GetDisplayMode(_displayIndex, _modeIndex, _mode))
}

func GetDesktopDisplayMode(displayIndex int, mode *DisplayMode) int {
	_displayIndex := (C.int) (displayIndex)
	_mode := (*C.SDL_DisplayMode) (unsafe.Pointer(mode))
	return (int) (C.SDL_GetDesktopDisplayMode(_displayIndex, _mode))
}

func GetCurrentDisplayMode(displayIndex int, mode *DisplayMode) int {
	_displayIndex := (C.int) (displayIndex)
	_mode := (*C.SDL_DisplayMode) (unsafe.Pointer(mode))
	return (int) (C.SDL_GetCurrentDisplayMode(_displayIndex, _mode))
}

func GetClosestDisplayMode(displayIndex int, mode *DisplayMode, closest *DisplayMode) *DisplayMode {
	_displayIndex := (C.int) (displayIndex)
	_mode := (*C.SDL_DisplayMode) (unsafe.Pointer(mode))
	_closest := (*C.SDL_DisplayMode) (unsafe.Pointer(closest))
	return (*DisplayMode) (unsafe.Pointer((C.SDL_GetClosestDisplayMode(_displayIndex,
									_mode, _closest))))
}

func GetWindowDisplayIndex(window *Window) int {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	return (int) (C.SDL_GetWindowDisplayIndex(_window))
}

func SetWindowDisplayMode(window *Window, mode *DisplayMode) int {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_mode := (*C.SDL_DisplayMode) (unsafe.Pointer(mode))
	return (int) (C.SDL_SetWindowDisplayMode(_window, _mode))
}

func GetWindowDisplayMode(window *Window, mode *DisplayMode) int {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_mode := (*C.SDL_DisplayMode) (unsafe.Pointer(mode))
	return (int) (C.SDL_GetWindowDisplayMode(_window, _mode))
}

func (window *Window) GetPixelFormat() uint32 {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	return (uint32) (C.SDL_GetWindowPixelFormat(_window))
}

func CreateWindow(title string, x int, y int, w int, h int, flags uint32) *Window {
	var window = C.SDL_CreateWindow(C.CString(title), C.int(x), C.int(y),
					C.int(w), C.int(h), C.Uint32(flags))
	return (*Window) (unsafe.Pointer(window))
}

func CreateWindowFrom(data unsafe.Pointer) *Window {
	return (*Window) (unsafe.Pointer(C.SDL_CreateWindowFrom(data)))
}

func (window *Window) GetID() uint32 {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	return (uint32) (C.SDL_GetWindowID(_window))
}

func GetWindowFromID(id uint32) *Window {
	_id := (C.Uint32) (id)
	return (*Window) (unsafe.Pointer((C.SDL_GetWindowFromID(_id))))
}

func (window *Window) GetFlags() uint32 {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	return (uint32) (C.SDL_GetWindowFlags(_window))
}

func (window *Window) SetTitle(title string) {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_title := (C.CString) (title)
	C.SDL_SetWindowTitle(_window, _title)
}

func (window *Window) GetTitle() string {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	return (string) (C.GoString(C.SDL_GetWindowTitle(_window)))
}

func (window *Window) SetIcon(icon *Surface) {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_icon := (*C.SDL_Surface) (unsafe.Pointer(window))
	C.SDL_SetWindowIcon(_window, _icon)
}

func (window *Window) SetData(name string, userdata unsafe.Pointer) unsafe.Pointer {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_name := (C.CString) (name)
	return (unsafe.Pointer) (C.SDL_SetWindowData(_window, _name, userdata))
}

func (window *Window) GetData(name string, userdata unsafe.Pointer) unsafe.Pointer {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_name := (C.CString) (name)
	return (unsafe.Pointer) (C.SDL_GetWindowData(_window, _name))
}

func (window *Window) SetPosition(x int, y int) {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	C.SDL_SetWindowPosition(_window, C.int(x), C.int(y))
}

func (window *Window) GetPosition(x *int, y *int) {
	var _x, _y C.int;
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	C.SDL_GetWindowPosition(_window, &_x, &_y)
	*x = int(_x);
	*y = int(_y);
}

func (window *Window) SetSize(w int, h int) {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	C.SDL_SetWindowSize(_window, C.int(w), C.int(h))
}

func (window *Window) GetSize(w *int, h *int) {
	var _w, _h C.int;
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	C.SDL_GetWindowSize(_window, &_w, &_h)
	*w = int(_w);
	*h = int(_h);
}

func (window *Window) SetMinimumSize(min_w int, min_h int) {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_min_w := (C.int) (min_w)
	_min_h := (C.int) (min_h)
	C.SDL_SetWindowMinimumSize(_window, _min_w, _min_h)
}

func (window *Window) GetMinimumSize(w *int, h *int) {
	var _w, _h C.int;
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	C.SDL_GetWindowMinimumSize(_window, &_w, &_h)
	*w = int(_w);
	*h = int(_h);
}

func (window *Window) SetMaximumSize(max_w int, max_h int) {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_max_w := (C.int) (max_w)
	_max_h := (C.int) (max_h)
	C.SDL_SetWindowMaximumSize(_window, _max_w, _max_h)
}

func (window *Window) GetMaximumSize(w *int, h *int) {
	var _w, _h C.int;
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	C.SDL_GetWindowMaximumSize(_window, &_w, &_h)
	*w = int(_w);
	*h = int(_h);
}

func (window *Window) SetBordered(bordered bool) {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_bordered := (C.SDL_bool) (btoi(bordered))
	C.SDL_SetWindowBordered(_window, _bordered)
}

func (window *Window) Show() {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	C.SDL_ShowWindow(_window)
}

func (window *Window) Hide() {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	C.SDL_HideWindow(_window)
}

func (window *Window) Raise() {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	C.SDL_RaiseWindow(_window)
}

func (window *Window) Maximize() {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	C.SDL_MaximizeWindow(_window)
}

func (window *Window) Minimize() {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	C.SDL_MinimizeWindow(_window)
}

func (window *Window) Restore() {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	C.SDL_RestoreWindow(_window)
}

func (window *Window) SetFullscreen(flags uint32) int {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_flags := (C.Uint32) (flags)
	return (int) (C.SDL_SetWindowFullscreen(_window, _flags))
}

func (window *Window) GetSurface() *Surface {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	return (*Surface) (unsafe.Pointer(C.SDL_GetWindowSurface(_window)))
}

func (window *Window) UpdateSurface() int {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	return (int) (C.SDL_UpdateWindowSurface(_window))
}

func (window *Window) UpdateSurfaceRects(rects *Rect, numrects int) int {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_rects := (*C.SDL_Rect) (unsafe.Pointer(rects))
	_numrects := (C.int) (numrects)
	return (int) (C.SDL_UpdateWindowSurfaceRects(_window, _rects, _numrects))
}

func (window *Window) SetGrab(grabbed bool) {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_grabbed := (C.SDL_bool) (btoi(grabbed))
	C.SDL_SetWindowGrab(_window, _grabbed)
}

func (window *Window) GetGrab() bool {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	return (bool) (itob(int(C.SDL_GetWindowGrab(_window))))
}

func (window *Window) SetBrightness(brightness float32) int {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_brightness := (C.float) (brightness)
	return (int) (C.SDL_SetWindowBrightness(_window, _brightness))
}

func (window *Window) GetBrightness() float32 {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	return (float32) (C.SDL_GetWindowBrightness(_window))
}

func (window *Window) SetGammaRamp(red, green, blue *uint16) int {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_red := (*C.Uint16) (red)
	_green := (*C.Uint16) (red)
	_blue := (*C.Uint16) (blue)
	return (int) (C.SDL_SetWindowGammaRamp(_window, _red, _green, _blue))
}

func (window *Window) GetGammaRamp(red, green, blue *uint16) int {
	var _red, _green, _blue *C.Uint16
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	status := (int) (C.SDL_GetWindowGammaRamp(_window, _red, _green, _blue))
	*red = (uint16) (*_red)
	*green = (uint16) (*_green)
	*blue = (uint16) (*_blue)
	return status;
}

func IsScreenSaverEnabled() bool {
	return (bool) (itob(int(C.SDL_IsScreenSaverEnabled())))
}

func EnableScreenSaver() {
	C.SDL_EnableScreenSaver()
}

func DisableScreenSaver() {
	C.SDL_DisableScreenSaver()
}

func GL_LoadLibrary(path string) int {
	_path := (C.CString) (path)
	return (int) (C.SDL_GL_LoadLibrary(_path))
}

func GL_GetProcAddress(proc string) unsafe.Pointer {
	_proc := (C.CString) (proc)
	return (unsafe.Pointer) (C.SDL_GL_GetProcAddress(_proc))
}

func GL_UnloadLibrary() {
	C.SDL_GL_UnloadLibrary()
}

func GL_ExtensionSupported(extension string) bool {
	_extension := (C.CString) (extension)
	return (bool) (itob(int(C.SDL_GL_ExtensionSupported(_extension))))
}

func GL_SetAttribute(attribute uint32, value int) int {
	_attribute := (C.SDL_GLattr) (attribute)
	_value := (C.int) (value)
	return (int) (C.SDL_GL_SetAttribute(_attribute, _value))
}

func GL_GetAttribute(attribute uint32, value *int) int {
	_attribute := (C.SDL_GLattr) (attribute)
	_value := (*C.int) (unsafe.Pointer(value))
	return (int) (C.SDL_GL_GetAttribute(_attribute, _value))
}

func GL_CreateContext(window *Window) GLContext {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	return (GLContext) (C.SDL_GL_CreateContext(_window))
}

func GL_MakeCurrent(window *Window, glcontext GLContext) int {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_glcontext := (C.SDL_GLContext) (glcontext)
	return (int) (C.SDL_GL_MakeCurrent(_window, _glcontext))
}

func GL_SetSwapInterval(interval int) int {
	_interval := (C.int) (interval)
	return (int) (C.SDL_GL_SetSwapInterval(_interval))
}

func GL_GetSwapInterval() int {
	return (int) (C.SDL_GL_GetSwapInterval())
}

func GL_SwapWindow(window *Window) {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	C.SDL_GL_SwapWindow(_window)
}

func GL_DeleteContext(context GLContext) {
	_context := (C.SDL_GLContext) (context)
	C.SDL_GL_DeleteContext(_context)
}

/* SDL_error.h */
func GetError() string {
	return (string) (C.GoString(C.SDL_GetError()))
}

func ClearError() {
	C.SDL_ClearError()
}

func Error(code int) {
	_code := (C.SDL_errorcode) (C.int(code))
	C.SDL_Error(_code)
}

func OutOfMemory() {
	Error(ENOMEM)
}

func Unsupported() {
	Error(UNSUPPORTED)
}

/* SDL_render.h */
func GetNumRenderDrivers() int {
	return (int) (C.SDL_GetNumRenderDrivers())
}

func GetRenderDriverInfo(index int, info *RendererInfo) int {
	_index := (C.int) (index)
	_info := (*C.SDL_RendererInfo) (unsafe.Pointer(info))
	return (int) (C.SDL_GetRenderDriverInfo(_index, _info))
}

func CreateWindowAndRenderer(width int, height int, windowFlags uint32, window **Window, renderer **Renderer) int {
	_width := (C.int) (width)
	_height := (C.int) (height)
	_windowFlags := (C.Uint32) (windowFlags)
	_window := (**C.SDL_Window) (unsafe.Pointer(window))
	_renderer := (**C.SDL_Renderer) (unsafe.Pointer(renderer))
	return (int) (C.SDL_CreateWindowAndRenderer(_width, _height, _windowFlags, _window, _renderer))
}

func CreateRenderer(window *Window, index int, flags uint32) *Renderer {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_index := (C.int) (index)
	_flags := (C.Uint32) (flags)
	return (*Renderer) (unsafe.Pointer(C.SDL_CreateRenderer(_window, _index, _flags)))
}

func CreateSoftwareRenderer(surface *Surface) *Renderer {
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	return (*Renderer) (unsafe.Pointer(C.SDL_CreateSoftwareRenderer(_surface)))
}

func (window *Window) GetRenderer() *Renderer {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	return (*Renderer) (unsafe.Pointer(C.SDL_GetRenderer(_window)))
}

func (renderer *Renderer) GetRendererInfo(info *RendererInfo) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_info := (*C.SDL_RendererInfo) (unsafe.Pointer(info))
	return (int) (C.SDL_GetRendererInfo(_renderer, _info))
}

func (renderer *Renderer) GetRendererOutputSize(w *int, h *int) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_w := (*C.int) (unsafe.Pointer(w))
	_h := (*C.int) (unsafe.Pointer(h))
	return (int) (C.SDL_GetRendererOutputSize(_renderer, _w, _h))
}

func CreateTexture(renderer *Renderer, format uint32, access int, w int, h int) *Texture {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_format := (C.Uint32) (format)
	_access := (C.int) (access)
	_w := (C.int) (w)
	_h := (C.int) (h)
	return (*Texture) (unsafe.Pointer(C.SDL_CreateTexture(_renderer, _format, _access, _w, _h)))
}

func (renderer *Renderer) CreateTextureFromSurface(surface *Surface) *Texture {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_surface := (*C.SDL_Surface) (unsafe.Pointer(surface))
	return (*Texture) (unsafe.Pointer(C.SDL_CreateTextureFromSurface(_renderer, _surface)))
}

func QueryTexture(texture *Texture, format *uint32, access *int, w *int, h *int) int {
	_texture := (*C.SDL_Texture) (unsafe.Pointer(texture))
	_format := (*C.Uint32) (unsafe.Pointer(access))
	_access := (*C.int) (unsafe.Pointer(access))
	_w := (*C.int) (unsafe.Pointer(w))
	_h := (*C.int) (unsafe.Pointer(h))
	return (int) (C.SDL_QueryTexture(_texture, _format, _access, _w, _h))
}

func (texture *Texture) SetColorMod(r uint8, g uint8, b uint8) int {
	_texture := (*C.SDL_Texture) (unsafe.Pointer(texture))
	_r := (C.Uint8) (r)
	_g := (C.Uint8) (g)
	_b := (C.Uint8) (b)
	return (int) (C.SDL_SetTextureColorMod(_texture, _r, _g, _b))
}

func (texture *Texture) SetAlphaMod(alpha uint8) int {
	_texture := (*C.SDL_Texture) (unsafe.Pointer(texture))
	_alpha := (C.Uint8) (alpha)
	return (int) (C.SDL_SetTextureAlphaMod(_texture, _alpha))
}

func (texture *Texture) SetBlendMode(blendMode uint32) int {
	_texture := (*C.SDL_Texture) (unsafe.Pointer(texture))
	_blendMode := (C.SDL_BlendMode) (C.Uint32(blendMode))
	return (int) (C.SDL_SetTextureBlendMode(_texture, _blendMode))
}

func (texture *Texture) GetBlendMode(blendMode *uint32) int {
	_texture := (*C.SDL_Texture) (unsafe.Pointer(texture))
	_blendMode := (*C.SDL_BlendMode) (unsafe.Pointer(blendMode))
	return (int) (C.SDL_GetTextureBlendMode(_texture, _blendMode))
}

func (texture *Texture) Update(rect *Rect, pixels unsafe.Pointer, pitch int) int {
	_texture := (*C.SDL_Texture) (unsafe.Pointer(texture))
	_rect := (*C.SDL_Rect) (unsafe.Pointer(rect))
	_pitch := (C.int) (pitch)
	return (int) (C.SDL_UpdateTexture(_texture, _rect, pixels, _pitch))
}

func (texture *Texture) Lock(rect *Rect, pixels unsafe.Pointer, pitch *int) int {
	_texture := (*C.SDL_Texture) (unsafe.Pointer(texture))
	_rect := (*C.SDL_Rect) (unsafe.Pointer(rect))
	_pitch := (*C.int) (unsafe.Pointer(pitch))
	return (int) (C.SDL_LockTexture(_texture, _rect, &pixels, _pitch))
}

func (texture *Texture) Unlock() {
	_texture := (*C.SDL_Texture) (unsafe.Pointer(texture))
	(C.SDL_UnlockTexture(_texture))
}

func (renderer *Renderer) RenderTargetSupported() bool {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	return (bool) (itob(int(C.SDL_RenderTargetSupported(_renderer))))
}

func (renderer *Renderer) SetRenderTarget(texture *Texture) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_texture := (*C.SDL_Texture) (unsafe.Pointer(texture))
	return (int) (C.SDL_SetRenderTarget(_renderer, _texture))
}

func (renderer *Renderer) GetRenderTarget() *Texture {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	return (*Texture) (unsafe.Pointer(C.SDL_GetRenderTarget(_renderer)))
}

func (renderer *Renderer) SetLogicalSize(w int, h int) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_w := (C.int) (w)
	_h := (C.int) (h)
	return (int) (C.SDL_RenderSetLogicalSize(_renderer, _w, _h))
}

func (renderer *Renderer) SetViewport(rect *Rect) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_rect := (*C.SDL_Rect) (unsafe.Pointer(rect))
	return (int) (C.SDL_RenderSetViewport(_renderer, _rect))
}

func (renderer *Renderer) GetViewport(rect *Rect) {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_rect := (*C.SDL_Rect) (unsafe.Pointer(rect))
	C.SDL_RenderGetViewport(_renderer, _rect)
}

func (renderer *Renderer) SetClipRect(rect *Rect) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_rect := (*C.SDL_Rect) (unsafe.Pointer(rect))
	return (int) (C.SDL_RenderSetClipRect(_renderer, _rect))
}

func (renderer *Renderer) GetClipRect(rect *Rect) {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_rect := (*C.SDL_Rect) (unsafe.Pointer(rect))
	C.SDL_RenderGetClipRect(_renderer, _rect)
}

func (renderer *Renderer) SetScale(scaleX, scaleY float32) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_scaleX := (C.float) (scaleX)
	_scaleY := (C.float) (scaleY)
	return (int) (C.SDL_RenderSetScale(_renderer, _scaleX, _scaleY))
}

func (renderer *Renderer) GetScale(scaleX, scaleY *float32) {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_scaleX := (*C.float) (unsafe.Pointer(scaleX))
	_scaleY := (*C.float) (unsafe.Pointer(scaleY))
	C.SDL_RenderGetScale(_renderer, _scaleX, _scaleY)
}

func (renderer *Renderer) SetDrawColor(r, g, b, a uint8) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_r := (C.Uint8) (r)
	_g := (C.Uint8) (g)
	_b := (C.Uint8) (b)
	_a := (C.Uint8) (a)
	return (int) (C.SDL_SetRenderDrawColor(_renderer, _r, _g, _b, _a))
}

func (renderer *Renderer) GetDrawColor(r, g, b, a *uint8) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_r := (*C.Uint8) (unsafe.Pointer(r))
	_g := (*C.Uint8) (unsafe.Pointer(g))
	_b := (*C.Uint8) (unsafe.Pointer(b))
	_a := (*C.Uint8) (unsafe.Pointer(a))
	return (int) (C.SDL_GetRenderDrawColor(_renderer, _r, _g, _b, _a))
}

func (renderer *Renderer) SetDrawBlendMode(blendMode uint32) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_blendMode := (C.SDL_BlendMode) (blendMode)
	return (int) (C.SDL_SetRenderDrawBlendMode(_renderer, _blendMode))
}

func (renderer *Renderer) GetDrawBlendMode(blendMode *uint32) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_blendMode := (*C.SDL_BlendMode) (unsafe.Pointer(blendMode))
	return (int) (C.SDL_GetRenderDrawBlendMode(_renderer, _blendMode))
}

func (renderer *Renderer) Clear() int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	return (int) (C.SDL_RenderClear(_renderer))
}

func (renderer *Renderer) DrawPoint(x, y int) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_x := (C.int) (x)
	_y := (C.int) (y)
	return (int) (C.SDL_RenderDrawPoint(_renderer, _x, _y))
}

func (renderer *Renderer) DrawPoints(points *Point, count int) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_points := (*C.SDL_Point) (unsafe.Pointer(points))
	_count := (C.int) (count)
	return (int) (C.SDL_RenderDrawPoints(_renderer, _points, _count))
}

func (renderer *Renderer) DrawLine(x1, y1, x2, y2 int) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_x1 := (C.int) (x1)
	_y1 := (C.int) (y1)
	_x2 := (C.int) (x2)
	_y2 := (C.int) (y2)
	return (int) (C.SDL_RenderDrawLine(_renderer, _x1, _y1, _x2, _y2))
}

func (renderer *Renderer) DrawLines(points *[]Point, count int) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_points := (*C.SDL_Point) (unsafe.Pointer(&(*points)[0]))
	_count := (C.int) (count)
	return (int) (C.SDL_RenderDrawLines(_renderer, _points, _count))
}

func (renderer *Renderer) DrawRect(rect *Rect) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_rect := (*C.SDL_Rect) (unsafe.Pointer(rect))
	return (int) (C.SDL_RenderDrawRect(_renderer, _rect))
}

func (renderer *Renderer) DrawRects(rects *[]Rect, count int) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_rects := (*C.SDL_Rect) (unsafe.Pointer(&(*rects)[0]))
	_count := (C.int) (count)
	return (int) (C.SDL_RenderDrawRects(_renderer, _rects, _count))
}

func (renderer *Renderer) FillRect(rect *Rect) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_rect := (*C.SDL_Rect) (unsafe.Pointer(rect))
	return (int) (C.SDL_RenderFillRect(_renderer, _rect))
}

func (renderer *Renderer) FillRects(rects *[]Rect, count int) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_rects := (*C.SDL_Rect) (unsafe.Pointer(&(*rects)[0]))
	_count := (C.int) (count)
	return (int) (C.SDL_RenderFillRects(_renderer, _rects, _count))
}

func (renderer *Renderer) Copy(texture *Texture, srcrect, dstrect *Rect) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_texture := (*C.SDL_Texture) (unsafe.Pointer(texture))
	_srcrect := (*C.SDL_Rect) (unsafe.Pointer(srcrect))
	_dstrect := (*C.SDL_Rect) (unsafe.Pointer(dstrect))
	return (int) (C.SDL_RenderCopy(_renderer, _texture, _srcrect, _dstrect))
}

func (renderer *Renderer) CopyEx(texture *Texture, srcrect, dstrect *Rect, angle float64, center *Point, flip uint32) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_texture := (*C.SDL_Texture) (unsafe.Pointer(texture))
	_srcrect := (*C.SDL_Rect) (unsafe.Pointer(srcrect))
	_dstrect := (*C.SDL_Rect) (unsafe.Pointer(dstrect))
	_angle := (C.double) (angle)
	_center := (*C.SDL_Point) (unsafe.Pointer(center))
	_flip := (C.SDL_RendererFlip) (flip)
	return (int) (C.SDL_RenderCopyEx(_renderer, _texture, _srcrect, _dstrect, _angle, _center, _flip))
}

func (renderer *Renderer) ReadPixels(rect *Rect, format uint32, pixels unsafe.Pointer, pitch int) int {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	_rect := (*C.SDL_Rect) (unsafe.Pointer(rect))
	_format := (C.Uint32) (format)
	_pitch := (C.int) (pitch)
	return (int) (C.SDL_RenderReadPixels(_renderer, _rect, _format, pixels, _pitch))
}

func (renderer *Renderer) Present() {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	C.SDL_RenderPresent(_renderer)
}

func DestroyTexture(texture *Texture) {
	_texture := (*C.SDL_Texture) (unsafe.Pointer(texture))
	C.SDL_DestroyTexture(_texture)
}

func DestroyRenderer(renderer *Renderer) {
	_renderer := (*C.SDL_Renderer) (unsafe.Pointer(renderer))
	C.SDL_DestroyRenderer(_renderer)
}

func (texture *Texture) GL_BindTexture(texw, texh *float32) int {
	_texture := (*C.SDL_Texture) (unsafe.Pointer(texture))
	_texw := (*C.float) (unsafe.Pointer(texw))
	_texh := (*C.float) (unsafe.Pointer(texh))
	return (int) (C.SDL_GL_BindTexture(_texture, _texw, _texh))
}

func (texture *Texture) GL_UnbindTexture() int {
	_texture := (*C.SDL_Texture) (unsafe.Pointer(texture))
	return (int) (C.SDL_GL_UnbindTexture(_texture))
}

/* SDL_timer.h */
func Delay(ms uint32) {
	_ms := (C.Uint32) (ms)
	C.SDL_Delay(C.Uint32(_ms));
}
