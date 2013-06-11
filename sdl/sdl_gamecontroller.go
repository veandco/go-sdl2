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

func GameControllerAddMapping(mappingString string) int {
	_mappingString := (C.CString) (mappingString)
	defer C.free(unsafe.Pointer(_mappingString))
	return (int) (C.SDL_GameControllerAddMapping(_mappingString))
}

func GameControllerMappingForGUID(guid JoystickGUID) string {
	_guid := (C.SDL_JoystickGUID) (guid)
	return (C.GoString) (C.SDL_GameControllerMappingForGUID(_guid))
}

func GameControllerMapping(gamecontroller *GameController) string {
	_gamecontroller := (*C.SDL_GameController) (gamecontroller)
	return (C.GoString) (C.SDL_GameControllerMapping(_gamecontroller))
}

func IsGameController(joystick_index int) bool {
	_joystick_index := (C.int) (joystick_index)
	return C.SDL_IsGameController(_joystick_index) > 0
}

func GameControllerNameForIndex(joystick_index int) string {
	_joystick_index := (C.int) (joystick_index)
	return (C.GoString) (C.SDL_GameControllerNameForIndex(_joystick_index))
}

func GameControllerOpen(joystick_index int) *GameController {
	_joystick_index := (C.int) (joystick_index)
	return (*GameController) (unsafe.Pointer(C.SDL_GameControllerOpen(_joystick_index)))
}

func (gamecontroller *GameController) GetAttached() bool {
	_gamecontroller := (*C.SDL_GameController) (unsafe.Pointer(gamecontroller))
	return C.SDL_GameControllerGetAttached(_gamecontroller) > 0
}

func (gamecontroller *GameController) GetJoystick() *Joystick {
	_gamecontroller := (*C.SDL_GameController) (unsafe.Pointer(gamecontroller))
	return (*Joystick) (unsafe.Pointer(C.SDL_GameControllerGetJoystick(_gamecontroller)))
}

func GameControllerEventState(state int) int {
	_state := (C.int) (state)
	return (int) (C.SDL_GameControllerEventState(_state))
}

func GameControllerUpdate() {
	C.SDL_GameControllerUpdate()
}

func GameControllerGetAxisFromString(pchString string) GameControllerAxis {
	_pchString := (C.CString) (pchString)
	defer C.free(unsafe.Pointer(_pchString))
	return (GameControllerAxis) (C.SDL_GameControllerGetAxisFromString(_pchString))
}

func GameControllerGetStringForAxis(axis GameControllerAxis) string {
	_axis := (C.SDL_GameControllerAxis) (axis)
	return (C.GoString) (C.SDL_GameControllerGetStringForAxis(_axis))
}

func (gamecontroller *GameController) GetBindForAxis(axis GameControllerAxis) GameControllerButtonBind {
	_gamecontroller := (*C.SDL_GameController) (gamecontroller)
	_axis := (C.SDL_GameControllerAxis) (axis)
	return (GameControllerButtonBind) (C.SDL_GameControllerGetBindForAxis(_gamecontroller, _axis))
}

func (gamecontroller *GameController) GetAxis(axis GameControllerAxis) int16 {
	_gamecontroller := (*C.SDL_GameController) (gamecontroller)
	_axis := (C.SDL_GameControllerAxis) (axis)
	return (int16) (C.SDL_GameControllerGetAxis(_gamecontroller, _axis))
}

func GameControllerGetButtonFromString(pchString string) GameControllerButton {
	_pchString := (C.CString) (pchString)
	defer C.free(unsafe.Pointer(_pchString))
	return (GameControllerButton) (C.SDL_GameControllerGetButtonFromString(_pchString))
}

func GameControllerGetStringForButton(button GameControllerButton) string {
	_button := (C.SDL_GameControllerButton) (button)
	return (C.GoString) (C.SDL_GameControllerGetStringForButton(_button))
}

func (gamecontroller *GameController) GetBindForButton(button GameControllerButton) GameControllerButtonBind {
	_gamecontroller := (*C.SDL_GameController) (gamecontroller)
	_button := (C.SDL_GameControllerButton) (button)
	return (GameControllerButtonBind) (C.SDL_GameControllerGetBindForButton(_gamecontroller, _button))
}

func (gamecontroller *GameController) GetButton(button GameControllerButton) byte {
	_gamecontroller := (*C.SDL_GameController) (gamecontroller)
	_button := (C.SDL_GameControllerButton) (button)
	return (byte) (C.SDL_GameControllerGetButton(_gamecontroller, _button))
}

func (gamecontroller *GameController) Close() {
	_gamecontroller := (*C.SDL_GameController) (gamecontroller)
	C.SDL_GameControllerClose(_gamecontroller)
}
