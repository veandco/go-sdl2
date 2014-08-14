package sdl

import (
	"testing"
)

func TestPushEvent(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	in := UserEvent{
		Type: USEREVENT,
		Code: 42,
	}

	PushEvent(&in)

	out, ok := WaitEvent().(*UserEvent)
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

func TestEventFilterGetsSet(t *testing.T) {
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

func TestEventFilter(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	filterFunc := func(e Event) bool {
		ue, ok := e.(*UserEvent)
		if !ok {
			return false
		}

		return ue.Code == 42
	}
	SetEventFilterFunc(filterFunc)

	ins := []*UserEvent{
		&UserEvent{Type: USEREVENT, Code: 42},
		&UserEvent{Type: USEREVENT, Code: 41},
	}

	expectedOutCount := 0
	for _, in := range ins {
		PushEvent(in)
		if filterFunc(in) {
			expectedOutCount++
		}
	}

	outCount := 0
	for out := WaitEvent(); out != nil; out = PollEvent() {
		outCount++
	}

	if outCount != expectedOutCount {
		t.Errorf("Expected %d events to pass but got %d.", expectedOutCount, outCount)
	}
}

func TestEventFilterNilOnStartup(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	if GetEventFilter() != nil {
		t.Errorf("Event filter should be nil on startup.")
	}
	SetEventFilterFunc(func(e Event) bool {
		return true
	})

	Quit()
	Init(INIT_EVERYTHING)

	if GetEventFilter() != nil {
		t.Errorf("Event filter should be nil on startup.")
	}
}
