// Package ttf is a TrueType font rendering library that is used with the SDL library, and almost as portable. It depends on freetype2 to handle the TrueType font data. It allows a programmer to use multiple TrueType fonts without having to code a font rendering routine themselves. With the power of outline fonts and antialiasing, high quality text output can be obtained without much effort.
package ttf

/*
#include <stdlib.h>
#include "sdl_ttf_wrapper.h"
void Do_TTF_SetError(const char *str) {
    TTF_SetError("%s", str);
}

#if SDL_TTF_MAJOR_VERSION == 2 && SDL_TTF_MINOR_VERSION == 0 && SDL_TTF_PATCHLEVEL >= 18
static inline void ByteSwappedUNICODE(int swapped)
{
	TTF_ByteSwappedUNICODE(swapped ? SDL_TRUE : SDL_FALSE);
}
#else
static inline void ByteSwappedUNICODE(int swapped)
{
	TTF_ByteSwappedUNICODE(swapped);
}
#endif
*/
import "C"
import (
	"errors"
	"unsafe"

	"github.com/veandco/go-sdl2/sdl"
)

// Hinting settings.
const (
	HINTING_NORMAL = int(C.TTF_HINTING_NORMAL)
	HINTING_LIGHT  = int(C.TTF_HINTING_LIGHT)
	HINTING_MONO   = int(C.TTF_HINTING_MONO)
	HINTING_NONE   = int(C.TTF_HINTING_NONE)
)

// Font rendering styles.
const (
	STYLE_NORMAL        = 0
	STYLE_BOLD          = 0x01
	STYLE_ITALIC        = 0x02
	STYLE_UNDERLINE     = 0x04
	STYLE_STRIKETHROUGH = 0x08
)

// Font contains font information.
type Font struct {
	f *C.TTF_Font
}

// Init initializes the truetype font API. This must be called before using other functions in this library, except ttf.WasInit(). SDL does not have to be initialized before this call.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_8.html)
func Init() error {
	if C.TTF_Init() == -1 {
		return GetError()
	}
	return nil
}

// WasInit reports whether the truetype font API is initialized. Use this before ttf.Init() to avoid initializing twice in a row. Or use this to determine if you need to call ttf.Quit().
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_9.html)
func WasInit() bool {
	return int(C.TTF_WasInit()) != 0
}

// Quit shuts down and cleanups the truetype font API. After calling this the SDL_ttf functions should not be used, excepting ttf.WasInit(). You may, of course, use ttf.Init() to use the functionality again.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_10.html)
func Quit() {
	C.TTF_Quit()
}

// GetError returns the last error that occurred, or an empty string if there hasn't been an error message set since the last call to sdl.ClearError().
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_12.html)
func GetError() error {
	e := C.TTF_GetError()
	if e == nil {
		return nil
	}
	return errors.New(C.GoString(e))
}

// SetError sets the SDL error message to the specified string.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_11.html)
func SetError(err string) {
	_err := C.CString(err)
	defer C.free(unsafe.Pointer(_err))
	C.Do_TTF_SetError(_err)
}

// ByteSwappedUnicode tells SDL_ttf whether UNICODE (Uint16 per character) text is generally byteswapped.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_20.html)
func ByteSwappedUnicode(swap bool) {
	val := 0
	if swap {
		val = 1
	}
	C.ByteSwappedUNICODE(C.int(val))
}

// OpenFont loads file for use as a font, at the specified size. This is actually OpenFontIndex(file, size, 0). This can load TTF and FON files.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_14.html)
func OpenFont(file string, size int) (*Font, error) {
	_file := C.CString(file)
	defer C.free(unsafe.Pointer(_file))
	_size := (C.int)(size)
	f := (*C.TTF_Font)(C.TTF_OpenFont(_file, _size))

	if f == nil {
		return nil, GetError()
	}
	return &Font{f}, nil
}

// OpenFontIndex loads file, face index, for use as a font, at the specified size. This is actually OpenFontIndexRW(RWFromFile(file), size, index), but checks that the RWops it creates is not NULL. This can load TTF and FON files.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_16.html)
func OpenFontIndex(file string, size int, index int) (*Font, error) {
	_file := C.CString(file)
	defer C.free(unsafe.Pointer(_file))
	_size := (C.int)(size)
	_index := (C.long)(index)
	f := (*C.TTF_Font)(C.TTF_OpenFontIndex(_file, _size, _index))

	if f == nil {
		return nil, GetError()
	}
	return &Font{f}, nil
}

// OpenFontRW loads src for use as a font, at specified size. This is actually OpenFontIndexRW(src, freesrc, size, 0). This can load TTF and FON formats. Using SDL_RWops is not covered here, but they enable you to load from almost any source.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_15.html)
func OpenFontRW(src *sdl.RWops, freesrc, size int) (*Font, error) {
	return OpenFontIndexRW(src, freesrc, size, 0)
}

