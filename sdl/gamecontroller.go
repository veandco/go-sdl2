package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,4))

#if defined(WARN_OUTDATED)
#pragma message("SDL_GameControllerFromInstanceID is not supported before SDL 2.0.4")
#endif

static SDL_GameController* SDL_GameControllerFromInstanceID(SDL_JoystickID joyid)
{
	return NULL;
}

#endif


#if !(SDL_VERSION_ATLEAST(2,0,6))

#if defined(WARN_OUTDATED)
#pragma message("SDL_GameControllerGetVendor is not supported before SDL 2.0.6")
#pragma message("SDL_GameControllerGetProduct is not supported before SDL 2.0.6")
#pragma message("SDL_GameControllerGetProductVersion is not supported before SDL 2.0.6")
#pragma message("SDL_GameControllerNumMappings is not supported before SDL 2.0.6")
#pragma message("SDL_GameControllerMappingForIndex is not supported before SDL 2.0.6")
#endif

static Uint16 SDL_GameControllerGetVendor(SDL_GameController* gamecontroller)
{
	return 0;
}


static Uint16 SDL_GameControllerGetProduct(SDL_GameController* gamecontroller)
{
	return 0;
}

static Uint16 SDL_GameControllerGetProductVersion(SDL_GameController* gamecontroller)
{
	return 0;
}

static int SDL_GameControllerNumMappings(void)
{
	return 0;
}

static char* SDL_GameControllerMappingForIndex(int mapping_index)
{
	return NULL;
}
#endif


#if !(SDL_VERSION_ATLEAST(2,0,9))

#if defined(WARN_OUTDATED)
#pragma message("SDL_GameControllerGetPlayerIndex is not supported before SDL 2.0.9")
#pragma message("SDL_GameControllerRumble is not supported before SDL 2.0.9")
#pragma message("SDL_GameControllerMappingForDeviceIndex is not supported before SDL 2.0.9")
#endif

typedef enum
{
    SDL_SENSOR_INVALID = -1,
    SDL_SENSOR_UNKNOWN,
    SDL_SENSOR_ACCEL,
    SDL_SENSOR_GYRO
} SDL_SensorType;

static int SDL_GameControllerGetPlayerIndex(SDL_GameController *gamecontroller)
{
	return -1;
}

static int SDL_GameControllerRumble(SDL_GameController *gamecontroller, Uint16 low_frequency_rumble, Uint16 high_frequency_rumble, Uint32 duration_ms)
{
	return -1;
}

static char *SDL_GameControllerMappingForDeviceIndex(int joystick_index)
{
	return NULL;
}

#endif


#if !(SDL_VERSION_ATLEAST(2,0,12))

typedef enum
{
    SDL_CONTROLLER_TYPE_UNKNOWN = 0,
    SDL_CONTROLLER_TYPE_XBOX360,
    SDL_CONTROLLER_TYPE_XBOXONE,
    SDL_CONTROLLER_TYPE_PS3,
    SDL_CONTROLLER_TYPE_PS4,
    SDL_CONTROLLER_TYPE_NINTENDO_SWITCH_PRO
} SDL_GameControllerType;

#if defined(WARN_OUTDATED)
#pragma message("SDL_GameControllerTypeForIndex is not supported before SDL 2.0.12")
#pragma message("SDL_GameControllerGetType is not supported before SDL 2.0.12")
#pragma message("SDL_GameControllerFromPlayerIndex is not supported before SDL 2.0.12")
#pragma message("SDL_GameControllerSetPlayerIndex is not supported before SDL 2.0.12")
#endif

static SDL_GameControllerType SDL_GameControllerTypeForIndex(int joystick_index)
{
	return SDL_CONTROLLER_TYPE_UNKNOWN;
}

static SDL_GameControllerType SDL_GameControllerGetType(SDL_GameController *gamecontroller)
{
	return SDL_CONTROLLER_TYPE_UNKNOWN;
}

static SDL_GameController * SDL_GameControllerFromPlayerIndex(int player_index)
{
	return NULL;
}

