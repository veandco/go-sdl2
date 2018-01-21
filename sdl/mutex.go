package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// ThreadID is the thread identifier for a thread.
type ThreadID uint64

// Mutex is the SDL mutex structure.
type Mutex struct {
	Recursive int
	Owner     ThreadID
	Sem       *Sem
}

// Sem is the SDL semaphore structure.
type Sem struct {
	Count        uint32
	WaitersCount uint32
	CountLock    *Mutex
	CountNonzero *Cond
}

// Cond is the SDL condition variable structure.
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

// CreateMutex creates a new mutex.
// (https://wiki.libsdl.org/SDL_CreateMutex)
func CreateMutex() (*Mutex, error) {
	mutex := C.SDL_CreateMutex()
	if mutex == nil {
		return nil, GetError()
	}
	return (*Mutex)(unsafe.Pointer(mutex)), nil
}

// Lock locks a mutex created with CreateMutex().
// (https://wiki.libsdl.org/SDL_LockMutex)
func (mutex *Mutex) Lock() error {
	ret := int(C.SDL_LockMutex(mutex.cptr()))
	if ret != 0 {
		return GetError()
	}
	return nil
}

// TryLock tries to lock a mutex without blocking.
// (https://wiki.libsdl.org/SDL_TryLockMutex)
func (mutex *Mutex) TryLock() error {
	ret := int(C.SDL_TryLockMutex(mutex.cptr()))
	if ret != 0 {
		return GetError()
	}
	return nil
}

// Unlock unlocks a mutex created with CreateMutex().
// (https://wiki.libsdl.org/SDL_UnlockMutex)
func (mutex *Mutex) Unlock() error {
	ret := int(C.SDL_UnlockMutex(mutex.cptr()))
	if ret != 0 {
		return GetError()
	}
	return nil
}

// Destroy destroys a mutex created with CreateMutex().
// (https://wiki.libsdl.org/SDL_DestroyMutex)
func (mutex *Mutex) Destroy() {
	C.SDL_DestroyMutex(mutex.cptr())
}

// CreateSemaphore creates a semaphore.
// (https://wiki.libsdl.org/SDL_CreateSemaphore)
func CreateSemaphore(initialValue uint32) (*Sem, error) {
	sem := C.SDL_CreateSemaphore(C.Uint32(initialValue))
	if sem == nil {
		return nil, GetError()
	}
	return (*Sem)(unsafe.Pointer(sem)), nil
}

// Destroy destroys a semaphore.
// (https://wiki.libsdl.org/SDL_DestroySemaphore)
func (sem *Sem) Destroy() {
	C.SDL_DestroySemaphore(sem.cptr())
}

// Wait waits until a semaphore has a positive value and then decrements it.
// (https://wiki.libsdl.org/SDL_SemWait)
func (sem *Sem) Wait() error {
	ret := int(C.SDL_SemWait(sem.cptr()))
	if ret != 0 {
		return GetError()
	}
	return nil
}

// TryWait sees if a semaphore has a positive value and decrement it if it does.
// (https://wiki.libsdl.org/SDL_SemTryWait)
func (sem *Sem) TryWait() error {
	ret := int(C.SDL_SemTryWait(sem.cptr()))
	if ret != 0 {
		return GetError()
	}
	return nil
}

// WaitTimeout waits until a semaphore has a positive value and then decrements it.
// (https://wiki.libsdl.org/SDL_SemWaitTimeout)
func (sem *Sem) WaitTimeout(ms uint32) error {
	ret := int(C.SDL_SemWaitTimeout(sem.cptr(), C.Uint32(ms)))
	if ret != 0 {
		return GetError()
	}
	return nil
}

// Post atomically increments a semaphore's value and wake waiting threads.
// (https://wiki.libsdl.org/SDL_SemPost)
func (sem *Sem) Post() error {
	ret := int(C.SDL_SemPost(sem.cptr()))
	if ret != 0 {
		return GetError()
	}
	return nil
}

// Value returns the current value of a semaphore.
// (https://wiki.libsdl.org/SDL_SemValue)
func (sem *Sem) Value() uint32 {
	return uint32(C.SDL_SemValue(sem.cptr()))
}

// CreateCond (https://wiki.libsdl.org/SDL_CreateCond)
func CreateCond() *Cond {
	return (*Cond)(unsafe.Pointer(C.SDL_CreateCond()))
}

// Destroy creates a condition variable.
// (https://wiki.libsdl.org/SDL_DestroyCond)
func (cond *Cond) Destroy() {
	C.SDL_DestroyCond(cond.cptr())
}

// Signal restarts one of the threads that are waiting on the condition variable.
// (https://wiki.libsdl.org/SDL_CondSignal)
func (cond *Cond) Signal() error {
	ret := int(C.SDL_CondSignal(cond.cptr()))
	if ret != 0 {
		return GetError()
	}
	return nil
}

// Broadcast restarts all threads that are waiting on the condition variable.
// (https://wiki.libsdl.org/SDL_CondBroadcast)
func (cond *Cond) Broadcast() error {
	ret := int(C.SDL_CondBroadcast(cond.cptr()))
	if ret != 0 {
		return GetError()
	}
	return nil
}

// Wait waits until a condition variable is signaled.
// (https://wiki.libsdl.org/SDL_CondWait)
func (cond *Cond) Wait(mutex *Mutex) error {
	ret := int(C.SDL_CondWait(cond.cptr(), mutex.cptr()))
	if ret != 0 {
		return GetError()
	}
	return nil
}

// WaitTimeout waits until a condition variable is signaled or a specified amount of time has passed.
// (https://wiki.libsdl.org/SDL_CondWaitTimeout)
func (cond *Cond) WaitTimeout(mutex *Mutex, ms uint32) error {
	ret := int(C.SDL_CondWaitTimeout(cond.cptr(), mutex.cptr(), C.Uint32(ms)))
	if ret != 0 {
		return GetError()
	}
	return nil
}
