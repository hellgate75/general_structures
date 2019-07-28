package logic

import (
	"encoding/gob"
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
	"io"
	"reflect"
)

// Componponent that describe data that can be sent / received to / from the Bus
type BusinessObject interface {
	ExportableObject
	// Get the Map that describes the component fields/value
	// Returns:
	//    logic.ValueMap Map that contains business object field names and values
	GetValuesMap() ValuesMap
	// Get the Key array
	// Returns:
	//    logic.Keys Business object field name array
	GetKeys() Keys
	// Get the Business Object Type
	// Returns:
	//    reflect.Type Business object field type
	GetType() reflect.Type
	// Get the Number of Fields, saved into the Business Objetc
	// Returns:
	//    int Number of Fields
	Size() int
}

type __businessData struct {
	fieldsMap  map[string]interface{}
	objectType reflect.Type
}

// Export Data into the flow
// Parameters:
//    flows.Flow Input Stream
// Returns:
//    error Any error that occurs during the computation or nil
func (business *__businessData) Export(f io.Writer) error {
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
	encoder.Encode(*business)
	return err
}

// Import Data from the flow
// Parameters:
//    flows.Flow Output Stream
// Returns:
//    error Any error that occurs during the computation or nil
func (business *__businessData) Import(f io.Reader) error {
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
	decoder.Decode(business)
	return err
}

// Get the Map that describes the component fields/value
// Returns:
//    logic.ValueMap Map that contains business object field names and values
func (business *__businessData) GetValuesMap() ValuesMap {
	return ValuesMap(business.fieldsMap)
}

// Get the Key array
// Returns:
//    logic.Keys business object field name array
func (business *__businessData) GetKeys() Keys {
	var lst Keys = make(Keys, 0)
	for k, _ := range business.fieldsMap {
		//		fmt.Println("GetKeys() - k:", k)
		lst = append(lst, k)
	}
	return lst
}

func (business *__businessData) Size() int {
	return len(business.GetKeys())
}

// Get the Business Object Type
// Returns:
//    reflect.Type business object field type
func (business *__businessData) GetType() reflect.Type {
	return business.objectType
}
