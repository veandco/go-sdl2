package sdl

import (
	"bytes"
	"errors"
	"image/color"
	"reflect"
	"testing"
)

func TestRenderIntegerScale(t *testing.T) {
	_, r, err := CreateWindowAndRenderer(50, 50, 0)
	if err != nil {
		t.Fatalf("unable to create renderer: %s", err)
		return
	}

	errUnsupported := errors.New("That operation is not supported")

	if !VERSION_ATLEAST(2, 0, 5) {
		t.Run("GetIntegerScale on versions < 2.0.5 should return unsupported", func(t *testing.T) {
			_, got := r.GetIntegerScale()
			if got.Error() != errUnsupported.Error() {
				t.Fatalf("expected call to GetIntegerScale to fail with %s, got %s", errUnsupported, got)
			}
		})
		t.Run("SetIntegerScale on versions < 2.0.5 should return unsupported", func(t *testing.T) {
			got := r.SetIntegerScale(false)
			if got.Error() != errUnsupported.Error() {
				t.Fatalf("expected call to SetIntegerScale to fail with %s, got %s", errUnsupported, got)
			}
		})

		return
	}

	set := func(v bool, t *testing.T) {
		if err := r.SetIntegerScale(v); err != nil {
			t.Fatalf("unable to set scale %v: %s", v, err)
		}
	}
	check := func(want bool, t *testing.T) {
		got, err := r.GetIntegerScale()
		if err != nil {
			t.Fatalf("unable to get scale: %s", err)
		}

		if got != want {
			t.Fatalf("wanted renderer scale to be %v", want)
		}
	}

	t.Run("should start out as false", func(t *testing.T) {
		check(false, t)
	})

	t.Run("if set to true should return true", func(t *testing.T) {
		set(true, t)
		check(true, t)
	})

	t.Run("if set to false should return false", func(t *testing.T) {
		set(true, t)
		check(true, t)

		set(false, t)
		check(false, t)
	})

	t.Run("should handle SDL_FALSE return on GetIntegerScale", func(t *testing.T) {
		_, dummy, err := CreateWindowAndRenderer(50, 50, 0)
		if err != nil {
			t.Fatalf("unable to create renderer: %s", err)
			return
		}

		if err := dummy.SetIntegerScale(true); err != nil {
			t.Fatalf("expected call to SetIntegerScale to succeed, got %s", err)
		}
		if err := dummy.Destroy(); err != nil {
			t.Fatalf("expected call to Destroy to succeed, got %s", err)
		}

		v, err := dummy.GetIntegerScale()
		if err == nil {
			t.Fatalf("expected call to GetIntegerScale on ivalid renderer to fail")
		}
		if v != false {
			t.Fatalf("wanted renderer scale to be %v", false)
		}
	})
}

