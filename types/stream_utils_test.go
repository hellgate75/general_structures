package types

import (
	"fmt"
	"testing"
)

func TestArrayToCollectionElement(t *testing.T) {
	var array RowArray = __createArrayInRange(50, 0)
	sType, root := ArrayToCollectionElement(array)
	expectedTypeString := "int64"
	typeString := sType.String()
	if expectedTypeString != typeString {
		t.Fatal(fmt.Sprintf("ArrayToCollectionElement output type is invalid, Expected <%s> But Given <%s>", expectedTypeString, typeString))
	}
	size := __newStream(sType, root, false).Count().Int64()
	expectedSize := int64(50)
	if expectedSize != size {
		t.Fatal(fmt.Sprintf("ArrayToCollectionElement output type is invalid, Expected <%d> But Given <%d>", expectedSize, size))
	}
	firstValue := root.Element
	expctedFirstValue := array[0]
	if expctedFirstValue != firstValue {
		t.Fatal(fmt.Sprintf("ArrayToCollectionElement first value is invalid, Expected <%v> But Given <%v>", expctedFirstValue, firstValue))
	}
	last := root
	for last.Next != nil {
		last = last.Next
	}
	lastValue := last.Element
	expctedLastValue := array[len(array)-1]
	if expctedLastValue != lastValue {
		t.Fatal(fmt.Sprintf("ArrayToCollectionElement last value is invalid, Expected <%v> But Given <%v>", expctedLastValue, lastValue))
	}
}
