package logic

import (
	"github.com/hellgate75/general_structures/types"
	"github.com/hellgate75/general_utils/log"
	"io"
	"sync"
	"time"
)

var logger log.Logger

// Operation type, describes informtion about events on Transactions
type TransactionOperation int

//Transacton Operations Enumeration
const (
	// Describe Transaction creation event
	CREATED TransactionOperation = iota + 1
	// Describe Transaction add element event
	ADD_ELEMENT
	// Describe Transaction remove element event
	REMOVE_ELEMENT
	// Describe Transaction paused event
	PAUSED
	// Describe Transaction resume paused event
	RESUMED
	// Describe Transaction savepoint event
	SAVEDPOINT
	// Describe Transaction rollback event
	ROLLBACKED
	// Describe Transaction commit event
	COMMITTED
	// Describe Transaction deletion event
	DELETED
)

//Initialize package logger if not started
func InitLogger() {
	currentLogger, err := log.New("logic")
	if err != nil {
		panic(err.Error())
	}
	logger = currentLogger
}

// Business Objects channel that is used to send/receive messages, data, instructions
type BOChannel chan BusinessObject

// Unique Identifier
type UUID string

// Map that represents the data fields (name is the string key, and value is the returned interface)
type ValuesMap map[string]interface{}

// Array that contains key list
type Keys []string

// Describes a component that can be exported/imported to/from streams
type ExportableObject interface {
	// Export Data into the flow
	// Parameters:
	//    flows.Flow Input Stream
	// Returns:
	//    errors Any error that occurs during the computation or nil
	Export(f io.Writer) error
	// Import Data from the flow
	// Parameters:
	//    flows.Flow Output Stream
	// Returns:
	//    errors Any error that occurs during the computation or nil
	Import(f io.Reader) error
}

// Interface that describes receiver features
type BusinessConsumer interface {
	// Method that retrieve consumer label for logging purposes
	// Returns:
	//    string Consumer descriptive label
	Label() string
	// Method that retrieve consumer unique Identifier
	// Returns:
	//    logic.UUID Unique Consumer Identifier
	Id() UUID
	// Method that accepts a Business Object
	// Parameters:
	//    bo (logic.BusinessObject) Object to be consumed
	// Returns:
	//    bool Acceptance state
	Accept(bo BusinessObject) bool
	// Retrieve state if a consumer is ready to accept B.O.s
	// Returns:
	//    bool Acceptance ready state
	IsReady() bool
}

// Define Single Transaction Operation
type ESBTransactionElement struct {
	Id             UUID
	Consumer       *BusinessConsumer
	BusinessObject *BusinessObject
	Nextelement    *ESBTransactionElement
}

// Defines transaction point, it is a save point into the transaction. It's used to maintain an order in the transaction execution sequence.
// It keep trace of entry logic.ESBTransactionElement and current one, holding information about next execution point.
type ESBTransactionPoint struct {
	Id             UUID
	RootElement    *ESBTransactionElement
	CurrentElement *ESBTransactionElement
	NextPoint      *ESBTransactionPoint
	PreviousPoint  *ESBTransactionPoint
	__mutex        sync.RWMutex
}

// Transaction History item, containing information about events heppened on the Transaction life and his saved working points,
// retaining infomration about deleted elements or save points
type TransactionHistoryEntry struct {
	Id          UUID
	Date        time.Time
	Operation   TransactionOperation
	Description string
}

//Information about Transaction and hierarchy
type ESBTransactionInformationEntry struct {
	EntryPoint ESBTransactionPoint
	NextEntry  *ESBTransactionInformationEntry
	Created    time.Time
	Updated    time.Time
	Paused     bool
	Errors     types.List
}
