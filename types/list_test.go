package types

import (
	"fmt"
	"testing"
)

func __createNewComparator() ComparatorFunc {
	comparator := func(first RowElement, second RowElement) int {
		return first.(int) - second.(int)
	}
	return ComparatorFunc(comparator)
}

func __createNewMatcher() MatcherFunc {
	comparator := func(first RowElement, second RowElement) bool {
		return first.(int) == second.(int)
	}
	return MatcherFunc(comparator)
}

func TestNewList(t *testing.T) {
	list := NewList(__createBaseType())
	list.AddCollection(__createBaseCollection())
	list.Add(__createBaseDataElement())
	__runListTestCases(list, t)
}

func __createBaseListDataArray() RowArray {
	var elements []RowElement
	elements = append(elements, RowElement(0))
	elements = append(elements, RowElement(1))
	elements = append(elements, RowElement(5))
	elements = append(elements, RowElement(3))
	elements = append(elements, RowElement(2))
	elements = append(elements, RowElement(4))
	return RowArray(elements)
}

func __createBaseListCollection() Collection {
	var coll Collection = NewCollection(__createBaseType())
	coll.AddAll(__createBaseListDataArray())
	return coll
}
func TestNewListWithArray(t *testing.T) {
	list := NewListWithArray(__createBaseType(), __createBaseListDataArray())
	__runListTestCases(list, t)
}

func TestCollectionAsList(t *testing.T) {
	list := CollectionAsList(__createBaseListCollection())
	__runListTestCases(list, t)
}

func __createBaseListCollectionElement() *CollectionElement {
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
	baseLevel := CollectionElement{
		Element: RowElement(1),
		Next:    &firstLevel,
	}
	root := CollectionElement{
		Element: RowElement(0),
		Next:    &baseLevel,
	}
	return &root
}

func TestIteratorAsList(t *testing.T) {
	list := IteratorAsList(NewIterator(__createBaseType(), __createBaseListCollectionElement()))
	__runListTestCases(list, t)
}

