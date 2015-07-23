package sdl

import (
	"testing"
)

func TestRectEmpty(t *testing.T) {
	var tests = []struct {
		want bool
		r    *Rect
	}{
		{want: true, r: &Rect{}},
		{want: true, r: &Rect{W: 1}},
		{want: true, r: &Rect{H: 1}},
		{want: true, r: nil},
		{want: false, r: &Rect{W: 1, H: 1}},
	}

	for _, test := range tests {
		if got := test.r.Empty(); got != test.want {
			t.Errorf("%#v.Empty() = %v - want %v", test.r, got, test.want)
		}
	}
}

func TestIntersectLines(t *testing.T) {

	var tests = []struct {
		want bool
		in   struct{ x1, y1, x2, y2 int }
		out  struct{ x1, y1, x2, y2 int }
		r    *Rect
	}{
		{want: false,
			in:  struct{ x1, y1, x2, y2 int }{15, 15, 25, 25},
			out: struct{ x1, y1, x2, y2 int }{15, 15, 25, 25},
			r:   &Rect{X: 0, Y: 0, W: 10, H: 10}},
		{want: true,
			in:  struct{ x1, y1, x2, y2 int }{-1, -1, 11, 11},
			out: struct{ x1, y1, x2, y2 int }{0, 0, 9, 9},
			r:   &Rect{X: 0, Y: 0, W: 10, H: 10}}}

	for _, test := range tests {
		original := test.in
		if result := test.r.IntersectLine(&test.in.x1, &test.in.y1, &test.in.x2, &test.in.y2); result != test.want {
			t.Errorf("%v.IntersectLines(%v, %v, %v, %v) = %v - want %v",
				test.r, original.x1, original.y1, original.x2, original.y2, result, test.want)
		}

		if test.in != test.out {
			t.Errorf("%v.IntersectLines(%v, %v, %v, %v) gives %v - want %v",
				test.r, original.x1, original.y1, original.x2, original.y2, test.in, test.out)
		}
	}
}
