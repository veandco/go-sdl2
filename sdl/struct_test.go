package sdl

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

// TODO: SysWMInfo
// TODO: RendererInfo
// TODO: AudioCVT
func TestStructABI(t *testing.T) {
	var tests = []struct {
		gStruct interface{}
		cStruct interface{}
	}{
		{AudioSpec{}, cAudioSpec{}},
		{DisplayMode{}, cDisplayMode{}},
		{Palette{}, cPalette{}},
		{PixelFormat{}, cPixelFormat{}},
		{Surface{}, cSurface{}},
		{Version{}, cVersion{}},
		{WindowEvent{}, cWindowEvent{}},

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
	}

	for _, test := range tests {
		testABI(t, test.gStruct, test.cStruct)
	}
}

// TODO: mixer.MusicType
// TODO: mixer.Fading
func TestTypeABI(t *testing.T) {
	var tests = []struct {
		gType interface{}
		cType interface{}
	}{
		{AudioStatus(0), cAudioStatus(0)},
		{ErrorCode(0), cErrorCode(0)},
		{RendererFlip(0), cRendererFlip(0)},
	}

	for _, test := range tests {
		testABI(t, test.gType, test.cType)
	}
}

func testABI(t *testing.T, a, b interface{}) {
	ta, tb := reflect.TypeOf(a), reflect.TypeOf(b)
	if ta.Size() != tb.Size() {
		t.Fatalf("type size missmatch: %s(%d) != %s(%d)",
			ta.Name(), ta.Size(), tb.Name(), tb.Size())
	}
	if ta.Kind() != tb.Kind() {
		t.Fatalf("type kind missmatch: %s=%s != %s=%s",
			ta.Name(), ta.Kind(), tb.Name(), tb.Kind())
	}

	if ta.Kind() == reflect.Struct {
		for i := 0; i < ta.NumField(); i++ {
			f := ta.Field(i)
			if f.Name == "_" {
				// ignore padding fields
				continue
			}
			if err := verifyField(t, f, tb); err != nil {
				dumpStructFormat(t, ta)
				dumpStructFormat(t, tb)
				t.Error(err)
			}
		}
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

func verifyField(t *testing.T, gField reflect.StructField, cType reflect.Type) error {
	var cField = searchField(cType, gField.Name)
	if cField == nil {
		return fmt.Errorf("field not found: %q", gField.Name)
	}
	gOffset, gSize := gField.Offset, gField.Type.Size()
	cOffset, cSize := cField.Offset, cField.Type.Size()
	if cOffset != gOffset || cSize != gSize {
		return fmt.Errorf("field offset/size missmatch %s(%d, %d) != %s(%d, %d)",
			gField.Name, gOffset, gSize,
			cField.Name, cOffset, cSize)
	}
	return nil
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
