package types

import (
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
	"math/big"
	"reflect"
	"runtime"
	"strconv"
)

type MapperFunc func(arg0 RowElement) RowElement

type BiArgumentsFunc func(arg0 RowElement, arg1 RowElement) RowElement

type ArrayArgumentFunc func(arg0 ...RowElement) Iterator

func __listToSingular(args0 ...RowElement) Iterator {
	if len(args0) == 0 {
		return NewIterator(reflect.TypeOf(int64(0)), nil)
	}
	var l List = NewList(args0[0].(List).GetType())
	for _, v := range args0 {
		l.AddCollection(v.(Collection))
	}
	return l.Iterator()
}
func __collectionToSingular(args0 ...RowElement) Iterator {
	if len(args0) == 0 {
		return NewIterator(reflect.TypeOf(int64(0)), nil)
	}
	var l List = NewList(args0[0].(Collection).GetType())
	for _, v := range args0 {
		l.AddCollection(v.(Collection))
	}
	return l.Iterator()
}
func __iteratorToSingular(args0 ...RowElement) Iterator {
	if len(args0) == 0 {
		return NewIterator(reflect.TypeOf(int64(0)), nil)
	}
	var l List = NewList(args0[0].(Iterator).GetType())
	for _, v := range args0 {
		iter := v.(Iterator)
		for iter.HasNext() {
			l.Add(iter.Next())
		}
	}
	return l.Iterator()
}

func __arrayOfArrayToSingular(args0 ...RowElement) Iterator {
	if len(args0) == 0 {
		return NewIterator(reflect.TypeOf(int64(0)), nil)
	}
	var l List = NewList(reflect.TypeOf(args0))
	var t reflect.Type = nil
	for _, v := range args0 {
		var arr RowArray = v.(RowArray)
		for _, val := range arr {
			l.Add(val)
			if t != nil && v != nil {
				t = reflect.TypeOf(val)
			}
		}
	}
	ret := NewList(t)
	ret.AddCollection(l)
	return ret.Iterator()
}

func __arrayToSingular(args0 ...RowElement) Iterator {
	if len(args0) == 0 {
		return NewIterator(reflect.TypeOf(int64(0)), nil)
	}
	var l List = NewList(reflect.TypeOf(args0[0]))
	for _, val := range args0 {
		l.Add(val)
	}
	return l.Iterator()
}

var (
	//Function that map one or more List(s) to exploded contained elements, used in types.Stream.FlatMap()
	LIST_TO_SINGULAR_FUNC ArrayArgumentFunc = __listToSingular
	//Function that map one or more Collection(s) to exploded contained elements, used in types.Stream.FlatMap()
	COLLECTION_TO_SINGULAR_FUNC ArrayArgumentFunc = __collectionToSingular
	//Function that map one or more Iterator(s) to exploded contained elements, used in types.Stream.FlatMap()
	ITERATOR_TO_SINGULAR_FUNC ArrayArgumentFunc = __iteratorToSingular
	//Function that map one or more Array of Arrays(s) to exploded contained elements, used in types.Stream.FlatMap()
	ARRAYOFARRAY_TO_SINGULAR ArrayArgumentFunc = __arrayOfArrayToSingular
	//Function that map one or more Array of Elements to exploded contained elements, used in types.Stream.FlatMap()
	ARRAY_TO_SINGULAR_FUNC ArrayArgumentFunc = __arrayToSingular
)

//Describe Stream, dynamic approach to the Collections, within implem,ented dynamic filters, mapping (re-map types and values in the Collections), reductions (reducing values), summarize methods
type Stream interface {
	// Define parallel stream computation, accordingly to max number of CPUs
	// Returns:
	//    types.Stream Parallel Stream
	Parallel() Stream
	// Varify if current is a Parallel Stream
	// Returns:
	//    bool Parallel Stream Status
	IsParallel() bool
	// Map elements to other ones, maintaining the same number of elements in the Stream
	// Parameters:
	//    mapper (types.MapperFunc) Function that map one element to another one
	// Returns:
	//    types.Stream New Elements Stream
	Map(mapper MapperFunc) Stream
	// Map elements to other flat ones (when elements are Collections, Arrays, Or Plurals generally of elements), maintaining the same number of elements in the Stream
	// Parameters:
	//    mapper (types.ArrayArgumentFunc) Function that map a plural element to a singular one
	// Returns:
	//    types.Stream New Elements Stream
	// Raises:
	//    panic in case of not plural element list
	FlatMap(mapper ArrayArgumentFunc) Stream
	// Reduce elements to other flat ones (when elements are Collections, Arrays, Or Plurals generally of elements), maintaining the same number of elements in the Stream
	// Parameters:
	//    mapper (types.ArrayArgumentFunc) Function that map a plural element to a singular one
	// Returns:
	//    types.Stream New Elements Stream
	Reduce(accumulationElement RowElement, reducer BiArgumentsFunc) RowElement
	// Return a new List within stream content
	// Returns:
	//    types.List New Elements List instance
	ToList() List
	// Return a new Collection within stream content
	// Returns:
	//    types.Collection New Elements Collection instance
	ToCollection() Collection
	// Return a new Iterator within stream content
	// Returns:
	//    types.Iterator New Elements Iterator instance
	Iterator() Iterator
	// Return elements count
	// Returns:
	//    *big.Int Count of stream elements
	Count() *big.Int
	// Return elements average
	// Returns:
	//    *big.Float Average of stream elements
	Average() *big.Float
}