static void SDL_GameControllerSetPlayerIndex(SDL_GameController *gamecontroller, int player_index)
{
	// do nothing
}

#endif


#if !(SDL_VERSION_ATLEAST(2,0,14))

#define SDL_CONTROLLER_TYPE_VIRTUAL (6)
#define SDL_CONTROLLER_TYPE_PS5 (7)

#if defined(WARN_OUTDATED)
#pragma message("SDL_GameControllerGetSerial is not supported before SDL 2.0.14")
#pragma message("SDL_GameControllerHasAxis is not supported before SDL 2.0.14")
#pragma message("SDL_GameControllerHasButton is not supported before SDL 2.0.14")
#pragma message("SDL_GameControllerGetNumTouchpads is not supported before SDL 2.0.14")
#pragma message("SDL_GameControllerGetNumTouchpadFingers is not supported before SDL 2.0.14")
#pragma message("SDL_GameControllerGetTouchpadFinger is not supported before SDL 2.0.14")
#pragma message("SDL_GameControllerHasSensor is not supported before SDL 2.0.14")
#pragma message("SDL_GameControllerSetSensorEnabled is not supported before SDL 2.0.14")
#pragma message("SDL_GameControllerIsSensorEnabled is not supported before SDL 2.0.14")
#pragma message("SDL_GameControllerGetSensorData is not supported before SDL 2.0.14")
#pragma message("SDL_GameControllerRumbleTriggers is not supported before SDL 2.0.14")
#pragma message("SDL_GameControllerHasLED is not supported before SDL 2.0.14")
#pragma message("SDL_GameControllerSetLED is not supported before SDL 2.0.14")
#endif

static const char * SDLCALL SDL_GameControllerGetSerial(SDL_GameController *gamecontroller)
{
	return NULL;
}

static SDL_bool SDL_GameControllerHasAxis(SDL_GameController *gamecontroller, SDL_GameControllerAxis axis)
{
	return SDL_FALSE;
}

static SDL_bool SDLCALL SDL_GameControllerHasButton(SDL_GameController *gamecontroller, SDL_GameControllerButton button)
{
	return SDL_FALSE;
}

static int SDLCALL SDL_GameControllerGetNumTouchpads(SDL_GameController *gamecontroller)
{
	return 0;
}

static int SDL_GameControllerGetNumTouchpadFingers(SDL_GameController *gamecontroller, int touchpad)
{
	return 0;
}

static int SDL_GameControllerGetTouchpadFinger(SDL_GameController *gamecontroller, int touchpad, int finger, Uint8 *state, float *x, float *y, float *pressure)
{
	return -1;
}

static SDL_bool SDL_GameControllerHasSensor(SDL_GameController *gamecontroller, SDL_SensorType type)
{
	return SDL_FALSE;
}

static int SDL_GameControllerSetSensorEnabled(SDL_GameController *gamecontroller, SDL_SensorType type, SDL_bool enabled)
{
	return -1;
}

static SDL_bool SDL_GameControllerIsSensorEnabled(SDL_GameController *gamecontroller, SDL_SensorType type)
{
	return SDL_FALSE;
}

static int SDL_GameControllerGetSensorData(SDL_GameController *gamecontroller, SDL_SensorType type, float *data, int num_values)
{
	return -1;
}

static int SDL_GameControllerRumbleTriggers(SDL_GameController *gamecontroller, Uint16 left_rumble, Uint16 right_rumble, Uint32 duration_ms)
{
	return -1;
}

static SDL_bool SDL_GameControllerHasLED(SDL_GameController *gamecontroller)
{
	return SDL_FALSE;
}

static int SDL_GameControllerSetLED(SDL_GameController *gamecontroller, Uint8 red, Uint8 green, Uint8 blue)
{
	return -1;
}

#endif


#if !(SDL_VERSION_ATLEAST(2,0,16))

#if defined(WARN_OUTDATED)
#pragma message("SDL_GameControllerSendEffect is not supported before SDL 2.0.16")
#pragma message("SDL_GameControllerGetSensorDataRate is not supported before SDL 2.0.16")
#endif

