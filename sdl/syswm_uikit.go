// +build uikit

package sdl

import "C"
import "unsafe"

// UIKitMsg contains Apple iOS window information.
type UIKitMsg struct {
	dummy C.int
}

// UIKit() returns Apple iOS message.
func (msg *SysWMmsg) UIKit() *UIKitMsg {
	return (*UIKitMsg)(unsafe.Pointer(&msg.data[0]))
}
