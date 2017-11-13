package sdl

import (
	"sync"
	"testing"
)

const (
	TIMER_RESOLUTION = "32"
	DRIVER           = "software"
)

func TestSetHintWithPriority(t *testing.T) {
	Do(func() {
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
	})
}

func TestGetSetHint(t *testing.T) {
	Do(func() {
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
	})
}

func TestClearHints(t *testing.T) {
	Do(func() {
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
	})
}

func TestAddHintCallback(t *testing.T) {
	Do(func() {
		var wg sync.WaitGroup

		AddHintCallback(HINT_ALLOW_TOPMOST, func(data interface{}, name, oldValue, newValue string) {
			t.Log("HINT_ALLOW_TOPMOST:", data, name, oldValue, newValue)
			if newValue != "" {
				if newValue != "1" && data.(int) != 12 {
					t.Fail()
				}
				wg.Done()
			}
		}, 12)
		wg.Add(1)

		AddHintCallback(HINT_RENDER_VSYNC, func(data interface{}, name, oldValue, newValue string) {
			t.Log("HINT_RENDER_VSYNC:", data, name, oldValue, newValue)
			if newValue != "" {
				if newValue != "1" && data.(int) != 34 {
					t.Fail()
				}
				wg.Done()
			}
		}, 34)
		wg.Add(1)

		SetHint(HINT_ALLOW_TOPMOST, "1")
		SetHint(HINT_RENDER_VSYNC, "1")

		wg.Wait()
	})
}

func TestDelHintCallback(t *testing.T) {
	Do(func() {
		AddHintCallback(HINT_ALLOW_TOPMOST, func(data interface{}, name, oldValue, newValue string) {
			if newValue == "2" {
				t.Fail()
			}
		}, nil)

		AddHintCallback(HINT_RENDER_VSYNC, func(data interface{}, name, oldValue, newValue string) {
			if newValue == "2" {
				t.Fail()
			}
		}, nil)

		DelHintCallback(HINT_ALLOW_TOPMOST)
		DelHintCallback(HINT_RENDER_VSYNC)

		SetHint(HINT_ALLOW_TOPMOST, "2")
		SetHint(HINT_RENDER_VSYNC, "2")
	})
}