static int SDL_GameControllerSendEffect(SDL_GameController *gamecontroller, const void *data, int size)
{
	return -1;
}

static float SDL_GameControllerGetSensorDataRate(SDL_GameController *gamecontroller, SDL_SensorType type)
{
	return 0.0f;
}

#endif

#if !(SDL_VERSION_ATLEAST(2,0,18))

#if defined(WARN_OUTDATED)
#pragma message("SDL_GameControllerHasRumble is not supported before SDL 2.0.18")
#pragma message("SDL_GameControllerHasRumbleTriggers is not supported before SDL 2.0.18")
#pragma message("SDL_GameControllerGetAppleSFSymbolsNameForButton is not supported before SDL 2.0.18")
#pragma message("SDL_GameControllerGetAppleSFSymbolsNameForAxis is not supported before SDL 2.0.18")
#endif

static SDL_bool SDL_GameControllerHasRumble(SDL_GameController *gamecontroller)
{
	return SDL_FALSE;
}

static SDL_bool SDL_GameControllerHasRumbleTriggers(SDL_GameController *gamecontroller)
{
	return SDL_FALSE;
}

static const char* SDL_GameControllerGetAppleSFSymbolsNameForButton(SDL_GameController *gamecontroller, SDL_GameControllerButton button)
{
	return NULL;
}

static const char* SDL_GameControllerGetAppleSFSymbolsNameForAxis(SDL_GameController *gamecontroller, SDL_GameControllerAxis axis)
{
	return NULL;
}

#endif
*/
import "C"
import (
	"encoding/binary"
	"unsafe"
)

// Types of game controller inputs.
const (
	CONTROLLER_BINDTYPE_NONE   = C.SDL_CONTROLLER_BINDTYPE_NONE
	CONTROLLER_BINDTYPE_BUTTON = C.SDL_CONTROLLER_BINDTYPE_BUTTON
	CONTROLLER_BINDTYPE_AXIS   = C.SDL_CONTROLLER_BINDTYPE_AXIS
	CONTROLLER_BINDTYPE_HAT    = C.SDL_CONTROLLER_BINDTYPE_HAT
)

// An enumeration of axes available from a controller.
// (https://wiki.libsdl.org/SDL_GameControllerAxis)
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

// An enumeration of buttons available from a controller.
// (https://wiki.libsdl.org/SDL_GameControllerButton)
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

// GameControllerBindType is a type of game controller input.
type GameControllerBindType C.SDL_GameControllerBindType

// GameControllerAxis is an axis on a game controller.
// (https://wiki.libsdl.org/SDL_GameControllerAxis)
type GameControllerAxis C.SDL_GameControllerAxis

// GameControllerButton is a button on a game controller.
// (https://wiki.libsdl.org/SDL_GameControllerButton)
type GameControllerButton C.SDL_GameControllerButton

// GameController used to identify an SDL game controller.
type GameController C.SDL_GameController

// GameControllerButtonBind SDL joystick layer binding for controller button/axis mapping.
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

// GameControllerAddMapping adds support for controllers that SDL is unaware of or to cause an existing controller to have a different binding.
// (https://wiki.libsdl.org/SDL_GameControllerAddMapping)
func GameControllerAddMapping(mappingString string) int {
	_mappingString := C.CString(mappingString)
	defer C.free(unsafe.Pointer(_mappingString))
	return int(C.SDL_GameControllerAddMapping(_mappingString))
}

// GameControllerNumMappings returns the number of mappings installed.
func GameControllerNumMappings() int {
	return int(C.SDL_GameControllerNumMappings())
}

// GameControllerMappingForIndex returns the game controller mapping string at a particular index.
func GameControllerMappingForIndex(index int) string {
	mappingString := C.SDL_GameControllerMappingForIndex(C.int(index))
	defer C.free(unsafe.Pointer(mappingString))
	return C.GoString(mappingString)
}

// GameControllerMappingForGUID returns the game controller mapping string for a given GUID.
// (https://wiki.libsdl.org/SDL_GameControllerMappingForGUID)
func GameControllerMappingForGUID(guid JoystickGUID) string {
	mappingString := C.SDL_GameControllerMappingForGUID(guid.c())
	defer C.free(unsafe.Pointer(mappingString))
	return C.GoString(mappingString)
}

