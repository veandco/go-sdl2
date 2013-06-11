package sdl

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
