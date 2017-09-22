package gfx

import (
	"testing"
)

func TestMin(t *testing.T) {
	if a := []int{4, 2, 3, 5, 1, 6, 5, 7}; min(a...) != 1 {
		t.Fail()
	}

	if a := []int{5, 6, -2, 1, 8, 1, 4, 3}; min(a...) != -2 {
		t.Fail()
	}
}

func TestFPSmanager(t *testing.T) {
	fpsManager := FPSmanager{}
	InitFramerate(&fpsManager)

	if !SetFramerate(&fpsManager, 60) {
		t.Errorf("gfx.SetFramerate failed")
	}

	if fps, ok := GetFramerate(&fpsManager); !ok || fps != 60 {
		t.Errorf("gfx.GetFramerate() = (%v, %v) - want (60, true)", fps, ok)
	}

	// frame count is initialized to 0
	if count, ok := GetFramecount(&fpsManager); !ok || count != 0 {
		t.Errorf("gfx.GetFramecount() = (%v, %v) - want (0, true)", count, ok)
	}

	// count one frame
	FramerateDelay(&fpsManager)

	// frame count should now be 1
	if count, ok := GetFramecount(&fpsManager); !ok || count != 1 {
		t.Errorf("gfx.Framecount() = (%v, %v) - want (1, true)", count, ok)
	}
}
