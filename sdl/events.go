package sdl

/*
#include "sdl_wrapper.h"
#include "events.h"

#if !SDL_VERSION_ATLEAST(2,0,9)
#define SDL_DISPLAYEVENT (0x150)
#endif

#if !SDL_VERSION_ATLEAST(2,0,2)
#define SDL_RENDER_TARGETS_RESET (0x2000)
#endif

#if !SDL_VERSION_ATLEAST(2,0,4)

#if defined(WARN_OUTDATED)
#pragma message("SDL_KEYMAPCHANGED is not supported before SDL 2.0.4")
#endif

#define SDL_KEYMAPCHANGED (0x304)


#if defined(WARN_OUTDATED)
#pragma message("SDL_AUDIODEVICEADDED is not supported before SDL 2.0.4")
#endif

#define SDL_AUDIODEVICEADDED (0x1100)


#if defined(WARN_OUTDATED)
#pragma message("SDL_AUDIODEVICEREMOVED is not supported before SDL 2.0.4")
#endif

#define SDL_AUDIODEVICEREMOVED (0x1101)


#if defined(WARN_OUTDATED)
#pragma message("SDL_RENDER_DEVICE_RESET is not supported before SDL 2.0.4")
#endif

#define SDL_RENDER_DEVICE_RESET (0x2001)


#if defined(WARN_OUTDATED)
#pragma message("SDL_AudioDeviceEvent is not supported before SDL 2.0.4")
#endif

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
#endif

#define SDL_DROPTEXT (0x1001)


#if defined(WARN_OUTDATED)
#pragma message("SDL_DROPBEGIN is not supported before SDL 2.0.5")
#endif

#define SDL_DROPBEGIN (0x1002)


#if defined(WARN_OUTDATED)
#pragma message("SDL_DROPCOMPLETE is not supported before SDL 2.0.5")
#endif

#define SDL_DROPCOMPLETE (0x1003)
#endif

#if !SDL_VERSION_ATLEAST(2,0,9)
#define SDL_SENSORUPDATE (0x1200)

typedef struct SDL_SensorEvent {
    Uint32 type;
    Uint32 timestamp;
    Sint32 which;
    float data[6];
} SDL_SensorEvent;
#endif

#if !SDL_VERSION_ATLEAST(2,0,22)
#if defined(WARN_OUTDATED)
#pragma message("SDL_TEXTEDITING_EXT is not supported before SDL 2.0.22")
#endif

#define SDL_TEXTEDITING_EXT (0x305)
#endif

// NOTE: To prevent build from failing when using older SDL2, we create a
// structure definiton that directly maps to the SDL2 struct definition if
// using the latest SDL2. Otherwise, we copy the latest definition and paste
// it here.
#if SDL_VERSION_ATLEAST(2,0,22)
typedef SDL_MouseButtonEvent MouseButtonEvent;
typedef SDL_MouseWheelEvent MouseWheelEvent;
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
} MouseWheelEvent;

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
#pragma message("SDL_JOYBATTEYUPDATED is not supported before SDL 2.24.0")
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
const (
	FIRSTEVENT = C.SDL_FIRSTEVENT // do not remove (unused)

	// Application events
	QUIT = C.SDL_QUIT // user-requested quit

	// Android, iOS and WinRT events
	APP_TERMINATING         = C.SDL_APP_TERMINATING         // OS is terminating the application
	APP_LOWMEMORY           = C.SDL_APP_LOWMEMORY           // OS is low on memory; free some
	APP_WILLENTERBACKGROUND = C.SDL_APP_WILLENTERBACKGROUND // application is entering background
	APP_DIDENTERBACKGROUND  = C.SDL_APP_DIDENTERBACKGROUND  //application entered background
	APP_WILLENTERFOREGROUND = C.SDL_APP_WILLENTERFOREGROUND // application is entering foreground
	APP_DIDENTERFOREGROUND  = C.SDL_APP_DIDENTERFOREGROUND  // application entered foreground

	// Display events
	DISPLAYEVENT = C.SDL_DISPLAYEVENT // Display state change

	// Window events
	WINDOWEVENT = C.SDL_WINDOWEVENT // window state change
	SYSWMEVENT  = C.SDL_SYSWMEVENT  // system specific event

	// Keyboard events
	KEYDOWN         = C.SDL_KEYDOWN         // key pressed
	KEYUP           = C.SDL_KEYUP           // key released
	TEXTEDITING     = C.SDL_TEXTEDITING     // keyboard text editing (composition)
	TEXTINPUT       = C.SDL_TEXTINPUT       // keyboard text input
	TEXTEDITING_EXT = C.SDL_TEXTEDITING_EXT // keyboard text editing (composition)
	KEYMAPCHANGED   = C.SDL_KEYMAPCHANGED   // keymap changed due to a system event such as an input language or keyboard layout change (>= SDL 2.0.4)

	// Mouse events
	MOUSEMOTION     = C.SDL_MOUSEMOTION     // mouse moved
	MOUSEBUTTONDOWN = C.SDL_MOUSEBUTTONDOWN // mouse button pressed
	MOUSEBUTTONUP   = C.SDL_MOUSEBUTTONUP   // mouse button released
	MOUSEWHEEL      = C.SDL_MOUSEWHEEL      // mouse wheel motion

	// Joystick events
	JOYAXISMOTION    = C.SDL_JOYAXISMOTION    // joystick axis motion
	JOYBALLMOTION    = C.SDL_JOYBALLMOTION    // joystick trackball motion
	JOYHATMOTION     = C.SDL_JOYHATMOTION     // joystick hat position change
	JOYBUTTONDOWN    = C.SDL_JOYBUTTONDOWN    // joystick button pressed
	JOYBUTTONUP      = C.SDL_JOYBUTTONUP      // joystick button released
	JOYDEVICEADDED   = C.SDL_JOYDEVICEADDED   // joystick connected
	JOYDEVICEREMOVED = C.SDL_JOYDEVICEREMOVED // joystick disconnected

	// Game controller events
	CONTROLLERAXISMOTION     = C.SDL_CONTROLLERAXISMOTION     // controller axis motion
	CONTROLLERBUTTONDOWN     = C.SDL_CONTROLLERBUTTONDOWN     // controller button pressed
	CONTROLLERBUTTONUP       = C.SDL_CONTROLLERBUTTONUP       // controller button released
	CONTROLLERDEVICEADDED    = C.SDL_CONTROLLERDEVICEADDED    // controller connected
	CONTROLLERDEVICEREMOVED  = C.SDL_CONTROLLERDEVICEREMOVED  // controller disconnected
	CONTROLLERDEVICEREMAPPED = C.SDL_CONTROLLERDEVICEREMAPPED // controller mapping updated

	// Touch events
	FINGERDOWN   = C.SDL_FINGERDOWN   // user has touched input device
	FINGERUP     = C.SDL_FINGERUP     // user stopped touching input device
	FINGERMOTION = C.SDL_FINGERMOTION // user is dragging finger on input device

	// Gesture events
	DOLLARGESTURE = C.SDL_DOLLARGESTURE
	DOLLARRECORD  = C.SDL_DOLLARRECORD
	MULTIGESTURE  = C.SDL_MULTIGESTURE

	// Clipboard events
	CLIPBOARDUPDATE = C.SDL_CLIPBOARDUPDATE // the clipboard changed

	// Drag and drop events
	DROPFILE     = C.SDL_DROPFILE     // the system requests a file open
	DROPTEXT     = C.SDL_DROPTEXT     // text/plain drag-and-drop event
	DROPBEGIN    = C.SDL_DROPBEGIN    // a new set of drops is beginning (NULL filename)
	DROPCOMPLETE = C.SDL_DROPCOMPLETE // current set of drops is now complete (NULL filename)

	// Audio hotplug events
	AUDIODEVICEADDED   = C.SDL_AUDIODEVICEADDED   // a new audio device is available (>= SDL 2.0.4)
	AUDIODEVICEREMOVED = C.SDL_AUDIODEVICEREMOVED // an audio device has been removed (>= SDL 2.0.4)

	// Sensor events
	SENSORUPDATE = C.SDL_SENSORUPDATE // a sensor was updated

	// Render events
	RENDER_TARGETS_RESET = C.SDL_RENDER_TARGETS_RESET // the render targets have been reset and their contents need to be updated (>= SDL 2.0.2)
	RENDER_DEVICE_RESET  = C.SDL_RENDER_DEVICE_RESET  // the device has been reset and all textures need to be recreated (>= SDL 2.0.4)

	// These are for your use, and should be allocated with RegisterEvents()
	USEREVENT = C.SDL_USEREVENT // a user-specified event
	LASTEVENT = C.SDL_LASTEVENT // (only for bounding internal arrays)
)

// Actions for PeepEvents().
// (https://wiki.libsdl.org/SDL_PeepEvents)
const (
	ADDEVENT  = C.SDL_ADDEVENT  // up to numevents events will be added to the back of the event queue
	PEEKEVENT = C.SDL_PEEKEVENT // up to numevents events at the front of the event queue, within the specified minimum and maximum type, will be returned and will not be removed from the queue
	GETEVENT  = C.SDL_GETEVENT  // up to numevents events at the front of the event queue, within the specified minimum and maximum type, will be returned and will be removed from the queue
)

// Toggles for different event state functions.
const (
	QUERY   = C.SDL_QUERY
	IGNORE  = C.SDL_IGNORE
	DISABLE = C.SDL_DISABLE
	ENABLE  = C.SDL_ENABLE
)

// Event is a union of all event structures used in SDL.
// (https://wiki.libsdl.org/SDL_Event)
type Event interface {
	GetType() uint32      // GetType returns the event type
	GetTimestamp() uint32 // GetTimestamp returns the timestamp of the event
}

// CEvent is a union of all event structures used in SDL.
// (https://wiki.libsdl.org/SDL_Event)
type CEvent struct {
	Type uint32
	_    [52]byte // padding
}

// CommonEvent contains common event data.
// (https://wiki.libsdl.org/SDL_Event)
type CommonEvent struct {
	Type      uint32 // the event type
	Timestamp uint32 // timestamp of the event
}

// GetType returns the event type.
func (e *CommonEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *CommonEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// DisplayEvent contains common event data.
// (https://wiki.libsdl.org/SDL_Event)
type DisplayEvent struct {
	Type      uint32 // the event type
	Timestamp uint32 // timestamp of the event
	Display   uint32 // the associated display index
	Event     uint8  // TODO: (https://wiki.libsdl.org/SDL_DisplayEventID)
	_         uint8  // padding
	_         uint8  // padding
	_         uint8  // padding
	Data1     int32  // event dependent data
}

// GetType returns the event type.
func (e *DisplayEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *DisplayEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// WindowEvent contains window state change event data.
// (https://wiki.libsdl.org/SDL_WindowEvent)
type WindowEvent struct {
	Type      uint32 // WINDOWEVENT
	Timestamp uint32 // timestamp of the event
	WindowID  uint32 // the associated window
	Event     uint8  // (https://wiki.libsdl.org/SDL_WindowEventID)
	_         uint8  // padding
	_         uint8  // padding
	_         uint8  // padding
	Data1     int32  // event dependent data
	Data2     int32  // event dependent data
}
type cWindowEvent C.SDL_WindowEvent

// GetType returns the event type.
func (e *WindowEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *WindowEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// KeyboardEvent contains keyboard key down event information.
// (https://wiki.libsdl.org/SDL_KeyboardEvent)
type KeyboardEvent struct {
	Type      uint32 // KEYDOWN, KEYUP
	Timestamp uint32 // timestamp of the event
	WindowID  uint32 // the window with keyboard focus, if any
	State     uint8  // PRESSED, RELEASED
	Repeat    uint8  // non-zero if this is a key repeat
	_         uint8  // padding
	_         uint8  // padding
	Keysym    Keysym // Keysym representing the key that was pressed or released
}
type cKeyboardEvent C.SDL_KeyboardEvent

// GetType returns the event type.
func (e *KeyboardEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *KeyboardEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// TextEditingEvent contains keyboard text editing event information.
// (https://wiki.libsdl.org/SDL_TextEditingEvent)
type TextEditingEvent struct {
	Type      uint32                               // TEXTEDITING
	Timestamp uint32                               // timestamp of the event
	WindowID  uint32                               // the window with keyboard focus, if any
	Text      [C.SDL_TEXTINPUTEVENT_TEXT_SIZE]byte // the null-terminated editing text in UTF-8 encoding
	Start     int32                                // the location to begin editing from
	Length    int32                                // the number of characters to edit from the start point
}
type cTextEditingEvent C.SDL_TextEditingEvent

// GetType returns the event type.
func (e *TextEditingEvent) GetType() uint32 {
	return e.Type
}

// GetText returns the text as string
func (e *TextEditingEvent) GetText() string {
	length := func(buf []byte) int {
		for i := range buf {
			if buf[i] == 0 {
				return i
			}
		}

		return 0
	}(e.Text[:])

	text := e.Text[:length]
	return string(text)
}

// GetTimestamp returns the timestamp of the event.
func (e *TextEditingEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// TextInputEvent contains keyboard text input event information.
// (https://wiki.libsdl.org/SDL_TextInputEvent)
type TextInputEvent struct {
	Type      uint32                               // TEXTINPUT
	Timestamp uint32                               // timestamp of the event
	WindowID  uint32                               // the window with keyboard focus, if any
	Text      [C.SDL_TEXTINPUTEVENT_TEXT_SIZE]byte // the null-terminated input text in UTF-8 encoding
}
type cTextInputEvent C.SDL_TextInputEvent

// GetType returns the event type.
func (e *TextInputEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *TextInputEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// GetText returns the text as string
func (e *TextInputEvent) GetText() string {
	length := func(buf []byte) int {
		for i := range buf {
			if buf[i] == 0 {
				return i
			}
		}

		return 0
	}(e.Text[:])

	text := e.Text[:length]
	return string(text)
}

// MouseMotionEvent contains mouse motion event information.
// (https://wiki.libsdl.org/SDL_MouseMotionEvent)
type MouseMotionEvent struct {
	Type      uint32 // MOUSEMOTION
	Timestamp uint32 // timestamp of the event
	WindowID  uint32 // the window with mouse focus, if any
	Which     uint32 // the mouse instance id, or TOUCH_MOUSEID
	State     uint32 // BUTTON_LEFT, BUTTON_MIDDLE, BUTTON_RIGHT, BUTTON_X1, BUTTON_X2
	X         int32  // X coordinate, relative to window
	Y         int32  // Y coordinate, relative to window
	XRel      int32  // relative motion in the X direction
	YRel      int32  // relative motion in the Y direction
}
type cMouseMotionEvent C.SDL_MouseMotionEvent

// GetType returns the event type.
func (e *MouseMotionEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *MouseMotionEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// MouseButtonEvent contains mouse button event information.
// (https://wiki.libsdl.org/SDL_MouseButtonEvent)
type MouseButtonEvent struct {
	Type      uint32 // MOUSEBUTTONDOWN, MOUSEBUTTONUP
	Timestamp uint32 // timestamp of the event
	WindowID  uint32 // the window with mouse focus, if any
	Which     uint32 // the mouse instance id, or TOUCH_MOUSEID
	Button    uint8  // BUTTON_LEFT, BUTTON_MIDDLE, BUTTON_RIGHT, BUTTON_X1, BUTTON_X2
	State     uint8  // PRESSED, RELEASED
	Clicks    uint8  // 1 for single-click, 2 for double-click, etc. (>= SDL 2.0.2)
	_         uint8  // padding
	X         int32  // X coordinate, relative to window
	Y         int32  // Y coordinate, relative to window
}
type cMouseButtonEvent C.SDL_MouseButtonEvent

// GetType returns the event type.
func (e *MouseButtonEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *MouseButtonEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// MouseWheelEvent contains mouse wheel event information.
// (https://wiki.libsdl.org/SDL_MouseWheelEvent)
type MouseWheelEvent struct {
	Type      uint32  // MOUSEWHEEL
	Timestamp uint32  // timestamp of the event
	WindowID  uint32  // the window with mouse focus, if any
	Which     uint32  // the mouse instance id, or TOUCH_MOUSEID
	X         int32   // the amount scrolled horizontally, positive to the right and negative to the left
	Y         int32   // the amount scrolled vertically, positive away from the user and negative toward the user
	Direction uint32  // MOUSEWHEEL_NORMAL, MOUSEWHEEL_FLIPPED (>= SDL 2.0.4)
	PreciseX  float32 // The amount scrolled horizontally, positive to the right and negative to the left, with float precision (added in 2.0.18)
	PreciseY  float32 // The amount scrolled vertically, positive away from the user and negative toward the user, with float precision (added in 2.0.18)
}
type cMouseWheelEvent C.SDL_MouseWheelEvent

// GetType returns the event type.
func (e *MouseWheelEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *MouseWheelEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// JoyAxisEvent contains joystick axis motion event information.
// (https://wiki.libsdl.org/SDL_JoyAxisEvent)
type JoyAxisEvent struct {
	Type      uint32     // JOYAXISMOTION
	Timestamp uint32     // timestamp of the event
	Which     JoystickID // the instance id of the joystick that reported the event
	Axis      uint8      // the index of the axis that changed
	_         uint8      // padding
	_         uint8      // padding
	_         uint8      // padding
	Value     int16      // the current position of the axis (range: -32768 to 32767)
	_         uint16     // padding
}
type cJoyAxisEvent C.SDL_JoyAxisEvent

// GetType returns the event type.
func (e *JoyAxisEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *JoyAxisEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// JoyBallEvent contains joystick trackball motion event information.
// (https://wiki.libsdl.org/SDL_JoyBallEvent)
type JoyBallEvent struct {
	Type      uint32     // JOYBALLMOTION
	Timestamp uint32     // timestamp of the event
	Which     JoystickID // the instance id of the joystick that reported the event
	Ball      uint8      // the index of the trackball that changed
	_         uint8      // padding
	_         uint8      // padding
	_         uint8      // padding
	XRel      int16      // the relative motion in the X direction
	YRel      int16      // the relative motion in the Y direction
}
type cJoyBallEvent C.SDL_JoyBallEvent

// GetType returns the event type.
func (e *JoyBallEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *JoyBallEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// JoyHatEvent contains joystick hat position change event information.
// (https://wiki.libsdl.org/SDL_JoyHatEvent)
type JoyHatEvent struct {
	Type      uint32     // JOYHATMOTION
	Timestamp uint32     // timestamp of the event
	Which     JoystickID // the instance id of the joystick that reported the event
	Hat       uint8      // the index of the hat that changed
	Value     uint8      // HAT_LEFTUP, HAT_UP, HAT_RIGHTUP, HAT_LEFT, HAT_CENTERED, HAT_RIGHT, HAT_LEFTDOWN, HAT_DOWN, HAT_RIGHTDOWN
	_         uint8      // padding
	_         uint8      // padding
}
type cJoyHatEvent C.SDL_JoyHatEvent

// GetType returns the event type.
func (e *JoyHatEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *JoyHatEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// JoyButtonEvent contains joystick button event information.
// (https://wiki.libsdl.org/SDL_JoyButtonEvent)
type JoyButtonEvent struct {
	Type      uint32     // JOYBUTTONDOWN, JOYBUTTONUP
	Timestamp uint32     // timestamp of the event
	Which     JoystickID // the instance id of the joystick that reported the event
	Button    uint8      // the index of the button that changed
	State     uint8      // PRESSED, RELEASED
	_         uint8      // padding
	_         uint8      // padding
}
type cJoyButtonEvent C.SDL_JoyButtonEvent

// GetType returns the event type.
func (e *JoyButtonEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *JoyButtonEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// JoyDeviceAddedEvent contains joystick device event information.
// (https://wiki.libsdl.org/SDL_JoyDeviceEvent)
type JoyDeviceAddedEvent struct {
	Type      uint32     // JOYDEVICEADDED
	Timestamp uint32     // the timestamp of the event
	Which     JoystickID // the joystick device index
}

// GetType returns the event type.
func (e *JoyDeviceAddedEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *JoyDeviceAddedEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// JoyDeviceRemovedEvent contains joystick device event information.
// (https://wiki.libsdl.org/SDL_JoyDeviceEvent)
type JoyDeviceRemovedEvent struct {
	Type      uint32     // JOYDEVICEREMOVED
	Timestamp uint32     // the timestamp of the event
	Which     JoystickID // the instance id
}

// GetType returns the event type.
func (e *JoyDeviceRemovedEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *JoyDeviceRemovedEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// ControllerAxisEvent contains game controller axis motion event information.
// (https://wiki.libsdl.org/SDL_ControllerAxisEvent)
type ControllerAxisEvent struct {
	Type      uint32     // CONTROLLERAXISMOTION
	Timestamp uint32     // the timestamp of the event
	Which     JoystickID // the joystick instance id
	Axis      uint8      // the controller axis (https://wiki.libsdl.org/SDL_GameControllerAxis)
	_         uint8      // padding
	_         uint8      // padding
	_         uint8      // padding
	Value     int16      // the axis value (range: -32768 to 32767)
	_         uint16     // padding
}
type cControllerAxisEvent C.SDL_ControllerAxisEvent

// GetType returns the event type.
func (e *ControllerAxisEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *ControllerAxisEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// ControllerButtonEvent contains game controller button event information.
// (https://wiki.libsdl.org/SDL_ControllerButtonEvent)
type ControllerButtonEvent struct {
	Type      uint32     // CONTROLLERBUTTONDOWN, CONTROLLERBUTTONUP
	Timestamp uint32     // the timestamp of the event
	Which     JoystickID // the joystick instance id
	Button    uint8      // the controller button (https://wiki.libsdl.org/SDL_GameControllerButton)
	State     uint8      // PRESSED, RELEASED
	_         uint8      // padding
	_         uint8      // padding
}
type cControllerButtonEvent C.SDL_ControllerButtonEvent

// GetType returns the event type.
func (e *ControllerButtonEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *ControllerButtonEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// ControllerDeviceEvent contains controller device event information.
// (https://wiki.libsdl.org/SDL_ControllerDeviceEvent)
type ControllerDeviceEvent struct {
	Type      uint32     // CONTROLLERDEVICEADDED, CONTROLLERDEVICEREMOVED, SDL_CONTROLLERDEVICEREMAPPED
	Timestamp uint32     // the timestamp of the event
	Which     JoystickID // the joystick device index for the CONTROLLERDEVICEADDED event or instance id for the CONTROLLERDEVICEREMOVED or CONTROLLERDEVICEREMAPPED event
}
type cControllerDeviceEvent C.SDL_ControllerDeviceEvent

// GetType returns the event type.
func (e *ControllerDeviceEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *ControllerDeviceEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// AudioDeviceEvent contains audio device event information.
// (https://wiki.libsdl.org/SDL_AudioDeviceEvent)
type AudioDeviceEvent struct {
	Type      uint32 // AUDIODEVICEADDED, AUDIODEVICEREMOVED
	Timestamp uint32 // the timestamp of the event
	Which     uint32 // the audio device index for the AUDIODEVICEADDED event (valid until next GetNumAudioDevices() call), AudioDeviceID for the AUDIODEVICEREMOVED event
	IsCapture uint8  // zero if an audio output device, non-zero if an audio capture device
	_         uint8  // padding
	_         uint8  // padding
	_         uint8  // padding
}
type cAudioDeviceEvent C.SDL_AudioDeviceEvent

// GetType returns the event type.
func (e *AudioDeviceEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *AudioDeviceEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// TouchFingerEvent contains finger touch event information.
// (https://wiki.libsdl.org/SDL_TouchFingerEvent)
type TouchFingerEvent struct {
	Type      uint32   // FINGERMOTION, FINGERDOWN, FINGERUP
	Timestamp uint32   // timestamp of the event
	TouchID   TouchID  // the touch device id
	FingerID  FingerID // the finger id
	X         float32  // the x-axis location of the touch event, normalized (0...1)
	Y         float32  // the y-axis location of the touch event, normalized (0...1)
	DX        float32  // the distance moved in the x-axis, normalized (-1...1)
	DY        float32  // the distance moved in the y-axis, normalized (-1...1)
	Pressure  float32  // the quantity of pressure applied, normalized (0...1)
}
type cTouchFingerEvent C.SDL_TouchFingerEvent

// GetType returns the event type.
func (e *TouchFingerEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *TouchFingerEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// MultiGestureEvent contains multiple finger gesture event information.
// (https://wiki.libsdl.org/SDL_MultiGestureEvent)
type MultiGestureEvent struct {
	Type       uint32  // MULTIGESTURE
	Timestamp  uint32  // timestamp of the event
	TouchID    TouchID // the touch device id
	DTheta     float32 // the amount that the fingers rotated during this motion
	DDist      float32 // the amount that the fingers pinched during this motion
	X          float32 // the normalized center of gesture
	Y          float32 // the normalized center of gesture
	NumFingers uint16  // the number of fingers used in the gesture
	_          uint16  // padding
}
type cMultiGestureEvent C.SDL_MultiGestureEvent

// GetType returns the event type.
func (e *MultiGestureEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *MultiGestureEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// DollarGestureEvent contains complex gesture event information.
// (https://wiki.libsdl.org/SDL_DollarGestureEvent)
type DollarGestureEvent struct {
	Type       uint32    // DOLLARGESTURE, DOLLARRECORD
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
func (e *DollarGestureEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *DollarGestureEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// DropEvent contains an event used to request a file open by the system.
// (https://wiki.libsdl.org/SDL_DropEvent)
type DropEvent struct {
	Type      uint32 // DROPFILE, DROPTEXT, DROPBEGIN, DROPCOMPLETE
	Timestamp uint32 // timestamp of the event
	File      string // the file name
	WindowID  uint32 // the window that was dropped on, if any
}

type tDropEvent struct {
	Type      uint32
	Timestamp uint32
	File      unsafe.Pointer
	WindowID  uint32
}
type cDropEvent C.SDL_DropEvent

// GetType returns the event type.
func (e *DropEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *DropEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// SensorEvent contains data from sensors such as accelerometer and gyroscope
// (https://wiki.libsdl.org/SDL_SensorEvent)
type SensorEvent struct {
	Type      uint32     // SDL_SENSORUPDATE
	Timestamp uint32     // In milliseconds, populated using SDL_GetTicks()
	Which     int32      // The instance ID of the sensor
	Data      [6]float32 // Up to 6 values from the sensor - additional values can be queried using SDL_SensorGetData()
}
type cSensorEvent C.SDL_SensorEvent

// GetType returns the event type.
func (e *SensorEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *SensorEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// RenderEvent contains render event information.
// (https://wiki.libsdl.org/SDL_EventType)
type RenderEvent struct {
	Type      uint32 // the event type
	Timestamp uint32 // timestamp of the event
}

// GetType returns the event type.
func (e *RenderEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *RenderEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// QuitEvent contains the "quit requested" event.
// (https://wiki.libsdl.org/SDL_QuitEvent)
type QuitEvent struct {
	Type      uint32 // QUIT
	Timestamp uint32 // timestamp of the event
}

// GetType returns the event type.
func (e *QuitEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *QuitEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// OSEvent contains OS specific event information.
type OSEvent struct {
	Type      uint32 // the event type
	Timestamp uint32 // timestamp of the event
}

// GetType returns the event type.
func (e *OSEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *OSEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// ClipboardEvent contains clipboard event information.
// (https://wiki.libsdl.org/SDL_EventType)
type ClipboardEvent struct {
	Type      uint32 // CLIPBOARDUPDATE
	Timestamp uint32 // timestamp of the event
}

// GetType returns the event type.
func (e *ClipboardEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *ClipboardEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// UserEvent contains an application-defined event type.
// (https://wiki.libsdl.org/SDL_UserEvent)
type UserEvent struct {
	Type      uint32         // value obtained from RegisterEvents()
	Timestamp uint32         // timestamp of the event
	WindowID  uint32         // the associated window, if any
	Code      int32          // user defined event code
	Data1     unsafe.Pointer // user defined data pointer
	Data2     unsafe.Pointer // user defined data pointer
}
type cUserEvent C.SDL_UserEvent

// GetType returns the event type.
func (e *UserEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *UserEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// SysWMEvent contains a video driver dependent system event.
// (https://wiki.libsdl.org/SDL_SysWMEvent)
type SysWMEvent struct {
	Type      uint32    // SYSWMEVENT
	Timestamp uint32    // timestamp of the event
	Msg       *SysWMmsg // driver dependent data, defined in SDL_syswm.h
}
type cSysWMEvent C.SDL_SysWMEvent

// GetType returns the event type.
func (e *SysWMEvent) GetType() uint32 {
	return e.Type
}

// GetTimestamp returns the timestamp of the event.
func (e *SysWMEvent) GetTimestamp() uint32 {
	return e.Timestamp
}

// EventAction is the action to take in PeepEvents() function.
// (https://wiki.libsdl.org/SDL_PeepEvents)
type EventAction C.SDL_eventaction

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
func PeepEvents(events []Event, action EventAction, minType, maxType uint32) (storedEvents int, err error) {
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
func HasEvent(type_ uint32) bool {
	return C.SDL_HasEvent(C.Uint32(type_)) != 0
}

// HasEvents checks for the existence of a range of event types in the event queue.
// (https://wiki.libsdl.org/SDL_HasEvents)
func HasEvents(minType, maxType uint32) bool {
	return C.SDL_HasEvents(C.Uint32(minType), C.Uint32(maxType)) != 0
}

// FlushEvent clears events from the event queue.
// (https://wiki.libsdl.org/SDL_FlushEvent)
func FlushEvent(type_ uint32) {
	C.SDL_FlushEvent(C.Uint32(type_))
}

// FlushEvents clears events from the event queue.
// (https://wiki.libsdl.org/SDL_FlushEvents)
func FlushEvents(minType, maxType uint32) {
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
		return (*DisplayEvent)(unsafe.Pointer(cevent))
	case WINDOWEVENT:
		return (*WindowEvent)(unsafe.Pointer(cevent))
	case SYSWMEVENT:
		return (*SysWMEvent)(unsafe.Pointer(cevent))
	case KEYDOWN, KEYUP:
		return (*KeyboardEvent)(unsafe.Pointer(cevent))
	case TEXTEDITING:
		return (*TextEditingEvent)(unsafe.Pointer(cevent))
	case TEXTINPUT:
		return (*TextInputEvent)(unsafe.Pointer(cevent))
	case MOUSEMOTION:
		return (*MouseMotionEvent)(unsafe.Pointer(cevent))
	case MOUSEBUTTONDOWN, MOUSEBUTTONUP:
		return (*MouseButtonEvent)(unsafe.Pointer(cevent))
	case MOUSEWHEEL:
		return (*MouseWheelEvent)(unsafe.Pointer(cevent))
	case JOYAXISMOTION:
		return (*JoyAxisEvent)(unsafe.Pointer(cevent))
	case JOYBALLMOTION:
		return (*JoyBallEvent)(unsafe.Pointer(cevent))
	case JOYHATMOTION:
		return (*JoyHatEvent)(unsafe.Pointer(cevent))
	case JOYBUTTONDOWN, JOYBUTTONUP:
		return (*JoyButtonEvent)(unsafe.Pointer(cevent))
	case JOYDEVICEADDED:
		return (*JoyDeviceAddedEvent)(unsafe.Pointer(cevent))
	case JOYDEVICEREMOVED:
		return (*JoyDeviceRemovedEvent)(unsafe.Pointer(cevent))
	case CONTROLLERAXISMOTION:
		return (*ControllerAxisEvent)(unsafe.Pointer(cevent))
	case CONTROLLERBUTTONDOWN, CONTROLLERBUTTONUP:
		return (*ControllerButtonEvent)(unsafe.Pointer(cevent))
	case CONTROLLERDEVICEADDED, CONTROLLERDEVICEREMOVED, CONTROLLERDEVICEREMAPPED:
		return (*ControllerDeviceEvent)(unsafe.Pointer(cevent))
	case AUDIODEVICEADDED, AUDIODEVICEREMOVED:
		return (*AudioDeviceEvent)(unsafe.Pointer(cevent))
	case FINGERMOTION, FINGERDOWN, FINGERUP:
		return (*TouchFingerEvent)(unsafe.Pointer(cevent))
	case MULTIGESTURE:
		return (*MultiGestureEvent)(unsafe.Pointer(cevent))
	case DOLLARGESTURE, DOLLARRECORD:
		return (*DollarGestureEvent)(unsafe.Pointer(cevent))
	case DROPFILE, DROPTEXT, DROPBEGIN, DROPCOMPLETE:
		e := (*tDropEvent)(unsafe.Pointer(cevent))
		event := DropEvent{Type: e.Type, Timestamp: e.Timestamp, File: C.GoString((*C.char)(e.File)), WindowID: e.WindowID}
		C.SDL_free(e.File)
		return &event
	case SENSORUPDATE:
		return (*SensorEvent)(unsafe.Pointer(cevent))
	case RENDER_TARGETS_RESET, RENDER_DEVICE_RESET:
		return (*RenderEvent)(unsafe.Pointer(cevent))
	case QUIT:
		return (*QuitEvent)(unsafe.Pointer(cevent))
	case CLIPBOARDUPDATE:
		return (*ClipboardEvent)(unsafe.Pointer(cevent))
	default:
		if cevent.Type >= USEREVENT {
			// all events beyond USEREVENT are UserEvents to be registered with RegisterEvents
			return (*UserEvent)(unsafe.Pointer(cevent))
		}
		return (*CommonEvent)(unsafe.Pointer(cevent))
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
func EventState(type_ uint32, state int) uint8 {
	return uint8(C.SDL_EventState(C.Uint32(type_), C.int(state)))
}

// GetEventState returns the current processing state of the specified event
// (https://wiki.libsdl.org/SDL_EventState)
func GetEventState(type_ uint32) uint8 {
	return uint8(C.SDL_EventState(C.Uint32(type_), QUERY))
}

// RegisterEvents allocates a set of user-defined events, and return the beginning event number for that set of events.
// (https://wiki.libsdl.org/SDL_RegisterEvents)
func RegisterEvents(numEvents int) uint32 {
	return uint32(C.SDL_RegisterEvents(C.int(numEvents)))
}