// IsGameController reports whether the given joystick is supported by the game controller interface.
// (https://wiki.libsdl.org/SDL_IsGameController)
func IsGameController(index int) bool {
	return C.SDL_IsGameController(C.int(index)) == C.SDL_TRUE
}

// GameControllerNameForIndex returns the implementation dependent name for the game controller.
// (https://wiki.libsdl.org/SDL_GameControllerNameForIndex)
func GameControllerNameForIndex(index int) string {
	return C.GoString(C.SDL_GameControllerNameForIndex(C.int(index)))
}

// GameControllerMappingForDeviceIndex returns the game controller mapping string at a particular index.
func GameControllerMappingForDeviceIndex(index int) string {
	mappingString := C.SDL_GameControllerMappingForDeviceIndex(C.int(index))
	defer C.free(unsafe.Pointer(mappingString))
	return C.GoString(mappingString)
}

// GameControllerOpen opens a gamecontroller for use.
// (https://wiki.libsdl.org/SDL_GameControllerOpen)
func GameControllerOpen(index int) *GameController {
	return (*GameController)(C.SDL_GameControllerOpen(C.int(index)))
}

// GameControllerFromInstanceID returns the GameController associated with an instance id.
// (https://wiki.libsdl.org/SDL_GameControllerFromInstanceID)
func GameControllerFromInstanceID(joyid JoystickID) *GameController {
	return (*GameController)(C.SDL_GameControllerFromInstanceID(joyid.c()))
}

// Name returns the implementation dependent name for an opened game controller.
// (https://wiki.libsdl.org/SDL_GameControllerName)
func (ctrl *GameController) Name() string {
	return C.GoString(C.SDL_GameControllerName(ctrl.cptr()))
}

// PlayerIndex the player index of an opened game controller, or -1 if it's not available.
// TODO: (https://wiki.libsdl.org/SDL_GameControllerGetPlayerIndex)
func (ctrl *GameController) PlayerIndex() int {
	return int(C.SDL_GameControllerGetPlayerIndex(ctrl.cptr()))
}

// Vendor returns the USB vendor ID of an opened controller, if available, 0 otherwise.
func (ctrl *GameController) Vendor() int {
	return int(C.SDL_GameControllerGetVendor(ctrl.cptr()))
}

// Product returns the USB product ID of an opened controller, if available, 0 otherwise.
func (ctrl *GameController) Product() int {
	return int(C.SDL_GameControllerGetProduct(ctrl.cptr()))
}

// ProductVersion returns the product version of an opened controller, if available, 0 otherwise.
func (ctrl *GameController) ProductVersion() int {
	return int(C.SDL_GameControllerGetProductVersion(ctrl.cptr()))
}

// Attached reports whether a controller has been opened and is currently connected.
// (https://wiki.libsdl.org/SDL_GameControllerGetAttached)
func (ctrl *GameController) Attached() bool {
	return C.SDL_GameControllerGetAttached(ctrl.cptr()) == C.SDL_TRUE
}

// Mapping returns the current mapping of a Game Controller.
// (https://wiki.libsdl.org/SDL_GameControllerMapping)
func (ctrl *GameController) Mapping() string {
	mappingString := C.SDL_GameControllerMapping(ctrl.cptr())
	defer C.free(unsafe.Pointer(mappingString))
	return C.GoString(mappingString)
}

// Joystick returns the Joystick ID from a Game Controller. The game controller builds on the Joystick API, but to be able to use the Joystick's functions with a gamepad, you need to use this first to get the joystick object.
// (https://wiki.libsdl.org/SDL_GameControllerGetJoystick)
func (ctrl *GameController) Joystick() *Joystick {
	return (*Joystick)(unsafe.Pointer(C.SDL_GameControllerGetJoystick(ctrl.cptr())))
}

