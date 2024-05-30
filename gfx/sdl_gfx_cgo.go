// +build !static

package gfx

//#cgo windows LDFLAGS: -lSDL2 -lSDL2_gfx
//#cgo linux freebsd darwin pkg-config: SDL2_gfx
//#cgo linux freebsd darwin LDFLAGS: -lSDL2_gfx
import "C"
