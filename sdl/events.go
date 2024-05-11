package sdl

/*
#include "sdl_wrapper.h"
#include "events.h"

// Used by Go functions like SDL_SensorEvent from SDL2 2.26.0 to enable backward compatibility when building with older SDL2
typedef struct SensorEvent {
    Uint32 type;
    Uint32 timestamp;
    Sint32 which;
    float data[6];
    Uint64 timestamp_us;
} SensorEvent;

typedef struct ControllerSensorEvent {
    Uint32 type;
    Uint32 timestamp;
    SDL_JoystickID which;
    Sint32 sensor;
    float data[3];
    Uint64 timestamp_us;
} ControllerSensorEvent;

typedef struct MouseWheelEvent
{
    Uint32 type;        // ::SDL_MOUSEWHEEL
    Uint32 timestamp;   // In milliseconds, populated using SDL_GetTicks()
    Uint32 windowID;    // The window with mouse focus, if any
    Uint32 which;       // The mouse instance id, or SDL_TOUCH_MOUSEID
    Sint32 x;           // The amount scrolled horizontally, positive to the right and negative to the left
    Sint32 y;           // The amount scrolled vertically, positive away from the user and negative toward the user
    Uint32 direction;   // Set to one of the SDL_MOUSEWHEEL_* defines. When FLIPPED the values in X and Y will be opposite. Multiply by -1 to change them back
    float preciseX;     // The amount scrolled horizontally, positive to the right and negative to the left, with float precision (added in 2.0.18)
    float preciseY;     // The amount scrolled vertically, positive away from the user and negative toward the user, with float precision (added in 2.0.18)
    Sint32 mouseX;       // X coordinate, relative to window (added in 2.26.0)
    Sint32 mouseY;       // Y coordinate, relative to window (added in 2.26.0)
} MouseWheelEvent;

#if !SDL_VERSION_ATLEAST(2,0,9)
#if defined(WARN_OUTDATED)
#pragma message("SDL_DISPLAYEVENT is not supported before SDL 2.0.9")
#pragma message("SDL_DisplayEvent is not supported before SDL 2.0.9")
#endif

#define SDL_DISPLAYEVENT (0x150)

typedef struct SDL_DisplayEvent
{
    Uint32 type;        // ::SDL_DISPLAYEVENT
    Uint32 timestamp;   // In milliseconds, populated using SDL_GetTicks()
    Uint32 display;     // The associated display index
    Uint8 event;        // ::SDL_DisplayEventID
    Uint8 padding1;
    Uint8 padding2;
    Uint8 padding3;
    Sint32 data1;       // event dependent data
} SDL_DisplayEvent;
#endif

#if !SDL_VERSION_ATLEAST(2,0,2)
#if defined(WARN_OUTDATED)
#pragma message("SDL_RENDER_TARGETS_RESET is not supported before SDL 2.0.2")
#endif

#define SDL_RENDER_TARGETS_RESET (0x2000)
#endif

#if !SDL_VERSION_ATLEAST(2,0,4)

#if defined(WARN_OUTDATED)
#pragma message("SDL_KEYMAPCHANGED is not supported before SDL 2.0.4")
#pragma message("SDL_AUDIODEVICEADDED is not supported before SDL 2.0.4")
#pragma message("SDL_AUDIODEVICEREMOVED is not supported before SDL 2.0.4")
#pragma message("SDL_RENDER_DEVICE_RESET is not supported before SDL 2.0.4")
#pragma message("SDL_AudioDeviceEvent is not supported before SDL 2.0.4")
#endif

#define SDL_KEYMAPCHANGED (0x304)
#define SDL_AUDIODEVICEADDED (0x1100)
#define SDL_AUDIODEVICEREMOVED (0x1101)
#define SDL_RENDER_DEVICE_RESET (0x2001)

typedef struct SDL_AudioDeviceEvent
{
    Uint32 type;
    Uint32 timestamp;
    Uint32 which;
    Uint8  iscapture;
    Uint8  padding1;
    Uint8  padding2;
    Uint8  padding3;
} SDL_AudioDeviceEvent;

#endif

#if !SDL_VERSION_ATLEAST(2,0,5)

#if defined(WARN_OUTDATED)
#pragma message("SDL_DROPTEXT is not supported before SDL 2.0.5")
#pragma message("SDL_DROPBEGIN is not supported before SDL 2.0.5")
#pragma message("SDL_DROPCOMPLETE is not supported before SDL 2.0.5")
#endif

#define SDL_DROPTEXT (0x1001)
#define SDL_DROPBEGIN (0x1002)
#define SDL_DROPCOMPLETE (0x1003)

#endif

#if !SDL_VERSION_ATLEAST(2,0,9)
#define SDL_SENSORUPDATE (0x1200)
#endif

#if !SDL_VERSION_ATLEAST(2,0,22)
#if defined(WARN_OUTDATED)
#pragma message("SDL_TEXTEDITING_EXT is not supported before SDL 2.0.22")
#endif

#if !SDL_VERSION_ATLEAST(2,0,14)
#define SDL_CONTROLLERTOUCHPADDOWN (0x656)
#define SDL_CONTROLLERTOUCHPADMOTION (0x657)
#define SDL_CONTROLLERTOUCHPADUP (0x658)
#define SDL_CONTROLLERSENSORUPDATE (0x659)
#endif

#define SDL_TEXTEDITING_EXT (0x305)
#endif

// NOTE: To prevent build from failing when using older SDL2, we create a
// structure definiton that directly maps to the SDL2 struct definition if
// using the latest SDL2. Otherwise, we copy the latest definition and paste
// it here.
#if SDL_VERSION_ATLEAST(2,0,22)
typedef SDL_MouseButtonEvent MouseButtonEvent;
typedef SDL_TouchFingerEvent TouchFingerEvent;
typedef SDL_DropEvent DropEvent;
#else
typedef struct MouseButtonEvent
{
    Uint32 type;        // ::SDL_MOUSEBUTTONDOWN or ::SDL_MOUSEBUTTONUP
    Uint32 timestamp;
    Uint32 windowID;    // The window with mouse focus, if any
    Uint32 which;       // The mouse instance id, or SDL_TOUCH_MOUSEID
    Uint8 button;       // The mouse button index
    Uint8 state;        // ::SDL_PRESSED or ::SDL_RELEASED
    Uint8 clicks;       // 1 for single-click, 2 for double-click, etc.
    Uint8 padding1;
    Sint32 x;           // X coordinate, relative to window
    Sint32 y;           // Y coordinate, relative to window
} MouseButtonEvent;

typedef struct TouchFingerEvent
{
    Uint32 type;           // ::SDL_FINGERMOTION or ::SDL_FINGERDOWN or ::SDL_FINGERUP
    Uint32 timestamp;      // In milliseconds, populated using SDL_GetTicks()
    SDL_TouchID touchId;   // The touch device id
    SDL_FingerID fingerId; //
    float x;               // Normalized in the range 0...1
    float y;               // Normalized in the range 0...1
    float dx;              // Normalized in the range -1...1
    float dy;              // Normalized in the range -1...1
    float pressure;        // Normalized in the range 0...1
    Uint32 windowID;       // The window underneath the finger, if any
} TouchFingerEvent;

typedef struct DropEvent
{
    Uint32 type;        // ::SDL_DROPBEGIN or ::SDL_DROPFILE or ::SDL_DROPTEXT or ::SDL_DROPCOMPLETE
    Uint32 timestamp;   // In milliseconds, populated using SDL_GetTicks()
    char *file;         // The file name, which should be freed with SDL_free(), is NULL on begin/complete
    Uint32 windowID;    // The window that was dropped on, if any
} DropEvent;

#endif

#if !SDL_VERSION_ATLEAST(2,24,0)
#if defined(WARN_OUTDATED)
#pragma message("SDL_JOYBATTERYUPDATED is not supported before SDL 2.24.0")
#endif

#define SDL_JOYBATTERYUPDATED (1543)

#if !SDL_VERSION_ATLEAST(2,0,4)
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
#endif

typedef struct SDL_JoyBatteryEvent
{
    Uint32 type;
    Uint32 timestamp;
    SDL_JoystickID which;
    SDL_JoystickPowerLevel level;
} SDL_JoyBatteryEvent;

#endif

#if !SDL_VERSION_ATLEAST(2,30,0)
#if defined(WARN_OUTDATED)
#pragma message("SDL_CONTROLLERSTEAMHANDLEUPDATED is not supported before SDL 2.30.0")
#endif
#endif
*/
import "C"
import (
	"reflect"
	"sync"
	"unsafe"
)

