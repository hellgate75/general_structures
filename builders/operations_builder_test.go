package builders

import (
	"errors"
	"fmt"
	"testing"
)

type __sampleOperStruct struct {
	valueInt      int64
	valueFloat    float64
	operations    int
	lastOperation string
	mainOperstion string
}

func __createOperationBuilder() Builder {
	builderConfig := __sampleOperStruct{
		0,
		0.0,
		0,
		"",
		"",
	}
	divideByCount := OperationBuilderItem{
		Name: "DivideByCount",
		Function: func(builder interface{}, args ...Value) (interface{}, error) {
			if len(args) < 1 {
				return builder, errors.New("Invalid number of parameters in feature : Divide by Count, expected 1 parameter: bool")
			}
			itf := builder.(__sampleOperStruct)
			var value float64 = float64(itf.valueInt)
			var divisor float64 = float64(len(args))
			itf.valueFloat = value / divisor
			itf.operations++
			itf.lastOperation = "divideByCount"
			return itf, nil
		},
		Active: true,
	}
	sumAllArgs := OperationBuilderItem{
		Name: "Average",
		Function: func(builder interface{}, args ...Value) (interface{}, error) {
			if len(args) < 1 {
				return builder, errors.New("Invalid number of parameters in feature : Sum, expected 1 parameter: bool")
			}
			itf := builder.(__sampleOperStruct)
			var sum int64 = int64(itf.valueInt)
			for _, n := range args {
				sum += n.(int64)
			}
			itf.valueInt = sum
			itf.operations++
			itf.mainOperstion = "average"
			itf.lastOperation = "average"
			return itf, nil
		},
		SubOperation: &divideByCount,
		Active:       true,
	}

	features := make([]OperationBuilderItem, 1)
	features[0] = sumAllArgs
	builderFunction := func(builder interface{}, args ...Value) (interface{}, error) {
		itf := builder.(__sampleOperStruct)
		value := itf.valueFloat
		return value, nil
	}
	return NewOperationBuilder(builderConfig, features, builderFunction)
}

func TestOperationBuilder(t *testing.T) {
	//average builder will call first sum of all arguments as first operation and with the
	// outcome it will call divide the sum by the length of summed numbers
	averageBuilder := __createOperationBuilder()

	list1 := fmt.Sprintf("%v", averageBuilder.Features())
	expectedList1 := "[Average->DivideByCount]"
	if list1 != expectedList1 {
		t.Fatal(fmt.Sprintf("Error retrieving current features : Expected <%s> but Given <%s>", expectedList1, list1))
	}
	averageBuilder.Apply("Average", int64(10), int64(12), int64(14), int64(16), int64(19))
	errorList1 := averageBuilder.Errors()
	if len(errorList1) > 0 {
		t.Fatal(fmt.Sprintf("Not expected errors, instead was given <%v>", errorList1))
	}
	outcome1, _ := averageBuilder.Build()
	expectedOutcome1 := float64(14.2)
	if outcome1.(float64) != expectedOutcome1 {
		t.Fatal(fmt.Sprintf("Error retrieving outcome : Expected <%f> but Given <%f>", expectedOutcome1, outcome1))
	}
	//Resetted
	outcome2, _ := averageBuilder.Apply("Average", int64(10), int64(12), int64(14), int64(16), int64(19)).Build()
	expectedOutcome2 := float64(14.2)
	if outcome2.(float64) != expectedOutcome2 {
		t.Fatal(fmt.Sprintf("Error retrieving outcome : Expected <%f> but Given <%f>", expectedOutcome2, outcome2))
	}
}
