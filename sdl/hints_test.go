package sdl

import "testing"

const (
	TIMER_RESOLUTION = "32"
	DRIVER           = "software"
)

func TestSetHintWithPriority(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	if !SetHintWithPriority(HINT_TIMER_RESOLUTION, TIMER_RESOLUTION, HINT_DEFAULT) {
		t.Errorf("return value for SetHintWithPriority(HINT_TIMER_RESOLUTION, TIMER_RESOLUTION, HINT_DEFAULT) is wrong")
	}
	if GetHint(HINT_TIMER_RESOLUTION) != TIMER_RESOLUTION {
		t.Errorf("SetHintWithPriority() did not set value properly")
	}
	if !SetHintWithPriority(HINT_RENDER_DRIVER, DRIVER, HINT_NORMAL) {
		t.Errorf("return value for SetHintWithPriority(HINT_RENDER_DRIVER, 'software', HINT_NORMAL) is wrong")
	}
	if GetHint(HINT_RENDER_DRIVER) != DRIVER {
		t.Errorf("SetHintWithPriority() [HINT_NORMAL] did not set value properly")
	}
	if SetHintWithPriority(HINT_RENDER_DRIVER, DRIVER, HINT_DEFAULT) {
		t.Errorf("return value for SetHintWithPriority(HINT_RENDER_DRIVER, DRIVER, HINT_DEFAULT) is wrong")
	}
	if GetHint(HINT_RENDER_DRIVER) != DRIVER {
		t.Errorf("SetHintWithPriority() [HINT_DEFAULT] did not set value properly")
	}

}

func TestGetSetHint(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	if !SetHint(HINT_TIMER_RESOLUTION, TIMER_RESOLUTION) {
		t.Errorf("return value for SetHint(HINT_TIMER_RESOLUTION, RESOLUTION) is wrong")
	}
	if GetHint(HINT_TIMER_RESOLUTION) != TIMER_RESOLUTION {
		t.Errorf("return value for GetHint(HINT_TIMER_RESOLUTION) is wrong")
	}

	if !SetHint(HINT_RENDER_DRIVER, DRIVER) {
		t.Errorf("return value for SetHint(HINT_RENDER_DRIVER, DRIVER) is wrong")
	}
	if GetHint(HINT_RENDER_DRIVER) != DRIVER {
		t.Errorf("return value for GetHint(HINT_RENDER_DRIVER) is wrong")
	}
}

func TestClearHints(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	if !SetHint(HINT_TIMER_RESOLUTION, TIMER_RESOLUTION) {
		t.Errorf("return value for SetHint(HINT_TIMER_RESOLUTION, TIMER_RESOLUTION) is wrong")
	}
	if GetHint(HINT_TIMER_RESOLUTION) != TIMER_RESOLUTION {
		t.Errorf("return value for GetHint(HINT_TIMER_RESOLUTION) is wrong")
	}
	ClearHints()
	if GetHint(HINT_TIMER_RESOLUTION) != "" {
		t.Errorf("return value for GetHint(HINT_TIMER_RESOLUTION) after ClearHints() is wrong")
	}
}
