package ttf

//#cgo LDFLAGS: -lSDL2 -lSDL2_ttf
//#include <SDL2/SDL_ttf.h>
import "C"
import "github.com/jackyb/go-sdl2/sdl"
import "unsafe"
import "errors"

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

func OpenFont(file string, size int) (*Font,error) {
	_file := (C.CString) (file)
	_size := (C.int) (size)
	f := (*C.TTF_Font) (C.TTF_OpenFont(_file, _size))
	C.free(unsafe.Pointer(_file))
	if f == nil {
		return nil,GetError()
	}
	return &Font{f},nil
}

func (f *Font) RenderText_Solid(text string, color sdl.Color) *sdl.Surface {
	_text := C.CString(text)
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	surface := (*sdl.Surface) (unsafe.Pointer(C.TTF_RenderText_Solid(f.f, _text, _c)))
	return surface
}

func (f *Font) RenderText_Blended(text string, color sdl.Color) *sdl.Surface {
	_text := C.CString(text)
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	surface := (*sdl.Surface) (unsafe.Pointer(C.TTF_RenderText_Blended(f.f, _text, _c)))
	return surface
}

func (f *Font) RenderText_Shaded(text string, fg, bg sdl.Color) *sdl.Surface {
	_text := C.CString(text)
	_fg := C.SDL_Color{C.Uint8(fg.R), C.Uint8(fg.G), C.Uint8(fg.B), C.Uint8(fg.A)}
	_bg := C.SDL_Color{C.Uint8(bg.R), C.Uint8(bg.G), C.Uint8(bg.B), C.Uint8(bg.A)}
	surface := (*sdl.Surface) (unsafe.Pointer(C.TTF_RenderText_Shaded(f.f, _text, _fg, _bg)))
	return surface
}

func (f *Font) Close() { C.TTF_CloseFont(f.f) }

func (f *Font) Height() int { return int(C.TTF_FontHeight(f.f)) }
func (f *Font) Ascent() int { return int(C.TTF_FontAscent(f.f)) }
func (f *Font) Descent() int { return int(C.TTF_FontDescent(f.f)) }
func (f *Font) LineSkip() int { return int(C.TTF_FontLineSkip(f.f)) }
func (f *Font) Faces() int { return int(C.TTF_FontFaces(f.f)) }

func (f *Font) FaceIsFixedWidth() bool {
	return int(C.TTF_FontFaceIsFixedWidth(f.f)) != 0
}

func (f *Font) FaceFamilyName() string {
	_fname := C.TTF_FontFaceFamilyName(f.f)
	fname := C.GoString(_fname)
	C.free(unsafe.Pointer(_fname))
	return fname
}
