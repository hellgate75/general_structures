package events

import (
	"fmt"
	"github.com/hellgate75/general_utils/log"
	"strings"
)

var logger log.Logger

func InitLogger() {
	currentLogger, err := log.New("events")
	if err != nil {
		panic(err.Error())
	}
	logger = currentLogger
}

// Message Type
type Message interface{}

// Value Type
type Value interface{}

// Message Channel
type MessageBus chan interface{}

// Message Channel
type TriggerFunc func(Message, []Value) error

// Event Target Type, describes information about target
type EventTarget struct {
	Target    string
	Operation string
}

// Equals method to match 2 instances or instances of 2 extensions of events.EventTarget
// Parameters:
//    t (events.EventTarget) Event Target istance to be compared
// Returns:
//    bool Equality response
func (et *EventTarget) Equals(t EventTarget) bool {
	return et.Target == t.Target
}

// String representation of the element
// Returns:
//    string String representation of the instance
func (et *EventTarget) String() string {
	return fmt.Sprintf("EventTarget{Target='%s', Operation='%s'}", et.Target, et.Operation)
}

// Event Source Type, describes information about source
type EventSource struct {
	Source     string
	CallOrigin string
	EventName  string
}

// Equals method to match 2 instances or instances of 2 extensions of events.EventSource
// Parameters:
//    s (events.EventSource) Event Source istance to be compared
// Returns:
//    bool Equality response
func (es *EventSource) Equals(s EventSource) bool {
	return es.EventName == s.EventName && es.Source == s.Source
}

// String representation of the element
// Returns:
//    string String representation of the instance
func (es *EventSource) String() string {
	return fmt.Sprintf("EventSource{Source='%s', EventName='%s', CallOrigin='%s'}", es.Source, es.EventName, es.CallOrigin)
}

// Event Message Wrapper, describes Message Elements, Source, support values
type EventMessage struct {
	Message   Message
	Arguments []Value
	Source    EventSource
}

// Equals method to match 2 instances or instances of 2 extensions of events.EventMessage
// Parameters:
//    s (events.EventMessage) Event Message istance to be compared
// Returns:
//    bool Equality response
func (em *EventMessage) Equals(m EventMessage) bool {
	var args []string = make([]string, len(em.Arguments))
	for i := 0; i < len(em.Arguments); i++ {
		args[i] = fmt.Sprintf("%v", em.Arguments[i])
	}
	var argsX []string = make([]string, len(m.Arguments))
	for i := 0; i < len(m.Arguments); i++ {
		argsX[i] = fmt.Sprintf("%v", m.Arguments[i])
	}
	return em.Source.Equals(m.Source) && em.Message == m.Message &&
		strings.Join(args, ",") == strings.Join(argsX, ",")
}

// String representation of the element
// Returns:
//    string String representation of the instance
func (em *EventMessage) String() string {
	var args []string = make([]string, len(em.Arguments))
	for i := 0; i < len(em.Arguments); i++ {
		args[i] = fmt.Sprintf("%v", em.Arguments[i])
	}
	return fmt.Sprintf("EventMessage{Source='%s', Message='%s', Arguments=[%s]}", em.Source, em.Message, strings.Join(args, ", "))
}

// Event Message Wrapper, describes Message Elements, Source, support values
type MessageType struct {
	EventName string
	Targets   []EventTarget
	Action    TriggerFunc
}

// Consumer Type describes Event consumer interface. Cosumer is the final receiver for the message
// broker, that cosumer messages.
type EventConsumer interface {
	// Consume message, the event manager uses this method to broadcast the message
	// Parameters:
	//    message (events.EventMessage) Message for the Consumer
	Consume(message EventMessage)
	// Consume message, the event broker uses this method to broadcast the message
	// Returns:
	//    []events.EventTarget Event specified targets
	GetTargets() []EventTarget
}

// Producer Type describes Event producer interface. Producer is the component that creates specific
// messages related to the Source and reporting message details.
type EventProducer interface {
	// Produce message, the event broker uses message for Consumers
	// Parameters:
	//    eventName (string) Event name that will be triggered from message broker
	// Returns:
	//    events.EventMessage Event message used from the message broker
	Produce(eventName string) EventMessage
}

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

// Registry Type describes Event registration interface. Registry registers events and actions into the message broker
type EventRegistry interface {
	// Registers a message into the event broker.
	// Parameters:
	//    eventName (string) Event name that will be registered at the message broker
	// Returns:
	//    events.EventRegistryClosure Event Registration Closure user interface
	On(eventName string) EventRegistryClosure
}

// Trigger Type describes Event trigger interface. Trigger activate an action on an event,
// calling the message broker and triggering the event action
type EventTrigger interface {
	// Triggers messages to the event broker
	// Parameters:
	//    eventName (string) Event name that will be registered at the message broker
	//    message (events.Message) Message that should be sent to the message broker
	//    values (variadic events.Value array) Message  Values that should be sent to the message broker
	// Returns:
	//    error Any error that can occur during computation
	Trigger(eventName string, message Message, values ...Value) error
}

// Event Broker describes features used from the Event Engine
type EventBroker interface {
	// Starts the event broker instance
	// Returns:
	//    error Any error that can occur during computation
	Start() error
	// Gets if the event broker instance is running
	// Returns:
	//    bool Event broker running state
	IsRunning() bool
	// Stops the event broker instance
	// Returns:
	//    error Any error that can occur during computation
	Stop() error
	// Destroys the event broker instance
	// Returns:
	//    error Any error that can occur during computation
	Destroy() error
	// Registers a message in the event broker instance
	// Parameters:
	//    message (events.Message) Message that should be sent to the message broker
	// Returns:
	//    error Any error that can occur during computation
	Register(eType MessageType) error
	// Registers a message in the event broker instance
	// Parameters:
	//    eventName (string) Event name that will be registered at the message broker
	//    eventTarget (events.EventTarget) Event target used to find the target Consumers
	//    message (events.Message) Message that should be sent to the message broker
	//    arguments ([]events.Value) Message Arguments
	// Returns:
	//    error Any error that can occur during computation
	FireEvent(eventName string, eventTarget EventTarget, message Message, arguments []Value) error
}
