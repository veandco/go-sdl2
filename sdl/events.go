package sdl

// #include <SDL2/SDL.h>
import "C"
import "unsafe"
import "reflect"
import "fmt"

const (
	FIRSTEVENT              = 0
	QUIT                    = 0x100
	APP_TERMINATING         = 0x101
	APP_LOWMEMORY           = 0x102
	APP_WILLENTERBACKGROUND = 0x103
	APP_DIDENTERBACKGROUND  = 0x104
	APP_WILLENTERFOREGROUND = 0x105
	APP_DIDENTERFOREGROUND  = 0x106

	/* Window events */
	WINDOWEVENT = 0x200
	SYSWMEVENT  = 0x201

	/* Keyboard events */
	KEYDOWN     = 0x300
	KEYUP       = 0x301
	TEXTEDITING = 0x302
	TEXTINPUT   = 0x303

	/* Mouse events */
	MOUSEMOTION     = 0x400
	MOUSEBUTTONDOWN = 0x401
	MOUSEBUTTONUP   = 0x402
	MOUSEWHEEL      = 0x403

	/* Joystick events */
	JOYAXISMOTION    = 0x600
	JOYBALLMOTION    = 0x601
	JOYHATMOTION     = 0x602
	JOYBUTTONDOWN    = 0x603
	JOYBUTTONUP      = 0x604
	JOYDEVICEADDED   = 0x605
	JOYDEVICEREMOVED = 0x606

	/* Game controller events */
	CONTROLLERAXISMOTION     = 0x650
	CONTROLLERBUTTONDOWN     = 0x651
	CONTROLLERBUTTONUP       = 0x652
	CONTROLLERDEVICEADDED    = 0x653
	CONTROLLERDEVICEREMOVED  = 0x654
	CONTROLLERDEVICEREMAPPED = 0x655

	/* Touch events */
	FINGERDOWN   = 0x700
	FINGERUP     = 0x701
	FINGERMOTION = 0x702

	/* Gesture events */
	DOLLARGESTURE = 0x800
	DOLLARRECORD  = 0x801
	MULTIGESTURE  = 0x802

	/* Clipboard events */
	CLIPBOARDUPDATE = 0x900

	/* Drag and drop events */
	DROPFILE = 0x1000

	USEREVENT = 0x8000
	LASTEVENT = 0xFFFF
)

const (
	ADDEVENT = iota
	PEEKEVENT
	GETEVENT
)

const (
	QUERY   = -1
	IGNORE  = 0
	DISABLE = 0
	ENABLE  = 1
)

type Event interface{}

type CEvent struct {
	Type     uint32
	padding1 [52]byte
}

type CommonEvent struct {
	Type      uint32
	Timestamp uint32
}

type WindowEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Event     uint8
	padding1  uint8
	padding2  uint8
	padding3  uint8
	Data1     int32
	Data2     int32
}

type KeyDownEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	State     uint8
	Repeat    uint8
	padding1  uint8
	padding2  uint8
	Keysym    Keysym
}

type KeyUpEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	State     uint8
	Repeat    uint8
	padding1  uint8
	padding2  uint8
	Keysym    Keysym
}

type TextEditingEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Text      [C.SDL_TEXTINPUTEVENT_TEXT_SIZE]byte
	Start     int32
	Length    int32
}

type TextInputEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Text      [C.SDL_TEXTINPUTEVENT_TEXT_SIZE]byte
}

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

type MouseButtonEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	Button    uint8
	State     uint8
	padding1  uint8
	padding2  uint8
	X         int32
	Y         int32
}

type MouseWheelEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Which     uint32
	X         int32
	Y         int32
}

type JoyAxisEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
	Axis      uint8
	padding1  uint8
	padding2  uint8
	padding3  uint8
	Value     int16
	padding4  uint16
}

type JoyBallEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
	Ball      uint8
	padding1  uint8
	padding2  uint8
	padding3  uint8
	XRel      int16
	YRel      int16
}

type JoyHatEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
	Hat       uint8
	Value     uint8
	padding1  uint8
	padding2  uint8
}

type JoyButtonEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
	Button    uint8
	State     uint8
	padding1  uint8
	padding2  uint8
}

type JoyDeviceEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
}

type ControllerAxisEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
	Axis      uint8
	padding1  uint8
	padding2  uint8
	padding3  uint8
	Value     int16
	padding4  uint16
}

type ControllerButtonEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
	Button    uint8
	State     uint8
	padding1  uint8
	padding2  uint8
}