// GameControllerEventState returns the current state of, enable, or disable events dealing with Game Controllers. This will not disable Joystick events, which can also be fired by a controller (see https://wiki.libsdl.org/SDL_JoystickEventState).
// (https://wiki.libsdl.org/SDL_GameControllerEventState)
func GameControllerEventState(state int) int {
	return int(C.SDL_GameControllerEventState(C.int(state)))
}

// GameControllerUpdate manually pumps game controller updates if not using the loop.
// (https://wiki.libsdl.org/SDL_GameControllerUpdate)
func GameControllerUpdate() {
	C.SDL_GameControllerUpdate()
}

// GameControllerGetAxisFromString converts a string into an enum representation for a GameControllerAxis.
// (https://wiki.libsdl.org/SDL_GameControllerGetAxisFromString)
func GameControllerGetAxisFromString(pchString string) GameControllerAxis {
	_pchString := C.CString(pchString)
	defer C.free(unsafe.Pointer(_pchString))
	return GameControllerAxis(C.SDL_GameControllerGetAxisFromString(_pchString))
}

// GameControllerGetStringForAxis converts from an axis enum to a string.
// (https://wiki.libsdl.org/SDL_GameControllerGetStringForAxis)
func GameControllerGetStringForAxis(axis GameControllerAxis) string {
	return C.GoString(C.SDL_GameControllerGetStringForAxis(axis.c()))
}

// BindForAxis returns the SDL joystick layer binding for a controller button mapping.
// (https://wiki.libsdl.org/SDL_GameControllerGetBindForAxis)
func (ctrl *GameController) BindForAxis(axis GameControllerAxis) GameControllerButtonBind {
	return GameControllerButtonBind(C.SDL_GameControllerGetBindForAxis(ctrl.cptr(), axis.c()))
}

// Axis returns the current state of an axis control on a game controller.
// (https://wiki.libsdl.org/SDL_GameControllerGetAxis)
func (ctrl *GameController) Axis(axis GameControllerAxis) int16 {
	return int16(C.SDL_GameControllerGetAxis(ctrl.cptr(), axis.c()))
}

// GameControllerGetButtonFromString turns a string into a button mapping.
// (https://wiki.libsdl.org/SDL_GameControllerGetButtonFromString)
func GameControllerGetButtonFromString(pchString string) GameControllerButton {
	_pchString := C.CString(pchString)
	defer C.free(unsafe.Pointer(_pchString))
	return GameControllerButton(C.SDL_GameControllerGetButtonFromString(_pchString))
}

// GameControllerGetStringForButton turns a button enum into a string mapping.
// (https://wiki.libsdl.org/SDL_GameControllerGetStringForButton)
func GameControllerGetStringForButton(btn GameControllerButton) string {
	return C.GoString(C.SDL_GameControllerGetStringForButton(btn.c()))
}

// BindForButton returns the SDL joystick layer binding for this controller button mapping.
// (https://wiki.libsdl.org/SDL_GameControllerGetBindForButton)
func (ctrl *GameController) BindForButton(btn GameControllerButton) GameControllerButtonBind {
	return GameControllerButtonBind(C.SDL_GameControllerGetBindForButton(ctrl.cptr(), btn.c()))
}

// Rumble triggers a rumble effect
// Each call to this function cancels any previous rumble effect, and calling it with 0 intensity stops any rumbling.
//
// lowFrequencyRumble - The intensity of the low frequency (left) rumble motor, from 0 to 0xFFFF
// highFrequencyRumble - The intensity of the high frequency (right) rumble motor, from 0 to 0xFFFF
// durationMS - The duration of the rumble effect, in milliseconds
//
// Returns error if rumble isn't supported on this joystick.
//
// TODO: (https://wiki.libsdl.org/SDL_GameControllerRumble)
func (ctrl *GameController) Rumble(lowFrequencyRumble, highFrequencyRumble uint16, durationMS uint32) error {
	return errorFromInt(int(C.SDL_GameControllerRumble(ctrl.cptr(), C.Uint16(lowFrequencyRumble), C.Uint16(highFrequencyRumble), C.Uint32(durationMS))))
}

