package sdl

/*
#include "sdl_wrapper.h"
#include "hints.h"

#if !(SDL_VERSION_ATLEAST(2,0,4))
#define SDL_HINT_NO_SIGNAL_HANDLERS ""
#define SDL_HINT_THREAD_STACK_SIZE ""
#define SDL_HINT_WINDOW_FRAME_USABLE_WHILE_CURSOR_HIDDEN ""
#define SDL_HINT_WINDOWS_ENABLE_MESSAGELOOP ""
#define SDL_HINT_WINDOWS_NO_CLOSE_ON_ALT_F4 ""
#define SDL_HINT_XINPUT_USE_OLD_JOYSTICK_MAPPING ""
#define SDL_HINT_MAC_BACKGROUND_APP ""
#define SDL_HINT_IME_INTERNAL_EDITING ""
#define SDL_HINT_VIDEO_X11_NET_WM_PING ""
#define SDL_HINT_ANDROID_SEPARATE_MOUSE_AND_TOUCH ""
#define SDL_HINT_ANDROID_APK_EXPANSION_MAIN_FILE_VERSION ""
#define SDL_HINT_ANDROID_APK_EXPANSION_PATCH_FILE_VERSION ""
#endif
#if !(SDL_VERSION_ATLEAST(2,0,3))
#define SDL_HINT_WINRT_PRIVACY_POLICY_URL ""
#define SDL_HINT_WINRT_PRIVACY_POLICY_LABEL ""
#define SDL_HINT_WINRT_HANDLE_BACK_BUTTON ""
#define SDL_HINT_RENDER_DIRECT3D11_DEBUG ""
#endif
#if !(SDL_VERSION_ATLEAST(2,0,2))
#define SDL_HINT_ACCELEROMETER_AS_JOYSTICK ""
#define SDL_HINT_MAC_CTRL_CLICK_EMULATE_RIGHT_CLICK ""
#define SDL_HINT_VIDEO_ALLOW_SCREENSAVER ""
#define SDL_HINT_MOUSE_RELATIVE_MODE_WARP ""
#define SDL_HINT_VIDEO_WIN_D3DCOMPILER ""
#define SDL_HINT_VIDEO_WINDOW_SHARE_PIXEL_FORMAT ""
#define SDL_HINT_VIDEO_MAC_FULLSCREEN_SPACES ""
#endif
#if !(SDL_VERSION_ATLEAST(2,0,1))
#define SDL_HINT_RENDER_DIRECT3D_THREADSAFE ""
#define SDL_HINT_VIDEO_HIGHDPI_DISABLED ""
#endif
*/
import "C"
import "unsafe"

