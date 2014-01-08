package sdl

// #include <SDL2/SDL_joystick.h>
import "C"
import "unsafe"

const (
	HAT_CENTERED	= 0x00
	HAT_UP		= 0x01
	HAT_RIGHT	= 0x02
	HAT_DOWN	= 0x04
	HAT_LEFT	= 0x08
	HAT_RIGHTUP     = HAT_RIGHT | HAT_UP
	HAT_RIGHTDOWN   = HAT_RIGHT | HAT_DOWN
	HAT_LEFTUP      = HAT_LEFT | HAT_UP
	HAT_LEFTDOWN    = HAT_LEFT | HAT_DOWN
)

type Joystick C.SDL_Joystick
type JoystickGUID C.SDL_JoystickGUID
type JoystickID C.SDL_JoystickID

func NumJoysticks() int {
	return (int) (C.SDL_NumJoysticks())
}

func JoystickNameForIndex(device_index int) string {
	_device_index := (C.int) (device_index)
	return (C.GoString) (C.SDL_JoystickNameForIndex(_device_index))
}

func JoystickOpen(device_index int) *Joystick {
	_device_index := (C.int) (device_index)
	return (*Joystick) (C.SDL_JoystickOpen(_device_index))
}

func (joystick *Joystick) Name() string {
	_joystick := (*C.SDL_Joystick) (joystick)
	return (C.GoString) (C.SDL_JoystickName(_joystick))
}

func JoystickGetDeviceGUID(device_index int) JoystickGUID {
	_device_index := (C.int) (device_index)
	return (JoystickGUID) (C.SDL_JoystickGetDeviceGUID(_device_index))
}

func (joystick *Joystick) GetGUID() JoystickGUID {
	_joystick := (*C.SDL_Joystick) (joystick)
	return (JoystickGUID) (C.SDL_JoystickGetGUID(_joystick))
}

func JoystickGetGUIDString(guid JoystickGUID, pszGUID string, cbGUID int) {
	_guid := (C.SDL_JoystickGUID) (guid)
	_pszGUID := (C.CString) (pszGUID)
	_cbGUID := (C.int) (cbGUID)
	defer C.SDL_free(unsafe.Pointer(_pszGUID))
	C.SDL_JoystickGetGUIDString(_guid, _pszGUID, _cbGUID)
}

func JoystickGetGUIDFromString(pchGUID string) JoystickGUID {
	_pchGUID := (C.CString) (pchGUID)
	defer C.SDL_free(unsafe.Pointer(_pchGUID))
	return (JoystickGUID) (C.SDL_JoystickGetGUIDFromString(_pchGUID))
}

func (joystick *Joystick) GetAttached() bool {
	_joystick := (*C.SDL_Joystick) (joystick)
	return C.SDL_JoystickGetAttached(_joystick) > 0
}

func (joystick *Joystick) InstanceID() JoystickID {
	_joystick := (*C.SDL_Joystick) (joystick)
	return (JoystickID) (C.SDL_JoystickInstanceID(_joystick))
}

func (joystick *Joystick) NumAxes() int {
	_joystick := (*C.SDL_Joystick) (joystick)
	return (int) (C.SDL_JoystickNumAxes(_joystick))
}

func (joystick *Joystick) NumBalls() int {
	_joystick := (*C.SDL_Joystick) (joystick)
	return (int) (C.SDL_JoystickNumBalls(_joystick))
}

func (joystick *Joystick) NumHats() int {
	_joystick := (*C.SDL_Joystick) (joystick)
	return (int) (C.SDL_JoystickNumHats(_joystick))
}

func (joystick *Joystick) NumButtons() int {
	_joystick := (*C.SDL_Joystick) (joystick)
	return (int) (C.SDL_JoystickNumButtons(_joystick))
}

func Update() {
	C.SDL_JoystickUpdate()
}

func JoystickEventState(state int) int {
	_state := (C.int) (state)
	return (int) (C.SDL_JoystickEventState(_state))
}

func (joystick *Joystick) GetAxis(axis int) int16 {
	_joystick := (*C.SDL_Joystick) (joystick)
	_axis := (C.int) (axis)
	return (int16) (C.SDL_JoystickGetAxis(_joystick, _axis))
}

func (joystick *Joystick) GetHat(hat int) byte {
	_joystick := (*C.SDL_Joystick) (joystick)
	_hat := (C.int) (hat)
	return (byte) (C.SDL_JoystickGetHat(_joystick, _hat))
}

func (joystick *Joystick) GetBall(ball int, dx, dy *int) int {
	_joystick := (*C.SDL_Joystick) (joystick)
	_ball := (C.int) (ball)
	_dx := (*C.int) (unsafe.Pointer(dx))
	_dy := (*C.int) (unsafe.Pointer(dy))
	return (int) (C.SDL_JoystickGetBall(_joystick, _ball, _dx, _dy))
}

func (joystick *Joystick) GetButton(button int) byte {
	_joystick := (*C.SDL_Joystick) (joystick)
	_button := (C.int) (button)
	return (byte) (C.SDL_JoystickGetButton(_joystick, _button))
}

func (joystick *Joystick) Close() {
	_joystick := (*C.SDL_Joystick) (joystick)
	C.SDL_JoystickClose(_joystick)
}