var (
	eventFilterCache          EventFilter
	eventWatches              = make(map[EventWatchHandle]*eventFilterCallbackContext)
	lastEventWatchHandleMutex sync.Mutex
	lastEventWatchHandle      EventWatchHandle
	cevent                    C.SDL_Event
)

// Enumeration of the types of events that can be delivered.
// (https://wiki.libsdl.org/SDL_EventType)
type EventType uint32

const (
	FIRSTEVENT EventType = C.SDL_FIRSTEVENT // do not remove (unused)

	// Application events
	QUIT EventType = C.SDL_QUIT // user-requested quit

	// Android, iOS and WinRT events
	APP_TERMINATING         EventType = C.SDL_APP_TERMINATING         // OS is terminating the application
	APP_LOWMEMORY           EventType = C.SDL_APP_LOWMEMORY           // OS is low on memory; free some
	APP_WILLENTERBACKGROUND EventType = C.SDL_APP_WILLENTERBACKGROUND // application is entering background
	APP_DIDENTERBACKGROUND  EventType = C.SDL_APP_DIDENTERBACKGROUND  //application entered background
	APP_WILLENTERFOREGROUND EventType = C.SDL_APP_WILLENTERFOREGROUND // application is entering foreground
	APP_DIDENTERFOREGROUND  EventType = C.SDL_APP_DIDENTERFOREGROUND  // application entered foreground

	// Display events
	DISPLAYEVENT EventType = C.SDL_DISPLAYEVENT // Display state change

	// Window events
	WINDOWEVENT EventType = C.SDL_WINDOWEVENT // window state change
	SYSWMEVENT  EventType = C.SDL_SYSWMEVENT  // system specific event

	// Keyboard events
	KEYDOWN         EventType = C.SDL_KEYDOWN         // key pressed
	KEYUP           EventType = C.SDL_KEYUP           // key released
	TEXTEDITING     EventType = C.SDL_TEXTEDITING     // keyboard text editing (composition)
	TEXTINPUT       EventType = C.SDL_TEXTINPUT       // keyboard text input
	TEXTEDITING_EXT EventType = C.SDL_TEXTEDITING_EXT // keyboard text editing (composition)
	KEYMAPCHANGED   EventType = C.SDL_KEYMAPCHANGED   // keymap changed due to a system event such as an input language or keyboard layout change (>= SDL 2.0.4)

	// Mouse events
	MOUSEMOTION     EventType = C.SDL_MOUSEMOTION     // mouse moved
	MOUSEBUTTONDOWN EventType = C.SDL_MOUSEBUTTONDOWN // mouse button pressed
	MOUSEBUTTONUP   EventType = C.SDL_MOUSEBUTTONUP   // mouse button released
	MOUSEWHEEL      EventType = C.SDL_MOUSEWHEEL      // mouse wheel motion

	// Joystick events
	JOYAXISMOTION     EventType = C.SDL_JOYAXISMOTION     // joystick axis motion
	JOYBALLMOTION     EventType = C.SDL_JOYBALLMOTION     // joystick trackball motion
	JOYHATMOTION      EventType = C.SDL_JOYHATMOTION      // joystick hat position change
	JOYBUTTONDOWN     EventType = C.SDL_JOYBUTTONDOWN     // joystick button pressed
	JOYBUTTONUP       EventType = C.SDL_JOYBUTTONUP       // joystick button released
	JOYDEVICEADDED    EventType = C.SDL_JOYDEVICEADDED    // joystick connected
	JOYDEVICEREMOVED  EventType = C.SDL_JOYDEVICEREMOVED  // joystick disconnected
	JOYBATTERYUPDATED EventType = C.SDL_JOYBATTERYUPDATED // joystick battery level change

	// Game controller events
	CONTROLLERAXISMOTION     EventType = C.SDL_CONTROLLERAXISMOTION     // controller axis motion
	CONTROLLERBUTTONDOWN     EventType = C.SDL_CONTROLLERBUTTONDOWN     // controller button pressed
	CONTROLLERBUTTONUP       EventType = C.SDL_CONTROLLERBUTTONUP       // controller button released
	CONTROLLERDEVICEADDED    EventType = C.SDL_CONTROLLERDEVICEADDED    // controller connected
	CONTROLLERDEVICEREMOVED  EventType = C.SDL_CONTROLLERDEVICEREMOVED  // controller disconnected
	CONTROLLERDEVICEREMAPPED EventType = C.SDL_CONTROLLERDEVICEREMAPPED // controller mapping updated
    CONTROLLERTOUCHPADDOWN   EventType = C.SDL_CONTROLLERTOUCHPADDOWN   // Game controller touchpad was touched
    CONTROLLERTOUCHPADMOTION EventType = C.SDL_CONTROLLERTOUCHPADMOTION // Game controller touchpad finger was moved
    CONTROLLERTOUCHPADUP     EventType = C.SDL_CONTROLLERTOUCHPADUP     // Game controller touchpad finger was lifted
    CONTROLLERSENSORUPDATE   EventType = C.SDL_CONTROLLERSENSORUPDATE   // Game controller sensor was updated

	// Touch events
	FINGERDOWN   EventType = C.SDL_FINGERDOWN   // user has touched input device
	FINGERUP     EventType = C.SDL_FINGERUP     // user stopped touching input device
	FINGERMOTION EventType = C.SDL_FINGERMOTION // user is dragging finger on input device

	// Gesture events
	DOLLARGESTURE EventType = C.SDL_DOLLARGESTURE
	DOLLARRECORD  EventType = C.SDL_DOLLARRECORD
	MULTIGESTURE  EventType = C.SDL_MULTIGESTURE

	// Clipboard events
	CLIPBOARDUPDATE EventType = C.SDL_CLIPBOARDUPDATE // the clipboard changed

	// Drag and drop events
	DROPFILE     EventType = C.SDL_DROPFILE     // the system requests a file open
	DROPTEXT     EventType = C.SDL_DROPTEXT     // text/plain drag-and-drop event
	DROPBEGIN    EventType = C.SDL_DROPBEGIN    // a new set of drops is beginning (NULL filename)
	DROPCOMPLETE EventType = C.SDL_DROPCOMPLETE // current set of drops is now complete (NULL filename)

	// Audio hotplug events
	AUDIODEVICEADDED   EventType = C.SDL_AUDIODEVICEADDED   // a new audio device is available (>= SDL 2.0.4)
	AUDIODEVICEREMOVED EventType = C.SDL_AUDIODEVICEREMOVED // an audio device has been removed (>= SDL 2.0.4)

	// Sensor events
	SENSORUPDATE EventType = C.SDL_SENSORUPDATE // a sensor was updated

	// Render events
	RENDER_TARGETS_RESET EventType = C.SDL_RENDER_TARGETS_RESET // the render targets have been reset and their contents need to be updated (>= SDL 2.0.2)
	RENDER_DEVICE_RESET  EventType = C.SDL_RENDER_DEVICE_RESET  // the device has been reset and all textures need to be recreated (>= SDL 2.0.4)

	// These are for your use, and should be allocated with RegisterEvents()
	USEREVENT EventType = C.SDL_USEREVENT // a user-specified event
	LASTEVENT EventType = C.SDL_LASTEVENT // (only for bounding internal arrays)
)

