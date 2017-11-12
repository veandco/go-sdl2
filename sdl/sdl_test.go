package sdl

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func init() {
	if !VERSION_ATLEAST(2, 0, 4) {
		fmt.Printf("Warning: Your SDL version (%d.%d.%d) is outdated and may fail to pass tests written for 2.0.4\n", MAJOR_VERSION, MINOR_VERSION, PATCHLEVEL)
	}
}

/* Set up main thread during testing to run any thread-stubborn test code.
 * This is necessary as Go runs test functions on different threads.
 * Examples of functions that are best run in main:
 *     Init, Quit, PollEvent, WaitEvent
 * Failure to run such functions in main during testing may produce insidious,
 * hard-to-reproduce bugs.
 */
func TestMain(m *testing.M) {
	var exitcode int
	Main(func() {
		exitcode = m.Run()
	})
	os.Exit(exitcode)
}

func TestInitQuit(t *testing.T) {
	Do(func() {
		Init(0)
		subs := []uint32{INIT_TIMER, INIT_AUDIO, INIT_VIDEO, INIT_JOYSTICK,
			INIT_HAPTIC, INIT_GAMECONTROLLER}

		for i := 0; i < len(subs); i++ {
			if (runtime.GOOS == "freebsd") && (subs[i] == INIT_HAPTIC) {
				// FreeBSD does not support the haptic subsystem
				continue
			}
			if err := Init(subs[i]); err != nil {
				t.Errorf("Error on Init(%d): %s", subs[i], err)
			}
			if WasInit(subs[i]) != subs[i] {
				t.Errorf("Init(%d): subsystem not initialized", subs[i])
			}
			QuitSubSystem(subs[i])
			if WasInit(subs[i]) == subs[i] {
				t.Errorf("QuitSubSystem(%d): subsystem still initialized", subs[i])
			}
			if err := InitSubSystem(subs[i]); err != nil {
				t.Errorf("Error on Init(%d): %s", subs[i], err)
			}
			if WasInit(subs[i]) != subs[i] {
				t.Errorf("InitSubSystem(%d): subsystem not initialized", subs[i])
			}
			QuitSubSystem(subs[i])
			if WasInit(subs[i]) == subs[i] {
				t.Errorf("QuitSubSystem(%d): subsystem still initialized", subs[i])
			}
		}
		Quit()
	})
}

func TestGetPlatform(t *testing.T) {
	goos := runtime.GOOS
	value := GetPlatform()
	if (goos == "linux" && value != "Linux") ||
		(goos == "freebsd" && value != "FreeBSD") ||
		(goos == "windows" && value != "Windows") ||
		(goos == "darwin" && value != "Mac OS X") {
		t.Errorf("platform mismatch: '%s' != '%s'", goos, value)
	}
}
