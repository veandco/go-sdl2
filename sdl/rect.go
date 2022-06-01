package sdl

// #include "sdl_wrapper.h"
import "C"
import (
	"math"
	"unsafe"
)

// Point defines a two dimensional point.
// (https://wiki.libsdl.org/SDL_Point)
type Point struct {
	X int32 // the x coordinate of the point
	Y int32 // the y coordinate of the point
}

// Rect contains the definition of a rectangle, with the origin at the upper left.
// (https://wiki.libsdl.org/SDL_Rect)
type Rect struct {
	X int32 // the x location of the rectangle's upper left corner
	Y int32 // the y location of the rectangle's upper left corner
	W int32 // the width of the rectangle
	H int32 // the height of the rectangle
}

// FPoint defines a two dimensional point.
// TODO: (https://wiki.libsdl.org/SDL_FPoint)
type FPoint struct {
	X float32 // the x coordinate of the point
	Y float32 // the y coordinate of the point
}

// FRect contains the definition of a rectangle, with the origin at the upper left.
// TODO: (https://wiki.libsdl.org/SDL_FRect)
type FRect struct {
	X float32 // the x location of the rectangle's upper left corner
	Y float32 // the y location of the rectangle's upper left corner
	W float32 // the width of the rectangle
	H float32 // the height of the rectangle
}

func (p *Point) cptr() *C.SDL_Point {
	return (*C.SDL_Point)(unsafe.Pointer(p))
}

func (a *Rect) cptr() *C.SDL_Rect {
	return (*C.SDL_Rect)(unsafe.Pointer(a))
}

func (p *FPoint) cptr() *C.SDL_FPoint {
	return (*C.SDL_FPoint)(unsafe.Pointer(p))
}

func (a *FRect) cptr() *C.SDL_FRect {
	return (*C.SDL_FRect)(unsafe.Pointer(a))
}

// InRect reports whether the point resides inside a rectangle.
// (https://wiki.libsdl.org/SDL_PointInRect)
func (p *Point) InRect(r *Rect) bool {
	if (p.X >= r.X) && (p.X < (r.X + r.W)) &&
		(p.Y >= r.Y) && (p.Y < (r.Y + r.H)) {
		return true
	}
	return false
}

// InRect reports whether the point resides inside a rectangle.
// (https://wiki.libsdl.org/SDL_PointInRect)
func (p *FPoint) InRect(r *FRect) bool {
	if (p.X >= r.X) && (p.X < (r.X + r.W)) &&
		(p.Y >= r.Y) && (p.Y < (r.Y + r.H)) {
		return true
	}
	return false
}

// Empty reports whether a rectangle has no area.
// (https://wiki.libsdl.org/SDL_RectEmpty)
func (a *Rect) Empty() bool {
	return a == nil || a.W <= 0 || a.H <= 0
}

// Equals reports whether two rectangles are equal.
// (https://wiki.libsdl.org/SDL_RectEquals)
func (a *Rect) Equals(b *Rect) bool {
	if (a != nil) && (b != nil) &&
		(a.X == b.X) && (a.Y == b.Y) &&
		(a.W == b.W) && (a.H == b.H) {
		return true
	}
	return false
}

// HasIntersection reports whether two rectangles intersect.
// (https://wiki.libsdl.org/SDL_HasIntersection)
func (a *Rect) HasIntersection(b *Rect) bool {
	if a == nil || b == nil {
		return false
	}

	// Special case for empty rects
	if a.Empty() || b.Empty() {
		return false
	}

	if a.X >= b.X+b.W || a.X+a.W <= b.X || a.Y >= b.Y+b.H || a.Y+a.H <= b.Y {
		return false
	}

	return true
}

