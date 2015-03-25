package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

type SharedObject uintptr

// LoadObject (https://wiki.libsdl.org/SDL_LoadObject)
func LoadObject(sofile string) SharedObject {
	_sofile := C.CString(sofile)
	defer C.free(unsafe.Pointer(_sofile))
	return (SharedObject)(C.SDL_LoadObject(_sofile))
}

// LoadFunction (https://wiki.libsdl.org/SDL_LoadFunction)
func (handle SharedObject) LoadFunction(name string) unsafe.Pointer {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	return (unsafe.Pointer)(C.SDL_LoadFunction((unsafe.Pointer)(handle), _name))
}

// UnloadObject (https://wiki.libsdl.org/SDL_UnloadObject)
func (handle SharedObject) Unload() {
	C.SDL_UnloadObject((unsafe.Pointer)(handle))
}
