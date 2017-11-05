package sdl

/*
#include "sdl_wrapper.h"
#include "events.h"

#if !SDL_VERSION_ATLEAST(2,0,2)
#define SDL_RENDER_TARGETS_RESET (0x2000)
#endif
*/
import "C"
import "unsafe"
import "reflect"
import "sync"

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

	// Window events
	WINDOWEVENT = C.SDL_WINDOWEVENT // window state change
	SYSWMEVENT  = C.SDL_SYSWMEVENT  // system specific event

	// Keyboard events
	KEYDOWN     = C.SDL_KEYDOWN     // key pressed
	KEYUP       = C.SDL_KEYUP       // key released
	TEXTEDITING = C.SDL_TEXTEDITING // keyboard text editing (composition)
	TEXTINPUT   = C.SDL_TEXTINPUT   // keyboard text input

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
	DROPFILE = C.SDL_DROPFILE // the system requests a file open

	// Render events
	RENDER_TARGETS_RESET = C.SDL_RENDER_TARGETS_RESET // the render targets have been reset and their contents need to be updated (>= SDL 2.0.2)

	// These are for your use, and should be allocated with RegisterEvents()
	USEREVENT = C.SDL_USEREVENT // a user-specified event
	LASTEVENT = C.SDL_LASTEVENT // only for bounding internal arrays
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
type Event interface{}

// CEvent is a union of all event structures used in SDL.
// (https://wiki.libsdl.org/SDL_Event)
type CEvent struct {
	Type uint32
	_    [52]byte // padding
}

// CommonEvent contains common event data.
// (https://wiki.libsdl.org/SDL_Event)
type CommonEvent struct {
	Type      uint32
	Timestamp uint32
}

// WindowEvent contains window state change event data.
// (https://wiki.libsdl.org/SDL_WindowEvent)
type WindowEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Event     uint8
	_         uint8 // padding
	_         uint8 // padding
	_         uint8 // padding
	Data1     int32
	Data2     int32
}
type cWindowEvent C.SDL_WindowEvent

//TODO: Key{Up,Down}Event only differ in 'Type' - a boolean field would be enough to distinguish up/down events

// KeyDownEvent contains keyboard key down event information.
// (https://wiki.libsdl.org/SDL_KeyboardEvent)
type KeyDownEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	State     uint8
	Repeat    uint8
	_         uint8 // padding
	_         uint8 // padding
	Keysym    Keysym
}
type cKeyboardEvent C.SDL_KeyboardEvent

// KeyUpEvent contains keyboard key up event information.
// (https://wiki.libsdl.org/SDL_KeyboardEvent)
type KeyUpEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	State     uint8
	Repeat    uint8
	_         uint8 // padding
	_         uint8 // padding
	Keysym    Keysym
}

// TextEditingEvent contains keyboard text editing event information.
// (https://wiki.libsdl.org/SDL_TextEditingEvent)
type TextEditingEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Text      [C.SDL_TEXTINPUTEVENT_TEXT_SIZE]byte
	Start     int32
	Length    int32
}
type cTextEditingEvent C.SDL_TextEditingEvent

// TextInputEvent contains keyboard text input event information.
// (https://wiki.libsdl.org/SDL_TextInputEvent)
type TextInputEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Text      [C.SDL_TEXTINPUTEVENT_TEXT_SIZE]byte
}
type cTextInputEvent C.SDL_TextInputEvent

// MouseMotionEvent contains mouse motion event information.
// (https://wiki.libsdl.org/SDL_MouseMotionEvent)
type MouseMotionEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	State     uint32
	X         int32
	Y         int32
	XRel      int32
	YRel      int32
}
type cMouseMotionEvent C.SDL_MouseMotionEvent

// MouseButtonEvent contains mouse button event information.
// (https://wiki.libsdl.org/SDL_MouseButtonEvent)
type MouseButtonEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	Button    uint8
	State     uint8
	_         uint8 // padding
	_         uint8 // padding
	X         int32
	Y         int32
}
type cMouseButtonEvent C.SDL_MouseButtonEvent

