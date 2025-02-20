package sdl

import (
	"errors"
	"image/color"
	"reflect"
	"testing"
	"unsafe"
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
	srcFormat := uint32(PIXELFORMAT_RGBA32)
	s, err := CreateRGBSurfaceWithFormat(0, 50, 50, 0, srcFormat)
	if err != nil {
		t.Fatalf("unable to create surface: %s", err)
	}
	defer s.Free()

	r, err := CreateSoftwareRenderer(s)
	if err != nil {
		t.Fatalf("unable to create renderer: %s", err)
	}
	defer r.Destroy()

	type quadrant struct {
		c      color.RGBA
		r      *Rect
		pitch  int
		pixels []byte
	}

	newQuad := func(c color.RGBA, r *Rect) quadrant {
		s, err := CreateRGBSurfaceWithFormat(0, int32(r.W), int32(r.H), 8, srcFormat)
		if err != nil {
			t.Fatalf("unable to create rgba surface: %s", err)
		}
		defer s.Free()

		if err := s.FillRect(nil, MapRGBA(s.Format, c.R, c.G, c.B, c.A)); err != nil {
			t.Fatalf("unable to fill rect: %s", err)
		}

		src := s.Pixels()
		q := quadrant{
			c:      c,
			r:      r,
			pitch:  int(s.Pitch),
			pixels: make([]byte, len(src)),
		}
		copy(q.pixels, src)
		return q
	}

	conv := func(q quadrant, format uint32) []byte {
		ret := make([]byte, int(q.r.W*q.r.H)*BytesPerPixel(format))
		if err := ConvertPixels(
			q.r.W,
			q.r.H,
			srcFormat,
			unsafe.Pointer(&q.pixels[0]),
			q.pitch,
			format,
			unsafe.Pointer(&ret[0]),
			int(q.r.W)*BytesPerPixel(format),
		); err != nil {
			t.Fatalf("unable to convert pixels to format %s: %s", GetPixelFormatName(uint(format)), err)
		}

		return ret
	}

	mergeH := func(q1, q2 quadrant, format uint32) []byte {
		a := conv(q1, format)
		b := conv(q2, format)
		pitch := int(q1.r.W) * BytesPerPixel(format)

		ret := make([]byte, 0, len(a)+len(b))
		rows := len(a) / pitch
		for i := 0; i < rows; i++ {
			ret = append(ret, a[:pitch]...)
			ret = append(ret, b[:pitch]...)
			a = a[pitch:]
			b = b[pitch:]
		}
		return ret
	}

	mergeV := func(q1, q2 quadrant, format uint32) []byte {
		a := conv(q1, format)
		b := conv(q2, format)
		return append(a, b...)
	}

	merge := func(q1, q2, q3, q4 quadrant, format uint32) []byte {
		a := mergeH(q1, q2, format)
		b := mergeH(q3, q4, format)

		return append(a, b...)
	}

	q1 := newQuad(color.RGBA{151, 237, 158, 255}, &Rect{X: 0, Y: 0, W: 25, H: 25})
	q2 := newQuad(color.RGBA{151, 213, 237, 255}, &Rect{X: 25, Y: 0, W: 25, H: 25})
	q3 := newQuad(color.RGBA{224, 151, 237, 255}, &Rect{X: 0, Y: 25, W: 25, H: 25})
	q4 := newQuad(color.RGBA{237, 151, 151, 255}, &Rect{X: 25, Y: 25, W: 25, H: 25})

	type test struct {
		name string
		rect *Rect
		fmt  uint32
		want []byte
	}

	block := func(format uint32) []test {
		name := GetPixelFormatName(uint(format))
		return []test{
			{"q1 " + name, q1.r, format, conv(q1, format)},
			{"q2 " + name, q2.r, format, conv(q2, format)},
			{"q3 " + name, q3.r, format, conv(q3, format)},
			{"q4 " + name, q4.r, format, conv(q4, format)},
			{"q1 + q2 " + name, &Rect{X: 0, Y: 0, W: 50, H: 25}, format, mergeH(q1, q2, format)},
			{"q3 + q4 " + name, &Rect{X: 0, Y: 25, W: 50, H: 25}, format, mergeH(q3, q4, format)},
			{"q1 + q3 " + name, &Rect{X: 0, Y: 0, W: 25, H: 50}, format, mergeV(q1, q3, format)},
			{"q2 + q4 " + name, &Rect{X: 25, Y: 0, W: 25, H: 50}, format, mergeV(q2, q4, format)},
			{"full " + name, &Rect{X: 0, Y: 0, W: 50, H: 50}, format, merge(q1, q2, q3, q4, format)},
		}
	}

	for _, q := range []quadrant{q1, q2, q3, q4} {
		if err := r.SetDrawColor(q.c.R, q.c.G, q.c.B, q.c.A); err != nil {
			t.Fatalf("unable to set draw color: %s", err)
		}
		if err := r.FillRect(q.r); err != nil {
			t.Fatalf("unable to fill rect: %s", err)
		}
	}

	var tests []test
	// tests = append(tests, block(PIXELFORMAT_UNKNOWN)...)
	// tests = append(tests, block(PIXELFORMAT_INDEX1LSB)...)
	// tests = append(tests, block(PIXELFORMAT_INDEX1MSB)...)
	// tests = append(tests, block(PIXELFORMAT_INDEX4LSB)...)
	// tests = append(tests, block(PIXELFORMAT_INDEX4MSB)...)
	// tests = append(tests, block(PIXELFORMAT_INDEX8)...)

	tests = append(tests, block(PIXELFORMAT_RGB332)...)         // ok
	tests = append(tests, block(PIXELFORMAT_RGB444)...)         // ok
	tests = append(tests, block(PIXELFORMAT_RGB555)...)         // ok
	tests = append(tests, block(PIXELFORMAT_BGR555)...)         // ok
	tests = append(tests, block(PIXELFORMAT_ARGB4444)...)       // ok
	tests = append(tests, block(PIXELFORMAT_RGBA4444)...)       // ok
	tests = append(tests, block(PIXELFORMAT_ABGR4444)...)       // ok
	tests = append(tests, block(PIXELFORMAT_BGRA4444)...)       // ok
	tests = append(tests, block(PIXELFORMAT_ARGB1555)...)       // ok
	tests = append(tests, block(PIXELFORMAT_RGBA5551)...)       // ok
	tests = append(tests, block(PIXELFORMAT_ABGR1555)...)       // ok
	tests = append(tests, block(PIXELFORMAT_BGRA5551)...)       // ok
	tests = append(tests, block(PIXELFORMAT_RGB565)...)         // ok
	tests = append(tests, block(PIXELFORMAT_BGR565)...)         // ok
	tests = append(tests, block(uint32(PIXELFORMAT_RGB24))...)  // ok
	tests = append(tests, block(uint32(PIXELFORMAT_BGR24))...)  // ok
	tests = append(tests, block(uint32(PIXELFORMAT_RGBA32))...) // ok
	tests = append(tests, block(uint32(PIXELFORMAT_ARGB32))...) // ok
	tests = append(tests, block(uint32(PIXELFORMAT_BGRA32))...) // ok
	tests = append(tests, block(uint32(PIXELFORMAT_ABGR32))...) // ok
	tests = append(tests, block(PIXELFORMAT_RGB888)...)         // ok
	tests = append(tests, block(PIXELFORMAT_RGBX8888)...)       // ok
	tests = append(tests, block(PIXELFORMAT_BGR888)...)         // ok
	tests = append(tests, block(PIXELFORMAT_BGRX8888)...)       // ok
	tests = append(tests, block(PIXELFORMAT_ARGB8888)...)       // ok
	tests = append(tests, block(PIXELFORMAT_RGBA8888)...)       // ok
	tests = append(tests, block(PIXELFORMAT_ABGR8888)...)       // ok
	tests = append(tests, block(PIXELFORMAT_BGRA8888)...)       // ok
	// tests = append(tests, block(PIXELFORMAT_ARGB2101010)...) // fail

	// tests = append(tests, block(PIXELFORMAT_YV12)...)
	// tests = append(tests, block(PIXELFORMAT_IYUV)...)
	// tests = append(tests, block(PIXELFORMAT_YUY2)...)
	// tests = append(tests, block(PIXELFORMAT_UYVY)...)
	// tests = append(tests, block(PIXELFORMAT_YVYU)...)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := r.GetPixels(tt.rect, tt.fmt)
			if err != nil {
				t.Fatalf("unable to get pixels: %s", err)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("want %v len(%d), got %v len(%d)", tt.want, len(tt.want), got, len(got))
			}
		})
	}
}
