package sdl

/*
#include "sdl_wrapper.h"
static Sint64 RWsize(SDL_RWops *ctx)
{
	return ctx->size(ctx);
}

static Sint64 RWseek(SDL_RWops *ctx, Sint64 offset, int whence)
{
	return ctx->seek(ctx, offset, whence);
}

static size_t RWread(SDL_RWops *ctx, void *ptr, size_t size, size_t maxnum)
{
	return ctx->read(ctx, ptr, size, maxnum);
}

static size_t RWwrite(SDL_RWops *ctx, void *ptr, size_t size, size_t num)
{
	return ctx->write(ctx, ptr, size, num);
}

static int RWclose(SDL_RWops *ctx)
{
	return ctx->close(ctx);
}
*/
import "C"
import "unsafe"

/* RWops Types */
const (
	RWOPS_UNKNOWN   = 0
	RWOPS_WINFILE   = 1
	RWOPS_STDFILE   = 2
	RWOPS_JNIFILE   = 3
	RWOPS_MEMORY    = 4
	RWOPS_MEMORY_RO = 5
)

const (
	RW_SEEK_SET = C.RW_SEEK_SET
	RW_SEEK_CUR = C.RW_SEEK_CUR
	RW_SEEK_END = C.RW_SEEK_END
)

// RWops (https://wiki.libsdl.org/SDL_RWops)
type RWops C.SDL_RWops

func (rw *RWops) cptr() *C.SDL_RWops {
	return (*C.SDL_RWops)(rw)
}

// RWFromFile (https://wiki.libsdl.org/SDL_RWFromFile)
func RWFromFile(file, mode string) *RWops {
	_file := C.CString(file)
	_mode := C.CString(mode)
	defer C.free(unsafe.Pointer(_file))
	defer C.free(unsafe.Pointer(_mode))
	return (*RWops)(unsafe.Pointer(C.SDL_RWFromFile(_file, _mode)))
}

// RWFromMem (https://wiki.libsdl.org/SDL_RWFromMem)
func RWFromMem(mem unsafe.Pointer, size int) *RWops {
	if mem == nil {
		return nil
	}
	return (*RWops)(unsafe.Pointer(C.SDL_RWFromMem(mem, C.int(size))))
}

// AllocRW (https://wiki.libsdl.org/SDL_AllocRW)
func AllocRW() *RWops {
	return (*RWops)(unsafe.Pointer(C.SDL_AllocRW()))
}

// FreeRW (https://wiki.libsdl.org/SDL_FreeRW)
func (area *RWops) FreeRW() {
	if area == nil {
		return
	}
	C.SDL_FreeRW(area.cptr())
}

// RWsize (https://wiki.libsdl.org/SDL_RWsize)
func (ctx *RWops) RWsize() int64 {
	return int64(C.RWsize(ctx.cptr()))
}

// RWseek (https://wiki.libsdl.org/SDL_RWseek)
func (ctx *RWops) RWseek(offset int64, whence int) int64 {
	if ctx == nil {
		return -1
	}
	return int64(C.RWseek(ctx.cptr(), C.Sint64(offset), C.int(whence)))
}

// RWread (https://wiki.libsdl.org/SDL_RWread)
func (ctx *RWops) RWread(ptr unsafe.Pointer, size, maxnum uint) uint {
	if ctx == nil {
		return 0
	}
	return uint(C.RWread(ctx.cptr(), ptr, C.size_t(size), C.size_t(maxnum)))
}

// RWtell (https://wiki.libsdl.org/SDL_RWtell)
func (ctx *RWops) RWtell() int64 {
	if ctx == nil {
		return 0
	}
	return int64(C.RWseek(ctx.cptr(), 0, RW_SEEK_CUR))
}

// RWwrite (https://wiki.libsdl.org/SDL_RWwrite)
func (ctx *RWops) RWwrite(ptr unsafe.Pointer, size, num uint) uint {
	if ctx == nil {
		return 0
	}
	if ptr == nil {
		return 0
	}
	return uint(C.RWwrite(ctx.cptr(), ptr, C.size_t(size), C.size_t(size)))
}

