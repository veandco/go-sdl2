package sdl

import (
	"reflect"
	"strings"
	"testing"
)

func TestStructABI(t *testing.T) {
	var tests = []struct {
		gStruct interface{}
		cStruct interface{}
	}{
		{AudioSpec{}, cAudioSpec{}},
		// TODO: AudioCVT is a packed struct in C - proper conversion needs some more work
		//{ AudioCVT{}, cAudioCVT{} },
		{DisplayMode{}, cDisplayMode{}},

		// EVENTS
		{KeyDownEvent{}, cKeyboardEvent{}},
		{KeyUpEvent{}, cKeyboardEvent{}},
		{MouseButtonEvent{}, cMouseButtonEvent{}},
		{MouseMotionEvent{}, cMouseMotionEvent{}},
		{MouseWheelEvent{}, cMouseWheelEvent{}},
		{TextEditingEvent{}, cTextEditingEvent{}},
		{TextInputEvent{}, cTextInputEvent{}},
		{JoyAxisEvent{}, cJoyAxisEvent{}},
		{JoyBallEvent{}, cJoyBallEvent{}},
		{JoyHatEvent{}, cJoyHatEvent{}},
		{JoyButtonEvent{}, cJoyButtonEvent{}},
		{ControllerAxisEvent{}, cControllerAxisEvent{}},
		{ControllerButtonEvent{}, cControllerButtonEvent{}},
		{ControllerDeviceEvent{}, cControllerDeviceEvent{}},
		{TouchFingerEvent{}, cTouchFingerEvent{}},
		{MultiGestureEvent{}, cMultiGestureEvent{}},
		{DollarGestureEvent{}, cDollarGestureEvent{}},
		{DropEvent{}, cDropEvent{}},
		{UserEvent{}, cUserEvent{}},
		{SysWMEvent{}, cSysWMEvent{}},

		{Palette{}, cPalette{}},
		{PixelFormat{}, cPixelFormat{}},
		{Surface{}, cSurface{}},
		{Version{}, cVersion{}},
		{WindowEvent{}, cWindowEvent{}},

		// TODO: embedded structures make testing difficult
		// { SysWMInfo{}, cSysWMinfo{} },
		// { RendererInfo{}, realcRendererInfo{} },
	}

	for _, test := range tests {
		testStructABI(t, test.gStruct, test.cStruct)
	}
}
func testStructABI(t *testing.T, a, b interface{}) {
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
