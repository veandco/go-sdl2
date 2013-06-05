package sdl

import "unsafe"

type DisplayMode struct {
	Format uint32
	W int
	H int
	RefreshRate int
	DriverData unsafe.Pointer
}

type Rect struct {
	X int
	Y int
	W int
	H int
}

type WindowShaper struct {
	Window *Window
	UserX uint32
	UserY uint32
	Mode *uint
	HasShape bool
	DriverData unsafe.Pointer
}

type WindowUserData struct {
	Name string
	Data unsafe.Pointer
	Next *WindowUserData
}

type Surface struct {
	Flags uint32
	Format PixelFormat
	W int
	H int
	Pitch int
	Pixels unsafe.Pointer
	UserData unsafe.Pointer
	Locked int
	LockData unsafe.Pointer
	ClipRect Rect
	_map *[0]byte
	RefCount int
}

type PixelFormat struct {
	Format uint32
	Palette *Palette
	BitsPerPixels uint8
	BytesPerPixel uint8
	Padding [2]uint8
	Rmask uint32
	Gmask uint32
	Bmask uint32
	Amask uint32
	Rloss uint8
	Gloss uint8
	Bloss uint8
	Aloss uint8
	Rshift uint8
	Gshift uint8
	Bshift uint8
	Ashift uint8
	RefCount int
	next *PixelFormat
}

type Palette struct {
	Ncolors int
	Colors *Color
	Version uint32
	RefCount int
}

type Color struct {
	r uint8
	g uint8
	b uint8
	a uint8
}

type Window struct {
	Magic unsafe.Pointer
	Id uint32
	Title string
	X int
	Y int
	W int
	H int
	MinW int
	Minh int
	MaxW int
	MaxH int
	Flags uint32
	Windowed Rect
	FullscreenMode DisplayMode
	Brightness float64
	gamma *uint16
	SavedGamma *uint16
	Surface Surface
	SurfaceValid bool
	Shaper WindowShaper
	Data WindowUserData
	DriverData unsafe.Pointer
	Prev *Window
	Next *Window
}

type Version struct {
	Major uint8
	Minor uint8
	Patch uint8
}
