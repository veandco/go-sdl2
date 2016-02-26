package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	buttons := []sdl.MessageBoxButtonData{
		{ 0, 0, "no" },
		{ sdl.MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT, 1, "yes" },
		{ sdl.MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT, 2, "cancel" },
	}

	colorScheme := sdl.MessageBoxColorScheme{
		Colors: [5]sdl.MessageBoxColor{
			sdl.MessageBoxColor{ 255,   0,   0 },
			sdl.MessageBoxColor{   0, 255,   0 },
			sdl.MessageBoxColor{ 255, 255,   0 },
			sdl.MessageBoxColor{   0,   0, 255 },
			sdl.MessageBoxColor{ 255,   0, 255 },
		},
	}

	messageboxdata := sdl.MessageBoxData{
		sdl.MESSAGEBOX_INFORMATION,
		nil,
		"example message box",
		"select a button",
		int32(len(buttons)),
		buttons,
		&colorScheme,
	}

	var buttonid int32
	var err error
	if err, buttonid = sdl.ShowMessageBox(&messageboxdata); err != nil {
		fmt.Println("error displaying message box")
		return
	}

	if buttonid == -1 {
		fmt.Println("no selection")
	} else {
		fmt.Println("selection was", buttons[buttonid].Text)
	}
}
