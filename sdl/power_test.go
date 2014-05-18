package sdl

import "testing"

func TestGetPowerInfo(t *testing.T) {
	state, seconds, percent := GetPowerInfo()
	strstate := "Unknown"
	minutes := -1
	if state == POWERSTATE_CHARGED {
		strstate = "Battery charged"
	} else if state == POWERSTATE_CHARGING {
		strstate = "Battery charging"
	} else if state == POWERSTATE_NO_BATTERY {
		strstate = "No battery"
	} else if state == POWERSTATE_ON_BATTERY {
		strstate = "On battery"
	}
	if seconds != -1 {
		minutes = seconds / 60
	}
	t.Log("Power status report, please check manually")
	t.Logf("Status: %s", strstate)
	t.Logf("Minutes left (-1 = undetermined): %d", minutes)
	t.Logf("Percent left (-1 = undetermined): %d", percent)
}