// Intersect calculates the intersection of two rectangles.
// (https://wiki.libsdl.org/SDL_IntersectRect)
func (a *Rect) Intersect(b *Rect) (Rect, bool) {
	var result Rect

	if a == nil || b == nil {
		return result, false
	}

	// Special case for empty rects
	if a.Empty() || b.Empty() {
		result.W = 0
		result.H = 0
		return result, false
	}

	aMin := a.X
	aMax := aMin + a.W
	bMin := b.X
	bMax := bMin + b.W
	if bMin > aMin {
		aMin = bMin
	}
	result.X = aMin
	if bMax < aMax {
		aMax = bMax
	}
	result.W = aMax - aMin

	aMin = a.Y
	aMax = aMin + a.H
	bMin = b.Y
	bMax = bMin + b.H
	if bMin > aMin {
		aMin = bMin
	}
	result.Y = aMin
	if bMax < aMax {
		aMax = bMax
	}
	result.H = aMax - aMin

	return result, !result.Empty()
}

// Union calculates the union of two rectangles.
// (https://wiki.libsdl.org/SDL_UnionRect)
func (a *Rect) Union(b *Rect) Rect {
	var result Rect

	if a == nil || b == nil {
		return result
	}

	// Special case for empty rects
	if a.Empty() {
		return *b
	} else if b.Empty() {
		return *a
	} else if a.Empty() && b.Empty() {
		return result
	}

	aMin := a.X
	aMax := aMin + a.W
	bMin := b.X
	bMax := bMin + b.W
	if bMin < aMin {
		aMin = bMin
	}
	result.X = aMin
	if bMax > aMax {
		aMax = bMax
	}
	result.W = aMax - aMin

	aMin = a.Y
	aMax = aMin + a.H
	bMin = b.Y
	bMax = bMin + b.H
	if bMin < aMin {
		aMin = bMin
	}
	result.Y = aMin
	if bMax > aMax {
		aMax = bMax
	}
	result.H = aMax - aMin

	return result
}

// Empty reports whether a rectangle has no area.
// (https://wiki.libsdl.org/SDL_RectEmpty)
func (a *FRect) Empty() bool {
	return a == nil || a.W <= 0 || a.H <= 0
}

// Equals reports whether two rectangles are equal.
// (https://wiki.libsdl.org/SDL_RectEquals)
func (a *FRect) Equals(b *FRect) bool {
	if (a != nil) && (b != nil) &&
		(a.X == b.X) && (a.Y == b.Y) &&
		(a.W == b.W) && (a.H == b.H) {
		return true
	}
	return false
}

// EqualsEpsilon returns true if the two rectangles are equal, within some given epsilon.
// (https://wiki.libsdl.org/SDL_FRectEqualsEpsilon)
func (a *FRect) EqualsEpsilon(b *FRect, epsilon float32) bool {
	if (a != nil) && (b != nil) && (a == b ||
		(float32(math.Abs(float64(a.X-b.X))) <= epsilon) &&
			(float32(math.Abs(float64(a.Y-b.Y))) <= epsilon) &&
			(float32(math.Abs(float64(a.W-b.W))) <= epsilon) &&
			(float32(math.Abs(float64(a.H-b.H))) <= epsilon)) {
		return true
	}
	return false
}

// HasIntersection reports whether two rectangles intersect.
// (https://wiki.libsdl.org/SDL_HasIntersection)
func (a *FRect) HasIntersection(b *FRect) bool {
	if a == nil || b == nil {
		return false
	}

	// Special case for empty rects
	if a.Empty() || b.Empty() {
		return false
	}

	if a.X >= b.X+b.W || a.X+a.W <= b.X || a.Y >= b.Y+b.H || a.Y+a.H <= b.Y {
		return false
	}

	return true
}

// Intersect calculates the intersection of two rectangles.
// (https://wiki.libsdl.org/SDL_IntersectRect)
func (a *FRect) Intersect(b *FRect) (FRect, bool) {
	var result FRect

	if a == nil || b == nil {
		return result, false
	}

	// Special case for empty rects
	if a.Empty() || b.Empty() {
		result.W = 0
		result.H = 0
		return result, false
	}

	aMin := a.X
	aMax := aMin + a.W
	bMin := b.X
	bMax := bMin + b.W
	if bMin > aMin {
		aMin = bMin
	}
	result.X = aMin
	if bMax < aMax {
		aMax = bMax
	}
	result.W = aMax - aMin

	aMin = a.Y
	aMax = aMin + a.H
	bMin = b.Y
	bMax = bMin + b.H
	if bMin > aMin {
		aMin = bMin
	}
	result.Y = aMin
	if bMax < aMax {
		aMax = bMax
	}
	result.H = aMax - aMin

	return result, !result.Empty()
}

