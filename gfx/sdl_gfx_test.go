package gfx

import (
	"testing"
)

func TestMin(t *testing.T) {
	if a := []int{4, 2, 3, 5, 1, 6, 5, 7}; min(a...) != 1 {
		t.Fail()
	}

	if a := []int{5, 6, -2, 1, 8, 1, 4, 3}; min(a...) != -2 {
		t.Fail()
	}
}
