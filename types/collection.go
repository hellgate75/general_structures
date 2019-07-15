package types

import (
	//	"fmt"
	"reflect"
	"sync"
)

//Generic Item
type RowElement interface{}

//Generic Array
type RowArray []RowElement

// Collection Component containing unlimited size of elements and
// providing mulitple features
type Collection interface {
	// Retrieve Component Element Type
	// Returns:
	//   reflect.Type Component Type
	GetType() reflect.Type
	// Returns the array iterator providing elements in the Collection
	// Returns:
	//    Iterator Array iterator component that contains all list elements
	Iterator() Iterator
	// Returns the size of the Collection
	// Returns:
	//    int64 Number of elements in the Collection
	Size() int64
	// Add a new element in the Collection
	// Parameters:
	//    item (types.RowElement) Element to have been added into the list
	// Returns:
	//    bool Added success indicator
	Add(item RowElement) bool
	// Add all elements in the Collection
	// Parameters:
	//    items (types.RowArray) Elements to have been added into the Collection
	// Returns:
	//    bool Added success indicator
	AddAll(items RowArray) bool
	// Add all elements in the Collection
	// Parameters:
	//    collection (types.Collection) Colelction of elements to have been added into the Collection
	// Returns:
	//    bool Added success indicator
	AddCollection(collection Collection) bool
	// Remove an element from the Collection
	// Parameters:
	//    item (types.RowElement) Element to have been removed from the Collection
	// Returns:
	//    bool Added success indicator
	Remove(item RowElement) bool
	// Remove all occurancies of an element from the Collection
	// Parameters:
	//    item (types.RowElement) Element to have been removed from the Collection
	// Returns:
	//    int64 Number of elements removed from the Collection
	RemoveAll(item RowElement) int64
	// Get First Element of the Collection

	// Returns:
	//    types.RowElement First Element of the Collection
	First() RowElement
	// Get Last Element of the Collection
	// Returns:
	//    types.RowElement Last Element of the Collection
	Last() RowElement
}

type __collectionStruct struct {
	__rootElement *CollectionElement
	__lastElement *CollectionElement
	__type        reflect.Type
	__size        int64
	__lock        sync.RWMutex
}

func (collection *__collectionStruct) GetType() reflect.Type {
	return collection.__type
}

func (collection *__collectionStruct) Iterator() Iterator {
	return NewIterator(collection.__type, collection.__rootElement)
}

func (collection *__collectionStruct) Size() int64 {
	return collection.__size
}

func (collection *__collectionStruct) Add(item RowElement) bool {
	defer func() {
		recover()
		collection.__lock.Unlock()
	}()
	collection.__lock.Lock()
	element := &CollectionElement{
		Element: item,
		Next:    nil,
	}
	if collection.__rootElement == nil {
		collection.__rootElement = element
		collection.__size++
	} else if collection.__lastElement == nil {
		collection.__rootElement.Next = element
		collection.__lastElement = element
		collection.__size++
	} else {
		collection.__lastElement.Next = element
		collection.__lastElement = element
		collection.__size++
	}
	return true
}

func (collection *__collectionStruct) AddAll(items RowArray) bool {
	var state bool = true
	for _, v := range items {
		if !collection.Add(v) {
			state = false
		}
	}
	return state
}

func (collection *__collectionStruct) AddCollection(coll Collection) bool {
	var state bool = true
	var iter Iterator = coll.Iterator()
	for iter.HasNext() {
		if !collection.Add(iter.Next()) {
			state = false
		}
	}
	return state
}

func __containsMethod(obj interface{}, method string) bool {
	if obj == nil {
		return false
	}
	st := reflect.TypeOf(obj)
	_, ok := st.MethodByName(method)
	return ok

}

func __equals(first RowElement, second RowElement) bool {
	if first == nil && second == nil {
		return true
	} else if first == nil {
		return false
	} else if second == nil {
		return false
	}
	m1, hasM1 := reflect.TypeOf(first).MethodByName("Equals")

	if !hasM1 {
		return first == second
	} else {
		var args []reflect.Value = make([]reflect.Value, 0)
		args = append(args, reflect.ValueOf(second).Elem())
		values := m1.Func.Call(args)
		if len(values) > 0 {
			return values[0].Elem().Bool()
		}
		return false
	}
}

