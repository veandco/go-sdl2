// +build !nomidi,!android

package mix

//#cgo windows LDFLAGS: -lSDL2 -lSDL2_mixer
//#cgo linux freebsd darwin pkg-config: sdl2
//#cgo linux freebsd darwin LDFLAGS: -lSDL2_mixer
//#include <stdlib.h>
//#include "sdl_mixer_wrapper.h"
//extern int callEachSoundFont(char* str, void* udata);
import "C"
import "unsafe"

var eachSoundFontFunc func(string) int

//export callEachSoundFont
func callEachSoundFont(str *C.char, udata unsafe.Pointer) C.int {
	return C.int(eachSoundFontFunc(C.GoString(str)))
}

// undocumented
func EachSoundFont(function func(string) int) int {
	eachSoundFontFunc = function
	return int(C.Mix_EachSoundFont((*[0]byte)(C.callEachSoundFont), nil))
}

// undocumented
func SetSoundFonts(paths string) bool {
	_paths := C.CString(paths)
	defer C.free(unsafe.Pointer(_paths))
	return int(C.Mix_SetSoundFonts(_paths)) == 0
}

// undocumented
func GetSoundFonts() string {
	return (string)(C.GoString(C.Mix_GetSoundFonts()))
}
