package sdl

// #include "sdl_wrapper.h"
import "C"
import "unsafe"
import "encoding/binary"

const (
	CONTROLLER_BINDTYPE_NONE   = C.SDL_CONTROLLER_BINDTYPE_NONE
	CONTROLLER_BINDTYPE_BUTTON = C.SDL_CONTROLLER_BINDTYPE_BUTTON
	CONTROLLER_BINDTYPE_AXIS   = C.SDL_CONTROLLER_BINDTYPE_AXIS
	CONTROLLER_BINDTYPE_HAT    = C.SDL_CONTROLLER_BINDTYPE_HAT
)

const (
	CONTROLLER_AXIS_INVALID      = C.SDL_CONTROLLER_AXIS_INVALID
	CONTROLLER_AXIS_LEFTX        = C.SDL_CONTROLLER_AXIS_LEFTX
	CONTROLLER_AXIS_LEFTY        = C.SDL_CONTROLLER_AXIS_LEFTY
	CONTROLLER_AXIS_RIGHTX       = C.SDL_CONTROLLER_AXIS_RIGHTX
	CONTROLLER_AXIS_RIGHTY       = C.SDL_CONTROLLER_AXIS_RIGHTY
	CONTROLLER_AXIS_TRIGGERLEFT  = C.SDL_CONTROLLER_AXIS_TRIGGERLEFT
	CONTROLLER_AXIS_TRIGGERRIGHT = C.SDL_CONTROLLER_AXIS_TRIGGERRIGHT
	CONTROLLER_AXIS_MAX          = C.SDL_CONTROLLER_AXIS_MAX
)

