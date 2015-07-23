package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// Point (https://wiki.libsdl.org/SDL_Point)
type Point struct {
	X int32
	Y int32
}

// Rect (https://wiki.libsdl.org/SDL_Rect)
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

// Rect (https://wiki.libsdl.org/SDL_RectEmpty)
func (r *Rect) Empty() bool {
	return r == nil || r.W <= 0 || r.H <= 0
}

// Rect (https://wiki.libsdl.org/SDL_RectEquals)
func (a *Rect) Equals(b *Rect) bool {
	if (a != nil) && (b != nil) &&
		(a.X == b.X) && (a.Y == a.Y) &&
		(a.W == b.W) && (a.H == b.H) {
		return true
	}
	return false
}

// Rect (https://wiki.libsdl.org/SDL_HasIntersection)
func (a *Rect) HasIntersection(b *Rect) bool {
	return C.SDL_HasIntersection(a.cptr(), b.cptr()) > 0
}

// Rect (https://wiki.libsdl.org/SDL_IntersectRect)
func (a *Rect) Intersect(b *Rect) (result Rect, ok bool) {
	ok = C.SDL_IntersectRect(a.cptr(), b.cptr(), result.cptr()) > 0
	return
}

// Rect (https://wiki.libsdl.org/SDL_UnionRect)
func (a *Rect) Union(b *Rect) (result Rect) {
	C.SDL_UnionRect(a.cptr(), b.cptr(), result.cptr())
	return
}

// EnclosePoints (https://wiki.libsdl.org/SDL_EnclosePoints)
func EnclosePoints(points []Point, clip *Rect) (result Rect, ok bool) {
	ok = C.SDL_EnclosePoints(points[0].cptr(), C.int(len(points)), clip.cptr(), result.cptr()) > 0
	return
}

// IntersectLine (https://wiki.libsdl.org/SDL_IntersectRectAndLine)
func (r *Rect) IntersectLine(X1, Y1, X2, Y2 *int) bool {
	X1_32, Y1_32, X2_32, Y2_32 := int32(*X1), int32(*Y1), int32(*X2), int32(*Y2)

	_X1 := (*C.int)(unsafe.Pointer(&X1_32))
	_Y1 := (*C.int)(unsafe.Pointer(&Y1_32))
	_X2 := (*C.int)(unsafe.Pointer(&X2_32))
	_Y2 := (*C.int)(unsafe.Pointer(&Y2_32))

	result := C.SDL_IntersectRectAndLine(r.cptr(), _X1, _Y1, _X2, _Y2) == C.SDL_TRUE

	*X1 = int(X1_32)
	*Y1 = int(Y1_32)
	*X2 = int(X2_32)
	*Y2 = int(Y2_32)

	return result
}
