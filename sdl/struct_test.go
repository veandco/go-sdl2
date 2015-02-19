package sdl

import (
	"reflect"
	"strings"
	"testing"
)

// TODO: AudioCVT is a packed struct in C - proper conversion needs some more work
//func TestStructAudioCVT(t *testing.T)  { testStructEquality(t, AudioCVT{}, cAudioCVT{}) }
func TestStructAudioSpec(t *testing.T) { testStructEquality(t, AudioSpec{}, cAudioSpec{}) }

func TestStructKeyDownEvent(t *testing.T) { testStructEquality(t, KeyDownEvent{}, cKeyboardEvent{}) }
func TestStructKeyUpEvent(t *testing.T)   { testStructEquality(t, KeyUpEvent{}, cKeyboardEvent{}) }
func TestStructMouseButtonEvent(t *testing.T) {
	testStructEquality(t, MouseButtonEvent{}, cMouseButtonEvent{})
}
func TestStructMouseMotionEvent(t *testing.T) {
	testStructEquality(t, MouseMotionEvent{}, cMouseMotionEvent{})
}
func TestStructMouseWheelEvent(t *testing.T) {
	testStructEquality(t, MouseWheelEvent{}, cMouseWheelEvent{})
}
func TestStructTextEditingEvent(t *testing.T) {
	testStructEquality(t, TextEditingEvent{}, cTextEditingEvent{})
}
func TestStructTextInputEvent(t *testing.T) {
	testStructEquality(t, TextInputEvent{}, cTextInputEvent{})
}

func TestStructPixelFormat(t *testing.T) { testStructEquality(t, PixelFormat{}, cPixelFormat{}) }
func TestStructSurface(t *testing.T)     { testStructEquality(t, Surface{}, cSurface{}) }
func TestStructVersion(t *testing.T)     { testStructEquality(t, Version{}, cVersion{}) }
func TestStructWindowEvent(t *testing.T) { testStructEquality(t, WindowEvent{}, cWindowEvent{}) }

//TODO
//func TestStructSysWmInfo(t *testing.T)   { testStructEquality(t, SysWMInfo{}, cSysWMinfo{}) }

func testStructEquality(t *testing.T, a, b interface{}) {
	ta, tb := reflect.TypeOf(a), reflect.TypeOf(b)
	if ta.Size() != tb.Size() {
		t.Errorf("struct size missmatch: %s(%d) != %s(%d)",
			ta.Name(), ta.Size(), tb.Name(), tb.Size())
	}

	dumpStructFormat(t, ta)
	dumpStructFormat(t, tb)

	for i := 0; i < ta.NumField(); i++ {
		f := ta.Field(i)
		if f.Name == "_" {
			// ignore padding fields
			continue
		}
		verifyField(t, f, tb)
	}
}

func dumpStructFormat(t *testing.T, s reflect.Type) {
	const dFormat = "%5d %30s %8d %8d"
	var sFormat = strings.Replace(dFormat, "d", "s", -1)

	t.Logf(sFormat, "index", "name", "offset", "size")
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		t.Logf(dFormat, i, s.Name()+"."+f.Name, f.Offset, f.Type.Size())
	}
}

func verifyField(t *testing.T, gField reflect.StructField, cType reflect.Type) {
	var cField = searchField(cType, gField.Name)
	if cField == nil {
		t.Error("field not found:", gField.Name)
		return
	}
	gOffset, gSize := gField.Offset, gField.Type.Size()
	cOffset, cSize := cField.Offset, cField.Type.Size()
	if cOffset != gOffset || cSize != gSize {
		t.Errorf("field offset/size missmatch %s(%d, %d) != %s(%d, %d)",
			gField.Name, gOffset, gSize,
			cField.Name, cOffset, cSize)
	}
}

func searchField(t reflect.Type, gName string) *reflect.StructField {
	for i := 0; i < t.NumField(); i++ {
		var f = t.Field(i)
		// convert c field_name to fieldname
		var cName = strings.Replace(f.Name, "_", "", -1)
		if strings.EqualFold(gName, cName) {
			return &f
		}
	}
	return nil
}
