package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

const (
	WINDOW_FULLSCREEN         = C.SDL_WINDOW_FULLSCREEN
	WINDOW_OPENGL             = C.SDL_WINDOW_OPENGL
	WINDOW_SHOWN              = C.SDL_WINDOW_SHOWN
	WINDOW_HIDDEN             = C.SDL_WINDOW_HIDDEN
	WINDOW_BORDERLESS         = C.SDL_WINDOW_BORDERLESS
	WINDOW_RESIZABLE          = C.SDL_WINDOW_RESIZABLE
	WINDOW_MINIMIZED          = C.SDL_WINDOW_MINIMIZED
	WINDOW_MAXIMIZED          = C.SDL_WINDOW_MAXIMIZED
	WINDOW_INPUT_GRABBED      = C.SDL_WINDOW_INPUT_GRABBED
	WINDOW_INPUT_FOCUS        = C.SDL_WINDOW_INPUT_FOCUS
	WINDOW_MOUSE_FOCUS        = C.SDL_WINDOW_MOUSE_FOCUS
	WINDOW_FULLSCREEN_DESKTOP = C.SDL_WINDOW_FULLSCREEN_DESKTOP
	WINDOW_FOREIGN            = C.SDL_WINDOW_FOREIGN
)

const (
	WINDOWEVENT_NONE         = C.SDL_WINDOWEVENT_NONE
	WINDOWEVENT_SHOWN        = C.SDL_WINDOWEVENT_SHOWN
	WINDOWEVENT_HIDDEN       = C.SDL_WINDOWEVENT_HIDDEN
	WINDOWEVENT_EXPOSED      = C.SDL_WINDOWEVENT_EXPOSED
	WINDOWEVENT_MOVED        = C.SDL_WINDOWEVENT_MOVED
	WINDOWEVENT_RESIZED      = C.SDL_WINDOWEVENT_RESIZED
	WINDOWEVENT_SIZE_CHANGED = C.SDL_WINDOWEVENT_SIZE_CHANGED
	WINDOWEVENT_MINIMIZED    = C.SDL_WINDOWEVENT_MINIMIZED
	WINDOWEVENT_MAXIMIZED    = C.SDL_WINDOWEVENT_MAXIMIZED
	WINDOWEVENT_RESTORED     = C.SDL_WINDOWEVENT_RESTORED
	WINDOWEVENT_ENTER        = C.SDL_WINDOWEVENT_ENTER
	WINDOWEVENT_LEAVE        = C.SDL_WINDOWEVENT_LEAVE
	WINDOWEVENT_FOCUS_GAINED = C.SDL_WINDOWEVENT_FOCUS_GAINED
	WINDOWEVENT_FOCUS_LOST   = C.SDL_WINDOWEVENT_FOCUS_LOST
	WINDOWEVENT_CLOSE        = C.SDL_WINDOWEVENT_CLOSE
)

const (
	WINDOWPOS_UNDEFINED_MASK = C.SDL_WINDOWPOS_UNDEFINED_MASK
	WINDOWPOS_UNDEFINED      = C.SDL_WINDOWPOS_UNDEFINED
	WINDOWPOS_CENTERED_MASK  = C.SDL_WINDOWPOS_CENTERED_MASK
	WINDOWPOS_CENTERED       = C.SDL_WINDOWPOS_CENTERED
)