// OpenFontIndexRW loads src, face index, for use as a font, at the specified size. This can load TTF and FON formats. Using SDL_RWops is not covered here, but they enable you to load from almost any source.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_17.html)
func OpenFontIndexRW(src *sdl.RWops, freesrc, size, index int) (*Font, error) {
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	_freesrc := (C.int)(freesrc)
	_size := (C.int)(size)
	_index := (C.long)(index)
	f := (*C.TTF_Font)(C.TTF_OpenFontIndexRW(_src, _freesrc, _size, _index))

	if f == nil {
		return nil, GetError()
	}
	return &Font{f}, nil
}

// RenderUTF8Solid creates an 8-bit palettized surface and render the given text at fast quality with the given font and color.  The 0 pixel is the colorkey, giving a transparent background, and the 1 pixel is set to the text color.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_44.html)
func (f *Font) RenderUTF8Solid(text string, color sdl.Color) (*sdl.Surface, error) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	surface := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Solid(f.f, _text, _c)))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// RenderUTF8Shaded creates an 8-bit palettized surface and render the given text at high quality with the given font and colors. The 0 pixel is background, while other pixels have varying degrees of the foreground color.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_48.html)
func (f *Font) RenderUTF8Shaded(text string, fg, bg sdl.Color) (*sdl.Surface, error) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	_fg := C.SDL_Color{C.Uint8(fg.R), C.Uint8(fg.G), C.Uint8(fg.B), C.Uint8(fg.A)}
	_bg := C.SDL_Color{C.Uint8(bg.R), C.Uint8(bg.G), C.Uint8(bg.B), C.Uint8(bg.A)}
	surface := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Shaded(f.f, _text, _fg, _bg)))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// RenderUTF8Blended creates a 32-bit ARGB surface and render the given text at high quality, using alpha blending to dither the font with the given color.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_52.html)
func (f *Font) RenderUTF8Blended(text string, color sdl.Color) (*sdl.Surface, error) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	surface := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Blended(f.f, _text, _c)))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// RenderUTF8BlendedWrapped creates a 32-bit ARGB surface and render the given text at high quality, using alpha blending to dither the font with the given color. Text is wrapped to multiple lines on line endings and on word boundaries if it extends beyond wrapLength in pixels.
func (f *Font) RenderUTF8BlendedWrapped(text string, fg sdl.Color, wrapLength int) (*sdl.Surface, error) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	_c := C.SDL_Color{C.Uint8(fg.R), C.Uint8(fg.G), C.Uint8(fg.B), C.Uint8(fg.A)}
	surface := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Blended_Wrapped(f.f, _text, _c, C.Uint32(wrapLength))))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// SizeUTF8 returns the resulting surface size (width and height) of the UTF8 encoded text rendered using font. No actual rendering is done, however correct kerning is done to get the actual width. The height returned in h is the same as you can get using ttf.Height().
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_40.html)
func (f *Font) SizeUTF8(text string) (int, int, error) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	var w C.int
	var h C.int
	result := C.TTF_SizeUTF8(f.f, _text, &w, &h)
	if result == 0 {
		return int(w), int(h), nil
	}
	return int(w), int(h), GetError()
}

// RenderGlyphSolid renders the glyph for the UNICODE ch using font with fg color onto a new surface, using the Solid mode. The caller is responsible for freeing any returned surface.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf.html#SEC46)
func (f *Font) RenderGlyphSolid(ch rune, fg sdl.Color) (*sdl.Surface, error) {
	_ch := C.Uint16(ch)
	_fg := C.SDL_Color{C.Uint8(fg.R), C.Uint8(fg.G), C.Uint8(fg.B), C.Uint8(fg.A)}
	surface := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderGlyph_Solid(f.f, _ch, _fg)))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// RenderGlyphShaded renders the glyph for the UNICODE ch using font with fg color onto a new surface filled with the bg color, using the Shaded mode. The caller is responsible for freeing any returned surface.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf.html#SEC50)
func (f *Font) RenderGlyphShaded(ch rune, fg, bg sdl.Color) (*sdl.Surface, error) {
	_ch := C.Uint16(ch)
	_fg := C.SDL_Color{C.Uint8(fg.R), C.Uint8(fg.G), C.Uint8(fg.B), C.Uint8(fg.A)}
	_bg := C.SDL_Color{C.Uint8(bg.R), C.Uint8(bg.G), C.Uint8(bg.B), C.Uint8(bg.A)}
	surface := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderGlyph_Shaded(f.f, _ch, _fg, _bg)))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// RenderGlyphBlended renders the glyph for the UNICODE ch using font with fg color onto a new surface, using the Blended mode. The caller is responsible for freeing any returned surface.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf.html#SEC54)
func (f *Font) RenderGlyphBlended(ch rune, fg sdl.Color) (*sdl.Surface, error) {
	_ch := C.Uint16(ch)
	_fg := C.SDL_Color{C.Uint8(fg.R), C.Uint8(fg.G), C.Uint8(fg.B), C.Uint8(fg.A)}
	surface := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderGlyph_Blended(f.f, _ch, _fg)))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