// Union calculates the union of two rectangles.
// (https://wiki.libsdl.org/SDL_UnionRect)
func (a *FRect) Union(b *FRect) FRect {
	var result FRect

	if a == nil || b == nil {
		return result
	}

	// Special case for empty rects
	if a.Empty() {
		return *b
	} else if b.Empty() {
		return *a
	} else if a.Empty() && b.Empty() {
		return result
	}

	aMin := a.X
	aMax := aMin + a.W
	bMin := b.X
	bMax := bMin + b.W
	if bMin < aMin {
		aMin = bMin
	}
	result.X = aMin
	if bMax > aMax {
		aMax = bMax
	}
	result.W = aMax - aMin

	aMin = a.Y
	aMax = aMin + a.H
	bMin = b.Y
	bMax = bMin + b.H
	if bMin < aMin {
		aMin = bMin
	}
	result.Y = aMin
	if bMax > aMax {
		aMax = bMax
	}
	result.H = aMax - aMin

	return result
}

// EnclosePoints calculates a minimal rectangle that encloses a set of points.
// (https://wiki.libsdl.org/SDL_EnclosePoints)
func EnclosePoints(points []Point, clip *Rect) (Rect, bool) {
	var result Rect

	if len(points) == 0 {
		return result, false
	}

	var minX, minY, maxX, maxY int32
	if clip != nil {
		added := false
		clipMinX := clip.X
		clipMinY := clip.Y
		clipMaxX := clip.X + clip.W - 1
		clipMaxY := clip.Y + clip.H - 1

		// If the clip has no size, we're done
		if clip.Empty() {
			return result, false
		}

		for _, val := range points {
			// Check if the point is inside the clip rect
			if val.X < clipMinX || val.X > clipMaxX || val.Y < clipMinY || val.Y > clipMaxY {
				continue
			}

			if !added {
				// If it's the first point
				minX = val.X
				maxX = val.X
				minY = val.Y
				maxY = val.Y
				added = true
			}

			// Find mins and maxes
			if val.X < minX {
				minX = val.X
			} else if val.X > maxX {
				maxX = val.X
			}
			if val.Y < minY {
				minY = val.Y
			} else if val.Y > maxY {
				maxY = val.Y
			}
		}
	} else {
		for i, val := range points {
			if i == 0 {
				// Populate the first point
				minX = val.X
				maxX = val.X
				minY = val.Y
				maxY = val.Y
				continue
			}

			// Find mins and maxes
			if val.X < minX {
				minX = val.X
			} else if val.X > maxX {
				maxX = val.X
			}
			if val.Y < minY {
				minY = val.Y
			} else if val.Y > maxY {
				maxY = val.Y
			}
		}
	}

	result.X = minX
	result.Y = minY
	result.W = (maxX - minX) + 1
	result.H = (maxY - minY) + 1

	return result, true
}

