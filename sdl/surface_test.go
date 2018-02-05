package sdl

import (
	"testing"
)

func TestSurface(t *testing.T) {
	Do(func() {
		var window *Window
		var surface *Surface
		var image *Surface
		var err error

		Init(INIT_EVERYTHING)
		defer Quit()

		if window, err = CreateWindow("Hello!", 100, 100, 800, 600, WINDOW_SHOWN); err != nil {
			t.Error(err)
		}
		defer window.Destroy()

		if surface, err = window.GetSurface(); err != nil {
			t.Error(err)
		}

		pixels := surface.Pixels()
		for i := range pixels {
			pixels[i] = 0xFF
		}

		if err = window.UpdateSurface(); err != nil {
			t.Error(err)
		}

		Delay(50)

		if image, err = LoadBMP("../.go-sdl2-examples/assets/test.bmp"); err != nil {
			t.Error(err)
		}
		defer image.Free()

		if err = image.Blit(&Rect{0, 0, 512, 512}, surface, &Rect{0, 0, 512, 512}); err != nil {
			t.Error(err)
		}

		if err = window.UpdateSurface(); err != nil {
			t.Error(err)
		}

		Delay(50)
	})
}
