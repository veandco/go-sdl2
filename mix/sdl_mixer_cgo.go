// +build !static

package mix

//#cgo windows LDFLAGS: -lSDL2 -lSDL2_mixer
//#cgo linux freebsd darwin pkg-config: sdl2
//#cgo linux freebsd darwin LDFLAGS: -lSDL2_mixer
import "C"
