package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,0,4))

#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickPowerLevel is not supported before SDL 2.0.4")
#endif

typedef enum
{
    SDL_JOYSTICK_POWER_UNKNOWN = -1,
    SDL_JOYSTICK_POWER_EMPTY,
    SDL_JOYSTICK_POWER_LOW,
    SDL_JOYSTICK_POWER_MEDIUM,
    SDL_JOYSTICK_POWER_FULL,
    SDL_JOYSTICK_POWER_WIRED,
    SDL_JOYSTICK_POWER_MAX
} SDL_JoystickPowerLevel;


#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickCurrentPowerLevel is not supported before SDL 2.0.4")
#endif

static SDL_JoystickPowerLevel SDL_JoystickCurrentPowerLevel(SDL_Joystick* joystick)
{
	return SDL_JOYSTICK_POWER_UNKNOWN;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickFromInstanceID is not supported before SDL 2.0.4")
#endif

static SDL_Joystick* SDL_JoystickFromInstanceID(SDL_JoystickID joyid)
{
	return NULL;
}
#endif

#if !(SDL_VERSION_ATLEAST(2,0,6))

#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickType is not supported before SDL 2.0.6")
#endif

typedef enum
{
	SDL_JOYSTICK_TYPE_UNKNOWN,
	SDL_JOYSTICK_TYPE_GAMECONTROLLER,
	SDL_JOYSTICK_TYPE_WHEEL,
	SDL_JOYSTICK_TYPE_ARCADE_STICK,
	SDL_JOYSTICK_TYPE_FLIGHT_STICK,
	SDL_JOYSTICK_TYPE_DANCE_PAD,
	SDL_JOYSTICK_TYPE_GUITAR,
	SDL_JOYSTICK_TYPE_DRUM_KIT,
	SDL_JOYSTICK_TYPE_ARCADE_PAD,
	SDL_JOYSTICK_TYPE_THROTTLE
} SDL_JoystickType;


#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickGetDeviceVendor is not supported before SDL 2.0.6")
#endif

static Uint16 SDL_JoystickGetDeviceVendor(int device_index)
{
	return 0;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickGetDeviceProduct is not supported before SDL 2.0.6")
#endif

static Uint16 SDL_JoystickGetDeviceProduct(int device_index)
{
	return 0;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickGetDeviceProductVersion is not supported before SDL 2.0.6")
#endif

static Uint16 SDL_JoystickGetDeviceProductVersion(int device_index)
{
	return 0;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickGetDeviceType is not supported before SDL 2.0.6")
#endif

static SDL_JoystickType SDL_JoystickGetDeviceType(int device_index)
{
	return SDL_JOYSTICK_TYPE_UNKNOWN;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickGetDeviceInstanceID is not supported before SDL 2.0.6")
#endif

static SDL_JoystickID SDL_JoystickGetDeviceInstanceID(int device_index)
{
	return 0;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickGetVendor is not supported before SDL 2.0.6")
#endif

static Uint16 SDL_JoystickGetVendor(SDL_Joystick* joystick)
{
	return 0;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickGetProduct is not supported before SDL 2.0.6")
#endif

static Uint16 SDL_JoystickGetProduct(SDL_Joystick* joystick)
{
	return 0;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickGetProductVersion is not supported before SDL 2.0.6")
#endif

static Uint16 SDL_JoystickGetProductVersion(SDL_Joystick* joystick)
{
	return 0;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickGetType is not supported before SDL 2.0.6")
#endif

static SDL_JoystickType SDL_JoystickGetType(SDL_Joystick* joystick)
{
	return SDL_JOYSTICK_TYPE_UNKNOWN;
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickGetAxisInitialState is not supported before SDL 2.0.6")
#endif

static SDL_bool SDL_JoystickGetAxisInitialState(SDL_Joystick* joystick, int axis, Sint16* state)
{
	return SDL_FALSE;
}
#endif

#if !(SDL_VERSION_ATLEAST(2,0,7))


#if defined(WARN_OUTDATED)
#pragma message("SDL_LockJoysticks is not supported before SDL 2.0.7")
#endif

static void SDL_LockJoysticks()
{
}


#if defined(WARN_OUTDATED)
#pragma message("SDL_UnlockJoysticks is not supported before SDL 2.0.7")
#endif

static void SDL_UnlockJoysticks()
{
}

#endif

#if !(SDL_VERSION_ATLEAST(2,0,9))


#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickGetDevicePlayerIndex is not supported before SDL 2.0.9")
#endif

static int SDL_JoystickGetDevicePlayerIndex(int device_index)
{
	return 0;
}

#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickGetPlayerIndex is not supported before SDL 2.0.9")
#endif

static int SDL_JoystickGetPlayerIndex(SDL_Joystick *joystick)
{
	return 0;
}

#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickRumble is not supported before SDL 2.0.9")
#endif

static int SDL_JoystickRumble(SDL_Joystick *joystick, Uint16 low_frequency_rumble, Uint16 high_frequency_rumble, Uint32 duration_ms)
{
	return -1;
}

#endif

#if !(SDL_VERSION_ATLEAST(2,0,18))

#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickHasRumble is not supported before SDL 2.0.18")
#pragma message("SDL_JoystickHasRumbleTriggers is not supported before SDL 2.0.18")
#endif

static SDL_bool SDL_JoystickHasRumble(SDL_Joystick *joystick)
{
	return SDL_FALSE;
}

static SDL_bool SDL_JoystickHasRumbleTriggers(SDL_Joystick *joystick)
{
	return SDL_FALSE;
}

#endif
*/
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

// Types of a joystick.
const (
	JOYSTICK_TYPE_UNKNOWN        = C.SDL_JOYSTICK_TYPE_UNKNOWN
	JOYSTICK_TYPE_GAMECONTROLLER = C.SDL_JOYSTICK_TYPE_GAMECONTROLLER
	JOYSTICK_TYPE_WHEEL          = C.SDL_JOYSTICK_TYPE_WHEEL
	JOYSTICK_TYPE_ARCADE_STICK   = C.SDL_JOYSTICK_TYPE_ARCADE_STICK
	JOYSTICK_TYPE_FLIGHT_STICK   = C.SDL_JOYSTICK_TYPE_FLIGHT_STICK
	JOYSTICK_TYPE_DANCE_PAD      = C.SDL_JOYSTICK_TYPE_DANCE_PAD
	JOYSTICK_TYPE_GUITAR         = C.SDL_JOYSTICK_TYPE_GUITAR
	JOYSTICK_TYPE_DRUM_KIT       = C.SDL_JOYSTICK_TYPE_DRUM_KIT
	JOYSTICK_TYPE_ARCADE_PAD     = C.SDL_JOYSTICK_TYPE_ARCADE_PAD
	JOYSTICK_TYPE_THROTTLE       = C.SDL_JOYSTICK_TYPE_THROTTLE
)

// An enumeration of battery levels of a joystick.
// (https://wiki.libsdl.org/SDL_JoystickPowerLevel)
const (
	JOYSTICK_POWER_UNKNOWN = C.SDL_JOYSTICK_POWER_UNKNOWN
	JOYSTICK_POWER_EMPTY   = C.SDL_JOYSTICK_POWER_EMPTY
	JOYSTICK_POWER_LOW     = C.SDL_JOYSTICK_POWER_LOW
	JOYSTICK_POWER_MEDIUM  = C.SDL_JOYSTICK_POWER_MEDIUM
	JOYSTICK_POWER_FULL    = C.SDL_JOYSTICK_POWER_FULL
	JOYSTICK_POWER_WIRED   = C.SDL_JOYSTICK_POWER_WIRED
	JOYSTICK_POWER_MAX     = C.SDL_JOYSTICK_POWER_MAX
)

// Joystick is an SDL joystick.
type Joystick C.SDL_Joystick

// JoystickGUID is a stable unique id for a joystick device.
type JoystickGUID C.SDL_JoystickGUID

// JoystickID is joystick's instance id.
type JoystickID C.SDL_JoystickID

// JoystickType is a type of a joystick.
type JoystickType C.SDL_JoystickType

// JoystickPowerLevel is a battery level of a joystick.
type JoystickPowerLevel C.SDL_JoystickPowerLevel

func (joy *Joystick) cptr() *C.SDL_Joystick {
	return (*C.SDL_Joystick)(unsafe.Pointer(joy))
}

func (guid JoystickGUID) c() C.SDL_JoystickGUID {
	return C.SDL_JoystickGUID(guid)
}

func (joyid JoystickID) c() C.SDL_JoystickID {
	return C.SDL_JoystickID(joyid)
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

// JoystickGetDevicePlayerIndex returns the player index of a joystick, or -1 if it's not available
// TODO: (https://wiki.libsdl.org/SDL_JoystickGetDevicePlayerIndex)
func JoystickGetDevicePlayerIndex(index int) int {
	return int(C.SDL_JoystickGetDevicePlayerIndex(C.int(index)))
}

// JoystickGetDeviceGUID returns the implementation dependent GUID for the joystick at a given device index.
// (https://wiki.libsdl.org/SDL_JoystickGetDeviceGUID)
func JoystickGetDeviceGUID(index int) JoystickGUID {
	return (JoystickGUID)(C.SDL_JoystickGetDeviceGUID(C.int(index)))
}

// JoystickGetDeviceVendor returns the USB vendor ID of a joystick, if available, 0 otherwise.
func JoystickGetDeviceVendor(index int) int {
	return int(C.SDL_JoystickGetDeviceVendor(C.int(index)))
}

// JoystickGetDeviceProduct returns the USB product ID of a joystick, if available, 0 otherwise.
func JoystickGetDeviceProduct(index int) int {
	return int(C.SDL_JoystickGetDeviceProduct(C.int(index)))
}

// JoystickGetDeviceProductVersion returns the product version of a joystick, if available, 0 otherwise.
func JoystickGetDeviceProductVersion(index int) int {
	return int(C.SDL_JoystickGetDeviceProductVersion(C.int(index)))
}

// JoystickGetDeviceType returns the type of a joystick.
func JoystickGetDeviceType(index int) JoystickType {
	return JoystickType(C.SDL_JoystickGetDeviceType(C.int(index)))
}

// JoystickGetDeviceInstanceID returns the instance ID of a joystick.
func JoystickGetDeviceInstanceID(index int) JoystickID {
	return JoystickID(C.SDL_JoystickGetDeviceInstanceID(C.int(index)))
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

// JoystickOpen opens a joystick for use.
// (https://wiki.libsdl.org/SDL_JoystickOpen)
func JoystickOpen(index int) *Joystick {
	return (*Joystick)(C.SDL_JoystickOpen(C.int(index)))
}

// JoystickFromInstanceID returns the Joystick associated with an instance id.
// (https://wiki.libsdl.org/SDL_JoystickFromInstanceID)
func JoystickFromInstanceID(joyid JoystickID) *Joystick {
	return (*Joystick)(C.SDL_JoystickFromInstanceID(joyid.c()))
}

// LockJoysticks locks joysticks for multi-threaded access to the joystick API
// TODO: (https://wiki.libsdl.org/SDL_LockJoysticks)
func LockJoysticks() {
	C.SDL_LockJoysticks()
}

// UnlockJoysticks unlocks joysticks for multi-threaded access to the joystick API
// TODO: (https://wiki.libsdl.org/SDL_UnlockJoysticks)
func UnlockJoysticks() {
	C.SDL_UnlockJoysticks()
}

// Name returns the implementation dependent name of a joystick.
// (https://wiki.libsdl.org/SDL_JoystickName)
func (joy *Joystick) Name() string {
	return (C.GoString)(C.SDL_JoystickName(joy.cptr()))
}

// PlayerIndex returns the player index of an opened joystick, or -1 if it's not available.
// (https://wiki.libsdl.org/SDL_JoystickGetPlayerIndex)
func (joy *Joystick) PlayerIndex() int {
	return int(C.SDL_JoystickGetPlayerIndex(joy.cptr()))
}

// GUID returns the implementation-dependent GUID for the joystick.
// (https://wiki.libsdl.org/SDL_JoystickGetGUID)
func (joy *Joystick) GUID() JoystickGUID {
	return (JoystickGUID)(C.SDL_JoystickGetGUID(joy.cptr()))
}

// Vendor returns the USB vendor ID of an opened joystick, if available, 0 otherwise.
func (joy *Joystick) Vendor() int {
	return int(C.SDL_JoystickGetVendor(joy.cptr()))
}

// Product returns the USB product ID of an opened joystick, if available, 0 otherwise.
func (joy *Joystick) Product() int {
	return int(C.SDL_JoystickGetProduct(joy.cptr()))
}

// ProductVersion returns the product version of an opened joystick, if available, 0 otherwise.
func (joy *Joystick) ProductVersion() int {
	return int(C.SDL_JoystickGetProductVersion(joy.cptr()))
}

// Type returns the the type of an opened joystick.
func (joy *Joystick) Type() JoystickType {
	return JoystickType(C.SDL_JoystickGetType(joy.cptr()))
}

// Attached returns the status of a specified joystick.
// (https://wiki.libsdl.org/SDL_JoystickGetAttached)
func (joy *Joystick) Attached() bool {
	return C.SDL_JoystickGetAttached(joy.cptr()) == C.SDL_TRUE
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

// Axis returns the current state of an axis control on a joystick.
// (https://wiki.libsdl.org/SDL_JoystickGetAxis)
func (joy *Joystick) Axis(axis int) int16 {
	return (int16)(C.SDL_JoystickGetAxis(joy.cptr(), C.int(axis)))
}

// AxisInitialState returns the initial state of an axis control on a joystick, ok is true if this axis has any initial value.
func (joy *Joystick) AxisInitialState(axis int) (state int16, ok bool) {
	ok = C.SDL_JoystickGetAxisInitialState(joy.cptr(), C.int(axis), (*C.Sint16)(&state)) == C.SDL_TRUE
	return
}

// Hat returns the current state of a POV hat on a joystick.
// (https://wiki.libsdl.org/SDL_JoystickGetHat)
func (joy *Joystick) Hat(hat int) byte {
	return (byte)(C.SDL_JoystickGetHat(joy.cptr(), C.int(hat)))
}

// Ball returns the ball axis change since the last poll.
// (https://wiki.libsdl.org/SDL_JoystickGetBall)
func (joy *Joystick) Ball(ball int, dx, dy *int32) int {
	_dx := (*C.int)(unsafe.Pointer(dx))
	_dy := (*C.int)(unsafe.Pointer(dy))
	return (int)(C.SDL_JoystickGetBall(joy.cptr(), C.int(ball), _dx, _dy))
}

// Button the current state of a button on a joystick.
// (https://wiki.libsdl.org/SDL_JoystickGetButton)
func (joy *Joystick) Button(button int) byte {
	return (byte)(C.SDL_JoystickGetButton(joy.cptr(), C.int(button)))
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
// TODO: (https://wiki.libsdl.org/SDL_JoystickRumble)
func (joy *Joystick) Rumble(lowFrequencyRumble, highFrequencyRumble uint16, durationMS uint32) error {
	return errorFromInt(int(C.SDL_JoystickRumble(joy.cptr(), C.Uint16(lowFrequencyRumble), C.Uint16(highFrequencyRumble), C.Uint32(durationMS))))
}

// Close closes a joystick previously opened with JoystickOpen().
// (https://wiki.libsdl.org/SDL_JoystickClose)
func (joy *Joystick) Close() {
	C.SDL_JoystickClose(joy.cptr())
}

// CurrentPowerLevel returns the battery level of a joystick as JoystickPowerLevel.
// (https://wiki.libsdl.org/SDL_JoystickCurrentPowerLevel)
func (joy *Joystick) CurrentPowerLevel() JoystickPowerLevel {
	return JoystickPowerLevel(C.SDL_JoystickCurrentPowerLevel(joy.cptr()))
}

// HasRumble queries whether a game controller has rumble support.
// (https://wiki.libsdl.org/SDL_JoystickHasRumble)
func (ctrl *Joystick) HasRumble() bool {
	return C.SDL_JoystickHasRumble(ctrl.cptr()) == C.SDL_TRUE
}

// HasRumbleTriggers queries whether a game controller has rumble support on triggers.
// (https://wiki.libsdl.org/SDL_JoystickHasRumbleTriggers)
func (ctrl *Joystick) HasRumbleTriggers() bool {
	return C.SDL_JoystickHasRumbleTriggers(ctrl.cptr()) == C.SDL_TRUE
}