// MouseWheelEvent contains mouse wheel event information.
// (https://wiki.libsdl.org/SDL_MouseWheelEvent)
type MouseWheelEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	X         int32
	Y         int32
	Direction uint32
}
type cMouseWheelEvent C.SDL_MouseWheelEvent

// JoyAxisEvent contains joystick axis motion event information.
// (https://wiki.libsdl.org/SDL_JoyAxisEvent)
type JoyAxisEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
	Axis      uint8
	_         uint8 // padding
	_         uint8 // padding
	_         uint8 // padding
	Value     int16
	_         uint16 // padding
}
type cJoyAxisEvent C.SDL_JoyAxisEvent

// JoyBallEvent contains joystick trackball motion event information.
// (https://wiki.libsdl.org/SDL_JoyBallEvent)
type JoyBallEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
	Ball      uint8
	_         uint8 // padding
	_         uint8 // padding
	_         uint8 // padding
	XRel      int16
	YRel      int16
}
type cJoyBallEvent C.SDL_JoyBallEvent

// JoyHatEvent contains joystick hat position change event information.
// (https://wiki.libsdl.org/SDL_JoyHatEvent)
type JoyHatEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
	Hat       uint8
	Value     uint8
	_         uint8 // padding
	_         uint8 // padding
}
type cJoyHatEvent C.SDL_JoyHatEvent

// JoyButtonEvent contains joystick button event information.
// (https://wiki.libsdl.org/SDL_JoyButtonEvent)
type JoyButtonEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
	Button    uint8
	State     uint8
	_         uint8 // padding
	_         uint8 // padding
}
type cJoyButtonEvent C.SDL_JoyButtonEvent

// JoyDeviceEvent contains joystick device event information.
// (https://wiki.libsdl.org/SDL_JoyDeviceEvent)
type JoyDeviceEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
}

// ControllerAxisEvent contains game controller axis motion event information.
// (https://wiki.libsdl.org/SDL_ControllerAxisEvent)
type ControllerAxisEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
	Axis      uint8
	_         uint8 // padding
	_         uint8 // padding
	_         uint8
	Value     int16
	_         uint16 // padding
}
type cControllerAxisEvent C.SDL_ControllerAxisEvent

// ControllerButtonEvent contains game controller button event information.
// (https://wiki.libsdl.org/SDL_ControllerButtonEvent)
type ControllerButtonEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
	Button    uint8
	State     uint8
	_         uint8 // padding
	_         uint8 // padding
}
type cControllerButtonEvent C.SDL_ControllerButtonEvent

// ControllerDeviceEvent contains controller device event information.
// (https://wiki.libsdl.org/SDL_ControllerDeviceEvent)
type ControllerDeviceEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
}
type cControllerDeviceEvent C.SDL_ControllerDeviceEvent

// TouchFingerEvent contains finger touch event information.
// (https://wiki.libsdl.org/SDL_TouchFingerEvent)
type TouchFingerEvent struct {
	Type      uint32
	Timestamp uint32
	TouchID   TouchID
	FingerID  FingerID
	X         float32
	Y         float32
	DX        float32
	DY        float32
	Pressure  float32
}
type cTouchFingerEvent C.SDL_TouchFingerEvent

// MultiGestureEvent contains multiple finger gesture event information.
// (https://wiki.libsdl.org/SDL_MultiGestureEvent)
type MultiGestureEvent struct {
	Type       uint32
	Timestamp  uint32
	TouchID    TouchID
	DTheta     float32
	DDist      float32
	X          float32
	Y          float32
	NumFingers uint16
	_          uint16 // padding
}
type cMultiGestureEvent C.SDL_MultiGestureEvent

// DollarGestureEvent contains complex gesture event information.
// (https://wiki.libsdl.org/SDL_DollarGestureEvent)
type DollarGestureEvent struct {
	Type       uint32
	Timestamp  uint32
	TouchID    TouchID
	GestureID  GestureID
	NumFingers uint32
	Error      float32
	X          float32
	Y          float32
}
type cDollarGestureEvent C.SDL_DollarGestureEvent

