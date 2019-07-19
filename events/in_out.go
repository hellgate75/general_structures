package events

import (
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
)

var (
	//Function that Do Notjhing on the Message Arrival, for Consumers that prefer istead use events.MessageBus channels, for message broadcasting
	DO_NOTHING_CONSUMER_FUNC func(EventMessage) error = func(e EventMessage) error { return nil }
)

// Consumer Type describes Event consumer interface. Cosumer is the final receiver for the message
// broker, that cosumer messages.
type EventConsumer interface {
	// Consume message, the event manager uses this method to broadcast the message
	// Parameters:
	//    message (events.EventMessage) Message for the Consumer
	Consume(message *EventMessage)
	// Consume message, the event broker uses this method to broadcast the message
	// Returns:
	//    []events.EventTarget Event specified targets
	GetTargets() []EventTarget
	// Register New Message Bus, Event Engine Communication
	// Parameters:
	//    bus (events.MessageBus) Communication Bus
	// Returns:
	//    error Any Error that can occur during computation
	RegisterBus(bus MessageBus) error
}

type __consumerStruct struct {
	buses        []MessageBus
	targets      []EventTarget
	consumerFunc func(EventMessage) error
}

func (cons *__consumerStruct) Consume(message *EventMessage) {
	defer func() {
		var err error = nil
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				err = itf.(error)
			} else {
				err = errors.New(fmt.Sprintf("%v", itf))
			}
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error processing message '%s'", err.Error()))
			}
		}
	}()
	errF := cons.consumerFunc(*message)
	if errF != nil {
		if logger != nil {
			logger.ErrorS(fmt.Sprintf("Error processing message '%s'", errF.Error()))
		}
	}
	for _, bus := range cons.buses {
		go func() {
			bus <- *message
		}()
	}
}
func (cons *__consumerStruct) GetTargets() []EventTarget {
	return cons.targets
}

func (cons *__consumerStruct) RegisterBus(bus MessageBus) error {
	var err error = nil
	defer func() {
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				err = itf.(error)
			} else {
				err = errors.New(fmt.Sprintf("%v", itf))
			}
		}
	}()
	cons.buses = append(cons.buses, bus)
	return err
}

// Create New Event Consumer
// Parameters:
//    targets ([]eventsEventTarget) Array of Event Targets
//    consumerFunc (func(EventMessage) error) Function that describe consume of arriving data
// Returns:
//    events.EventConsumer Event Consumer instance
func NewEventConsumer(targets []EventTarget, consumerFunc func(EventMessage) error) EventConsumer {
	return &__consumerStruct{
		consumerFunc: consumerFunc,
		targets:      targets,
		buses:        make([]MessageBus, 0),
	}
}

// Producer Type describes Event producer interface. Producer is the component that creates specific
// messages related to the Source and reporting message details.
type EventProducer interface {
	// Produce message, the event broker uses message for Consumers
	// Parameters:
	//    eventName (string) Event name that will be triggered from message broker
	//    message (events.Message) Message of the event
	// Returns:
	//    events.EventMessage Event message used from the message broker
	Produce(eventName string, message Message) EventMessage
}

type __producerStruct struct {
	producerFunc func(string, Message) EventMessage
}

func (prod *__producerStruct) Produce(eventName string, message Message) EventMessage {
	return prod.producerFunc(eventName, message)
}

// Create New Event Producer
// Parameters:
//    producerFunc (func(string, Message) EventMessage) Function that descrives production of new EventMessage
// Returns:
//    events.EventConsumer Event Consumer instance
func NewEventProducer(producerFunc func(string, Message) EventMessage) EventProducer {
	return &__producerStruct{
		producerFunc: producerFunc,
	}
}
