package sdl

/*
#include "sdl_wrapper.h"
#include "log.h"

static inline void _SDL_Log(const char *fmt)
{
    SDL_Log("%s", fmt);
}

static inline void _SDL_LogVerbose(int category, const char *fmt)
{
    SDL_LogVerbose(category, "%s", fmt);
}

static inline void _SDL_LogDebug(int category, const char *fmt)
{
    SDL_LogDebug(category, "%s", fmt);
}

static inline void _SDL_LogInfo(int category, const char *fmt)
{
    SDL_LogInfo(category, "%s", fmt);
}

static inline void _SDL_LogWarn(int category, const char *fmt)
{
    SDL_LogWarn(category, "%s", fmt);
}

static inline void _SDL_LogError(int category, const char *fmt)
{
    SDL_LogError(category, "%s", fmt);
}

static inline void _SDL_LogCritical(int category, const char *fmt)
{
    SDL_LogCritical(category, "%s", fmt);
}

static inline void _SDL_LogMessage(int category, SDL_LogPriority priority, const char *fmt)
{
    SDL_LogCritical(category, "%s", fmt);
}
*/
import "C"
import "fmt"
import "unsafe"

// An enumeration of the predefined log categories.
// (https://wiki.libsdl.org/SDL_LOG_CATEGORY)
const (
	LOG_CATEGORY_APPLICATION = iota // application log
	LOG_CATEGORY_ERROR              // error log
	LOG_CATEGORY_ASSERT             // assert log
	LOG_CATEGORY_SYSTEM             // system log
	LOG_CATEGORY_AUDIO              // audio log
	LOG_CATEGORY_VIDEO              // video log
	LOG_CATEGORY_RENDER             // render log
	LOG_CATEGORY_INPUT              // input log
	LOG_CATEGORY_TEST               // test log
	LOG_CATEGORY_RESERVED1          // reserved for future SDL library use
	LOG_CATEGORY_RESERVED2          // reserved for future SDL library use
	LOG_CATEGORY_RESERVED3          // reserved for future SDL library use
	LOG_CATEGORY_RESERVED4          // reserved for future SDL library use
	LOG_CATEGORY_RESERVED5          // reserved for future SDL library use
	LOG_CATEGORY_RESERVED6          // reserved for future SDL library use
	LOG_CATEGORY_RESERVED7          // reserved for future SDL library use
	LOG_CATEGORY_RESERVED8          // reserved for future SDL library use
	LOG_CATEGORY_RESERVED9          // reserved for future SDL library use
	LOG_CATEGORY_RESERVED10         // reserved for future SDL library use
	LOG_CATEGORY_CUSTOM             // reserved for application use
)

// An enumeration of the predefined log priorities.
// (https://wiki.libsdl.org/SDL_LogPriority)
const (
	LOG_PRIORITY_VERBOSE  = iota + 1 // verbose
	LOG_PRIORITY_DEBUG               // debug
	LOG_PRIORITY_INFO                // info
	LOG_PRIORITY_WARN                // warn
	LOG_PRIORITY_ERROR               // error
	LOG_PRIORITY_CRITICAL            // critical
	NUM_LOG_PRIORITIES               // (internal use)
)

// LogPriority is a predefined log priority.
// (https://wiki.libsdl.org/SDL_LogPriority)
type LogPriority C.SDL_LogPriority

func (p LogPriority) c() C.SDL_LogPriority {
	return C.SDL_LogPriority(p)
}

// LogSetAllPriority sets the priority of all log categories.
// (https://wiki.libsdl.org/SDL_LogSetAllPriority)
func LogSetAllPriority(p LogPriority) {
	C.SDL_LogSetAllPriority(p.c())
}

// LogSetPriority sets the priority of a particular log category.
// (https://wiki.libsdl.org/SDL_LogSetPriority)
func LogSetPriority(category int, p LogPriority) {
	C.SDL_LogSetPriority(C.int(category), p.c())
}

// LogGetPriority returns the priority of a particular log category.
// (https://wiki.libsdl.org/SDL_LogGetPriority)
func LogGetPriority(category int) LogPriority {
	return LogPriority(C.SDL_LogGetPriority(C.int(category)))
}

// LogResetPriorities resets all priorities to default.
// (https://wiki.libsdl.org/SDL_LogResetPriorities)
func LogResetPriorities() {
	C.SDL_LogResetPriorities()
}

// Log logs a message with LOG_CATEGORY_APPLICATION and LOG_PRIORITY_INFO.
// (https://wiki.libsdl.org/SDL_Log)
func Log(str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_Log(cstr)
}

// LogVerbose logs a message with LOG_PRIORITY_VERBOSE.
// (https://wiki.libsdl.org/SDL_LogVerbose)
func LogVerbose(category int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogVerbose(C.int(category), cstr)
}

// LogDebug logs a message with LOG_PRIORITY_DEBUG.
// (https://wiki.libsdl.org/SDL_LogDebug)
func LogDebug(category int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogDebug(C.int(category), cstr)
}

// LogInfo logs a message with LOG_PRIORITY_INFO.
// (https://wiki.libsdl.org/SDL_LogInfo)
func LogInfo(category int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogInfo(C.int(category), cstr)
}

// LogWarn logs a message with LOG_PRIORITY_WARN.
// (https://wiki.libsdl.org/SDL_LogWarn)
func LogWarn(category int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogWarn(C.int(category), cstr)
}

// LogError logs a message with LOG_PRIORITY_ERROR.
// (https://wiki.libsdl.org/SDL_LogError)
func LogError(category int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogError(C.int(category), cstr)
}

// LogCritical logs a message with LOG_PRIORITY_CRITICAL.
// (https://wiki.libsdl.org/SDL_LogCritical)
func LogCritical(category int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogCritical(C.int(category), cstr)
}

// LogMessage logs a message with the specified category and priority.
// (https://wiki.libsdl.org/SDL_LogMessage)
func LogMessage(category int, pri LogPriority, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogMessage(C.int(category), C.SDL_LogPriority(pri), cstr)
}

// LogOutputFunction is the function to call instead of the default
type LogOutputFunction func(data interface{}, category int, pri LogPriority, message string)

type logOutputFunctionCtx struct {
	f LogOutputFunction
	d interface{}
}

// Yissakhar Z. Beck (DeedleFake)'s implementation
//
//export logOutputFunction
func logOutputFunction(data unsafe.Pointer, category C.int, pri C.SDL_LogPriority, message *C.char) {
	ctx := (*logOutputFunctionCtx)(data)

	ctx.f(ctx.d, int(category), LogPriority(pri), C.GoString(message))
}

var (
	logOutputFunctionCache LogOutputFunction
	logOutputDataCache     interface{}
)

// LogGetOutputFunction returns the current log output function.
// (https://wiki.libsdl.org/SDL_LogGetOutputFunction)
func LogGetOutputFunction() (LogOutputFunction, interface{}) {
	return logOutputFunctionCache, logOutputDataCache
}

// LogSetOutputFunction replaces the default log output function with one of your own.
// (https://wiki.libsdl.org/SDL_LogSetOutputFunction)
func LogSetOutputFunction(f LogOutputFunction, data interface{}) {
	ctx := &logOutputFunctionCtx{
		f: f,
		d: data,
	}

	C.LogSetOutputFunction(unsafe.Pointer(ctx))

	logOutputFunctionCache = f
	logOutputDataCache = data
}
