package types

import (
	"reflect"
)

//Tranforms an array to types.CollectionElement, used for filling internally Collections and derivates
// Parameters:
//    arr (type.RowArray) Array of data
// Returns:
//   (reflect.Type List item Type,
//   *types.CollectionElement Pointer to Root of new Data Sequence)
func ArrayToCollectionElement(arr RowArray) (reflect.Type, *CollectionElement) {
	var t reflect.Type = nil
	var last *CollectionElement = nil
	for i := len(arr) - 1; i >= 0; i-- {
		newElem := __makeBaseCollectionElement(arr[i])
		if t == nil && newElem.Element != nil {
			t = reflect.TypeOf(newElem.Element)
		}
		newElem.Next = last
		last = newElem
	}
	return t, last
}