func (collection *__collectionStruct) Remove(item RowElement) bool {
	var state bool = false
	if collection.__rootElement == nil {
		return state
	}
	if __equals(collection.__rootElement.Element, item) {
		collection.__rootElement = collection.__rootElement.Next
		if collection.__rootElement == nil || collection.__rootElement.Next == nil {
			collection.__lastElement = nil
		}
		collection.__size--
		state = true
	} else {
		previous := collection.__rootElement
		elem := collection.__rootElement.Next
		for elem != nil {
			if __equals(elem.Element, item) {
				previous.Next = elem.Next
				if elem.Next == nil {
					collection.__lastElement = previous
				}
				collection.__size--
				state = true
				break
			}
			previous = elem
			elem = elem.Next
		}
	}
	return state
}

func (collection *__collectionStruct) RemoveAll(item RowElement) int64 {
	var state int64 = int64(0)
	for collection.__rootElement != nil && __equals(collection.__rootElement.Element, item) {
		collection.__rootElement = collection.__rootElement.Next
		collection.__size--
		state++
	}
	if collection.__rootElement == nil {
		collection.__lastElement = nil
		return state
	}
	if collection.__rootElement.Next == nil {
		collection.__lastElement = nil
	}
	previous := collection.__rootElement
	elem := collection.__rootElement.Next
	for elem != nil {
		//		fmt.Println(fmt.Sprintf("BEFORE - Prev (p:%p) : %v", previous, *previous))
		//		fmt.Println(fmt.Sprintf("BEFORE - Elem (p:%p) : %v", elem, *elem))
		//		fmt.Println(fmt.Sprintf("BEFORE - Size : %x", collection.__size))
		if __equals(elem.Element, item) {
			previous.Next = elem.Next
			collection.__size--
			state++
		}
		//		fmt.Println(fmt.Sprintf("AFTER - Prev (p:%p) : %v", previous, *previous))
		//		fmt.Println(fmt.Sprintf("AFTER - Elem (p:%p) : %v", elem, *elem))
		//		fmt.Println(fmt.Sprintf("AFTER - Size : %x", collection.__size))
		previous = elem
		elem = elem.Next
	}
	return state
}

func (collection *__collectionStruct) First() RowElement {
	return collection.__rootElement.Element
}

func (collection *__collectionStruct) Last() RowElement {
	return collection.__lastElement.Element
}

// Create New Collection component
// Parameters:
//   t (reflect.Type) Component Type
// Returns:
//   types.Collection Collection component instance
func NewCollection(t reflect.Type) Collection {
	return &__collectionStruct{
		__rootElement: nil,
		__lastElement: nil,
		__type:        t,
		__size:        int64(0),
	}
}

// Create New Collection component filled with elements from an array
// Parameters:
//   t (reflect.Type) Component Type
//   arr (types.RowArray) Array of generic type elements
// Returns:
//   types.Collection Collection component instance
func NewCollectionWithArray(t reflect.Type, arr RowArray) Collection {
	coll := __collectionStruct{
		__rootElement: nil,
		__lastElement: nil,
		__type:        t,
		__size:        int64(0),
	}
	coll.AddAll(arr)
	return &coll
}

// Clone a collection, creating a new instamce same to origin one
// Parameters:
//   coll (types.Collection) Origin Collection component instance
// Returns:
//   types.Collection Collection component instance
func CloneCollection(coll Collection) Collection {
	collect := __collectionStruct{
		__rootElement: nil,
		__lastElement: nil,
		__type:        coll.GetType(),
		__size:        int64(0),
	}
	collect.AddCollection(coll)
	return &collect
}

// Create a collection, using an Iterator component instance
// Parameters:
//   iterator (types.Iterator) Origin Iterator component instance
// Returns:
//   types.Collection Collection component instance
func IteratorToCollection(iterator Iterator) Collection {
	colleaction := NewCollection(iterator.GetType())
	for iterator.HasNext() {
		colleaction.Add(iterator.Next())
	}
	return colleaction
}
