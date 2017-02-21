package sdl

// #include "sdl_wrapper.h"
import "C"
import "errors"

const (
	ENOMEM      ErrorCode = C.SDL_ENOMEM
	EFREAD                = C.SDL_EFREAD
	EFWRITE               = C.SDL_EFWRITE
	EFSEEK                = C.SDL_EFSEEK
	UNSUPPORTED           = C.SDL_UNSUPPORTED
	LASTERROR             = C.SDL_LASTERROR
)

type ErrorCode uint32
type cErrorCode C.SDL_errorcode

func (ec ErrorCode) c() C.SDL_errorcode {
	return C.SDL_errorcode(ec)
}

// GetError (https://wiki.libsdl.org/SDL_GetError)
func GetError() error {
	if err := C.SDL_GetError(); err != nil {
		gostr := C.GoString(err)
		// SDL_GetError returns "an empty string if there hasn't been an error message"
		if len(gostr) > 0 {
			return errors.New(gostr)
		}
	}
	return nil
}

// ClearError (https://wiki.libsdl.org/SDL_ClearError)
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
