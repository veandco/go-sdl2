package sdl

import (
	"testing"
)

func TestEventsPushEvent(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	in := UserEvent{
		Type: USEREVENT,
		Code: 42,
	}

	PushEvent(&in)

	out, ok := PollEvent().(*UserEvent)
	if !ok {
		t.Errorf("Failed to cast event to *UserEvent")
	}
	if out.Code != in.Code {
		t.Errorf("Expected event code %d but got %d", in.Code, out.Code)
	}
}

type simpleTestFilter struct{}

func (s *simpleTestFilter) FilterEvent(e Event) bool {
	return true
}

func TestEventsSetGetEventFilter(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	filter := &simpleTestFilter{}
	SetEventFilter(filter)

	if filter != GetEventFilter() {
		t.Errorf("Could not round-trip the event filter.")
	}

	if !isCEventFilterSet() {
		t.Errorf("Event filter was not actually set in C.")
	}

	SetEventFilter(nil)

	if nil != GetEventFilter() {
		t.Errorf("Event filter was not cleared.")
	}

	if isCEventFilterSet() {
		t.Errorf("Event filter was not actually cleared in C.")
	}
}

func countEventsInQ(wait bool) int {
	var e Event
	if wait {
		e = WaitEvent()
	} else {
		e = PollEvent()
	}

	count := 0
	for ; e != nil; e = PollEvent() {
		count++
	}
	return count
}

func TestEventsSetEventFilter(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	filterFunc := func(e Event, log bool) bool {
		if log {
			t.Log("TestSetEventFilter received", e)
		}

		ue, ok := e.(*UserEvent)
		if !ok {
			return false
		}

		return ue.Code == 42
	}
	SetEventFilterFunc(func(e Event) bool { return filterFunc(e, true) })

	ins := []*UserEvent{
		&UserEvent{Type: USEREVENT, Code: 42},
		&UserEvent{Type: USEREVENT, Code: 41},
	}

	expectedOutCount := 0
	for _, in := range ins {
		PushEvent(in)
		if filterFunc(in, false) {
			expectedOutCount++
		}
	}

	outCount := countEventsInQ(false)

	if outCount != expectedOutCount {
		t.Errorf("Expected %d events to pass but got %d.", expectedOutCount, outCount)
	}
}

func TestEventsGetEventFilterNilOnStartup(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	if GetEventFilter() != nil {
		t.Errorf("Event filter should be nil on startup.")
	}
	SetEventFilterFunc(func(_ Event) bool {
		return true
	})

	Quit()
	Init(INIT_EVERYTHING)

	if GetEventFilter() != nil {
		t.Errorf("Event filter should be nil on startup.")
	}
}

func TestEventsFilterEventsFuncQ(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	ins := []*UserEvent{
		&UserEvent{Type: USEREVENT, Code: 42},
		&UserEvent{Type: USEREVENT, Code: 41},
	}

	filterFunc := func(e Event, log bool) bool {
		if log {
			t.Log("TestEventsFilterEventsFuncQ received", e)
		}

		ue, ok := e.(*UserEvent)
		if !ok {
			return false
		}

		return ue.Code == 42
	}

	expectedOutCount := 0
	for _, in := range ins {
		PushEvent(in)
		if filterFunc(in, false) {
			expectedOutCount++
		}
	}

	FilterEventsFunc(func(e Event) bool { return filterFunc(e, true) })

	outCount := countEventsInQ(false)

	if outCount != expectedOutCount {
		t.Errorf("Expected %d events to pass but got %d.", expectedOutCount, outCount)
	}
}

func TestEventsAddEventWatch(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	in := UserEvent{
		Type: USEREVENT,
		Code: 42,
	}

	out := Event(nil)
	watch := func(e Event) bool {
		t.Log("TestAddEventWatch received", e)
		out = e
		return true
	}

	AddEventWatchFunc(watch)
	PushEvent(&in)

	if out == nil {
		t.Errorf("Event from event watch was nil but expected it to be non-nil.")
	}

	outue, ok := out.(*UserEvent)
	if !ok {
		t.Errorf("Failed to cast event to *UserEvent")
	}
	if outue.Code != in.Code {
		t.Errorf("Expected event code %d but got %d", in.Code, outue.Code)
	}
}

func TestEventsEventWatchClearOnStartup(t *testing.T) {
	Init(INIT_EVERYTHING)

	AddEventWatchFunc(func(_ Event) bool {
		return true
	})

	Quit()

	if len(eventWatchesCache) != 0 {
		t.Errorf("Expected go event watches cache to be cleared but it contains %d contexts", len(eventWatchesCache))
	}
}

func TestEventsAddDelEventWatch(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	in := UserEvent{
		Type: USEREVENT,
		Code: 42,
	}

	out := Event(nil)
	watch := func(e Event) bool {
		t.Log("TestAddDelEventWatch received", e)
		out = e
		return true
	}

	handle := AddEventWatchFunc(watch)
	DelEventWatch(handle)
	PushEvent(&in)

	if out != nil {
		t.Errorf("Event was received from event watch after it had been removed.")
	}
}
