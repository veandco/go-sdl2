package sdl

/*
#include "sdl_wrapper.h"
static inline Sint32 ShowMessageBox(SDL_MessageBoxData data)
{
	Sint32 buttonid;
	SDL_ShowMessageBox(&data, &buttonid);
	return buttonid;
}

#if !(SDL_VERSION_ATLEAST(2,0,1))
static void SDL_GL_GetDrawableSize(SDL_Window *window, int *w, int *h)
{
	*w = 0;
	*h = 0;
}
#endif

#if !(SDL_VERSION_ATLEAST(2,0,5))
#pragma message("SDL_SetWindowResizable is not supported before SDL 2.0.5")
static void SDL_SetWindowResizable(SDL_Window *window, SDL_bool resizable)
{
}
#endif
*/
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
	WINDOW_ALLOW_HIGHDPI      = C.SDL_WINDOW_ALLOW_HIGHDPI
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
	MESSAGEBOX_ERROR       = C.SDL_MESSAGEBOX_ERROR
	MESSAGEBOX_WARNING     = C.SDL_MESSAGEBOX_WARNING
	MESSAGEBOX_INFORMATION = C.SDL_MESSAGEBOX_INFORMATION
)

const (
	MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT = C.SDL_MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT
	MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT = C.SDL_MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT
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

// MessageBoxColor (https://wiki.libsdl.org/SDL_MessageBoxColor)
type MessageBoxColor struct {
	R uint8
	G uint8
	B uint8
}
type cMessageBoxColor C.SDL_MessageBoxColor

// MessageBoxColorScheme (https://wiki.libsdl.org/SDL_MessageBoxColorScheme)
type MessageBoxColorScheme struct {
	Colors [5]MessageBoxColor
}
type cMessageBoxColorScheme C.SDL_MessageBoxColorScheme

// MessageBoxButtonData (https://wiki.libsdl.org/SDL_MessageBoxButtonData)
type MessageBoxButtonData struct {
	Flags    uint32
	ButtonId int32
	Text     string
}

// MessageBoxData (https://wiki.libsdl.org/SDL_MessageBoxData)
type MessageBoxData struct {
	Flags       uint32
	Window      *Window
	Title       string
	Message     string
	NumButtons  int32
	Buttons     []MessageBoxButtonData
	ColorScheme *MessageBoxColorScheme
}

type cMessageBoxData struct {
	Flags       uint32
	Window      *C.SDL_Window
	Title       *C.char
	Message     *C.char
	NumButtons  int32
	Buttons     *C.SDL_MessageBoxButtonData
	ColorScheme *C.SDL_MessageBoxColorScheme
}

func (w *Window) cptr() *C.SDL_Window {
	return (*C.SDL_Window)(unsafe.Pointer(w))
}

func (dm *DisplayMode) cptr() *C.SDL_DisplayMode {
	return (*C.SDL_DisplayMode)(unsafe.Pointer(dm))
}

func (mc *MessageBoxColor) cptr() *C.SDL_MessageBoxColor {
	return (*C.SDL_MessageBoxColor)(unsafe.Pointer(mc))
}

func (mcs *MessageBoxColorScheme) cptr() *C.SDL_MessageBoxColorScheme {
	return (*C.SDL_MessageBoxColorScheme)(unsafe.Pointer(mcs))
}

func (mbd *MessageBoxButtonData) cptr() *C.SDL_MessageBoxButtonData {
	return (*C.SDL_MessageBoxButtonData)(unsafe.Pointer(mbd))
}

func (md *MessageBoxData) cptr() *C.SDL_MessageBoxData {
	return (*C.SDL_MessageBoxData)(unsafe.Pointer(md))
}

func (attr GLattr) c() C.SDL_GLattr {
	return C.SDL_GLattr(attr)
}

// GetDisplayName (https://wiki.libsdl.org/SDL_GetDisplayName)
func GetDisplayName(displayIndex int) string {
	return C.GoString(C.SDL_GetDisplayName(C.int(displayIndex)))
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
func GetNumVideoDrivers() (int, error) {
	n := int(C.SDL_GetNumVideoDrivers())
	if n < 0 {
		return n, GetError()
	}
	return n, nil
}

// GetVideoDriver (https://wiki.libsdl.org/SDL_GetVideoDriver)
func GetVideoDriver(index int) string {
	return string(C.GoString(C.SDL_GetVideoDriver(C.int(index))))
}

// VideoInit (https://wiki.libsdl.org/SDL_VideoInit)
func VideoInit(driverName string) error {
	if C.SDL_VideoInit(C.CString(driverName)) != 0 {
		return GetError()
	}
	return nil
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
func GetNumDisplayModes(displayIndex int) (int, error) {
	n := int(C.SDL_GetNumDisplayModes(C.int(displayIndex)))
	if n < 0 {
		return n, GetError()
	}
	return n, nil
}

// GetDisplayBounds (https://wiki.libsdl.org/SDL_GetDisplayBounds)
func GetDisplayBounds(displayIndex int, rect *Rect) error {
	if C.SDL_GetDisplayBounds(C.int(displayIndex), rect.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// GetDisplayMode (https://wiki.libsdl.org/SDL_GetDisplayMode)
func GetDisplayMode(displayIndex int, modeIndex int, mode *DisplayMode) error {
	if C.SDL_GetDisplayMode(C.int(displayIndex), C.int(modeIndex), mode.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// GetDesktopDisplayMode (https://wiki.libsdl.org/SDL_GetDesktopDisplayMode)
func GetDesktopDisplayMode(displayIndex int, mode *DisplayMode) error {
	if C.SDL_GetDesktopDisplayMode(C.int(displayIndex), mode.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// GetCurrentDisplayMode (https://wiki.libsdl.org/SDL_GetCurrentDisplayMode)
func GetCurrentDisplayMode(displayIndex int, mode *DisplayMode) error {
	if C.SDL_GetCurrentDisplayMode(C.int(displayIndex), mode.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// GetClosestDisplayMode (https://wiki.libsdl.org/SDL_GetClosestDisplayMode)
func GetClosestDisplayMode(displayIndex int, mode *DisplayMode, closest *DisplayMode) (*DisplayMode, error) {
	m := (*DisplayMode)(unsafe.Pointer((C.SDL_GetClosestDisplayMode(C.int(displayIndex), mode.cptr(), closest.cptr()))))
	if m == nil {
		return nil, GetError()
	}
	return m, nil
}

// Window (https://wiki.libsdl.org/SDL_GetWindowDisplayIndex)
func (window *Window) GetDisplayIndex() (int, error) {
	i := int(C.SDL_GetWindowDisplayIndex(window.cptr()))
	if i < 0 {
		return i, GetError()
	}
	return i, nil
}

// Window (https://wiki.libsdl.org/SDL_SetWindowDisplayMode)
func (window *Window) SetDisplayMode(mode *DisplayMode) error {
	if C.SDL_SetWindowDisplayMode(window.cptr(), mode.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Window (https://wiki.libsdl.org/SDL_GetWindowDisplayMode)
func (window *Window) GetDisplayMode(mode *DisplayMode) error {
	if C.SDL_GetWindowDisplayMode(window.cptr(), mode.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Window (https://wiki.libsdl.org/SDL_GetWindowPixelFormat)
func (window *Window) GetPixelFormat() (uint32, error) {
	f := (uint32)(C.SDL_GetWindowPixelFormat(window.cptr()))
	if f == PIXELFORMAT_UNKNOWN {
		return f, GetError()
	}
	return f, nil
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

// Window (https://wiki.libsdl.org/SDL_SetWindowResizable)
func (window *Window) SetResizable(resizable bool) {
    var _resizable C.SDL_bool = C.SDL_FALSE
    if resizable {
        _resizable = C.SDL_TRUE
    }
    C.SDL_SetWindowResizable(window.cptr(), _resizable)
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
func (window *Window) SetFullscreen(flags uint32) error {
	if C.SDL_SetWindowFullscreen(window.cptr(), C.Uint32(flags)) != 0 {
		return GetError()
	}
	return nil
}

// Window (https://wiki.libsdl.org/SDL_GetWindowSurface)
func (window *Window) GetSurface() (*Surface, error) {
	surface := (*Surface)(unsafe.Pointer(C.SDL_GetWindowSurface(window.cptr())))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// Window (https://wiki.libsdl.org/SDL_UpdateWindowSurface)
func (window *Window) UpdateSurface() error {
	if C.SDL_UpdateWindowSurface(window.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// Window (https://wiki.libsdl.org/SDL_UpdateWindowSurfaceRects)
func (window *Window) UpdateSurfaceRects(rects []Rect) error {
	if C.SDL_UpdateWindowSurfaceRects(window.cptr(), rects[0].cptr(), C.int(len(rects))) != 0 {
		return GetError()
	}
	return nil
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
func (window *Window) SetBrightness(brightness float32) error {
	if C.SDL_SetWindowBrightness(window.cptr(), C.float(brightness)) != 0 {
		return GetError()
	}
	return nil
}

// Window (https://wiki.libsdl.org/SDL_GetWindowBrightness)
func (window *Window) GetBrightness() float32 {
	return float32(C.SDL_GetWindowBrightness(window.cptr()))
}

// Window (https://wiki.libsdl.org/SDL_SetWindowGammaRamp)
func (window *Window) SetGammaRamp(red, green, blue *[256]uint16) error {
	if C.SDL_SetWindowGammaRamp(
		window.cptr(),
		(*C.Uint16)(unsafe.Pointer(red)),
		(*C.Uint16)(unsafe.Pointer(green)),
		(*C.Uint16)(unsafe.Pointer(blue))) != 0 {

		return GetError()
	}
	return nil
}

// Window (https://wiki.libsdl.org/SDL_GetWindowGammaRamp)
func (window *Window) GetGammaRamp() (red, green, blue *[256]uint16, err error) {
	if C.SDL_GetWindowGammaRamp(
		window.cptr(),
		(*C.Uint16)(unsafe.Pointer(red)),
		(*C.Uint16)(unsafe.Pointer(green)),
		(*C.Uint16)(unsafe.Pointer(blue))) != 0 {

		return red, green, blue, GetError()
	}
	return red, green, blue, nil
}

// ShowSimpleMessageBox (https://wiki.libsdl.org/SDL_ShowSimpleMessageBox)
func ShowSimpleMessageBox(flags uint32, title, message string, window *Window) error {
	_title := C.CString(title)
	defer C.free(unsafe.Pointer(_title))
	_message := C.CString(message)
	defer C.free(unsafe.Pointer(_message))
	if (int)(C.SDL_ShowSimpleMessageBox(C.Uint32(flags), _title, _message, window.cptr())) < 0 {
		return GetError()
	}
	return nil
}

// ShowMessageBox (https://wiki.libsdl.org/SDL_ShowMessageBox)
func ShowMessageBox(data *MessageBoxData) (err error, buttonid int32) {
	_title := C.CString(data.Title)
	defer C.free(unsafe.Pointer(_title))
	_message := C.CString(data.Message)
	defer C.free(unsafe.Pointer(_message))

	var cbuttons []C.SDL_MessageBoxButtonData
	var cbtntexts []*C.char
	defer func(texts []*C.char) {
		for _, t := range texts {
			C.free(unsafe.Pointer(t))
		}
	}(cbtntexts)

	for _, btn := range data.Buttons {
		ctext := C.CString(btn.Text)
		cbtn := C.SDL_MessageBoxButtonData{
			flags:    C.Uint32(btn.Flags),
			buttonid: C.int(btn.ButtonId),
			text:     ctext,
		}

		cbuttons = append(cbuttons, cbtn)
		cbtntexts = append(cbtntexts, ctext)
	}

	cdata := C.SDL_MessageBoxData{
		flags:       C.Uint32(data.Flags),
		window:      data.Window.cptr(),
		title:       _title,
		message:     _message,
		numbuttons:  C.int(data.NumButtons),
		buttons:     &cbuttons[0],
		colorScheme: data.ColorScheme.cptr(),
	}

	if buttonid := int32(C.ShowMessageBox(cdata)); buttonid < 0 {
		return GetError(), buttonid
	} else {
		return nil, buttonid
	}
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
func GL_LoadLibrary(path string) error {
	_path := C.CString(path)
	if C.SDL_GL_LoadLibrary(_path) != 0 {
		return GetError()
	}
	return nil
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
func GL_SetAttribute(attr GLattr, value int) error {
	if C.SDL_GL_SetAttribute(attr.c(), C.int(value)) != 0 {
		return GetError()
	}
	return nil
}

// GL_GetAttribute (https://wiki.libsdl.org/SDL_GL_GetAttribute)
func GL_GetAttribute(attr GLattr) (int, error) {
	var _value C.int
	if C.SDL_GL_GetAttribute(attr.c(), &_value) != 0 {
		return int(_value), GetError()
	}
	return int(_value), nil
}

// GL_CreateContext (https://wiki.libsdl.org/SDL_GL_CreateContext)
func GL_CreateContext(window *Window) (GLContext, error) {
	c := GLContext(C.SDL_GL_CreateContext(window.cptr()))
	if c == nil {
		return nil, GetError()
	}
	return c, nil
}

// GL_MakeCurrent (https://wiki.libsdl.org/SDL_GL_MakeCurrent)
func GL_MakeCurrent(window *Window, glcontext GLContext) error {
	if C.SDL_GL_MakeCurrent(window.cptr(), C.SDL_GLContext(glcontext)) != 0 {
		return GetError()
	}
	return nil
}

// GL_SetSwapInterval (https://wiki.libsdl.org/SDL_GL_SetSwapInterval)
func GL_SetSwapInterval(interval int) error {
	if C.SDL_GL_SetSwapInterval(C.int(interval)) != 0 {
		return GetError()
	}
	return nil
}

// GL_GetSwapInterval (https://wiki.libsdl.org/SDL_GL_GetSwapInterval)
func GL_GetSwapInterval() (int, error) {
	i := int(C.SDL_GL_GetSwapInterval())
	if i == -1 {
		return i, GetError()
	}
	return i, nil
}

// GL_GetDrawableSize (https://wiki.libsdl.org/SDL_GL_GetDrawableSize)
func GL_GetDrawableSize(window *Window) (w, h int) {
	var _w, _h C.int
	C.SDL_GL_GetDrawableSize(window.cptr(), &_w, &_h)
	return int(_w), int(_h)
}

// GL_SwapWindow (https://wiki.libsdl.org/SDL_GL_SwapWindow)
func GL_SwapWindow(window *Window) {
	C.SDL_GL_SwapWindow(window.cptr())
}

// GL_DeleteContext (https://wiki.libsdl.org/SDL_GL_DeleteContext)
func GL_DeleteContext(context GLContext) {
	C.SDL_GL_DeleteContext(C.SDL_GLContext(context))
}
