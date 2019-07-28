package logic

import (
	"errors"
	"fmt"
	"github.com/hellgate75/general_structures/types"
	errs "github.com/hellgate75/general_utils/errors"
	"time"
)

var BUSINESSOBJECT_BUS_TIMEOUT time.Duration = 30 * time.Second
var CONSUMERS_MAX_WAITING_TIMEOUT time.Duration = 5 * time.Minute
var CONSUMERS_SUSPEND_TIMEOUT time.Duration = 10 * time.Second

// Describes Element that mask the connection with consumers and provides multiple consume strategies : Consumed by Only First, Consumed by All.
// Any Consumer can be registered on the Service Bus, and automatically will dispatch Business Object through first available or all consumers, when ready.
// Persistency strategy will depend on the Service Buss Implementation.
type ServiceBus interface {
	// Offers a business object only at first available consumer
	// Parameters:
	//    bo (logic.BusinessObject) Object to be consumed
	// Retruns:
	//    error Any error that occurs during the computation or nil
	OfferFirst(bo BusinessObject) error
	// Offers first communication channel
	// Retruns:
	//    *BOChannel Pointer to Business Object Channel
	OfferFirstChannel() *BOChannel
	// Offers a business object to any available consumer, and retrieve to not available accordingly to consumer strategy implementation
	// Parameters:
	//    bo (logic.BusinessObject) Object to be consumed
	// Retruns:
	//    error Any error that occurs during the computation or nil
	Offer(bo BusinessObject) error
	// Offers all communication channel
	// Retruns:
	//    *BOChannel Pointer to Business Object Channel
	OfferChannel() *BOChannel
	// Register a new Consumer into the Service Bus
	// Parameters:
	//    consumer (logic.BusinessConsumer) Consumer to be registered
	// Retruns:
	//    bool Registration success state
	RegisterConsumer(consumer BusinessConsumer) bool
	// Remove an evistingConsumer from the Service Bus
	// Parameters:
	//    consumer (logic.BusinessConsumer) Consumer to be removed
	// Retruns:
	//    bool Removal success state
	RemoveConsumer(consumer BusinessConsumer) bool
	// Remove an evistingConsumer from the Service Bus
	// Parameters:
	//    id (UUID) Unique identifier of consumer to be removed
	// Retruns:
	//    bool Removal success state
	RemoveConsumerById(id UUID) bool
	//Execute Transaction Element
	// Parameters:
	//    te (logic.ESBTransactionElement) Transaction Element to be executed
	// Retruns:
	//    error Any error that occurs during the computation or nil
	__execute(te ESBTransactionElement) error
}

type __serviceBusStruct struct {
	firstInChannel BOChannel
	allInChannel   BOChannel
	consumers      types.List
	firstInRunning bool
	allInRunning   bool
	busIsRunning   bool
}

func (esb *__serviceBusStruct) __readFirstInChannel() {
	esb.firstInRunning = true
	defer func() {
		var err error
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				err = itf.(error)
			} else {
				err = errors.New(fmt.Sprintf("%v", itf))
			}
		}
		if logger != nil {
			logger.Error(err)
		}
		esb.firstInRunning = false
	}()
	for esb.busIsRunning {
		select {
		case in_bo := <-esb.firstInChannel:
			esb.OfferFirst(in_bo)
		case <-time.After(BUSINESSOBJECT_BUS_TIMEOUT):
		}
	}
	esb.firstInRunning = false
}

func (esb *__serviceBusStruct) OfferFirst(bo BusinessObject) error {
	return nil
}

func (esb *__serviceBusStruct) OfferFirstChannel() *BOChannel {
	if !esb.firstInRunning {
		go esb.__readFirstInChannel()
	}
	return &esb.firstInChannel
}

func (esb *__serviceBusStruct) __readAllInChannel() {
	esb.allInRunning = true
	defer func() {
		var err error
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				err = itf.(error)
			} else {
				err = errors.New(fmt.Sprintf("%v", itf))
			}
		}
		if logger != nil {
			logger.Error(err)
		}
		esb.allInRunning = false
	}()
	for esb.busIsRunning {
		select {
		case in_bo := <-esb.allInChannel:
			esb.Offer(in_bo)
		case <-time.After(BUSINESSOBJECT_BUS_TIMEOUT):
		}
	}
	esb.allInRunning = false
}

func (esb *__serviceBusStruct) Offer(bo BusinessObject) error {
	return nil
}

func (esb *__serviceBusStruct) OfferChannel() *BOChannel {
	if !esb.allInRunning {
		go esb.__readAllInChannel()
	}
	return &esb.allInChannel
}

func (esb *__serviceBusStruct) RegisterConsumer(consumer BusinessConsumer) bool {
	return esb.consumers.Add(consumer)
}

func (esb *__serviceBusStruct) RemoveConsumerById(id UUID) bool {
	var consumer BusinessConsumer
	iter := esb.consumers.Iterator()
	for iter.HasNext() {
		item := iter.Next()
		if item != nil {
			xCons := item.(BusinessConsumer)
			uuid := xCons.Id()
			if uuid == id {
				consumer = item.(BusinessConsumer)
			}
		}
	}
	if consumer == nil {
		return false
	}
	return esb.consumers.Remove(consumer)
}

func (esb *__serviceBusStruct) RemoveConsumer(consumer BusinessConsumer) bool {
	return esb.consumers.Remove(consumer)
}

func (esb *__serviceBusStruct) __execute(te ESBTransactionElement) error {
	var err error
	defer func() {
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				err = itf.(error)
			} else {
				err = errors.New(fmt.Sprintf("%v", itf))
			}
		}
		if logger != nil {
			logger.Error(err)
		}
	}()
	consumer := *te.Consumer
	var waitingTime time.Time = time.Now()
	var maxWaitingTime time.Time = waitingTime.Add(CONSUMERS_MAX_WAITING_TIMEOUT)
	for !consumer.IsReady() {
		time.Sleep(CONSUMERS_SUSPEND_TIMEOUT)
		if waitingTime.Sub(maxWaitingTime).Nanoseconds() > int64(0) {
			err = errors.New(fmt.Sprintf("Max unavailability timeout accessing to Consumer %s (id: %s)", consumer.Label(), consumer.Id()))
			break
		}
	}
	bo := *te.BusinessObject
	accepted := consumer.Accept(bo)
	if !accepted {
		err = errors.New(fmt.Sprintf("Error for Consumer %s (id: %s) accepting Businees Object of type %s", consumer.Label(), consumer.Id(), bo.GetType().Kind().String()))
	}
	return err
}
