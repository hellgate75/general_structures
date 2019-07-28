package logic

import (
	"encoding/gob"
	"errors"
	"fmt"
	"github.com/hellgate75/general_structures/types"
	errs "github.com/hellgate75/general_utils/errors"
	"io"
	"reflect"
	"sync"
	"time"
)

// Describes behavior of Transaction Manager. This component manage and hold serveral Transactions and Save points, giving an easy access to transaction commands.
// It's used by the logic.ServiceBus to manage transaction according to implememtation policies.
type ESBTransactionManager interface {
	ExportableObject
	// Creates and starts a New Empty Transaction
	// Returns:
	//    logic.UUID New Transaction Id
	NewTransaction() UUID
	// Starts a paused Transaction by logic.UUID
	// Parameters:
	//    transactionId (logic.UUID) Existing or New Trasaction Id
	// Returns:
	//    error Any error that occurs during the computation or nil
	Start(transactionId UUID) error
	// Pauses a Transaction by logic.UUID, any cancellation/finalization activities will be denied
	// Parameters:
	//    transactionId (logic.UUID) Existing Trasaction Id
	// Returns:
	//    error Any error that occurs during the computation or nil
	Pause(transactionId UUID) error
	// Verify if a Transaction is started (if existing) by logic.UUID
	// Parameters:
	//    transactionId (logic.UUID) Existing Trasaction Id
	// Returns:
	//    bool Transaction running state or false if not exists
	IsStarted(transactionId UUID) bool
	// Verify if a Transaction exists by logic.UUID
	// Parameters:
	//    transactionId (logic.UUID) Existing Trasaction Id
	// Returns:
	//    bool Transaction existance state
	Exists(transactionId UUID) bool
	// Adds an element to a Transaction by logic.UUID
	// Parameters:
	//    transactionId (logic.UUID) Existing Trasaction Id
	//    element (logic.ESBTransactionElement) Element to be added to the transaction as next step
	// Returns:
	//    (logic.UUID Transaction Working Point Save Unique Identifier,
	//     bool Transaction element insert state)
	AddElement(transactionId UUID, element ESBTransactionElement) (UUID, bool)
	// Save Working point of a Transaction by logic.UUID, it can resetted by Rollback command
	// Parameters:
	//    transactionId (logic.UUID) Existing Trasaction Id
	// Returns:
	//    (logic.UUID Transaction Working Point Save Unique Identifier,
	//     error Any error that occurs during the computation or nil)
	SavePoint(transactionId UUID) (UUID, error)
	// Remove all trasaction elements since last save point or since the beginning if no working point has been saved
	// Parameters:
	//    transactionId (logic.UUID) Existing Trasaction Id
	// Returns:
	//    error Any error that occurs during the computation or nil
	Rollback(transactionId UUID) error
	// Remove all trasaction elements since given save point logic.UUID or nothing if savepoint doesn't exist
	// Parameters:
	//    transactionId (logic.UUID) Existing Trasaction Identifier
	//    savePointId (logic.UUID) Saved Working Point Unique Identifier
	// Returns:
	//    error Any error that occurs during the computation or nil
	RollbackTo(transactionId UUID, savePointId UUID) error
	// Apply all transactions commands in all savepoints since the beginning, removing all transaction history, in case of success
	// Parameters:
	//    transactionId (logic.UUID) Existing Trasaction Id
	// Returns:
	//    error Any error that occurs during the computation or nil
	Commit(transactionId UUID) error
	// Removing all transaction history and any saved working points
	// Parameters:
	//    transactionId (logic.UUID) Existing Trasaction Id
	// Returns:
	//    error Any error that occurs during the computation or nil
	Delete(transactionId UUID) error
	// Retrive sequence of transaction information, and any saved working point
	// Parameters:
	//    transactionId (logic.UUID) Existing Trasaction Id
	// Returns:
	//    *logic.ESBTransactionInformationEntry Pointer to entry (root) Transaction Point
	GetInformation(transactionId UUID) *ESBTransactionInformationEntry
	// Retrive list of transaction history, and for any saved working point
	// Parameters:
	//    transactionId (logic.UUID) Existing Trasaction Id
	// Returns:
	//    types.List List conaining all history entries (logic.TransactionHistoryEntry)
	GetHistory(transactionId UUID) types.List
}

// Single Transaction Item Container
type __transactionItemStruct struct {
	entryPoint      ESBTransactionPoint
	currentPoint    *ESBTransactionPoint
	history         types.List
	transactionUUID UUID
	Created         time.Time
	Updated         time.Time
	Paused          bool
	Errors          types.List
}