const (
	CONTROLLER_BUTTON_INVALID       = C.SDL_CONTROLLER_BUTTON_INVALID
	CONTROLLER_BUTTON_A             = C.SDL_CONTROLLER_BUTTON_A
	CONTROLLER_BUTTON_B             = C.SDL_CONTROLLER_BUTTON_B
	CONTROLLER_BUTTON_X             = C.SDL_CONTROLLER_BUTTON_X
	CONTROLLER_BUTTON_Y             = C.SDL_CONTROLLER_BUTTON_Y
	CONTROLLER_BUTTON_BACK          = C.SDL_CONTROLLER_BUTTON_BACK
	CONTROLLER_BUTTON_GUIDE         = C.SDL_CONTROLLER_BUTTON_GUIDE
	CONTROLLER_BUTTON_START         = C.SDL_CONTROLLER_BUTTON_START
	CONTROLLER_BUTTON_LEFTSTICK     = C.SDL_CONTROLLER_BUTTON_LEFTSTICK
	CONTROLLER_BUTTON_RIGHTSTICK    = C.SDL_CONTROLLER_BUTTON_RIGHTSTICK
	CONTROLLER_BUTTON_LEFTSHOULDER  = C.SDL_CONTROLLER_BUTTON_LEFTSHOULDER
	CONTROLLER_BUTTON_RIGHTSHOULDER = C.SDL_CONTROLLER_BUTTON_RIGHTSHOULDER
	CONTROLLER_BUTTON_DPAD_UP       = C.SDL_CONTROLLER_BUTTON_DPAD_UP
	CONTROLLER_BUTTON_DPAD_DOWN     = C.SDL_CONTROLLER_BUTTON_DPAD_DOWN
	CONTROLLER_BUTTON_DPAD_LEFT     = C.SDL_CONTROLLER_BUTTON_DPAD_LEFT
	CONTROLLER_BUTTON_DPAD_RIGHT    = C.SDL_CONTROLLER_BUTTON_DPAD_RIGHT
	CONTROLLER_BUTTON_MAX           = C.SDL_CONTROLLER_BUTTON_MAX
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

// GameControllerAddMapping (https://wiki.libsdl.org/SDL_GameControllerAddMapping)
func GameControllerAddMapping(mappingString string) int {
	_mappingString := C.CString(mappingString)
	defer C.free(unsafe.Pointer(_mappingString))
	return int(C.SDL_GameControllerAddMapping(_mappingString))
}

// GameControllerMappingForGUID (https://wiki.libsdl.org/SDL_GameControllerMappingForGUID)
func GameControllerMappingForGUID(guid JoystickGUID) string {
	return C.GoString(C.SDL_GameControllerMappingForGUID(guid.c()))
}

// GameControllerMapping (https://wiki.libsdl.org/SDL_GameControllerMapping)
func GameControllerMapping(ctrl *GameController) string {
	return C.GoString(C.SDL_GameControllerMapping(ctrl.cptr()))
}

// IsGameController (https://wiki.libsdl.org/SDL_IsGameController)
func IsGameController(index int) bool {
	return C.SDL_IsGameController(C.int(index)) > 0
}

// GameControllerNameForIndex (https://wiki.libsdl.org/SDL_GameControllerNameForIndex)
func GameControllerNameForIndex(index int) string {
	return C.GoString(C.SDL_GameControllerNameForIndex(C.int(index)))
}

// GameControllerOpen (https://wiki.libsdl.org/SDL_GameControllerOpen)
func GameControllerOpen(index int) *GameController {
	return (*GameController)(unsafe.Pointer(C.SDL_GameControllerOpen(C.int(index))))
}

// GameController (https://wiki.libsdl.org/SDL_GameControllerName)
func (ctrl *GameController) Name() string {
	return C.GoString(C.SDL_GameControllerName(ctrl.cptr()))
}

// GameController (https://wiki.libsdl.org/SDL_GameControllerGetAttached)
func (ctrl *GameController) GetAttached() bool {
	return C.SDL_GameControllerGetAttached(ctrl.cptr()) > 0
}

// GameController (https://wiki.libsdl.org/SDL_GameControllerGetJoystick)
func (ctrl *GameController) GetJoystick() *Joystick {
	return (*Joystick)(unsafe.Pointer(C.SDL_GameControllerGetJoystick(ctrl.cptr())))
}

// GameControllerEventState (https://wiki.libsdl.org/SDL_GameControllerEventState)
func GameControllerEventState(state int) int {
	return int(C.SDL_GameControllerEventState(C.int(state)))
}

// GameControllerUpdate (https://wiki.libsdl.org/SDL_GameControllerUpdate)
func GameControllerUpdate() {
	C.SDL_GameControllerUpdate()
}

// GameControllerGetAxisFromString (https://wiki.libsdl.org/SDL_GameControllerGetAxisFromString)
func GameControllerGetAxisFromString(pchString string) GameControllerAxis {
	_pchString := C.CString(pchString)
	defer C.free(unsafe.Pointer(_pchString))
	return GameControllerAxis(C.SDL_GameControllerGetAxisFromString(_pchString))
}

// GameControllerGetStringForAxis (https://wiki.libsdl.org/SDL_GameControllerGetStringForAxis)
func GameControllerGetStringForAxis(axis GameControllerAxis) string {
	return C.GoString(C.SDL_GameControllerGetStringForAxis(axis.c()))
}

// GameController (https://wiki.libsdl.org/SDL_GameControllerGetBindForAxis)
func (ctrl *GameController) GetBindForAxis(axis GameControllerAxis) GameControllerButtonBind {
	return GameControllerButtonBind(C.SDL_GameControllerGetBindForAxis(ctrl.cptr(), axis.c()))
}

// GameController (https://wiki.libsdl.org/SDL_GameControllerGetAxis)
func (ctrl *GameController) GetAxis(axis GameControllerAxis) int16 {
	return int16(C.SDL_GameControllerGetAxis(ctrl.cptr(), axis.c()))
}

// GameControllerGetButtonFromString (https://wiki.libsdl.org/SDL_GameControllerGetButtonFromString)
func GameControllerGetButtonFromString(pchString string) GameControllerButton {
	_pchString := C.CString(pchString)
	defer C.free(unsafe.Pointer(_pchString))
	return GameControllerButton(C.SDL_GameControllerGetButtonFromString(_pchString))
}

// GameControllerGetStringForButton (https://wiki.libsdl.org/SDL_GameControllerGetStringForButton)
func GameControllerGetStringForButton(btn GameControllerButton) string {
	return C.GoString(C.SDL_GameControllerGetStringForButton(btn.c()))
}

// GameController (https://wiki.libsdl.org/SDL_GameControllerGetBindForButton)
func (ctrl *GameController) GetBindForButton(btn GameControllerButton) GameControllerButtonBind {
	return GameControllerButtonBind(C.SDL_GameControllerGetBindForButton(ctrl.cptr(), btn.c()))
}

// GameController (https://wiki.libsdl.org/SDL_GameControllerGetButton)
func (ctrl *GameController) GetButton(btn GameControllerButton) byte {
	return byte(C.SDL_GameControllerGetButton(ctrl.cptr(), btn.c()))
}

// GameController (https://wiki.libsdl.org/SDL_GameControllerClose)
func (ctrl *GameController) Close() {
	C.SDL_GameControllerClose(ctrl.cptr())
}

func (bind *GameControllerButtonBind) Type() int {
	return int(bind.bindType)
}

func (bind *GameControllerButtonBind) Button() int {
	val, _ := binary.Varint(bind.value[:4])
	return int(val)
}

func (bind *GameControllerButtonBind) Axis() int {
	val, _ := binary.Varint(bind.value[:4])
	return int(val)
}

func (bind *GameControllerButtonBind) Hat() int {
	val, _ := binary.Varint(bind.value[:4])
	return int(val)
}

func (bind *GameControllerButtonBind) HatMask() int {
	val, _ := binary.Varint(bind.value[4:8])
	return int(val)
}
