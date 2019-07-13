## package builders // import "github.com/hellgate75/general_structures/builders"


### FUNCTIONS

#### func InitLogger()
     Initialize package logger if not started

### TYPES

##### type Builder interface {
##### 	Function that apply arguments to build internal builder elements
##### 	Parameters:
 	   feature (string) Label for the required feature
 	   args (Value variadic array) Arguments if the feature
##### 	Returns:
       builders.Builder the builder instance
##### 	Apply(feature string, args ...Value) Builder
    	Retrieves the errors occured during the build and/or the application of features
##### 	Returns:
      []error Array of occured errors during features application and/or build
##### 	Errors() []error
      List of names related to current available features
##### 	Returns:
      []string List of available features
##### 	Features() []string
      Build the content running the features in the provided sequence, creating the output value
##### 	Returns:
      (builders.Value The built element value,
       error Any error that occurs during build)
##### 	Build() (Value, error)
##### }
    Interface that describes base feature of generic Builder Pattern Applying
    sequentially the builder features it create base configuration and
    definition of Builder configuration. The execution of the build function
    will create the builder target value related to the Builder scope

#### func NewBuilder(builderConfig Value, features []BuilderItem, builderFunction BuilderFunction, repeatable bool) Builder
    Creates New Generic Buider
##### Parameters:

       builderConfig (builders.Value) Generic Builder Configuration Value
       features ([]builders.BuilderItem) List of Features available for the Builder
       builderFunction (builders.BuilderFunction) Function that build the outcome
       repeatable (bool) Attribute that allows the builder to be re-executed and re-featured (after build builder will reset)
#####     Returns:
       builders.Builder Generic Builder instance

#### func NewOperationBuilder(builderConfig Value, operations []OperationBuilderItem, builderFunction BuilderFunction) Builder
    Creates New Operation Builder
#####    Parameters:
       builderConfig (builders.Value) Generic Builder Configuration Value
       operations ([]builders.OperationBuilderItem) List of Operations available for the Builder
       builderFunction (builders.BuilderFunction) Function that build the outcome
#####    Returns:
       builders.Builder Generic Builder instance

#### type BuilderFunction func(builder interface{}, args ...Value) (interface{}, error)
    Function that is used to execute a feature

##### type BuilderItem struct {
##### 	Name         string
##### 	Function     BuilderFunction
##### 	Dependancies map[string]BuilderItem
##### }
    Builder Element that contains Feature Name, Computation Function and map of
    dependant sub features, available since the feature is invoked, and removed
    after selection of one of the available, passing to dependant dependancies

##### type OperationBuilderItem struct {
##### 	Name         string
##### 	Function     BuilderFunction
##### 	SubOperation *OperationBuilderItem
##### 	Active       bool
##### }
    Operation Builder Element, that containes operation name, computation
    function and sub call operation, that is called passing outcome of main, and
    all arguments. It's active when flag is true

##### type Value interface{}
    Default Builder generic value