// RWclose (https://wiki.libsdl.org/SDL_RWclose)
func (ctx *RWops) RWclose() error {
	if ctx != nil && C.RWclose(ctx.cptr()) != 0 {
		return GetError()
	}
	return nil
}

func (src *RWops) ReadU8() uint8 {
	if src == nil {
		return 0
	}
	return uint8(C.SDL_ReadU8(src.cptr()))
}

// ReadLE16 (https://wiki.libsdl.org/SDL_ReadLE16)
func (src *RWops) ReadLE16() uint16 {
	if src == nil {
		return 0
	}
	return uint16(C.SDL_ReadLE16(src.cptr()))
}

// ReadBE16 (https://wiki.libsdl.org/SDL_ReadBE16)
func (src *RWops) ReadBE16() uint16 {
	if src == nil {
		return 0
	}
	return uint16(C.SDL_ReadBE16(src.cptr()))
}

// ReadLE32 (https://wiki.libsdl.org/SDL_ReadLE32)
func (src *RWops) ReadLE32() uint32 {
	if src == nil {
		return 0
	}
	return uint32(C.SDL_ReadLE32(src.cptr()))
}

// ReadBE32 (https://wiki.libsdl.org/SDL_ReadBE32)
func (src *RWops) ReadBE32() uint32 {
	if src == nil {
		return 0
	}
	return uint32(C.SDL_ReadBE32(src.cptr()))
}

// ReadLE64 (https://wiki.libsdl.org/SDL_ReadLE64)
func (src *RWops) ReadLE64() uint64 {
	if src == nil {
		return 0
	}
	return uint64(C.SDL_ReadLE64(src.cptr()))
}

// ReadBE64 (https://wiki.libsdl.org/SDL_ReadBE64)
func (src *RWops) ReadBE64() uint64 {
	if src == nil {
		return 0
	}
	return uint64(C.SDL_ReadBE64(src.cptr()))
}

func (dst *RWops) WriteU8(value uint8) uint {
	if dst == nil {
		return 0
	}
	return uint(C.SDL_WriteU8(dst.cptr(), C.Uint8(value)))
}

// WriteLE16 (https://wiki.libsdl.org/SDL_WriteLE16)
func (dst *RWops) WriteLE16(value uint16) uint {
	if dst == nil {
		return 0
	}
	return uint(C.SDL_WriteLE16(dst.cptr(), C.Uint16(value)))
}

// WriteBE16 (https://wiki.libsdl.org/SDL_WriteBE16)
func (dst *RWops) WriteBE16(value uint16) uint {
	if dst == nil {
		return 0
	}
	return uint(C.SDL_WriteBE16(dst.cptr(), C.Uint16(value)))
}

// WriteLE32 (https://wiki.libsdl.org/SDL_WriteLE32)
func (dst *RWops) WriteLE32(value uint32) uint {
	if dst == nil {
		return 0
	}
	return uint(C.SDL_WriteLE32(dst.cptr(), C.Uint32(value)))
}

// WriteBE32 (https://wiki.libsdl.org/SDL_WriteBE32)
func (dst *RWops) WriteBE32(value uint32) uint {
	if dst == nil {
		return 0
	}
	return uint(C.SDL_WriteBE32(dst.cptr(), C.Uint32(value)))
}

// WriteLE64 (https://wiki.libsdl.org/SDL_WriteLE64)
func (dst *RWops) WriteLE64(value uint64) uint {
	if dst == nil {
		return 0
	}
	return uint(C.SDL_WriteLE64(dst.cptr(), C.Uint64(value)))
}

// WriteBE64 (https://wiki.libsdl.org/SDL_WriteBE64)
func (dst *RWops) WriteBE64(value uint64) uint {
	if dst == nil {
		return 0
	}
	return uint(C.SDL_WriteBE64(dst.cptr(), C.Uint64(value)))
}
