package logic

import (
	"errors"
	"fmt"
	"github.com/hellgate75/general_structures/types"
	"reflect"
	"sync"
)

// Create a Business Object from elements (Business Obect Map and Type)
// Parameters:
//    boMap (logic.ValuesMap) Map of object fields (name, and value entries)
//    boType (reflect.Type) Type related to the original type
// Returns:
//    logic.BusinessObject Outcome Business Object
func NewBusinessObject(boMap ValuesMap, boType reflect.Type) BusinessObject {
	return &__businessData{
		fieldsMap:  boMap,
		objectType: boType,
	}
}

// Convert a structure to a Business Object (logic.BusinessObject) or for any other type will be saved in "Val" field
// Parameters:
//    itf (interface{}) Input component/structure
// Returns:
//    (logic.BusinessObject Object containing the input fields or the 'Val' field in case of non strcuture input element,
//     error Any error that occurs during the computation or nil)
func ConvertStructure(itf interface{}) (BusinessObject, error) {
	var tp reflect.Type = reflect.TypeOf(itf)
	var val reflect.Value = reflect.ValueOf(itf)
	var mp map[string]interface{} = make(map[string]interface{})

	if tp.Kind() == reflect.Struct {
		numFields := tp.NumField()
		for i := 0; i < numFields; i++ {
			sf := tp.Field(i)
			name := sf.Name
			mp[name] = val.Field(i).Interface()
		}
	} else {
		return nil, errors.New(fmt.Sprintf("Provided Object Type <%s> is not a structure!!", tp.Kind().String()))
	}
	return NewBusinessObject(ValuesMap(mp), tp), nil
}

// Convert a Bussiness Object into a provided structure, in case it is suitable for the logic.BusinessObject
// Parameters:
//    o (logic.BusinessObject) Reference Business Object
//    itf (interface{}) Pointer to structure
// Returns:
//    error Any error that occurs during the computation or nil
func ConvertBusinessObject(o BusinessObject, itf interface{}) error {
	mSize := o.Size()
	if mSize == 0 {
		return errors.New(fmt.Sprintf("Input Business Object has %d Fields, No output is available", mSize))
	}
	mType := o.GetType()
	itfType := reflect.TypeOf(itf)
	if itfType.Kind() == reflect.Ptr {
		itfType = itfType.Elem()
	}
	//	fmt.Println(mType)
	//	fmt.Println(itfType)
	//	if !mType.ConvertibleTo(itfType) {
	//		return errors.New(fmt.Sprintf("Unable to assign map made of type <%s> to type <%s>", itfType.Kind().String(), mType.Kind().String()))
	//	}
	itfVal := reflect.ValueOf(itf)
	if itfVal.Kind() == reflect.Ptr {
		itfVal = itfVal.Elem()
	}
	mFld := o.GetValuesMap()
	var output string = ""
	for k, v := range mFld {
		if _, ok := itfType.FieldByName(k); ok {
			fld := itfVal.FieldByName(k)
			if fld.Kind() == reflect.Ptr {
				fld = fld.Elem()
			}
			if !fld.IsValid() {
				output = fmt.Sprintf("%sStruct Field <%s> is not valid type <%s>,", output, k, mType.Kind().String())
			} else if !fld.CanSet() {
				output = fmt.Sprintf("%sStruct Field <%s> cannot be setted up in type <%s>,", output, k, mType.Kind().String())
			} else {
				val := reflect.ValueOf(v)
				if val.Kind() == reflect.Ptr {
					val = val.Elem()
				}
				fld.Set(val)
			}
		} else {
			output = fmt.Sprintf("%sMap Field <%s> is not contained in the structure <%s>,", output, k, mType.Kind().String())
		}
	}
	if len(output) > 0 {
		output = output[:len(output)-1]
		return errors.New(output)
	}
	return nil
}

// Create New Service Bus Transaction Manager
// Parameters:
//    bus (logic.ServiceBus) Service Bus instance
// Returns:
//    logic.ESBTransactionManager Service Bus Transaction Manager instance
func NewESBTransactionManager(bus ServiceBus) ESBTransactionManager {
	return &__transactionManagerStruct{
		serviceBus:       bus,
		transactionItems: types.NewList(reflect.TypeOf(__transactionItemStruct{})),
		__mutex:          sync.RWMutex{},
	}
}