const (
	GL_RED_SIZE                   = C.SDL_GL_RED_SIZE
	GL_GREEN_SIZE                 = C.SDL_GL_GREEN_SIZE
	GL_BLUE_SIZE                  = C.SDL_GL_BLUE_SIZE
	GL_ALPHA_SIZE                 = C.SDL_GL_ALPHA_SIZE
	GL_BUFFER_SIZE                = C.SDL_GL_BUFFER_SIZE
	GL_DOUBLEBUFFER               = C.SDL_GL_DOUBLEBUFFER
	GL_DEPTH_SIZE                 = C.SDL_GL_DEPTH_SIZE
	GL_STENCIL_SIZE               = C.SDL_GL_STENCIL_SIZE
	GL_ACCUM_RED_SIZE             = C.SDL_GL_ACCUM_RED_SIZE
	GL_ACCUM_GREEN_SIZE           = C.SDL_GL_ACCUM_GREEN_SIZE
	GL_ACCUM_BLUE_SIZE            = C.SDL_GL_ACCUM_BLUE_SIZE
	GL_ACCUM_ALPHA_SIZE           = C.SDL_GL_ALPHA_SIZE
	GL_STEREO                     = C.SDL_GL_STEREO
	GL_MULTISAMPLEBUFFERS         = C.SDL_GL_MULTISAMPLEBUFFERS
	GL_MULTISAMPLESAMPLES         = C.SDL_GL_MULTISAMPLESAMPLES
	GL_ACCELERATED_VISUAL         = C.SDL_GL_ACCELERATED_VISUAL
	GL_RETAINED_BACKING           = C.SDL_GL_RETAINED_BACKING
	GL_CONTEXT_MAJOR_VERSION      = C.SDL_GL_CONTEXT_MAJOR_VERSION
	GL_CONTEXT_MINOR_VERSION      = C.SDL_GL_CONTEXT_MINOR_VERSION
	GL_CONTEXT_EGL                = C.SDL_GL_CONTEXT_EGL
	GL_CONTEXT_FLAGS              = C.SDL_GL_CONTEXT_FLAGS
	GL_CONTEXT_PROFILE_MASK       = C.SDL_GL_CONTEXT_PROFILE_MASK
	GL_SHARE_WITH_CURRENT_CONTEXT = C.SDL_GL_SHARE_WITH_CURRENT_CONTEXT
)

const (
	GL_CONTEXT_PROFILE_CORE          = C.SDL_GL_CONTEXT_PROFILE_CORE
	GL_CONTEXT_PROFILE_COMPATIBILITY = C.SDL_GL_CONTEXT_PROFILE_COMPATIBILITY
	GL_CONTEXT_PROFILE_ES            = C.SDL_GL_CONTEXT_PROFILE_ES
)

const (
	GL_CONTEXT_DEBUG_FLAG              = C.SDL_GL_CONTEXT_DEBUG_FLAG
	GL_CONTEXT_FORWARD_COMPATIBLE_FLAG = C.SDL_GL_CONTEXT_FORWARD_COMPATIBLE_FLAG
	GL_CONTEXT_ROBUST_ACCESS_FLAG      = C.SDL_GL_CONTEXT_ROBUST_ACCESS_FLAG
	GL_CONTEXT_RESET_ISOLATION_FLAG    = C.SDL_GL_CONTEXT_RESET_ISOLATION_FLAG
)

// DisplayMode (https://wiki.libsdl.org/SDL_DisplayMode)
type DisplayMode struct {
	Format      uint32
	W           int32
	H           int32
	RefreshRate int32
	DriverData  unsafe.Pointer
}
type cDisplayMode C.SDL_DisplayMode

type Window C.SDL_Window
type GLContext C.SDL_GLContext

// GLattr (https://wiki.libsdl.org/SDL_GLattr)
type GLattr C.SDL_GLattr

func (w *Window) cptr() *C.SDL_Window {
	return (*C.SDL_Window)(unsafe.Pointer(w))
}

func (dm *DisplayMode) cptr() *C.SDL_DisplayMode {
	return (*C.SDL_DisplayMode)(unsafe.Pointer(dm))
}

func (attr GLattr) c() C.SDL_GLattr {
	return C.SDL_GLattr(attr)
}

// GetNumVideoDisplays (https://wiki.libsdl.org/SDL_GetNumVideoDisplays)
func GetNumVideoDisplays() (int, error) {
	n := int(C.SDL_GetNumVideoDisplays())
	if n < 0 {
		return n, GetError()
	}
	return n, nil
}

// GetNumVideoDrivers (https://wiki.libsdl.org/SDL_GetNumVideoDrivers)
func GetNumVideoDrivers() int {
	return int(C.SDL_GetNumVideoDrivers())
}

