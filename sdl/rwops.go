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

#if !(SDL_VERSION_ATLEAST(2,0,6))

#if defined(WARN_OUTDATED)
#pragma message("SDL_LoadFile_RW is not supported before SDL 2.0.6")
#endif

static void * SDL_LoadFile_RW(SDL_RWops * src, size_t *datasize, int freesrc)
{
	return 0;
}
#endif
*/
import "C"
import (
	"reflect"
	"unsafe"
)

// RWops types
const (
	RWOPS_UNKNOWN   = 0 // unknown stream type
	RWOPS_WINFILE   = 1 // win32 file
	RWOPS_STDFILE   = 2 // stdio file
	RWOPS_JNIFILE   = 3 // android asset
	RWOPS_MEMORY    = 4 // memory stream
	RWOPS_MEMORY_RO = 5 // read-only memory stream
)

// RWops seek from
const (
	RW_SEEK_SET = C.RW_SEEK_SET // seek from the beginning of data
	RW_SEEK_CUR = C.RW_SEEK_CUR // seek relative to current read point
	RW_SEEK_END = C.RW_SEEK_END // seek relative to the end of data
)

// RWops provides an abstract interface to stream I/O. Applications can generally ignore the specifics of this structure's internals and treat them as opaque pointers. The details are important to lower-level code that might need to implement one of these, however.
// (https://wiki.libsdl.org/SDL_RWops)
type RWops C.SDL_RWops

func (rwops *RWops) cptr() *C.SDL_RWops {
	return (*C.SDL_RWops)(rwops)
}

// RWFromFile creates a new RWops structure for reading from and/or writing to a named file.
// (https://wiki.libsdl.org/SDL_RWFromFile)
func RWFromFile(file, mode string) *RWops {
	_file := C.CString(file)
	_mode := C.CString(mode)
	defer C.free(unsafe.Pointer(_file))
	defer C.free(unsafe.Pointer(_mode))
	return (*RWops)(unsafe.Pointer(C.SDL_RWFromFile(_file, _mode)))
}

// RWFromMem prepares a read-write memory buffer for use with RWops.
// (https://wiki.libsdl.org/SDL_RWFromMem)
func RWFromMem(mem []byte) (*RWops, error) {
	if mem == nil {
		return nil, ErrInvalidParameters
	}

	header := (*reflect.SliceHeader)(unsafe.Pointer(&mem))
	_mem := unsafe.Pointer(header.Data)

	rwops := (*RWops)(unsafe.Pointer(C.SDL_RWFromMem(_mem, C.int(len(mem)))))
	if rwops == nil {
		return nil, GetError()
	}
	return rwops, nil
}

// AllocRW allocates an empty, unpopulated RWops structure.
// (https://wiki.libsdl.org/SDL_AllocRW)
func AllocRW() *RWops {
	return (*RWops)(unsafe.Pointer(C.SDL_AllocRW()))
}

// Free frees the RWops structure allocated by AllocRW().
// (https://wiki.libsdl.org/SDL_FreeRW)
func (rwops *RWops) Free() error {
	if rwops == nil {
		return ErrInvalidParameters
	}

	C.SDL_FreeRW(rwops.cptr())
	return nil
}

// Size returns the size of the data stream in the RWops.
// (https://wiki.libsdl.org/SDL_RWsize)
func (rwops *RWops) Size() (int64, error) {
	n := int64(C.RWsize(rwops.cptr()))
	if n < 0 {
		return n, GetError()
	}
	return n, nil
}

// Seek seeks within the RWops data stream.
// (https://wiki.libsdl.org/SDL_RWseek)
func (rwops *RWops) Seek(offset int64, whence int) (int64, error) {
	if rwops == nil {
		return -1, ErrInvalidParameters
	}

	ret := int64(C.RWseek(rwops.cptr(), C.Sint64(offset), C.int(whence)))
	if ret < 0 {
		return ret, GetError()
	}
	return ret, nil
}

// Read reads from a data source.
// (https://wiki.libsdl.org/SDL_RWread)
func (rwops *RWops) Read(buf []byte) (n int, err error) {
	return rwops.Read2(buf, 1, uint(len(buf)))
}

// Read2 reads from a data source (native).
// (https://wiki.libsdl.org/SDL_RWread)
func (rwops *RWops) Read2(buf []byte, size, maxnum uint) (n int, err error) {
	if rwops == nil || buf == nil {
		return 0, ErrInvalidParameters
	}

	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	_data := unsafe.Pointer(header.Data)

	n = int(C.RWread(rwops.cptr(), _data, C.size_t(size), C.size_t(maxnum)))
	if n == 0 {
		err = GetError()
	}
	return
}

// Tell returns the current read/write offset in the RWops data stream.
// (https://wiki.libsdl.org/SDL_RWtell)
func (rwops *RWops) Tell() (int64, error) {
	if rwops == nil {
		return 0, ErrInvalidParameters
	}

	ret := int64(C.RWseek(rwops.cptr(), 0, C.int(RW_SEEK_CUR)))
	if ret < 0 {
		return ret, GetError()
	}
	return ret, nil
}

// Write writes to the RWops data stream.
// (https://wiki.libsdl.org/SDL_RWwrite)
func (rwops *RWops) Write(buf []byte) (n int, err error) {
	return rwops.Write2(buf, 1, uint(len(buf)))
}

