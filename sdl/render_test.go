package sdl

import (
	"errors"
	"testing"
)

func TestRenderIntegerScale(t *testing.T) {
	_, r, err := CreateWindowAndRenderer(50, 50, 0)
	if err != nil {
		t.Fatalf("unable to create renderer: %s", err)
		return
	}

	errUnsupported := errors.New("That operation is not supported")

	if !VERSION_ATLEAST(2, 0, 5) {
		t.Run("GetIntegerScale on versions < 2.0.5 should return unsupported", func(t *testing.T) {
			_, got := r.GetIntegerScale()
			if got.Error() != errUnsupported.Error() {
				t.Fatalf("expected call to GetIntegerScale to fail with %s, got %s", errUnsupported, got)
			}
		})
		t.Run("SetIntegerScale on versions < 2.0.5 should return unsupported", func(t *testing.T) {
			got := r.SetIntegerScale(false)
			if got.Error() != errUnsupported.Error() {
				t.Fatalf("expected call to SetIntegerScale to fail with %s, got %s", errUnsupported, got)
			}
		})

		return
	}

	set := func(v bool, t *testing.T) {
		if err := r.SetIntegerScale(v); err != nil {
			t.Fatalf("unable to set scale %v: %s", v, err)
		}
	}
	check := func(want bool, t *testing.T) {
		got, err := r.GetIntegerScale()
		if err != nil {
			t.Fatalf("unable to get scale: %s", err)
		}

		if got != want {
			t.Fatalf("wanted renderer scale to be %v", want)
		}
	}

	t.Run("should start out as false", func(t *testing.T) {
		check(false, t)
	})

	t.Run("if set to true should return true", func(t *testing.T) {
		set(true, t)
		check(true, t)
	})

	t.Run("if set to false should return false", func(t *testing.T) {
		set(true, t)
		check(true, t)

		set(false, t)
		check(false, t)
	})

	t.Run("should handle SDL_FALSE return on GetIntegerScale", func(t *testing.T) {
		_, dummy, err := CreateWindowAndRenderer(50, 50, 0)
		if err != nil {
			t.Fatalf("unable to create renderer: %s", err)
			return
		}

		if err := dummy.SetIntegerScale(true); err != nil {
			t.Fatalf("expected call to SetIntegerScale to succeed, got %s", err)
		}
		if err := dummy.Destroy(); err != nil {
			t.Fatalf("expected call to Destroy to succeed, got %s", err)
		}

		v, err := dummy.GetIntegerScale()
		if err == nil {
			t.Fatalf("expected call to GetIntegerScale on ivalid renderer to fail")
		}
		if v != false {
			t.Fatalf("wanted renderer scale to be %v", false)
		}
	})
}
