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

const (
	LOG_CATEGORY_APPLICATION = iota
	LOG_CATEGORY_ERROR
	LOG_CATEGORY_ASSERT
	LOG_CATEGORY_SYSTEM
	LOG_CATEGORY_AUDIO
	LOG_CATEGORY_VIDEO
	LOG_CATEGORY_RENDER
	LOG_CATEGORY_INPUT
	LOG_CATEGORY_TEST
	LOG_CATEGORY_RESERVED1
	LOG_CATEGORY_RESERVED2
	LOG_CATEGORY_RESERVED3
	LOG_CATEGORY_RESERVED4
	LOG_CATEGORY_RESERVED5
	LOG_CATEGORY_RESERVED6
	LOG_CATEGORY_RESERVED7
	LOG_CATEGORY_RESERVED8
	LOG_CATEGORY_RESERVED9
	LOG_CATEGORY_RESERVED10
	LOG_CATEGORY_CUSTOM
)

const (
	LOG_PRIORITY_VERBOSE = iota + 1
	LOG_PRIORITY_DEBUG
	LOG_PRIORITY_INFO
	LOG_PRIORITY_WARN
	LOG_PRIORITY_ERROR
	LOG_PRIORITY_CRITICAL
	NUM_LOG_PRIORITIES
)

// LogPriority (https://wiki.libsdl.org/SDL_LogPriority)
type LogPriority C.SDL_LogPriority

func (p LogPriority) c() C.SDL_LogPriority {
	return C.SDL_LogPriority(p)
}

// LogSetAllPriority (https://wiki.libsdl.org/SDL_LogSetAllPriority)
func LogSetAllPriority(p LogPriority) {
	C.SDL_LogSetAllPriority(p.c())
}

// LogSetPriority (https://wiki.libsdl.org/SDL_LogSetPriority)
func LogSetPriority(category int, p LogPriority) {
	C.SDL_LogSetPriority(C.int(category), p.c())
}

// LogGetPriority (https://wiki.libsdl.org/SDL_LogGetPriority)
func LogGetPriority(category int) LogPriority {
	return LogPriority(C.SDL_LogGetPriority(C.int(category)))
}

// LogResetPriorities (https://wiki.libsdl.org/SDL_LogResetPriorities)
func LogResetPriorities() {
	C.SDL_LogResetPriorities()
}

// Log (https://wiki.libsdl.org/SDL_Log)
func Log(str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_Log(cstr)
}

// LogVerbose (https://wiki.libsdl.org/SDL_LogVerbose)
func LogVerbose(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogVerbose(C.int(cat), cstr)
}

// LogDebug (https://wiki.libsdl.org/SDL_LogDebug)
func LogDebug(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogDebug(C.int(cat), cstr)
}

// LogInfo (https://wiki.libsdl.org/SDL_LogInfo)
func LogInfo(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogInfo(C.int(cat), cstr)
}

// LogWarn (https://wiki.libsdl.org/SDL_LogWarn)
func LogWarn(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogWarn(C.int(cat), cstr)
}

// LogError (https://wiki.libsdl.org/SDL_LogError)
func LogError(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogError(C.int(cat), cstr)
}

// LogCritical (https://wiki.libsdl.org/SDL_LogCritical)
func LogCritical(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogCritical(C.int(cat), cstr)
}

// LogMessage (https://wiki.libsdl.org/SDL_LogMessage)
func LogMessage(cat int, pri LogPriority, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogMessage(C.int(cat), C.SDL_LogPriority(pri), cstr)
}

type LogOutputFunction func(data interface{}, cat int, pri LogPriority, message string)

type logOutputFunctionCtx struct {
	f LogOutputFunction
	d interface{}
}

// Yissakhar Z. Beck (DeedleFake)'s implementation
//
//export logOutputFunction
func logOutputFunction(data unsafe.Pointer, cat C.int, pri C.SDL_LogPriority, message *C.char) {
	ctx := (*logOutputFunctionCtx)(data)

	ctx.f(ctx.d, int(cat), LogPriority(pri), C.GoString(message))
}

var (
	logOutputFunctionCache LogOutputFunction
	logOutputDataCache     interface{}
)

// LogGetOutputFunction (https://wiki.libsdl.org/SDL_LogGetOutputFunction)
func LogGetOutputFunction() (LogOutputFunction, interface{}) {
	return logOutputFunctionCache, logOutputDataCache
}

// LogSetOutputFunction (https://wiki.libsdl.org/SDL_LogSetOutputFunction)
func LogSetOutputFunction(f LogOutputFunction, data interface{}) {
	ctx := &logOutputFunctionCtx{
		f: f,
		d: data,
	}

	C.LogSetOutputFunction(unsafe.Pointer(ctx))

	logOutputFunctionCache = f
	logOutputDataCache = data
}
