package sdl

/*
#include "sdl_wrapper.h"

#if !(SDL_VERSION_ATLEAST(2,26,0))
#if defined(WARN_OUTDATED)
#pragma message("SDL_GetJoystickGUIDInfo is not supported before SDL 2.26.0")
#endif

static inline void SDL_GetJoystickGUIDInfo(SDL_JoystickGUID guid, Uint16 *vendor, Uint16 *product, Uint16 *version, Uint16 *crc16)
{
    // do nothing
}
#endif

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

#if !(SDL_VERSION_ATLEAST(2,0,14))

#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickGetSerial is not supported before SDL 2.0.14")
#pragma message("SDL_JoystickRumbleTriggers is not supported before SDL 2.0.14")
#pragma message("SDL_JoystickHasLED is not supported before SDL 2.0.14")
#pragma message("SDL_JoystickSetLED is not supported before SDL 2.0.14")
#pragma message("SDL_JoystickAttachVirtual is not supported before SDL 2.0.14")
#pragma message("SDL_JoystickDetachVirtual is not supported before SDL 2.0.14")
#pragma message("SDL_JoystickIsVirtual is not supported before SDL 2.0.14")
#pragma message("SDL_JoystickSetVirtualAxis is not supported before SDL 2.0.14")
#pragma message("SDL_JoystickSetVirtualButton is not supported before SDL 2.0.14")
#pragma message("SDL_JoystickSetVirtualHat is not supported before SDL 2.0.14")
#endif

static inline const char * SDL_JoystickGetSerial(SDL_Joystick *joystick)
{
	return NULL;
}

static inline int SDL_JoystickRumbleTriggers(SDL_Joystick *joystick, Uint16 left_rumble, Uint16 right_rumble, Uint32 duration_ms)
{
	return -1;
}

static SDL_bool SDL_JoystickHasLED(SDL_Joystick *joystick)
{
	return SDL_FALSE;
}

static int SDL_JoystickSetLED(SDL_Joystick *joystick, Uint8 red, Uint8 green, Uint8 blue)
{
	return -1;
}

static int SDL_JoystickAttachVirtual(SDL_JoystickType type, int naxes, int nbuttons, int nhats)
{
	return -1;
}

static int SDL_JoystickDetachVirtual(int device_index)
{
	return -1;
}

static SDL_bool SDL_JoystickIsVirtual(int device_index)
{
	return SDL_FALSE;
}

static int SDL_JoystickSetVirtualAxis(SDL_Joystick *joystick, int axis, Sint16 value)
{
	return -1;
}

static int SDL_JoystickSetVirtualButton(SDL_Joystick *joystick, int button, Uint8 value)
{
	return -1;
}

static int SDL_JoystickSetVirtualHat(SDL_Joystick *joystick, int hat, Uint8 value)
{
	return -1;
}

#endif

#if !(SDL_VERSION_ATLEAST(2,0,12))

#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickFromPlayerIndex is not supported before SDL 2.0.12")
#pragma message("SDL_JoystickSetPlayerIndex is not supported before SDL 2.0.12")
#endif

static inline SDL_Joystick *SDL_JoystickFromPlayerIndex(int player_index)
{
	return NULL;
}

static inline void SDL_JoystickSetPlayerIndex(SDL_Joystick *joystick, int player_index)
{
	// do nothing
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

#if !(SDL_VERSION_ATLEAST(2,24,0))

#if defined(WARN_OUTDATED)
#pragma message("SDL_JoystickPathForIndex is not supported before SDL 2.24.0")
#pragma message("SDL_JoystickPath is not supported before SDL 2.24.0")
#pragma message("SDL_JoystickGetFirmwareVersion is not supported before SDL 2.24.0")
#pragma message("SDL_JoystickAttachVirtualEx is not supported before SDL 2.24.0")
#endif

typedef struct SDL_VirtualJoystickDesc
{
    Uint16 version;     // `SDL_VIRTUAL_JOYSTICK_DESC_VERSION`
    Uint16 type;        // `SDL_JoystickType`
    Uint16 naxes;       // the number of axes on this joystick
    Uint16 nbuttons;    // the number of buttons on this joystick
    Uint16 nhats;       // the number of hats on this joystick
    Uint16 vendor_id;   // the USB vendor ID of this joystick
    Uint16 product_id;  // the USB product ID of this joystick
    Uint16 padding;     // unused
    Uint32 button_mask; // A mask of which buttons are valid for this controller e.g. (1 << SDL_CONTROLLER_BUTTON_A)
    Uint32 axis_mask;   // A mask of which axes are valid for this controller e.g. (1 << SDL_CONTROLLER_AXIS_LEFTX)
    const char *name;   // the name of the joystick

    void *userdata;     // User data pointer passed to callbacks
    void (SDLCALL *Update)(void *userdata); // Called when the joystick state should be updated
    void (SDLCALL *SetPlayerIndex)(void *userdata, int player_index); // Called when the player index is set
    int (SDLCALL *Rumble)(void *userdata, Uint16 low_frequency_rumble, Uint16 high_frequency_rumble); // Implements SDL_JoystickRumble()
    int (SDLCALL *RumbleTriggers)(void *userdata, Uint16 left_rumble, Uint16 right_rumble); // Implements SDL_JoystickRumbleTriggers()
    int (SDLCALL *SetLED)(void *userdata, Uint8 red, Uint8 green, Uint8 blue); // Implements SDL_JoystickSetLED()
    int (SDLCALL *SendEffect)(void *userdata, const void *data, int size); // Implements SDL_JoystickSendEffect()

} SDL_VirtualJoystickDesc;

static const char *SDL_JoystickPathForIndex(int device_index)
{
	return NULL;
}

static const char *SDL_JoystickPath(SDL_Joystick *joystick)
{
	return NULL;
}

static Uint16 SDL_JoystickGetFirmwareVersion(SDL_Joystick *joystick)
{
	return 0;
}

static inline int SDL_JoystickAttachVirtualEx(const SDL_VirtualJoystickDesc *desc)
{
	return -1;
}

#endif
*/
import "C"
import "unsafe"

