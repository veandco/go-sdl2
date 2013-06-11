package sdl

import "unsafe"

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


