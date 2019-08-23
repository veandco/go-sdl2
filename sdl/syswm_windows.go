// +build windows

package sdl

/*
#include <windef.h>
*/
import "C"
import "unsafe"

// WindowsMsg contains Microsoft Windows window information.
type WindowsMsg struct {
	Hwnd   C.HWND
	Msg    C.UINT
	WParam C.WPARAM
	LParam C.LPARAM
}

// Windows() returns Microsoft Windows message.
func (msg *SysWMmsg) Windows() *WindowsMsg {
	return (*WindowsMsg)(unsafe.Pointer(&msg.data[0]))
}