func (te __transactionItemStruct) Equals(o __transactionItemStruct) bool {
	return te.transactionUUID == o.transactionUUID
}

// Strcuture that represent base structure for the logic.TransactionManager component
type __transactionManagerStruct struct {
	transactionItems types.List
	serviceBus       ServiceBus
	__mutex          sync.RWMutex
}

func (tm *__transactionManagerStruct) NewTransaction() UUID {
	timeIn := time.Now()
	var uuid UUID = NewUUID()
	rootPoint := ESBTransactionPoint{
		Id:      uuid,
		__mutex: sync.RWMutex{},
	}
	item := __transactionItemStruct{
		entryPoint:      rootPoint,
		Created:         timeIn,
		Updated:         timeIn,
		Errors:          types.NewList(reflect.TypeOf(errors.New(""))),
		Paused:          false,
		history:         types.NewList(reflect.TypeOf(TransactionHistoryEntry{})),
		transactionUUID: uuid,
	}
	item.currentPoint = &item.entryPoint
	tm.transactionItems.Add(item)
	return item.transactionUUID
}

func __deleteFromList(list types.List, id UUID) error {
	for list.Iterator().HasNext() {
		itf := list.Iterator().Next()
		tis := itf.(__transactionItemStruct)
		if tis.transactionUUID == id {
			list.Remove(tis)
			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to delete transaction %v, uuid not found!!", id))
}

func __getFromList(list types.List, id UUID) *__transactionItemStruct {
	for list.Iterator().HasNext() {
		itf := list.Iterator().Next()
		tis := itf.(__transactionItemStruct)
		if tis.transactionUUID == id {
			return &tis
		}
	}
	return nil
}

func (tm *__transactionManagerStruct) Start(transactionId UUID) error {
	var err error = nil
	tis := __getFromList(tm.transactionItems, transactionId)
	if tis != nil {
		if tis.Paused {
			tis.Paused = false
		}
	} else {
		err = errors.New(fmt.Sprintf("Unable to find transaction with id: %v", transactionId))
	}
	return err
}

func (tm *__transactionManagerStruct) Pause(transactionId UUID) error {
	var err error = nil
	tis := __getFromList(tm.transactionItems, transactionId)
	if tis != nil {
		if tis.Paused {
			tis.Paused = true
		}
	} else {
		err = errors.New(fmt.Sprintf("Unable to find transaction with id: %v", transactionId))
	}
	return err
}

func (tm *__transactionManagerStruct) IsStarted(transactionId UUID) bool {
	tis := __getFromList(tm.transactionItems, transactionId)
	if tis != nil {
		return !tis.Paused
	}
	return false
}
func (tm __transactionManagerStruct) Exists(transactionId UUID) bool {
	tis := __getFromList(tm.transactionItems, transactionId)
	return tis != nil
}

func (tm *__transactionManagerStruct) AddElement(transactionId UUID, element ESBTransactionElement) (UUID, bool) {
	var valid bool = true
	var uuid UUID = NewUUID()

	tis := __getFromList(tm.transactionItems, transactionId)
	if tis != nil {
		if tis.currentPoint.RootElement == nil {
			element.Id = uuid
			tis.currentPoint.RootElement = &element
			tis.currentPoint.CurrentElement = &element
		} else {
			element.Id = uuid
			tis.currentPoint.CurrentElement.Nextelement = &element
			tis.currentPoint.CurrentElement = &element
		}
	} else {
		uuid = ""
		valid = false
	}
	return uuid, valid
}

func (tm *__transactionManagerStruct) SavePoint(transactionId UUID) (UUID, error) {
	var err error = nil
	var uuid UUID = NewUUID()

	tis := __getFromList(tm.transactionItems, transactionId)
	if tis != nil {
		newPoint := ESBTransactionPoint{
			Id:      uuid,
			__mutex: sync.RWMutex{},
		}
		tis.currentPoint.NextPoint = &newPoint
		newPoint.PreviousPoint = tis.currentPoint
		tis.currentPoint = &newPoint
		tis.Updated = time.Now()
	} else {
		uuid = ""
		err = errors.New(fmt.Sprintf("Unable to find transaction with id: %v", transactionId))
	}
	return uuid, err
}

func (tm *__transactionManagerStruct) Rollback(transactionId UUID) error {
	var err error = nil

	tis := __getFromList(tm.transactionItems, transactionId)
	if tis != nil {
		if tis.entryPoint.Id == tis.currentPoint.Id {
			var uuid UUID = NewUUID()
			rootPoint := ESBTransactionPoint{
				Id:      uuid,
				__mutex: sync.RWMutex{},
			}
			tis.entryPoint = rootPoint
			tis.currentPoint = &tis.entryPoint
			tis.transactionUUID = uuid
			return nil
		}
		tis.currentPoint = tis.currentPoint.PreviousPoint
		tis.Updated = time.Now()
	} else {
		err = errors.New(fmt.Sprintf("Unable to find transaction with id: %v", transactionId))
	}
	return err
}

func (tm *__transactionManagerStruct) RollbackTo(transactionId UUID, savePointId UUID) error {
	var err error = nil
	tis := __getFromList(tm.transactionItems, transactionId)
	if tis != nil {
		if tis.entryPoint.Id == savePointId {
			var uuid UUID = NewUUID()
			rootPoint := ESBTransactionPoint{
				Id:      uuid,
				__mutex: sync.RWMutex{},
			}
			tis.entryPoint = rootPoint
			tis.currentPoint = &tis.entryPoint
			tis.transactionUUID = uuid
			return nil
		}
		leaf := tis.currentPoint
		for leaf != nil && leaf.Id != savePointId {
			leaf = leaf.PreviousPoint
		}
		if leaf != nil {
			tis.currentPoint = leaf.PreviousPoint
			if tis.currentPoint == nil {
				tis.currentPoint = &tis.entryPoint
			}
		} else {
			err = errors.New(fmt.Sprintf("Unable to find in transaction %v save point with id: %v", transactionId, savePointId))

		}
	} else {
		err = errors.New(fmt.Sprintf("Unable to find transaction with id: %v", transactionId))
	}
	return err

}

func (tm *__transactionManagerStruct) Commit(transactionId UUID) error {
	var err error = nil

	if tm.serviceBus == nil {
		return errors.New("Service Bus is not setled up, no commit available!!")
	}

	tis := __getFromList(tm.transactionItems, transactionId)
	if tis != nil {

		leaf := &tis.entryPoint
		for leaf != nil {
			elem := leaf.RootElement
			for elem != nil {
				err := tm.serviceBus.__execute(*elem)
				if err != nil {
					return errors.New(fmt.Sprintf("Error executing transaction %s, savepoint %s, error is %s", transactionId, leaf.Id, err.Error()))
				}
				elem = (*elem).Nextelement
			}
			leaf = leaf.NextPoint
		}
	} else {
		err = errors.New(fmt.Sprintf("Unable to find transaction with id: %v", transactionId))
	}
	return err

}

func (tm *__transactionManagerStruct) Delete(transactionId UUID) error {
	var err error
	err = __deleteFromList(tm.transactionItems, transactionId)
	return err
}

func __createESBTransactionInformationEntry(tp ESBTransactionPoint, created time.Time, updated time.Time, paused bool, errors types.List) ESBTransactionInformationEntry {
	return ESBTransactionInformationEntry{
		EntryPoint: tp,
		NextEntry:  nil,
		Created:    created,
		Updated:    updated,
		Paused:     paused,
		Errors:     errors,
	}
}

func (tm *__transactionManagerStruct) GetInformation(transactionId UUID) *ESBTransactionInformationEntry {
	tis := __getFromList(tm.transactionItems, transactionId)
	if tis != nil {
		elem := &tis.entryPoint
		root := __createESBTransactionInformationEntry(*elem, tis.Created, tis.Updated, tis.Paused, tis.Errors)
		prev := root
		elem = elem.NextPoint
		for elem != nil {
			next := __createESBTransactionInformationEntry(*elem, tis.Created, tis.Updated, tis.Paused, tis.Errors)
			prev.NextEntry = &next
			prev = next
			elem = elem.NextPoint
		}
		return &root
	}
	return nil
}

func (tm *__transactionManagerStruct) GetHistory(transactionId UUID) types.List {
	tis := __getFromList(tm.transactionItems, transactionId)
	if tis != nil {
		return tis.history
	}
	return nil
}

func (tm *__transactionManagerStruct) Export(f io.Writer) error {
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
	}()
	encoder := gob.NewEncoder(f)
	encoder.Encode(tm)
	return err
}
func (tm *__transactionManagerStruct) Import(f io.Reader) error {
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
	}()
	decoder := gob.NewDecoder(f)
	decoder.Decode(tm)
	return err
}