func TestGetPixels(t *testing.T) {
	repeat := bytes.Repeat
	line := func(a []byte, b []byte, len int) []byte {
		return append(repeat(a, len), repeat(b, len)...)
	}
	concat := func(a []byte, b []byte) []byte {
		return append(a, b...)
	}
	rgba := func(c color.RGBA) []byte {
		return []byte{c.R, c.G, c.B, c.A}
	}
	bgra := func(c color.RGBA) []byte {
		return []byte{c.A, c.R, c.G, c.B}
	}
	drawRect := func(renderer *Renderer, rect *Rect, c color.RGBA) {
		if err := renderer.SetDrawColor(c.R, c.G, c.B, c.A); err != nil {
			t.Fatalf("unable to set draw color: %s", err)
		}
		if err := renderer.FillRect(rect); err != nil {
			t.Fatalf("unable to fill rect: %s", err)
		}
	}

	_, r, err := CreateWindowAndRenderer(50, 50, 0)
	if err != nil {
		t.Fatalf("unable to create renderer: %s", err)
	}

	c1, r1 := color.RGBA{151, 237, 158, 255}, &Rect{X: 0, Y: 0, W: 25, H: 25}
	c2, r2 := color.RGBA{151, 213, 237, 255}, &Rect{X: 25, Y: 0, W: 25, H: 25}
	c3, r3 := color.RGBA{224, 151, 237, 255}, &Rect{X: 0, Y: 25, W: 25, H: 25}
	c4, r4 := color.RGBA{237, 151, 151, 255}, &Rect{X: 25, Y: 25, W: 25, H: 25}

	drawRect(r, r1, c1)
	drawRect(r, r2, c2)
	drawRect(r, r3, c3)
	drawRect(r, r4, c4)

	tests := []struct {
		name string
		rect *Rect
		fmt  uint32
		want []byte
	}{
		{"r1 ABGR8888", r1, PIXELFORMAT_ABGR8888, repeat(rgba(c1), 25*25)},
		{"r2 ABGR8888", r2, PIXELFORMAT_ABGR8888, repeat(rgba(c2), 25*25)},
		{"r3 ABGR8888", r3, PIXELFORMAT_ABGR8888, repeat(rgba(c3), 25*25)},
		{"r4 ABGR8888", r4, PIXELFORMAT_ABGR8888, repeat(rgba(c4), 25*25)},

		{"r1 + r2 ABGR8888", &Rect{X: 0, Y: 0, W: 50, H: 25}, PIXELFORMAT_ABGR8888, repeat(line(rgba(c1), rgba(c2), 25), 25)},
		{"r3 + r4 ABGR8888", &Rect{X: 0, Y: 25, W: 50, H: 25}, PIXELFORMAT_ABGR8888, repeat(line(rgba(c3), rgba(c4), 25), 25)},
		{"r1 + r3 ABGR8888", &Rect{X: 0, Y: 0, W: 25, H: 50}, PIXELFORMAT_ABGR8888, concat(repeat(rgba(c1), 25*25), repeat(rgba(c3), 25*25))},
		{"r2 + r4 ABGR8888", &Rect{X: 25, Y: 0, W: 25, H: 50}, PIXELFORMAT_ABGR8888, concat(repeat(rgba(c2), 25*25), repeat(rgba(c4), 25*25))},

		{"full ABGR8888", &Rect{X: 0, Y: 0, W: 50, H: 50}, PIXELFORMAT_ABGR8888, concat(repeat(line(rgba(c1), rgba(c2), 25), 25), repeat(line(rgba(c3), rgba(c4), 25), 25))},

		{"r1 BGRA8888", r1, PIXELFORMAT_BGRA8888, repeat(bgra(c1), 25*25)},
		{"r2 BGRA8888", r2, PIXELFORMAT_BGRA8888, repeat(bgra(c2), 25*25)},
		{"r3 BGRA8888", r3, PIXELFORMAT_BGRA8888, repeat(bgra(c3), 25*25)},
		{"r4 BGRA8888", r4, PIXELFORMAT_BGRA8888, repeat(bgra(c4), 25*25)},

		{"r1 + r2 ABGR8888", &Rect{X: 0, Y: 0, W: 50, H: 25}, PIXELFORMAT_BGRA8888, repeat(line(bgra(c1), bgra(c2), 25), 25)},
		{"r3 + r4 ABGR8888", &Rect{X: 0, Y: 25, W: 50, H: 25}, PIXELFORMAT_BGRA8888, repeat(line(bgra(c3), bgra(c4), 25), 25)},
		{"r1 + r3 ABGR8888", &Rect{X: 0, Y: 0, W: 25, H: 50}, PIXELFORMAT_BGRA8888, concat(repeat(bgra(c1), 25*25), repeat(bgra(c3), 25*25))},
		{"r2 + r4 ABGR8888", &Rect{X: 25, Y: 0, W: 25, H: 50}, PIXELFORMAT_BGRA8888, concat(repeat(bgra(c2), 25*25), repeat(bgra(c4), 25*25))},

		{"full ABGR8888", &Rect{X: 0, Y: 0, W: 50, H: 50}, PIXELFORMAT_BGRA8888, concat(repeat(line(bgra(c1), bgra(c2), 25), 25), repeat(line(bgra(c3), bgra(c4), 25), 25))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetPixels(tt.rect, tt.fmt)
			if err != nil {
				t.Fatalf("unable to get pixels for r1: %s", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("wanted %v got: %v", tt.want, got)
			}
		})
	}
}
