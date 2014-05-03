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

func U8To32Array(buf []byte) []uint32 {
        var ret []uint32
        header := (*reflect.SliceHeader)(unsafe.Pointer(&ret))
        header.Cap = len(buf) / 4 + 1
        header.Len = len(buf) / 4 + 1
        header.Data = uintptr(unsafe.Pointer(&buf[0]))
        return ret
}

func Endian() int {
	var x uint32 = 0x01020304
	var t int
	switch *(*byte)(unsafe.Pointer(&x)) {
	case 0x01:
		t = BIG_ENDIAN
	case 0x04:
		t = LIL_ENDIAN
	}

	return t
}