// Write2 writes to the RWops data stream (native).
// (https://wiki.libsdl.org/SDL_RWwrite)
func (rwops *RWops) Write2(buf []byte, size, num uint) (n int, err error) {
	if rwops == nil || buf == nil {
		return 0, ErrInvalidParameters
	}

	header := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	_data := unsafe.Pointer(header.Data)

	n = int(C.RWwrite(rwops.cptr(), _data, C.size_t(size), C.size_t(num)))
	if n < int(num) {
		err = GetError()
	}
	return
}

// Close closes and frees the allocated RWops structure.
// (https://wiki.libsdl.org/SDL_RWclose)
func (rwops *RWops) Close() error {
	if rwops != nil && C.RWclose(rwops.cptr()) != 0 {
		return GetError()
	}
	return nil
}

// ReadU8 reads a byte from the RWops.
// (https://wiki.libsdl.org/SDL_ReadU8)
func (rwops *RWops) ReadU8() uint8 {
	if rwops == nil {
		return 0
	}
	return uint8(C.SDL_ReadU8(rwops.cptr()))
}

// ReadLE16 reads 16 bits of little-endian data from the RWops and returns in native format.
// (https://wiki.libsdl.org/SDL_ReadLE16)
func (rwops *RWops) ReadLE16() uint16 {
	if rwops == nil {
		return 0
	}
	return uint16(C.SDL_ReadLE16(rwops.cptr()))
}

// ReadBE16 read 16 bits of big-endian data from the RWops and returns in native format.
// (https://wiki.libsdl.org/SDL_ReadBE16)
func (rwops *RWops) ReadBE16() uint16 {
	if rwops == nil {
		return 0
	}
	return uint16(C.SDL_ReadBE16(rwops.cptr()))
}

// ReadLE32 reads 32 bits of little-endian data from the RWops and returns in native format.
// (https://wiki.libsdl.org/SDL_ReadLE32)
func (rwops *RWops) ReadLE32() uint32 {
	if rwops == nil {
		return 0
	}
	return uint32(C.SDL_ReadLE32(rwops.cptr()))
}

// ReadBE32 reads 32 bits of big-endian data from the RWops and returns in native format.
// (https://wiki.libsdl.org/SDL_ReadBE32)
func (rwops *RWops) ReadBE32() uint32 {
	if rwops == nil {
		return 0
	}
	return uint32(C.SDL_ReadBE32(rwops.cptr()))
}

// ReadLE64 reads 64 bits of little-endian data from the RWops and returns in native format.
// (https://wiki.libsdl.org/SDL_ReadLE64)
func (rwops *RWops) ReadLE64() uint64 {
	if rwops == nil {
		return 0
	}
	return uint64(C.SDL_ReadLE64(rwops.cptr()))
}

// ReadBE64 reads 64 bits of big-endian data from the RWops and returns in native format.
// (https://wiki.libsdl.org/SDL_ReadBE64)
func (rwops *RWops) ReadBE64() uint64 {
	if rwops == nil {
		return 0
	}
	return uint64(C.SDL_ReadBE64(rwops.cptr()))
}

// LoadFile_RW loads all the data from an SDL data stream.
// (https://wiki.libsdl.org/SDL_LoadFile_RW)
func (src *RWops) LoadFileRW(freesrc bool) (data []byte, size int) {
	var _size C.size_t
	var _freesrc C.int = 0

	if freesrc {
		_freesrc = 1
	}

	_data := C.SDL_LoadFile_RW(src.cptr(), &_size, _freesrc)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	sliceHeader.Cap = int(_size)
	sliceHeader.Len = int(_size)
	sliceHeader.Data = uintptr(_data)
	size = int(_size)
	return
}

// LoadFile loads an entire file
// (https://wiki.libsdl.org/SDL_LoadFile)
func LoadFile(file string) (data []byte, size int) {
	return RWFromFile(file, "rb").LoadFileRW(true)
}

// WriteU8 writes a byte to the RWops.
// (https://wiki.libsdl.org/SDL_WriteU8)
func (rwops *RWops) WriteU8(value uint8) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteU8(rwops.cptr(), C.Uint8(value)))
}

// WriteLE16 writes 16 bits in native format to the RWops as little-endian data.
// (https://wiki.libsdl.org/SDL_WriteLE16)
func (rwops *RWops) WriteLE16(value uint16) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteLE16(rwops.cptr(), C.Uint16(value)))
}

// WriteBE16 writes 16 bits in native format to the RWops as big-endian data.
// (https://wiki.libsdl.org/SDL_WriteBE16)
func (rwops *RWops) WriteBE16(value uint16) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteBE16(rwops.cptr(), C.Uint16(value)))
}

// WriteLE32 writes 32 bits in native format to the RWops as little-endian data.
// (https://wiki.libsdl.org/SDL_WriteLE32)
func (rwops *RWops) WriteLE32(value uint32) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteLE32(rwops.cptr(), C.Uint32(value)))
}

// WriteBE32 writes 32 bits in native format to the RWops as big-endian data.
// (https://wiki.libsdl.org/SDL_WriteBE32)
func (rwops *RWops) WriteBE32(value uint32) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteBE32(rwops.cptr(), C.Uint32(value)))
}

// WriteLE64 writes 64 bits in native format to the RWops as little-endian data.
// (https://wiki.libsdl.org/SDL_WriteLE64)
func (rwops *RWops) WriteLE64(value uint64) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteLE64(rwops.cptr(), C.Uint64(value)))
}

// WriteBE64 writes 64 bits in native format to the RWops as big-endian data.
// (https://wiki.libsdl.org/SDL_WriteBE64)
func (rwops *RWops) WriteBE64(value uint64) uint {
	if rwops == nil {
		return 0
	}
	return uint(C.SDL_WriteBE64(rwops.cptr(), C.Uint64(value)))
}
