// +build x11

package sdl

/*
#include <X11/Xlib.h>
*/
import "C"
import "unsafe"

// X11Msg contains X Window System window information.
type X11Msg struct {
	Event C.XEvent
}

// X11() returns X Window System message.
func (msg *SysWMmsg) X11() *X11Msg {
	return (*X11Msg)(unsafe.Pointer(&msg.data[0]))
}