// Button returns the current state of a button on a game controller.
// (https://wiki.libsdl.org/SDL_GameControllerGetButton)
func (ctrl *GameController) Button(btn GameControllerButton) byte {
	return byte(C.SDL_GameControllerGetButton(ctrl.cptr(), btn.c()))
}

// Close closes a game controller previously opened with GameControllerOpen().
// (https://wiki.libsdl.org/SDL_GameControllerClose)
func (ctrl *GameController) Close() {
	C.SDL_GameControllerClose(ctrl.cptr())
}

// Type returns the type of game controller input for this SDL joystick layer binding.
func (bind *GameControllerButtonBind) Type() int {
	return int(bind.bindType)
}

// Button returns button mapped for this SDL joystick layer binding.
func (bind *GameControllerButtonBind) Button() int {
	val, _ := binary.Varint(bind.value[:4])
	return int(val)
}

// Axis returns axis mapped for this SDL joystick layer binding.
func (bind *GameControllerButtonBind) Axis() int {
	val, _ := binary.Varint(bind.value[:4])
	return int(val)
}

// Hat returns hat mapped for this SDL joystick layer binding.
func (bind *GameControllerButtonBind) Hat() int {
	val, _ := binary.Varint(bind.value[:4])
	return int(val)
}

// HatMask returns hat mask for this SDL joystick layer binding.
func (bind *GameControllerButtonBind) HatMask() int {
	val, _ := binary.Varint(bind.value[4:8])
	return int(val)
}

// SendEffect sends a controller specific effect packet.
// (https://wiki.libsdl.org/SDL_GameControllerSendEffect)
func (ctrl *GameController) SendEffect(data []byte) (err error) {
	_size := C.int(len(data))
	return errorFromInt(int(C.SDL_GameControllerSendEffect(ctrl.cptr(), unsafe.Pointer(&data[0]), _size)))
}

// GetSensorDataRate gets the data rate (number of events per second) of a game controller sensor.
// (https://wiki.libsdl.org/SDL_GameControllerGetSensorDataRate)
func (ctrl *GameController) SensorDataRate(typ SensorType) (rate float32) {
	return float32(C.SDL_GameControllerGetSensorDataRate(ctrl.cptr(), C.SDL_SensorType(typ)))
}

// HasRumble queries whether a game controller has rumble support.
// (https://wiki.libsdl.org/SDL_GameControllerHasRumble)
func (ctrl *GameController) HasRumble() bool {
	return C.SDL_GameControllerHasRumble(ctrl.cptr()) == C.SDL_TRUE
}

// HasRumbleTriggers queries whether a game controller has rumble support on triggers.
// (https://wiki.libsdl.org/SDL_GameControllerHasRumbleTriggers)
func (ctrl *GameController) HasRumbleTriggers() bool {
	return C.SDL_GameControllerHasRumbleTriggers(ctrl.cptr()) == C.SDL_TRUE
}

// GetAppleSFSymbolsNameForButton returns the sfSymbolsName for a given button on a game controller on Apple platforms.
// (https://wiki.libsdl.org/SDL_GameControllerGetAppleSFSymbolsNameForButton)
func (ctrl *GameController) GetAppleSFSymbolsNameForButton(button GameControllerButton) (sfSymbolsName string) {
	_button := C.SDL_GameControllerButton(button)
	_sfSymbolsName := C.SDL_GameControllerGetAppleSFSymbolsNameForButton(ctrl.cptr(), _button)
	sfSymbolsName = C.GoString(_sfSymbolsName)
	return
}

// GetAppleSFSymbolsNameForAxis returns the sfSymbolsName for a given axis on a game controller on Apple platforms.
// (https://wiki.libsdl.org/SDL_GameControllerGetAppleSFSymbolsNameForAxis)
func (ctrl *GameController) SDL_GameControllerGetAppleSFSymbolsNameForAxis(axis GameControllerAxis) (sfSymbolsName string) {
	_axis := C.SDL_GameControllerAxis(axis)
	_sfSymbolsName := C.SDL_GameControllerGetAppleSFSymbolsNameForAxis(ctrl.cptr(), _axis)
	sfSymbolsName = C.GoString(_sfSymbolsName)
	return
}
