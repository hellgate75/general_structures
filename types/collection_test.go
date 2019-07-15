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

func __createBaseSecondCollection() Collection {
	var elements []RowElement
	elements = append(elements, RowElement(6))
	elements = append(elements, RowElement(7))
	var coll Collection = NewCollection(__createBaseType())
	coll.AddAll(elements)
	return coll
}

func TestNewCollection(t *testing.T) {
	collection := NewCollection(__createBaseType())
	collection.Add(__createBaseDataElement())
	collection.AddAll(__createBaseDataArray())
	if collection.Size() != int64(6) {
		t.Fatal(fmt.Sprintf("Invalid Collection size %x, instead of %x", collection.Size(), 6))
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
	if collection.Size() != int64(5) {
		t.Fatal(fmt.Sprintf("Invalid Collection size %x, instead of %x", collection.Size(), 5))
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
	if collection.Size() != int64(5) {
		t.Fatal(fmt.Sprintf("Invalid Collection size %x, instead of %x", collection.Size(), 5))
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
	if collection.Size() != int64(7) {
		t.Fatal(fmt.Sprintf("Invalid Collection size %x, instead of %x", collection.Size(), 7))
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
	if collection.Size() != int64(5) {
		t.Fatal(fmt.Sprintf("Invalid Collection size %x, instead of %x", collection.Size(), 5))
	}
}
