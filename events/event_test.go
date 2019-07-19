package events

import (
	"fmt"
	"testing"
)

func TestEventTarget(t *testing.T) {
	target := EventTarget{
		AppCode:   "MyApp",
		Target:    "Event",
		Operation: "Oper",
	}
	target2 := EventTarget{
		AppCode:   "MyApp",
		Target:    "Event2",
		Operation: "Oper",
	}
	target3 := EventTarget{
		AppCode:   "MyApp",
		Target:    "Event",
		Operation: "Oper",
	}
	if !target.Equals(target3) {
		t.Fatal("Event Equals method in equality doeasn't work properly!!")
	}
	if target.Equals(target2) {
		t.Fatal("Event Equals method in disequality doeasn't work properly!!")
	}
	expectedValueStr := "EventTarget{AppCode='MyApp', Target='Event', Operation='Oper'}"
	valueStr := fmt.Sprintf("%s", target.String())
	if expectedValueStr != valueStr {
		t.Fatal(fmt.Sprintf("Wrong value for the String() call, Expected : <%s> but Given: <%s>!!", expectedValueStr, valueStr))
	}
}

func TestEventSource(t *testing.T) {
	target := EventSource{
		Source:     "Source1",
		EventName:  "Event1",
		CallOrigin: "Call1",
	}
	target2 := EventSource{
		Source:     "Source1",
		EventName:  "Event2",
		CallOrigin: "Call2",
	}
	target3 := EventSource{
		Source:     "Source1",
		EventName:  "Event1",
		CallOrigin: "Call1",
	}
	if !target.Equals(target3) {
		t.Fatal("Event Equals method in equality doeasn't work properly!!")
	}
	if target.Equals(target2) {
		t.Fatal("Event Equals method in disequality doeasn't work properly!!")
	}
	expectedValueStr := "EventSource{Source='Source1', EventName='Event1', CallOrigin='Call1'}"
	valueStr := fmt.Sprintf("%s", target.String())
	if expectedValueStr != valueStr {
		t.Fatal(fmt.Sprintf("Wrong value for the String() call, Expected : <%s> but Given: <%s>!!", expectedValueStr, valueStr))
	}
}

func TestEventMessage(t *testing.T) {
	target := EventMessage{
		Source: EventSource{
			Source:     "Source1",
			EventName:  "Event1",
			CallOrigin: "Call1",
		},
		Message:   "Message1",
		Arguments: []Value{"arg0", "arg1", "arg2"},
	}
	target2 := EventMessage{
		Source: EventSource{
			Source:     "Source2",
			EventName:  "Event2",
			CallOrigin: "Call2",
		},
		Message:   "Message1",
		Arguments: []Value{"arg0", "arg1", "arg2"},
	}
	target3 := EventMessage{
		Source: EventSource{
			Source:     "Source1",
			EventName:  "Event1",
			CallOrigin: "Call1",
		},
		Message:   "Message1",
		Arguments: []Value{"arg0", "arg1", "arg2"},
	}
	if !target.Equals(target3) {
		t.Fatal("Event Equals method in equality doeasn't work properly!!")
	}
	if target.Equals(target2) {
		t.Fatal("Event Equals method in disequality doeasn't work properly!!")
	}
	expectedValueStr := "EventMessage{Source='EventSource{Source='Source1', EventName='Event1', CallOrigin='Call1'}', Message='Message1', Arguments=[arg0, arg1, arg2]}"
	valueStr := fmt.Sprintf("%s", target.String())
	if expectedValueStr != valueStr {
		t.Fatal(fmt.Sprintf("Wrong value for the String() call, Expected : <%s> but Given: <%s>!!", expectedValueStr, valueStr))
	}
}

func TestMessageType(t *testing.T) {
	target := MessageType{
		EventName: "Event1",
		Action:    TriggerFunc(func(Message, []Value) error { return nil }),
		Targets: []EventTarget{
			EventTarget{
				AppCode:   "MyApp",
				Target:    "Event",
				Operation: "Oper",
			},
		},
	}
	target2 := MessageType{
		EventName: "Event1",
		Action:    TriggerFunc(func(Message, []Value) error { return nil }),
		Targets: []EventTarget{
			EventTarget{
				AppCode:   "MyApp",
				Target:    "Event2",
				Operation: "Oper",
			},
		},
	}
	target3 := MessageType{
		EventName: "Event1",
		Action:    TriggerFunc(func(Message, []Value) error { return nil }),
		Targets: []EventTarget{
			EventTarget{
				AppCode:   "MyApp",
				Target:    "Event",
				Operation: "Oper",
			},
		},
	}
	if !target.Equals(target3) {
		t.Fatal("Event Equals method in equality doeasn't work properly!!")
	}
	if target.Equals(target2) {
		t.Fatal("Event Equals method in disequality doeasn't work properly!!")
	}
	valueStr := fmt.Sprintf("%s", target.String())
	expectedValueStr := fmt.Sprintf("EventMessage{EventName='Event1', Action='%p', Targets=[EventTarget{AppCode='MyApp', Target='Event', Operation='Oper'}]}", target.Action)
	if expectedValueStr != valueStr {
		t.Fatal(fmt.Sprintf("Wrong value for the String() call, Expected : <%s> but Given: <%s>!!", expectedValueStr, valueStr))
	}
}
