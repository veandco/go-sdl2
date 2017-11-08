package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"

// Hat positions.
// (https://wiki.libsdl.org/SDL_JoystickGetHat)
const (
	HAT_CENTERED  = C.SDL_HAT_CENTERED
	HAT_UP        = C.SDL_HAT_UP
	HAT_RIGHT     = C.SDL_HAT_RIGHT
	HAT_DOWN      = C.SDL_HAT_DOWN
	HAT_LEFT      = C.SDL_HAT_LEFT
	HAT_RIGHTUP   = C.SDL_HAT_RIGHTUP
	HAT_RIGHTDOWN = C.SDL_HAT_RIGHTDOWN
	HAT_LEFTUP    = C.SDL_HAT_LEFTUP
	HAT_LEFTDOWN  = C.SDL_HAT_LEFTDOWN
)

// Joystick is an SDL joystick.
type Joystick C.SDL_Joystick

// JoystickGUID is a stable unique id for a joystick device.
type JoystickGUID C.SDL_JoystickGUID

// JoystickID is joystick's instance id.
type JoystickID C.SDL_JoystickID

func (joy *Joystick) cptr() *C.SDL_Joystick {
	return (*C.SDL_Joystick)(unsafe.Pointer(joy))
}

func (guid JoystickGUID) c() C.SDL_JoystickGUID {
	return C.SDL_JoystickGUID(guid)
}

// NumJoysticks returns the number of joysticks attached to the system.
// (https://wiki.libsdl.org/SDL_NumJoysticks)
func NumJoysticks() int {
	return (int)(C.SDL_NumJoysticks())
}

// JoystickNameForIndex returns the implementation dependent name of a joystick.
// (https://wiki.libsdl.org/SDL_JoystickNameForIndex)
func JoystickNameForIndex(index int) string {
	return (C.GoString)(C.SDL_JoystickNameForIndex(C.int(index)))
}

// JoystickOpen opens a joystick for use.
// (https://wiki.libsdl.org/SDL_JoystickOpen)
func JoystickOpen(index JoystickID) *Joystick {
	return (*Joystick)(C.SDL_JoystickOpen(C.int(index)))
}

// Name returns the implementation dependent name of a joystick.
// (https://wiki.libsdl.org/SDL_JoystickName)
func (joy *Joystick) Name() string {
	return (C.GoString)(C.SDL_JoystickName(joy.cptr()))
}

// JoystickGetDeviceGUID returns the implementation dependent GUID for the joystick at a given device index.
// (https://wiki.libsdl.org/SDL_JoystickGetDeviceGUID)
func JoystickGetDeviceGUID(index int) JoystickGUID {
	return (JoystickGUID)(C.SDL_JoystickGetDeviceGUID(C.int(index)))
}

// GetGUID returns the implementation-dependent GUID for the joystick.
// (https://wiki.libsdl.org/SDL_JoystickGetGUID)
func (joy *Joystick) GetGUID() JoystickGUID {
	return (JoystickGUID)(C.SDL_JoystickGetGUID(joy.cptr()))
}

// JoystickGetGUIDString returns an ASCII string representation for a given JoystickGUID.
// (https://wiki.libsdl.org/SDL_JoystickGetGUIDString)
func JoystickGetGUIDString(guid JoystickGUID) string {
	_pszGUID := make([]rune, 1024)
	pszGUID := C.CString(string(_pszGUID[:]))
	defer C.free(unsafe.Pointer(pszGUID))
	C.SDL_JoystickGetGUIDString(guid.c(), pszGUID, C.int(unsafe.Sizeof(_pszGUID)))
	return C.GoString(pszGUID)
}

// JoystickGetGUIDFromString converts a GUID string into a JoystickGUID structure.
// (https://wiki.libsdl.org/SDL_JoystickGetGUIDFromString)
func JoystickGetGUIDFromString(pchGUID string) JoystickGUID {
	_pchGUID := C.CString(pchGUID)
	defer C.free(unsafe.Pointer(_pchGUID))
	return (JoystickGUID)(C.SDL_JoystickGetGUIDFromString(_pchGUID))
}

// GetAttached returns the status of a specified joystick.
// (https://wiki.libsdl.org/SDL_JoystickGetAttached)
func (joy *Joystick) GetAttached() bool {
	return C.SDL_JoystickGetAttached(joy.cptr()) > 0
}

// InstanceID returns the instance ID of an opened joystick.
// (https://wiki.libsdl.org/SDL_JoystickInstanceID)
func (joy *Joystick) InstanceID() JoystickID {
	return (JoystickID)(C.SDL_JoystickInstanceID(joy.cptr()))
}

// NumAxes returns the number of general axis controls on a joystick.
// (https://wiki.libsdl.org/SDL_JoystickNumAxes)
func (joy *Joystick) NumAxes() int {
	return (int)(C.SDL_JoystickNumAxes(joy.cptr()))
}

// NumBalls returns the number of trackballs on a joystick.
// (https://wiki.libsdl.org/SDL_JoystickNumBalls)
func (joy *Joystick) NumBalls() int {
	return (int)(C.SDL_JoystickNumBalls(joy.cptr()))
}

// NumHats returns the number of POV hats on a joystick.
// (https://wiki.libsdl.org/SDL_JoystickNumHats)
func (joy *Joystick) NumHats() int {
	return (int)(C.SDL_JoystickNumHats(joy.cptr()))
}

// NumButtons returns the number of buttons on a joystick.
// (https://wiki.libsdl.org/SDL_JoystickNumButtons)
func (joy *Joystick) NumButtons() int {
	return (int)(C.SDL_JoystickNumButtons(joy.cptr()))
}

// JoystickUpdate updates the current state of the open joysticks.
// (https://wiki.libsdl.org/SDL_JoystickUpdate)
func JoystickUpdate() {
	C.SDL_JoystickUpdate()
}

// JoystickEventState enables or disables joystick event polling.
// (https://wiki.libsdl.org/SDL_JoystickEventState)
func JoystickEventState(state int) int {
	return (int)(C.SDL_JoystickEventState(C.int(state)))
}

// GetAxis returns the current state of an axis control on a joystick.
// (https://wiki.libsdl.org/SDL_JoystickGetAxis)
func (joy *Joystick) GetAxis(axis int) int16 {
	return (int16)(C.SDL_JoystickGetAxis(joy.cptr(), C.int(axis)))
}

// GetHat returns the current state of a POV hat on a joystick.
// (https://wiki.libsdl.org/SDL_JoystickGetHat)
func (joy *Joystick) GetHat(hat int) byte {
	return (byte)(C.SDL_JoystickGetHat(joy.cptr(), C.int(hat)))
}

// GetBall returns the ball axis change since the last poll.
// (https://wiki.libsdl.org/SDL_JoystickGetBall)
func (joy *Joystick) GetBall(ball int, dx, dy *int) int {
	_dx := (*C.int)(unsafe.Pointer(dx))
	_dy := (*C.int)(unsafe.Pointer(dy))
	return (int)(C.SDL_JoystickGetBall(joy.cptr(), C.int(ball), _dx, _dy))
}

// GetButton the current state of a button on a joystick.
// (https://wiki.libsdl.org/SDL_JoystickGetButton)
func (joy *Joystick) GetButton(button int) byte {
	return (byte)(C.SDL_JoystickGetButton(joy.cptr(), C.int(button)))
}

// Close closes a joystick previously opened with JoystickOpen().
// (https://wiki.libsdl.org/SDL_JoystickClose)
func (joy *Joystick) Close() {
	C.SDL_JoystickClose(joy.cptr())
}
