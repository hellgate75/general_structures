## package events // import "github.com/hellgate75/general_structures/events"


### VARIABLES

##### var (
##### 	//Function that Do Nothing on the Message Arrival, for Consumers that prefer istead use events.MessageBus channels, for message broadcasting
##### 	DO_NOTHING_CONSUMER_FUNC func(EventMessage) error
##### )


### FUNCTIONS

#### func InitLogger()
     Initialize package logger if not started


### TYPES

##### type EventBroker interface {
##### 	// Registers a message in the event broker instance
##### 	// Parameters:
##### 	//    message (events.Message) Message that should be sent to the message broker
##### 	// Returns:
##### 	//    error Any error that can occur during computation
##### 	RegisterMessage(eType MessageType) error
##### 	// Registers a Consumer into the Event Broker
##### 	// Parameters:
##### 	//    eventTarget (events.EventTarget) Event target used that is treated by the target
##### 	//    consumer (events.EventConsumer) Consumer used to receive messages
##### 	// Returns:
##### 	//    error Any error that can occur during computation
##### 	RegisterConsumer(eventTarget EventTarget, consumer EventConsumer) error
##### 	// Registers a message in the event broker instance
##### 	// Parameters:
##### 	//    eventName (string) Event name that will be registered at the message broker
##### 	//    eventTarget (events.EventTarget) Event target used to find the target Consumers
##### 	//    source (events.EventSource) Event source used to qualify message at the Consumers
##### 	//    message (events.Message) Message that should be sent to the message broker
##### 	//    arguments ([]events.Value) Message Arguments
##### 	// Returns:
##### 	//    error Any error that can occur during computation
##### 	FireEvent(eventName string, eventTarget EventTarget, source EventSource, message Message, arguments []Value) error
##### 	// Registers a message in the event broker instance
##### 	// Parameters:
##### 	//    eventName (string) Event name that will be registered at the message broker
##### 	//    eventTarget (events.EventTarget) Event target used to find the target Consumers
##### 	//    eventMessage (events.EventMessage) Event Message consumed by the Consumers
##### 	// Returns:
##### 	//    error Any error that can occur during computation
##### 	FireEventMessage(eventName string, eventTarget EventTarget, eventMessage EventMessage) error
##### 	// Recover Trigger related to the specified target
##### 	// Parameters:
##### 	//    eventTarget (events.EventTarget) Event target used that is treated by the target
##### 	// Returns:
##### 	//    (events.EventTrigger Trigger Element related to the specified target,
##### 	//     error Any error that can occur during computation)
##### 	GetTrigger(eventTarget EventTarget) EventTrigger
##### }
    Event Broker describes features used from the Event Engine

##### func NewEventBroker() EventBroker
      Create a New Event Registry.
#####       Returns:
         events.EventRegistry Just created EventRegistry instance

##### type EventConsumer interface {
##### 	// Consume message, the event manager uses this method to broadcast the message
##### 	// Parameters:
##### 	//    message (events.EventMessage) Message for the Consumer
##### 	Consume(message *EventMessage)
##### 	// Consume message, the event broker uses this method to broadcast the message
##### 	// Returns:
##### 	//    []events.EventTarget Event specified targets
##### 	GetTargets() []EventTarget
##### 	// Register New Message Bus, Event Engine Communication
##### 	// Parameters:
##### 	//    bus (events.MessageBus) Communication Bus
##### 	// Returns:
##### 	//    error Any Error that can occur during computation
##### 	RegisterBus(bus MessageBus) error
##### }
    Consumer Type describes Event consumer interface. Cosumer is the final
    receiver for the message broker, that cosumer messages.

##### func NewEventConsumer(targets []EventTarget, consumerFunc func(EventMessage) error) EventConsumer
    Create New Event Consumer
#####     Parameters:
       targets ([]eventsEventTarget) Array of Event Targets
        consumerFunc (func(EventMessage) error) Function that describe consume of arriving data
#####     Returns:
       events.EventConsumer Event Consumer instance

##### type EventMessage struct {
##### 	Message   Message
##### 	Arguments []Value
##### 	Source    EventSource
##### }
    Event Message Wrapper, describes Message Elements, Source, support values

##### func (em *EventMessage) Equals(m EventMessage) bool
    Equals method to match 2 instances or instances of 2 extensions of
    events.EventMessage
#####     Parameters:
       s (events.EventMessage) Event Message istance to be compared
#####     Returns:
       bool Equality response

##### func (em *EventMessage) String() string
    String representation of the element
#####     Returns:
       string String representation of the instance

