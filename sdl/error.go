package sdl

// #include "sdl_wrapper.h"
import "C"
import "errors"

const (
	ENOMEM      = C.SDL_ENOMEM
	EFREAD      = C.SDL_EFREAD
	EFWRITE     = C.SDL_EFWRITE
	EFSEEK      = C.SDL_EFSEEK
	UNSUPPORTED = C.SDL_UNSUPPORTED
	LASTERROR   = C.SDL_LASTERROR
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
