package sdl

// #include <SDL2/SDL_error.h>
import "C"
import "errors"

const (
	ENOMEM				= 0x00000000
	EFREAD				= 0x00000001
	EFWRITE				= 0x00000002
	EFSEEK				= 0x00000003
	UNSUPPORTED			= 0x00000004
	LASTERROR			= 0x00000005
)

func GetError() error {
	_c_err := C.SDL_GetError()
	if _c_err == nil {
		return nil
	}
	return errors.New(C.GoString(_c_err))
}

func ClearError() {
	C.SDL_ClearError()
}

func Error(code int) {
	_code := (C.SDL_errorcode) (C.int(code))
	C.SDL_Error(_code)
}

func OutOfMemory() {
	Error(ENOMEM)
}

func Unsupported() {
	Error(UNSUPPORTED)
}
