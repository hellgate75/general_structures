package events

import (
	"fmt"
	"testing"
	"time"
)

func TestNewEventConsumer(t *testing.T) {
	var messageFunc EventMessage
	var messageBus EventMessage
	var targets []EventTarget = []EventTarget{
		EventTarget{
			AppCode:   "App1",
			Target:    "Target1",
			Operation: "Oper1",
		},
	}
	var consumerFunc func(EventMessage) error = func(e EventMessage) error {
		messageFunc = e
		//		fmt.Println("Message arrived in functio!!")
		//		fmt.Println(messageFunc)
		return nil
	}
	var consumer EventConsumer = NewEventConsumer(targets, consumerFunc)
	consTargets := consumer.GetTargets()
	var bus MessageBus = make(MessageBus)
	consumer.RegisterBus(bus)
	go func() {
		select {
		case messageItf := <-bus:
			messageBus = messageItf.(EventMessage)
			//			fmt.Println("Message arrived in bus!!")
			//			fmt.Println(messageBus)
		}
	}()
	var eventMessage EventMessage = EventMessage{
		Source: EventSource{
			CallOrigin: "Origin1",
			EventName:  "OnTest1",
			Source:     "Source1",
		},
		Message:   Message("Test Message"),
		Arguments: []Value{Value("Arg1"), Value("Arg2"), Value("Arg3")},
	}
	consumer.Consume(&eventMessage)
	time.Sleep(2500 * time.Millisecond)
	if len(targets) != len(consTargets) || !targets[0].Equals(consTargets[0]) {
		t.Fatal("Consumer targets doesn't match the expectancies!!!")
	}
	if !messageFunc.Equals(eventMessage) {
		t.Fatal(fmt.Sprintf("Message provided by function desn't match with sent one, Expected <%s> but Given <%s>", eventMessage.String(), messageFunc.String()))
	}
	if !messageBus.Equals(eventMessage) {
		t.Fatal(fmt.Sprintf("Message provided by bus desn't match with sent one, Expected <%s> but Given <%s>", eventMessage.String(), messageBus.String()))
	}
}

func TestNewEventProducer(t *testing.T) {
	var producerFunc func(string, Message) EventMessage = func(event string, message Message) EventMessage {
		return EventMessage{
			Source: EventSource{
				CallOrigin: "Origin1",
				EventName:  event,
				Source:     "Source1",
			},
			Message:   message,
			Arguments: []Value{Value("Arg1"), Value("Arg2"), Value("Arg3")},
		}

	}
	var producer EventProducer = NewEventProducer(producerFunc)
	eventName := "Event1"
	message := Message("Message1")
	var em EventMessage = producer.Produce(eventName, message)
	if eventName != em.Source.EventName {
		t.Fatal(fmt.Sprintf("Message preduced event name is Wrong, Expected <%s> but Given <%s>", eventName, em.Source.EventName))
	}
	if message != em.Message {
		t.Fatal(fmt.Sprintf("Message preduced message is Wrong, Expected <%s> but Given <%s>", message, em.Message))
	}

}
