package sdl

/*
#include <SDL2/SDL.h>

static inline _SDL_Log(const char *fmt)
{
    SDL_Log(fmt);
}

static inline _SDL_LogVerbose(int category, const char *fmt)
{
    SDL_LogVerbose(category, fmt);
}

static inline _SDL_LogDebug(int category, const char *fmt)
{
    SDL_LogDebug(category, fmt);
}

static inline _SDL_LogInfo(int category, const char *fmt)
{
    SDL_LogInfo(category, fmt);
}

static inline _SDL_LogWarn(int category, const char *fmt)
{
    SDL_LogWarn(category, fmt);
}

static inline _SDL_LogError(int category, const char *fmt)
{
    SDL_LogError(category, fmt);
}

static inline _SDL_LogCritical(int category, const char *fmt)
{
    SDL_LogCritical(category, fmt);
}

static inline _SDL_LogMessage(int category, SDL_LogPriority priority, const char *fmt)
{
    SDL_LogCritical(category, fmt);
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

type LogPriority C.SDL_LogPriority

func (p LogPriority) c() C.SDL_LogPriority {
	return C.SDL_LogPriority(p)
}

func LogSetAllPriority(p LogPriority) {
	C.SDL_LogSetAllPriority(p.c())
}

func LogSetPriority(category int, p LogPriority) {
	C.SDL_LogSetPriority(C.int(category), p.c())
}

func LogGetPriority(category int) LogPriority {
	return LogPriority(C.SDL_LogGetPriority(C.int(category)))
}

func LogResetPriorities() {
	C.SDL_LogResetPriorities()
}

func Log(str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_Log(cstr)
}

func LogVerbose(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogVerbose(C.int(cat), cstr)
}

func LogDebug(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogDebug(C.int(cat), cstr)
}

func LogInfo(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogInfo(C.int(cat), cstr)
}

func LogWarn(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogWarn(C.int(cat), cstr)
}

func LogError(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogError(C.int(cat), cstr)
}

func LogCritical(cat int, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogCritical(C.int(cat), cstr)
}

func LogMessage(cat int, pri LogPriority, str string, args ...interface{}) {
	str = fmt.Sprintf(str, args...)

	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	C._SDL_LogMessage(C.int(cat), C.SDL_LogPriority(pri), cstr)
}
