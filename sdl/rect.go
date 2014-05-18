package sdl

// #include <SDL2/SDL.h>
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

func (p *Point) cptr() *C.SDL_Point {
	return (*C.SDL_Point)(unsafe.Pointer(p))
}

func (r *Rect) cptr() *C.SDL_Rect {
	return (*C.SDL_Rect)(unsafe.Pointer(r))
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
	return C.SDL_HasIntersection(a.cptr(), b.cptr()) > 0
}

func (a *Rect) Intersect(b, result *Rect) bool {
	return C.SDL_IntersectRect(a.cptr(), b.cptr(), result.cptr()) > 0
}

func (a *Rect) Union(b, result *Rect) {
	C.SDL_UnionRect(a.cptr(), b.cptr(), result.cptr())
}

func EnclosePoints(points []Point, clip, result *Rect) bool {
	return C.SDL_EnclosePoints(points[0].cptr(), C.int(len(points)), clip.cptr(), result.cptr()) > 0
}

func (rect *Rect) IntersectLine(X1, Y1, X2, Y2 *int) bool {
	_X1 := (*C.int)(unsafe.Pointer(X1))
	_Y1 := (*C.int)(unsafe.Pointer(Y1))
	_X2 := (*C.int)(unsafe.Pointer(X2))
	_Y2 := (*C.int)(unsafe.Pointer(Y2))
	return C.SDL_IntersectRectAndLine(rect.cptr(), _X1, _Y1, _X2, _Y2) > 0
}
