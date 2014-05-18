package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"

const (
	CONTROLLER_BINDTYPE_NONE = iota
	CONTROLLER_BINDTYPE_BUTTON
	CONTROLLER_BINDTYPE_AXIS
	CONTROLLER_BINDTYPE_HAT
)

const (
	CONTROLLER_AXIS_INVALID = iota - 1
	CONTROLLER_AXIS_LEFTX
	CONTROLLER_AXIS_LEFTY
	CONTROLLER_AXIS_RIGHTX
	CONTROLLER_AXIS_RIGHTY
	CONTROLLER_AXIS_TRIGGERLEFT
	CONTROLLER_AXIS_TRIGGERRIGHT
	CONTROLLER_AXIS_MAX
)

const (
	CONTROLLER_BUTTON_INVALID = iota - 1
	CONTROLLER_BUTTON_A
	CONTROLLER_BUTTON_B
	CONTROLLER_BUTTON_X
	CONTROLLER_BUTTON_Y
	CONTROLLER_BUTTON_BACK
	CONTROLLER_BUTTON_GUIDE
	CONTROLLER_BUTTON_START
	CONTROLLER_BUTTON_LEFTSTICK
	CONTROLLER_BUTTON_RIGHTSTICK
	CONTROLLER_BUTTON_LEFTSHOULDER
	CONTROLLER_BUTTON_RIGHTSHOULDER
	CONTROLLER_BUTTON_DPAD_UP
	CONTROLLER_BUTTON_DPAD_DOWN
	CONTROLLER_BUTTON_DPAD_LEFT
	CONTROLLER_BUTTON_DPAD_RIGHT
	CONTROLLER_BUTTON_MAX
)

type GameControllerBindType C.SDL_GameControllerBindType
type GameControllerAxis C.SDL_GameControllerAxis
type GameControllerButton C.SDL_GameControllerButton
type GameController C.SDL_GameController
type GameControllerButtonBind C.SDL_GameControllerButtonBind

func (ctrl *GameController) cptr() *C.SDL_GameController {
    return (*C.SDL_GameController)(unsafe.Pointer(ctrl))
}

func (axis GameControllerAxis) c() C.SDL_GameControllerAxis {
    return C.SDL_GameControllerAxis(axis)
}

func (btn GameControllerButton) c() C.SDL_GameControllerButton {
    return C.SDL_GameControllerButton(btn)
}

func GameControllerAddMapping(mappingString string) int {
	_mappingString := C.CString(mappingString)
	defer C.free(unsafe.Pointer(_mappingString))
	return int(C.SDL_GameControllerAddMapping(_mappingString))
}

func GameControllerMappingForGUID(guid JoystickGUID) string {
	return C.GoString(C.SDL_GameControllerMappingForGUID(guid.c()))
}

func GameControllerMapping(ctrl *GameController) string {
	return C.GoString(C.SDL_GameControllerMapping(ctrl.cptr()))
}

func IsGameController(index int) bool {
	return C.SDL_IsGameController(C.int(index)) > 0
}

func GameControllerNameForIndex(index int) string {
	return C.GoString(C.SDL_GameControllerNameForIndex(C.int(index)))
}

func GameControllerOpen(index int) *GameController {
	return (*GameController)(unsafe.Pointer(C.SDL_GameControllerOpen(C.int(index))))
}

func (ctrl *GameController) GetAttached() bool {
	return C.SDL_GameControllerGetAttached(ctrl.cptr()) > 0
}

func (ctrl *GameController) GetJoystick() *Joystick {
	return (*Joystick)(unsafe.Pointer(C.SDL_GameControllerGetJoystick(ctrl.cptr())))
}

func GameControllerEventState(state int) int {
	return int(C.SDL_GameControllerEventState(C.int(state)))
}

func GameControllerUpdate() {
	C.SDL_GameControllerUpdate()
}

func GameControllerGetAxisFromString(pchString string) GameControllerAxis {
	_pchString := C.CString(pchString)
	defer C.free(unsafe.Pointer(_pchString))
	return GameControllerAxis(C.SDL_GameControllerGetAxisFromString(_pchString))
}

func GameControllerGetStringForAxis(axis GameControllerAxis) string {
	return C.GoString(C.SDL_GameControllerGetStringForAxis(axis.c()))
}

func (ctrl *GameController) GetBindForAxis(axis GameControllerAxis) GameControllerButtonBind {
	return GameControllerButtonBind(C.SDL_GameControllerGetBindForAxis(ctrl.cptr(), axis.c()))
}

func (ctrl *GameController) GetAxis(axis GameControllerAxis) int16 {
	return int16(C.SDL_GameControllerGetAxis(ctrl.cptr(), axis.c()))
}

func GameControllerGetButtonFromString(pchString string) GameControllerButton {
	_pchString := C.CString(pchString)
	defer C.free(unsafe.Pointer(_pchString))
	return GameControllerButton(C.SDL_GameControllerGetButtonFromString(_pchString))
}

func GameControllerGetStringForButton(btn GameControllerButton) string {
	return C.GoString(C.SDL_GameControllerGetStringForButton(btn.c()))
}

func (ctrl *GameController) GetBindForButton(btn GameControllerButton) GameControllerButtonBind {
	return GameControllerButtonBind(C.SDL_GameControllerGetBindForButton(ctrl.cptr(), btn.c()))
}

func (ctrl *GameController) GetButton(btn GameControllerButton) byte {
	return byte(C.SDL_GameControllerGetButton(ctrl.cptr(), btn.c()))
}

func (ctrl *GameController) Close() {
	C.SDL_GameControllerClose(ctrl.cptr())
}
