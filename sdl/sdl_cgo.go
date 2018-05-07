// +build !static

package sdl

//#cgo windows LDFLAGS: -lSDL2
//#cgo linux freebsd darwin pkg-config: sdl2
import "C"
