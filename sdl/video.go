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


#if defined(WARN_OUTDATED)
#pragma message("SDL_WINDOW_ALLOW_HIGHDPI is not supported before SDL 2.0.1")
#endif

#define SDL_WINDOW_ALLOW_HIGHDPI (0)


#if defined(WARN_OUTDATED)
#pragma message("SDL_GL_FRAMEBUFFER_SRGB_CAPABLE is not supported before SDL 2.0.1")
#endif

#define SDL_GL_FRAMEBUFFER_SRGB_CAPABLE (0)
#endif

#if !(SDL_VERSION_ATLEAST(2,0,4))

#if defined(WARN_OUTDATED)
#pragma message("SDL_WINDOW_MOUSE_CAPTURE is not supported before SDL 2.0.4")
#endif

#define SDL_WINDOW_MOUSE_CAPTURE (0)


#if defined(WARN_OUTDATED)
#pragma message("SDL_GL_CONTEXT_RELEASE_BEHAVIOR is not supported before SDL 2.0.4")
#endif

#define SDL_GL_CONTEXT_RELEASE_BEHAVIOR (0)


#if defined(WARN_OUTDATED)
#pragma message("SDL_GetDisplayDPI is not supported before SDL 2.0.4")
#endif

static int SDL_GetDisplayDPI(int displayIndex, float* ddpi, float* hdpi, float* vdpi)
{
	return -1;
}

#endif

#if !(SDL_VERSION_ATLEAST(2,0,5))

#if defined(WARN_OUTDATED)
#pragma message("SDL_SetWindowResizable is not supported before SDL 2.0.5")
#endif

