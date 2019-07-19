package events

import (
	"fmt"
	//	"github.com/hellgate75/general_structures/types"
	"strings"
)

// Event Target Type, describes information about target
type EventTarget struct {
	AppCode   string
	Target    string
	Operation string
}

// Equals method to match 2 instances or instances of 2 extensions of events.EventTarget
// Parameters:
//    t (events.EventTarget) Event Target istance to be compared
// Returns:
//    bool Equality response
func (et *EventTarget) Equals(t EventTarget) bool {
	return et.AppCode == t.AppCode && et.Target == t.Target &&
		et.Operation == t.Operation
}

// String representation of the element
// Returns:
//    string String representation of the instance
func (et *EventTarget) String() string {
	return fmt.Sprintf("EventTarget{AppCode='%s', Target='%s', Operation='%s'}", et.AppCode, et.Target, et.Operation)
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
	return fmt.Sprintf("EventMessage{Source='%s', Message='%s', Arguments=[%s]}", em.Source.String(), em.Message, strings.Join(args, ", "))
}

// Event Message Wrapper, describes Message Elements, Source, support values
type MessageType struct {
	EventName string
	Targets   []EventTarget
	Action    TriggerFunc
}

// Equals method to match 2 instances or instances of 2 extensions of events.MessageType
// Parameters:
//    s (events.MessageType) Message Type istance to be compared
// Returns:
//    bool Equality response
func (mt *MessageType) Equals(m MessageType) bool {
	if len(mt.Targets) != len(m.Targets) {
		return false
	}
	for i := 0; i < len(mt.Targets); i++ {
		if !mt.Targets[i].Equals(m.Targets[i]) {
			return false
		}
	}
	return mt.EventName == m.EventName
}

// String representation of the element
// Returns:
//    string String representation of the instance
func (mt *MessageType) String() string {
	var args []string = make([]string, len(mt.Targets))
	for i := 0; i < len(mt.Targets); i++ {
		args[i] = fmt.Sprintf("%s", mt.Targets[i].String())
	}
	return fmt.Sprintf("EventMessage{EventName='%s', Action='%v', Targets=[%s]}", mt.EventName, mt.Action, strings.Join(args, ", "))
}
