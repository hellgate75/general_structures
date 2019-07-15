package types

import (
	"fmt"
	"reflect"
	"testing"
)

func __createBaseCollectionElement() *CollectionElement {
	fourthLevel := CollectionElement{
		Element: RowElement(4),
		Next:    nil,
	}
	tirdLevel := CollectionElement{
		Element: RowElement(2),
		Next:    &fourthLevel,
	}
	secondLevel := CollectionElement{
		Element: RowElement(3),
		Next:    &tirdLevel,
	}
	firstLevel := CollectionElement{
		Element: RowElement(5),
		Next:    &secondLevel,
	}
	root := CollectionElement{
		Element: RowElement(1),
		Next:    &firstLevel,
	}
	return &root
}

func TestNewIterator(t *testing.T) {
	v := RowElement(0)
	tp := reflect.ValueOf(v).Type()
	iter := NewIterator(tp, __createBaseCollectionElement())
	if iter.HasNext() {
		current := iter.Next()
		expected := RowElement(1)
		if current != expected {
			t.Fatal(fmt.Sprintf("Wrong value at level <%x> - Given : <%v> Expected <%v>", 0, current, expected))
		}
	} else {
		t.Fatal(fmt.Sprintf("Unexpected unavailable element at level <%x>", 0))
	}
	if iter.HasNext() {
		current := iter.Next()
		expected := RowElement(5)
		if current != expected {
			t.Fatal(fmt.Sprintf("Wrong value at level <%x> - Given : <%v> Expected <%v>", 1, current, expected))
		}
	} else {
		t.Fatal(fmt.Sprintf("Unexpected unavailable element at level <%x>", 1))
	}
	if iter.HasNext() {
		current := iter.Next()
		expected := RowElement(3)
		if current != expected {
			t.Fatal(fmt.Sprintf("Wrong value at level <%x> - Given : <%v> Expected <%v>", 2, current, expected))
		}
	} else {
		t.Fatal(fmt.Sprintf("Unexpected unavailable element at level <%x>", 2))
	}
	if iter.HasNext() {
		current := iter.Next()
		expected := RowElement(2)
		if current != expected {
			t.Fatal(fmt.Sprintf("Wrong value at level <%x> - Given : <%v> Expected <%v>", 3, current, expected))
		}
	} else {
		t.Fatal(fmt.Sprintf("Unexpected unavailable element at level <%x>", 3))
	}
	if iter.HasNext() {
		current := iter.Next()
		expected := RowElement(4)
		if current != expected {
			t.Fatal(fmt.Sprintf("Wrong value at level <%x> - Given : <%v> Expected <%v>", 4, current, expected))
		}
	} else {
		t.Fatal(fmt.Sprintf("Unexpected unavailable element at level <%x>", 4))
	}
	if iter.HasNext() {
		t.Fatal(fmt.Sprintf("Unexpected available element at not existing level <%x>", 5))
	}
}
