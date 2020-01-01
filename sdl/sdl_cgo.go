// +build !static

package sdl

//#cgo windows LDFLAGS: -lSDL2
//#cgo linux freebsd darwin openbsd pkg-config: sdl2
import "C"