// Configuration hints
// (https://wiki.libsdl.org/CategoryHints)
const (
	HINT_FRAMEBUFFER_ACCELERATION                 = C.SDL_HINT_FRAMEBUFFER_ACCELERATION                 // specifies how 3D acceleration is used with Window.GetSurface()
	HINT_RENDER_DRIVER                            = C.SDL_HINT_RENDER_DRIVER                            // specifies which render driver to use
	HINT_RENDER_OPENGL_SHADERS                    = C.SDL_HINT_RENDER_OPENGL_SHADERS                    // specifies whether the OpenGL render driver uses shaders
	HINT_RENDER_DIRECT3D_THREADSAFE               = C.SDL_HINT_RENDER_DIRECT3D_THREADSAFE               // specifies whether the Direct3D device is initialized for thread-safe operations
	HINT_RENDER_DIRECT3D11_DEBUG                  = C.SDL_HINT_RENDER_DIRECT3D11_DEBUG                  // specifies a variable controlling whether to enable Direct3D 11+'s Debug Layer
	HINT_RENDER_SCALE_QUALITY                     = C.SDL_HINT_RENDER_SCALE_QUALITY                     // specifies scaling quality
	HINT_RENDER_VSYNC                             = C.SDL_HINT_RENDER_VSYNC                             // specifies whether sync to vertical refresh is enabled or disabled in CreateRenderer() to avoid tearing
	HINT_VIDEO_ALLOW_SCREENSAVER                  = C.SDL_HINT_VIDEO_ALLOW_SCREENSAVER                  // specifies whether the screensaver is enabled
	HINT_VIDEO_X11_NET_WM_PING                    = C.SDL_HINT_VIDEO_X11_NET_WM_PING                    // specifies whether the X11 _NET_WM_PING protocol should be supported
	HINT_VIDEO_X11_XVIDMODE                       = C.SDL_HINT_VIDEO_X11_XVIDMODE                       // specifies whether the X11 VidMode extension should be used
	HINT_VIDEO_X11_XINERAMA                       = C.SDL_HINT_VIDEO_X11_XINERAMA                       // specifies whether the X11 Xinerama extension should be used
	HINT_VIDEO_X11_XRANDR                         = C.SDL_HINT_VIDEO_X11_XRANDR                         // specifies whether the X11 XRandR extension should be used
	HINT_GRAB_KEYBOARD                            = C.SDL_HINT_GRAB_KEYBOARD                            // specifies whether grabbing input grabs the keyboard
	HINT_MOUSE_RELATIVE_MODE_WARP                 = C.SDL_HINT_MOUSE_RELATIVE_MODE_WARP                 // specifies whether relative mouse mode is implemented using mouse warping
	HINT_VIDEO_MINIMIZE_ON_FOCUS_LOSS             = C.SDL_HINT_VIDEO_MINIMIZE_ON_FOCUS_LOSS             // specifies if a Window is minimized if it loses key focus when in fullscreen mode
	HINT_IDLE_TIMER_DISABLED                      = C.SDL_HINT_IDLE_TIMER_DISABLED                      // specifies a variable controlling whether the idle timer is disabled on iOS
	HINT_IME_INTERNAL_EDITING                     = C.SDL_HINT_IME_INTERNAL_EDITING                     // specifies whether certain IMEs should handle text editing internally instead of sending TextEditingEvents
	HINT_ORIENTATIONS                             = C.SDL_HINT_ORIENTATIONS                             // specifies a variable controlling which orientations are allowed on iOS
	HINT_ACCELEROMETER_AS_JOYSTICK                = C.SDL_HINT_ACCELEROMETER_AS_JOYSTICK                // specifies whether the Android / iOS built-in accelerometer should be listed as a joystick device, rather than listing actual joysticks only
	HINT_XINPUT_ENABLED                           = C.SDL_HINT_XINPUT_ENABLED                           // specifies if Xinput gamepad devices are detected
	HINT_XINPUT_USE_OLD_JOYSTICK_MAPPING          = C.SDL_HINT_XINPUT_USE_OLD_JOYSTICK_MAPPING          // specifies that SDL should use the old axis and button mapping for XInput devices
	HINT_GAMECONTROLLERCONFIG                     = C.SDL_HINT_GAMECONTROLLERCONFIG                     // specifies extra gamecontroller db entries
	HINT_JOYSTICK_ALLOW_BACKGROUND_EVENTS         = C.SDL_HINT_JOYSTICK_ALLOW_BACKGROUND_EVENTS         // specifies if joystick (and gamecontroller) events are enabled even when the application is in the background
	HINT_ALLOW_TOPMOST                            = C.SDL_HINT_ALLOW_TOPMOST                            // specifies if top most bit on an SDL Window can be set
	HINT_THREAD_STACK_SIZE                        = C.SDL_HINT_THREAD_STACK_SIZE                        // specifies a variable specifying SDL's threads stack size in bytes or "0" for the backend's default size
	HINT_TIMER_RESOLUTION                         = C.SDL_HINT_TIMER_RESOLUTION                         // specifies the timer resolution in milliseconds
	HINT_VIDEO_HIGHDPI_DISABLED                   = C.SDL_HINT_VIDEO_HIGHDPI_DISABLED                   // specifies if high-DPI windows ("Retina" on Mac and iOS) are not allowed
	HINT_MAC_BACKGROUND_APP                       = C.SDL_HINT_MAC_BACKGROUND_APP                       // specifies if the SDL app should not be forced to become a foreground process on Mac OS X
	HINT_MAC_CTRL_CLICK_EMULATE_RIGHT_CLICK       = C.SDL_HINT_MAC_CTRL_CLICK_EMULATE_RIGHT_CLICK       // specifies whether ctrl+click should generate a right-click event on Mac
	HINT_VIDEO_WIN_D3DCOMPILER                    = C.SDL_HINT_VIDEO_WIN_D3DCOMPILER                    // specifies which shader compiler to preload when using the Chrome ANGLE binaries
	HINT_VIDEO_WINDOW_SHARE_PIXEL_FORMAT          = C.SDL_HINT_VIDEO_WINDOW_SHARE_PIXEL_FORMAT          // specifies the address of another Window* (as a hex string formatted with "%p")
	HINT_WINRT_PRIVACY_POLICY_URL                 = C.SDL_HINT_WINRT_PRIVACY_POLICY_URL                 // specifies a URL to a WinRT app's privacy policy
	HINT_WINRT_PRIVACY_POLICY_LABEL               = C.SDL_HINT_WINRT_PRIVACY_POLICY_LABEL               // specifies a label text for a WinRT app's privacy policy link
	HINT_WINRT_HANDLE_BACK_BUTTON                 = C.SDL_HINT_WINRT_HANDLE_BACK_BUTTON                 // specifies a variable to allow back-button-press events on Windows Phone to be marked as handled
	HINT_VIDEO_MAC_FULLSCREEN_SPACES              = C.SDL_HINT_VIDEO_MAC_FULLSCREEN_SPACES              // specifies policy for fullscreen Spaces on Mac OS X
	HINT_NO_SIGNAL_HANDLERS                       = C.SDL_HINT_NO_SIGNAL_HANDLERS                       // specifies not to catch the SIGINT or SIGTERM signals
	HINT_WINDOW_FRAME_USABLE_WHILE_CURSOR_HIDDEN  = C.SDL_HINT_WINDOW_FRAME_USABLE_WHILE_CURSOR_HIDDEN  // specifies whether the window frame and title bar are interactive when the cursor is hidden
	HINT_WINDOWS_ENABLE_MESSAGELOOP               = C.SDL_HINT_WINDOWS_ENABLE_MESSAGELOOP               // specifies whether the windows message loop is processed by SDL
	HINT_WINDOWS_NO_CLOSE_ON_ALT_F4               = C.SDL_HINT_WINDOWS_NO_CLOSE_ON_ALT_F4               // specifies that SDL should not to generate WINDOWEVENT_CLOSE events for Alt+F4 on Microsoft Windows
	HINT_ANDROID_SEPARATE_MOUSE_AND_TOUCH         = C.SDL_HINT_ANDROID_SEPARATE_MOUSE_AND_TOUCH         // specifies a variable to control whether mouse and touch events are to be treated together or separately
	HINT_ANDROID_APK_EXPANSION_MAIN_FILE_VERSION  = C.SDL_HINT_ANDROID_APK_EXPANSION_MAIN_FILE_VERSION  // specifies the Android APK expansion main file version
	HINT_ANDROID_APK_EXPANSION_PATCH_FILE_VERSION = C.SDL_HINT_ANDROID_APK_EXPANSION_PATCH_FILE_VERSION // specifies the Android APK expansion patch file version
)