type __streamStruct struct {
	__root     *CollectionElement
	__parallel bool
	__type     reflect.Type
	__count    big.Int
	__sum      big.Int
	__average  big.Float
	__cpus     int
}

func __cloneCollectionElementDescending(element *CollectionElement) *CollectionElement {
	if element == nil {
		return nil
	}
	var root *CollectionElement = __makeBaseCollectionElement(element.Element)
	var previous *CollectionElement = root
	var elem *CollectionElement = element.Next
	for elem != nil {
		newElem := __makeBaseCollectionElement(elem.Element)
		previous.Next = newElem
		previous = newElem
		elem = elem.Next
	}
	return root
}

func (stream *__streamStruct) Parallel() Stream {
	return __newStream(stream.__type, __cloneCollectionElementDescending(stream.__root), true)
}

func (stream *__streamStruct) IsParallel() bool {
	return stream.__parallel
}

func (stream *__streamStruct) Map(mapper MapperFunc) Stream {
	var t reflect.Type = nil
	if stream.__root == nil {
		return __newStream(stream.__type, stream.__root, stream.__parallel)
	}
	var root *CollectionElement = __cloneCollectionElementDescending(stream.__root)
	root.Element = mapper(root.Element)
	if root.Element != nil {
		t = reflect.TypeOf(root.Element)
	}
	elem := root.Next
	for elem != nil {
		elem.Element = mapper(elem.Element)
		if t == nil && elem.Element != nil {
			t = reflect.TypeOf(elem.Element)
		}
		elem = elem.Next
	}
	if t == nil {
		t = stream.__type
	}
	return __newStream(t, root, stream.__parallel)
}

func __checkPlural(element RowElement) {
	if element == nil {
		return
	}
	t := reflect.TypeOf(element)
	switch t.Kind() {
	case reflect.Array:
		return
	case reflect.Slice:
		return
	case reflect.Map:
		return
	case reflect.Interface:
		if t.Name() == reflect.TypeOf(NewList(t)).Name() {
			return
		} else if t.Name() == reflect.TypeOf(NewCollection(t)).Name() {
			return
		} else if t.Name() == reflect.TypeOf(NewIterator(t, nil)).Name() {
			return
		}
	default:
	}
	panic(fmt.Sprintf("Element of type <%v> is not plural!!", t))
}

func (stream *__streamStruct) FlatMap(mapper ArrayArgumentFunc) Stream {
	if stream.__root == nil {
		return __newStream(stream.__type, stream.__root, stream.__parallel)
	}
	var t reflect.Type = nil
	var root *CollectionElement = __cloneCollectionElementDescending(stream.__root)
	__checkPlural(root.Element)
	var rootElements Iterator = mapper(root.Element)
	for !rootElements.HasNext() {
		root = root.Next
		if root == nil {
			return __newStream(stream.__type, root, stream.__parallel)
		}
		rootElements = mapper(root.Element)
	}
	rootNext := root.Next
	var rootElem *CollectionElement = nil
	for rootElements.HasNext() {
		newElem := __makeBaseCollectionElement(rootElements.Next())
		if rootElem != nil {
			rootElem.Next = newElem
		} else {
			root = newElem
		}
		rootElem = newElem
		if t == nil && newElem.Element != nil {
			t = reflect.TypeOf(newElem.Element)
		}
	}
	rootElem.Next = rootNext
	prev := rootElem
	elem := rootElem.Next
	for elem != nil {
		__checkPlural(elem.Element)
		var elements Iterator = mapper(elem.Element)
		//Remove Previous Element and use new ones
		next := elem.Next
		if !elements.HasNext() {
			// Remove Element that has no elements in the list and set
			// previous as current, so next will be linked to the previous
			// and we continue with the next
			prev.Next = next
			elem = prev
		} else {
			// Select previous as current, and start attaching to that first element
			// of the new array, so use that element to link the next in the elements list
			// sequence
			elem = prev
			for elements.HasNext() {
				newElem := __makeBaseCollectionElement(elements.Next())
				elem.Next = newElem
				elem = newElem
				if t == nil && newElem.Element != nil {
					t = reflect.TypeOf(newElem.Element)
				}
			}
			//Connect next sequence to next last element of the list
			elem.Next = next
		}
		prev = elem
		elem = next
	}
	if t == nil {
		t = stream.__type
	}
	return __newStream(t, root, stream.__parallel)
}

