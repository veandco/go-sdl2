package sdl

import (
	"bytes"
	"testing"
)

var squareWave = []byte("RIFF,\x00\x00\x00WAVEfmt \x10\x00\x00\x00\x01\x00" +
	"\x01\x00\xab \x00\x00\xab \x00\x00\x01\x00\b\x00data\b\x00\x00\x00\x00" +
	"\x00\x00\x00\xff\xff\xff\xff")

func TestAudioDevices(t *testing.T) {
	for _, isCapture := range []bool{false, true} {
		n := GetNumAudioDevices(isCapture)
		for i := 0; i < n; i++ {
			if GetAudioDeviceName(i, isCapture) == "" {
				t.Errorf("GetAudioDeviceName(%v, %v) returned empty string",
					i, isCapture)
			}
		}

		// cover GetAudioDeviceName() even if n < 1
		if name := GetAudioDeviceName(-1, isCapture); name != "" {
			t.Errorf("GetAudioDeviceName(-1, %v) == %#v; want empty string",
				isCapture, name)
		}
	}
}

func TestAudioInitQuit(t *testing.T) {
	Do(func() {
		// figure out what driver will work
		if err := InitSubSystem(INIT_AUDIO); err != nil {
			t.Fatalf("InitSubSystem(INIT_AUDIO): %v", err)
		}
		driver := GetCurrentAudioDriver()

		// reset audio subsystem
		QuitSubSystem(INIT_AUDIO)
		if GetCurrentAudioDriver() != "" {
			t.Fatalf(`GetCurrentAudioDriver() != "" after QuitSubSystem()`)
		}

		// test valid init
		if err := AudioInit(driver); err != nil {
			t.Errorf("AudioInit(%s) returned error: %v", driver, err)
		} else {
			if got := GetCurrentAudioDriver(); got != driver {
				t.Errorf("GetCurrentAudioDriver() == %#v; want %#v", got, driver)
			}

			// test quit
			AudioQuit()
			if GetCurrentAudioDriver() != "" {
				t.Fatalf(`GetCurrentAudioDriver() != "" after AudioQuit()`)
			}
		}

		// test invalid init
		driver = "bogus"
		if err := AudioInit(driver); err == nil {
			t.Errorf("AudioInit(%s) did not return error", driver)
		}
	})
}

func TestLoadWAVRW(t *testing.T) {
	// load WAV from *RWOps pointing to WAV data
	src, _ := RWFromMem(squareWave)
	buf, spec := LoadWAVRW(src, false)

	// test returned []byte
	want := []byte{0, 0, 0, 0, 255, 255, 255, 255}
	if buf == nil {
		t.Errorf("LoadWAVRW() returned nil []byte")
	} else {
		if bytes.Compare(buf, want) != 0 {
			t.Errorf("LoadWAVRW() returned %v; want %v", buf, want)
		}
		FreeWAV(buf) // make sure we can free without error
	}

	// test returned *AudioSpec
	if spec == nil {
		t.Errorf("LoadWAVRW() returned nil *AudioSpec")
	} else if spec.Freq != 8363 { // no need to test all the spec fields
		t.Errorf("LoadWAVRW() returned *AudioSpec with Freq %d; want %d",
			spec.Freq, 8363)
	}
}
