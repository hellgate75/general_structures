package events

import (
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
)

// Trigger Type describes Event trigger interface. Trigger activate an action on an event,
// calling the message broker and triggering the event action
type EventTrigger interface {
	// Triggers messages to the event broker
	// Parameters:
	//    eventName (string) Event name that will be registered at the message broker
	//    source (events.EventSource) Event source used to qualify message at the Consumers
	//    message (events.Message) Message that should be sent to the message broker
	//    values (variadic events.Value array) Message  Values that should be sent to the message broker
	// Returns:
	//    error Any error that can occur during computation
	Trigger(eventName string, source EventSource, message Message, values ...Value) error
}

type __triggerStruct struct {
	target EventTarget
	broker EventBroker
}

func (et *__triggerStruct) Trigger(eventName string, source EventSource, message Message, values ...Value) error {
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
	et.broker.FireEvent(eventName, et.target, source, message, values)
	return err
}

// Creates new event trigger element, linked to an event broker
// Parameters:
//    eventTarget (events.EventTarget) Event target used that is treated by the target
//    broker (*events.EventBroker) Event Broker instance
// Returns:
//    events.EventTrigger Trigger Linked to the Event Broker
func NewEventTrigger(target EventTarget, broker EventBroker) EventTrigger {
	return &__triggerStruct{
		target: target,
		broker: broker,
	}
}

// Event Broker describes features used from the Event Engine
type EventBroker interface {
	// Registers a message in the event broker instance
	// Parameters:
	//    message (events.Message) Message that should be sent to the message broker
	// Returns:
	//    error Any error that can occur during computation
	RegisterMessage(eType MessageType) error
	// Registers a Consumer into the Event Broker
	// Parameters:
	//    eventTarget (events.EventTarget) Event target used that is treated by the target
	//    consumer (events.EventConsumer) Consumer used to receive messages
	// Returns:
	//    error Any error that can occur during computation
	RegisterConsumer(eventTarget EventTarget, consumer EventConsumer) error
	// Registers a message in the event broker instance
	// Parameters:
	//    eventName (string) Event name that will be registered at the message broker
	//    eventTarget (events.EventTarget) Event target used to find the target Consumers
	//    source (events.EventSource) Event source used to qualify message at the Consumers
	//    message (events.Message) Message that should be sent to the message broker
	//    arguments ([]events.Value) Message Arguments
	// Returns:
	//    error Any error that can occur during computation
	FireEvent(eventName string, eventTarget EventTarget, source EventSource, message Message, arguments []Value) error
	// Registers a message in the event broker instance
	// Parameters:
	//    eventName (string) Event name that will be registered at the message broker
	//    eventTarget (events.EventTarget) Event target used to find the target Consumers
	//    eventMessage (events.EventMessage) Event Message consumed by the Consumers
	// Returns:
	//    error Any error that can occur during computation
	FireEventMessage(eventName string, eventTarget EventTarget, eventMessage EventMessage) error
	// Recover Trigger related to the specified target
	// Parameters:
	//    eventTarget (events.EventTarget) Event target used that is treated by the target
	// Returns:
	//    (events.EventTrigger Trigger Element related to the specified target,
	//     error Any error that can occur during computation)
	GetTrigger(eventTarget EventTarget) EventTrigger
}

func __encodeEventTarget(t EventTarget) string {
	return fmt.Sprintf("%s::%s::%s", t.AppCode, t.Target, t.Operation)
}

type __eventBrokerStruct struct {
	//Encoded Event Target and EventConsumer list (for any event target overlapping case) map
	consumersMap map[string][]EventConsumer
	//Encoded Event Target and MessageType list (for any event name overlapping case) map
	messagesMap map[string][]MessageType
	//event name and MessageType list (for any event name overlapping case) map
	messagesEventMap map[string][]MessageType
}

func (broker *__eventBrokerStruct) RegisterMessage(eType MessageType) error {
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
	for i := 0; i < len(eType.Targets); i++ {
		var key string = __encodeEventTarget(eType.Targets[i])
		mList, ok := broker.messagesMap[key]
		if !ok {
			mList = make([]MessageType, 0)
		}
		mList = append(mList, eType)
		broker.messagesMap[key] = mList
	}
	var eventKey string = eType.EventName
	mList, ok := broker.messagesEventMap[eventKey]
	if !ok {
		mList = make([]MessageType, 0)
	}
	mList = append(mList, eType)
	broker.messagesEventMap[eventKey] = mList
	return err
}

func (broker *__eventBrokerStruct) RegisterConsumer(eventTarget EventTarget, consumer EventConsumer) error {
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
	var key string = __encodeEventTarget(eventTarget)
	mList, ok := broker.consumersMap[key]
	if !ok {
		mList = make([]EventConsumer, 0)
	}
	mList = append(mList, consumer)
	broker.consumersMap[key] = mList
	return err
}

func (broker *__eventBrokerStruct) FireEventMessage(eventName string, eventTarget EventTarget, eventMessage EventMessage) error {
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
	var key string = __encodeEventTarget(eventTarget)
	mList, ok := broker.consumersMap[key]
	if !ok {
		return errors.New(fmt.Sprintf("Unable to find the specified target : %v", eventTarget))
	}
	for _, consumer := range mList {
		go func() {
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
			consumer.Consume(&eventMessage)
		}()
	}
	return err
}
func (broker *__eventBrokerStruct) FireEvent(eventName string, eventTarget EventTarget, source EventSource, message Message, arguments []Value) error {
	eventMessage := EventMessage{
		Source:    source,
		Message:   message,
		Arguments: arguments,
	}
	return broker.FireEventMessage(eventName, eventTarget, eventMessage)
}

func (broker *__eventBrokerStruct) GetTrigger(eventTarget EventTarget) EventTrigger {
	return NewEventTrigger(eventTarget, EventBroker(broker))

}

// Creates new event Event Broker
// Returns:
//    events.EventBroker Event Broker instance
func NewEventBroker() EventBroker {
	return &__eventBrokerStruct{
		consumersMap:     make(map[string][]EventConsumer),
		messagesMap:      make(map[string][]MessageType),
		messagesEventMap: make(map[string][]MessageType),
	}
}
