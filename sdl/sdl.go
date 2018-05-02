// Package sdl is SDL2 wrapped for Go users. It enables interoperability between Go and the SDL2 library which is written in C. That means the original SDL2 installation is required for this to work. SDL2 is a cross-platform development library designed to provide low level access to audio, keyboard, mouse, joystick, and graphics hardware via OpenGL and Direct3D.
package sdl

//#include "sdl_wrapper.h"
import "C"

import (
	"runtime"
)

// These are the flags which may be passed to SDL_Init().
// (https://wiki.libsdl.org/SDL_Init)
const (
	INIT_TIMER          = 0x00000001 // timer subsystem
	INIT_AUDIO          = 0x00000010 // audio subsystem
	INIT_VIDEO          = 0x00000020 // video subsystem; automatically initializes the events subsystem
	INIT_JOYSTICK       = 0x00000200 // joystick subsystem; automatically initializes the events subsystem
	INIT_HAPTIC         = 0x00001000 // haptic (force feedback) subsystem
	INIT_GAMECONTROLLER = 0x00002000 // controller subsystem; automatically initializes the joystick subsystem
	INIT_EVENTS         = 0x00004000 // events subsystem
	INIT_NOPARACHUTE    = 0x00100000 // compatibility; this flag is ignored

	INIT_EVERYTHING = INIT_TIMER | INIT_AUDIO | INIT_VIDEO | INIT_EVENTS | INIT_JOYSTICK | INIT_HAPTIC | INIT_GAMECONTROLLER // all of the above subsystems
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

// errorFromInt returns GetError() if passed negative value, otherwise it returns nil.
func errorFromInt(code int) error {
	if code < 0 {
		return GetError()
	}
	return nil
}
