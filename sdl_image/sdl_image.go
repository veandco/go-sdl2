package img

//#cgo LDFLAGS: -lSDL2_image
//#include <SDL2/SDL_image.h>
import "C"
import "unsafe"
import "go-sdl2/sdl"

func Load(file string) *sdl.Surface {
	return (*sdl.Surface) (unsafe.Pointer(C.IMG_Load(C.CString(file))))
}
