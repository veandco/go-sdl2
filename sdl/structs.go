package sdl

import "unsafe"

type DisplayMode struct {
	Format uint32
	W int
	H int
	RefreshRate int
	DriverData unsafe.Pointer
}

type Point struct {
	X int32
	Y int32
}

type Rect struct {
	X int32
	Y int32
	W int32
	H int32
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

type Version struct {
	Major uint8
	Minor uint8
	Patch uint8
}

/* SDL_render.h */
type RendererInfo struct {
	Name string
	Flags uint32
	NumTextureFormats uint32
	TextureFormats [16]int32
	MaxTextureWidth int
	MaxTextureHeight int
}

/* SDL_sysrender.h */
type FPoint struct {
	x float32
	y float32
}

type FRect struct {
	x float32
	y float32
	w float32
	h float32
}

type Texture struct {
	Magic unsafe.Pointer
	Format uint32
	Access int
	W int
	H int
	ModMode int
	BlendMode BlendMode
	R uint8
	G uint8
	B uint8
	A uint8
	Renderer Renderer
	Native *Texture
	Yuv *SW_YUVTexture
	Pixels unsafe.Pointer
	Pitch int
	LockedRect Rect
	DriverData Rect
	Prev *Texture
	Next *Texture
}

type Renderer struct {
	magic unsafe.Pointer
	WindowEvent unsafe.Pointer
	GetOutputSize unsafe.Pointer
	CreateTexture unsafe.Pointer
	SetTextureColorMod unsafe.Pointer
	SetTextureAlphaMod unsafe.Pointer
	UpdateTexture unsafe.Pointer
	LockTexture unsafe.Pointer
	UnlockTexture unsafe.Pointer
	_SetRenderTarget unsafe.Pointer
	UpdateViewport unsafe.Pointer
	UpdateClipRect unsafe.Pointer
	RenderClear unsafe.Pointer
	RenderDrawPoints unsafe.Pointer
	RenderDrawLines unsafe.Pointer
	RenderFillRects unsafe.Pointer
	RenderCopy unsafe.Pointer
	RenderCopyEx unsafe.Pointer
	RenderReadPixels unsafe.Pointer
	RenderPresent unsafe.Pointer
	DestroyTexture unsafe.Pointer
	DestroyRenderer unsafe.Pointer
	GL_BindTexture unsafe.Pointer
	GL_UnbindTexture unsafe.Pointer
	Info RendererInfo
	Window *Window
	Hidden bool
	LogicalW int
	LogicalH int
	LogicalWBackup int
	LogicalHBackup int
	Viewport Rect
	ViewportBackup Rect
	ClipRect Rect
	ClipRectBackup Rect
	Scale FPoint
	ScaleBackup FPoint
	Textures *Texture
	Targets *Texture
	R uint8
	G uint8
	B uint8
	A uint8
	BlendMode BlendMode
	DriverData unsafe.Pointer
}

/* SDL_yuv_sw_c.h */
type SW_YUVTexture struct {
	Format uint32
	TargetFormat uint32
	W int
	H int
	Pixels *uint8
	ColorTab *int
	Rgb2Pix *uint32
	Display1X unsafe.Pointer
	Display2x unsafe.Pointer
	Pitches [3]uint16
	Planes *[3]uint8
	Stretch *Surface
	Display *Surface
}
