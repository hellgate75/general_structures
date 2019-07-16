package types

import (
	"reflect"
)

// Element that Contains information about data and next node, it's used to maintain a hierarchy and sequence into the Collections components
type CollectionElement struct {
	Element RowElement
	Next    *CollectionElement
}

// Collection Component containing unlimited size of elements and
// providing mulitple features
type Iterator interface {
	// Assign Component Element Type
	// Parameters:
	//   t (reflect.Type) Component Type
	Type(t reflect.Type)
	// Retrieve Component Element Type
	// Returns:
	//   reflect.Type Component Type
	GetType() reflect.Type
	// Retrieve if iterator has a next element to merge
	// Returns:
	//   bool Iterator next
	HasNext() bool
	// Retrieve if iterator has a next element
	// Returns:
	//   types.RowElement next element if present or nil elsewise
	Next() RowElement
}

type __iteratorType struct {
	__rootElement    *CollectionElement
	__currentElement *CollectionElement
	__type           reflect.Type
	__inited         bool
}

func (iterator *__iteratorType) Type(t reflect.Type) {
	iterator.__type = t
}

func (iterator *__iteratorType) GetType() reflect.Type {
	return iterator.__type
}

func (iterator *__iteratorType) HasNext() bool {
	if iterator.__currentElement == nil {
		return !iterator.__inited
	}
	return iterator.__currentElement.Next != nil
}

func (iterator *__iteratorType) Next() RowElement {
	if iterator.__currentElement == nil && !iterator.__inited {
		iterator.__inited = true
		iterator.__currentElement = iterator.__rootElement
		if iterator.__currentElement == nil {
			panic("No Element in the iterator root!!")
		}
		return iterator.__currentElement.Element
	} else if iterator.__currentElement != nil {
		iterator.__currentElement = iterator.__currentElement.Next
		if iterator.__currentElement == nil {
			panic("Element is nil, maybe iterator is out of scope!!")
		}
		return iterator.__currentElement.Element
	}
	panic("No Element in the iterator root!!")
}

// Create New Iterator based on the Root of a Collection Elements hierarchy
// Parameters:
//   t (reflect.Type) Component Type
//   rootElement (*CollectionElement) Pointer to Collection Elements hierarchy root item
// Returns:
//   types.Iterator Iterator component instance
func NewIterator(t reflect.Type, rootElement *CollectionElement) Iterator {
	return &__iteratorType{
		__rootElement: rootElement,
		__type:        t,
	}
}
