// +build static

package img

//#cgo CFLAGS: -I${SRCDIR}/../_libs/include -I${SRCDIR}/../_libs/include/SDL2
//#cgo LDFLAGS: -L${SRCDIR}/../_libs
//#cgo linux,386 LDFLAGS: -lSDL2_image_linux_386 -Wl,--no-undefined -lpng_linux_386 -ljpeg_linux_386 -lSDL2_linux_386 -lm -ldl -lz -lasound -lm -ldl -lpthread -lX11 -lXext -lXcursor -lXinerama -lXi -lXrandr -lXss -lXxf86vm -lpthread -lrt
//#cgo linux,amd64 LDFLAGS: -lSDL2_image_linux_amd64 -Wl,--no-undefined -lpng_linux_amd64 -ljpeg_linux_amd64 -lSDL2_linux_amd64 -lm -ldl -lz -lasound -lm -ldl -lpthread -lX11 -lXext -lXcursor -lXinerama -lXi -lXrandr -lXss -lXxf86vm -lpthread -lrt
//#cgo windows,386 LDFLAGS: -lSDL2_image_windows_386 -Wl,--no-undefined -lpng_windows_386 -ljpeg_windows_386 -lz_windows_386 -lSDL2_windows_386 -lSDL2main_windows_386 -Wl,--no-undefined -lm -ldinput8 -ldxguid -ldxerr8 -luser32 -lgdi32 -lwinmm -limm32 -lole32 -loleaut32 -lshell32 -lversion -luuid -lsetupapi -static-libgcc
//#cgo windows,amd64 LDFLAGS: -lSDL2_image_windows_amd64 -Wl,--no-undefined -lpng_windows_amd64 -ljpeg_windows_amd64 -lz_windows_amd64 -lSDL2_windows_amd64 -lSDL2main_windows_amd64 -Wl,--no-undefined -lm -ldinput8 -ldxguid -ldxerr8 -luser32 -lgdi32 -lwinmm -limm32 -lole32 -loleaut32 -lshell32 -lversion -luuid -lsetupapi -static-libgcc
//#cgo darwin,amd64 LDFLAGS: -lSDL2_image_darwin_amd64 -lSDL2_darwin_amd64 -lpng_darwin_amd64 -ljpeg_darwin_amd64 -lm -lz -liconv -Wl,-framework,CoreAudio -Wl,-framework,AudioToolbox -Wl,-framework,ForceFeedback -lobjc -Wl,-framework,CoreVideo -Wl,-framework,Cocoa -Wl,-framework,Carbon -Wl,-framework,IOKit -Wl,-framework,Metal
//#cgo darwin,arm64 LDFLAGS: -lSDL2_image_darwin_arm64 -lSDL2_darwin_arm64 -lpng_darwin_arm64 -ljpeg_darwin_arm64 -lm -lz -liconv -Wl,-framework,CoreAudio -Wl,-framework,AudioToolbox -Wl,-framework,ForceFeedback -lobjc -Wl,-framework,CoreVideo -Wl,-framework,Cocoa -Wl,-framework,Carbon -Wl,-framework,IOKit -Wl,-framework,Metal
//#cgo android,arm LDFLAGS: -lSDL2_image_android_arm -Wl,--no-undefined -lpng_android_arm -ljpeg_android_arm -lSDL2_android_arm -lm -ldl -lz -llog -landroid -lGLESv2 -lGLESv1_CM
//#cgo linux,arm,!android LDFLAGS: -L/opt/vc/lib -L/opt/vc/lib64 -lSDL2_image_linux_arm -Wl,--no-undefined -lpng_linux_arm -ljpeg_linux_arm -lSDL2_linux_arm -lm -ldl -lz -liconv -lbcm_host -lvcos -lvchiq_arm -pthread
import "C"