// GetVideoDriver (https://wiki.libsdl.org/SDL_GetVideoDriver)
func GetVideoDriver(index int) string {
	return string(C.GoString(C.SDL_GetVideoDriver(C.int(index))))
}

// VideoInit (https://wiki.libsdl.org/SDL_VideoInit)
func VideoInit(driverName string) int {
	_driverName := C.CString(driverName)
	return int(C.SDL_VideoInit(_driverName))
}

// VideoQuit (https://wiki.libsdl.org/SDL_VideoQuit)
func VideoQuit() {
	C.SDL_VideoQuit()
}

// GetCurrentVideoDriver (https://wiki.libsdl.org/SDL_GetCurrentVideoDriver)
func GetCurrentVideoDriver() string {
	return string(C.GoString(C.SDL_GetCurrentVideoDriver()))
}

// GetNumDisplayModes (https://wiki.libsdl.org/SDL_GetNumDisplayModes)
func GetNumDisplayModes(displayIndex int) int {
	return int(C.SDL_GetNumDisplayModes(C.int(displayIndex)))
}

// GetDisplayBounds (https://wiki.libsdl.org/SDL_GetDisplayBounds)
func GetDisplayBounds(displayIndex int, rect *Rect) int {
	return int(C.SDL_GetDisplayBounds(C.int(displayIndex), rect.cptr()))
}

// GetDisplayMode (https://wiki.libsdl.org/SDL_GetDisplayMode)
func GetDisplayMode(displayIndex int, modeIndex int, mode *DisplayMode) int {
	return int(C.SDL_GetDisplayMode(C.int(displayIndex), C.int(modeIndex), mode.cptr()))
}

// GetDesktopDisplayMode (https://wiki.libsdl.org/SDL_GetDesktopDisplayMode)
func GetDesktopDisplayMode(displayIndex int, mode *DisplayMode) int {
	return int(C.SDL_GetDesktopDisplayMode(C.int(displayIndex), mode.cptr()))
}

// GetCurrentDisplayMode (https://wiki.libsdl.org/SDL_GetCurrentDisplayMode)
func GetCurrentDisplayMode(displayIndex int, mode *DisplayMode) int {
	return int(C.SDL_GetCurrentDisplayMode(C.int(displayIndex), mode.cptr()))
}

// GetClosestDisplayMode (https://wiki.libsdl.org/SDL_GetClosestDisplayMode)
func GetClosestDisplayMode(displayIndex int, mode *DisplayMode, closest *DisplayMode) *DisplayMode {
	return (*DisplayMode)(unsafe.Pointer((C.SDL_GetClosestDisplayMode(C.int(displayIndex), mode.cptr(), closest.cptr()))))
}

// Window (https://wiki.libsdl.org/SDL_GetWindowDisplayIndex)
func (window *Window) GetDisplayIndex() int {
	return int(C.SDL_GetWindowDisplayIndex(window.cptr()))
}

// Window (https://wiki.libsdl.org/SDL_SetWindowDisplayMode)
func (window *Window) SetDisplayMode(mode *DisplayMode) int {
	return int(C.SDL_SetWindowDisplayMode(window.cptr(), mode.cptr()))
}

// Window (https://wiki.libsdl.org/SDL_GetWindowDisplayMode)
func (window *Window) GetDisplayMode(mode *DisplayMode) int {
	return int(C.SDL_GetWindowDisplayMode(window.cptr(), mode.cptr()))
}

// Window (https://wiki.libsdl.org/SDL_GetWindowPixelFormat)
func (window *Window) GetPixelFormat() uint32 {
	return (uint32)(C.SDL_GetWindowPixelFormat(window.cptr()))
}

// CreateWindow (https://wiki.libsdl.org/SDL_CreateWindow)
func CreateWindow(title string, x int, y int, w int, h int, flags uint32) (*Window, error) {
	_title := C.CString(title)
	var _window = C.SDL_CreateWindow(_title, C.int(x), C.int(y), C.int(w), C.int(h), C.Uint32(flags))
	if _window == nil {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer(_window)), nil
}

