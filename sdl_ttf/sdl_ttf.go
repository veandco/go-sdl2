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
	if f == nil {
		return nil,GetError()
	}
	return &Font{f},nil
}

func (f *Font) RenderText_Solid(text string, color sdl.Color) *sdl.Surface {
	_text := (C.CString) (text)
	// _color := (*C.SDL_Color) (unsafe.Pointer(&color))
	_c := C.SDL_Color{C.Uint8(color.R), C.Uint8(color.G), C.Uint8(color.B), C.Uint8(color.A)}
	surface := (*sdl.Surface) (unsafe.Pointer(C.TTF_RenderText_Solid(f.f, _text, _c)))
	return surface
}
