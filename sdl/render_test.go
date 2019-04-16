package sdl

import "testing"

func TestRenderIntegerScale(t *testing.T) {
	_, r, err := CreateWindowAndRenderer(50, 50, 0)
	if err != nil {
		t.Fatalf("unable to create renderer: %s", err)
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
}