type ControllerDeviceEvent struct {
	Type      uint32
	Timestamp uint32
	Which     JoystickID
}

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

type MultiGestureEvent struct {
	Type       uint32
	Timestamp  uint32
	TouchId    TouchID
	DTheta     float32
	DDist      float32
	X          float32
	Y          float32
	NumFingers uint16
	padding    uint16
}

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

type DropEvent struct {
	Type      uint32
	Timestamp uint32
	file      unsafe.Pointer
}

type QuitEvent struct {
	Type      uint32
	Timestamp uint32
}

type OSEvent struct {
	Type      uint32
	Timestamp uint32
}

type ClipboardEvent struct {
	Type      uint32
	Timestamp uint32
}

type UserEvent struct {
	Type      uint32
	Timestamp uint32
	WindowID  uint32
	Code      int32
	Data1     unsafe.Pointer
	Data2     unsafe.Pointer
}

type SysWMEvent struct {
	Type      uint32
	Timestamp uint32
	msg       unsafe.Pointer
}

type EventFilter C.SDL_EventFilter
type EventAction C.SDL_eventaction

func (action EventAction) c() C.SDL_eventaction {
    return C.SDL_eventaction(action)
}

func PumpEvents() {
	C.SDL_PumpEvents()
}

func PeepEvents(events []Event, action EventAction, minType, maxType uint32) int {
	var _events []CEvent = make([]CEvent, len(events))

	if action == ADDEVENT { // the contents of _events matter if they are to be added
		for i := 0; i < len(events); i++ {
			_events[i] = *cEvent(events[i])
		}
	}

	_pevents := (*C.SDL_Event)(unsafe.Pointer(&_events[0]))
	retVal := int(C.SDL_PeepEvents(_pevents, C.int(len(events)), action.c(), C.Uint32(minType), C.Uint32(maxType)))

	if action != ADDEVENT { // put events into slice, events unchanged if action = ADDEVENT
		for i := 0; i < retVal; i++ {
			events[i] = goEvent(&_events[i])
		}
	}
	return retVal
}

func HasEvent(type_ uint32) bool {
	return C.SDL_HasEvent(C.Uint32(type_)) != 0
}

func HasEvents(minType, maxType uint32) bool {
	return C.SDL_HasEvents(C.Uint32(minType), C.Uint32(maxType)) != 0
}

func FlushEvent(type_ uint32) {
	C.SDL_FlushEvent(C.Uint32(type_))
}

func FlushEvents(minType, maxType uint32) {
	C.SDL_FlushEvents(C.Uint32(minType), C.Uint32(maxType))
}

func PollEvent() Event {
	var cevent C.SDL_Event
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
	case QUIT:
		return (*QuitEvent)(unsafe.Pointer(cevent))
	case USEREVENT:
		return (*UserEvent)(unsafe.Pointer(cevent))
	case CLIPBOARDUPDATE:
		return (*ClipboardEvent)(unsafe.Pointer(cevent))
	}

	panic(fmt.Errorf("Unknown event type: %v", cevent.Type))
}

func cEvent(event Event) *CEvent {
	evv := reflect.ValueOf(event)
	p := evv.Elem()
	return (*CEvent)(unsafe.Pointer(p.UnsafeAddr()))
}

func WaitEventTimeout(timeout int) Event {
	var cevent CEvent
	_event := (*C.SDL_Event)(unsafe.Pointer(&cevent))
	ok := int(C.SDL_WaitEventTimeout(_event, C.int(timeout)))
	if ok == 0 {
		return nil
	}
	return goEvent(&cevent)
}

func WaitEvent() Event {
	var cevent CEvent
	_event := (*C.SDL_Event)(unsafe.Pointer(&cevent))
	ok := int(C.SDL_WaitEvent(_event))
	if ok == 0 {
		return nil
	}
	return goEvent(&cevent)
}

func PushEvent(event Event) int {
	_event := (*C.SDL_Event)(unsafe.Pointer(cEvent(event)))
	return int(C.SDL_PushEvent(_event))
}

/* TODO: implement SDL_EventFilter functions */

func EventState(type_ uint32, state int) uint8 {
	return uint8(C.SDL_EventState(C.Uint32(type_), C.int(state)))
}

func GetEventState(type_ uint32) uint8 {
	return uint8(C.SDL_EventState(C.Uint32(type_), QUERY))
}

func RegisterEvents(numEvents int) uint32 {
	return uint32(C.SDL_RegisterEvents(C.int(numEvents)))
}