// EncloseFPoints calculates a minimal rectangle that encloses a set of points with float precision.
// (https://wiki.libsdl.org/SDL_EncloseFPoints)
func EncloseFPoints(points []FPoint, clip *FRect) (result FRect, enclosed bool) {
	if len(points) == 0 {
		return result, false
	}

	var minX, minY, maxX, maxY float32
	if clip != nil {
		added := false
		clipMinX := clip.X
		clipMinY := clip.Y
		clipMaxX := clip.X + clip.W - 1
		clipMaxY := clip.Y + clip.H - 1

		// If the clip has no size, we're done
		if clip.Empty() {
			return result, false
		}

		for _, val := range points {
			// Check if the point is inside the clip rect
			if val.X < clipMinX || val.X > clipMaxX || val.Y < clipMinY || val.Y > clipMaxY {
				continue
			}

			if !added {
				// If it's the first point
				minX = val.X
				maxX = val.X
				minY = val.Y
				maxY = val.Y
				added = true
			}

			// Find mins and maxes
			if val.X < minX {
				minX = val.X
			} else if val.X > maxX {
				maxX = val.X
			}
			if val.Y < minY {
				minY = val.Y
			} else if val.Y > maxY {
				maxY = val.Y
			}
		}
	} else {
		for i, val := range points {
			if i == 0 {
				// Populate the first point
				minX = val.X
				maxX = val.X
				minY = val.Y
				maxY = val.Y
				continue
			}

			// Find mins and maxes
			if val.X < minX {
				minX = val.X
			} else if val.X > maxX {
				maxX = val.X
			}
			if val.Y < minY {
				minY = val.Y
			} else if val.Y > maxY {
				maxY = val.Y
			}
		}
	}

	result.X = minX
	result.Y = minY
	result.W = (maxX - minX) + 1
	result.H = (maxY - minY) + 1

	return result, true
}

const (
	codeBottom = 1
	codeTop    = 2
	codeLeft   = 4
	codeRight  = 8
)

func computeOutCode(rect *Rect, x, y int32) int {
	code := 0
	if y < rect.Y {
		code |= codeTop
	} else if y >= rect.Y+rect.H {
		code |= codeBottom
	}
	if x < rect.X {
		code |= codeLeft
	} else if x >= rect.X+rect.W {
		code |= codeRight
	}

	return code
}

func computeFOutCode(rect *FRect, x, y float32) int {
	code := 0
	if y < rect.Y {
		code |= codeTop
	} else if y >= rect.Y+rect.H {
		code |= codeBottom
	}
	if x < rect.X {
		code |= codeLeft
	} else if x >= rect.X+rect.W {
		code |= codeRight
	}

	return code
}