##### type EventProducer interface {
##### 	// Produce message, the event broker uses message for Consumers
##### 	// Parameters:
##### 	//    eventName (string) Event name that will be triggered from message broker
##### 	//    message (events.Message) Message of the event
##### 	// Returns:
##### 	//    events.EventMessage Event message used from the message broker
##### 	Produce(eventName string, message Message) EventMessage
##### }
    Producer Type describes Event producer interface. Producer is the component
    that creates specific messages related to the Source and reporting message
    details.

##### func NewEventProducer(producerFunc func(string, Message) EventMessage) EventProducer
    Create New Event Producer
#####     Parameters:
       producerFunc (func(string, Message) EventMessage) Function that descrives production of new EventMessage
#####     Returns:
       events.EventConsumer Event Consumer instance

##### type EventRegistry interface {
##### 	// Registers a message into the event broker.
##### 	// Parameters:
##### 	//    eventName (string) Event name that will be registered at the message broker
##### 	// Returns:
##### 	//    events.EventRegistryClosure Event Registration Closure user interface
##### 	On(eventName string) EventRegistryClosure
##### 	// Get the list of Error Types related to an event name.
##### 	// Parameters:
##### 	//    eventName (string) Event name that will be registered at the message broker
##### 	// Returns:
##### 	//    ( []events.MessageType Array of MessageType registered on given event name,
##### 	//    error Any error that can occur during computation)
##### 	Get(eventName string) ([]MessageType, error)
##### }
    Registry Type describes Event registration interface. Registry registers
    events and actions into the message broker

##### func NewEventRegistry() EventRegistry
    Create a New Event Registry.
#####     Returns:
       events.EventRegistry Just created EventRegistry instance

##### type EventRegistryClosure interface {
##### 	// Registers a message into the event broker.
##### 	// Parameters:
##### 	//    eventFunc (events.TriggerFunc) Event function that will be registered at the message broker, with the specified message name
##### 	//    targets ([]events.EventTarget) List of available event targets
##### 	// Returns:
##### 	//    error Any error that can occur during computation
##### 	Do(eventFunc TriggerFunc, targets []EventTarget) error
##### }
    Registry Event store closure interface, used to complete registration of one
    event.

##### type EventSource struct {
##### 	Source     string
##### 	CallOrigin string
##### 	EventName  string
##### }
    Event Source Type, describes information about source

##### func (es *EventSource) Equals(s EventSource) bool
    Equals method to match 2 instances or instances of 2 extensions of
    events.EventSource
#####     Parameters:
       s (events.EventSource) Event Source istance to be compared
#####     Returns:
       bool Equality response

##### func (es *EventSource) String() string
    String representation of the element
#####     Returns:
       string String representation of the instance

##### type EventTarget struct {
##### 	AppCode   string
##### 	Target    string
##### 	Operation string
##### }
    Event Target Type, describes information about target

##### func (et *EventTarget) Equals(t EventTarget) bool
    Equals method to match 2 instances or instances of 2 extensions of
    events.EventTarget
#####     Parameters:
       t (events.EventTarget) Event Target istance to be compared
#####     Returns:
       bool Equality response

##### func (et *EventTarget) String() string
    String representation of the element
#####     Returns:
       string String representation of the instance

##### type EventTrigger interface {
##### 	// Triggers messages to the event broker
##### 	// Parameters:
##### 	//    eventName (string) Event name that will be registered at the message broker
##### 	//    source (events.EventSource) Event source used to qualify message at the Consumers
##### 	//    message (events.Message) Message that should be sent to the message broker
##### 	//    values (variadic events.Value array) Message  Values that should be sent to the message broker
##### 	// Returns:
##### 	//    error Any error that can occur during computation
##### 	Trigger(eventName string, source EventSource, message Message, values ...Value) error
##### }
    Trigger Type describes Event trigger interface. Trigger activate an action
    on an event, calling the message broker and triggering the event action

##### func NewEventTrigger(target EventTarget, broker EventBroker) EventTrigger
    Creates new event trigger element, linked to an event broker
#####     Parameters:
       eventTarget (events.EventTarget) Event target used that is treated by the target
       broker (*events.EventBroker) Event Broker instance
#####     Returns:
       events.EventTrigger Trigger Linked to the Event Broker

##### type Message interface{}
    Message Type

##### type MessageBus chan interface{}
    Message Channel

##### type MessageType struct {
##### 	EventName string
##### 	Targets   []EventTarget
##### 	Action    TriggerFunc
##### }
    Event Message Wrapper, describes Message Elements, Source, support values

##### func (mt *MessageType) Equals(m MessageType) bool
    Equals method to match 2 instances or instances of 2 extensions of
    events.MessageType
#####     Parameters:
       s (events.MessageType) Message Type instance to be compared
#####     Returns:
       bool Equality response

##### func (mt *MessageType) String() string
    String representation of the element
#####     Returns:
       string String representation of the instance

##### type TriggerFunc func(Message, []Value) error
    Message Channel

##### type Value interface{}
    Value Type

