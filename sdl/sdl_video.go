package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"

type GLContext unsafe.Pointer

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
