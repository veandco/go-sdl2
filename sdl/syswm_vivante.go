// +build vivante

package sdl

import "C"
import "unsafe"

// VivanteKitMsg contains Vivante window information.
type VivanteKitMsg struct {
	dummy C.int
}

// Vivante() returns Vivante message.
func (msg *SysWMmsg) Vivante() *VivanteMsg {
	return (*VivanteMsg)(unsafe.Pointer(&msg.data[0]))
}
