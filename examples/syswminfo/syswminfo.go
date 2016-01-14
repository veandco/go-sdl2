// author: Jacky Boen

package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

func main() {
	var window *sdl.Window
	var info *sdl.SysWMInfo
	var subsystem string
	var err error

	window, err = sdl.CreateWindow("", 0, 0, 0, 0, sdl.WINDOW_HIDDEN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		os.Exit(1)
	}
	defer window.Destroy()

	info, err = window.GetWMInfo()
	if err == nil {
		switch info.Subsystem {
		case sdl.SYSWM_UNKNOWN:
			subsystem = "An unknown system!"
		case sdl.SYSWM_WINDOWS:
			subsystem = "Microsoft Windows(TM)"
		case sdl.SYSWM_X11:
			subsystem = "X Window System"
		case sdl.SYSWM_DIRECTFB:
			subsystem = "DirectFB"
		case sdl.SYSWM_COCOA:
			subsystem = "Apple OS X"
		case sdl.SYSWM_UIKIT:
			subsystem = "UIKit"
		}

		fmt.Printf("This program is running SDL version %d.%d.%d on %s\n",
			info.Version.Major,
			info.Version.Minor,
			info.Version.Patch,
			subsystem)
	} else {
		fmt.Fprintf(os.Stderr, "Couldn't get window information: %s\n", err)
	}
}
