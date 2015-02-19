package sdl

import (
	"reflect"
	"strings"
	"testing"
)

func TestVersionSize(t *testing.T)     { testStructEquality(t, Version{}, cVersion{}) }
func TestPixelFormatSize(t *testing.T) { testStructEquality(t, PixelFormat{}, cPixelFormat{}) }

//func TestSysWmInfoSize(t *testing.T) { testStructEquality(t, SysWMInfo{}, cSysWMinfo{}) }
// Surface
// WindowEvent

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
	const format = "%25s%10d%10d"
	t.Logf("%25s%10s%10s", "name", "offset", "size")
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if f.Name == "_" {
			// ignore padding fields
			continue
		}
		t.Logf(format, s.Name()+"."+f.Name, f.Offset, f.Type.Size())
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
