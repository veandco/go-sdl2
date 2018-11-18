// +build static

package sdl

//#cgo CFLAGS: -I${SRCDIR}/../.go-sdl2-libs/include -I${SRCDIR}/../.go-sdl2-libs/include/SDL2
//#cgo LDFLAGS: -L${SRCDIR}/../.go-sdl2-libs
//#cgo linux,386 LDFLAGS: -lSDL2_linux_386 -Wl,--no-undefined -lm -ldl -lasound -lm -ldl -lpthread -lX11 -lXext -lXcursor -lXinerama -lXi -lXrandr -lXss -lXxf86vm -lpthread -lrt
//#cgo linux,amd64 LDFLAGS: -lSDL2_linux_amd64 -Wl,--no-undefined -lm -ldl -lasound -lm -ldl -lpthread -lX11 -lXext -lXcursor -lXinerama -lXi -lXrandr -lXss -lXxf86vm -lpthread -lrt
//#cgo windows,386 LDFLAGS: -lSDL2_windows_386 -lSDL2main_windows_386 -mwindows -Wl,--no-undefined -lm -ldinput8 -ldxguid -ldxerr8 -luser32 -lgdi32 -lwinmm -limm32 -lole32 -loleaut32 -lshell32 -lsetupapi -lversion -luuid -static-libgcc
//#cgo windows,amd64 LDFLAGS: -lSDL2_windows_amd64 -lSDL2main_windows_amd64 -mwindows -Wl,--no-undefined -lm -ldinput8 -ldxguid -ldxerr8 -luser32 -lgdi32 -lwinmm -limm32 -lole32 -loleaut32 -lshell32 -lversion -luuid -lsetupapi -static-libgcc
//#cgo darwin,amd64 LDFLAGS: -lSDL2_darwin_amd64 -lm -liconv -Wl,-framework,CoreAudio -Wl,-framework,AudioToolbox -Wl,-framework,ForceFeedback -lobjc -Wl,-framework,CoreVideo -Wl,-framework,Cocoa -Wl,-framework,Carbon -Wl,-framework,IOKit -Wl,-framework,Metal
//#cgo android,arm LDFLAGS: -lSDL2_android_arm -Wl,--no-undefined -lm -ldl -llog -landroid -lGLESv2
//#cgo linux,arm,!android LDFLAGS: -L/opt/vc/lib -L/opt/vc/lib64 -lSDL2_linux_arm -Wl,--no-undefined -lm -ldl -liconv -lbcm_host -lvcos -lvchiq_arm -pthread
import "C"
