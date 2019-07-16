package types

import (
	"fmt"
	"reflect"
)

// Comparator functin that returns 1, 0, -1 in case first item suites bigger, same or lower position in the list
type ComparatorFunc func(first RowElement, second RowElement) int

//Returns if an element is equal to another
type MatcherFunc func(first RowElement, second RowElement) bool

// List Component containing unlimited size of elements and
// providing mulitple features
type List interface {
	Collection
	// Sort list by Equals() bool method of the interface component or majority comparator
	// Parameters:
	//    compare (types.ComparatorFunc) Function used to compare elements in the List
	// Returns:
	//    bool Sort success indicator
	Sort(compare ComparatorFunc) bool
	// Verify if an element is inlcuded in the list using equality comparator
	// Parameters:
	//    item (types.RowElement) Element to have been seeked for into the list
	// Returns:
	//    bool Contains success indicator
	Contains(item RowElement) bool
	// Verify if an element is inlcuded in the list using MatcherFunc
	// Parameters:
	//    item (types.RowElement) Element to have been seeked for into the list
	//    matcher (types.MatcherFunc) Function that equals
	// Returns:
	//    bool Contains success indicator
	ContainsAs(item RowElement, matcher MatcherFunc) bool
	// Get an element based on his index, or panic in case index is out of range
	// Parameters:
	//    index (int64) Element index in the List
	// Returns:
	//    types.RowElement Element found in the List
	Get(index int64) RowElement
	// Set an element based on his index, or panic in case index is out of range
	// Parameters:
	//    index (int64) Element index in the List
	//    item (types.RowElement) Element to have been setted up into the list
	// Returns:
	//    bool Operation susccess status
	Set(index int64, item RowElement) bool
	// Add an element after specified index, or panic in case index is out of range
	// Parameters:
	//    index (int64) Element index in the List
	//    item (types.RowElement) Element to have been setted up into the list
	// Returns:
	//    bool Operation susccess status
	AddAfter(index int64, item RowElement) bool
	// Add an element after specified index, or panic in case index is out of range
	// Parameters:
	//    index (int64) Element index in the List
	//    items (types.RowArray) Element list to have been setted up into the list
	// Returns:
	//    int64 Number of elements Added into the List
	AddAllAfter(index int64, items RowArray) int64
	// Add an element after specified index, or panic in case index is out of range
	// Parameters:
	//    index (int64) Element index in the List
	//    coll (types.Collection) Collection of Elements to have been setted up into the list
	// Returns:
	//    int64 Number of elements Added into the List
	AddCollectionAfter(index int64, coll Collection) int64
	// Remove Element at specified index, or panic in case index is out of range
	// Parameters:
	//    index (int64) Element index in the List
	// Returns:
	//    bool Operation susccess status
	RemoveAt(index int64) bool
	// Returns the size of the List
	// Returns:
	//    int64 Number of elements in the List
	Size() int64
	// Returns a sub list of the current List
	// Parameters:
	//    start (int64) First Element index in the new Sub List
	//    end (int64) Last Element index in the new Sub List (excluded)
	// Returns:
	//    List New List containing elements between start (included) and end (excluded)
	SubList(start int64, end int64) List
}

type __ListStruct struct {
	__collectionStruct
}

func __sortListElements(root *CollectionElement, compare ComparatorFunc) bool {
	var element *CollectionElement = root
	var changes bool = false
	next := element.Next
	for next != nil {
		//		fmt.Println(fmt.Sprintf("BEFORE - Elem (p:%p) : %v", element, *element))
		//		fmt.Println(fmt.Sprintf("BEFORE - Next (p:%p) : %v", next, *next))
		val := compare(element.Element, next.Element)
		if val > 0 {
			value := element.Element
			element.Element = next.Element
			next.Element = value
			// Restart from scratch
			changes = true
			//			fmt.Println(fmt.Sprintf("AFTER - Elem (p:%p) : %v", element, *element))
			//			fmt.Println(fmt.Sprintf("AFTER - Next (p:%p) : %v", next, *next))
			break
		}
		element = next
		next = next.Next
	}
	if changes {
		__sortListElements(root, compare)
	}
	return changes
}

func (list *__ListStruct) Sort(compare ComparatorFunc) bool {
	defer func() {
		recover()
		list.__lock.Unlock()
	}()
	list.__lock.Lock()
	state := __sortListElements(list.__rootElement, compare)
	return state
}

func (list *__ListStruct) Size() int64 {
	return list.__size
}

func __containsListElements(root *CollectionElement, value RowElement, matcher *MatcherFunc) bool {
	var element *CollectionElement = root
	if matcher != nil {
		matcherFunc := *matcher
		if matcherFunc(element.Element, value) {
			return true
		}
		next := element.Next
		for next != nil {
			if matcherFunc(next.Element, value) {
				return true
			}
			element = next
			next = element.Next
		}
	} else {
		if element.Element == value {
			return true
		}
		next := element.Next
		for next != nil {
			if next.Element == value {
				return true
			}
			element = next
			next = element.Next
		}
	}
	return false
}

