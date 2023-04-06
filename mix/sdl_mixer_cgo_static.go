// +build static

package mix

//#cgo CFLAGS: -I${SRCDIR}/../_libs/include -I${SRCDIR}/../_libs/include/SDL2
//#cgo LDFLAGS: -L${SRCDIR}/../_libs
//#cgo linux,386 LDFLAGS: -lSDL2_mixer_linux_386 -Wl,--no-undefined -lmpg123_linux_386 -lvorbisfile_linux_386 -lvorbis_linux_386 -logg_linux_386 -lSDL2_linux_386 -lm -ldl -lasound -lm -ldl -lstdc++ -lpthread -lX11 -lXext -lXcursor -lXinerama -lXi -lXrandr -lXss -lXxf86vm -lpthread -lrt
//#cgo linux,amd64 LDFLAGS: -lSDL2_mixer_linux_amd64 -Wl,--no-undefined -lmpg123_linux_amd64 -lvorbisfile_linux_amd64 -lvorbis_linux_amd64 -logg_linux_amd64 -lSDL2_linux_amd64 -lm -ldl -lstdc++ -lasound -lm -ldl -lpthread -lX11 -lXext -lXcursor -lXinerama -lXi -lXrandr -lXss -lXxf86vm -lpthread -lrt
//#cgo windows,386 LDFLAGS: -lSDL2_mixer_windows_386 -Wl,--no-undefined -lmpg123_windows_386 -lvorbisfile_windows_386 -lvorbis_windows_386 -logg_windows_386 -lSDL2_windows_386 -lSDL2main_windows_386 -lstdc++ -Wl,--no-undefined -lm -ldinput8 -ldxguid -ldxerr8 -lshlwapi -luser32 -lgdi32 -lwinmm -limm32 -lole32 -loleaut32 -lshell32 -lversion -luuid -lsetupapi -static-libgcc
//#cgo windows,amd64 LDFLAGS: -lSDL2_mixer_windows_amd64 -Wl,--no-undefined -lmpg123_windows_amd64 -lvorbisfile_windows_amd64 -lvorbis_windows_amd64 -logg_windows_amd64 -lSDL2_windows_amd64 -lSDL2main_windows_amd64 -lstdc++ -Wl,--no-undefined -lm -ldinput8 -ldxguid -ldxerr8 -lshlwapi -luser32 -lgdi32 -lwinmm -limm32 -lole32 -loleaut32 -lshell32 -lversion -luuid -lsetupapi -static-libgcc
//#cgo darwin,amd64 LDFLAGS: -lSDL2_mixer_darwin_amd64 -lmpg123_darwin_amd64 -lvorbisfile_darwin_amd64 -lvorbis_darwin_amd64 -logg_darwin_amd64 -lSDL2_darwin_amd64 -lm -liconv -lfreetype_darwin_amd64 -lz_darwin_amd64 -lSDL2_darwin_amd64 -Wl,-framework,CoreAudio -Wl,-framework,AudioToolbox -Wl,-framework,ForceFeedback -lobjc -Wl,-framework,CoreVideo -Wl,-framework,Cocoa -Wl,-framework,Carbon -Wl,-framework,IOKit -Wl,-framework,Metal -Wl,-framework,AudioUnit
//#cgo darwin,arm64 LDFLAGS: -lSDL2_mixer_darwin_arm64 -lmpg123_darwin_arm64 -lvorbisfile_darwin_arm64 -lvorbis_darwin_arm64 -logg_darwin_arm64 -lSDL2_darwin_arm64 -lm -liconv -lfreetype_darwin_arm64 -lz -lSDL2_darwin_arm64 -Wl,-framework,CoreAudio -Wl,-framework,AudioToolbox -Wl,-framework,ForceFeedback -lobjc -Wl,-framework,CoreVideo -Wl,-framework,Cocoa -Wl,-framework,Carbon -Wl,-framework,IOKit -Wl,-framework,Metal -Wl,-framework,AudioUnit
//#cgo android,arm LDFLAGS: -lSDL2_mixer_android_arm -Wl,--no-undefined -lmpg123_android_arm -lvorbisfile_android_arm -lvorbis_android_arm -logg_android_arm -lSDL2_android_arm -lm -ldl -llog -landroid -lGLESv2 -lGLESv1_CM
//#cgo linux,arm,!android LDFLAGS: -L/opt/vc/lib -L/opt/vc/lib64 -lSDL2_mixer_linux_arm -Wl,--no-undefined -lmpg123_linux_arm -lvorbisfile_linux_arm -lvorbis_linux_arm -logg_linux_arm -lSDL2_linux_arm -lm -ldl -liconv -lbcm_host -lvcos -lvchiq_arm -pthread
import "C"
