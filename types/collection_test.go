package types

import (
	"fmt"
	"reflect"
	"testing"
)

func __createBaseType() reflect.Type {
	i := RowElement(0)
	return reflect.ValueOf(i).Type()
}

func __createBaseDataElement() RowElement {
	return RowElement(0)
}

func __createBaseIterator() Iterator {
	return NewIterator(__createBaseType(), __createBaseCollectionElement())
}

func __createBaseDataArray() RowArray {
	var elements []RowElement
	elements = append(elements, RowElement(1))
	elements = append(elements, RowElement(5))
	elements = append(elements, RowElement(3))
	elements = append(elements, RowElement(2))
	elements = append(elements, RowElement(4))
	return RowArray(elements)
}

func __createBaseCollection() Collection {
	var coll Collection = NewCollection(__createBaseType())
	coll.AddAll(__createBaseDataArray())
	return coll
}

func __createBaseSecondDataArray() RowArray {
	var elements []RowElement
	elements = append(elements, RowElement(6))
	elements = append(elements, RowElement(7))
	return RowArray(elements)
}

func __createBaseSecondCollection() Collection {
	var coll Collection = NewCollection(__createBaseType())
	coll.AddAll(__createBaseSecondDataArray())
	return coll
}

func TestNewCollection(t *testing.T) {
	collection := NewCollection(__createBaseType())
	collection.Add(__createBaseDataElement())
	collection.AddAll(__createBaseDataArray())
	if collection.Empty() {
		t.Fatal("It is Invalid that Collection is Empty!!!")
	}
	if !collection.Iterator().HasNext() {
		t.Fatal("Empty iterator found!!!")
	}
	if collection.First() == nil {
		t.Fatal("First element is nil!!!")
	}
	if collection.Last() == nil {
		t.Fatal("First element is nil!!!")
	}
	if collection.GetType() != __createBaseType() {
		t.Fatal(fmt.Sprintf("Collection type is invalid, Expected <%v> But Given <%v>", __createBaseType(), collection.GetType()))
	}
	collection.Remove(__createBaseDataElement())
	if collection.Empty() {
		t.Fatal("It is Invalid that Collection is Empty!!!")
	}
}

func TestNewCollectionWithArray(t *testing.T) {
	collection := NewCollectionWithArray(__createBaseType(), __createBaseDataArray())
	if !collection.Iterator().HasNext() {
		t.Fatal("Empty iterator found!!!")
	}
	if collection.First() == nil {
		t.Fatal("First element is nil!!!")
	}
	if collection.Last() == nil {
		t.Fatal("First element is nil!!!")
	}
	if collection.GetType() != __createBaseType() {
		t.Fatal(fmt.Sprintf("Collection type is invalid, Expected <%v> But Given <%v>", __createBaseType(), collection.GetType()))
	}
	if collection.Empty() {
		t.Fatal("It is Invalid that Collection is Empty!!!")
	}
	expectedString := "Collection{size: '5', sampleData: <1 5 3 2 4 >}"
	collectionString := fmt.Sprintf("%v", collection)
	if collectionString != expectedString {
		t.Fatal(fmt.Sprintf("Collection type is invalid, Expected <%v> But Given <%v>", expectedString, collectionString))
	}
}

func TestCloneCollection(t *testing.T) {
	collection := CloneCollection(__createBaseCollection())
	collection.AddCollection(__createBaseSecondCollection())
	if !collection.Iterator().HasNext() {
		t.Fatal("Empty iterator found!!!")
	}
	if collection.First() == nil {
		t.Fatal("First element is nil!!!")
	}
	if collection.Last() == nil {
		t.Fatal("First element is nil!!!")
	}
	if collection.GetType() != __createBaseType() {
		t.Fatal(fmt.Sprintf("Collection type is invalid, Expected <%v> But Given <%v>", __createBaseType(), collection.GetType()))
	}
	if collection.Empty() {
		t.Fatal("It is Invalid that Collection is Empty!!!")
	}
	expectedString := "Collection{size: '7', sampleData: <1 5 3 2 4 ...>}"
	collectionString := fmt.Sprintf("%v", collection)
	if collectionString != expectedString {
		t.Fatal(fmt.Sprintf("Collection type is invalid, Expected <%v> But Given <%v>", expectedString, collectionString))
	}
	collSize := collection.Stream().Count().Int64()
	expectedSize := int64(7)
	if expectedSize != collSize {
		t.Fatal(fmt.Sprintf("Collection stream size is wrong, Expected <%v> But Given <%v>", expectedSize, collSize))
	}
}

func TestIteratorToCollection(t *testing.T) {
	collection := IteratorToCollection(__createBaseIterator())
	if !collection.Iterator().HasNext() {
		t.Fatal("Empty iterator found!!!")
	}
	if collection.First() == nil {
		t.Fatal("First element is nil!!!")
	}
	if collection.Last() == nil {
		t.Fatal("First element is nil!!!")
	}
	if collection.GetType() != __createBaseType() {
		t.Fatal(fmt.Sprintf("Collection type is invalid, Expected <%v> But Given <%v>", __createBaseType(), collection.GetType()))
	}
	if collection.Empty() {
		t.Fatal("It is Invalid that Collection is Empty!!!")
	}
	collection.Clear()
	if !collection.Empty() {
		t.Fatal("It is Invalid that Collection is not Empty!!!")
	}
	expectedFirst := RowElement(nil)
	first := collection.First()
	if first != expectedFirst {
		t.Fatal(fmt.Sprintf("Collection first value after Clear() is invalid, Expected <%v> But Given <%v>", expectedFirst, first))
	}
	expectedLast := RowElement(nil)
	last := collection.Last()
	if last != expectedLast {
		t.Fatal(fmt.Sprintf("Collection last value after Clear() is invalid, Expected <%v> But Given <%v>", expectedLast, last))
	}
	expectedString := "Collection{size: '0', sampleData: <>}"
	collectionString := fmt.Sprintf("%v", collection)
	if collectionString != expectedString {
		t.Fatal(fmt.Sprintf("Collection String() value is wrong, Expected <%v> But Given <%v>", expectedString, collectionString))
	}
}
