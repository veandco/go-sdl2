package sdl

/*
#include "sdl_wrapper.h"
#include "system.h"

#if !(SDL_VERSION_ATLEAST(2,0,16))

#if defined(WARN_OUTDATED)
#pragma message("SDL_RenderGetD3D11Device is not supported before SDL 2.0.16")
#endif

typedef struct ID3D11Device ID3D11Device;

static ID3D11Device* SDL_RenderGetD3D11Device(SDL_Renderer * renderer)
{
	return NULL;
}

#endif
*/
import "C"
import (
	"unsafe"
)

type WindowsMessageHook func(userdata interface{}, hWnd unsafe.Pointer, message uint32, wParam uint64, lParam int64)
var windowsMessageHook WindowsMessageHook

type ID3D11Device C.ID3D11Device;

// SetWindowsMessageHook sets a callback for every Windows message, run before TranslateMessage().
// (https://wiki.libsdl.org/SDL_SetWindowsMessageHook)
func SetWindowsMessageHook(callback WindowsMessageHook, userdata interface{}) {
	windowsMessageHook = callback
	C.SetWindowsMessageHook()
}

//export goWindowsMessageHook
func goWindowsMessageHook(userdata interface{}, hWnd unsafe.Pointer, message uint32, wParam uint64, lParam int64) {
	if windowsMessageHook == nil {
		return
	}
	windowsMessageHook(userdata, hWnd, message, wParam, lParam)
}

// SDL_RenderGetD3D11Device gets the D3D11 device associated with a renderer.
// (https://wiki.libsdl.org/SDL_RenderGetD3D11Device)
func (renderer *Renderer) GetD3D11Device() (device *ID3D11Device, err error) {
	device = (*ID3D11Device)(C.SDL_RenderGetD3D11Device(renderer.cptr()))
	if device == nil {
		err = GetError()
	}
	return
}
