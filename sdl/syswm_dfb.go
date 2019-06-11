// +build dfb

package sdl

import "C"
import "unsafe"

// DFBMsg contains DirectFB window information.
type DFBMsg struct {
	Event C.DFBEvent
}

// DFB() returns DirectFB message.
func (msg *SysWMmsg) DFB() *DFBMsg {
	return (*DFBMsg)(unsafe.Pointer(&msg.data[0]))
}