// Hat positions.
// (https://wiki.libsdl.org/SDL_JoystickGetHat)
type JoystickHat int

const (
	HAT_CENTERED  JoystickHat = C.SDL_HAT_CENTERED
	HAT_UP        JoystickHat = C.SDL_HAT_UP
	HAT_RIGHT     JoystickHat = C.SDL_HAT_RIGHT
	HAT_DOWN      JoystickHat = C.SDL_HAT_DOWN
	HAT_LEFT      JoystickHat = C.SDL_HAT_LEFT
	HAT_RIGHTUP   JoystickHat = C.SDL_HAT_RIGHTUP
	HAT_RIGHTDOWN JoystickHat = C.SDL_HAT_RIGHTDOWN
	HAT_LEFTUP    JoystickHat = C.SDL_HAT_LEFTUP
	HAT_LEFTDOWN  JoystickHat = C.SDL_HAT_LEFTDOWN
)

// Types of a joystick.
type JoystickType C.SDL_JoystickType

const (
	JOYSTICK_TYPE_UNKNOWN        JoystickType = C.SDL_JOYSTICK_TYPE_UNKNOWN
	JOYSTICK_TYPE_GAMECONTROLLER JoystickType = C.SDL_JOYSTICK_TYPE_GAMECONTROLLER
	JOYSTICK_TYPE_WHEEL          JoystickType = C.SDL_JOYSTICK_TYPE_WHEEL
	JOYSTICK_TYPE_ARCADE_STICK   JoystickType = C.SDL_JOYSTICK_TYPE_ARCADE_STICK
	JOYSTICK_TYPE_FLIGHT_STICK   JoystickType = C.SDL_JOYSTICK_TYPE_FLIGHT_STICK
	JOYSTICK_TYPE_DANCE_PAD      JoystickType = C.SDL_JOYSTICK_TYPE_DANCE_PAD
	JOYSTICK_TYPE_GUITAR         JoystickType = C.SDL_JOYSTICK_TYPE_GUITAR
	JOYSTICK_TYPE_DRUM_KIT       JoystickType = C.SDL_JOYSTICK_TYPE_DRUM_KIT
	JOYSTICK_TYPE_ARCADE_PAD     JoystickType = C.SDL_JOYSTICK_TYPE_ARCADE_PAD
	JOYSTICK_TYPE_THROTTLE       JoystickType = C.SDL_JOYSTICK_TYPE_THROTTLE
)

// An enumeration of battery levels of a joystick.
// (https://wiki.libsdl.org/SDL_JoystickPowerLevel)
type JoystickPowerLevel C.SDL_JoystickPowerLevel

