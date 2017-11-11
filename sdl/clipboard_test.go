package sdl

import "testing"

func TestClipboard(t *testing.T) {
	Do(func() {
		Init(INIT_VIDEO)
		defer Quit()

		// need a window to own the clipboard
		if win, err := CreateWindow("", 0, 0, 0, 0, 0); err == nil {
			defer win.Destroy()
		} else {
			t.Skipf("Could not CreateWindow(): %v", err)
		}

		wantString, wantBool := []string{"", "naïveté"}, []bool{false, true}
		for i := 0; i < 2; i++ {
			if err := SetClipboardText(wantString[i]); err != nil {
				t.Skipf("Could not SetClipboardText(): %v", err)
			}
			if got := HasClipboardText(); got != wantBool[i] {
				t.Errorf("HasClipboardText() == %v; want %v", got, wantBool[i])
			}
			if got, err := GetClipboardText(); err == nil && got != wantString[i] {
				t.Errorf("GetClipboardText() == (nil, %#v); want (nil, %#v)",
					got, wantString[i])
			}
		}
	})
}