// Actions for PeepEvents().
// (https://wiki.libsdl.org/SDL_PeepEvents)
type EventAction C.SDL_eventaction

const (
	ADDEVENT  EventAction = C.SDL_ADDEVENT  // up to numevents events will be added to the back of the event queue
	PEEKEVENT EventAction = C.SDL_PEEKEVENT // up to numevents events at the front of the event queue, within the specified minimum and maximum type, will be returned and will not be removed from the queue
	GETEVENT  EventAction = C.SDL_GETEVENT  // up to numevents events at the front of the event queue, within the specified minimum and maximum type, will be returned and will be removed from the queue
)

// Toggles for different event state functions.
// (https://wiki.libsdl.org/SDL_EventState)
type EventStateConstant int

const (
	QUERY   EventStateConstant = C.SDL_QUERY
	IGNORE  EventStateConstant = C.SDL_IGNORE
	DISABLE EventStateConstant = C.SDL_DISABLE
	ENABLE  EventStateConstant = C.SDL_ENABLE
)

// Event is a union of all event structures used in SDL.
// (https://wiki.libsdl.org/SDL_Event)
type Event interface {
	GetType() EventType   // GetType returns the event type
	GetTimestamp() uint32 // GetTimestamp returns the timestamp of the event
}

// CEvent is a union of all event structures used in SDL.
// (https://wiki.libsdl.org/SDL_Event)
type CEvent struct {
	Type EventType
	_    [52]byte // padding
}

// CommonEvent contains common event data.
// (https://wiki.libsdl.org/SDL_Event)
type CommonEvent struct {
	Type      EventType // the event type
	Timestamp uint32    // timestamp of the event
}

type cCommonEvent C.SDL_CommonEvent

