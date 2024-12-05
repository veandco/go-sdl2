// +build !nomidi,!android

package mix

//#include <stdlib.h>
//#include "sdl_mixer_wrapper.h"
//extern int callEachSoundFont(char* str, void* udata);
//
//#if !(SDL_MIXER_VERSION_ATLEAST(2,6,0))
//
//#if defined(WARN_OUTDATED)
//#pragma message("Mix_SetTimidityCfg is not supported before SDL 2.6.0")
//#endif
//
//int Mix_SetTimidityCfg(const char *path)
//{
//	return 0;
//}
//
//#endif
import "C"
import "unsafe"

var eachSoundFontFunc func(string) int

//export callEachSoundFont
func callEachSoundFont(str *C.char, udata unsafe.Pointer) C.int {
	return C.int(eachSoundFontFunc(C.GoString(str)))
}

// EachSoundFont iterates over SoundFonts paths to use by supported MIDI backends.
// (https://wiki.libsdl.org/SDL2_mixer/Mix_EachSoundFont)
func EachSoundFont(function func(string) int) int {
	eachSoundFontFunc = function
	return int(C.Mix_EachSoundFont((*[0]byte)(C.callEachSoundFont), nil))
}

// SetSoundFonts sets SoundFonts paths to use by supported MIDI backends.
// (https://wiki.libsdl.org/SDL2_mixer/Mix_SetSoundFonts)
func SetSoundFonts(paths string) bool {
	_paths := C.CString(paths)
	defer C.free(unsafe.Pointer(_paths))
	return int(C.Mix_SetSoundFonts(_paths)) == 0
}

// GetSoundFonts returns SoundFonts paths to use by supported MIDI backends.
// (https://wiki.libsdl.org/SDL2_mixer/Mix_GetSoundFonts)
func GetSoundFonts() string {
	return (string)(C.GoString(C.Mix_GetSoundFonts()))
}

// Set full path of the Timidity config file. Returns true if successful, false on error. This is obviously only useful if SDL_mixer is using Timidity internally to play MIDI files.
// (https://wiki.libsdl.org/SDL2_mixer/Mix_SetTimidityCfg)
func SetTimidityCfg(path string) bool {
	_path := C.CString(path)
	defer C.free(unsafe.Pointer(_path))
	return int(C.Mix_SetTimidityCfg(_path)) == 0
}
