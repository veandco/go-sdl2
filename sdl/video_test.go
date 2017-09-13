package sdl

import (
	"testing"
)

func TestWindow(t *testing.T) {
	Do(func() {
		var window *Window
		var err error

		Init(INIT_EVERYTHING)
		defer Quit()

		if window, err = CreateWindow("Hello!", 100, 100, 800, 600, WINDOW_SHOWN); err != nil {
			t.Error(err)
		}
		defer window.Destroy()

		if title := window.GetTitle(); title != "Hello!" {
			t.Error("Window title is incorrect!")
		}

		if x, y := window.GetPosition(); x != 100 || y != 100 {
			t.Error("Window position is incorrect!")
		}

		if w, h := window.GetSize(); w != 800 || h != 600 {
			t.Error("Window size is incorrect!")
		}

		Delay(100)
	})
}
