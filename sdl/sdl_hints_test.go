package sdl

import "testing"

func TestSetHintWithPriority(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	if !SetHintWithPriority("TEST", "32", HINT_DEFAULT) {
		t.Errorf("return value for SetHintWithPriority('TEST', '32', HINT_DEFAULT) is wrong")
	}
	if GetHint("TEST") != "32" {
		t.Errorf("SetHintWithPriority() did not set value properly")
	}
	if !SetHintWithPriority("TEST", "abc", HINT_NORMAL) {
		t.Errorf("return value for SetHintWithPriority('TEST', 'abc', HINT_NORMAL) is wrong")
	}
	if GetHint("TEST") != "abc" {
		t.Errorf("SetHintWithPriority() [HINT_NORMAL] did not set value properly")
	}
	if SetHintWithPriority("TEST", "xyz", HINT_DEFAULT) {
		t.Errorf("return value for SetHintWithPriority('TEST', 'xyz', HINT_DEFAULT) is wrong")
	}
	if GetHint("TEST") != "abc" {
		t.Errorf("SetHintWithPriority() [HINT_DEFAULT] did not set value properly")
	}

}

func TestGetSetHint(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	if !SetHint("TEST", "32") {
		t.Errorf("return value for SetHint('TEST', '32') is wrong")
	}
	if GetHint("TEST") != "32" {
		t.Errorf("return value for GetHint('TEST') is wrong")
	}
	if !SetHint(HINT_RENDER_DRIVER, "dummy") {
		t.Errorf("return value for SetHint(HINT_RENDER_DRIVER, 'dummy') is wrong")
	}
	if GetHint(HINT_RENDER_DRIVER) != "dummy" {
		t.Errorf("return value for GetHint(HINT_RENDER_DRIVER) is wrong")
	}
}

func TestClearHints(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	if !SetHint("TEST", "32") {
		t.Errorf("return value for SetHint('TEST', '32') is wrong")
	}
	if GetHint("TEST") != "32" {
		t.Errorf("return value for GetHint('TEST') is wrong")
	}
	ClearHints()
	if GetHint("TEST") != "" {
		t.Errorf("return value for GetHint('TEST') after ClearHints() is wrong")
	}
}
