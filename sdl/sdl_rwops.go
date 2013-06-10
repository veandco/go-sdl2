package sdl

/*
#include <SDL2/SDL.h>
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
	RW_SEEK_SET			= 0
	RW_SEEK_CUR			= 1
	RW_SEEK_END			= 2
)

type RWops C.SDL_RWops

func RWFromFile(file, mode string) *RWops {
	_file := (C.CString) (file)
	_mode := (C.CString) (mode)
	return (*RWops) (unsafe.Pointer(C.SDL_RWFromFile(_file, _mode)))
}

func RWFromMem(mem unsafe.Pointer, size int) *RWops {
	_size := (C.int) (size)
	return (*RWops) (unsafe.Pointer(C.SDL_RWFromMem(mem, _size)))
}

func AllocRW() *RWops {
	return (*RWops) (unsafe.Pointer(C.SDL_AllocRW()))
}

func FreeRw(area *RWops) {
	_area := (*C.SDL_RWops) (area)
	C.SDL_FreeRW(_area)
}

func RWsize(ctx *RWops) int64 {
	_ctx := (*C.SDL_RWops) (unsafe.Pointer(ctx))
	return (int64) (C.RWsize(_ctx))
}

func RWseek(ctx *RWops, offset int64, whence int) int64 {
	_ctx := (*C.SDL_RWops) (unsafe.Pointer(ctx))
	_offset := (C.Sint64) (offset)
	_whence := (C.int) (whence)
	return (int64) (C.RWseek(_ctx, _offset, _whence))
}

func RWread(ctx *RWops, ptr unsafe.Pointer, size, maxnum uint) uint {
	_ctx := (*C.SDL_RWops) (unsafe.Pointer(ctx))
	_size := (C.size_t) (size)
	_maxnum := (C.size_t) (maxnum)
	return (uint) (C.RWread(_ctx, ptr, _size, _maxnum))
}

func RWtell(ctx *RWops) int64 {
	_ctx := (*C.SDL_RWops) (unsafe.Pointer(ctx))
	return (int64) (C.RWseek(_ctx, 0, RW_SEEK_CUR))
}

func RWwrite(ctx *RWops, ptr unsafe.Pointer, size, num uint) uint {
	_ctx := (*C.SDL_RWops) (unsafe.Pointer(ctx))
	_size := (C.size_t) (size)
	_num := (C.size_t) (num)
	return (uint) (C.RWwrite(_ctx, ptr, _size, _num))
}

func RWclose(ctx *RWops) int {
	_ctx := (*C.SDL_RWops) (unsafe.Pointer(ctx))
	return (int) (C.RWclose(_ctx))
}

func ReadU8(src *RWops) uint8 {
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	return (uint8) (C.SDL_ReadU8(_src))
}

func ReadLE16(src *RWops) uint16 {
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	return (uint16) (C.SDL_ReadLE16(_src))
}

func ReadBE16(src *RWops) uint16 {
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	return (uint16) (C.SDL_ReadBE16(_src))
}

func ReadLE32(src *RWops) uint32 {
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	return (uint32) (C.SDL_ReadLE32(_src))
}

func ReadBE32(src *RWops) uint32 {
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	return (uint32) (C.SDL_ReadBE32(_src))
}

func ReadLE64(src *RWops) uint64 {
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	return (uint64) (C.SDL_ReadLE64(_src))
}

func ReadBE64(src *RWops) uint64 {
	_src := (*C.SDL_RWops) (unsafe.Pointer(src))
	return (uint64) (C.SDL_ReadBE64(_src))
}

func WriteU8(dst *RWops, value uint8) uint {
	_dst := (*C.SDL_RWops) (unsafe.Pointer(dst))
	_value := (C.Uint8) (value)
	return (uint) (C.SDL_WriteU8(_dst, _value))
}

func WriteLE16(dst *RWops, value uint16) uint {
	_dst := (*C.SDL_RWops) (unsafe.Pointer(dst))
	_value := (C.Uint16) (value)
	return (uint) (C.SDL_WriteLE16(_dst, _value))
}

func WriteBE16(dst *RWops, value uint16) uint {
	_dst := (*C.SDL_RWops) (unsafe.Pointer(dst))
	_value := (C.Uint16) (value)
	return (uint) (C.SDL_WriteBE16(_dst, _value))
}

func WriteLE32(dst *RWops, value uint32) uint {
	_dst := (*C.SDL_RWops) (unsafe.Pointer(dst))
	_value := (C.Uint32) (value)
	return (uint) (C.SDL_WriteLE32(_dst, _value))
}

func WriteBE32(dst *RWops, value uint32) uint {
	_dst := (*C.SDL_RWops) (unsafe.Pointer(dst))
	_value := (C.Uint32) (value)
	return (uint) (C.SDL_WriteBE32(_dst, _value))
}

func WriteLE64(dst *RWops, value uint64) uint {
	_dst := (*C.SDL_RWops) (unsafe.Pointer(dst))
	_value := (C.Uint64) (value)
	return (uint) (C.SDL_WriteLE64(_dst, _value))
}

func WriteBE64(dst *RWops, value uint64) uint {
	_dst := (*C.SDL_RWops) (unsafe.Pointer(dst))
	_value := (C.Uint64) (value)
	return (uint) (C.SDL_WriteBE64(_dst, _value))
}
