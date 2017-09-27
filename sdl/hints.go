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

// Hint Names
const (
	HINT_FRAMEBUFFER_ACCELERATION                 = C.SDL_HINT_FRAMEBUFFER_ACCELERATION
	HINT_RENDER_DRIVER                            = C.SDL_HINT_RENDER_DRIVER
	HINT_RENDER_OPENGL_SHADERS                    = C.SDL_HINT_RENDER_OPENGL_SHADERS
	HINT_RENDER_DIRECT3D_THREADSAFE               = C.SDL_HINT_RENDER_DIRECT3D_THREADSAFE
	HINT_RENDER_DIRECT3D11_DEBUG                  = C.SDL_HINT_RENDER_DIRECT3D11_DEBUG
	HINT_RENDER_SCALE_QUALITY                     = C.SDL_HINT_RENDER_SCALE_QUALITY
	HINT_RENDER_VSYNC                             = C.SDL_HINT_RENDER_VSYNC
	HINT_VIDEO_ALLOW_SCREENSAVER                  = C.SDL_HINT_VIDEO_ALLOW_SCREENSAVER
	HINT_VIDEO_X11_NET_WM_PING                    = C.SDL_HINT_VIDEO_X11_NET_WM_PING
	HINT_VIDEO_X11_XVIDMODE                       = C.SDL_HINT_VIDEO_X11_XVIDMODE
	HINT_VIDEO_X11_XINERAMA                       = C.SDL_HINT_VIDEO_X11_XINERAMA
	HINT_VIDEO_X11_XRANDR                         = C.SDL_HINT_VIDEO_X11_XRANDR
	HINT_GRAB_KEYBOARD                            = C.SDL_HINT_GRAB_KEYBOARD
	HINT_MOUSE_RELATIVE_MODE_WARP                 = C.SDL_HINT_MOUSE_RELATIVE_MODE_WARP
	HINT_VIDEO_MINIMIZE_ON_FOCUS_LOSS             = C.SDL_HINT_VIDEO_MINIMIZE_ON_FOCUS_LOSS
	HINT_IDLE_TIMER_DISABLED                      = C.SDL_HINT_IDLE_TIMER_DISABLED
	HINT_IME_INTERNAL_EDITING                     = C.SDL_HINT_IME_INTERNAL_EDITING
	HINT_ORIENTATIONS                             = C.SDL_HINT_ORIENTATIONS
	HINT_ACCELEROMETER_AS_JOYSTICK                = C.SDL_HINT_ACCELEROMETER_AS_JOYSTICK
	HINT_XINPUT_ENABLED                           = C.SDL_HINT_XINPUT_ENABLED
	HINT_XINPUT_USE_OLD_JOYSTICK_MAPPING          = C.SDL_HINT_XINPUT_USE_OLD_JOYSTICK_MAPPING
	HINT_GAMECONTROLLERCONFIG                     = C.SDL_HINT_GAMECONTROLLERCONFIG
	HINT_JOYSTICK_ALLOW_BACKGROUND_EVENTS         = C.SDL_HINT_JOYSTICK_ALLOW_BACKGROUND_EVENTS
	HINT_ALLOW_TOPMOST                            = C.SDL_HINT_ALLOW_TOPMOST
	HINT_THREAD_STACK_SIZE                        = C.SDL_HINT_THREAD_STACK_SIZE
	HINT_TIMER_RESOLUTION                         = C.SDL_HINT_TIMER_RESOLUTION
	HINT_VIDEO_HIGHDPI_DISABLED                   = C.SDL_HINT_VIDEO_HIGHDPI_DISABLED
	HINT_MAC_BACKGROUND_APP                       = C.SDL_HINT_MAC_BACKGROUND_APP
	HINT_MAC_CTRL_CLICK_EMULATE_RIGHT_CLICK       = C.SDL_HINT_MAC_CTRL_CLICK_EMULATE_RIGHT_CLICK
	HINT_VIDEO_WIN_D3DCOMPILER                    = C.SDL_HINT_VIDEO_WIN_D3DCOMPILER
	HINT_VIDEO_WINDOW_SHARE_PIXEL_FORMAT          = C.SDL_HINT_VIDEO_WINDOW_SHARE_PIXEL_FORMAT
	HINT_WINRT_PRIVACY_POLICY_URL                 = C.SDL_HINT_WINRT_PRIVACY_POLICY_URL
	HINT_WINRT_PRIVACY_POLICY_LABEL               = C.SDL_HINT_WINRT_PRIVACY_POLICY_LABEL
	HINT_WINRT_HANDLE_BACK_BUTTON                 = C.SDL_HINT_WINRT_HANDLE_BACK_BUTTON
	HINT_VIDEO_MAC_FULLSCREEN_SPACES              = C.SDL_HINT_VIDEO_MAC_FULLSCREEN_SPACES
	HINT_NO_SIGNAL_HANDLERS                       = C.SDL_HINT_NO_SIGNAL_HANDLERS
	HINT_WINDOW_FRAME_USABLE_WHILE_CURSOR_HIDDEN  = C.SDL_HINT_WINDOW_FRAME_USABLE_WHILE_CURSOR_HIDDEN
	HINT_WINDOWS_ENABLE_MESSAGELOOP               = C.SDL_HINT_WINDOWS_ENABLE_MESSAGELOOP
	HINT_WINDOWS_NO_CLOSE_ON_ALT_F4               = C.SDL_HINT_WINDOWS_NO_CLOSE_ON_ALT_F4
	HINT_ANDROID_SEPARATE_MOUSE_AND_TOUCH         = C.SDL_HINT_ANDROID_SEPARATE_MOUSE_AND_TOUCH
	HINT_ANDROID_APK_EXPANSION_MAIN_FILE_VERSION  = C.SDL_HINT_ANDROID_APK_EXPANSION_MAIN_FILE_VERSION
	HINT_ANDROID_APK_EXPANSION_PATCH_FILE_VERSION = C.SDL_HINT_ANDROID_APK_EXPANSION_PATCH_FILE_VERSION
)

