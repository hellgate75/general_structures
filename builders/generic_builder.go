package builders

import (
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
	"math"
)

//Builder Element that contains Feature Name, Computation Function and map of dependant sub features, available since the feature is invoked, and removed after selection of one of the available, passing to dependant dependancies
type BuilderItem struct {
	Name         string
	Function     BuilderFunction
	Dependancies map[string]BuilderItem
}

//Structure for Generic Builder
type __BuilderStruct struct {
	errorArray      []error
	built           bool
	changeable      bool
	lastItem        BuilderItem
	Builder         interface{}
	OriginalBuilder interface{}
	FeaturesList    map[string]BuilderItem
	BuildFunction   BuilderFunction
}

func (gb *__BuilderStruct) Apply(feature string, args ...Value) Builder {
	defer func() {
		var err error
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error in Builder::Apply(), message : %s", err.Error()))
			}
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error in Builder::Apply(), message : %v", itf))
			}
		}
		if err != nil {
			gb.errorArray = append(gb.errorArray, err)
		}

	}()
	if gb.built && !gb.changeable {
		gb.errorArray = append(gb.errorArray, errors.New(fmt.Sprintf("Unable to apply feature named : '%s' in built Builder!!", feature)))
		return gb
	}
	if function, ok := gb.FeaturesList[feature]; ok {
		gb.lastItem = function
		builder, err := function.Function(gb.Builder, args...)
		if err != nil {
			gb.errorArray = append(gb.errorArray, err)
		} else {
			gb.Builder = builder
		}
	} else {
		if gb.lastItem.Function != nil {
			if function, ok := gb.lastItem.Dependancies[feature]; ok {
				gb.lastItem = function
				builder, err := function.Function(gb.Builder, args...)
				if err != nil {
					gb.errorArray = append(gb.errorArray, err)
				} else {
					gb.Builder = builder
				}
			} else {
				gb.errorArray = append(gb.errorArray, errors.New(fmt.Sprintf("Unable to find feature named : '%s' in Builder Features!!", feature)))

			}
		} else {
			gb.errorArray = append(gb.errorArray, errors.New(fmt.Sprintf("Unable to find feature named : '%s' in Builder Features!!", feature)))
		}
	}
	return gb
}

func __filterErrors(in []error) []error {
	var list []error
	for _, err := range in {
		if err != nil && err.Error() != "<nil>" {
			list = append(list, err)
		}
	}
	return list
}

func (gb *__BuilderStruct) Errors() []error {
	gb.errorArray = __filterErrors(gb.errorArray)
	return gb.errorArray
}

func __rank(start int, end int, in []string) []string {
	if start == end {
		return in
	}
	idx := int(math.Round(float64(end / 2)))
	if idx <= start {
		return in
	}
	left_start := start
	left_end := idx
	right_start := idx
	right_end := end
	if in[left_start] > in[left_end] {
		tmp := in[left_start]
		in[left_start] = in[left_end]
		in[left_end] = tmp
		in = __rank(start, end, in)
	} else {
		in = __rank(left_start+1, left_end, in)
	}
	if in[right_start] > in[right_end] {
		tmp := in[right_start]
		in[right_start] = in[right_end]
		in[right_end] = tmp
		in = __rank(start, end, in)
	} else {
		in = __rank(right_start, right_end-1, in)
	}
	return in
}

func sortStringList(in []string) []string {
	for i := 0; i < len(in); i++ {
		in = __rank(i, len(in)-1, in)
	}
	return in
}

func (gb *__BuilderStruct) Features() []string {
	var list []string
	for _, f := range gb.FeaturesList {
		list = append(list, f.Name)
	}
	if gb.lastItem.Function != nil {
		for _, f := range gb.lastItem.Dependancies {
			list = append(list, f.Name)
		}
	}
	return sortStringList(list)
}

func (gb *__BuilderStruct) Build() (Value, error) {
	var err error
	defer func() {
		itf := recover()
		if errs.IsError(itf) {
			err = itf.(error)
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error in Builder::Build(), message : %s", err.Error()))
			}
		} else {
			err = errors.New(fmt.Sprintf("%v", itf))
			if logger != nil {
				logger.ErrorS(fmt.Sprintf("Error in Builder::Build(), message : %v", itf))
			}
		}
	}()
	if gb.built && !gb.changeable {
		errX := errors.New("Unable to build a built Builder!!")
		gb.errorArray = append(gb.errorArray, errX)
		return nil, errX
	}
	outcome, err := gb.BuildFunction(gb.Builder)
	gb.built = true
	if gb.changeable {
		gb.Builder = gb.OriginalBuilder
		gb.built = false
		gb.errorArray = make([]error, 0)
	}
	return outcome, err
}

// Creates New Generic Buider
// Parameters:
//    builderConfig (builders.Value) Generic Builder Configuration Value
//    features ([]builders.BuilderItem) List of Features available for the Builder
//    builderFunction (builders.BuilderFunction) Function that build the outcome
//    repeatable (bool) Attribute that allows the builder to be re-executed and re-featured (after build builder will reset)
// Returns:
//    builders.Builder Generic Builder instance
func NewBuilder(builderConfig Value, features []BuilderItem, builderFunction BuilderFunction, repeatable bool) Builder {
	var featuresMap map[string]BuilderItem = make(map[string]BuilderItem)
	for _, v := range features {
		featuresMap[v.Name] = v
	}
	return &__BuilderStruct{
		Builder:         builderConfig,
		OriginalBuilder: builderConfig,
		BuildFunction:   builderFunction,
		FeaturesList:    featuresMap,
		changeable:      repeatable,
		errorArray:      make([]error, 0),
	}
}
