// +build !static

package img

//#cgo linux freebsd darwin pkg-config: sdl2
//#cgo linux freebsd darwin LDFLAGS: -lSDL2_image
//#cgo windows LDFLAGS: -lSDL2 -lSDL2_image
import "C"
