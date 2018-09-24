package sdl

import (
	"testing"
)

func TestGetBasePath(t *testing.T) {
	if !VERSION_ATLEAST(2, 0, 1) {
		return
	}

	Do(func() {
		Init(INIT_EVERYTHING)
		path := GetBasePath()
		// pathetic, but a comparison to runtime.Caller() does not help here,
		// since this would return go's testing package
		if path == "" {
			t.Error("GetBasePath() did not return anything")
		}
		Quit()
	})
}

// Do not test GetPrefPath(), since SDL_GetPrefPath() also
// *creates* the necessary directories, polluting the user's
// system with unnecessary stuff.
// func TestGetPrefPath(t *testing.T) {
// 	path := GetPrefPath("go-sdl2", "test")
// 	t.Error(path)
// }
