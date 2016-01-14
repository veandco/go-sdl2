package sdl

/*
#include "sdl_wrapper.h"

#if defined(__WIN32)
#include <SDL2/SDL_syswm.h>
#else
#include <SDL_syswm.h>
#endif

#if !(SDL_VERSION_ATLEAST(2,0,2))
#define SDL_SYSWM_WAYLAND SDL_SYSWM_UNKNOWN
#define SDL_SYSWM_MIR SDL_SYSWM_UNKNOWN
#endif
*/
import "C"
import "unsafe"

const (
	SYSWM_UNKNOWN  = C.SDL_SYSWM_UNKNOWN
	SYSWM_WINDOWS  = C.SDL_SYSWM_WINDOWS
	SYSWM_X11      = C.SDL_SYSWM_X11
	SYSWM_DIRECTFB = C.SDL_SYSWM_DIRECTFB
	SYSWM_COCOA    = C.SDL_SYSWM_COCOA
	SYSWM_UIKIT    = C.SDL_SYSWM_UIKIT
	SYSWM_WAYLAND  = C.SDL_SYSWM_WAYLAND
	SYSWM_MIR      = C.SDL_SYSWM_MIR
)

// SysWMInfo (https://wiki.libsdl.org/SDL_SysWMinfo)
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

// Window (https://wiki.libsdl.org/SDL_GetWindowWMInfo)
func (window *Window) GetWMInfo() (*SysWMInfo, error) {
	var info SysWMInfo
	VERSION(&info.Version)
	if C.SDL_GetWindowWMInfo(window.cptr(), info.cptr()) == 0 {
		return nil, GetError()
	}
	return &info, nil
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