func (list *__ListStruct) Contains(item RowElement) bool {
	var state bool = false
	defer func() {
		recover()
		list.__lock.RUnlock()
	}()
	list.__lock.RLock()
	state = __containsListElements(list.__rootElement, item, nil)
	return state
}

func (list *__ListStruct) ContainsAs(item RowElement, matcher MatcherFunc) bool {
	var state bool = false
	defer func() {
		recover()
		list.__lock.RUnlock()
	}()
	list.__lock.RLock()
	state = __containsListElements(list.__rootElement, item, &matcher)
	return state
}

func (list *__ListStruct) Get(index int64) RowElement {
	if index < int64(0) || index >= list.__size {
		panic(fmt.Sprintf("Index <%x> out of bounds %x <-> %x", index, 0, list.__size-1))
	}
	currentIndex := int64(0)
	elem := list.__rootElement
	for elem != nil && currentIndex < index {
		elem = elem.Next
		currentIndex++
	}
	if currentIndex != index {
		panic(fmt.Sprintf("Index <%x> not reachable, too few elements in the List!!!", index))
	}
	if elem == nil {
		panic(fmt.Sprintf("Element at index <%x> seems nil in the List!!!", index))
	}
	return elem.Element
}

func (list *__ListStruct) Set(index int64, item RowElement) bool {
	if index < int64(0) || index >= list.__size {
		panic(fmt.Sprintf("Index <%x> out of bounds %x <-> %x", index, 0, list.__size-1))
	}
	currentIndex := int64(0)
	elem := list.__rootElement
	for elem != nil && currentIndex < index {
		elem = elem.Next
		currentIndex++
	}
	if currentIndex != index {
		panic(fmt.Sprintf("Index <%x> not reachable, too few elements in the List!!!", index))
	}
	if elem == nil {
		panic(fmt.Sprintf("Element at index <%x> seems nil in the List!!!", index))
	}
	elem.Element = item
	return list.Get(index) == item
}

func (list *__ListStruct) AddAfter(index int64, item RowElement) bool {
	if index < int64(0) || index >= list.__size {
		panic(fmt.Sprintf("Index <%x> out of bounds %x <-> %x", index, 0, list.__size-1))
	}
	currentIndex := int64(0)
	elem := list.__rootElement
	for elem != nil && currentIndex < index {
		elem = elem.Next
		currentIndex++
	}
	if currentIndex != index {
		panic(fmt.Sprintf("Index <%x> not reachable, too few elements in the List!!!", index))
	}
	if elem == nil {
		panic(fmt.Sprintf("Element at index <%x> seems nil in the List!!!", index))
	}
	if elem == nil {
		return false
	}
	defer func() {
		recover()
		list.__lock.Unlock()
	}()
	list.__lock.Lock()
	element := __makeBaseCollectionElement(item)
	next := elem.Next
	elem.Next = element
	element.Next = next
	list.__size++
	return list.Get(index+1) == item
}

func (list *__ListStruct) AddAllAfter(index int64, items RowArray) int64 {
	if index < int64(0) || index >= list.__size {
		panic(fmt.Sprintf("Index <%x> out of bounds %x <-> %x", index, 0, list.__size-1))
	}
	currentIndex := int64(0)
	elem := list.__rootElement
	for elem != nil && currentIndex < index {
		elem = elem.Next
		currentIndex++
	}
	if currentIndex != index {
		panic(fmt.Sprintf("Index <%x> not reachable, too few elements in the List!!!", index))
	}
	if elem == nil {
		panic(fmt.Sprintf("Element at index <%x> seems nil in the List!!!", index))
	}
	if elem == nil {
		return int64(0)
	}
	defer func() {
		recover()
		list.__lock.Unlock()
	}()
	list.__lock.Lock()
	var count int64 = int64(0)
	for _, item := range items {
		element := __makeBaseCollectionElement(item)
		next := elem.Next
		elem.Next = element
		element.Next = next
		elem = element
		count++
		list.__size++
	}
	return count
}