const (
	JOYSTICK_POWER_UNKNOWN JoystickPowerLevel = C.SDL_JOYSTICK_POWER_UNKNOWN
	JOYSTICK_POWER_EMPTY   JoystickPowerLevel = C.SDL_JOYSTICK_POWER_EMPTY
	JOYSTICK_POWER_LOW     JoystickPowerLevel = C.SDL_JOYSTICK_POWER_LOW
	JOYSTICK_POWER_MEDIUM  JoystickPowerLevel = C.SDL_JOYSTICK_POWER_MEDIUM
	JOYSTICK_POWER_FULL    JoystickPowerLevel = C.SDL_JOYSTICK_POWER_FULL
	JOYSTICK_POWER_WIRED   JoystickPowerLevel = C.SDL_JOYSTICK_POWER_WIRED
	JOYSTICK_POWER_MAX     JoystickPowerLevel = C.SDL_JOYSTICK_POWER_MAX
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

func (joyid JoystickID) c() C.SDL_JoystickID {
	return C.SDL_JoystickID(joyid)
}

// GetInfo returns the device information encoded in a JoystickGUID structure.
// (https://wiki.libsdl.org/SDL_GetJoystickGUIDInfo)
func (guid JoystickGUID) GetInfo() (vendor, product, version, crc16 uint16) {
    _vendor := (*C.Uint16)(&vendor)
    _product := (*C.Uint16)(&product)
    _version := (*C.Uint16)(&version)
    _crc16 := (*C.Uint16)(&crc16)
    C.SDL_GetJoystickGUIDInfo(guid.c(), _vendor, _product, _version, _crc16)
    return
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

// JoystickPathForIndex returns the implementation dependent path of a joystick.
// (https://wiki.libsdl.org/SDL_JoystickPathForIndex)
func JoystickPathForIndex(index int) string {
	return (C.GoString)(C.SDL_JoystickPathForIndex(C.int(index)))
}

// JoystickGetDevicePlayerIndex returns the player index of a joystick, or -1 if it's not available
// (https://wiki.libsdl.org/SDL_JoystickGetDevicePlayerIndex)
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

// JoystickFromPlayerIndex returns the Joystick associated with a player index.
// (https://wiki.libsdl.org/SDL_JoystickFromPlayerIndex)
func JoystickFromPlayerIndex(playerIndex int) *Joystick {
	return (*Joystick)(C.SDL_JoystickFromPlayerIndex(C.int(playerIndex)))
}

// JoystickAttachVirtual attaches a new virtual joystick.
// (https://wiki.libsdl.org/SDL_JoystickAttachVirtual)
func JoystickAttachVirtual(typ JoystickType, naxes, nbuttons, nhats int) (deviceIndex int, err error) {
	deviceIndex = int(C.SDL_JoystickAttachVirtual(C.SDL_JoystickType(typ), C.int(naxes), C.int(nbuttons), C.int(nhats)))
	err = errorFromInt(deviceIndex)
	return
}

// JoystickAttachVirtualEx attaches a new virtual joystick with extended properties.
// (https://wiki.libsdl.org/SDL_JoystickAttachVirtualEx)
func JoystickAttachVirtualEx(desc *C.SDL_VirtualJoystickDesc) (deviceIndex int, err error) {
	deviceIndex = int(C.SDL_JoystickAttachVirtualEx(desc))
	err = errorFromInt(deviceIndex)
	return
}

// JoystickDetachVirtual detaches a virtual joystick.
// (https://wiki.libsdl.org/SDL_JoystickDetachVirtual)
func JoystickDetachVirtual(deviceIndex int) error {
	return errorFromInt(int(C.SDL_JoystickDetachVirtual(C.int(deviceIndex))))
}

// JoystickIsVirtual indicates whether or not a virtual-joystick is at a given device index.
// (https://wiki.libsdl.org/SDL_JoystickIsVirtual)
func JoystickIsVirtual(deviceIndex int) bool {
	return C.SDL_JoystickIsVirtual(C.int(deviceIndex)) == C.SDL_TRUE
}

// SetVirtualAxis sets axis value on an opened, virtual-joystick's controls.
// Please note that values set here will not be applied until the next
// call to SDL_JoystickUpdate, which can either be called directly,
// or can be called indirectly through various other SDL APIS,
// including, but not limited to the following: SDL_PollEvent,
// SDL_PumpEvents, SDL_WaitEventTimeout, SDL_WaitEvent..
// (https://wiki.libsdl.org/SDL_JoystickSetVirtualAxis)
func (joy *Joystick) SetVirtualAxis(axis int, value int16) error {
	return errorFromInt(int(C.SDL_JoystickSetVirtualAxis(joy.cptr(), C.int(axis), C.Sint16(value))))
}

// SetVirtualButton sets button value on an opened, virtual-joystick's controls.
// Please note that values set here will not be applied until the next
// call to SDL_JoystickUpdate, which can either be called directly,
// or can be called indirectly through various other SDL APIS,
// including, but not limited to the following: SDL_PollEvent,
// SDL_PumpEvents, SDL_WaitEventTimeout, SDL_WaitEvent..
// (https://wiki.libsdl.org/SDL_JoystickSetVirtualButton)
func (joy *Joystick) SetVirtualButton(button int, value uint8) error {
	return errorFromInt(int(C.SDL_JoystickSetVirtualButton(joy.cptr(), C.int(button), C.Uint8(value))))
}

// SetVirtualHat sets hat value on an opened, virtual-joystick's controls.
// Please note that values set here will not be applied until the next
// call to SDL_JoystickUpdate, which can either be called directly,
// or can be called indirectly through various other SDL APIS,
// including, but not limited to the following: SDL_PollEvent,
// SDL_PumpEvents, SDL_WaitEventTimeout, SDL_WaitEvent..
// (https://wiki.libsdl.org/SDL_JoystickSetVirtualHat)
func (joy *Joystick) SetVirtualHat(hat int, value uint8) error {
	return errorFromInt(int(C.SDL_JoystickSetVirtualHat(joy.cptr(), C.int(hat), C.Uint8(value))))
}

// LockJoysticks locks joysticks for multi-threaded access to the joystick API
// (https://wiki.libsdl.org/SDL_LockJoysticks)
func LockJoysticks() {
	C.SDL_LockJoysticks()
}

// UnlockJoysticks unlocks joysticks for multi-threaded access to the joystick API
// (https://wiki.libsdl.org/SDL_UnlockJoysticks)
func UnlockJoysticks() {
	C.SDL_UnlockJoysticks()
}

// Name returns the implementation dependent name of a joystick.
// (https://wiki.libsdl.org/SDL_JoystickName)
func (joy *Joystick) Name() string {
	return (C.GoString)(C.SDL_JoystickName(joy.cptr()))
}

// Path returns the implementation dependent path of a joystick.
// (https://wiki.libsdl.org/SDL_JoystickPath)
func (joy *Joystick) Path() string {
	return (C.GoString)(C.SDL_JoystickPath(joy.cptr()))
}

// PlayerIndex returns the player index of an opened joystick, or -1 if it's not available.
// (https://wiki.libsdl.org/SDL_JoystickGetPlayerIndex)
func (joy *Joystick) PlayerIndex() int {
	return int(C.SDL_JoystickGetPlayerIndex(joy.cptr()))
}

// SetPlayerIndex returns set the player index of an opened joystick.
// (https://wiki.libsdl.org/SDL_JoystickSetPlayerIndex)
func (joy *Joystick) SetPlayerIndex(playerIndex int) {
	C.SDL_JoystickSetPlayerIndex(joy.cptr(), C.int(playerIndex))
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

// FirmwareVersion returns the firmware version of an opened joystick, if available.
// (https://wiki.libsdl.org/SDL_JoystickGetFirmwareVersion)
func (joy *Joystick) FirmwareVersion() uint16 {
	return uint16(C.SDL_JoystickGetFirmwareVersion(joy.cptr()))
}

// Serial returns the serial number of an opened joystick, if available.
// (https://wiki.libsdl.org/SDL_JoystickGetSerial)
func (joy *Joystick) Serial() string {
	return C.GoString(C.SDL_JoystickGetSerial(joy.cptr()))
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
func (joy *Joystick) Hat(hat JoystickHat) byte {
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
// (https://wiki.libsdl.org/SDL_JoystickRumble)
func (joy *Joystick) Rumble(lowFrequencyRumble, highFrequencyRumble uint16, durationMS uint32) error {
	return errorFromInt(int(C.SDL_JoystickRumble(joy.cptr(), C.Uint16(lowFrequencyRumble), C.Uint16(highFrequencyRumble), C.Uint32(durationMS))))
}

// RumbleTriggers starts a rumble effect in the joystick's triggers.
// (https://wiki.libsdl.org/SDL_JoystickRumbleTriggers)
func (joy *Joystick) RumbleTriggers(leftRumble, rightRumble uint16, durationMS uint32) error {
	return errorFromInt(int(C.SDL_JoystickRumbleTriggers(joy.cptr(), C.Uint16(leftRumble), C.Uint16(rightRumble), C.Uint32(durationMS))))
}

// HasLED returns whether a joystick has an LED.
// (https://wiki.libsdl.org/SDL_JoystickHasLED)
func (joy *Joystick) HasLED() bool {
	return C.SDL_JoystickHasLED(joy.cptr()) == C.SDL_TRUE
}

// SetLED updates a joystick's LED color.
// (https://wiki.libsdl.org/SDL_JoystickSetLED)
func (joy *Joystick) SetLED(red, green, blue uint8) error {
	return errorFromInt(int(C.SDL_JoystickSetLED(joy.cptr(), C.Uint8(red), C.Uint8(green), C.Uint8(blue))))
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
