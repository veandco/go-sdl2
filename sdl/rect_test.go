package sdl

import (
	"testing"
)

func TestRectHasIntersection(t *testing.T) {
	var tests = []struct {
		want bool
		rA   *Rect
		rB   *Rect
	}{
		{want: true, rA: &Rect{0, 0, 10, 10}, rB: &Rect{6, 7, 6, 7}},
		{want: false, rA: &Rect{0, 0, 10, 10}, rB: &Rect{20, 20, 5, 5}},
		{want: false, rA: &Rect{0, 0, 10, 10}, rB: &Rect{10, 0, 5, 5}},
	}

	for _, test := range tests {
		if result := test.rA.HasIntersection(test.rB); result != test.want {
			t.Errorf("%v.HasIntersection(%v) = %v - want %v", test.rA, test.rB, result, test.want)
		}
		if result := test.rB.HasIntersection(test.rA); result != test.want {
			t.Errorf("%v.HasIntersection(%v) = %v - want %v", test.rB, test.rA, result, test.want)
		}
	}
}

func TestRectIntersect(t *testing.T) {
	var tests = []struct {
		want   bool
		rA     *Rect
		rB     *Rect
		result Rect
	}{
		{want: true, rA: &Rect{0, 0, 10, 10}, rB: &Rect{6, 7, 6, 7}, result: Rect{6, 7, 4, 3}},
		{want: false, rA: &Rect{0, 0, 10, 10}, rB: &Rect{20, 20, 5, 5}, result: Rect{20, 20, -10, -10}},
		{want: false, rA: &Rect{0, 0, 10, 10}, rB: &Rect{10, 0, 5, 5}, result: Rect{10, 0, 0, 5}},
	}

	for _, test := range tests {
		result, intersects := test.rA.Intersect(test.rB)
		if intersects != test.want || result != test.result {
			t.Errorf("%v.Intersects(%v) = %v, %v - want %v, %v", test.rA, test.rB, intersects, result, test.want, test.result)
		}
		result, intersects = test.rB.Intersect(test.rA)
		if intersects != test.want || result != test.result {
			t.Errorf("%v.Intersects(%v) = %v, %v - want %v, %v", test.rB, test.rA, intersects, result, test.want, test.result)
		}
	}
}

func TestRectUnion(t *testing.T) {
	var tests = []struct {
		rA     *Rect
		rB     *Rect
		result Rect
	}{
		{rA: &Rect{0, 0, 10, 10}, rB: &Rect{6, 7, 6, 7}, result: Rect{0, 0, 12, 14}},
		{rA: &Rect{0, 0, 10, 10}, rB: &Rect{20, 20, 5, 5}, result: Rect{0, 0, 25, 25}},
		{rA: &Rect{0, 0, 10, 10}, rB: &Rect{10, 0, 5, 5}, result: Rect{0, 0, 15, 10}},
	}

	for _, test := range tests {
		result := test.rA.Union(test.rB)
		if result != test.result {
			t.Errorf("%v.Union(%v) = %v - want %v", test.rA, test.rB, result, test.result)
		}
		result = test.rB.Union(test.rA)
		if result != test.result {
			t.Errorf("%v.Union(%v) = %v - want %v", test.rB, test.rA, result, test.result)
		}
	}
}

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

func TestEnclosurePoints(t *testing.T) {
	var tests = []struct {
		want      bool
		points    []Point
		clip      *Rect
		enclosure Rect
	}{
		{want: false, points: make([]Point, 0), clip: nil, enclosure: Rect{}},
		{want: true, points: []Point{{5, 5}, {3, 7}, {4, 2}}, clip: nil, enclosure: Rect{3, 2, 3, 6}},
		{want: true, points: []Point{{5, 5}, {3, 7}, {4, 2}}, clip: &Rect{0, 0, 6, 6}, enclosure: Rect{4, 2, 2, 4}},
	}

	for _, test := range tests {
		enclosure, result := EnclosePoints(test.points, test.clip)
		if enclosure != test.enclosure || result != test.want {
			t.Errorf("EnclosurePoints(%v, %v) = %v, %v - want %v, %v",
				test.points, test.clip, enclosure, result, test.enclosure, test.want)
		}
	}
}

func TestIntersectLines(t *testing.T) {

	var tests = []struct {
		want bool
		in   struct{ x1, y1, x2, y2 int32 }
		out  struct{ x1, y1, x2, y2 int32 }
		r    *Rect
	}{
		{want: false,
			in:  struct{ x1, y1, x2, y2 int32 }{15, 15, 25, 25},
			out: struct{ x1, y1, x2, y2 int32 }{15, 15, 25, 25},
			r:   &Rect{X: 0, Y: 0, W: 10, H: 10}},
		{want: true,
			in:  struct{ x1, y1, x2, y2 int32 }{-1, -1, 11, 11},
			out: struct{ x1, y1, x2, y2 int32 }{0, 0, 9, 9},
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