// CreateWindowFrom (https://wiki.libsdl.org/SDL_CreateWindowFrom)
func CreateWindowFrom(data unsafe.Pointer) (*Window, error) {
	_window := C.SDL_CreateWindowFrom(data)
	if _window == nil {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer(_window)), nil
}

// Window (https://wiki.libsdl.org/SDL_DestroyWindow)
func (window *Window) Destroy() {
	C.SDL_DestroyWindow(window.cptr())
}

// Window (https://wiki.libsdl.org/SDL_GetWindowID)
func (window *Window) GetID() uint32 {
	return (uint32)(C.SDL_GetWindowID(window.cptr()))
}

// GetWindowFromID (https://wiki.libsdl.org/SDL_GetWindowFromID)
func GetWindowFromID(id uint32) (*Window, error) {
	_window := C.SDL_GetWindowFromID(C.Uint32(id))
	if _window == nil {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer((_window))), nil
}

// Window (https://wiki.libsdl.org/SDL_GetWindowFlags)
func (window *Window) GetFlags() uint32 {
	return (uint32)(C.SDL_GetWindowFlags(window.cptr()))
}

// Window (https://wiki.libsdl.org/SDL_SetWindowTitle)
func (window *Window) SetTitle(title string) {
	_title := C.CString(title)
	C.SDL_SetWindowTitle(window.cptr(), _title)
}

// Window (https://wiki.libsdl.org/SDL_GetWindowTitle)
func (window *Window) GetTitle() string {
	return string(C.GoString(C.SDL_GetWindowTitle(window.cptr())))
}

// Window (https://wiki.libsdl.org/SDL_SetWindowIcon)
func (window *Window) SetIcon(icon *Surface) {
	C.SDL_SetWindowIcon(window.cptr(), icon.cptr())
}

// Window (https://wiki.libsdl.org/SDL_SetWindowData)
func (window *Window) SetData(name string, userdata unsafe.Pointer) unsafe.Pointer {
	_name := C.CString(name)
	return unsafe.Pointer(C.SDL_SetWindowData(window.cptr(), _name, userdata))
}

// Window (https://wiki.libsdl.org/SDL_GetWindowData)
func (window *Window) GetData(name string) unsafe.Pointer {
	_name := C.CString(name)
	return unsafe.Pointer(C.SDL_GetWindowData(window.cptr(), _name))
}

// Window (https://wiki.libsdl.org/SDL_SetWindowPosition)
func (window *Window) SetPosition(x int, y int) {
	C.SDL_SetWindowPosition(window.cptr(), C.int(x), C.int(y))
}

// Window (https://wiki.libsdl.org/SDL_GetWindowPosition)
func (window *Window) GetPosition() (x, y int) {
	var _x, _y C.int
	C.SDL_GetWindowPosition(window.cptr(), &_x, &_y)
	return int(_x), int(_y)
}

// Window (https://wiki.libsdl.org/SDL_SetWindowSize)
func (window *Window) SetSize(w int, h int) {
	C.SDL_SetWindowSize(window.cptr(), C.int(w), C.int(h))
}

// Window (https://wiki.libsdl.org/SDL_GetWindowSize)
func (window *Window) GetSize() (w, h int) {
	var _w, _h C.int
	C.SDL_GetWindowSize(window.cptr(), &_w, &_h)
	return int(_w), int(_h)
}

// Window (https://wiki.libsdl.org/SDL_SetWindowMinimumSize)
func (window *Window) SetMinimumSize(minW int, minH int) {
	C.SDL_SetWindowMinimumSize(window.cptr(), C.int(minW), C.int(minH))
}

// Window (https://wiki.libsdl.org/SDL_GetWindowMinimumSize)
func (window *Window) GetMinimumSize() (w, h int) {
	var _w, _h C.int
	C.SDL_GetWindowMinimumSize(window.cptr(), &_w, &_h)
	return int(_w), int(_h)
}

