package sdl

import (
	"reflect"
	"strings"
	"testing"
)

// TODO: AudioCVT is a packed struct in C - proper conversion needs some more work
// func TestStructAudioCVT(t *testing.T)  { testStructEq(t, AudioCVT{}, cAudioCVT{}) }
func TestStructAudioSpec(t *testing.T)   { testStructEq(t, AudioSpec{}, cAudioSpec{}) }
func TestStructDisplayMode(t *testing.T) { testStructEq(t, DisplayMode{}, cDisplayMode{}) }

// EVENTS

func TestStructKeyDown(t *testing.T)     { testStructEq(t, KeyDownEvent{}, cKeyboardEvent{}) }
func TestStructKeyUp(t *testing.T)       { testStructEq(t, KeyUpEvent{}, cKeyboardEvent{}) }
func TestStructMouseButton(t *testing.T) { testStructEq(t, MouseButtonEvent{}, cMouseButtonEvent{}) }
func TestStructMouseMotion(t *testing.T) { testStructEq(t, MouseMotionEvent{}, cMouseMotionEvent{}) }
func TestStructMouseWheel(t *testing.T)  { testStructEq(t, MouseWheelEvent{}, cMouseWheelEvent{}) }
func TestStructTextEditing(t *testing.T) { testStructEq(t, TextEditingEvent{}, cTextEditingEvent{}) }
func TestStructTextInput(t *testing.T)   { testStructEq(t, TextInputEvent{}, cTextInputEvent{}) }
func TestStructJoyAxis(t *testing.T)     { testStructEq(t, JoyAxisEvent{}, cJoyAxisEvent{}) }
func TestStructJoyBall(t *testing.T)     { testStructEq(t, JoyBallEvent{}, cJoyBallEvent{}) }
func TestStructJoyHat(t *testing.T)      { testStructEq(t, JoyHatEvent{}, cJoyHatEvent{}) }
func TestStructJoyButton(t *testing.T)   { testStructEq(t, JoyButtonEvent{}, cJoyButtonEvent{}) }
func TestStructControllerAxis(t *testing.T) {
	testStructEq(t, ControllerAxisEvent{}, cControllerAxisEvent{})
}
func TestStructControllerBtn(t *testing.T) {
	testStructEq(t, ControllerButtonEvent{}, cControllerButtonEvent{})
}
func TestStructControllerDev(t *testing.T) {
	testStructEq(t, ControllerDeviceEvent{}, cControllerDeviceEvent{})
}
func TestStructTouchFinger(t *testing.T)  { testStructEq(t, TouchFingerEvent{}, cTouchFingerEvent{}) }
func TestStructMultiGesture(t *testing.T) { testStructEq(t, MultiGestureEvent{}, cMultiGestureEvent{}) }
func TestStructDollarGesture(t *testing.T) {
	testStructEq(t, DollarGestureEvent{}, cDollarGestureEvent{})
}
func TestStructDrop(t *testing.T)  { testStructEq(t, DropEvent{}, cDropEvent{}) }
func TestStructUser(t *testing.T)  { testStructEq(t, UserEvent{}, cUserEvent{}) }
func TestStructSysWM(t *testing.T) { testStructEq(t, SysWMEvent{}, cSysWMEvent{}) }

// END OF EVENTS

func TestStructPalette(t *testing.T)     { testStructEq(t, Palette{}, cPalette{}) }
func TestStructPixelFormat(t *testing.T) { testStructEq(t, PixelFormat{}, cPixelFormat{}) }
func TestStructSurface(t *testing.T)     { testStructEq(t, Surface{}, cSurface{}) }
func TestStructVersion(t *testing.T)     { testStructEq(t, Version{}, cVersion{}) }
func TestStructWindowEvent(t *testing.T) { testStructEq(t, WindowEvent{}, cWindowEvent{}) }

// TODO: embedded structures make testing difficult
// func TestStructSysWmInfo(t *testing.T)   { testStructEq(t, SysWMInfo{}, cSysWMinfo{}) }
// func TestStructRendererInfo(t *testing.T) {	testStructEq(t, RendererInfo{}, realcRendererInfo{})}

func testStructEq(t *testing.T, a, b interface{}) {
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
