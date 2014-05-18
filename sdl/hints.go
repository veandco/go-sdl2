package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"

const (
	HINT_FRAMEBUFFER_ACCELERATION           = "SDL_FRAMEBUFFER_ACCELERATION"
	HINT_RENDER_DRIVER                      = "SDL_RENDER_DRIVER"
	HINT_RENDER_OPENGL_SHADERS              = "SDL_RENDER_OPENGL_SHADERS"
	HINT_RENDER_DIRECT3D_THREADSAFE         = "SDL_RENDER_DIRECT3D_THREADSAFE"
	HINT_RENDER_DIRECT3D11_DEBUG            = "SDL_HINT_RENDER_DIRECT3D11_DEBUG"
	HINT_RENDER_SCALE_QUALITY               = "SDL_RENDER_SCALE_QUALITY"
	HINT_RENDER_VSYNC                       = "SDL_RENDER_VSYNC"
	HINT_VIDEO_ALLOW_SCREENSAVER            = "SDL_VIDEO_ALLOW_SCREENSAVER"
	HINT_VIDEO_X11_XVIDMODE                 = "SDL_VIDEO_X11_XVIDMODE"
	HINT_VIDEO_X11_XINERAMA                 = "SDL_VIDEO_X11_XINERAMA"
	HINT_VIDEO_X11_XRANDR                   = "SDL_VIDEO_X11_XRANDR"
	HINT_GRAB_KEYBOARD                      = "SDL_GRAB_KEYBOARD"
	HINT_MOUSE_RELATIVE_MODE_WARP           = "SDL_MOUSE_RELATIVE_MODE_WARP"
	HINT_VIDEO_MINIMIZE_ON_FOCUS_LOSS       = "SDL_VIDEO_MINIMIZE_ON_FOCUS_LOSS"
	HINT_IDLE_TIMER_DISABLED                = "SDL_IOS_IDLE_TIMER_DISABLED"
	HINT_ORIENTATIONS                       = "SDL_IOS_ORIENTATIONS"
	HINT_ACCELEROMETER_AS_JOYSTICK          = "SDL_ACCELEROMETER_AS_JOYSTICK"
	HINT_XINPUT_ENABLED                     = "SDL_XINPUT_ENABLED"
	HINT_GAMECONTROLLERCONFIG               = "SDL_GAMECONTROLLERCONFIG"
	HINT_JOYSTICK_ALLOW_BACKGROUND_EVENTS   = "SDL_JOYSTICK_ALLOW_BACKGROUND_EVENTS"
	HINT_ALLOW_TOPMOST                      = "SDL_ALLOW_TOPMOST"
	HINT_TIMER_RESOLUTION                   = "SDL_TIMER_RESOLUTION"
	HINT_VIDEO_HIGHDPI_DISABLED             = "SDL_VIDEO_HIGHDPI_DISABLED"
	HINT_MAC_CTRL_CLICK_EMULATE_RIGHT_CLICK = "SDL_MAC_CTRL_CLICK_EMULATE_RIGHT_CLICK"
	HINT_VIDEO_WIN_D3DCOMPILER              = "SDL_VIDEO_WIN_D3DCOMPILER"
	HINT_VIDEO_WINDOW_SHARE_PIXEL_FORMAT    = "SDL_VIDEO_WINDOW_SHARE_PIXEL_FORMAT"
	HINT_WINRT_PRIVACY_POLICY_URL           = "SDL_HINT_WINRT_PRIVACY_POLICY_URL"
	HINT_WINRT_PRIVACY_POLICY_LABEL         = "SDL_HINT_WINRT_PRIVACY_POLICY_LABEL"
	HINT_WINRT_HANDLE_BACK_BUTTON           = "SDL_HINT_WINRT_HANDLE_BACK_BUTTON"
	HINT_VIDEO_MAC_FULLSCREEN_SPACES        = "SDL_VIDEO_MAC_FULLSCREEN_SPACES"
)

const (
	HINT_DEFAULT = iota
	HINT_NORMAL
	HINT_OVERRIDE
)

type HintPriority C.SDL_HintPriority

func (hp HintPriority) c() C.SDL_HintPriority {
    return C.SDL_HintPriority(hp)
}

func SetHintWithPriority(name string, value string, hp HintPriority) bool {
	_name := C.CString(name)
	_value := C.CString(value)
	defer C.free(unsafe.Pointer(_name))
	defer C.free(unsafe.Pointer(_value))
	return C.SDL_SetHintWithPriority(_name, _value, hp.c()) > 0
}

func SetHint(name string, value string) bool {
	_name := C.CString(name)
	_value := C.CString(value)
	defer C.free(unsafe.Pointer(_name))
	defer C.free(unsafe.Pointer(_value))
	return C.SDL_SetHint(_name, _value) > 0
}

func GetHint(name string) string {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	return C.GoString(C.SDL_GetHint(_name))
}

func ClearHints() {
	C.SDL_ClearHints()
}

/* TODO:
typedef void (*SDL_HintCallback)(void *userdata, const char *name, const char *oldValue, const char *newValue);
extern DECLSPEC void SDLCALL SDL_AddHintCallback(const char *name, SDL_HintCallback callback, void *userdata);
extern DECLSPEC void SDLCALL SDL_DelHintCallback(const char *name, SDL_HintCallback callback, void *userdata);
*/