func __runListTestCases(list List, t *testing.T) {
	if list.Size() != int64(6) {
		t.Fatal(fmt.Sprintf("Invalid List size %x, instead of %x", list.Size(), 6))
	}
	if !list.Iterator().HasNext() {
		t.Fatal("Empty iterator found!!!")
	}
	if list.First() == nil {
		t.Fatal("First element is nil!!!")
	}
	if list.Last() == nil {
		t.Fatal("First element is nil!!!")
	}
	if list.GetType() != __createBaseType() {
		t.Fatal(fmt.Sprintf("List type is invalid, Expected <%v> But Given <%v>", __createBaseType(), list.GetType()))
	}
	list.Remove(__createBaseDataElement())
	if list.Size() != int64(5) {
		t.Fatal(fmt.Sprintf("Invalid List size %x, instead of %x", list.Size(), 5))
	}
	list.RemoveAll(RowElement(5))
	if list.Size() != int64(4) {
		t.Fatal(fmt.Sprintf("Invalid List size %x, instead of %x", list.Size(), 4))
	}
	list.Add(RowElement(5))
	if list.Size() != int64(5) {
		t.Fatal(fmt.Sprintf("Invalid List size %x, instead of %x", list.Size(), 5))
	}
	changes := list.Sort(__createNewComparator())
	if !changes {
		t.Fatal("List was not sorted, so a sort is expected!!")
	}
	iter := list.Iterator()
	if !iter.HasNext() {
		t.Fatal(fmt.Sprintf("Unavailable value at level <%x>!!", 0))
	}
	if n := iter.Next(); n != RowElement(1) {
		t.Fatal(fmt.Sprintf("Wrong value at level <%x>, Expected <%v>, But Given <%v> !!", 0, RowElement(1), n))
	}
	if !iter.HasNext() {
		t.Fatal(fmt.Sprintf("Unavailable value at level <%x>!!", 1))
	}
	if n := iter.Next(); n != RowElement(2) {
		t.Fatal(fmt.Sprintf("Wrong value at level <%x>, Expected <%v>, But Given <%v> !!", 1, RowElement(2), n))
	}
	if !iter.HasNext() {
		t.Fatal(fmt.Sprintf("Unavailable value at level <%x>!!", 2))
	}
	if n := iter.Next(); n != RowElement(3) {
		t.Fatal(fmt.Sprintf("Wrong value at level <%x>, Expected <%v>, But Given <%v> !!", 2, RowElement(3), n))
	}
	if !iter.HasNext() {
		t.Fatal(fmt.Sprintf("Unavailable value at level <%x>!!", 3))
	}
	if n := iter.Next(); n != RowElement(4) {
		t.Fatal(fmt.Sprintf("Wrong value at level <%x>, Expected <%v>, But Given <%v> !!", 3, RowElement(4), n))
	}
	if !iter.HasNext() {
		t.Fatal(fmt.Sprintf("Unavailable value at level <%x>!!", 4))
	}
	if n := iter.Next(); n != RowElement(5) {
		t.Fatal(fmt.Sprintf("Wrong value at level <%x>, Expected <%v>, But Given <%v> !!", 4, RowElement(5), n))
	}
	match1 := list.Contains(RowElement(5))
	if !match1 {
		t.Fatal(fmt.Sprintf("Element <%v> is contained in the list: default matcher doesn't work!!", RowElement(5)))
	}
	match2 := list.ContainsAs(RowElement(5), __createNewMatcher())
	if !match2 {
		t.Fatal(fmt.Sprintf("Element <%v> is contained in the list: custom matcher doesn't work!!", RowElement(5)))
	}
	value1 := list.Get(1)
	expectedValue1 := RowElement(2)
	if value1 != expectedValue1 {
		t.Fatal(fmt.Sprintf("Wrong value Getting at index <%x>, Expected <%v>, But Given <%v> !!", 1, expectedValue1, value1))
	}
	settedValue1 := RowElement(6)
	settedUp1 := list.Set(1, settedValue1)
	if !settedUp1 {
		t.Fatal(fmt.Sprintf("Wrong outcome Setting Element at index <%x>, From <%v>, To <%v> !!", 1, value1, settedValue1))
	}
	value2 := list.Get(1)
	expectedValue2 := settedValue1
	if value2 != expectedValue2 {
		t.Fatal(fmt.Sprintf("Wrong value Getting at index <%x>, Expected <%v>, But Given <%v> !!", 1, expectedValue2, value2))
	}
	addedValue1 := RowElement(9)
	added1 := list.AddAfter(1, addedValue1)
	if !added1 {
		t.Fatal(fmt.Sprintf("Wrong outcome Adding Element after index <%x>, After Value <%v>, With Value <%v> !!", 1, value2, addedValue1))
	}
	value3 := list.Get(2)
	expectedValue3 := addedValue1
	if value3 != expectedValue3 {
		t.Fatal(fmt.Sprintf("Wrong value Getting at index <%x>, Expected <%v>, But Given <%v> !!", 1, expectedValue3, value3))
	}
	expectedValue4 := RowElement(9)
	removed1 := list.RemoveAt(1)
	if !removed1 {
		t.Fatal(fmt.Sprintf("Wrong outcome Removing Element after index <%x>, Removing Value <%v>, Replacing With <%v> !!", 1, value2, expectedValue4))
	}
	value4 := list.Get(1)
	if value4 != expectedValue4 {
		t.Fatal(fmt.Sprintf("Wrong value Removing at index <%x>, Expected <%v>, But Given <%v> !!", 1, expectedValue4, value4))
	}
	listAddingAll1 := __createBaseSecondDataArray()
	listAddingAllLen1 := int64(len(listAddingAll1))
	added2 := list.AddAllAfter(1, listAddingAll1)
	if listAddingAllLen1 != added2 {
		t.Fatal(fmt.Sprintf("Wrong value for Number of Added Elements at index <%x>, Expected <%v>, But Given <%v> !!", 1, listAddingAllLen1, added2))
	}
	expectedValue5 := listAddingAll1[0]
	value5 := list.Get(2)
	if expectedValue5 != value5 {
		t.Fatal(fmt.Sprintf("Wrong value of Added Element taken at index <%x>, Expected <%v>, But Given <%v> !!", 2, expectedValue5, value5))
	}
	expectedValue6 := listAddingAll1[1]
	value6 := list.Get(3)
	if expectedValue6 != value6 {
		t.Fatal(fmt.Sprintf("Wrong value of Added Element taken at index <%x>, Expected <%v>, But Given <%v> !!", 3, expectedValue6, value6))
	}
	listAddingAll2 := __createBaseSecondCollection()
	listAddingAllLen2 := int64(len(listAddingAll1))
	added3 := list.AddCollectionAfter(3, listAddingAll2)
	if listAddingAllLen2 != added3 {
		t.Fatal(fmt.Sprintf("Wrong value for Number of Added Collection Elements at index <%x>, Expected <%v>, But Given <%v> !!", 1, listAddingAllLen2, added3))
	}
	expectedValue7 := listAddingAll1[0]
	value7 := list.Get(4)
	if expectedValue7 != value7 {
		t.Fatal(fmt.Sprintf("Wrong value of Added Collection Element taken at index <%x>, Expected <%v>, But Given <%v> !!", 4, expectedValue7, value7))
	}
	expectedValue8 := listAddingAll1[1]
	value8 := list.Get(5)
	if expectedValue8 != value8 {
		t.Fatal(fmt.Sprintf("Wrong value of Added Collection Element taken at index <%x>, Expected <%v>, But Given <%v> !!", 5, expectedValue8, value8))
	}
	expectedString := "List{size: '9', sampleData: <1 9 6 7 6 ...>}"
	listString := fmt.Sprintf("%v", list)
	if listString != expectedString {
		t.Fatal(fmt.Sprintf("List String() value is wrong, Expected <%v> But Given <%v>", expectedString, listString))
	}

}
