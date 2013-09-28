package sdl

import (
	"unsafe"
	"reflect"
)

func Btoi(b bool) int {
	if b == true {
		return 1
	}

	return 0
}

func Itob(i int) bool {
	if i > 0 {
		return true
	}

	return false
}

func U8To32Array(buf []byte) []uint32 {
        var ret []uint32
        header := (*reflect.SliceHeader)(unsafe.Pointer(&ret))
        header.Cap = len(ret) / 4 + 1
        header.Len = len(ret) / 4 + 1
        header.Data = uintptr(unsafe.Pointer(&buf[0]))
        return ret
}
