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