func (list *__ListStruct) AddCollectionAfter(index int64, coll Collection) int64 {
	if index < int64(0) || index >= list.__size {
		panic(fmt.Sprintf("Index <%x> out of bounds %x <-> %x", index, 0, list.__size-1))
	}
	currentIndex := int64(0)
	elem := list.__rootElement
	for elem != nil && currentIndex < index {
		elem = elem.Next
		currentIndex++
	}
	if currentIndex != index {
		panic(fmt.Sprintf("Index <%x> not reachable, too few elements in the List!!!", index))
	}
	if elem == nil {
		panic(fmt.Sprintf("Element at index <%x> seems nil in the List!!!", index))
	}
	if elem == nil {
		return int64(0)
	}
	defer func() {
		recover()
		list.__lock.Unlock()
	}()
	list.__lock.Lock()
	var count int64 = int64(0)
	iter := coll.Iterator()
	for iter.HasNext() {
		element := __makeBaseCollectionElement(iter.Next())
		next := elem.Next
		elem.Next = element
		element.Next = next
		elem = element
		count++
		list.__size++
	}
	return count
}

func (list *__ListStruct) RemoveAt(index int64) bool {
	var out bool = true
	if index < int64(0) || index >= list.__size {
		panic(fmt.Sprintf("Index <%x> out of bounds %x <-> %x", index, 0, list.__size-1))
	}
	currentIndex := int64(0)
	var previous *CollectionElement = nil
	elem := list.__rootElement
	for elem != nil && currentIndex < index {
		previous = elem
		elem = elem.Next
		currentIndex++
	}
	if currentIndex != index {
		panic(fmt.Sprintf("Index <%x> not reachable, too few elements in the List!!!", index))
	}
	if elem == nil {
		panic(fmt.Sprintf("Element at index <%x> seems nil in the List!!!", index))
	}
	if elem == nil {
		return false
	}
	defer func() {
		r := recover()
		if r != nil {
			out = false
		}
		list.__lock.Unlock()
	}()
	list.__lock.Lock()
	if previous != nil {
		previous.Next = elem.Next
		list.__size--
	} else if previous == nil && index == int64(0) {
		next := list.__rootElement.Next
		if next == nil {
			list.Clear()
		} else {
			list.__rootElement = next
			list.__size--
		}
	} else {
		return false
	}
	return out
}

func (list *__ListStruct) SubList(start int64, end int64) List {
	if start < int64(0) || start >= list.__size {
		panic(fmt.Sprintf("Start Index <%x> out of bounds %x <-> %x", start, 0, list.__size-1))
	}
	if end < int64(0) || end >= list.__size {
		panic(fmt.Sprintf("End Index <%x> out of bounds %x <-> %x", end, 0, list.__size-1))
	}
	if start > end {
		//Swap indexes in case start is bigger than end, and both are good
		tmp := start
		start = end
		end = tmp
	}
	defer func() {
		recover()
		list.__lock.RUnlock()
	}()
	list.__lock.RLock()
	// We use scrolling into CollectionElement sequence, because we want use huge sublist features, so no int size limitation, but in64 size limitation
	var newRoot *CollectionElement = nil
	var newLast *CollectionElement = nil
	index := int64(0)
	iter := list.Iterator()
	for iter.HasNext() {
		if index >= start && index < end {
			item := iter.Next()
			if newRoot == nil {
				newRoot = __makeBaseCollectionElement(item)
				newLast = newRoot
			} else {
				newNode := __makeBaseCollectionElement(item)
				newLast.Next = newNode
				newLast = newNode
			}
		} else if index >= end {
			break
		}
		index++
	}
	newListIterator := NewIterator(list.GetType(), newRoot)
	return IteratorAsList(newListIterator)
}

// Create New List component
// Parameters:
//   t (reflect.Type) Component Type
// Returns:
//   types.List List component instance
func NewList(t reflect.Type) List {
	return &__ListStruct{
		__collectionStruct{
			__rootElement:   nil,
			__lastElement:   nil,
			__type:          t,
			__size:          int64(0),
			__componentName: "List",
		},
	}
}

// Create New List component filled with elements from an array
// Parameters:
//   t (reflect.Type) Component Type
//   arr (types.RowArray) Array of generic type elements
// Returns:
//   types.List List component instance
func NewListWithArray(t reflect.Type, arr RowArray) List {
	coll := __ListStruct{
		__collectionStruct{
			__rootElement:   nil,
			__lastElement:   nil,
			__type:          t,
			__size:          int64(0),
			__componentName: "List",
		},
	}
	coll.AddAll(arr)
	return &coll
}

// Create a List using a collection instance
// Parameters:
//   coll (types.Collection) Origin Collection component instance
// Returns:
//   types.List List component instance
func CollectionAsList(coll Collection) List {
	collect := __ListStruct{
		__collectionStruct{
			__rootElement:   nil,
			__lastElement:   nil,
			__type:          coll.GetType(),
			__size:          int64(0),
			__componentName: "List",
		},
	}
	collect.AddCollection(coll)
	return &collect
}

// Create a List, using an Iterator component instance
// Parameters:
//   iterator (types.Iterator) Origin Iterator component instance
// Returns:
//   types.List List component instance
func IteratorAsList(iterator Iterator) List {
	return CollectionAsList(IteratorToCollection(iterator))
}
