// +build go1.4

package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// External storage states. See the official Android developer guide for more information.
// (http://developer.android.com/guide/topics/data/data-storage.html)
const (
	ANDROID_EXTERNAL_STORAGE_READ  = C.SDL_ANDROID_EXTERNAL_STORAGE_READ
	ANDROID_EXTERNAL_STORAGE_WRITE = C.SDL_ANDROID_EXTERNAL_STORAGE_WRITE
)

// AndroidGetInternalStoragePath returns the path used for internal storage for this application.
// (https://wiki.libsdl.org/SDL_AndroidGetInternalStoragePath)
func AndroidGetInternalStoragePath() string {
	return C.GoString(C.SDL_AndroidGetInternalStoragePath())
}

// AndroidGetExternalStoragePath returns the path used for external storage for this application.
// (https://wiki.libsdl.org/SDL_AndroidGetExternalStoragePath)
func AndroidGetExternalStoragePath() string {
	return C.GoString(C.SDL_AndroidGetExternalStoragePath())
}

// AndroidGetExternalStorageState returns the current state of external storage.
// (https://wiki.libsdl.org/SDL_AndroidGetExternalStorageState)
func AndroidGetExternalStorageState() int {
	return int(C.SDL_AndroidGetExternalStorageState())
}

// AndroidGetJNIEnv returns the Java native interface object (JNIEnv) of the current thread on Android builds.
// (https://wiki.libsdl.org/SDL_AndroidGetJNIEnv)
func AndroidGetJNIEnv() unsafe.Pointer {
	return unsafe.Pointer(C.SDL_AndroidGetJNIEnv())
}

// AndroidGetActivity returns the Java instance of the activity class in an Android application.
// (https://wiki.libsdl.org/SDL_AndroidGetActivity)
func AndroidGetActivity() unsafe.Pointer {
	return unsafe.Pointer(C.SDL_AndroidGetActivity())
}
