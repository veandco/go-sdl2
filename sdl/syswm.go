package sdl

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
	Version   Version
	Subsystem uint32
	dummy     [24]byte
}

type WindowsInfo struct {
	Window unsafe.Pointer
}

type X11Info struct {
	Display unsafe.Pointer
	Window  uint
}

type DFBInfo struct {
	Dfb     unsafe.Pointer
	Window  unsafe.Pointer
	Surface unsafe.Pointer
}

type CocoaInfo struct {
	Window unsafe.Pointer
}

type UIKitInfo struct {
	Window unsafe.Pointer
}

func (info *SysWMInfo) cptr() *C.SDL_SysWMinfo {
    return (*C.SDL_SysWMinfo)(unsafe.Pointer(info))
}

func (window *Window) GetWMInfo(info *SysWMInfo) bool {
	return C.SDL_GetWindowWMInfo(window.cptr(), info.cptr()) == 1
}

func (info *SysWMInfo) GetWindowsInfo() *WindowsInfo {
	return (*WindowsInfo)(unsafe.Pointer(&info.dummy[0]))
}

func (info *SysWMInfo) GetX11Info() *X11Info {
	return (*X11Info)(unsafe.Pointer(&info.dummy[0]))
}

func (info *SysWMInfo) GetDFBInfo() *DFBInfo {
	return (*DFBInfo)(unsafe.Pointer(&info.dummy[0]))
}

func (info *SysWMInfo) GetCocoaInfo() *CocoaInfo {
	return (*CocoaInfo)(unsafe.Pointer(&info.dummy[0]))
}

func (info *SysWMInfo) GetUIKitInfo() *UIKitInfo {
	return (*UIKitInfo)(unsafe.Pointer(&info.dummy[0]))
}
