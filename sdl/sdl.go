// Package sdl is SDL2 wrapped for Go users. It enables interoperability between Go and the SDL2 library which is written in C. That means the original SDL2 installation is required for this to work. SDL2 is a cross-platform development library designed to provide low level access to audio, keyboard, mouse, joystick, and graphics hardware via OpenGL and Direct3D.
package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,9))

#if defined(WARN_OUTDATED)
#pragma message("SDL_INIT_SENSOR is not supported before SDL 2.0.9")
#endif

#define SDL_INIT_SENSOR (0x00008000u)
#endif
*/
import "C"

import (
	"runtime"
)

// These are the flags which may be passed to SDL_Init().
// (https://wiki.libsdl.org/SDL_Init)
const (
	INIT_TIMER          = C.SDL_INIT_TIMER          // timer subsystem
	INIT_AUDIO          = C.SDL_INIT_AUDIO          // audio subsystem
	INIT_VIDEO          = C.SDL_INIT_VIDEO          // video subsystem; automatically initializes the events subsystem
	INIT_JOYSTICK       = C.SDL_INIT_JOYSTICK       // joystick subsystem; automatically initializes the events subsystem
	INIT_HAPTIC         = C.SDL_INIT_HAPTIC         // haptic (force feedback) subsystem
	INIT_GAMECONTROLLER = C.SDL_INIT_GAMECONTROLLER // controller subsystem; automatically initializes the joystick subsystem
	INIT_EVENTS         = C.SDL_INIT_EVENTS         // events subsystem
	INIT_NOPARACHUTE    = C.SDL_INIT_NOPARACHUTE    // compatibility; this flag is ignored
	INIT_SENSOR         = C.SDL_INIT_SENSOR         // sensor subsystem
	INIT_EVERYTHING     = C.SDL_INIT_EVERYTHING     // all of the above subsystems
)

const (
	RELEASED = 0
	PRESSED  = 1
)

// Calls a function in the main thread. It is only properly initialized inside
// sdl.Main(..). As a default, it panics. It is used by sdl.Do(..) below.
var callInMain = func(f func()) {
	panic("sdl.Main(main func()) must be called before sdl.Do(f func())")
}

func init() {
	// Make sure the main goroutine is bound to the main thread.
	runtime.LockOSThread()
}

// Main entry point. Run this function at the beginning of main(), and pass your
// own main body to it as a function. E.g.:
//
// 	func main() {
// 		sdl.Main(func() {
// 			// Your code here....
// 			// [....]
//
// 			// Calls to SDL can be made by any goroutine, but always guarded by sdl.Do()
// 			sdl.Do(func() {
// 				sdl.Init(0)
// 			})
// 		})
// 	}
//
// Avoid calling functions like os.Exit(..) within your passed-in function since
// they don't respect deferred calls. Instead, do this:
//
// 	func main() {
// 		var exitcode int
// 		sdl.Main(func() {
// 			exitcode = run()) // assuming run has signature func() int
// 		})
// 		os.Exit(exitcode)
// 	}
func Main(main func()) {
	// Queue of functions that are thread-sensitive
	callQueue := make(chan func())

	// Properly initialize callInMain for use by sdl.Do(..)
	callInMain = func(f func()) {
		done := make(chan bool, 1)
		callQueue <- func() {
			f()
			done <- true
		}
		<-done
	}

	go func() {
		main()
		// fmt.Println("END") // to check if os.Exit(..) is called by main() above
		close(callQueue)
	}()

	for f := range callQueue {
		f()
	}
}

// Do the specified function in the main thread.
// For this function to work, you must have correctly used sdl.Main(..) in your
// main() function. Calling this function before/without sdl.Main(..) will cause
// a panic.
func Do(f func()) {
	callInMain(f)
}

// Init initialize the SDL library. This must be called before using most other SDL functions.
// (https://wiki.libsdl.org/SDL_Init)
func Init(flags uint32) error {
	if C.SDL_Init(C.Uint32(flags)) != 0 {
		return GetError()
	}
	return nil
}

// Quit cleans up all initialized subsystems. You should call it upon all exit conditions.
// (https://wiki.libsdl.org/SDL_Quit)
func Quit() {
	C.SDL_Quit()

	eventFilterCache = nil
	for k := range eventWatches {
		delete(eventWatches, k)
	}
}

// InitSubSystem initializes specific SDL subsystems.
// (https://wiki.libsdl.org/SDL_InitSubSystem)
func InitSubSystem(flags uint32) error {
	if C.SDL_InitSubSystem(C.Uint32(flags)) != 0 {
		return GetError()
	}
	return nil
}

// QuitSubSystem shuts down specific SDL subsystems.
// (https://wiki.libsdl.org/SDL_QuitSubSystem)
func QuitSubSystem(flags uint32) {
	C.SDL_QuitSubSystem(C.Uint32(flags))
}

// WasInit returns a mask of the specified subsystems which have previously been initialized.
// (https://wiki.libsdl.org/SDL_WasInit)
func WasInit(flags uint32) uint32 {
	return uint32(C.SDL_WasInit(C.Uint32(flags)))
}

// GetPlatform returns the name of the platform.
// (https://wiki.libsdl.org/SDL_GetPlatform)
func GetPlatform() string {
	return string(C.GoString(C.SDL_GetPlatform()))
}
