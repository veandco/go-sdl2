package sdl

/*
#include <SDL2/SDL_rwops.h>
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

const (
	RW_SEEK_SET = 0
	RW_SEEK_CUR = 1
	RW_SEEK_END = 2
)

type RWops C.SDL_RWops

func RWFromFile(file, mode string) *RWops {
	_file := C.CString(file)
	defer C.free(unsafe.Pointer(_file))
	_mode := C.CString(mode)
	defer C.free(unsafe.Pointer(_mode))
	return (*RWops)(unsafe.Pointer(C.SDL_RWFromFile(_file, _mode)))
}

func RWFromMem(mem unsafe.Pointer, size int) *RWops {
	if mem == nil {
		return nil
	}
	_size := (C.int)(size)
	return (*RWops)(unsafe.Pointer(C.SDL_RWFromMem(mem, _size)))
}

func AllocRW() *RWops {
	return (*RWops)(unsafe.Pointer(C.SDL_AllocRW()))
}

func (area *RWops) FreeRW() {
	if area == nil {
		return
	}
	_area := (*C.SDL_RWops)(area)
	C.SDL_FreeRW(_area)
}

func (ctx *RWops) RWsize() int64 {
	_ctx := (*C.SDL_RWops)(unsafe.Pointer(ctx))
	return (int64)(C.RWsize(_ctx))
}

func (ctx *RWops) RWseek(offset int64, whence int) int64 {
	if ctx == nil {
		return -1
	}
	_ctx := (*C.SDL_RWops)(unsafe.Pointer(ctx))
	_offset := (C.Sint64)(offset)
	_whence := (C.int)(whence)
	return (int64)(C.RWseek(_ctx, _offset, _whence))
}

func (ctx *RWops) RWread(ptr unsafe.Pointer, size, maxnum uint) uint {
	if ctx == nil {
		return 0
	}
	_ctx := (*C.SDL_RWops)(unsafe.Pointer(ctx))
	_size := (C.size_t)(size)
	_maxnum := (C.size_t)(maxnum)
	return (uint)(C.RWread(_ctx, ptr, _size, _maxnum))
}

func (ctx *RWops) RWtell() int64 {
	if ctx == nil {
		return 0
	}
	_ctx := (*C.SDL_RWops)(unsafe.Pointer(ctx))
	return (int64)(C.RWseek(_ctx, 0, RW_SEEK_CUR))
}

func (ctx *RWops) RWwrite(ptr unsafe.Pointer, size, num uint) uint {
	if ctx == nil {
		return 0
	}
	if ptr == nil {
		return 0
	}
	_ctx := (*C.SDL_RWops)(unsafe.Pointer(ctx))
	_size := (C.size_t)(size)
	_num := (C.size_t)(num)
	return (uint)(C.RWwrite(_ctx, ptr, _size, _num))
}

func (ctx *RWops) RWclose() int {
	if ctx == nil {
		return 0
	}
	_ctx := (*C.SDL_RWops)(unsafe.Pointer(ctx))
	return (int)(C.RWclose(_ctx))
}

func (src *RWops) ReadU8() uint8 {
	if src == nil {
		return 0
	}
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return (uint8)(C.SDL_ReadU8(_src))
}

func (src *RWops) ReadLE16() uint16 {
	if src == nil {
		return 0
	}
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return (uint16)(C.SDL_ReadLE16(_src))
}

func (src *RWops) ReadBE16() uint16 {
	if src == nil {
		return 0
	}
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return (uint16)(C.SDL_ReadBE16(_src))
}

func (src *RWops) ReadLE32() uint32 {
	if src == nil {
		return 0
	}
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return (uint32)(C.SDL_ReadLE32(_src))
}

func (src *RWops) ReadBE32() uint32 {
	if src == nil {
		return 0
	}
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return (uint32)(C.SDL_ReadBE32(_src))
}

func (src *RWops) ReadLE64() uint64 {
	if src == nil {
		return 0
	}
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return (uint64)(C.SDL_ReadLE64(_src))
}

func (src *RWops) ReadBE64() uint64 {
	if src == nil {
		return 0
	}
	_src := (*C.SDL_RWops)(unsafe.Pointer(src))
	return (uint64)(C.SDL_ReadBE64(_src))
}

func (dst *RWops) WriteU8(value uint8) uint {
	if dst == nil {
		return 0
	}
	_dst := (*C.SDL_RWops)(unsafe.Pointer(dst))
	_value := (C.Uint8)(value)
	return (uint)(C.SDL_WriteU8(_dst, _value))
}

func (dst *RWops) WriteLE16(value uint16) uint {
	if dst == nil {
		return 0
	}
	_dst := (*C.SDL_RWops)(unsafe.Pointer(dst))
	_value := (C.Uint16)(value)
	return (uint)(C.SDL_WriteLE16(_dst, _value))
}

func (dst *RWops) WriteBE16(value uint16) uint {
	if dst == nil {
		return 0
	}
	_dst := (*C.SDL_RWops)(unsafe.Pointer(dst))
	_value := (C.Uint16)(value)
	return (uint)(C.SDL_WriteBE16(_dst, _value))
}

func (dst *RWops) WriteLE32(value uint32) uint {
	if dst == nil {
		return 0
	}
	_dst := (*C.SDL_RWops)(unsafe.Pointer(dst))
	_value := (C.Uint32)(value)
	return (uint)(C.SDL_WriteLE32(_dst, _value))
}

func (dst *RWops) WriteBE32(value uint32) uint {
	if dst == nil {
		return 0
	}
	_dst := (*C.SDL_RWops)(unsafe.Pointer(dst))
	_value := (C.Uint32)(value)
	return (uint)(C.SDL_WriteBE32(_dst, _value))
}

func (dst *RWops) WriteLE64(value uint64) uint {
	if dst == nil {
		return 0
	}
	_dst := (*C.SDL_RWops)(unsafe.Pointer(dst))
	_value := (C.Uint64)(value)
	return (uint)(C.SDL_WriteLE64(_dst, _value))
}

func (dst *RWops) WriteBE64(value uint64) uint {
	if dst == nil {
		return 0
	}
	_dst := (*C.SDL_RWops)(unsafe.Pointer(dst))
	_value := (C.Uint64)(value)
	return (uint)(C.SDL_WriteBE64(_dst, _value))
}
