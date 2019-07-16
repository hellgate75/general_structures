package types

import (
	"fmt"
	"testing"
)

const (
	__MAX_TEST_STREAM_ITEMS         = 1000
	__MAX_ELEMENTS_TEST_STREAM_ITEM = 100
)

func __createHugeElementsPlain() *CollectionElement {
	var root *CollectionElement = nil
	var lastCreated *CollectionElement = nil
	for i := __MAX_TEST_STREAM_ITEMS - 1; i >= 0; i-- {
		newElem := __makeBaseCollectionElement(RowElement(int64(i + 1)))
		newElem.Next = lastCreated
		if i == 0 {
			root = newElem
		}
		lastCreated = newElem
	}
	return root
}

func __createArrayInRange(max int, diff int) RowArray {
	out := make([]RowElement, max)
	for i := 0; i < max; i++ {
		out[i] = RowElement(int64(i + 1 + diff))
	}
	return RowArray(out)
}

func __createHugeElementsArray() *CollectionElement {
	var root *CollectionElement = nil
	var lastCreated *CollectionElement = nil
	for i := __MAX_TEST_STREAM_ITEMS - 1; i >= 0; i-- {
		newElem := __makeBaseCollectionElement(RowElement(__createArrayInRange(__MAX_ELEMENTS_TEST_STREAM_ITEM, __MAX_ELEMENTS_TEST_STREAM_ITEM*i)))
		newElem.Next = lastCreated
		if i == 0 {
			root = newElem
		}
		lastCreated = newElem
	}
	return root
}

func TestNewStream(t *testing.T) {
	stream := __newStream(__createBaseType(), __createHugeElementsPlain(), false)
	expectedCount := int64(__MAX_TEST_STREAM_ITEMS)
	count := stream.Count().Int64()
	if expectedCount != count {
		t.Fatal(fmt.Sprintf("Wrong count of the stream, Expected <%d>, But Given <%d> !!", expectedCount, count))
	}
	reducer := BiArgumentsFunc(func(arg0 RowElement, arg1 RowElement) RowElement {
		var accumulator int64 = arg0.(int64)
		var value int64 = arg1.(int64)
		return accumulator + value
	})
	valueReduced := stream.Reduce(int64(0), reducer)
	expectedValueReduced := int64(500500)
	if expectedValueReduced != valueReduced {
		t.Fatal(fmt.Sprintf("Wrong reduced element of the stream, Expected <%d>, But Given <%d> !!", expectedValueReduced, valueReduced))
	}
	valueAverage, _ := stream.Average().Float64()
	expectedValueAverage := float64(500.5)
	if expectedValueAverage != valueAverage {
		t.Fatal(fmt.Sprintf("Wrong average in elements of the stream, Expected <%f>, But Given <%f> !!", expectedValueAverage, valueAverage))
	}
	mapper := MapperFunc(func(arg0 RowElement) RowElement {
		return fmt.Sprintf("%v", arg0)
	})
	newStream := stream.Map(mapper)
	newCount := newStream.Count().Int64()
	if count != newCount {
		t.Fatal(fmt.Sprintf("Wrong count in elements mapping to another stream, Expected <%d>, But Given <%d> !!", count, newCount))
	}
	newList := newStream.ToList()
	newListCount := newList.Size()
	if count != newListCount {
		t.Fatal(fmt.Sprintf("Wrong count in elements mapping stream to List, Expected <%d>, But Given <%d> !!", count, newListCount))
	}
	expectedType := "string"
	newListType := newList.GetType().String()
	if expectedType != newListType {
		t.Fatal(fmt.Sprintf("Wrong type in stream to List, Expected <%s>, But Given <%s> !!", expectedType, newListType))
	}
	listIterator := newList.Iterator()
	iterator := newStream.Iterator()
	valIter01 := listIterator.Next()
	valIter02 := iterator.Next()
	if valIter01 != valIter02 {
		t.Fatal(fmt.Sprintf("Wrong values in stream to Iterator, Expected <%v>, But Given <%v> !!", valIter01, valIter02))
	}
	listIterator = newList.Iterator()
	newStreamColl := newStream.ToCollection()
	valColl01 := listIterator.Next()
	valColl02 := newStreamColl.Iterator().Next()
	if valColl01 != valColl02 {
		t.Fatal(fmt.Sprintf("Wrong values in stream to Collection, Expected <%v>, But Given <%v> !!", valColl01, valColl02))
	}
}

func TestNewStreamFlatMap(t *testing.T) {
	stream := __newStream(__createBaseType(), __createHugeElementsArray(), false)
	expectedCount := int64(__MAX_TEST_STREAM_ITEMS)
	count := stream.Count().Int64()
	if expectedCount != count {
		t.Fatal(fmt.Sprintf("Wrong count of the stream made of Array of Arrays, Expected <%d>, But Given <%d> !!", expectedCount, count))
	}
	newStream := stream.FlatMap(ARRAYOFARRAY_TO_SINGULAR)
	explodedArraySize := newStream.Count().Int64()
	expectedArraySize := int64(__MAX_TEST_STREAM_ITEMS * __MAX_ELEMENTS_TEST_STREAM_ITEM)
	if expectedArraySize != explodedArraySize {
		t.Fatal(fmt.Sprintf("Wrong count of the stream made of flatMap Sream, Expected <%d>, But Given <%d> !!", expectedArraySize, explodedArraySize))
	}
}
