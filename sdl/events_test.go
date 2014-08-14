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

func TestSetGetEventFilter(t *testing.T) {
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

func TestSetEventFilterFuncUserData(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	data := int(42)
	SetEventFilterFunc(func(userdata interface{}, e Event) bool {
		i, ok := userdata.(int)
		if !ok {
			t.Errorf("Failed to cast userdata to int")
		}
		if i != data {
			t.Errorf("Expected userdata to be %d but was %d", data, i)
		}

		return true
	}, data)

	in := UserEvent{Type: USEREVENT, Code: 42}
	PushEvent(&in)
	WaitEvent()
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

func TestSetEventFilter(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	filterFunc := func(_ interface{}, e Event) bool {
		ue, ok := e.(*UserEvent)
		if !ok {
			return false
		}

		return ue.Code == 42
	}
	SetEventFilterFunc(filterFunc, nil)

	ins := []*UserEvent{
		&UserEvent{Type: USEREVENT, Code: 42},
		&UserEvent{Type: USEREVENT, Code: 41},
	}

	expectedOutCount := 0
	for _, in := range ins {
		PushEvent(in)
		if filterFunc(nil, in) {
			expectedOutCount++
		}
	}

	outCount := countEventsInQ(false)

	if outCount != expectedOutCount {
		t.Errorf("Expected %d events to pass but got %d.", expectedOutCount, outCount)
	}
}

func TestGetEventFilterNilOnStartup(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	if GetEventFilter() != nil {
		t.Errorf("Event filter should be nil on startup.")
	}
	SetEventFilterFunc(func(_ interface{}, _ Event) bool {
		return true
	}, nil)

	Quit()
	Init(INIT_EVERYTHING)

	if GetEventFilter() != nil {
		t.Errorf("Event filter should be nil on startup.")
	}
}

func TestFilterEventsFuncQ(t *testing.T) {
	Init(INIT_EVERYTHING)
	defer Quit()

	ins := []*UserEvent{
		&UserEvent{Type: USEREVENT, Code: 42},
		&UserEvent{Type: USEREVENT, Code: 41},
	}
	data := int(42)

	filterFunc := func(userdata interface{}, e Event) bool {
		i, ok := userdata.(int)
		if !ok {
			t.Errorf("Failed to cast userdata to int")
			return false
		}
		if i != data {
			t.Errorf("Expected userdata to be %d but was %d", data, i)
			return false
		}

		ue, ok := e.(*UserEvent)
		if !ok {
			t.Errorf("Failed to cast event to *UserEvent")
			return false
		}

		return ue.Code == int32(i)
	}

	expectedOutCount := 0
	for _, in := range ins {
		PushEvent(in)
		if filterFunc(data, in) {
			expectedOutCount++
		}
	}

	FilterEventsFunc(filterFunc, data)

	outCount := countEventsInQ(false)

	if outCount != expectedOutCount {
		t.Errorf("Expected %d events to pass but got %d.", expectedOutCount, outCount)
	}
}