// GetType returns the event type.
func (e CommonEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e CommonEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// DisplayEvent contains common event data.
// (https://wiki.libsdl.org/SDL_Event)
type DisplayEvent struct {
	Type      EventType // the event type
	Timestamp uint32    // timestamp of the event
	Display   uint32    // the associated display index
	Event     uint8     // event subtype for display events (https://wiki.libsdl.org/SDL_DisplayEventID)
	Data1     int32     // event dependent data
}

type cDisplayEvent C.SDL_DisplayEvent

// GetType returns the event type.
func (e DisplayEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e DisplayEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// WindowEvent contains window state change event data.
// (https://wiki.libsdl.org/SDL_WindowEvent)
type WindowEvent struct {
	Type      EventType     // WINDOWEVENT
	Timestamp uint32        // timestamp of the event
	WindowID  uint32        // the associated window
	Event     WindowEventID // (https://wiki.libsdl.org/SDL_WindowEventID)
	Data1     int32         // event dependent data
	Data2     int32         // event dependent data
}
type cWindowEvent C.SDL_WindowEvent

// GetType returns the event type.
func (e WindowEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e WindowEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// KeyboardEvent contains keyboard key down event information.
// (https://wiki.libsdl.org/SDL_KeyboardEvent)
type KeyboardEvent struct {
	Type      EventType   // KEYDOWN, KEYUP
	Timestamp uint32      // timestamp of the event
	WindowID  uint32      // the window with keyboard focus, if any
	State     ButtonState // PRESSED, RELEASED
	Repeat    uint8       // non-zero if this is a key repeat
	Keysym    Keysym      // Keysym representing the key that was pressed or released
}
type cKeyboardEvent C.SDL_KeyboardEvent

// GetType returns the event type.
func (e KeyboardEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e KeyboardEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// TextEditingEvent contains keyboard text editing event information.
// (https://wiki.libsdl.org/SDL_TextEditingEvent)
type TextEditingEvent struct {
	Type      EventType // TEXTEDITING
	Timestamp uint32    // timestamp of the event
	WindowID  uint32    // the window with keyboard focus, if any
	Text      string    // the null-terminated editing text in UTF-8 encoding
	Start     int32     // the location to begin editing from
	Length    int32     // the number of characters to edit from the start point
}
type cTextEditingEvent C.SDL_TextEditingEvent

// GetType returns the event type.
func (e TextEditingEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e TextEditingEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// GetText returns the text as string
func (e TextEditingEvent) GetText() string {
	return e.Text
}

// TextInputEvent contains keyboard text input event information.
// (https://wiki.libsdl.org/SDL_TextInputEvent)
type TextInputEvent struct {
	Type      EventType // TEXTINPUT
	Timestamp uint32    // timestamp of the event
	WindowID  uint32    // the window with keyboard focus, if any
	Text      string    // the null-terminated input text in UTF-8 encoding
}
type cTextInputEvent C.SDL_TextInputEvent

// GetType returns the event type.
func (e TextInputEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e TextInputEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// GetText returns the text as string
func (e TextInputEvent) GetText() string {
	return e.Text
}

// MouseMotionEvent contains mouse motion event information.
// (https://wiki.libsdl.org/SDL_MouseMotionEvent)
type MouseMotionEvent struct {
	Type      EventType  // MOUSEMOTION
	Timestamp uint32     // timestamp of the event
	WindowID  uint32     // the window with mouse focus, if any
	Which     uint32     // the mouse instance id, or TOUCH_MOUSEID
	State     ButtonMask // masks for BUTTON_LEFT, BUTTON_MIDDLE, BUTTON_RIGHT, BUTTON_X1, BUTTON_X2
	X         int32      // X coordinate, relative to window
	Y         int32      // Y coordinate, relative to window
	XRel      int32      // relative motion in the X direction
	YRel      int32      // relative motion in the Y direction
}
type cMouseMotionEvent C.SDL_MouseMotionEvent

// GetType returns the event type.
func (e MouseMotionEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e MouseMotionEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// MouseButtonEvent contains mouse button event information.
// (https://wiki.libsdl.org/SDL_MouseButtonEvent)
type MouseButtonEvent struct {
	Type      EventType   // MOUSEBUTTONDOWN, MOUSEBUTTONUP
	Timestamp uint32      // timestamp of the event
	WindowID  uint32      // the window with mouse focus, if any
	Which     uint32      // the mouse instance id, or TOUCH_MOUSEID
	Button    Button      // BUTTON_LEFT, BUTTON_MIDDLE, BUTTON_RIGHT, BUTTON_X1, BUTTON_X2
	State     ButtonState // PRESSED, RELEASED
	Clicks    uint8       // 1 for single-click, 2 for double-click, etc. (>= SDL 2.0.2)
	X         int32       // X coordinate, relative to window
	Y         int32       // Y coordinate, relative to window
}
type cMouseButtonEvent C.MouseButtonEvent

// GetType returns the event type.
func (e MouseButtonEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e MouseButtonEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// MouseWheelEvent contains mouse wheel event information.
// (https://wiki.libsdl.org/SDL_MouseWheelEvent)
type MouseWheelEvent struct {
	Type      EventType // MOUSEWHEEL
	Timestamp uint32    // timestamp of the event
	WindowID  uint32    // the window with mouse focus, if any
	Which     uint32    // the mouse instance id, or TOUCH_MOUSEID
	X         int32     // the amount scrolled horizontally, positive to the right and negative to the left
	Y         int32     // the amount scrolled vertically, positive away from the user and negative toward the user
	Direction uint32    // MOUSEWHEEL_NORMAL, MOUSEWHEEL_FLIPPED (>= SDL 2.0.4)
	PreciseX  float32   // The amount scrolled horizontally, positive to the right and negative to the left, with float precision (added in 2.0.18)
	PreciseY  float32   // The amount scrolled vertically, positive away from the user and negative toward the user, with float precision (added in 2.0.18)
	MouseX    int32     // X coordinate, relative to window (added in 2.26.0)
	MouseY    int32     // Y coordinate, relative to window (added in 2.26.0)
}
type cMouseWheelEvent C.MouseWheelEvent

// GetType returns the event type.
func (e MouseWheelEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e MouseWheelEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// JoyAxisEvent contains joystick axis motion event information.
// (https://wiki.libsdl.org/SDL_JoyAxisEvent)
type JoyAxisEvent struct {
	Type      EventType  // JOYAXISMOTION
	Timestamp uint32     // timestamp of the event
	Which     JoystickID // the instance id of the joystick that reported the event
	Axis      uint8      // the index of the axis that changed
	Value     int16      // the current position of the axis (range: -32768 to 32767)
}
type cJoyAxisEvent C.SDL_JoyAxisEvent

// GetType returns the event type.
func (e JoyAxisEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e JoyAxisEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// JoyBallEvent contains joystick trackball motion event information.
// (https://wiki.libsdl.org/SDL_JoyBallEvent)
type JoyBallEvent struct {
	Type      EventType  // JOYBALLMOTION
	Timestamp uint32     // timestamp of the event
	Which     JoystickID // the instance id of the joystick that reported the event
	Ball      uint8      // the index of the trackball that changed
	XRel      int16      // the relative motion in the X direction
	YRel      int16      // the relative motion in the Y direction
}
type cJoyBallEvent C.SDL_JoyBallEvent

// GetType returns the event type.
func (e JoyBallEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e JoyBallEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// JoyHatEvent contains joystick hat position change event information.
// (https://wiki.libsdl.org/SDL_JoyHatEvent)
type JoyHatEvent struct {
	Type      EventType   // JOYHATMOTION
	Timestamp uint32      // timestamp of the event
	Which     JoystickID  // the instance id of the joystick that reported the event
	Hat       uint8       // the index of the hat that changed
	Value     JoystickHat // HAT_LEFTUP, HAT_UP, HAT_RIGHTUP, HAT_LEFT, HAT_CENTERED, HAT_RIGHT, HAT_LEFTDOWN, HAT_DOWN, HAT_RIGHTDOWN
}
type cJoyHatEvent C.SDL_JoyHatEvent

// GetType returns the event type.
func (e JoyHatEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e JoyHatEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// JoyButtonEvent contains joystick button event information.
// (https://wiki.libsdl.org/SDL_JoyButtonEvent)
type JoyButtonEvent struct {
	Type      EventType   // JOYBUTTONDOWN, JOYBUTTONUP
	Timestamp uint32      // timestamp of the event
	Which     JoystickID  // the instance id of the joystick that reported the event
	Button    uint8       // the index of the button that changed
	State     ButtonState // PRESSED, RELEASED
}
type cJoyButtonEvent C.SDL_JoyButtonEvent

// GetType returns the event type.
func (e JoyButtonEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e JoyButtonEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// JoyDeviceAddedEvent contains joystick device event information.
// (https://wiki.libsdl.org/SDL_JoyDeviceEvent)
type JoyDeviceAddedEvent struct {
	Type      EventType  // JOYDEVICEADDED
	Timestamp uint32     // the timestamp of the event
	Which     JoystickID // the joystick device index
}

type cJoyDeviceEvent C.SDL_JoyDeviceEvent

// GetType returns the event type.
func (e JoyDeviceAddedEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e JoyDeviceAddedEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// JoyDeviceRemovedEvent contains joystick device event information.
// (https://wiki.libsdl.org/SDL_JoyDeviceEvent)
type JoyDeviceRemovedEvent struct {
	Type      EventType  // JOYDEVICEREMOVED
	Timestamp uint32     // the timestamp of the event
	Which     JoystickID // the instance id
}

// GetType returns the event type.
func (e JoyDeviceRemovedEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e JoyDeviceRemovedEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// JoyBatteryEvent contains joystick button event information.
// (https://wiki.libsdl.org/SDL_JoyBatteryEvent)
type JoyBatteryEvent struct {
	Type      EventType          // JOYBATTERYUPDATED
	Timestamp uint32             // timestamp of the event
	Which     JoystickID         // the instance id of the joystick that reported the event
	Level     JoystickPowerLevel // the joystick battery level
}
type cJoyBatteryEvent C.SDL_JoyBatteryEvent

// GetType returns the event type.
func (e JoyBatteryEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e JoyBatteryEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// ControllerAxisEvent contains game controller axis motion event information.
// (https://wiki.libsdl.org/SDL_ControllerAxisEvent)
type ControllerAxisEvent struct {
	Type      EventType          // CONTROLLERAXISMOTION
	Timestamp uint32             // the timestamp of the event
	Which     JoystickID         // the joystick instance id
	Axis      GameControllerAxis // the controller axis (https://wiki.libsdl.org/SDL_GameControllerAxis)
	Value     int16              // the axis value (range: -32768 to 32767)
}
type cControllerAxisEvent C.SDL_ControllerAxisEvent

// GetType returns the event type.
func (e ControllerAxisEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e ControllerAxisEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// ControllerButtonEvent contains game controller button event information.
// (https://wiki.libsdl.org/SDL_ControllerButtonEvent)
type ControllerButtonEvent struct {
	Type      EventType            // CONTROLLERBUTTONDOWN, CONTROLLERBUTTONUP
	Timestamp uint32               // the timestamp of the event
	Which     JoystickID           // the joystick instance id
	Button    GameControllerButton // the controller button (https://wiki.libsdl.org/SDL_GameControllerButton)
	State     ButtonState          // PRESSED, RELEASED
}
type cControllerButtonEvent C.SDL_ControllerButtonEvent

// GetType returns the event type.
func (e ControllerButtonEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e ControllerButtonEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// ControllerDeviceEvent contains controller device event information.
// (https://wiki.libsdl.org/SDL_ControllerDeviceEvent)
type ControllerDeviceEvent struct {
	Type      EventType  // CONTROLLERDEVICEADDED, CONTROLLERDEVICEREMOVED, SDL_CONTROLLERDEVICEREMAPPED
	Timestamp uint32     // the timestamp of the event
	Which     JoystickID // the joystick device index for the CONTROLLERDEVICEADDED event or instance id for the CONTROLLERDEVICEREMOVED or CONTROLLERDEVICEREMAPPED event
}
type cControllerDeviceEvent C.SDL_ControllerDeviceEvent

// GetType returns the event type.
func (e ControllerDeviceEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e ControllerDeviceEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// ControllerSensorEvent contains data from sensors such as accelerometer and gyroscope
// (https://wiki.libsdl.org/SDL_ControllerSensorEvent)
type ControllerSensorEvent struct {
	Type        EventType  // SDL_CONTROLLERSENSORUPDATE
	Timestamp   uint32     // In milliseconds, populated using SDL_GetTicks()
	Which       JoystickID // The joystick instance id
    Sensor      int32      // The type of the sensor, one of the values of SensorType
	Data        [3]float32 // Up to 3 values from the sensor - additional values can be queried using SDL_SensorGetData()
    TimestampUs uint64     // The timestamp of the sensor reading in microseconds, if the hardware provides this information.
}
type cControllerSensorEvent C.ControllerSensorEvent

// GetType returns the event type.
func (e ControllerSensorEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e ControllerSensorEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// AudioDeviceEvent contains audio device event information.
// (https://wiki.libsdl.org/SDL_AudioDeviceEvent)
type AudioDeviceEvent struct {
	Type      EventType // AUDIODEVICEADDED, AUDIODEVICEREMOVED
	Timestamp uint32    // the timestamp of the event
	Which     uint32    // the audio device index for the AUDIODEVICEADDED event (valid until next GetNumAudioDevices() call), AudioDeviceID for the AUDIODEVICEREMOVED event
	IsCapture uint8     // zero if an audio output device, non-zero if an audio capture device
}
type cAudioDeviceEvent C.SDL_AudioDeviceEvent

// GetType returns the event type.
func (e AudioDeviceEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e AudioDeviceEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// TouchFingerEvent contains finger touch event information.
// (https://wiki.libsdl.org/SDL_TouchFingerEvent)
type TouchFingerEvent struct {
	Type      EventType // FINGERMOTION, FINGERDOWN, FINGERUP
	Timestamp uint32    // timestamp of the event
	TouchID   TouchID   // the touch device id
	FingerID  FingerID  // the finger id
	X         float32   // the x-axis location of the touch event, normalized (0...1)
	Y         float32   // the y-axis location of the touch event, normalized (0...1)
	DX        float32   // the distance moved in the x-axis, normalized (-1...1)
	DY        float32   // the distance moved in the y-axis, normalized (-1...1)
	Pressure  float32   // the quantity of pressure applied, normalized (0...1)
	WindowID  uint32    // the window underneath the finger, if any
}
type cTouchFingerEvent C.TouchFingerEvent

// GetType returns the event type.
func (e TouchFingerEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e TouchFingerEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// MultiGestureEvent contains multiple finger gesture event information.
// (https://wiki.libsdl.org/SDL_MultiGestureEvent)
type MultiGestureEvent struct {
	Type       EventType // MULTIGESTURE
	Timestamp  uint32    // timestamp of the event
	TouchID    TouchID   // the touch device id
	DTheta     float32   // the amount that the fingers rotated during this motion
	DDist      float32   // the amount that the fingers pinched during this motion
	X          float32   // the normalized center of gesture
	Y          float32   // the normalized center of gesture
	NumFingers uint16    // the number of fingers used in the gesture
}
type cMultiGestureEvent C.SDL_MultiGestureEvent

// GetType returns the event type.
func (e MultiGestureEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e MultiGestureEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// DollarGestureEvent contains complex gesture event information.
// (https://wiki.libsdl.org/SDL_DollarGestureEvent)
type DollarGestureEvent struct {
	Type       EventType // DOLLARGESTURE, DOLLARRECORD
	Timestamp  uint32    // timestamp of the event
	TouchID    TouchID   // the touch device id
	GestureID  GestureID // the unique id of the closest gesture to the performed stroke
	NumFingers uint32    // the number of fingers used to draw the stroke
	Error      float32   // the difference between the gesture template and the actual performed gesture (lower error is a better match)
	X          float32   // the normalized center of gesture
	Y          float32   // the normalized center of gesture
}
type cDollarGestureEvent C.SDL_DollarGestureEvent

// GetType returns the event type.
func (e DollarGestureEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e DollarGestureEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// DropEvent contains an event used to request a file open by the system.
// (https://wiki.libsdl.org/SDL_DropEvent)
type DropEvent struct {
	Type      EventType // DROPFILE, DROPTEXT, DROPBEGIN, DROPCOMPLETE
	Timestamp uint32    // timestamp of the event
	File      string    // the file name
	WindowID  uint32    // the window that was dropped on, if any
}
type cDropEvent C.DropEvent

// GetType returns the event type.
func (e DropEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e DropEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// SensorEvent contains data from sensors such as accelerometer and gyroscope
// (https://wiki.libsdl.org/SDL_SensorEvent)
type SensorEvent struct {
	Type        EventType  // SDL_SENSORUPDATE
	Timestamp   uint32     // In milliseconds, populated using SDL_GetTicks()
	Which       int32      // The instance ID of the sensor
	Data        [6]float32 // Up to 6 values from the sensor - additional values can be queried using SDL_SensorGetData()
    TimestampUs uint64     // The timestamp of the sensor reading in microseconds, if the hardware provides this information.
}
type cSensorEvent C.SensorEvent

// GetType returns the event type.
func (e SensorEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e SensorEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// RenderEvent contains render event information.
// (https://wiki.libsdl.org/SDL_EventType)
type RenderEvent struct {
	Type      EventType // the event type
	Timestamp uint32    // timestamp of the event
}

// GetType returns the event type.
func (e RenderEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e RenderEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// QuitEvent contains the "quit requested" event.
// (https://wiki.libsdl.org/SDL_QuitEvent)
type QuitEvent struct {
	Type      EventType // QUIT
	Timestamp uint32    // timestamp of the event
}

// GetType returns the event type.
func (e QuitEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e QuitEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// OSEvent contains OS specific event information.
type OSEvent struct {
	Type      EventType // the event type
	Timestamp uint32    // timestamp of the event
}

// GetType returns the event type.
func (e OSEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e OSEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// ClipboardEvent contains clipboard event information.
// (https://wiki.libsdl.org/SDL_EventType)
type ClipboardEvent struct {
	Type      EventType // CLIPBOARDUPDATE
	Timestamp uint32    // timestamp of the event
}

// GetType returns the event type.
func (e ClipboardEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e ClipboardEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// UserEvent contains an application-defined event type.
// (https://wiki.libsdl.org/SDL_UserEvent)
type UserEvent struct {
	Type      EventType      // value obtained from RegisterEvents()
	Timestamp uint32         // timestamp of the event
	WindowID  uint32         // the associated window, if any
	Code      int32          // user defined event code
	Data1     unsafe.Pointer // user defined data pointer
	Data2     unsafe.Pointer // user defined data pointer
}
type cUserEvent C.SDL_UserEvent

// GetType returns the event type.
func (e UserEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e UserEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// SysWMEvent contains a video driver dependent system event.
// (https://wiki.libsdl.org/SDL_SysWMEvent)
type SysWMEvent struct {
	Type      EventType // SYSWMEVENT
	Timestamp uint32    // timestamp of the event
	Msg       *SysWMmsg // driver dependent data, defined in SDL_syswm.h
}
type cSysWMEvent C.SDL_SysWMEvent

// GetType returns the event type.
func (e SysWMEvent) GetType() EventType {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e SysWMEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// EventFilter is the function to call when an event happens.
// (https://wiki.libsdl.org/SDL_SetEventFilter)
type EventFilter interface {
	FilterEvent(e Event, userdata interface{}) bool
}

type eventFilterFunc func(Event, interface{}) bool

type eventFilterCallbackContext struct {
	filter   EventFilter
	handle   EventWatchHandle
	userdata interface{}
}

// EventWatchHandle is an event watch callback added with AddEventWatch().
type EventWatchHandle uintptr

func (action EventAction) c() C.SDL_eventaction {
	return C.SDL_eventaction(action)
}

// PumpEvents pumps the event loop, gathering events from the input devices.
// (https://wiki.libsdl.org/SDL_PumpEvents)
func PumpEvents() {
	C.SDL_PumpEvents()
}

// PeepEvents checks the event queue for messages and optionally return them.
// (https://wiki.libsdl.org/SDL_PeepEvents)
func PeepEvents(events []Event, action EventAction, minType, maxType EventType) (storedEvents int, err error) {
	if events == nil {
		return 0, nil
	}

	var _events []CEvent = make([]CEvent, len(events))

	if action == ADDEVENT { // the contents of _events matter if they are to be added
		for i := 0; i < len(events); i++ {
			_events[i] = *cEvent(events[i])
		}
	}

	_pevents := (*C.SDL_Event)(unsafe.Pointer(&_events[0]))
	storedEvents = int(C.SDL_PeepEvents(_pevents, C.int(len(events)), action.c(), C.Uint32(minType), C.Uint32(maxType)))

	if action != ADDEVENT { // put events into slice, events unchanged if action = ADDEVENT
		for i := 0; i < storedEvents; i++ {
			events[i] = goEvent(&_events[i])
		}
	}

	if storedEvents < 0 {
		err = GetError()
	}

	return
}

// HasEvent checks for the existence of certain event types in the event queue.
// (https://wiki.libsdl.org/SDL_HasEvent)
func HasEvent(type_ EventType) bool {
	return C.SDL_HasEvent(C.Uint32(type_)) != 0
}

// HasEvents checks for the existence of a range of event types in the event queue.
// (https://wiki.libsdl.org/SDL_HasEvents)
func HasEvents(minType, maxType EventType) bool {
	return C.SDL_HasEvents(C.Uint32(minType), C.Uint32(maxType)) != 0
}

// FlushEvent clears events from the event queue.
// (https://wiki.libsdl.org/SDL_FlushEvent)
func FlushEvent(type_ uint32) {
	C.SDL_FlushEvent(C.Uint32(type_))
}

// FlushEvents clears events from the event queue.
// (https://wiki.libsdl.org/SDL_FlushEvents)
func FlushEvents(minType, maxType EventType) {
	C.SDL_FlushEvents(C.Uint32(minType), C.Uint32(maxType))
}

// PollEvent polls for currently pending events.
// (https://wiki.libsdl.org/SDL_PollEvent)
func PollEvent() Event {
	ret := C.PollEvent()
	if ret == 0 {
		return nil
	}
	return goEvent((*CEvent)(unsafe.Pointer(&C.event)))
}

func goEvent(cevent *CEvent) Event {
	switch cevent.Type {
	case DISPLAYEVENT:
		e := (*cDisplayEvent)(unsafe.Pointer(cevent))
		return DisplayEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Display:   uint32(e.display),
			Event:     uint8(e.event),
			Data1:     int32(e.data1),
		}
	case WINDOWEVENT:
		e := (*cWindowEvent)(unsafe.Pointer(cevent))
		return WindowEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			WindowID:  uint32(e.windowID),
			Event:     WindowEventID(e.event),
			Data1:     int32(e.data1),
			Data2:     int32(e.data2),
		}
	case SYSWMEVENT:
		e := (*cSysWMEvent)(unsafe.Pointer(cevent))
		return SysWMEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Msg:       (*SysWMmsg)(unsafe.Pointer(e.msg)),
		}
	case KEYDOWN, KEYUP:
		e := (*cKeyboardEvent)(unsafe.Pointer(cevent))
		return KeyboardEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			WindowID:  uint32(e.windowID),
			State:     ButtonState(e.state),
			Keysym: Keysym{
				Mod:      Keymod(e.keysym.mod),
				Sym:      Keycode(e.keysym.sym),
				Scancode: Scancode(e.keysym.scancode),
			},
		}
	case TEXTEDITING:
		e := (*cTextEditingEvent)(unsafe.Pointer(cevent))
		return TextEditingEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			WindowID:  uint32(e.windowID),
			Text:      C.GoString(&e.text[0]), // text is null-terminated
			Start:     int32(e.start),
			Length:    int32(e.length),
		}
	case TEXTINPUT:
		e := (*cTextInputEvent)(unsafe.Pointer(cevent))
		return TextInputEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			WindowID:  uint32(e.windowID),
			Text:      C.GoString(&e.text[0]),
		}
	case MOUSEMOTION:
		e := (*cMouseMotionEvent)(unsafe.Pointer(cevent))
		return MouseMotionEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			WindowID:  uint32(e.windowID),
			Which:     uint32(e.which),
			State:     ButtonMask(e.state),
			X:         int32(e.x),
			Y:         int32(e.y),
			XRel:      int32(e.xrel),
			YRel:      int32(e.yrel),
		}
	case MOUSEBUTTONDOWN, MOUSEBUTTONUP:
		e := (*cMouseButtonEvent)(unsafe.Pointer(cevent))
		return MouseButtonEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			WindowID:  uint32(e.windowID),
			Which:     uint32(e.which),
			Button:    Button(e.button),
			State:     ButtonState(e.state),
			Clicks:    uint8(e.clicks),
			X:         int32(e.x),
			Y:         int32(e.y),
		}
	case MOUSEWHEEL:
		e := (*cMouseWheelEvent)(unsafe.Pointer(cevent))
		return MouseWheelEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			WindowID:  uint32(e.windowID),
			Which:     uint32(e.which),
			X:         int32(e.x),
			Y:         int32(e.y),
			Direction: uint32(e.direction),
			PreciseX:  float32(e.preciseX),
			PreciseY:  float32(e.preciseY),
			MouseX:    int32(e.mouseX),
			MouseY:    int32(e.mouseY),
		}
	case JOYAXISMOTION:
		e := (*cJoyAxisEvent)(unsafe.Pointer(cevent))
		return JoyAxisEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Which:     JoystickID(e.which),
			Axis:      uint8(e.axis),
			Value:     int16(e.value),
		}
	case JOYBALLMOTION:
		e := (*cJoyBallEvent)(unsafe.Pointer(cevent))
		return JoyBallEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Which:     JoystickID(e.which),
			Ball:      uint8(e.ball),
			XRel:      int16(e.xrel),
			YRel:      int16(e.yrel),
		}
	case JOYHATMOTION:
		e := (*cJoyHatEvent)(unsafe.Pointer(cevent))
		return JoyHatEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Which:     JoystickID(e.which),
			Hat:       uint8(e.hat),
			Value:     JoystickHat(e.value),
		}
	case JOYBUTTONDOWN, JOYBUTTONUP:
		e := (*cJoyButtonEvent)(unsafe.Pointer(cevent))
		return JoyButtonEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Which:     JoystickID(e.which),
			Button:    uint8(e.button),
			State:     ButtonState(e.state),
		}
	case JOYDEVICEADDED:
		e := (*cJoyDeviceEvent)(unsafe.Pointer(cevent))
		return JoyDeviceAddedEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Which:     JoystickID(e.which),
		}
	case JOYDEVICEREMOVED:
		e := (*cJoyDeviceEvent)(unsafe.Pointer(cevent))
		return JoyDeviceRemovedEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Which:     JoystickID(e.which),
		}
	case JOYBATTERYUPDATED:
		e := (*cJoyBatteryEvent)(unsafe.Pointer(cevent))
		return JoyBatteryEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Which:     JoystickID(e.which),
			Level:     JoystickPowerLevel(e.level),
		}
	case CONTROLLERAXISMOTION:
		e := (*cControllerAxisEvent)(unsafe.Pointer(cevent))
		return ControllerAxisEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Which:     JoystickID(e.which),
			Axis:      GameControllerAxis(e.axis),
			Value:     int16(e.value),
		}
	case CONTROLLERBUTTONDOWN, CONTROLLERBUTTONUP:
		e := (*cControllerButtonEvent)(unsafe.Pointer(cevent))
		return ControllerButtonEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Which:     JoystickID(e.which),
			Button:    GameControllerButton(e.button),
			State:     ButtonState(e.state),
		}
	case CONTROLLERDEVICEADDED, CONTROLLERDEVICEREMOVED, CONTROLLERDEVICEREMAPPED:
		e := (*cControllerDeviceEvent)(unsafe.Pointer(cevent))
		return ControllerDeviceEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Which:     JoystickID(e.which),
		}
	case CONTROLLERSENSORUPDATE:
		e := (*cControllerSensorEvent)(unsafe.Pointer(cevent))
		return ControllerSensorEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Which:     JoystickID(e.which),
            Sensor:    int32(e.sensor),
			Data: [3]float32{
				float32(e.data[0]),
				float32(e.data[1]),
				float32(e.data[2]),
			},
            TimestampUs: uint64(e.timestamp_us),
		}
	case AUDIODEVICEADDED, AUDIODEVICEREMOVED:
		e := (*cAudioDeviceEvent)(unsafe.Pointer(cevent))
		return AudioDeviceEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Which:     uint32(e.which),
			IsCapture: uint8(e.iscapture),
		}
	case FINGERMOTION, FINGERDOWN, FINGERUP:
		e := (*cTouchFingerEvent)(unsafe.Pointer(cevent))
		return TouchFingerEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			TouchID:   TouchID(e.touchId),
			FingerID:  FingerID(e.fingerId),
			X:         float32(e.x),
			Y:         float32(e.y),
			DX:        float32(e.dx),
			DY:        float32(e.dy),
			Pressure:  float32(e.pressure),
			WindowID:  uint32(e.windowID),
		}
	case MULTIGESTURE:
		e := (*cMultiGestureEvent)(unsafe.Pointer(cevent))
		return MultiGestureEvent{
			Type:       EventType(e._type),
			Timestamp:  uint32(e.timestamp),
			TouchID:    TouchID(e.touchId),
			DTheta:     float32(e.dTheta),
			DDist:      float32(e.dDist),
			X:          float32(e.x),
			Y:          float32(e.y),
			NumFingers: uint16(e.numFingers),
		}
	case DOLLARGESTURE, DOLLARRECORD:
		e := (*cDollarGestureEvent)(unsafe.Pointer(cevent))
		return DollarGestureEvent{
			Type:       EventType(e._type),
			Timestamp:  uint32(e.timestamp),
			TouchID:    TouchID(e.touchId),
			GestureID:  GestureID(e.gestureId),
			NumFingers: uint32(e.numFingers),
			Error:      float32(e.error),
			X:          float32(e.x),
			Y:          float32(e.y),
		}
	case DROPFILE, DROPTEXT, DROPBEGIN, DROPCOMPLETE:
		e := (*cDropEvent)(unsafe.Pointer(cevent))
		// From SDL doc: the file name, which should be freed with SDL_free(), is NULL on BEGIN/COMPLETE.
		defer func() {
			if e.file != nil {
				C.SDL_free(unsafe.Pointer(e.file))
			}
		}()
		return DropEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			File:      C.GoString(e.file),
			WindowID:  uint32(e.windowID),
		}
	case SENSORUPDATE:
		e := (*cSensorEvent)(unsafe.Pointer(cevent))
		return SensorEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
			Which:     int32(e.which),
			Data: [6]float32{
				float32(e.data[0]),
				float32(e.data[1]),
				float32(e.data[2]),
				float32(e.data[3]),
				float32(e.data[4]),
				float32(e.data[5]),
			},
            TimestampUs: uint64(e.timestamp_us),
		}
	case RENDER_TARGETS_RESET, RENDER_DEVICE_RESET:
		// This is a CommonEvent as it doesn't currently (sdl-v2.0.22) have any additional fields
		e := (*cCommonEvent)(unsafe.Pointer(cevent))
		return RenderEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
		}
	case QUIT:
		// This is a CommonEvent as it doesn't currently (sdl-v2.0.22) have any additional fields
		e := (*cCommonEvent)(unsafe.Pointer(cevent))
		return QuitEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
		}
	case CLIPBOARDUPDATE:
		// This is a CommonEvent as it doesn't currently (sdl-v2.0.22) have any additional fields
		e := (*cCommonEvent)(unsafe.Pointer(cevent))
		return ClipboardEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
		}
	default:
		if cevent.Type >= USEREVENT {
			// all events beyond USEREVENT are UserEvents to be registered with RegisterEvents
			e := (*cUserEvent)(unsafe.Pointer(cevent))
			return UserEvent{
				Type:      EventType(e._type),
				Timestamp: uint32(e.timestamp),
				WindowID:  uint32(e.windowID),
				Code:      int32(e.code),
				Data1:     e.data1,
				Data2:     e.data2,
			}
		}
		e := (*cCommonEvent)(unsafe.Pointer(cevent))
		return CommonEvent{
			Type:      EventType(e._type),
			Timestamp: uint32(e.timestamp),
		}
	}
}

func cEvent(event Event) *CEvent {
	evv := reflect.ValueOf(event)
	p := evv.Elem()
	return (*CEvent)(unsafe.Pointer(p.UnsafeAddr()))
}

// WaitEventTimeout waits until the specified timeout (in milliseconds) for the next available event.
// (https://wiki.libsdl.org/SDL_WaitEventTimeout)
func WaitEventTimeout(timeout int) Event {
	var cevent CEvent
	_event := (*C.SDL_Event)(unsafe.Pointer(&cevent))
	ok := int(C.SDL_WaitEventTimeout(_event, C.int(timeout)))
	if ok == 0 {
		return nil
	}
	return goEvent(&cevent)
}

// WaitEvent waits indefinitely for the next available event.
// (https://wiki.libsdl.org/SDL_WaitEvent)
func WaitEvent() Event {
	var cevent CEvent
	_event := (*C.SDL_Event)(unsafe.Pointer(&cevent))
	ok := int(C.SDL_WaitEvent(_event))
	if ok == 0 {
		return nil
	}
	return goEvent(&cevent)
}

// PushEvent adds an event to the event queue.
// (https://wiki.libsdl.org/SDL_PushEvent)
func PushEvent(event Event) (filtered bool, err error) {
	_event := (*C.SDL_Event)(unsafe.Pointer(cEvent(event)))
	if ok := int(C.SDL_PushEvent(_event)); ok < 0 {
		filtered, err = false, GetError()
	} else if ok == 0 {
		filtered, err = true, nil
	}
	return
}

func (ef eventFilterFunc) FilterEvent(e Event, userdata interface{}) bool {
	return ef(e, userdata)
}

func newEventFilterCallbackContext(filter EventFilter, userdata interface{}) *eventFilterCallbackContext {
	lastEventWatchHandleMutex.Lock()
	defer lastEventWatchHandleMutex.Unlock()
	// Look for the next available watch handle (this should be immediate
	// unless you're creating a LOT of handlers).
	for {
		if _, ok := eventWatches[lastEventWatchHandle]; !ok {
			break
		}
		lastEventWatchHandle++
	}
	e := &eventFilterCallbackContext{filter, lastEventWatchHandle, userdata}
	eventWatches[lastEventWatchHandle] = e
	lastEventWatchHandle++
	return e
}

func (e *eventFilterCallbackContext) cptr() unsafe.Pointer {
	return unsafe.Pointer(e.handle)
}

//export goSetEventFilterCallback
func goSetEventFilterCallback(data unsafe.Pointer, e *C.SDL_Event) C.int {
	// No check for eventFilterCache != nil. Why? because it should never be
	// nil since the callback is set/unset based on the last filter being nil
	// /non-nil. If there is an issue, then it should panic here so we can
	// figure out why that is.

	return wrapEventFilterCallback(eventFilterCache, e, nil)
}

//export goEventFilterCallback
func goEventFilterCallback(userdata unsafe.Pointer, e *C.SDL_Event) C.int {
	// same sort of reasoning with goSetEventFilterCallback, userdata should
	// always be non-nil and represent a valid eventFilterCallbackContext. If
	// it doesn't a panic will let us know that there something wrong and the
	// problem can be fixed.

	context := eventWatches[EventWatchHandle(userdata)]
	return wrapEventFilterCallback(context.filter, e, context.userdata)
}

func wrapEventFilterCallback(filter EventFilter, e *C.SDL_Event, userdata interface{}) C.int {
	gev := goEvent((*CEvent)(unsafe.Pointer(e)))
	result := filter.FilterEvent(gev, userdata)

	if result {
		return C.SDL_TRUE
	}
	return C.SDL_FALSE
}

// SetEventFilter sets up a filter to process all events before they change internal state and are posted to the internal event queue.
// (https://wiki.libsdl.org/SDL_SetEventFilter)
func SetEventFilter(filter EventFilter, userdata interface{}) {
	if eventFilterCache == nil && filter == nil {
		// nothing to do...
		return
	}

	if eventFilterCache == nil && filter != nil {
		// We had no event filter before and do now; lets set
		// goSetEventFilterCallback() as the event filter.
		C.setEventFilter()
	} else if eventFilterCache != nil && filter == nil {
		// We had an event filter before, but no longer do, lets clear the
		// event filter
		C.clearEventFilter()
	}

	eventFilterCache = filter
}

// SetEventFilterFunc sets up a function to process all events before they change internal state and are posted to the internal event queue.
// (https://wiki.libsdl.org/SDL_SetEventFilter)
func SetEventFilterFunc(filterFunc eventFilterFunc, userdata interface{}) {
	SetEventFilter(filterFunc, userdata)
}

// GetEventFilter queries the current event filter.
// (https://wiki.libsdl.org/SDL_GetEventFilter)
func GetEventFilter() EventFilter {
	return eventFilterCache
}

func isCEventFilterSet() bool {
	return C.SDL_GetEventFilter(nil, nil) == C.SDL_TRUE
}

// FilterEvents run a specific filter function on the current event queue, removing any events for which the filter returns 0.
// (https://wiki.libsdl.org/SDL_FilterEvents)
func FilterEvents(filter EventFilter, userdata interface{}) {
	context := newEventFilterCallbackContext(filter, userdata)
	C.filterEvents(context.cptr())
}

// FilterEventsFunc run a specific function on the current event queue, removing any events for which the filter returns 0.
// (https://wiki.libsdl.org/SDL_FilterEvents)
func FilterEventsFunc(filter eventFilterFunc, userdata interface{}) {
	FilterEvents(filter, userdata)
}

// AddEventWatch adds a callback to be triggered when an event is added to the event queue.
// (https://wiki.libsdl.org/SDL_AddEventWatch)
func AddEventWatch(filter EventFilter, userdata interface{}) EventWatchHandle {
	context := newEventFilterCallbackContext(filter, userdata)
	C.addEventWatch(context.cptr())
	return context.handle
}

// AddEventWatchFunc adds a callback function to be triggered when an event is added to the event queue.
// (https://wiki.libsdl.org/SDL_AddEventWatch)
func AddEventWatchFunc(filterFunc eventFilterFunc, userdata interface{}) EventWatchHandle {
	return AddEventWatch(filterFunc, userdata)
}

// DelEventWatch removes an event watch callback added with AddEventWatch().
// (https://wiki.libsdl.org/SDL_DelEventWatch)
func DelEventWatch(handle EventWatchHandle) {
	context, ok := eventWatches[handle]
	if !ok {
		return
	}
	delete(eventWatches, context.handle)
	C.delEventWatch(context.cptr())
}

// EventState sets the state of processing events by type.
// (https://wiki.libsdl.org/SDL_EventState)
func EventState(type_ EventType, state EventStateConstant) uint8 {
	return uint8(C.SDL_EventState(C.Uint32(type_), C.int(state)))
}

// GetEventState returns the current processing state of the specified event
// (https://wiki.libsdl.org/SDL_EventState)
func GetEventState(type_ EventType) uint8 {
	return uint8(C.SDL_EventState(C.Uint32(type_), C.int(QUERY)))
}

// RegisterEvents allocates a set of user-defined events, and return the beginning event number for that set of events.
// (https://wiki.libsdl.org/SDL_RegisterEvents)
func RegisterEvents(numEvents int) uint32 {
	return uint32(C.SDL_RegisterEvents(C.int(numEvents)))
}
