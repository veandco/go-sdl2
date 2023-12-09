package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,14))

#if defined(WARN_OUTDATED)
#pragma message("SDL_GetErrorMsg is not supported before SDL 2.0.14")
#endif

static char * SDL_GetErrorMsg(char *errstr, int maxlen)
{
	return errstr;
}

#endif

void GoSetError(const char *fmt) {
  SDL_SetError("%s", fmt);
}

*/
import "C"
import (
	"errors"
	"unsafe"
)

var emptyCString *C.char = C.CString("")
var ErrInvalidParameters = errors.New("Invalid Parameters")

// SDL error codes with their corresponding predefined strings.
const (
	ENOMEM      ErrorCode = C.SDL_ENOMEM      // out of memory
	EFREAD      ErrorCode = C.SDL_EFREAD      // error reading from datastream
	EFWRITE     ErrorCode = C.SDL_EFWRITE     // error writing to datastream
	EFSEEK      ErrorCode = C.SDL_EFSEEK      // error seeking in datastream
	UNSUPPORTED ErrorCode = C.SDL_UNSUPPORTED // that operation is not supported
	LASTERROR   ErrorCode = C.SDL_LASTERROR   // the highest numbered predefined error
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

// GetErrorMsg returns the last error message that was set for the current thread.
// (https://wiki.libsdl.org/SDL_GetErrorMsg)
func GetErrorMsg(maxlen int) error {
	_buf := C.SDL_malloc(C.size_t(maxlen))
	if err := C.SDL_GetErrorMsg((*C.char)(_buf), C.int(maxlen)); err != nil {
		gostr := C.GoString(err)
		// SDL_GetErrorMsg returns "an empty string if there hasn't been an error message"
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
		_err := C.CString(err.Error())
		defer C.free(unsafe.Pointer(_err))
		C.GoSetError(_err)
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
