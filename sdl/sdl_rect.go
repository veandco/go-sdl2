package sdl

// #include <SDL2/SDL_rect.h>
import "C"
import "unsafe"

type Point struct {
	X int32
	Y int32
}

type Rect struct {
	X int32
	Y int32
	W int32
	H int32
}

func (r *Rect) Empty() bool {
	if (r != nil) || (r.W <= 0) || (r.H <= 0) {
		return true
	}
	return false
}

func (a *Rect) Equals(b *Rect) bool {
	if (a != nil) && (b != nil) &&
		(a.X == b.X) && (a.Y == a.Y) &&
		(a.W == b.W) && (a.H == b.H) {
		return true
	}
	return false
}

func (a *Rect) HasIntersection(b *Rect) bool {
	_a := (*C.SDL_Rect) (unsafe.Pointer(a))
	_b := (*C.SDL_Rect) (unsafe.Pointer(b))
	return C.SDL_HasIntersection(_a, _b) > 0
}

func (a *Rect) Intersect(b, result *Rect) bool {
	_a := (*C.SDL_Rect) (unsafe.Pointer(a))
	_b := (*C.SDL_Rect) (unsafe.Pointer(b))
	_result := (*C.SDL_Rect) (unsafe.Pointer(result))
	return C.SDL_IntersectRect(_a, _b, _result) > 0
}

func (a *Rect) Union(b, result *Rect) {
	_a := (*C.SDL_Rect) (unsafe.Pointer(a))
	_b := (*C.SDL_Rect) (unsafe.Pointer(b))
	_result := (*C.SDL_Rect) (unsafe.Pointer(result))
	C.SDL_UnionRect(_a, _b, _result)
}

func EnclosePoints(points *Point, count int, clip, result *Rect) bool {
	_points := (*C.SDL_Point) (unsafe.Pointer(points))
	_count := (C.int) (count)
	_clip := (*C.SDL_Rect) (unsafe.Pointer(clip))
	_result := (*C.SDL_Rect) (unsafe.Pointer(result))
	return C.SDL_EnclosePoints(_points, _count, _clip, _result) > 0
}

func (rect *Rect) IntersectLine(X1, Y1, X2, Y2 *int) bool {
	_rect := (*C.SDL_Rect) (unsafe.Pointer(rect))
	_X1 := (*C.int) (unsafe.Pointer(X1))
	_Y1 := (*C.int) (unsafe.Pointer(Y1))
	_X2 := (*C.int) (unsafe.Pointer(X2))
	_Y2 := (*C.int) (unsafe.Pointer(Y2))
	return C.SDL_IntersectRectAndLine(_rect, _X1, _Y1, _X2, _Y2) > 0
}
