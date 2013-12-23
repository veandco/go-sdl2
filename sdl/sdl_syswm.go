package sdl

// #cgo LDFLAGS: -lSDL2
// #include <SDL2/SDL.h>
// #include <SDL2/SDL_syswm.h>
import "C"
import "unsafe"

const (
    SYSWM_UNKNOWN = iota
    SYSWM_WINDOWS
    SYSWM_X11
    SYSWM_DIRECTFB
    SYSWM_COCOA
    SYSWM_UIKIT
)

type SysWMInfo struct {
	Version	  Version
	Subsystem uint32
	dummy	  [32]byte
}

func (window *Window) GetWMInfo(info *SysWMInfo) bool {
	_window := (*C.SDL_Window) (unsafe.Pointer(window))
	_info := (*C.SDL_SysWMinfo) (unsafe.Pointer(info))
	return C.SDL_GetWindowWMInfo(_window, _info) == 1
}