// IntersectLine calculates the intersection of a rectangle and a line segment.
// (https://wiki.libsdl.org/SDL_IntersectRectAndLine)
func (a *Rect) IntersectLine(X1, Y1, X2, Y2 *int32) bool {
	if a.Empty() {
		return false
	}

	x1 := *X1
	y1 := *Y1
	x2 := *X2
	y2 := *Y2
	rectX1 := a.X
	rectY1 := a.Y
	rectX2 := a.X + a.W - 1
	rectY2 := a.Y + a.H - 1

	// Check if the line is entirely inside the rect
	if x1 >= rectX1 && x1 <= rectX2 && x2 >= rectX1 && x2 <= rectX2 &&
		y1 >= rectY1 && y1 <= rectY2 && y2 >= rectY1 && y2 <= rectY2 {
		return true
	}

	// Check if the line is entirely outside the rect
	if (x1 < rectX1 && x2 < rectX1) || (x1 > rectX2 && x2 > rectX2) ||
		(y1 < rectY1 && y2 < rectY1) || (y1 > rectY2 && y2 > rectY2) {
		return false
	}

	// Check if the line is horizontal
	if y1 == y2 {
		if x1 < rectX1 {
			*X1 = rectX1
		} else if x1 > rectX2 {
			*X1 = rectX2
		}
		if x2 < rectX1 {
			*X2 = rectX1
		} else if x2 > rectX2 {
			*X2 = rectX2
		}

		return true
	}

	// Check if the line is vertical
	if x1 == x2 {
		if y1 < rectY1 {
			*Y1 = rectY1
		} else if y1 > rectY2 {
			*Y1 = rectY2
		}
		if y2 < rectY1 {
			*Y2 = rectY1
		} else if y2 > rectY2 {
			*Y2 = rectY2
		}

		return true
	}

	// Use Cohen-Sutherland algorithm when all shortcuts fail
	outCode1 := computeOutCode(a, x1, y1)
	outCode2 := computeOutCode(a, x2, y2)
	for outCode1 != 0 || outCode2 != 0 {
		if outCode1&outCode2 != 0 {
			return false
		}

		if outCode1 != 0 {
			var x, y int32
			if outCode1&codeTop != 0 {
				y = rectY1
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode1&codeBottom != 0 {
				y = rectY2
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode1&codeLeft != 0 {
				x = rectX1
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			} else if outCode1&codeRight != 0 {
				x = rectX2
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			}

			x1 = x
			y1 = y
			outCode1 = computeOutCode(a, x, y)
		} else {
			var x, y int32
			if outCode2&codeTop != 0 {
				y = rectY1
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode2&codeBottom != 0 {
				y = rectY2
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode2&codeLeft != 0 {
				x = rectX1
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			} else if outCode2&codeRight != 0 {
				x = rectX2
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			}

			x2 = x
			y2 = y
			outCode2 = computeOutCode(a, x, y)
		}
	}

	*X1 = x1
	*Y1 = y1
	*X2 = x2
	*Y2 = y2

	return true
}

// IntersectLine calculates the intersection of a rectangle and a line segment.
// (https://wiki.libsdl.org/SDL_IntersectFRectAndLine)
func (a *FRect) IntersectLine(X1, Y1, X2, Y2 *float32) bool {
	if a.Empty() {
		return false
	}

	x1 := *X1
	y1 := *Y1
	x2 := *X2
	y2 := *Y2
	rectX1 := a.X
	rectY1 := a.Y
	rectX2 := a.X + a.W - 1
	rectY2 := a.Y + a.H - 1

	// Check if the line is entirely inside the rect
	if x1 >= rectX1 && x1 <= rectX2 && x2 >= rectX1 && x2 <= rectX2 &&
		y1 >= rectY1 && y1 <= rectY2 && y2 >= rectY1 && y2 <= rectY2 {
		return true
	}

	// Check if the line is entirely outside the rect
	if (x1 < rectX1 && x2 < rectX1) || (x1 > rectX2 && x2 > rectX2) ||
		(y1 < rectY1 && y2 < rectY1) || (y1 > rectY2 && y2 > rectY2) {
		return false
	}

	// Check if the line is horizontal
	if y1 == y2 {
		if x1 < rectX1 {
			*X1 = rectX1
		} else if x1 > rectX2 {
			*X1 = rectX2
		}
		if x2 < rectX1 {
			*X2 = rectX1
		} else if x2 > rectX2 {
			*X2 = rectX2
		}

		return true
	}

	// Check if the line is vertical
	if x1 == x2 {
		if y1 < rectY1 {
			*Y1 = rectY1
		} else if y1 > rectY2 {
			*Y1 = rectY2
		}
		if y2 < rectY1 {
			*Y2 = rectY1
		} else if y2 > rectY2 {
			*Y2 = rectY2
		}

		return true
	}

	// Use Cohen-Sutherland algorithm when all shortcuts fail
	outCode1 := computeFOutCode(a, x1, y1)
	outCode2 := computeFOutCode(a, x2, y2)
	for outCode1 != 0 || outCode2 != 0 {
		if outCode1&outCode2 != 0 {
			return false
		}

		if outCode1 != 0 {
			var x, y float32
			if outCode1&codeTop != 0 {
				y = rectY1
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode1&codeBottom != 0 {
				y = rectY2
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode1&codeLeft != 0 {
				x = rectX1
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			} else if outCode1&codeRight != 0 {
				x = rectX2
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			}

			x1 = x
			y1 = y
			outCode1 = computeFOutCode(a, x, y)
		} else {
			var x, y float32
			if outCode2&codeTop != 0 {
				y = rectY1
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode2&codeBottom != 0 {
				y = rectY2
				x = x1 + ((x2-x1)*(y-y1))/(y2-y1)
			} else if outCode2&codeLeft != 0 {
				x = rectX1
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			} else if outCode2&codeRight != 0 {
				x = rectX2
				y = y1 + ((y2-y1)*(x-x1))/(x2-x1)
			}

			x2 = x
			y2 = y
			outCode2 = computeFOutCode(a, x, y)
		}
	}

	*X1 = x1
	*Y1 = y1
	*X2 = x2
	*Y2 = y2

	return true
}
