package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// SharedObject is a pointer to the object handle.
type SharedObject uintptr

// LoadObject dynamically loads a shared object and returns a pointer to the object handle.
// (https://wiki.libsdl.org/SDL_LoadObject)
func LoadObject(sofile string) SharedObject {
	_sofile := C.CString(sofile)
	defer C.free(unsafe.Pointer(_sofile))
	return (SharedObject)(C.SDL_LoadObject(_sofile))
}

// LoadFunction returns a pointer to the named function from the shared object.
// (https://wiki.libsdl.org/SDL_LoadFunction)
func (handle SharedObject) LoadFunction(name string) unsafe.Pointer {
	_name := C.CString(name)
	defer C.free(unsafe.Pointer(_name))
	return (unsafe.Pointer)(C.SDL_LoadFunction((unsafe.Pointer)(handle), _name))
}

// Unload unloads a shared object from memory.
// (https://wiki.libsdl.org/SDL_UnloadObject)
func (handle SharedObject) Unload() {
	C.SDL_UnloadObject((unsafe.Pointer)(handle))
}
