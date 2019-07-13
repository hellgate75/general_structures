package builders

import (
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
)

//Operation Builder Element, that containes operation name, computation function and sub call operation, that is called passing outcome of main, and all arguments. It's active when flag is true
type OperationBuilderItem struct {
	Name         string
	Function     BuilderFunction
	SubOperation *OperationBuilderItem
	Active       bool
}

//Structure for Operation Builder
type __OperationBuilderStruct struct {
	errorArray      []error
	Builder         interface{}
	OriginalBuilder interface{}
	OperationsList  map[string]OperationBuilderItem
	BuildFunction   BuilderFunction
}

func (gb *__OperationBuilderStruct) Apply(feature string, args ...Value) Builder {
	defer func() {
		var err error
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error in OperationBuilder::Apply(), message : %s", err.Error()))
			}
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error in OperationBuilder::Apply(), message : %v", itf))
			}
		}
		if err != nil {
			gb.errorArray = append(gb.errorArray, err)
		}

	}()
	if function, ok := gb.OperationsList[feature]; ok {
		//Execute sub-sequentially operations from main to siblings
		for function.Active {
			builder, err := function.Function(gb.Builder, args...)
			if err != nil {
				gb.errorArray = append(gb.errorArray, err)
			} else {
				gb.Builder = builder
			}
			if function.SubOperation != nil {
				function = *function.SubOperation
			} else {
				break
			}
		}
	} else {
		gb.errorArray = append(gb.errorArray, errors.New(fmt.Sprintf("Unable to find operation named : '%s' in Builder Operations!!", feature)))
	}
	return gb
}

func (gb *__OperationBuilderStruct) Errors() []error {
	gb.errorArray = __filterErrors(gb.errorArray)
	return gb.errorArray
}

func __buildOperationString(item OperationBuilderItem) string {
	itemVal := item
	out := ""
	for itemVal.Active {
		out += itemVal.Name
		if itemVal.SubOperation != nil {
			itemVal = *itemVal.SubOperation
			out += "->"
		} else {
			break
		}
	}
	return out
}

func (gb *__OperationBuilderStruct) Features() []string {
	var list []string
	for _, f := range gb.OperationsList {
		if f.Active {
			list = append(list, __buildOperationString(f))
		}
	}
	return sortStringList(list)
}

func (gb *__OperationBuilderStruct) Build() (Value, error) {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error in OperationBuilder::Build(), message : %s", err.Error()))
			}
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error in OperationBuilder::Build(), message : %v", itf))
			}
		}
	}()
	outcome, err := gb.BuildFunction(gb.Builder)
	gb.Builder = gb.OriginalBuilder
	gb.errorArray = make([]error, 0)
	return outcome, err
}

// Creates New Operation Buider
// Parameters:
//    builderConfig (builders.Value) Generic Builder Configuration Value
//    operations ([]builders.OperationBuilderItem) List of Operations available for the Builder
//    builderFunction (builders.BuilderFunction) Function that build the outcome
// Returns:
//    builders.Builder Generic Builder instance
func NewOperationBuilder(builderConfig Value, operations []OperationBuilderItem, builderFunction BuilderFunction) Builder {
	var featuresMap map[string]OperationBuilderItem = make(map[string]OperationBuilderItem)
	for _, v := range operations {
		featuresMap[v.Name] = v
	}
	return &__OperationBuilderStruct{
		Builder:         builderConfig,
		OriginalBuilder: builderConfig,
		BuildFunction:   builderFunction,
		OperationsList:  featuresMap,
		errorArray:      make([]error, 0),
	}
}
