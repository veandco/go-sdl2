package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,26,0))

#if defined(WARN_OUTDATED)
#pragma message("SDL_SetPrimarySelectionText is not supported before SDL 2.26.0")
#pragma message("SDL_GetPrimarySelectionText is not supported before SDL 2.26.0")
#pragma message("SDL_HasPrimarySelectionText is not supported before SDL 2.26.0")
#endif

static inline int SDL_SetPrimarySelectionText(const char *text)
{
    return -1;
}

static inline char * SDL_GetPrimarySelectionText(void)
{
    return NULL;
}

static inline SDL_bool SDL_HasPrimarySelectionText(void)
{
    return SDL_FALSE;
}

#endif
*/
import "C"
import "unsafe"

// SetClipboardText puts UTF-8 text into the clipboard.
// (https://wiki.libsdl.org/SDL_SetClipboardText)
func SetClipboardText(text string) error {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	if C.SDL_SetClipboardText(_text) < 0 {
		return GetError()
	}
	return nil
}

// GetClipboardText returns UTF-8 text from the clipboard.
// (https://wiki.libsdl.org/SDL_GetClipboardText)
func GetClipboardText() (string, error) {
	text := C.SDL_GetClipboardText()
	if text == nil {
		return "", GetError()
	}
	defer C.SDL_free(unsafe.Pointer(text))
	_text := C.GoString(text)
	return _text, nil
}

// HasClipboardText reports whether the clipboard exists and contains a text string that is non-empty.
// (https://wiki.libsdl.org/SDL_HasClipboardText)
func HasClipboardText() bool {
	return C.SDL_HasClipboardText() > 0
}

// SetPrimarySelectionText puts UTF-8 text into the primary selection.
// (https://wiki.libsdl.org/SDL_SetPrimarySelectionText)
func SetPrimarySelectionText(text string) error {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	if C.SDL_SetPrimarySelectionText(_text) < 0 {
		return GetError()
	}
	return nil
}

// GetPrimarySelectionText gets UTF-8 text from the primary selection.
// (https://wiki.libsdl.org/SDL_GetPrimarySelectionText)
func GetPrimarySelectionText() (string, error) {
	text := C.SDL_GetPrimarySelectionText()
	if text == nil {
		return "", GetError()
	}
	defer C.SDL_free(unsafe.Pointer(text))
	_text := C.GoString(text)
	return _text, nil
}

// HasPrimarySelectionText queries whether the primary selection exists and contains a non-empty text string.
// (https://wiki.libsdl.org/SDL_HasPrimarySelectionText)
func HasPrimarySelectionText() bool {
	return C.SDL_HasPrimarySelectionText() == C.SDL_TRUE
}
