package events

import (
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
)

// Registry Event store closure interface, used to complete registration of one event.
type EventRegistryClosure interface {
	// Registers a message into the event broker.
	// Parameters:
	//    eventFunc (events.TriggerFunc) Event function that will be registered at the message broker, with the specified message name
	//    targets ([]events.EventTarget) List of available event targets
	// Returns:
	//    error Any error that can occur during computation
	Do(eventFunc TriggerFunc, targets []EventTarget) error
}

type __innerRegistryClosureStruct struct {
	self *__registryStruct
	name string
}

func (rc *__innerRegistryClosureStruct) Do(eventFunc TriggerFunc, targets []EventTarget) error {
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
	et := MessageType{
		EventName: rc.name,
		Targets:   targets,
		Action:    eventFunc,
	}
	if _, ok := rc.self.registry[rc.name]; !ok {
		rc.self.registry[rc.name] = make([]MessageType, 0)
	}
	rc.self.registry[rc.name] = append(rc.self.registry[rc.name], et)
	return err
}

func __newInnerEventRegistryClosure(eventName string, structure *__registryStruct) EventRegistryClosure {
	return &__innerRegistryClosureStruct{
		self: structure,
		name: eventName,
	}
}

// Registry Type describes Event registration interface. Registry registers events and actions into the message broker
type EventRegistry interface {
	// Registers a message into the event broker.
	// Parameters:
	//    eventName (string) Event name that will be registered at the message broker
	// Returns:
	//    events.EventRegistryClosure Event Registration Closure user interface
	On(eventName string) EventRegistryClosure
	// Get the list of Error Types related to an event name.
	// Parameters:
	//    eventName (string) Event name that will be registered at the message broker
	// Returns:
	//    ( []events.MessageType Array of MessageType registered on given event name,
	//    error Any error that can occur during computation)
	Get(eventName string) ([]MessageType, error)
}

type __registryStruct struct {
	registry map[string][]MessageType
}

func (r *__registryStruct) On(eventName string) EventRegistryClosure {
	return __newInnerEventRegistryClosure(eventName, r)
}

func (r *__registryStruct) Get(eventName string) ([]MessageType, error) {
	arr, ok := r.registry[eventName]
	if !ok {
		return make([]MessageType, 0), errors.New(fmt.Sprintf("Unable to find events associate to name : '%s'", eventName))
	}
	return arr, nil
}

// Create a New Event Registry.
// Returns:
//    events.EventRegistry Just created EventRegistry instance
func NewEventRegistry() EventRegistry {
	return &__registryStruct{
		registry: make(map[string][]MessageType),
	}
}
