package sdl

// #include <SDL2/SDL.h>
import "C"
import "errors"

const (
	ENOMEM      = 0x00000000
	EFREAD      = 0x00000001
	EFWRITE     = 0x00000002
	EFSEEK      = 0x00000003
	UNSUPPORTED = 0x00000004
	LASTERROR   = 0x00000005
)

type ErrorCode uint

func (ec ErrorCode) c() C.SDL_errorcode {
    return C.SDL_errorcode(ec)
}

func GetError() error {
	_err := C.SDL_GetError()
	if *_err == 0 {
		return nil
	}
	return errors.New(C.GoString(_err))
}

func ClearError() {
	C.SDL_ClearError()
}

func Error(code ErrorCode) {
	C.SDL_Error(code.c())
}

func OutOfMemory() {
	Error(ENOMEM)
}

func Unsupported() {
	Error(UNSUPPORTED)
}