func (stream *__streamStruct) Reduce(accumulationElement RowElement, reducer BiArgumentsFunc) RowElement {
	if stream.__root == nil {
		return __newStream(stream.__type, stream.__root, stream.__parallel)
	}
	var root *CollectionElement = __cloneCollectionElementDescending(stream.__root)
	accumulationElement = reducer(accumulationElement, root.Element)
	elem := root.Next
	for elem != nil {
		accumulationElement = reducer(accumulationElement, elem.Element)
		elem = elem.Next
	}
	return accumulationElement
}

func (stream *__streamStruct) ToList() List {
	return IteratorAsList(NewIterator(stream.__type, __cloneCollectionElementDescending(stream.__root)))
}

func (stream *__streamStruct) ToCollection() Collection {
	return IteratorToCollection(NewIterator(stream.__type, __cloneCollectionElementDescending(stream.__root)))
}

func (stream *__streamStruct) Iterator() Iterator {
	return NewIterator(stream.__type, stream.__root)
}

func (stream *__streamStruct) Count() *big.Int {
	return &stream.__count
}

func (stream *__streamStruct) Average() *big.Float {
	return &stream.__average
}

func __calculateStreamBaseInfo(element *CollectionElement) (big.Int, big.Int, big.Float) {
	var ONE *big.Int = big.NewInt(int64(1))
	var count *big.Int = big.NewInt(int64(0))
	var sum *big.Int = big.NewInt(int64(0))
	var average *big.Float = big.NewFloat(float64(0))
	defer func() {
		var message string = ""
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				message = fmt.Sprintf("Unable to calculate base, error : %s!!!", itf.(error).Error())
			} else {
				message = fmt.Sprintf("Unable to calculate base, error : %v!!!", itf)
			}
			if logger != nil {
				logger.ErrorS(message)
			} else {
				fmt.Println(message)
			}
			count = big.NewInt(int64(0))
			sum = big.NewInt(int64(0))
			average = big.NewFloat(float64(0))
		}
	}()
	if element != nil {
		count = count.Add(count, ONE)
		i1, err1 := strconv.ParseInt(fmt.Sprintf("%v", element.Element), 10, 64)
		if err1 != nil {
			i1 = int64(0)
		}
		sum = sum.Add(sum, big.NewInt(i1))
		elem := element.Next
		for elem != nil {
			count = count.Add(count, ONE)
			i2, err2 := strconv.ParseInt(fmt.Sprintf("%v", elem.Element), 10, 64)
			if err2 != nil {
				i2 = int64(0)
			}
			sum = sum.Add(sum, big.NewInt(i2))
			elem = elem.Next
		}
	}
	sumFloat, _, _ := big.ParseFloat(sum.String(), 10, 10, big.ToNearestAway)
	countFloat, _, _ := big.ParseFloat(count.String(), 10, 10, big.ToNearestAway)
	average = sumFloat.Quo(sumFloat, countFloat)
	return *count, *sum, *average
}

func __newStream(sType reflect.Type, root *CollectionElement, parallel bool) Stream {
	count, sum, average := __calculateStreamBaseInfo(root)
	return &__streamStruct{
		__root:     root,
		__parallel: parallel,
		__type:     sType,
		__count:    count,
		__sum:      sum,
		__average:  average,
		__cpus:     runtime.NumCPU(),
	}
}

//Retrieve A Stream related to provided array
// Returns:
//   types.Stream New Data Stream
func ArrayAsStream(arr RowArray) Stream {
	sType, root := ArrayToCollectionElement(arr)
	count, sum, average := __calculateStreamBaseInfo(root)
	return &__streamStruct{
		__root:     root,
		__parallel: false,
		__type:     sType,
		__count:    count,
		__sum:      sum,
		__average:  average,
		__cpus:     runtime.NumCPU(),
	}
}

//Retrieve A Stream related to provided array, with parallelized computation
// Returns:
//   types.Stream New Parallel Data Stream
func ArrayAsParallelStream(arr RowArray) Stream {
	sType, root := ArrayToCollectionElement(arr)
	count, sum, average := __calculateStreamBaseInfo(root)
	return &__streamStruct{
		__root:     root,
		__parallel: true,
		__type:     sType,
		__count:    count,
		__sum:      sum,
		__average:  average,
		__cpus:     runtime.NumCPU(),
	}
}

//Retrieve A Stream related to provided Iterator data
// Returns:
//   types.Stream New Data Stream
func IteratorAsStream(iter Iterator) Stream {
	var arr []RowElement
	for iter.HasNext() {
		arr = append(arr, iter.Next())
	}
	return ArrayAsStream(RowArray(arr))
}