// DropEvent contains an event used to request a file open by the system.
// (https://wiki.libsdl.org/SDL_DropEvent)
type DropEvent struct {
	Type      uint32
	Timestamp uint32
	File      unsafe.Pointer
	WindowID  uint32
}
type cDropEvent C.SDL_DropEvent

// RenderEvent contains render event information.
// (https://wiki.libsdl.org/SDL_EventType)
type RenderEvent struct {
	Type      uint32
	Timestamp uint32
}

// QuitEvent contains the "quit requested" event.
// (https://wiki.libsdl.org/SDL_QuitEvent)
type QuitEvent struct {
	Type      uint32
	Timestamp uint32
}

// OSEvent contains OS specific event information.
type OSEvent struct {
	Type      uint32
	Timestamp uint32
}

// ClipboardEvent contains clipboard event information.
// (https://wiki.libsdl.org/SDL_EventType)
type ClipboardEvent struct {
	Type      uint32
	Timestamp uint32
}

// UserEvent contains an application-defined event type.
// (https://wiki.libsdl.org/SDL_UserEvent)
type UserEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Code      int32
	Data1     unsafe.Pointer
	Data2     unsafe.Pointer
}
type cUserEvent C.SDL_UserEvent

// SysWMEvent contains a video driver dependent system event.
// (https://wiki.libsdl.org/SDL_SysWMEvent)
type SysWMEvent struct {
	Type      uint32
	Timestamp uint32
	msg       unsafe.Pointer
}
type cSysWMEvent C.SDL_SysWMEvent

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
	ret := C.SDL_PollEvent(&cevent)
	if ret == 0 {
		return nil
	}
	return goEvent((*CEvent)(unsafe.Pointer(&cevent)))
}

func goEvent(cevent *CEvent) Event {
	switch cevent.Type {
	case WINDOWEVENT:
		return (*WindowEvent)(unsafe.Pointer(cevent))
	case SYSWMEVENT:
		return (*SysWMEvent)(unsafe.Pointer(cevent))
	case KEYDOWN:
		return (*KeyDownEvent)(unsafe.Pointer(cevent))
	case KEYUP:
		return (*KeyUpEvent)(unsafe.Pointer(cevent))
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
	case JOYDEVICEADDED, JOYDEVICEREMOVED:
		return (*JoyDeviceEvent)(unsafe.Pointer(cevent))
	case CONTROLLERAXISMOTION:
		return (*ControllerAxisEvent)(unsafe.Pointer(cevent))
	case CONTROLLERBUTTONDOWN, CONTROLLERBUTTONUP:
		return (*ControllerButtonEvent)(unsafe.Pointer(cevent))
	case CONTROLLERDEVICEADDED, CONTROLLERDEVICEREMOVED, CONTROLLERDEVICEREMAPPED:
		return (*ControllerDeviceEvent)(unsafe.Pointer(cevent))
	case FINGERDOWN, FINGERUP, FINGERMOTION:
		return (*TouchFingerEvent)(unsafe.Pointer(cevent))
	case DOLLARGESTURE, DOLLARRECORD:
		return (*DollarGestureEvent)(unsafe.Pointer(cevent))
	case MULTIGESTURE:
		return (*MultiGestureEvent)(unsafe.Pointer(cevent))
	case DROPFILE:
		return (*DropEvent)(unsafe.Pointer(cevent))
	case RENDER_TARGETS_RESET:
		return (*RenderEvent)(unsafe.Pointer(cevent))
	case QUIT:
		return (*QuitEvent)(unsafe.Pointer(cevent))
	case USEREVENT:
		return (*UserEvent)(unsafe.Pointer(cevent))
	case CLIPBOARDUPDATE:
		return (*ClipboardEvent)(unsafe.Pointer(cevent))
	default:
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

// WaitEvent swait indefinitely for the next available event.
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
