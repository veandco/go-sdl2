package sdl

import (
	"bytes"
	"reflect"
	"testing"
	"unsafe"
)

var squareWave = []byte("RIFF,\x00\x00\x00WAVEfmt \x10\x00\x00\x00\x01\x00" +
	"\x01\x00\xab \x00\x00\xab \x00\x00\x01\x00\b\x00data\b\x00\x00\x00\x00" +
	"\x00\x00\x00\xff\xff\xff\xff")

func TestLoadWAV_RW(t *testing.T) {
	// load WAV from *RWOps pointing to WAV data
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&squareWave))
	src := RWFromMem(unsafe.Pointer(sliceHeader.Data), len(squareWave))
	buf, spec := LoadWAV_RW(src, 0, &AudioSpec{})

	// test returned []byte
	want := []byte{0, 0, 0, 0, 255, 255, 255, 255}
	if buf == nil {
		t.Errorf("LoadWAV_RW() returned nil []byte")
	} else {
		if bytes.Compare(buf, want) != 0 {
			t.Errorf("LoadWAV_RW() returned %v; want %v", buf, want)
		}
		FreeWAV(buf) // make sure we can free without error
	}

	// test returned *AudioSpec
	if spec == nil {
		t.Errorf("LoadWAV_RW() returned nil *AudioSpec")
	} else if spec.Freq != 8363 { // no need to test all the spec fields
		t.Errorf("LoadWAV_RW() returned *AudioSpec with Freq %d; want %d",
			spec.Freq, 8363)
	}
}
