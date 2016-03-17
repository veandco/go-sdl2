package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

type ThreadID uint64

type Mutex struct {
	Recursive int
	Owner     ThreadID
	Sem       *Sem
}

type Sem struct {
	Count        uint32
	WaitersCount uint32
	CountLock    *Mutex
	CountNonzero *Cond
}

type Cond struct {
	Lock     *Mutex
	Waiting  int
	Signals  int
	WaitSem  *Sem
	WaitDone *Sem
}

func (m *Mutex) cptr() *C.SDL_mutex {
	return (*C.SDL_mutex)(unsafe.Pointer(m))
}

func (s *Sem) cptr() *C.SDL_sem {
	return (*C.SDL_sem)(unsafe.Pointer(s))
}

func (c *Cond) cptr() *C.SDL_cond {
	return (*C.SDL_cond)(unsafe.Pointer(c))
}

// CreateMutex (https://wiki.libsdl.org/SDL_CreateMutex)
func CreateMutex() (*Mutex, error) {
	mutex := C.SDL_CreateMutex()
	if mutex == nil {
		return nil, GetError()
	}
	return (*Mutex)(unsafe.Pointer(mutex)), nil
}

// LockMutex (https://wiki.libsdl.org/SDL_LockMutex)
func LockMutex(mutex *Mutex) int {
	return int(C.SDL_LockMutex(mutex.cptr()))
}

// TryLockMutex (https://wiki.libsdl.org/SDL_TryLockMutex)
func TryLockMutex(mutex *Mutex) int {
	return int(C.SDL_TryLockMutex(mutex.cptr()))
}

// UnlockMutex (https://wiki.libsdl.org/SDL_UnlockMutex)
func UnlockMutex(mutex *Mutex) int {
	return int(C.SDL_UnlockMutex(mutex.cptr()))
}

// DestroyMutex (https://wiki.libsdl.org/SDL_DestroyMutex)
func DestroyMutex(mutex *Mutex) {
	C.SDL_DestroyMutex(mutex.cptr())
}

// CreateSemaphore (https://wiki.libsdl.org/SDL_CreateSemaphore)
func CreateSemaphore(initialValue uint32) (*Sem, error) {
	sem := C.SDL_CreateSemaphore(C.Uint32(initialValue))
	if sem == nil {
		return nil, GetError()
	}
	return (*Sem)(unsafe.Pointer(sem)), nil
}

// DestroySemaphore (https://wiki.libsdl.org/SDL_DestroySemaphore)
func DestroySemaphore(sem *Sem) {
	C.SDL_DestroySemaphore(sem.cptr())
}

// SemWait (https://wiki.libsdl.org/SDL_SemWait)
func SemWait(sem *Sem) int {
	return int(C.SDL_SemWait(sem.cptr()))
}

// SemTryWait (https://wiki.libsdl.org/SDL_SemTryWait)
func SemTryWait(sem *Sem) int {
	return int(C.SDL_SemTryWait(sem.cptr()))
}

// SemWaitTimeout (https://wiki.libsdl.org/SDL_SemWaitTimeout)
func SemWaitTimeout(sem *Sem, ms uint32) int {
	return int(C.SDL_SemWaitTimeout(sem.cptr(), C.Uint32(ms)))
}

// SemPost (https://wiki.libsdl.org/SDL_SemPost)
func SemPost(sem *Sem) int {
	return int(C.SDL_SemPost(sem.cptr()))
}

// SemValue (https://wiki.libsdl.org/SDL_SemValue)
func SemValue(sem *Sem) uint32 {
	return uint32(C.SDL_SemValue(sem.cptr()))
}

// CreateCond (https://wiki.libsdl.org/SDL_CreateCond)
func CreateCond() *Cond {
	return (*Cond)(unsafe.Pointer(C.SDL_CreateCond()))
}

// DestroyCond (https://wiki.libsdl.org/SDL_DestroyCond)
func DestroyCond(cond *Cond) {
	C.SDL_DestroyCond(cond.cptr())
}

// CondSignal (https://wiki.libsdl.org/SDL_CondSignal)
func CondSignal(cond *Cond) int {
	return int(C.SDL_CondSignal(cond.cptr()))
}

// CondBroadcast (https://wiki.libsdl.org/SDL_CondBroadcast)
func CondBroadcast(cond *Cond) int {
	return int(C.SDL_CondBroadcast(cond.cptr()))
}

// CondWait (https://wiki.libsdl.org/SDL_CondWait)
func CondWait(cond *Cond, mutex *Mutex) int {
	return int(C.SDL_CondWait(cond.cptr(), mutex.cptr()))
}

// CondWaitTimeout (https://wiki.libsdl.org/SDL_CondWaitTimeout)
func CondWaitTimeout(cond *Cond, mutex *Mutex, ms uint32) int {
	return int(C.SDL_CondWaitTimeout(cond.cptr(), mutex.cptr(), C.Uint32(ms)))
}
