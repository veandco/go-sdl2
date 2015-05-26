package ttf

//#cgo windows LDFLAGS: -lSDL2 -lSDL2_ttf
//#cgo darwin LDFLAGS: -framework SDL2 -framework SDL2_ttf
//#cgo linux freebsd pkg-config: sdl2
//#cgo linux freebsd LDFLAGS: -lSDL2_ttf
//#include <stdlib.h>
//#include "sdl_ttf_wrapper.h"
//void Do_TTF_SetError(const char *str) {
//    TTF_SetError(str);
//}
import "C"
import "github.com/veandco/go-sdl2/sdl"
import "unsafe"
import "errors"

//Font Hinting Types
const (
	HINTING_NORMAL = int(C.TTF_HINTING_NORMAL)
	HINTING_LIGHT  = int(C.TTF_HINTING_LIGHT)
	HINTING_MONO   = int(C.TTF_HINTING_MONO)
	HINTING_NONE   = int(C.TTF_HINTING_NONE)
)

//Font Style Types
const (
	STYLE_NORMAL        = 0
	STYLE_BOLD          = 0x01
	STYLE_ITALIC        = 0x02
	STYLE_UNDERLINE     = 0x04
	STYLE_STRIKETHROUGH = 0x08
)

type Font struct {
	f *C.TTF_Font
}

func Init() int {
	return int(C.TTF_Init())
}

func WasInit() bool {
	return int(C.TTF_WasInit()) != 0
}

func Quit() {
	C.TTF_Quit()
}

func GetError() error {
	e := C.TTF_GetError()
	if e == nil {
		return nil
	}
	return errors.New(C.GoString(e))
}

func SetError(err string) {
	_err := C.CString(err)
	defer C.free(unsafe.Pointer(_err))
	C.Do_TTF_SetError(_err)
}

func ByteSwappedUnicode(swap bool) {
	val := 0
	if swap {
		val = 1
	}
	C.TTF_ByteSwappedUNICODE(C.int(val))
}

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

func OpenFontRW(src *sdl.RWops, freesrc, size int) (*Font, error) {
	return OpenFontIndexRW(src, freesrc, size, 0)
}

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

func (f *Font) RenderUTF8_Solid(text string, color sdl.Color) (*sdl.Surface,error) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	surface := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Solid(f.f, _text, _c)))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}
func (f *Font) RenderUTF8_Shaded(text string, fg, bg sdl.Color) (*sdl.Surface,error) {
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
func (f *Font) RenderUTF8_Blended(text string, color sdl.Color) (*sdl.Surface,error) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	surface := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Blended(f.f, _text, _c)))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}

func (f *Font) RenderUTF8_Blended_Wrapped(text string, fg sdl.Color,wrapLength int) (*sdl.Surface,error) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	_c := C.SDL_Color{C.Uint8(fg.R), C.Uint8(fg.G), C.Uint8(fg.B), C.Uint8(fg.A)}
	surface := (*sdl.Surface)(unsafe.Pointer(C.TTF_RenderUTF8_Blended_Wrapped(f.f, _text, _c, C.Uint32(wrapLength))))
	if surface == nil {
		return nil, GetError()
	}
	return surface, nil
}


func (f *Font) SizeUTF8(text string) (int,int,error) {
	_text := C.CString(text)
	defer C.free(unsafe.Pointer(_text))
	var w C.int
	var h C.int
	result :=  C.TTF_SizeUTF8(f.f,_text,&w,&h)
	if result == 0 {
		return int(w),int(h),nil
	}
	return int(w),int(h),GetError()
}


func (f *Font) Close() {
	C.TTF_CloseFont(f.f)
	f.f = nil
}

func (f *Font) Height() int   { return int(C.TTF_FontHeight(f.f)) }
func (f *Font) Ascent() int   { return int(C.TTF_FontAscent(f.f)) }
func (f *Font) Descent() int  { return int(C.TTF_FontDescent(f.f)) }
func (f *Font) LineSkip() int { return int(C.TTF_FontLineSkip(f.f)) }
func (f *Font) Faces() int    { return int(C.TTF_FontFaces(f.f)) }

func (f *Font) GetStyle() int {
	return int(C.TTF_GetFontStyle(f.f))
}

func (f *Font) SetStyle(style int) {
	C.TTF_SetFontStyle(f.f, C.int(style))
}

func (f *Font) GetHinting() int {
	return int(C.TTF_GetFontHinting(f.f))
}

func (f *Font) SetHinting(hinting int) {
	C.TTF_SetFontHinting(f.f, C.int(hinting))
}

func (f *Font) GetKerning() bool {
	return int(C.TTF_GetFontKerning(f.f)) == 1
}

func (f *Font) SetKerning(allowed bool) {
	val := 0
	if allowed {
		val = 1
	}
	C.TTF_SetFontKerning(f.f, C.int(val))
}

func (f *Font) GetOutline() int {
	return int(C.TTF_GetFontOutline(f.f))
}

func (f *Font) SetOutline(outline int) {
	C.TTF_SetFontOutline(f.f, C.int(outline))
}

func (f *Font) FaceIsFixedWidth() bool {
	return int(C.TTF_FontFaceIsFixedWidth(f.f)) != 0
}

func (f *Font) FaceFamilyName() string {
	_fname := C.TTF_FontFaceFamilyName(f.f)
	fname := C.GoString(_fname)
	return fname
}
