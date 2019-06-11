// +build cocoa OR darwin

package sdl

import "C"
import "unsafe"

// CocoaMsg contains Apple Mac OS X window information.
type CocoaMsg struct {
	dummy C.int
}

// Cocoa() returns Apple Mac OS X message.
func (msg *SysWMmsg) Cocoa() *CocoaMsg {
	return (*CocoaMsg)(unsafe.Pointer(&msg.data[0]))
}
