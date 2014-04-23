package sdl

import (
	"runtime"
	"testing"
)

func TestCPUCount(t *testing.T) {
	if runtime.NumCPU() != GetCPUCount() {
		t.Error("GetCPUCount(): CPU count mismatch")
	}
}

func TestGetCPUCacheLineSize(t *testing.T) {
	if GetCPUCacheLineSize() == 0 {
		t.Error("GetCPUCacheLineSize(): cache line size wrong")
	}
}

func TestGetSystemRAM(t *testing.T) {
	if GetSystemRAM() <= 0 {
		t.Error("GetSystemRAM(): wrong return value")
	}
}