// Window (https://wiki.libsdl.org/SDL_SetWindowMaximumSize)
func (window *Window) SetMaximumSize(maxW int, maxH int) {
	C.SDL_SetWindowMaximumSize(window.cptr(), C.int(maxW), C.int(maxH))
}

// Window (https://wiki.libsdl.org/SDL_GetWindowMaximumSize)
func (window *Window) GetMaximumSize() (w, h int) {
	var _w, _h C.int
	C.SDL_GetWindowMaximumSize(window.cptr(), &_w, &_h)
	return int(_w), int(_h)
}

// Window (https://wiki.libsdl.org/SDL_SetWindowBordered)
func (window *Window) SetBordered(bordered bool) {
	C.SDL_SetWindowBordered(window.cptr(), (C.SDL_bool)(Btoi(bordered)))
}

// Window (https://wiki.libsdl.org/SDL_ShowWindow)
func (window *Window) Show() {
	C.SDL_ShowWindow(window.cptr())
}

// Window (https://wiki.libsdl.org/SDL_HideWindow)
func (window *Window) Hide() {
	C.SDL_HideWindow(window.cptr())
}

// Window (https://wiki.libsdl.org/SDL_RaiseWindow)
func (window *Window) Raise() {
	C.SDL_RaiseWindow(window.cptr())
}

// Window (https://wiki.libsdl.org/SDL_MaximizeWindow)
func (window *Window) Maximize() {
	C.SDL_MaximizeWindow(window.cptr())
}

// Window (https://wiki.libsdl.org/SDL_MinimizeWindow)
func (window *Window) Minimize() {
	C.SDL_MinimizeWindow(window.cptr())
}

// Window (https://wiki.libsdl.org/SDL_RestoreWindow)
func (window *Window) Restore() {
	C.SDL_RestoreWindow(window.cptr())
}

// Window (https://wiki.libsdl.org/SDL_SetWindowFullscreen)
func (window *Window) SetFullscreen(flags uint32) int {
	return int(C.SDL_SetWindowFullscreen(window.cptr(), C.Uint32(flags)))
}

// Window (https://wiki.libsdl.org/SDL_GetWindowSurface)
func (window *Window) GetSurface() *Surface {
	return (*Surface)(unsafe.Pointer(C.SDL_GetWindowSurface(window.cptr())))
}

