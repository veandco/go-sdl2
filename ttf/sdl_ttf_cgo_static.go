//go:build static
// +build static

package ttf

//#cgo CFLAGS: -I${SRCDIR}/../_libs/include -I${SRCDIR}/../_libs/include/SDL2
//#cgo LDFLAGS: -L${SRCDIR}/../_libs
//#cgo linux,386 LDFLAGS: -lSDL2_ttf_linux_386 -Wl,--no-undefined -lfreetype_linux_386 -lSDL2_linux_386 -lm -ldl -lasound -lm -ldl -lpthread -lX11 -lXext -lXcursor -lXinerama -lXi -lXrandr -lXss -lXxf86vm -lpthread -lrt
//#cgo linux,amd64 LDFLAGS: -lSDL2_ttf_linux_amd64 -Wl,--no-undefined -lfreetype_linux_amd64 -lSDL2_linux_amd64 -lm -ldl -lasound -lm -ldl -lpthread -lX11 -lXext -lXcursor -lXinerama -lXi -lXrandr -lXss -lXxf86vm -lpthread -lrt
//#cgo linux,mipsle LDFLAGS: -lSDL2_ttf_linux_mipsle -Wl,--no-undefined -lfreetype_linux_mipsle -lSDL2_linux_mipsle -lm -ldl
//#cgo windows,386 LDFLAGS: -lSDL2_ttf_windows_386 -Wl,--no-undefined -lfreetype_windows_386 -lSDL2_windows_386 -lSDL2main_windows_386 -Wl,--no-undefined -lm -ldinput8 -ldxguid -ldxerr8 -luser32 -lgdi32 -lwinmm -limm32 -lole32 -loleaut32 -lshell32 -lversion -luuid -lsetupapi -lrpcrt4 -static-libgcc
//#cgo windows,amd64 LDFLAGS: -lSDL2_ttf_windows_amd64 -Wl,--no-undefined -lfreetype_windows_amd64 -lSDL2_windows_amd64 -lSDL2main_windows_amd64 -Wl,--no-undefined -lm -ldinput8 -ldxguid -ldxerr8 -luser32 -lgdi32 -lwinmm -limm32 -lole32 -loleaut32 -lshell32 -lversion -luuid -lsetupapi -lrpcrt4 -static-libgcc
//#cgo darwin,amd64 LDFLAGS: -lSDL2_ttf_darwin_amd64 -lm -liconv -lfreetype_darwin_amd64 -lSDL2_darwin_amd64 -Wl,-framework,CoreAudio -Wl,-framework,AudioToolbox -Wl,-framework,ForceFeedback -lobjc -Wl,-framework,CoreVideo -Wl,-framework,Cocoa -Wl,-framework,Carbon -Wl,-framework,IOKit -Wl,-framework,Metal
//#cgo android,arm LDFLAGS: -lSDL2_ttf_android_arm -Wl,--no-undefined -lfreetype_android_arm -lSDL2_android_arm -lm -ldl -llog -landroid -lGLESv2 -lGLESv1_CM
//#cgo linux,arm,!android LDFLAGS: -L/opt/vc/lib -L/opt/vc/lib64 -lSDL2_ttf_linux_arm -Wl,--no-undefined -lfreetype_linux_arm -lSDL2_linux_arm -lm -ldl -liconv -lbcm_host -lvcos -lvchiq_arm -pthread
import "C"