// An enumeration of hint priorities.
// (https://wiki.libsdl.org/SDL_HintPriority)
const (
	HINT_DEFAULT  = C.SDL_HINT_DEFAULT  // low priority, used for default values
	HINT_NORMAL   = C.SDL_HINT_NORMAL   // medium priority
	HINT_OVERRIDE = C.SDL_HINT_OVERRIDE // high priority
)

// HintCallback is the function to call when the hint value changes.
type HintCallback func(data interface{}, name, oldValue, newValue string)

// HintCallbackAndData contains a callback function and userdata.
type HintCallbackAndData struct {
	callback HintCallback // the function to call when the hint value changes
	data     interface{}  // data to pass to the callback function
}

var hintCallbacks = make(map[string]HintCallbackAndData)

// HintPriority is a hint priority used in SetHintWithPriority().
// (https://wiki.libsdl.org/SDL_HintPriority)
type HintPriority C.SDL_HintPriority

func (hp HintPriority) c() C.SDL_HintPriority {
	return C.SDL_HintPriority(hp)
}

// SetHintWithPriority sets a hint with a specific priority.
// (https://wiki.libsdl.org/SDL_SetHintWithPriority)
func SetHintWithPriority(name, value string, hp HintPriority) bool {
	_name := C.CString(name)
	_value := C.CString(value)
	defer C.free(unsafe.Pointer(_name))
	defer C.free(unsafe.Pointer(_value))
	return C.SDL_SetHintWithPriority(_name, _value, hp.c()) > 0
}

// SetHint sets a hint with normal priority.
// (https://wiki.libsdl.org/SDL_SetHint)
func SetHint(name, value string) bool {
	_name := C.CString(name)
	_value := C.CString(value)
	defer C.free(unsafe.Pointer(_name))
	defer C.free(unsafe.Pointer(_value))
	return C.SDL_SetHint(_name, _value) > 0
}

// GetHint returns the value of a hint.
// (https://wiki.libsdl.org/SDL_GetHint)
func GetHint(name string) string {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	return C.GoString(C.SDL_GetHint(_name))
}

// ClearHints clears all hints.
// (https://wiki.libsdl.org/SDL_ClearHints)
func ClearHints() {
	C.SDL_ClearHints()
}

// AddHintCallback adds a function to watch a particular hint.
// (https://wiki.libsdl.org/SDL_AddHintCallback)
func AddHintCallback(name string, fn HintCallback, data interface{}) {
	_name := C.CString(name)
	hintCallbacks[name] = HintCallbackAndData{
		callback: fn,
		data:     data,
	}
	C.addHintCallback(_name)
}

// DelHintCallback removes a function watching a particular hint.
// (https://wiki.libsdl.org/SDL_DelHintCallback)
func DelHintCallback(name string) {
	_name := C.CString(name)
	delete(hintCallbacks, name)
	C.delHintCallback(_name)
}

//export goHintCallback
func goHintCallback(_name, _oldValue, _newValue *C.char) {
	name := C.GoString(_name)
	oldValue := C.GoString(_oldValue)
	newValue := C.GoString(_newValue)
	if cb, ok := hintCallbacks[name]; ok {
		cb.callback(cb.data, name, oldValue, newValue)
	}
}
