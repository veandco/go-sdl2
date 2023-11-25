package sdl

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type EdgeCaseFunc func(t *testing.T, a, b interface{})

func structMatch(t *testing.T, ta, tb reflect.Type) {
	structMatchGeneric(t, ta, tb, ta.Size(), tb.Size())
}

func structMatchFixed(t *testing.T, ta, tb reflect.Type, szb uintptr) {
	structMatchGeneric(t, ta, tb, ta.Size(), szb)
}

func structMatchGeneric(t *testing.T, ta, tb reflect.Type, sza, szb uintptr) {
	if sza != szb {
		t.Errorf("type size mismatch: %s(%d) != %s(%d)",
			ta.Name(), ta.Size(), tb.Name(), tb.Size())
		t.Fail()
	}

	if ta.Kind() != tb.Kind() {
		t.Errorf("type kind mismatch: %s=%s != %s=%s",
			ta.Name(), ta.Kind(), tb.Name(), tb.Kind())
		t.Fail()
	}
}

func mouseWheelEventEdgeCase(t *testing.T, a, b interface{}) {
	ta, tb := reflect.TypeOf(a), reflect.TypeOf(b)

	if VERSION_ATLEAST(2, 0, 4) {
		structMatchFixed(t, ta, tb, 28)
	} else {
		tf := reflect.TypeOf(MouseWheelEvent{}.Direction)
		structMatchFixed(t, ta, tb, tb.Size()+tf.Size())
	}
}

func dropEventEdgeCase(t *testing.T, a, b interface{}) {
	ta, tb := reflect.TypeOf(a), reflect.TypeOf(b)

	if VERSION_ATLEAST(2, 0, 5) {
		structMatch(t, ta, tb)
	} else {
		structMatchFixed(t, ta, tb, 24)
	}
}

// TODO: SysWMInfo
// TODO: RendererInfo
// TODO: AudioCVT
func TestStructABI(t *testing.T) {
	var tests = []struct {
		gStruct    interface{}
		cStruct    interface{}
		isEdgeCase EdgeCaseFunc
	}{
		{AudioSpec{}, cAudioSpec{}, nil},
		{DisplayMode{}, cDisplayMode{}, nil},
		{Palette{}, cPalette{}, nil},
		{PixelFormat{}, cPixelFormat{}, nil},
		{Surface{}, cSurface{}, nil},
		{Version{}, cVersion{}, nil},
		{WindowEvent{}, cWindowEvent{}, nil},

		// EVENTS
		{KeyboardEvent{}, cKeyboardEvent{}, nil},
		{MouseButtonEvent{}, cMouseButtonEvent{}, nil},
		{MouseMotionEvent{}, cMouseMotionEvent{}, nil},
		{MouseWheelEvent{}, cMouseWheelEvent{}, mouseWheelEventEdgeCase},
		{TextEditingEvent{}, cTextEditingEvent{}, nil},
		{TextInputEvent{}, cTextInputEvent{}, nil},
		{JoyAxisEvent{}, cJoyAxisEvent{}, nil},
		{JoyBallEvent{}, cJoyBallEvent{}, nil},
		{JoyHatEvent{}, cJoyHatEvent{}, nil},
		{JoyButtonEvent{}, cJoyButtonEvent{}, nil},
		{ControllerAxisEvent{}, cControllerAxisEvent{}, nil},
		{ControllerButtonEvent{}, cControllerButtonEvent{}, nil},
		{ControllerDeviceEvent{}, cControllerDeviceEvent{}, nil},
		{AudioDeviceEvent{}, cAudioDeviceEvent{}, nil},
		{TouchFingerEvent{}, cTouchFingerEvent{}, nil},
		{MultiGestureEvent{}, cMultiGestureEvent{}, nil},
		{DollarGestureEvent{}, cDollarGestureEvent{}, nil},
		{DropEvent{}, cDropEvent{}, dropEventEdgeCase},
		{UserEvent{}, cUserEvent{}, nil},
		{SysWMEvent{}, cSysWMEvent{}, nil},
	}

	for _, test := range tests {
		testABI(t, test.gStruct, test.cStruct, test.isEdgeCase)
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
		testABI(t, test.gType, test.cType, nil)
	}
}

func testABI(t *testing.T, a, b interface{}, f EdgeCaseFunc) {
	ta, tb := reflect.TypeOf(a), reflect.TypeOf(b)

	if f != nil {
		f(t, a, b)
	} else {
		structMatch(t, ta, tb)

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
		return fmt.Errorf("field offset/size mismatch %s(%d, %d) != %s(%d, %d)",
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