static void SDL_SetWindowResizable(SDL_Window *window, SDL_bool resizable)
{
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_SetWindowOpacity is not supported before SDL 2.0.5")
#endif

static int SDL_SetWindowOpacity(SDL_Window *window, float opacity)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_GetWindowOpacity is not supported before SDL 2.0.5")
#endif

static int SDL_GetWindowOpacity(SDL_Window *window, float *opacity)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_GetDisplayUsableBounds is not supported before SDL 2.0.5")
#endif

static int SDL_GetDisplayUsableBounds(int displayIndex, SDL_Rect* rect)
{
	return -1;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_WINDOW_ALWAYS_ON_TOP is not supported before SDL 2.0.5")
#endif

#define SDL_WINDOW_ALWAYS_ON_TOP (0)


#if defined(WARN_OUTDATED)
#pragma message("SDL_WINDOW_SKIP_TASKBAR is not supported before SDL 2.0.5")
#endif

#define SDL_WINDOW_SKIP_TASKBAR (0)


#if defined(WARN_OUTDATED)
#pragma message("SDL_WINDOW_UTILITY is not supported before SDL 2.0.5")
#endif

#define SDL_WINDOW_UTILITY (0)


#if defined(WARN_OUTDATED)
#pragma message("SDL_WINDOW_TOOLTIP is not supported before SDL 2.0.5")
#endif

#define SDL_WINDOW_TOOLTIP (0)


#if defined(WARN_OUTDATED)
#pragma message("SDL_WINDOW_POPUP_MENU is not supported before SDL 2.0.5")
#endif

#define SDL_WINDOW_POPUP_MENU (0)


#if defined(WARN_OUTDATED)
#pragma message("SDL_WINDOWEVENT_TAKE_FOCUS is not supported before SDL 2.0.5")
#endif

#define SDL_WINDOWEVENT_TAKE_FOCUS (0)


#if defined(WARN_OUTDATED)
#pragma message("SDL_WINDOWEVENT_HIT_TEST is not supported before SDL 2.0.5")
#endif

#define SDL_WINDOWEVENT_HIT_TEST (0)
#endif

#if !(SDL_VERSION_ATLEAST(2,0,6))

#if defined(WARN_OUTDATED)
#pragma message("SDL_WINDOW_VULKAN is not supported before SDL 2.0.6")
#endif

#define SDL_WINDOW_VULKAN (0)


#if defined(WARN_OUTDATED)
#pragma message("SDL_GL_CONTEXT_RESET_NOTIFICATION is not supported before SDL 2.0.6")
#endif

#define SDL_GL_CONTEXT_RESET_NOTIFICATION (0)


#if defined(WARN_OUTDATED)
#pragma message("SDL_GL_CONTEXT_NO_ERROR is not supported before SDL 2.0.6")
#endif

#define SDL_GL_CONTEXT_NO_ERROR (0)
#endif


#if !(SDL_VERSION_ATLEAST(2,0,16))

#if defined(WARN_OUTDATED)
#pragma message("SDL_FlashWindow is not supported before SDL 2.0.16")
#pragma message("SDL_SetWindowAlwaysOnTop is not supported before SDL 2.0.16")
#pragma message("SDL_SetWindowKeyboardGrab is not supported before SDL 2.0.16")
#endif

typedef enum
{
    SDL_FLASH_CANCEL,                   // Cancel any window flash state
    SDL_FLASH_BRIEFLY,                  // Flash the window briefly to get attention
    SDL_FLASH_UNTIL_FOCUSED,            // Flash the window until it gets focus
} SDL_FlashOperation;

static int SDL_FlashWindow(SDL_Window * window, SDL_FlashOperation operation)
{
	return -1;
}

static void SDL_SetWindowAlwaysOnTop(SDL_Window * window, SDL_bool on_top)
{
	return;
}

static void SDL_SetWindowKeyboardGrab(SDL_Window * window, SDL_bool grabbed)
{
	return;
}

#endif

#if !(SDL_VERSION_ATLEAST(2,0,18))

#if defined(WARN_OUTDATED)
#pragma message("SDL_GetWindowICCProfile is not supported before SDL 2.0.18")
#pragma message("SDL_SetWindowMouseRect is not supported before SDL 2.0.18")
#pragma message("SDL_GetWindowMouseRect is not supported before SDL 2.0.18")
#endif

#define SDL_WINDOWEVENT_ICCPROF_CHANGED (17) // The ICC profile of the window's display has changed.
#define SDL_WINDOWEVENT_DISPLAY_CHANGED (18) // Window has been moved to display data1.

static void* SDLCALL SDL_GetWindowICCProfile(SDL_Window * window, size_t* size)
{
	return NULL;
}

static int SDL_SetWindowMouseRect(SDL_Window * window, const SDL_Rect * rect)
{
	return -1;
}

static const SDL_Rect * SDLCALL SDL_GetWindowMouseRect(SDL_Window * window)
{
	return NULL;
}

#endif
*/
import "C"
import "unsafe"

// An enumeration of window states.
// (https://wiki.libsdl.org/SDL_WindowFlags)
const (
	WINDOW_FULLSCREEN         = C.SDL_WINDOW_FULLSCREEN         // fullscreen window
	WINDOW_OPENGL             = C.SDL_WINDOW_OPENGL             // window usable with OpenGL context
	WINDOW_SHOWN              = C.SDL_WINDOW_SHOWN              // window is visible
	WINDOW_HIDDEN             = C.SDL_WINDOW_HIDDEN             // window is not visible
	WINDOW_BORDERLESS         = C.SDL_WINDOW_BORDERLESS         // no window decoration
	WINDOW_RESIZABLE          = C.SDL_WINDOW_RESIZABLE          // window can be resized
	WINDOW_MINIMIZED          = C.SDL_WINDOW_MINIMIZED          // window is minimized
	WINDOW_MAXIMIZED          = C.SDL_WINDOW_MAXIMIZED          // window is maximized
	WINDOW_INPUT_GRABBED      = C.SDL_WINDOW_INPUT_GRABBED      // window has grabbed input focus
	WINDOW_INPUT_FOCUS        = C.SDL_WINDOW_INPUT_FOCUS        // window has input focus
	WINDOW_MOUSE_FOCUS        = C.SDL_WINDOW_MOUSE_FOCUS        // window has mouse focus
	WINDOW_FULLSCREEN_DESKTOP = C.SDL_WINDOW_FULLSCREEN_DESKTOP // fullscreen window at the current desktop resolution
	WINDOW_FOREIGN            = C.SDL_WINDOW_FOREIGN            // window not created by SDL
	WINDOW_ALLOW_HIGHDPI      = C.SDL_WINDOW_ALLOW_HIGHDPI      // window should be created in high-DPI mode if supported (>= SDL 2.0.1)
	WINDOW_MOUSE_CAPTURE      = C.SDL_WINDOW_MOUSE_CAPTURE      // window has mouse captured (unrelated to INPUT_GRABBED, >= SDL 2.0.4)
	WINDOW_ALWAYS_ON_TOP      = C.SDL_WINDOW_ALWAYS_ON_TOP      // window should always be above others (X11 only, >= SDL 2.0.5)
	WINDOW_SKIP_TASKBAR       = C.SDL_WINDOW_SKIP_TASKBAR       // window should not be added to the taskbar (X11 only, >= SDL 2.0.5)
	WINDOW_UTILITY            = C.SDL_WINDOW_UTILITY            // window should be treated as a utility window (X11 only, >= SDL 2.0.5)
	WINDOW_TOOLTIP            = C.SDL_WINDOW_TOOLTIP            // window should be treated as a tooltip (X11 only, >= SDL 2.0.5)
	WINDOW_POPUP_MENU         = C.SDL_WINDOW_POPUP_MENU         // window should be treated as a popup menu (X11 only, >= SDL 2.0.5)
	WINDOW_VULKAN             = C.SDL_WINDOW_VULKAN             // window usable for Vulkan surface (>= SDL 2.0.6)
)

// An enumeration of window events.
// (https://wiki.libsdl.org/SDL_WindowEventID)
const (
	WINDOWEVENT_NONE            = C.SDL_WINDOWEVENT_NONE            // (never used)
	WINDOWEVENT_SHOWN           = C.SDL_WINDOWEVENT_SHOWN           // window has been shown
	WINDOWEVENT_HIDDEN          = C.SDL_WINDOWEVENT_HIDDEN          // window has been hidden
	WINDOWEVENT_EXPOSED         = C.SDL_WINDOWEVENT_EXPOSED         // window has been exposed and should be redrawn
	WINDOWEVENT_MOVED           = C.SDL_WINDOWEVENT_MOVED           // window has been moved to data1, data2
	WINDOWEVENT_RESIZED         = C.SDL_WINDOWEVENT_RESIZED         // window has been resized to data1xdata2; this event is always preceded by WINDOWEVENT_SIZE_CHANGED
	WINDOWEVENT_SIZE_CHANGED    = C.SDL_WINDOWEVENT_SIZE_CHANGED    // window size has changed, either as a result of an API call or through the system or user changing the window size; this event is followed by WINDOWEVENT_RESIZED if the size was changed by an external event, i.e. the user or the window manager
	WINDOWEVENT_MINIMIZED       = C.SDL_WINDOWEVENT_MINIMIZED       // window has been minimized
	WINDOWEVENT_MAXIMIZED       = C.SDL_WINDOWEVENT_MAXIMIZED       // window has been maximized
	WINDOWEVENT_RESTORED        = C.SDL_WINDOWEVENT_RESTORED        // window has been restored to normal size and position
	WINDOWEVENT_ENTER           = C.SDL_WINDOWEVENT_ENTER           // window has gained mouse focus
	WINDOWEVENT_LEAVE           = C.SDL_WINDOWEVENT_LEAVE           // window has lost mouse focus
	WINDOWEVENT_FOCUS_GAINED    = C.SDL_WINDOWEVENT_FOCUS_GAINED    // window has gained keyboard focus
	WINDOWEVENT_FOCUS_LOST      = C.SDL_WINDOWEVENT_FOCUS_LOST      // window has lost keyboard focus
	WINDOWEVENT_CLOSE           = C.SDL_WINDOWEVENT_CLOSE           // the window manager requests that the window be closed
	WINDOWEVENT_TAKE_FOCUS      = C.SDL_WINDOWEVENT_TAKE_FOCUS      // window is being offered a focus (should SDL_SetWindowInputFocus() on itself or a subwindow, or ignore) (>= SDL 2.0.5)
	WINDOWEVENT_HIT_TEST        = C.SDL_WINDOWEVENT_HIT_TEST        // window had a hit test that wasn't SDL_HITTEST_NORMAL (>= SDL 2.0.5)
	WINDOWEVENT_ICCPROF_CHANGED = C.SDL_WINDOWEVENT_ICCPROF_CHANGED // the ICC profile of the window's display has changed
	WINDOWEVENT_DISPLAY_CHANGED = C.SDL_WINDOWEVENT_DISPLAY_CHANGED // window has been moved to display data1
)

// Window position flags.
// (https://wiki.libsdl.org/SDL_CreateWindow)
const (
	WINDOWPOS_UNDEFINED_MASK = C.SDL_WINDOWPOS_UNDEFINED_MASK // used to indicate that you don't care what the window position is
	WINDOWPOS_UNDEFINED      = C.SDL_WINDOWPOS_UNDEFINED      // used to indicate that you don't care what the window position is
	WINDOWPOS_CENTERED_MASK  = C.SDL_WINDOWPOS_CENTERED_MASK  // used to indicate that the window position should be centered
	WINDOWPOS_CENTERED       = C.SDL_WINDOWPOS_CENTERED       // used to indicate that the window position should be centered
)

// An enumeration of message box flags (e.g. if supported message box will display warning icon).
// (https://wiki.libsdl.org/SDL_MessageBoxFlags)
const (
	MESSAGEBOX_ERROR       = C.SDL_MESSAGEBOX_ERROR       // error dialog
	MESSAGEBOX_WARNING     = C.SDL_MESSAGEBOX_WARNING     // warning dialog
	MESSAGEBOX_INFORMATION = C.SDL_MESSAGEBOX_INFORMATION // informational dialog
)

// Flags for MessageBoxButtonData.
const (
	MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT = C.SDL_MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT // marks the default button when return is hit
	MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT = C.SDL_MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT // marks the default button when escape is hit
)

// OpenGL configuration attributes.
// (https://wiki.libsdl.org/SDL_GL_SetAttribute)
const (
	GL_RED_SIZE                   = C.SDL_GL_RED_SIZE                   // the minimum number of bits for the red channel of the color buffer; defaults to 3
	GL_GREEN_SIZE                 = C.SDL_GL_GREEN_SIZE                 // the minimum number of bits for the green channel of the color buffer; defaults to 3
	GL_BLUE_SIZE                  = C.SDL_GL_BLUE_SIZE                  // the minimum number of bits for the blue channel of the color buffer; defaults to 2
	GL_ALPHA_SIZE                 = C.SDL_GL_ALPHA_SIZE                 // the minimum number of bits for the alpha channel of the color buffer; defaults to 0
	GL_BUFFER_SIZE                = C.SDL_GL_BUFFER_SIZE                // the minimum number of bits for frame buffer size; defaults to 0
	GL_DOUBLEBUFFER               = C.SDL_GL_DOUBLEBUFFER               // whether the output is single or double buffered; defaults to double buffering on
	GL_DEPTH_SIZE                 = C.SDL_GL_DEPTH_SIZE                 // the minimum number of bits in the depth buffer; defaults to 16
	GL_STENCIL_SIZE               = C.SDL_GL_STENCIL_SIZE               // the minimum number of bits in the stencil buffer; defaults to 0
	GL_ACCUM_RED_SIZE             = C.SDL_GL_ACCUM_RED_SIZE             // the minimum number of bits for the red channel of the accumulation buffer; defaults to 0
	GL_ACCUM_GREEN_SIZE           = C.SDL_GL_ACCUM_GREEN_SIZE           // the minimum number of bits for the green channel of the accumulation buffer; defaults to 0
	GL_ACCUM_BLUE_SIZE            = C.SDL_GL_ACCUM_BLUE_SIZE            // the minimum number of bits for the blue channel of the accumulation buffer; defaults to 0
	GL_ACCUM_ALPHA_SIZE           = C.SDL_GL_ALPHA_SIZE                 // the minimum number of bits for the alpha channel of the accumulation buffer; defaults to 0
	GL_STEREO                     = C.SDL_GL_STEREO                     // whether the output is stereo 3D; defaults to off
	GL_MULTISAMPLEBUFFERS         = C.SDL_GL_MULTISAMPLEBUFFERS         // the number of buffers used for multisample anti-aliasing; defaults to 0; see Remarks for details
	GL_MULTISAMPLESAMPLES         = C.SDL_GL_MULTISAMPLESAMPLES         // the number of samples used around the current pixel used for multisample anti-aliasing; defaults to 0; see Remarks for details
	GL_ACCELERATED_VISUAL         = C.SDL_GL_ACCELERATED_VISUAL         // set to 1 to require hardware acceleration, set to 0 to force software rendering; defaults to allow either
	GL_RETAINED_BACKING           = C.SDL_GL_RETAINED_BACKING           // not used (deprecated)
	GL_CONTEXT_MAJOR_VERSION      = C.SDL_GL_CONTEXT_MAJOR_VERSION      // OpenGL context major version
	GL_CONTEXT_MINOR_VERSION      = C.SDL_GL_CONTEXT_MINOR_VERSION      // OpenGL context minor version
	GL_CONTEXT_EGL                = C.SDL_GL_CONTEXT_EGL                // not used (deprecated)
	GL_CONTEXT_FLAGS              = C.SDL_GL_CONTEXT_FLAGS              // some combination of 0 or more of elements of the GLcontextFlag enumeration; defaults to 0 (https://wiki.libsdl.org/SDL_GLcontextFlag)
	GL_CONTEXT_PROFILE_MASK       = C.SDL_GL_CONTEXT_PROFILE_MASK       // type of GL context (Core, Compatibility, ES); default value depends on platform (https://wiki.libsdl.org/SDL_GLprofile)
	GL_SHARE_WITH_CURRENT_CONTEXT = C.SDL_GL_SHARE_WITH_CURRENT_CONTEXT // OpenGL context sharing; defaults to 0
	GL_FRAMEBUFFER_SRGB_CAPABLE   = C.SDL_GL_FRAMEBUFFER_SRGB_CAPABLE   // requests sRGB capable visual; defaults to 0 (>= SDL 2.0.1)
	GL_CONTEXT_RELEASE_BEHAVIOR   = C.SDL_GL_CONTEXT_RELEASE_BEHAVIOR   // sets context the release behavior; defaults to 1 (>= SDL 2.0.4)
	GL_CONTEXT_RESET_NOTIFICATION = C.SDL_GL_CONTEXT_RESET_NOTIFICATION // (>= SDL 2.0.6)
	GL_CONTEXT_NO_ERROR           = C.SDL_GL_CONTEXT_NO_ERROR           // (>= SDL 2.0.6)
)

// An enumeration of OpenGL profiles.
// (https://wiki.libsdl.org/SDL_GLprofile)
const (
	GL_CONTEXT_PROFILE_CORE          = C.SDL_GL_CONTEXT_PROFILE_CORE          // OpenGL core profile - deprecated functions are disabled
	GL_CONTEXT_PROFILE_COMPATIBILITY = C.SDL_GL_CONTEXT_PROFILE_COMPATIBILITY // OpenGL compatibility profile - deprecated functions are allowed
	GL_CONTEXT_PROFILE_ES            = C.SDL_GL_CONTEXT_PROFILE_ES            // OpenGL ES profile - only a subset of the base OpenGL functionality is available
)

// An enumeration of OpenGL context configuration flags.
// (https://wiki.libsdl.org/SDL_GLcontextFlag)
const (
	GL_CONTEXT_DEBUG_FLAG              = C.SDL_GL_CONTEXT_DEBUG_FLAG              // intended to put the GL into a "debug" mode which might offer better developer insights, possibly at a loss of performance
	GL_CONTEXT_FORWARD_COMPATIBLE_FLAG = C.SDL_GL_CONTEXT_FORWARD_COMPATIBLE_FLAG // intended to put the GL into a "forward compatible" mode, which means that no deprecated functionality will be supported, possibly at a gain in performance, and only applies to GL 3.0 and later contexts
	GL_CONTEXT_ROBUST_ACCESS_FLAG      = C.SDL_GL_CONTEXT_ROBUST_ACCESS_FLAG      // intended to require a GL context that supports the GL_ARB_robustness extension--a mode that offers a few APIs that are safer than the usual defaults (think snprintf() vs sprintf())
	GL_CONTEXT_RESET_ISOLATION_FLAG    = C.SDL_GL_CONTEXT_RESET_ISOLATION_FLAG    // intended to require the GL to make promises about what to do in the face of driver or hardware failure
)

//
// Window flash operation
//
const (
	FLASH_CANCEL        FlashOperation = C.SDL_FLASH_CANCEL        // Cancel any window flash state
	FLASH_BRIEFLY                      = C.SDL_FLASH_BRIEFLY       // Flash the window briefly to get attention
	FLASH_UNTIL_FOCUSED                = C.SDL_FLASH_UNTIL_FOCUSED // Flash the window until it gets focus
)

type FlashOperation C.SDL_FlashOperation

// DisplayMode contains the description of a display mode.
// (https://wiki.libsdl.org/SDL_DisplayMode)
type DisplayMode struct {
	Format      uint32         // one of the PixelFormatEnum values (https://wiki.libsdl.org/SDL_PixelFormatEnum)
	W           int32          // width, in screen coordinates
	H           int32          // height, in screen coordinates
	RefreshRate int32          // refresh rate (in Hz), or 0 for unspecified
	DriverData  unsafe.Pointer // driver-specific data, initialize to 0
}
type cDisplayMode C.SDL_DisplayMode

// Window is a type used to identify a window.
type Window C.SDL_Window

// GLContext is an opaque handle to an OpenGL context.
type GLContext C.SDL_GLContext

// GLattr is an OpenGL configuration attribute.
//(https://wiki.libsdl.org/SDL_GLattr)
type GLattr C.SDL_GLattr

// MessageBoxColor contains RGB value used in an MessageBoxColorScheme.
// (https://wiki.libsdl.org/SDL_MessageBoxColor)
type MessageBoxColor struct {
	R uint8 // the red component in the range 0-255
	G uint8 // the green component in the range 0-255
	B uint8 // the blue component in the range 0-255
}
type cMessageBoxColor C.SDL_MessageBoxColor

// MessageBoxColorScheme contains a set of colors to use for message box dialogs.
// (https://wiki.libsdl.org/SDL_MessageBoxColorScheme)
type MessageBoxColorScheme struct {
	Colors [5]MessageBoxColor // background, text, button border, button background, button selected
}
type cMessageBoxColorScheme C.SDL_MessageBoxColorScheme

// MessageBoxButtonData contains individual button data for a message box.
// (https://wiki.libsdl.org/SDL_MessageBoxButtonData)
type MessageBoxButtonData struct {
	Flags    uint32 // MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT, MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT
	ButtonID int32  // user defined button id (value returned via ShowMessageBox())
	Text     string // the UTF-8 button text
}

// MessageBoxData contains title, text, window and other data for a message box.
// (https://wiki.libsdl.org/SDL_MessageBoxData)
type MessageBoxData struct {
	Flags       uint32  // MESSAGEBOX_ERROR, MESSAGEBOX_WARNING, MESSAGEBOX_INFORMATION
	Window      *Window // parent window or nil
	Title       string
	Message     string
	Buttons     []MessageBoxButtonData
	ColorScheme *MessageBoxColorScheme // nil to use system settings
}

func (window *Window) cptr() *C.SDL_Window {
	return (*C.SDL_Window)(unsafe.Pointer(window))
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

// GetDisplayName returns the name of a display in UTF-8 encoding.
// (https://wiki.libsdl.org/SDL_GetDisplayName)
func GetDisplayName(displayIndex int) (string, error) {
	name := C.SDL_GetDisplayName(C.int(displayIndex))
	if name == nil {
		return "", GetError()
	}
	return C.GoString(name), nil
}

// GetNumVideoDisplays returns the number of available video displays.
// (https://wiki.libsdl.org/SDL_GetNumVideoDisplays)
func GetNumVideoDisplays() (int, error) {
	n := int(C.SDL_GetNumVideoDisplays())
	return n, errorFromInt(n)
}

// GetNumVideoDrivers returns the number of video drivers compiled into SDL.
// (https://wiki.libsdl.org/SDL_GetNumVideoDrivers)
func GetNumVideoDrivers() (int, error) {
	n := int(C.SDL_GetNumVideoDrivers())
	return n, errorFromInt(n)
}

// GetVideoDriver returns the name of a built in video driver.
// (https://wiki.libsdl.org/SDL_GetVideoDriver)
func GetVideoDriver(index int) string {
	return string(C.GoString(C.SDL_GetVideoDriver(C.int(index))))
}

// VideoInit initializes the video subsystem, optionally specifying a video driver.
// (https://wiki.libsdl.org/SDL_VideoInit)
func VideoInit(driverName string) error {
	return errorFromInt(int(
		C.SDL_VideoInit(C.CString(driverName))))
}

// VideoQuit shuts down the video subsystem, if initialized with VideoInit().
// (https://wiki.libsdl.org/SDL_VideoQuit)
func VideoQuit() {
	C.SDL_VideoQuit()
}

// GetCurrentVideoDriver returns the name of the currently initialized video driver.
// (https://wiki.libsdl.org/SDL_GetCurrentVideoDriver)
func GetCurrentVideoDriver() (string, error) {
	name := C.SDL_GetCurrentVideoDriver()
	if name == nil {
		return "", GetError()
	}
	return C.GoString(name), nil
}

// GetNumDisplayModes returns the number of available display modes.
// (https://wiki.libsdl.org/SDL_GetNumDisplayModes)
func GetNumDisplayModes(displayIndex int) (int, error) {
	n := int(C.SDL_GetNumDisplayModes(C.int(displayIndex)))
	return n, errorFromInt(n)
}

// GetDisplayBounds returns the desktop area represented by a display, with the primary display located at 0,0.
// (https://wiki.libsdl.org/SDL_GetDisplayBounds)
func GetDisplayBounds(displayIndex int) (rect Rect, err error) {
	err = errorFromInt(int(
		C.SDL_GetDisplayBounds(C.int(displayIndex), (&rect).cptr())))
	return
}

// GetDisplayUsableBounds returns the usable desktop area represented by a display, with the primary display located at 0,0.
// (https://wiki.libsdl.org/SDL_GetDisplayUsableBounds)
func GetDisplayUsableBounds(displayIndex int) (rect Rect, err error) {
	err = errorFromInt(int(
		C.SDL_GetDisplayUsableBounds(C.int(displayIndex), rect.cptr())))
	return
}

// GetDisplayMode returns information about a specific display mode.
// (https://wiki.libsdl.org/SDL_GetDisplayMode)
func GetDisplayMode(displayIndex int, modeIndex int) (mode DisplayMode, err error) {
	err = errorFromInt(int(
		C.SDL_GetDisplayMode(C.int(displayIndex), C.int(modeIndex), (&mode).cptr())))
	return
}

// GetDesktopDisplayMode returns information about the desktop display mode.
// (https://wiki.libsdl.org/SDL_GetDesktopDisplayMode)
func GetDesktopDisplayMode(displayIndex int) (mode DisplayMode, err error) {
	err = errorFromInt(int(
		C.SDL_GetDesktopDisplayMode(C.int(displayIndex), (&mode).cptr())))
	return
}

// GetCurrentDisplayMode returns information about the current display mode.
// (https://wiki.libsdl.org/SDL_GetCurrentDisplayMode)
func GetCurrentDisplayMode(displayIndex int) (mode DisplayMode, err error) {
	err = errorFromInt(int(
		C.SDL_GetCurrentDisplayMode(C.int(displayIndex), (&mode).cptr())))
	return
}

// GetClosestDisplayMode returns the closest match to the requested display mode.
// (https://wiki.libsdl.org/SDL_GetClosestDisplayMode)
func GetClosestDisplayMode(displayIndex int, mode *DisplayMode, closest *DisplayMode) (*DisplayMode, error) {
	m := (*DisplayMode)(unsafe.Pointer((C.SDL_GetClosestDisplayMode(C.int(displayIndex), mode.cptr(), closest.cptr()))))
	if m == nil {
		return nil, GetError()
	}
	return m, nil
}

// GetDisplayDPI returns the dots/pixels-per-inch for a display.
// (https://wiki.libsdl.org/SDL_GetDisplayDPI)
func GetDisplayDPI(displayIndex int) (ddpi, hdpi, vdpi float32, err error) {
	err = errorFromInt(int(
		C.SDL_GetDisplayDPI(C.int(displayIndex), (*C.float)(unsafe.Pointer(&ddpi)), (*C.float)(unsafe.Pointer(&hdpi)), (*C.float)(unsafe.Pointer(&vdpi)))))
	return
}

// GetDisplayIndex returns the index of the display associated with the window.
// (https://wiki.libsdl.org/SDL_GetWindowDisplayIndex)
func (window *Window) GetDisplayIndex() (int, error) {
	i := int(C.SDL_GetWindowDisplayIndex(window.cptr()))
	return i, errorFromInt(i)
}

// SetDisplayMode sets the display mode to use when the window is visible at fullscreen.
// (https://wiki.libsdl.org/SDL_SetWindowDisplayMode)
func (window *Window) SetDisplayMode(mode *DisplayMode) error {
	return errorFromInt(int(
		C.SDL_SetWindowDisplayMode(window.cptr(), mode.cptr())))
}

// GetDisplayMode fills in information about the display mode to use when the window is visible at fullscreen.
// (https://wiki.libsdl.org/SDL_GetWindowDisplayMode)
func (window *Window) GetDisplayMode() (mode DisplayMode, err error) {
	err = errorFromInt(int(
		C.SDL_GetWindowDisplayMode(window.cptr(), (&mode).cptr())))
	return
}

// GetPixelFormat returns the pixel format associated with the window.
// (https://wiki.libsdl.org/SDL_GetWindowPixelFormat)
func (window *Window) GetPixelFormat() (uint32, error) {
	f := (uint32)(C.SDL_GetWindowPixelFormat(window.cptr()))
	if f == PIXELFORMAT_UNKNOWN {
		return f, GetError()
	}
	return f, nil
}

// CreateWindow creates a window with the specified position, dimensions, and flags.
// (https://wiki.libsdl.org/SDL_CreateWindow)
func CreateWindow(title string, x, y, w, h int32, flags uint32) (*Window, error) {
	var _window = C.SDL_CreateWindow(C.CString(title), C.int(x), C.int(y), C.int(w), C.int(h), C.Uint32(flags))
	if _window == nil {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer(_window)), nil
}

// CreateWindowFrom creates an SDL window from an existing native window.
// (https://wiki.libsdl.org/SDL_CreateWindowFrom)
func CreateWindowFrom(data unsafe.Pointer) (*Window, error) {
	_window := C.SDL_CreateWindowFrom(data)
	if _window == nil {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer(_window)), nil
}

// Destroy destroys the window.
// (https://wiki.libsdl.org/SDL_DestroyWindow)
func (window *Window) Destroy() error {
	lastErr := GetError()
	ClearError()
	C.SDL_DestroyWindow(window.cptr())
	err := GetError()
	if err != nil {
		return err
	}
	SetError(lastErr)
	return nil
}

// GetID returns the numeric ID of the window, for logging purposes.
//  (https://wiki.libsdl.org/SDL_GetWindowID)
func (window *Window) GetID() (uint32, error) {
	id := uint32(C.SDL_GetWindowID(window.cptr()))
	if id == 0 {
		return 0, GetError()
	}
	return id, nil
}

// GetWindowFromID returns a window from a stored ID.
// (https://wiki.libsdl.org/SDL_GetWindowFromID)
func GetWindowFromID(id uint32) (*Window, error) {
	_window := C.SDL_GetWindowFromID(C.Uint32(id))
	if _window == nil {
		return nil, GetError()
	}
	return (*Window)(unsafe.Pointer((_window))), nil
}

// GetFlags returns the window flags.
// (https://wiki.libsdl.org/SDL_GetWindowFlags)
func (window *Window) GetFlags() uint32 {
	return (uint32)(C.SDL_GetWindowFlags(window.cptr()))
}

// SetTitle sets the title of the window.
// (https://wiki.libsdl.org/SDL_SetWindowTitle)
func (window *Window) SetTitle(title string) {
	C.SDL_SetWindowTitle(window.cptr(), C.CString(title))
}

// GetTitle returns the title of the window.
// (https://wiki.libsdl.org/SDL_GetWindowTitle)
func (window *Window) GetTitle() string {
	return C.GoString(C.SDL_GetWindowTitle(window.cptr()))
}

// SetIcon sets the icon for the window.
// (https://wiki.libsdl.org/SDL_SetWindowIcon)
func (window *Window) SetIcon(icon *Surface) {
	C.SDL_SetWindowIcon(window.cptr(), icon.cptr())
}

// SetData associates an arbitrary named pointer with the window.
// (https://wiki.libsdl.org/SDL_SetWindowData)
func (window *Window) SetData(name string, userdata unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(C.SDL_SetWindowData(window.cptr(), C.CString(name), userdata))
}

// GetData returns the data pointer associated with the window.
// (https://wiki.libsdl.org/SDL_GetWindowData)
func (window *Window) GetData(name string) unsafe.Pointer {
	return unsafe.Pointer(C.SDL_GetWindowData(window.cptr(), C.CString(name)))
}

// SetPosition sets the position of the window.
// (https://wiki.libsdl.org/SDL_SetWindowPosition)
func (window *Window) SetPosition(x, y int32) {
	C.SDL_SetWindowPosition(window.cptr(), C.int(x), C.int(y))
}

// GetPosition returns the position of the window.
// (https://wiki.libsdl.org/SDL_GetWindowPosition)
func (window *Window) GetPosition() (x, y int32) {
	var _x, _y C.int
	C.SDL_GetWindowPosition(window.cptr(), &_x, &_y)
	return int32(_x), int32(_y)
}

// SetResizable sets the user-resizable state of the window.
// (https://wiki.libsdl.org/SDL_SetWindowResizable)
func (window *Window) SetResizable(resizable bool) {
	C.SDL_SetWindowResizable(window.cptr(), C.SDL_bool(Btoi(resizable)))
}

// SetSize sets the size of the window's client area.
// (https://wiki.libsdl.org/SDL_SetWindowSize)
func (window *Window) SetSize(w, h int32) {
	C.SDL_SetWindowSize(window.cptr(), C.int(w), C.int(h))
}

// GetSize returns the size of the window's client area.
// (https://wiki.libsdl.org/SDL_GetWindowSize)
func (window *Window) GetSize() (w, h int32) {
	var _w, _h C.int
	C.SDL_GetWindowSize(window.cptr(), &_w, &_h)
	return int32(_w), int32(_h)
}

// SetMinimumSize sets the minimum size of the window's client area.
// (https://wiki.libsdl.org/SDL_SetWindowMinimumSize)
func (window *Window) SetMinimumSize(minW, minH int32) {
	C.SDL_SetWindowMinimumSize(window.cptr(), C.int(minW), C.int(minH))
}

// GetMinimumSize returns the minimum size of the window's client area.
// (https://wiki.libsdl.org/SDL_GetWindowMinimumSize)
func (window *Window) GetMinimumSize() (w, h int32) {
	var _w, _h C.int
	C.SDL_GetWindowMinimumSize(window.cptr(), &_w, &_h)
	return int32(_w), int32(_h)
}

// SetMaximumSize sets the maximum size of the window's client area.
// (https://wiki.libsdl.org/SDL_SetWindowMaximumSize)
func (window *Window) SetMaximumSize(maxW, maxH int32) {
	C.SDL_SetWindowMaximumSize(window.cptr(), C.int(maxW), C.int(maxH))
}

// GetMaximumSize returns the maximum size of the window's client area.
// (https://wiki.libsdl.org/SDL_GetWindowMaximumSize)
func (window *Window) GetMaximumSize() (w, h int32) {
	var _w, _h C.int
	C.SDL_GetWindowMaximumSize(window.cptr(), &_w, &_h)
	return int32(_w), int32(_h)
}

// SetBordered sets the border state of the window.
// (https://wiki.libsdl.org/SDL_SetWindowBordered)
func (window *Window) SetBordered(bordered bool) {
	C.SDL_SetWindowBordered(window.cptr(), C.SDL_bool(Btoi(bordered)))
}

// Show shows the window.
// (https://wiki.libsdl.org/SDL_ShowWindow)
func (window *Window) Show() {
	C.SDL_ShowWindow(window.cptr())
}

// Hide hides the window.
// (https://wiki.libsdl.org/SDL_HideWindow)
func (window *Window) Hide() {
	C.SDL_HideWindow(window.cptr())
}

// Raise raises the window above other windows and set the input focus.
// (https://wiki.libsdl.org/SDL_RaiseWindow)
func (window *Window) Raise() {
	C.SDL_RaiseWindow(window.cptr())
}

// Maximize makes the window as large as possible.
// (https://wiki.libsdl.org/SDL_MaximizeWindow)
func (window *Window) Maximize() {
	C.SDL_MaximizeWindow(window.cptr())
}

// Minimize minimizes the window to an iconic representation.
// (https://wiki.libsdl.org/SDL_MinimizeWindow)
func (window *Window) Minimize() {
	C.SDL_MinimizeWindow(window.cptr())
}

// Restore restores the size and position of a minimized or maximized window.
// (https://wiki.libsdl.org/SDL_RestoreWindow)
func (window *Window) Restore() {
	C.SDL_RestoreWindow(window.cptr())
}

// SetFullscreen sets the window's fullscreen state.
// (https://wiki.libsdl.org/SDL_SetWindowFullscreen)
func (window *Window) SetFullscreen(flags uint32) error {
	return errorFromInt(int(
		C.SDL_SetWindowFullscreen(window.cptr(), C.Uint32(flags))))
}

// GetSurface returns the SDL surface associated with the window.
// (https://wiki.libsdl.org/SDL_GetWindowSurface)
func (window *Window) GetSurface() (*Surface, error) {
	surface := (*Surface)(unsafe.Pointer(C.SDL_GetWindowSurface(window.cptr())))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// UpdateSurface copies the window surface to the screen.
// (https://wiki.libsdl.org/SDL_UpdateWindowSurface)
func (window *Window) UpdateSurface() error {
	return errorFromInt(int(
		C.SDL_UpdateWindowSurface(window.cptr())))
}

// UpdateSurfaceRects copies areas of the window surface to the screen.
// (https://wiki.libsdl.org/SDL_UpdateWindowSurfaceRects)
func (window *Window) UpdateSurfaceRects(rects []Rect) error {
	return errorFromInt(int(
		C.SDL_UpdateWindowSurfaceRects(window.cptr(), rects[0].cptr(), C.int(len(rects)))))
}

// SetGrab sets the window's input grab mode.
// (https://wiki.libsdl.org/SDL_SetWindowGrab)
func (window *Window) SetGrab(grabbed bool) {
	C.SDL_SetWindowGrab(window.cptr(), C.SDL_bool((Btoi(grabbed))))
}

// GetGrab returns the window's input grab mode.
// (https://wiki.libsdl.org/SDL_GetWindowGrab)
func (window *Window) GetGrab() bool {
	return C.SDL_GetWindowGrab(window.cptr()) != 0
}

// SetBrightness sets the brightness (gamma multiplier) for the display that owns the given window.
// (https://wiki.libsdl.org/SDL_SetWindowBrightness)
func (window *Window) SetBrightness(brightness float32) error {
	return errorFromInt(int(
		C.SDL_SetWindowBrightness(window.cptr(), C.float(brightness))))
}

// GetBrightness returns the brightness (gamma multiplier) for the display that owns the given window.
// (https://wiki.libsdl.org/SDL_GetWindowBrightness)
func (window *Window) GetBrightness() float32 {
	return float32(C.SDL_GetWindowBrightness(window.cptr()))
}

// SetGammaRamp sets the gamma ramp for the display that owns the given window.
// (https://wiki.libsdl.org/SDL_SetWindowGammaRamp)
func (window *Window) SetGammaRamp(red, green, blue *[256]uint16) error {
	return errorFromInt(int(
		C.SDL_SetWindowGammaRamp(
			window.cptr(),
			(*C.Uint16)(unsafe.Pointer(red)),
			(*C.Uint16)(unsafe.Pointer(green)),
			(*C.Uint16)(unsafe.Pointer(blue)))))
}

// GetGammaRamp returns the gamma ramp for the display that owns a given window.
// (https://wiki.libsdl.org/SDL_GetWindowGammaRamp)
func (window *Window) GetGammaRamp() (red, green, blue *[256]uint16, err error) {
	code := int(C.SDL_GetWindowGammaRamp(
		window.cptr(),
		(*C.Uint16)(unsafe.Pointer(red)),
		(*C.Uint16)(unsafe.Pointer(green)),
		(*C.Uint16)(unsafe.Pointer(blue))))
	return red, green, blue, errorFromInt(code)
}

// SetWindowOpacity sets the opacity of the window.
// (https://wiki.libsdl.org/SDL_SetWindowOpacity)
func (window *Window) SetWindowOpacity(opacity float32) error {
	return errorFromInt(int(
		C.SDL_SetWindowOpacity(window.cptr(), C.float(opacity))))
}

// GetWindowOpacity returns the opacity of the window.
// (https://wiki.libsdl.org/SDL_GetWindowOpacity)
func (window *Window) GetWindowOpacity() (opacity float32, err error) {
	return opacity, errorFromInt(int(
		C.SDL_GetWindowOpacity(window.cptr(), (*C.float)(unsafe.Pointer(&opacity)))))
}

// ShowSimpleMessageBox displays a simple modal message box.
// (https://wiki.libsdl.org/SDL_ShowSimpleMessageBox)
func ShowSimpleMessageBox(flags uint32, title, message string, window *Window) error {
	_title := C.CString(title)
	defer C.free(unsafe.Pointer(_title))
	_message := C.CString(message)
	defer C.free(unsafe.Pointer(_message))
	return errorFromInt(int(
		C.SDL_ShowSimpleMessageBox(C.Uint32(flags), _title, _message, window.cptr())))
}

// ShowMessageBox creates a modal message box.
// (https://wiki.libsdl.org/SDL_ShowMessageBox)
func ShowMessageBox(data *MessageBoxData) (buttonid int32, err error) {
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
			buttonid: C.int(btn.ButtonID),
			text:     ctext,
		}

		cbuttons = append(cbuttons, cbtn)
		cbtntexts = append(cbtntexts, ctext)
	}

	var buttonPtr *C.SDL_MessageBoxButtonData
	if len(cbuttons) > 0 {
		buttonPtr = &cbuttons[0]
	}
	cdata := C.SDL_MessageBoxData{
		flags:       C.Uint32(data.Flags),
		window:      data.Window.cptr(),
		title:       _title,
		message:     _message,
		numbuttons:  C.int(len(data.Buttons)),
		buttons:     buttonPtr,
		colorScheme: data.ColorScheme.cptr(),
	}

	buttonid = int32(C.ShowMessageBox(cdata))
	return buttonid, errorFromInt(int(buttonid))
}

// IsScreenSaverEnabled reports whether the screensaver is currently enabled.
// (https://wiki.libsdl.org/SDL_IsScreenSaverEnabled)
func IsScreenSaverEnabled() bool {
	return C.SDL_IsScreenSaverEnabled() != 0
}

// EnableScreenSaver allows the screen to be blanked by a screen saver.
// (https://wiki.libsdl.org/SDL_EnableScreenSaver)
func EnableScreenSaver() {
	C.SDL_EnableScreenSaver()
}

// DisableScreenSaver prevents the screen from being blanked by a screen saver.
// (https://wiki.libsdl.org/SDL_DisableScreenSaver)
func DisableScreenSaver() {
	C.SDL_DisableScreenSaver()
}

// GLLoadLibrary dynamically loads an OpenGL library.
// (https://wiki.libsdl.org/SDL_GL_LoadLibrary)
func GLLoadLibrary(path string) error {
	return errorFromInt(int(
		C.SDL_GL_LoadLibrary(C.CString(path))))
}

// GLGetProcAddress returns an OpenGL function by name.
// (https://wiki.libsdl.org/SDL_GL_GetProcAddress)
func GLGetProcAddress(proc string) unsafe.Pointer {
	return unsafe.Pointer(C.SDL_GL_GetProcAddress(C.CString(proc)))
}

// GLUnloadLibrary unloads the OpenGL library previously loaded by GLLoadLibrary().
// (https://wiki.libsdl.org/SDL_GL_UnloadLibrary)
func GLUnloadLibrary() {
	C.SDL_GL_UnloadLibrary()
}

// GLExtensionSupported reports whether an OpenGL extension is supported for the current context.
// (https://wiki.libsdl.org/SDL_GL_ExtensionSupported)
func GLExtensionSupported(extension string) bool {
	return C.SDL_GL_ExtensionSupported(C.CString(extension)) != 0
}

// GLSetAttribute sets an OpenGL window attribute before window creation.
// (https://wiki.libsdl.org/SDL_GL_SetAttribute)
func GLSetAttribute(attr GLattr, value int) error {
	return errorFromInt(int(
		C.SDL_GL_SetAttribute(attr.c(), C.int(value))))
}

// GLGetAttribute returns the actual value for an attribute from the current context.
// (https://wiki.libsdl.org/SDL_GL_GetAttribute)
func GLGetAttribute(attr GLattr) (int, error) {
	var _value C.int
	if C.SDL_GL_GetAttribute(attr.c(), &_value) != 0 {
		return int(_value), GetError()
	}
	return int(_value), nil
}

// GLCreateContext creates an OpenGL context for use with an OpenGL window, and make it current.
// (https://wiki.libsdl.org/SDL_GL_CreateContext)
func (window *Window) GLCreateContext() (GLContext, error) {
	c := GLContext(C.SDL_GL_CreateContext(window.cptr()))
	if c == nil {
		return nil, GetError()
	}
	return c, nil
}

// GLMakeCurrent sets up an OpenGL context for rendering into an OpenGL window.
// (https://wiki.libsdl.org/SDL_GL_MakeCurrent)
func (window *Window) GLMakeCurrent(glcontext GLContext) error {
	return errorFromInt(int(
		C.SDL_GL_MakeCurrent(window.cptr(), C.SDL_GLContext(glcontext))))
}

// GLSetSwapInterval sets the swap interval for the current OpenGL context.
// (https://wiki.libsdl.org/SDL_GL_SetSwapInterval)
func GLSetSwapInterval(interval int) error {
	return errorFromInt(int(
		C.SDL_GL_SetSwapInterval(C.int(interval))))
}

// GLGetSwapInterval returns the swap interval for the current OpenGL context.
// (https://wiki.libsdl.org/SDL_GL_GetSwapInterval)
func GLGetSwapInterval() (int, error) {
	i := int(C.SDL_GL_GetSwapInterval())
	// -1 means adaptive vsync, not an error
	// 0 means vsync off
	// 1 means vsync on
	if i == -1 || i == 0 || i == 1 {
		return i, nil
	}
	// any other value should be an error
	return i, errorFromInt(i)
}

// GLGetDrawableSize returns the size of a window's underlying drawable in pixels (for use with glViewport).
// (https://wiki.libsdl.org/SDL_GL_GetDrawableSize)
func (window *Window) GLGetDrawableSize() (w, h int32) {
	var _w, _h C.int
	C.SDL_GL_GetDrawableSize(window.cptr(), &_w, &_h)
	return int32(_w), int32(_h)
}

// GLSwap updates a window with OpenGL rendering.
// (https://wiki.libsdl.org/SDL_GL_SwapWindow)
func (window *Window) GLSwap() {
	C.SDL_GL_SwapWindow(window.cptr())
}

// GLDeleteContext deletes an OpenGL context.
// (https://wiki.libsdl.org/SDL_GL_DeleteContext)
func GLDeleteContext(context GLContext) {
	C.SDL_GL_DeleteContext(C.SDL_GLContext(context))
}

// Flash requests the window to demand attention from the user.
// (https://wiki.libsdl.org/SDL_FlashWindow)
func (window *Window) Flash(operation FlashOperation) (err error) {
	return errorFromInt(int(C.SDL_FlashWindow(window.cptr(), C.SDL_FlashOperation(operation))))
}

// SetAlwaysOnTop sets the window to always be above the others.
// (https://wiki.libsdl.org/SDL_SetWindowAlwaysOnTop)
func (window *Window) SetAlwaysOnTop(onTop bool) {
	C.SDL_SetWindowAlwaysOnTop(window.cptr(), C.SDL_bool(Btoi(onTop)))
}

// SetKeyboardGrab sets a window's keyboard grab mode.
// (https://wiki.libsdl.org/SDL_GetWindowKeyboardGrab)
func (window *Window) SetKeyboardGrab(grabbed bool) {
	C.SDL_SetWindowKeyboardGrab(window.cptr(), C.SDL_bool(Btoi(grabbed)))
}

// GetICCProfile gets the raw ICC profile data for the screen the window is currently on.
//
// Data returned should be freed with SDL_free.
//
// (https://wiki.libsdl.org/SDL_GetWindowICCProfile)
func (window *Window) GetICCProfile() (iccProfile unsafe.Pointer, size uintptr, err error) {
	_size := (*C.size_t)(unsafe.Pointer(&size))
	iccProfile = C.SDL_GetWindowICCProfile(window.cptr(), _size)
	if iccProfile == nil {
		err = GetError()
	}
	return
}

// SetMouseRect confines the cursor to the specified area of a window.
//
// Note that this does NOT grab the cursor, it only defines the area a cursor
// is restricted to when the window has mouse focus.
//
// (https://wiki.libsdl.org/SDL_SetWindowMouseRect)
func (window *Window) SetMouseRect(rect Rect) (err error) {
	_rect := (*C.SDL_Rect)(unsafe.Pointer(&rect))
	err = errorFromInt(int(C.SDL_SetWindowMouseRect(window.cptr(), _rect)))
	return
}

// GetMouseRect gets the mouse confinement rectangle of a window.
// (https://wiki.libsdl.org/SDL_GetWindowMouseRect)
func (window *Window) GetMouseRect() (rect Rect) {
	_rect := C.SDL_GetWindowMouseRect(window.cptr())
	rect = *((*Rect)(unsafe.Pointer(_rect)))
	return
}
