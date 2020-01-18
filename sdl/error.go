package sdl

/*
#include "sdl_wrapper.h"

void GoSetError(const char *fmt) {
  SDL_SetError("%s", fmt);
}

*/
// #include "sdl_wrapper.h"
import "C"
import "errors"

var emptyCString *C.char = C.CString("")
var ErrInvalidParameters = errors.New("Invalid Parameters")

// SDL error codes with their corresponding predefined strings.
const (
	ENOMEM      ErrorCode = C.SDL_ENOMEM      // out of memory
	EFREAD                = C.SDL_EFREAD      // error reading from datastream
	EFWRITE               = C.SDL_EFWRITE     // error writing to datastream
	EFSEEK                = C.SDL_EFSEEK      // error seeking in datastream
	UNSUPPORTED           = C.SDL_UNSUPPORTED // that operation is not supported
	LASTERROR             = C.SDL_LASTERROR   // the highest numbered predefined error
)

// ErrorCode is an error code used in SDL error messages.
type ErrorCode uint32
type cErrorCode C.SDL_errorcode

func (ec ErrorCode) c() C.SDL_errorcode {
	return C.SDL_errorcode(ec)
}

// GetError returns the last error that occurred, or an empty string if there hasn't been an error message set since the last call to ClearError().
// (https://wiki.libsdl.org/SDL_GetError)
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

// SetError set the SDL error message.
// (https://wiki.libsdl.org/SDL_SetError)
func SetError(err error) {
	if err != nil {
		C.GoSetError(C.CString(err.Error()))
		return
	}
	C.GoSetError(emptyCString)
}

// ClearError clears any previous error message.
// (https://wiki.libsdl.org/SDL_ClearError)
func ClearError() {
	C.SDL_ClearError()
}

// Error sets the SDL error message to the specified error code.
func Error(code ErrorCode) {
	C.SDL_Error(code.c())
}

// OutOfMemory sets SDL error message to ENOMEM (out of memory).
func OutOfMemory() {
	Error(ENOMEM)
}

// Unsupported sets SDL error message to UNSUPPORTED (that operation is not supported).
func Unsupported() {
	Error(UNSUPPORTED)
}

// errorFromInt returns GetError() if passed negative value, otherwise it returns nil.
func errorFromInt(code int) (err error) {
	if code < 0 {
		err = GetError()
		if err == nil {
			err = errors.New("Unknown error (probably using old version of SDL2 and the function called is not supported?)")
		}
	}
	return
}
