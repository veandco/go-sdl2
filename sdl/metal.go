package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,12))

typedef void *SDL_MetalView;

#if defined(WARN_OUTDATED)
#pragma message("SDL_Metal_CreateView is not supported before SDL 2.0.12")
#pragma message("SDL_Metal_DestroyView is not supported before SDL 2.0.12")
#endif

static SDL_MetalView SDL_Metal_CreateView(SDL_Window * window)
{
	return NULL;
}

static void SDL_Metal_DestroyView(SDL_MetalView view)
{
	// do nothing
}

#endif

#if !(SDL_VERSION_ATLEAST(2,0,14))

#if defined(WARN_OUTDATED)
#pragma message("SDL_Metal_GetLayer is not supported before SDL 2.0.14")
#pragma message("SDL_Metal_GetDrawableSize is not supported before SDL 2.0.14")
#endif

static void * SDL_Metal_GetLayer(SDL_MetalView view)
{
	return NULL;
}

static void SDL_Metal_GetDrawableSize(SDL_Window* window, int *w, int *h)
{
	// do nothing
}

#endif
*/
import "C"
import "unsafe"

// A handle to a CAMetalLayer-backed NSView (macOS) or UIView (iOS/tvOS).
type MetalView C.SDL_MetalView

// Metal_CreateView creates a CAMetalLayer-backed NSView/UIView and attach it to the specified window.
// (https://wiki.libsdl.org/SDL_Metal_CreateView)
func Metal_CreateView(window *Window) MetalView {
	return MetalView(C.SDL_Metal_CreateView(window.cptr()))
}

// Metal_DestroyView Destroy an existing SDL_MetalView object.
// (https://wiki.libsdl.org/SDL_Metal_DestroyView)
func Metal_DestroyView(metalView MetalView) {
	C.SDL_Metal_DestroyView(C.SDL_MetalView(metalView))
}

// Metal_GetLayer gets a pointer to the backing CAMetalLayer for the given view.
// (https://wiki.libsdl.org/SDL_Metal_GetLayer)
func Metal_GetLayer(metalView MetalView) unsafe.Pointer {
	return C.SDL_Metal_GetLayer(C.SDL_MetalView(metalView))
}

// Metal_GetDrawableSize Get the size of a window's underlying drawable in pixels (for use with setting viewport, scissor & etc).
// (https://wiki.libsdl.org/SDL_Metal_GetDrawableSize)
func Metal_GetDrawableSize(window *Window) (w, h int) {
	C.SDL_Metal_GetDrawableSize(window.cptr(), (*C.int)(unsafe.Pointer(&w)), (*C.int)(unsafe.Pointer((&h))))
	return
}
