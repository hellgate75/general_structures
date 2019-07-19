package events

import (
	"fmt"
	"testing"
)

func TestNewInnerEventRegistryClosure(t *testing.T) {
	eventName := "Event1"
	structure := __registryStruct{
		registry: make(map[string][]MessageType),
	}
	ic := __newInnerEventRegistryClosure(eventName, &structure)
	ic.Do(TriggerFunc(func(m Message, a []Value) error { return nil }), []EventTarget{})
	if _, ok := structure.registry[eventName]; !ok {
		t.Fatal("Uncaught registration of a new MessageType!!")
	}
}

func TestNewEventRegistry(t *testing.T) {
	eventName := "Event1"
	var ir EventRegistry = NewEventRegistry()
	ir.On(eventName).Do(TriggerFunc(func(m Message, a []Value) error { return nil }), []EventTarget{})
	arr, err := ir.Get(eventName)
	if err != nil {
		t.Fatal(fmt.Sprintf("Uncaught error during events recovery, error is '%s'", err.Error()))
	}
	if len(arr) == 0 {
		t.Fatal(fmt.Sprintf("No element associated to event '%s'", eventName))
	}
	name := arr[0].EventName
	if eventName != name {
		t.Fatal(fmt.Sprintf("Wrong event name in the registry, Expected <%s> but Given <%s>", eventName, name))
	}
}
