package logic

import (
	"fmt"
	"reflect"
	"testing"
)

type TestStruct1 struct {
	Name string
	Age  int
}

type TestStruct2 struct {
	TestStruct1
	Surname string
}

type TestStruct3 struct {
	Name    string
	Surname string
}

func __createTestBusinessObject() BusinessObject {
	return NewBusinessObject(__createTestBOMap(), reflect.TypeOf(TestStruct1{}))
}

func __createTestErrorBusinessObject() BusinessObject {
	return NewBusinessObject(__createTestErrorBOMap(), reflect.TypeOf(TestStruct1{}))
}

func __createTestBOMap() ValuesMap {
	vm := make(ValuesMap)
	vm["Name"] = "Fabrizio"
	vm["Age"] = 44
	return vm
}

func __createTestErrorBOMap() ValuesMap {
	vm := make(ValuesMap)
	return vm
}

func TestNewBusinessObject(t *testing.T) {
	bo := NewBusinessObject(__createTestBOMap(), reflect.TypeOf(TestStruct1{}))
	var expectedLen int = 2
	var expectedFld1 string = "Name"
	var expectedFld2 string = "Age"
	var mLen int = bo.Size()
	if expectedLen != mLen {
		t.Fatal(fmt.Sprintf("BusinessObject has invalid Size(), Expcted <%d> but Given <%d>", expectedLen, mLen))
	}
	keys := bo.GetKeys()
	var mFld1 string = keys[0]
	var mFld2 string = keys[1]
	if expectedFld1 != mFld1 && expectedFld1 != mFld2 {
		t.Fatal(fmt.Sprintf("BusinessObject has invalid Field, Expected <%s> but Given <%s>", expectedFld1, mFld1))
	}
	if expectedFld2 != mFld2 && expectedFld2 != mFld1 {
		t.Fatal(fmt.Sprintf("BusinessObject has invalid Field, Expected <%s> but Given <%s>", expectedFld2, mFld2))
	}

}

func TestConvertStructure(t *testing.T) {
	tStrc1 := TestStruct1{
		Name: "Fabrizio",
		Age:  44,
	}
	bo, err := ConvertStructure(tStrc1)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error from the structure conversion, error is <%s>", err.Error()))
	}
	var expectedLen int = 2
	var expectedFld1 string = "Name"
	var expectedFld2 string = "Age"
	var mLen int = bo.Size()
	if expectedLen != mLen {
		t.Fatal(fmt.Sprintf("BusinessObject has invalid Size(), Expcted <%d> but Given <%d>", expectedLen, mLen))
	}
	keys := bo.GetKeys()
	var mFld1 string = keys[0]
	var mFld2 string = keys[1]
	if expectedFld1 != mFld1 && expectedFld1 != mFld2 {
		t.Fatal(fmt.Sprintf("BusinessObject has invalid Field, Expcted <%s> but Given <%s>", expectedFld1, mFld1))
	}
	if expectedFld2 != mFld2 && expectedFld2 != mFld1 {
		t.Fatal(fmt.Sprintf("BusinessObject has invalid Field, Expcted <%s> but Given <%s>", expectedFld2, mFld2))
	}
	expectedName := "Fabrizio"
	expectedAge := 44
	mFldName := bo.GetValuesMap()["Name"]
	mFldAge := bo.GetValuesMap()["Age"]
	if expectedName != mFldName {
		t.Fatal(fmt.Sprintf("BusinessObject has invalid Field Value, Expcted <%s> but Given <%s>", expectedName, mFldName))
	}
	if expectedAge != mFldAge {
		t.Fatal(fmt.Sprintf("BusinessObject has invalid Field Value, Expcted <%d> but Given <%d>", expectedAge, mFldAge))
	}
}

func TestErrorsInConvertStructure(t *testing.T) {
	tStrc1 := int64(34)
	_, err := ConvertStructure(tStrc1)
	if err == nil {
		t.Fatal(fmt.Sprintf("Unexpected nil error for non structure conversion : type : <%s>", reflect.TypeOf(tStrc1).Kind().String()))
	}
}

func TestConvertBusinessObject(t *testing.T) {
	itf := TestStruct1{}
	err := ConvertBusinessObject(__createTestBusinessObject(), &itf)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error from the B.O. conversion, error is <%s>", err.Error()))
	}
	expectedName := "Fabrizio"
	expectedAge := 44
	if expectedName != itf.Name {
		t.Fatal(fmt.Sprintf("BusinessObject has invalid Field Value, Expcted <%s> but Given <%s>", expectedName, itf.Name))
	}
	if expectedAge != itf.Age {
		t.Fatal(fmt.Sprintf("BusinessObject has invalid Field Value, Expcted <%d> but Given <%d>", expectedAge, itf.Age))
	}
}

func TestErrorsInConvertBusinessObject(t *testing.T) {
	itf := TestStruct3{}
	err := ConvertBusinessObject(__createTestBusinessObject(), &itf)
	if err == nil {
		t.Fatal(fmt.Sprintf("Unexpected nil error for non compatible structure conversion : type : <%s>", reflect.TypeOf(itf).Kind().String()))
	}
	itf2 := TestStruct1{}
	err2 := ConvertBusinessObject(__createTestErrorBusinessObject(), &itf2)
	if err2 == nil {
		t.Fatal(fmt.Sprintf("Unexpected nil error for empty fields structure conversion : type : <%s>", reflect.TypeOf(itf2).Kind().String()))
	}
}

func TestConvertCompatibilityBusinessObject(t *testing.T) {
	itf := TestStruct2{}
	err := ConvertBusinessObject(__createTestBusinessObject(), &itf)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error in compatibility derivated structure, error is <%s>", err.Error()))
	}
}
