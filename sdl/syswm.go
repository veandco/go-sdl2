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

#if !(SDL_VERSION_ATLEAST(2,0,3))

#if defined(WARN_OUTDATED)
#pragma message("SDL_SYSWM_WINRT is not supported before SDL 2.0.3")
#endif

#define SDL_SYSWM_WINRT (0)
#endif

#if !(SDL_VERSION_ATLEAST(2,0,4))

#if defined(WARN_OUTDATED)
#pragma message("SDL_SYSWM_ANDROID is not supported before SDL 2.0.4")
#endif

#define SDL_SYSWM_ANDROID (0)
#endif

#if !(SDL_VERSION_ATLEAST(2,0,5))

#if defined(WARN_OUTDATED)
#pragma message("SDL_SYSWM_VIVANTE is not supported before SDL 2.0.5")
#endif

#define SDL_SYSWM_VIVANTE (0)
#endif
*/
import "C"
import "unsafe"

// Various supported windowing subsystems.
const (
	SYSWM_UNKNOWN  = C.SDL_SYSWM_UNKNOWN
	SYSWM_WINDOWS  = C.SDL_SYSWM_WINDOWS  // Microsoft Windows
	SYSWM_X11      = C.SDL_SYSWM_X11      // X Window System
	SYSWM_DIRECTFB = C.SDL_SYSWM_DIRECTFB // DirectFB
	SYSWM_COCOA    = C.SDL_SYSWM_COCOA    // Apple Mac OS X
	SYSWM_UIKIT    = C.SDL_SYSWM_UIKIT    // Apple iOS
	SYSWM_WAYLAND  = C.SDL_SYSWM_WAYLAND  // Wayland (>= SDL 2.0.2)
	SYSWM_MIR      = C.SDL_SYSWM_MIR      // Mir (>= SDL 2.0.2)
	SYSWM_WINRT    = C.SDL_SYSWM_WINRT    // WinRT (>= SDL 2.0.3)
	SYSWM_ANDROID  = C.SDL_SYSWM_ANDROID  // Android (>= SDL 2.0.4)
	SYSWM_VIVANTE  = C.SDL_SYSWM_VIVANTE  // Vivante (>= SDL 2.0.5)
)

// SysWMInfo contains system-dependent information about a window.
// (https://wiki.libsdl.org/SDL_SysWMinfo)
type SysWMInfo struct {
	Version   Version  // a Version structure that contains the current SDL version
	Subsystem uint32   // the windowing system type
	dummy     [24]byte // unused (to help compilers when no specific system is available)
}

// WindowsInfo contains Microsoft Windows window information.
type WindowsInfo struct {
	Window unsafe.Pointer // the window handle
	DeviceContext unsafe.Pointer // the device context handle
	Instance unsafe.Pointer // the instance handle
}

// X11Info contains X Window System window information.
type X11Info struct {
	Display unsafe.Pointer // the X11 display
	Window  uint           // the X11 window
}

// DFBInfo contains DirectFB window information.
type DFBInfo struct {
	Dfb     unsafe.Pointer // the DirectFB main interface
	Window  unsafe.Pointer // the DirectFB window handle
	Surface unsafe.Pointer // the DirectFB client surface
}

// CocoaInfo contains Apple Mac OS X window information.
type CocoaInfo struct {
	Window unsafe.Pointer // the Cocoa window
}

// UIKitInfo contains Apple iOS window information.
type UIKitInfo struct {
	Window unsafe.Pointer // the UIKit window
}

// SysWMmsg contains system-dependent window manager messages.
// (https://wiki.libsdl.org/SDL_SysWMmsg)
type SysWMmsg struct {
	Version   Version  // a Version structure that contains the current SDL version
	Subsystem uint32   // the windowing system type
	data      [24]byte // internal data
}

func (info *SysWMInfo) cptr() *C.SDL_SysWMinfo {
	return (*C.SDL_SysWMinfo)(unsafe.Pointer(info))
}

// GetWMInfo returns driver specific information about a window.
// (https://wiki.libsdl.org/SDL_GetWindowWMInfo)
func (window *Window) GetWMInfo() (*SysWMInfo, error) {
	var info SysWMInfo
	VERSION(&info.Version)
	if C.SDL_GetWindowWMInfo(window.cptr(), info.cptr()) == 0 {
		return nil, GetError()
	}
	return &info, nil
}

// GetWindowsInfo returns Microsoft Windows window information.
func (info *SysWMInfo) GetWindowsInfo() *WindowsInfo {
	return (*WindowsInfo)(unsafe.Pointer(&info.dummy[0]))
}

// GetX11Info returns X Window System window information.
func (info *SysWMInfo) GetX11Info() *X11Info {
	return (*X11Info)(unsafe.Pointer(&info.dummy[0]))
}

// GetDFBInfo returns DirectFB window information.
func (info *SysWMInfo) GetDFBInfo() *DFBInfo {
	return (*DFBInfo)(unsafe.Pointer(&info.dummy[0]))
}

// GetCocoaInfo returns Apple Mac OS X window information.
func (info *SysWMInfo) GetCocoaInfo() *CocoaInfo {
	return (*CocoaInfo)(unsafe.Pointer(&info.dummy[0]))
}

// GetUIKitInfo returns Apple iOS window information.
func (info *SysWMInfo) GetUIKitInfo() *UIKitInfo {
	return (*UIKitInfo)(unsafe.Pointer(&info.dummy[0]))
}