// Close frees the memory used by font, and frees font itself as well. Do not use font after this without loading a new font to it.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_18.html)
func (f *Font) Close() {
	C.TTF_CloseFont(f.f)
	f.f = nil
}

// Height returns the maximum pixel height of all glyphs of the loaded font.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_29.html)
func (f *Font) Height() int { return int(C.TTF_FontHeight(f.f)) }

// Ascent returns the maximum pixel ascent of all glyphs of the loaded font. This can also be interpreted as the distance from the top of the font to the baseline.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_30.html)
func (f *Font) Ascent() int { return int(C.TTF_FontAscent(f.f)) }

// Descent returns the maximum pixel descent of all glyphs of the loaded font. This can also be interpreted as the distance from the baseline to the bottom of the font.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_31.html)
func (f *Font) Descent() int { return int(C.TTF_FontDescent(f.f)) }

// LineSkip returns the recommended pixel height of a rendered line of text of the loaded font.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_32.html)
func (f *Font) LineSkip() int { return int(C.TTF_FontLineSkip(f.f)) }

// Faces returns the number of faces ("sub-fonts") available in the loaded font.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_33.html)
func (f *Font) Faces() int { return int(C.TTF_FontFaces(f.f)) }

// GetStyle returns the rendering style of the loaded font.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_21.html)
func (f *Font) GetStyle() int {
	return int(C.TTF_GetFontStyle(f.f))
}

// SetStyle sets the rendering style of the loaded font.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_22.html)
func (f *Font) SetStyle(style int) {
	C.TTF_SetFontStyle(f.f, C.int(style))
}

// GetHinting returns the current hinting setting of the loaded font.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_25.html)
func (f *Font) GetHinting() int {
	return int(C.TTF_GetFontHinting(f.f))
}

// SetHinting sets the hinting of the loaded font.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_26.html)
func (f *Font) SetHinting(hinting int) {
	C.TTF_SetFontHinting(f.f, C.int(hinting))
}

// GetKerning returns the current kerning setting of the loaded font.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_27.html)
func (f *Font) GetKerning() bool {
	return int(C.TTF_GetFontKerning(f.f)) == 1
}

// SetKerning sets whther to use kerning when rendering the loaded font. This has no effect on individual glyphs, but rather when rendering whole strings of characters, at least a word at a time. Perhaps the only time to disable this is when kerning is not working for a specific font, resulting in overlapping glyphs or abnormal spacing within words.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_28.html)
func (f *Font) SetKerning(allowed bool) {
	val := 0
	if allowed {
		val = 1
	}
	C.TTF_SetFontKerning(f.f, C.int(val))
}

// GetOutline returns the current outline size of the loaded font.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_23.html)
func (f *Font) GetOutline() int {
	return int(C.TTF_GetFontOutline(f.f))
}

// SetOutline sets the outline pixel width of the loaded font.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_24.html)
func (f *Font) SetOutline(outline int) {
	C.TTF_SetFontOutline(f.f, C.int(outline))
}

// FaceIsFixedWidth reports whether the current font face of the loaded font is a fixed width font. Fixed width fonts are monospace, meaning every character that exists in the font is the same width, thus you can assume that a rendered string's width is going to be the result of a simple calculation: glyph_width * string_length.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_34.html)
func (f *Font) FaceIsFixedWidth() bool {
	return int(C.TTF_FontFaceIsFixedWidth(f.f)) != 0
}

// FaceFamilyName returns the current font face family name from the loaded font.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_35.html)
func (f *Font) FaceFamilyName() string {
	_fname := C.TTF_FontFaceFamilyName(f.f)
	fname := C.GoString(_fname)
	return fname
}

// FaceStyleName returns the current font face family's style name from the loaded font.
// (https://wiki.libsdl.org/SDL_ttf/TTF_FontFaceStyleName)
func (f *Font) FaceStyleName() string {
        _fname := C.TTF_FontFaceStyleName(f.f)
        fname := C.GoString(_fname)
        return fname
}

// GlyphMetrics contains glyph-specific rendering metrics.
type GlyphMetrics struct {
	MinX, MaxX int
	MinY, MaxY int
	Advance    int
}

// GlyphMetrics returns the desired glyph metrics of the UNICODE char given in ch from the loaded font.
// (https://www.libsdl.org/projects/SDL_ttf/docs/SDL_ttf_38.html)
func (f *Font) GlyphMetrics(ch rune) (*GlyphMetrics, error) {
	_ch := C.Uint16(ch)
	var minX, maxX, minY, maxY, adv C.int
	result := C.TTF_GlyphMetrics(f.f, _ch, &minX, &maxX, &minY, &maxY, &adv)
	if result == 0 {
		return &GlyphMetrics{int(minX), int(maxX), int(minY), int(maxY), int(adv)}, nil
	}
	return nil, GetError()
}
