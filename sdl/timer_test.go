package sdl

import (
	"testing"
	"time"
)

// Both time.Sleep() and Delay() are only guaranteed to wait *at least* the
// given duration, so all tests only check that time differentials are at least
// a certain value.

func TestGetTicks(t *testing.T) {
	before := GetTicks()
	ms := 1
	time.Sleep(time.Duration(ms) * time.Millisecond)
	after := GetTicks()
	if after-before < uint32(ms) {
		t.Errorf("GetTicks(): returned %d, then %d after %d ms",
			before, after, ms)
	}
}

func TestGetPerformance(t *testing.T) {
	freq := GetPerformanceFrequency()
	if freq <= 0 {
		t.Errorf("GetPerformanceFrequency(): returned %d", freq)
	}
	before := GetPerformanceCounter()
	ms := 10
	time.Sleep(time.Duration(ms) * time.Millisecond)
	after := GetPerformanceCounter()
	if after-before < freq*uint64(ms)/1000 {
		t.Errorf("GetPerformanceCounter(): returned %d, then %d after %d ms",
			before, after, ms)
	}
}

func TestDelay(t *testing.T) {
	for msDelay := 0; msDelay <= 20; msDelay += 10 {
		before := time.Now()
		Delay(uint32(msDelay))
		msSince := time.Since(before) / time.Millisecond
		if int64(msSince) < int64(msDelay) {
			t.Errorf("Delay(%d): delayed for only %d ms", msDelay, msSince)
		}
	}
}
