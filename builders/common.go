package builders

import (
	"github.com/hellgate75/general_utils/log"
)

var logger log.Logger

func InitLogger() {
	currentLogger, err := log.New("builders")
	if err != nil {
		panic(err.Error())
	}
	logger = currentLogger
}

// Default Builder generic value
type Value interface{}

// Function that is used to execute a feature
type BuilderFunction func(builder interface{}, args ...Value) (interface{}, error)

// Interface that describes base feature of generic Builder Pattern
// Applying sequentially the builder features it create base configuration and definition of
// Builder configuration. The execution of the build function will create the builder target
// value related to the Builder scope
type Builder interface {
	// Function that apply arguments to build internal builder elements
	// Parameters:
	//    feature (string) Label for the required feature
	//    args (Value variadic array) Arguments if the feature
	// Returns:
	//    builders.Builder the builder instance
	Apply(feature string, args ...Value) Builder
	// Retrieves the errors occured during the build and/or the application of features
	// Returns:
	// []arror Array of occured errors during features application and/or build
	Errors() []error
	// List of names related to current available features
	// Returns:
	// []string List of available features
	Features() []string
	// Build the content running the features in the provided sequence, creating the output value
	// Returns:
	// (builders.Value The built element value,
	// error Any error that occurs during build)
	Build() (Value, error)
}
