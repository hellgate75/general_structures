package events

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestNewEventTrigger(t *testing.T) {
	var eventName string = "Event1"
	var source EventSource = EventSource{
		CallOrigin: "Origin1",
		EventName:  eventName,
		Source:     "Source1",
	}
	var message Message = Message("Event Message")
	var values []Value = []Value{Value("Argq1"), Value("Argq2"), Value("Argq3")}
	var firedEventMessage EventMessage
	var target EventTarget = EventTarget{
		AppCode:   "App1",
		Target:    "Target1",
		Operation: "Oper1",
	}
	var broker EventBroker = NewEventBroker()
	var et EventTrigger = NewEventTrigger(target, broker)
	var targets []EventTarget = []EventTarget{target}
	var consumerFunc func(EventMessage) error = func(e EventMessage) error {
		firedEventMessage = e
		//		fmt.Println("Message arrived in functio!!")
		//		fmt.Println(messageFunc)
		return nil
	}
	var consumer EventConsumer = NewEventConsumer(targets, consumerFunc)
	broker.RegisterConsumer(target, consumer)
	et.Trigger(eventName, source, message, values)
	time.Sleep(1500 * time.Millisecond)
	if message != firedEventMessage.Message {
		t.Fatal(fmt.Sprintf("Error in Receive Trigger Call, Expected <%v> but Given <%v>", message, firedEventMessage.Message))
	}
}

func TestNewEventBroker(t *testing.T) {
	var broker EventBroker = NewEventBroker()
	var target EventTarget = EventTarget{
		AppCode:   "ScreenEcho",
		Operation: "ScreenRead",
		Target:    "LoadFile",
	}
	var target2 EventTarget = EventTarget{
		AppCode:   "ScreenEcho",
		Operation: "ScreenRead",
		Target:    "SaveFile",
	}
	var eType MessageType = MessageType{
		EventName: "OnClick",
		Targets:   []EventTarget{target, target2},
		Action: TriggerFunc(func(m Message, args []Value) error {
			return nil
		}),
	}
	broker.RegisterMessage(eType)
	var firedSave, firedLoad bool
	var message1, message2 Message
	var consumerFunc func(EventMessage) error = func(event EventMessage) error {
		if event.Source.Source == "SaveButton" {
			firedSave = true
			message1 = event.Message
			return nil
		} else if event.Source.Source == "LoadButton" {
			firedLoad = true
			message2 = event.Message
			return nil
		}
		return errors.New(fmt.Sprintf("Fired uncaught event source : <%s>", event.Source.Source))
	}
	expectedMessage1 := Message("Click On Save Button!!")
	expectedMessage2 := Message("Click On Load Button!!")
	var consumer EventConsumer = NewEventConsumer([]EventTarget{target, target2}, consumerFunc)
	broker.RegisterConsumer(target, consumer)
	broker.RegisterConsumer(target2, consumer)
	var saveTrigger EventTrigger = broker.GetTrigger(target)
	var loadTrigger EventTrigger = broker.GetTrigger(target2)
	var eventSource1 EventSource = EventSource{
		EventName:  "OnClick",
		CallOrigin: "SaveButton",
		Source:     "SaveButton",
	}
	var eventSource2 EventSource = EventSource{
		EventName:  "OnClick",
		CallOrigin: "LoadButton",
		Source:     "LoadButton",
	}
	saveTrigger.Trigger("OnClick", eventSource1, Message(expectedMessage1), []Value{})
	loadTrigger.Trigger("OnClick", eventSource2, Message(expectedMessage2), []Value{})
	time.Sleep(1500 * time.Millisecond)
	if !firedSave {
		t.Fatal("Event from Save Button didn't arrive to consumer!!")
	}
	if !firedLoad {
		t.Fatal("Event from Load Button didn't arrive to consumer!!")
	}
	if expectedMessage1 != message1 {
		t.Fatal(fmt.Sprintf("Error in Save Button Event Message, Expected <%v> but Given <%v>", expectedMessage1, message1))
	}
	if expectedMessage2 != message2 {
		t.Fatal(fmt.Sprintf("Error in Load Button Event Message, Expected <%v> but Given <%v>", expectedMessage2, message2))
	}
}
