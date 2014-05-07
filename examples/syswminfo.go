// author: Jacky Boen

package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"fmt"
	"os"
)

func main() {
	var window *sdl.Window
	var info sdl.SysWMInfo
	var subsystem string

	window = sdl.CreateWindow("", 0, 0, 0, 0, sdl.WINDOW_HIDDEN)
	if window == nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", sdl.GetError())
		os.Exit(1)
	}

	sdl.VERSION(&info.Version)

	if window.GetWMInfo(&info) {
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
		fmt.Fprintf(os.Stderr, "Couldn't get window information: %s\n", sdl.GetError())
	}

	window.Destroy()
}
