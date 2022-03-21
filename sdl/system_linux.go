// +build !android

package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,9))

#if defined(WARN_OUTDATED)
#pragma message("SDL_LinuxSetThreadPriority is not supported before SDL 2.0.9")
#endif

static int SDL_LinuxSetThreadPriority(Sint64 threadID, int priority)
{
	return -1;
}

#endif

#if !(SDL_VERSION_ATLEAST(2,0,18))

#if defined(WARN_OUTDATED)
#pragma message("SDL_LinuxSetThreadPriorityAndPolicy is not supported before SDL 2.0.18")
#endif

static int SDL_LinuxSetThreadPriorityAndPolicy(Sint64 threadID, int sdlPriority, int schedPolicy)
{
	return -1;
}

#endif
*/
import "C"

// LinuxSetThreadPriority sets the UNIX nice value for a thread.
//
// This uses setpriority() if possible, and RealtimeKit if available.
//
// (https://wiki.libsdl.org/SDL_LinuxSetThreadPriority)
func LinuxSetThreadPriority(threadID int64, priority int) (err error) {
	_threadID := C.Sint64(threadID)
	_priority := C.int(priority)
	return errorFromInt(int(C.SDL_LinuxSetThreadPriority(_threadID, _priority)))
}

// LinuxSetThreadPriority sets the priority (not nice level) and scheduling policy for a thread.
//
// This uses setpriority() if possible, and RealtimeKit if available.
//
// (https://wiki.libsdl.org/SDL_LinuxSetThreadPriorityAndPolicy)
func LinuxSetThreadPriorityAndPolicy(threadID int64, sdlPriority, schedPolicy int) (err error) {
	_threadID := C.Sint64(threadID)
	_sdlPriority := C.int(sdlPriority)
	_schedPolicy := C.int(schedPolicy)
	return errorFromInt(int(C.SDL_LinuxSetThreadPriorityAndPolicy(_threadID, _sdlPriority, _schedPolicy)))
}