const (
	HINT_DEFAULT  = C.SDL_HINT_DEFAULT
	HINT_NORMAL   = C.SDL_HINT_NORMAL
	HINT_OVERRIDE = C.SDL_HINT_OVERRIDE
)

type HintCallback func(data interface{}, name, oldValue, newValue string)
type HintCallbackAndData struct {
	callback HintCallback
	data interface{}
}

var hintCallbacks = make(map[string]HintCallbackAndData)

// HintPriority (https://wiki.libsdl.org/SDL_HintPriority)
type HintPriority C.SDL_HintPriority

func (hp HintPriority) c() C.SDL_HintPriority {
	return C.SDL_HintPriority(hp)
}

// SetHintWithPriority (https://wiki.libsdl.org/SDL_SetHintWithPriority)
func SetHintWithPriority(name, value string, hp HintPriority) bool {
	_name := C.CString(name)
	_value := C.CString(value)
	defer C.free(unsafe.Pointer(_name))
	defer C.free(unsafe.Pointer(_value))
	return C.SDL_SetHintWithPriority(_name, _value, hp.c()) > 0
}

// SetHint (https://wiki.libsdl.org/SDL_SetHint)
func SetHint(name, value string) bool {
	_name := C.CString(name)
	_value := C.CString(value)
	defer C.free(unsafe.Pointer(_name))
	defer C.free(unsafe.Pointer(_value))
	return C.SDL_SetHint(_name, _value) > 0
}

// GetHint (https://wiki.libsdl.org/SDL_GetHint)
func GetHint(name string) string {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	return C.GoString(C.SDL_GetHint(_name))
}

// ClearHints (https://wiki.libsdl.org/SDL_ClearHints)
func ClearHints() {
	C.SDL_ClearHints()
}

// AddHintCallback (https://wiki.libsdl.org/SDL_AddHintCallback)
func AddHintCallback(name string, fn HintCallback, data interface{}) {
	_name := C.CString(name)
	hintCallbacks[name] = HintCallbackAndData{
		callback: fn,
		data: data,
	}
	C.addHintCallback(_name)
}

// DelHintCallback (https://wiki.libsdl.org/SDL_DelHintCallback)
func DelHintCallback(name string) {
	_name := C.CString(name)
	delete(hintCallbacks, name)
	C.delHintCallback(_name)
}

//export GoHintCallback
func GoHintCallback(_name, _oldValue, _newValue *C.char) {
	name := C.GoString(_name)
	oldValue := C.GoString(_oldValue)
	newValue := C.GoString(_newValue)
	if cb, ok := hintCallbacks[name]; ok {
		cb.callback(cb.data, name, oldValue, newValue)
	}
}