// Window (https://wiki.libsdl.org/SDL_UpdateWindowSurface)
func (window *Window) UpdateSurface() error {
	if C.SDL_UpdateWindowSurface(window.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Window (https://wiki.libsdl.org/SDL_UpdateWindowSurfaceRects)
func (window *Window) UpdateSurfaceRects(rects []Rect) int {
	return int(C.SDL_UpdateWindowSurfaceRects(window.cptr(), rects[0].cptr(), C.int(len(rects))))
}

// Window (https://wiki.libsdl.org/SDL_SetWindowGrab)
func (window *Window) SetGrab(grabbed bool) {
	C.SDL_SetWindowGrab(window.cptr(), C.SDL_bool((Btoi(grabbed))))
}

// Window (https://wiki.libsdl.org/SDL_GetWindowGrab)
func (window *Window) GetGrab() bool {
	return C.SDL_GetWindowGrab(window.cptr()) != 0
}

// Window (https://wiki.libsdl.org/SDL_SetWindowBrightness)
func (window *Window) SetBrightness(brightness float32) int {
	return int(C.SDL_SetWindowBrightness(window.cptr(), C.float(brightness)))
}

// Window (https://wiki.libsdl.org/SDL_GetWindowBrightness)
func (window *Window) GetBrightness() float32 {
	return float32(C.SDL_GetWindowBrightness(window.cptr()))
}

// Window (https://wiki.libsdl.org/SDL_SetWindowGammaRamp)
func (window *Window) SetGammaRamp(red, green, blue *uint16) int {
	_red := (*C.Uint16)(red)
	_green := (*C.Uint16)(red)
	_blue := (*C.Uint16)(blue)
	return int(C.SDL_SetWindowGammaRamp(window.cptr(), _red, _green, _blue))
}

// Window (https://wiki.libsdl.org/SDL_GetWindowGammaRamp)
func (window *Window) GetGammaRamp() (red, green, blue uint16, status int) {
	var _red, _green, _blue C.Uint16
	_status := int(C.SDL_GetWindowGammaRamp(window.cptr(), &_red, &_green, &_blue))
	return uint16(_red), uint16(_green), uint16(_blue), _status
}

// IsScreenSaverEnabled (https://wiki.libsdl.org/SDL_IsScreenSaverEnabled)
func IsScreenSaverEnabled() bool {
	return C.SDL_IsScreenSaverEnabled() != 0
}

// EnableScreenSaver (https://wiki.libsdl.org/SDL_EnableScreenSaver)
func EnableScreenSaver() {
	C.SDL_EnableScreenSaver()
}

// DisableScreenSaver (https://wiki.libsdl.org/SDL_DisableScreenSaver)
func DisableScreenSaver() {
	C.SDL_DisableScreenSaver()
}

// GL_LoadLibrary (https://wiki.libsdl.org/SDL_GL_LoadLibrary)
func GL_LoadLibrary(path string) int {
	_path := C.CString(path)
	return int(C.SDL_GL_LoadLibrary(_path))
}

// GL_GetProcAddress (https://wiki.libsdl.org/SDL_GL_GetProcAddress)
func GL_GetProcAddress(proc string) unsafe.Pointer {
	_proc := C.CString(proc)
	return unsafe.Pointer(C.SDL_GL_GetProcAddress(_proc))
}

// GL_UnloadLibrary (https://wiki.libsdl.org/SDL_GL_UnloadLibrary)
func GL_UnloadLibrary() {
	C.SDL_GL_UnloadLibrary()
}

// GL_ExtensionSupported (https://wiki.libsdl.org/SDL_GL_ExtensionSupported)
func GL_ExtensionSupported(extension string) bool {
	_extension := C.CString(extension)
	return C.SDL_GL_ExtensionSupported(_extension) != 0
}

// GL_SetAttribute (https://wiki.libsdl.org/SDL_GL_SetAttribute)
func GL_SetAttribute(attr GLattr, value int) int {
	return int(C.SDL_GL_SetAttribute(attr.c(), C.int(value)))
}

// GL_GetAttribute (https://wiki.libsdl.org/SDL_GL_GetAttribute)
func GL_GetAttribute(attr GLattr) (value int, status int) {
	var _value, _status C.int
	_status = (C.SDL_GL_GetAttribute(attr.c(), &_value))
	return int(_value), int(_status)
}

// GL_CreateContext (https://wiki.libsdl.org/SDL_GL_CreateContext)
func GL_CreateContext(window *Window) GLContext {
	return GLContext(C.SDL_GL_CreateContext(window.cptr()))
}

// GL_MakeCurrent (https://wiki.libsdl.org/SDL_GL_MakeCurrent)
func GL_MakeCurrent(window *Window, glcontext GLContext) int {
	return int(C.SDL_GL_MakeCurrent(window.cptr(), C.SDL_GLContext(glcontext)))
}

// GL_SetSwapInterval (https://wiki.libsdl.org/SDL_GL_SetSwapInterval)
func GL_SetSwapInterval(interval int) int {
	return int(C.SDL_GL_SetSwapInterval(C.int(interval)))
}

// GL_GetSwapInterval (https://wiki.libsdl.org/SDL_GL_GetSwapInterval)
func GL_GetSwapInterval() int {
	return int(C.SDL_GL_GetSwapInterval())
}

// GL_SwapWindow (https://wiki.libsdl.org/SDL_GL_SwapWindow)
func GL_SwapWindow(window *Window) {
	C.SDL_GL_SwapWindow(window.cptr())
}

// GL_DeleteContext (https://wiki.libsdl.org/SDL_GL_DeleteContext)
func GL_DeleteContext(context GLContext) {
	C.SDL_GL_DeleteContext(C.SDL_GLContext(context))
}
