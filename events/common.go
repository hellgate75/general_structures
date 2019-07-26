package events

import (
	"github.com/hellgate75/general_utils/log"
)

var logger log.Logger

//Initialize package logger if not started
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
