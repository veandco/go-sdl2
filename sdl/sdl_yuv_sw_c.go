package sdl

import "unsafe"

type SW_YUVTexture struct {
	Format       uint32
	TargetFormat uint32
	W            int
	H            int
	Pixels       *uint8
	ColorTab     *int
	Rgb2Pix      *uint32
	Display1X    unsafe.Pointer
	Display2x    unsafe.Pointer
	Pitches      [3]uint16
	Planes       *[3]uint8
	Stretch      *Surface
	Display      *Surface
}
