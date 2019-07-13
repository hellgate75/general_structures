package builders

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

type __sampleStruct struct {
	upperCase bool
	lowerCase bool
	trim      bool
	value     string
}

func __createBuilder(repeat bool) Builder {
	builderConfig := __sampleStruct{
		false,
		false,
		false,
		"",
	}
	subFeatures := make(map[string]BuilderItem)
	subFeatures["Trim"] = BuilderItem{
		Name: "Trim",
		Function: func(builder interface{}, args ...Value) (interface{}, error) {
			if len(args) < 1 {
				return builder, errors.New("Invalid number of parameters in feature : Trim, expected 1 parameter: bool")
			}
			itf := builder.(__sampleStruct)
			itf.trim = args[0].(bool)
			return itf, nil
		},
	}

	features := make([]BuilderItem, 3)
	features[0] = BuilderItem{
		Name: "With",
		Function: func(builder interface{}, args ...Value) (interface{}, error) {
			if len(args) < 1 {
				return builder, errors.New("Invalid number of parameters in feature : With, expected 1 parameter: string")
			}
			itf := builder.(__sampleStruct)
			itf.value = args[0].(string)
			return itf, nil
		},
	}
	features[1] = BuilderItem{
		Name: "Lower",
		Function: func(builder interface{}, args ...Value) (interface{}, error) {
			if len(args) < 1 {
				return builder, errors.New("Invalid number of parameters in feature : Lower, expected 1 parameter: bool")
			}
			itf := builder.(__sampleStruct)
			itf.lowerCase = args[0].(bool)
			return itf, nil
		},
		Dependancies: subFeatures,
	}
	features[2] = BuilderItem{
		Name: "Upper",
		Function: func(builder interface{}, args ...Value) (interface{}, error) {
			if len(args) < 1 {
				return builder, errors.New("Invalid number of parameters in feature : Upper, expected 1 parameter: bool")
			}
			itf := builder.(__sampleStruct)
			itf.upperCase = args[0].(bool)
			return itf, nil
		},
		Dependancies: subFeatures,
	}
	builderFunction := func(builder interface{}, args ...Value) (interface{}, error) {
		itf := builder.(__sampleStruct)
		value := itf.value
		if itf.lowerCase {
			value = strings.ToLower(value)
		}
		if itf.upperCase {
			value = strings.ToUpper(value)
		}
		if itf.trim {
			value = strings.TrimSpace(value)
		}
		return value, nil
	}
	return NewBuilder(builderConfig, features, builderFunction, repeat)
}

func TestBuilder(t *testing.T) {
	builder := __createBuilder(true)
	list1 := fmt.Sprintf("%v", builder.Features())
	expectedList1 := "[Lower Upper With]"
	if list1 != expectedList1 {
		t.Fatal(fmt.Sprintf("Error retrieving current features : Expected <%s> but Given <%s>", expectedList1, list1))
	}
	builder.Apply("With", " String With Spaces ")
	list2 := fmt.Sprintf("%v", builder.Features())
	if list2 != expectedList1 {
		t.Fatal(fmt.Sprintf("Error retrieving current features : Expected <%s> but Given <%s>", expectedList1, list2))
	}
	builder.Apply("Lower", true)
	list3 := fmt.Sprintf("%v", builder.Features())
	expectedList3 := "[Lower Trim Upper With]"
	if list3 != expectedList3 {
		t.Fatal(fmt.Sprintf("Error retrieving current features : Expected <%s> but Given <%s>", expectedList3, list3))
	}
	errorList1 := builder.Errors()
	if len(errorList1) > 0 {
		t.Fatal(fmt.Sprintf("Not expected errors, instead was given <%v>", errorList1))
	}
	outcome1, _ := builder.Build()
	expectedOutcome1 := " string with spaces "
	if outcome1.(string) != expectedOutcome1 {
		t.Fatal(fmt.Sprintf("Error retrieving outcome : Expected <%s> but Given <%v>", expectedOutcome1, outcome1))
	}
	//Resetted
	builder.Apply("With", " String With Spaces ").Apply("Upper", true).Apply("Trim", true)
	errorList2 := builder.Errors()
	if len(errorList2) > 0 {
		t.Fatal(fmt.Sprintf("Not expected errors, instead was given <%v>", errorList2))
	}
	outcome2, _ := builder.Build()
	expectedOutcome2 := "STRING WITH SPACES"
	if outcome2.(string) != expectedOutcome2 {
		t.Fatal(fmt.Sprintf("Error retrieving outcome : Expected <%s> but Given <%v>", expectedOutcome2, outcome2))
	}
}
func TestBuilderErrors1(t *testing.T) {
	builder := __createBuilder(false)
	builder.Apply("With", "this").Apply("Cut", "that").Build()
	errorList1 := builder.Errors()
	if len(errorList1) != 1 {
		t.Fatal(fmt.Sprintf("Not expected errors, instead was given <%v>", errorList1))
	}
	expectedErrorList1 := "Unable to find feature named : 'Cut' in Builder Features!!"
	givenErrorList1 := errorList1[0].Error()
	if givenErrorList1 != expectedErrorList1 {
		t.Fatal(fmt.Sprintf("Wrong Error message : Expected <%s> but Given <%v>", givenErrorList1, expectedErrorList1))
	}

}
func TestBuilderErrors2(t *testing.T) {
	builder := __createBuilder(false)
	builder.Apply("With", "this").Build()
	builder.Apply("Lower", true).Build()
	errorList1 := builder.Errors()
	if len(errorList1) != 2 {
		t.Fatal(fmt.Sprintf("Not expected errors, instead was given <%v>", errorList1))
	}
	expectedErrorList1 := "Unable to apply feature named : 'Lower' in built Builder!!"
	givenErrorList1 := errorList1[0].Error()
	if givenErrorList1 != expectedErrorList1 {
		t.Fatal(fmt.Sprintf("Wrong Error message : Expected <%s> but Given <%v>", givenErrorList1, expectedErrorList1))
	}
	expectedErrorList2 := "Unable to build a built Builder!!"
	givenErrorList2 := errorList1[1].Error()
	if givenErrorList2 != expectedErrorList2 {
		t.Fatal(fmt.Sprintf("Wrong Error message : Expected <%s> but Given <%v>", givenErrorList2, expectedErrorList2))
	}

}
