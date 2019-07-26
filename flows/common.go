package flows

import (
	"errors"
	"fmt"
	"github.com/hellgate75/general_utils/log"
	"os"
)

var fileSep string = fmt.Sprintf("%c", os.PathSeparator)

var logger log.Logger

//Initialize package logger if not started
func InitLogger() {
	currentLogger, err := log.New("flows")
	if err != nil {
		panic(err.Error())
	}
	logger = currentLogger
}

// Context represents the Context Scope for the Executable run
type Context struct {
	// Map that contains workflow wnvironment variables or facts
	Environment map[string]interface{}
	// Context cache for the workflow in order to save teporary data
	Cache map[interface{}]interface{}
}

// Insert a new entry in the environment
// Parameters:
//    name (string) New Environmrnt entry name
//    value (interface) New Environmrnt entry value
func (ctx *Context) AddEnvEntry(name string, value interface{}) {
	ctx.Environment[name] = value
}

// Get an environment element
// Parameters:
//    name (string) Environmrnt entry name
// Returns:
//    (interface{} Environment element or nil in case name doesn't exists,
//    error Any arror that occurs during computation)
func (ctx *Context) GetEnvEntry(name string) (interface{}, error) {
	value, ok := ctx.Environment[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unable to find environment entry for name '%s'", name))
	}
	return value, nil
}

// Get an environment keys
// Returns:
//    ([]string Environment Keys,
//    error Any arror that occurs during computation)
func (ctx *Context) GetEnvKeys() ([]string, error) {
	var keys []string = make([]string, 0)
	for k, _ := range ctx.Environment {
		keys = append(keys, k)
	}
	if len(keys) == 0 {
		return keys, errors.New("Unable to find any environment key!!")
	}
	return keys, nil
}

// Insert a new entry in the cache
// Parameters:
//    name (string) New cache entry name
//    value (interface{}) New cache entry value
func (ctx *Context) AddCacheEntry(name interface{}, value interface{}) {
	ctx.Cache[name] = value
}

// Get a cache element
// Parameters:
//    name (interface{}) Cache entry name
// Returns:
//    error Any arror that occurs during computation
func (ctx *Context) GetCacheEntry(name interface{}) (interface{}, error) {
	value, ok := ctx.Cache[name]
	if !ok {
		return nil, errors.New(fmt.Sprintf("Unable to find environment entry for name '%s'", name))
	}
	return value, nil
}

// Get an cache keys
// Returns:
//    ([]interface{} Cache Keys,
//    error Any arror that occurs during computation)
func (ctx *Context) GetCacheKeys() ([]interface{}, error) {
	var keys []interface{} = make([]interface{}, 0)
	for k, _ := range ctx.Cache {
		keys = append(keys, k)
	}
	if len(keys) == 0 {
		return keys, errors.New("Unable to find any cache key!!")
	}
	return keys, nil
}

// Creates New Empty Context
// Returns:
//    flows.Context Context instance
func NewEmptyContext() Context {
	return Context{
		Cache:       make(map[interface{}]interface{}),
		Environment: make(map[string]interface{}),
	}
}

// Role type, that will be used for creating role codes in the flows.UserInfo component
type Role int

//User information
type UserInfo struct {
	// User identity
	ID string `json:"identity" yaml:"identity" xml:"identity"`
	// Username
	UserName string `json:"userName" yaml:"userName" xml:"username"`
	// User name
	Name string `json:"name" yaml:"name" xml:"name"`
	// User Surname
	Surname string `json:"surname" yaml:"surname" xml:"surname"`
	// User Role
	Role Role `json:"role" yaml:"role" xml:"role"`
}

//Session represents Jobs Execution reference component
type Session struct {
	//User Descriptor
	user *UserInfo `json:"user" yaml:"user" xml:"user"`
	//Current Execution Feature
	currentFeature *Feature `json:"feature" yaml:"feature" xml:"feature"`
	//Current Execution Feature Step
	currentStep *Step `json:"step" yaml:"step" xml:"step"`
	//Current Execution Feature Step Phase
	currentPhase *Phase `json:"phase" yaml:"phase" xml:"phase"`
	// Workflow Context
	context *Context `json:"context" yaml:"context" xml:"context"`
}

// Get User Information
// Returns:
//    *flows.UserInfo User descriptor pointer
func (session *Session) GetUserInfo() *UserInfo {
	return session.user
}

// Get Current Running Feature
// Returns:
//    *flows.Feature Current running feature pointer
func (session *Session) GetCurrentFeature() *Feature {
	return session.currentFeature
}

// Get Current Running Step
// Returns:
//    *flows.Step Current running step pointer
func (session *Session) GetCurrentStep() *Step {
	return session.currentStep
}

// Get Current Running Phase
// Returns:
//    *flows.Phase Current running phase pointer
func (session *Session) GetCurrentPhase() *Phase {
	return session.currentPhase
}

// Get Current Running Phase
// Returns:
//    *flows.Context Context pointer
func (session *Session) GetContext() *Context {
	return session.context
}

// Get Session Clone
// Returns:
//    flows.Session Session cloned instance
func (session *Session) Clone() Session {
	return Session{
		context:        session.context,
		currentFeature: session.currentFeature,
		currentPhase:   session.currentPhase,
		currentStep:    session.currentStep,
		user:           session.user,
	}
}

// Creates New Empty Session
// Parameters:
//    user (*flows.UserInfo) User information component instance pointer
//    context (*flows.Context) Session Context instance pointer
// Returns:
//    flows.Session Session instance
func NewEmptySession(user *UserInfo, context *Context) Session {
	return Session{
		user:    user,
		context: context,
	}
}

// Interface that describes features for any executable structures
type Executable interface {
	Printable
	// Execution method
	// Parameters:
	//    s (flows.Session) Is the current execution session
	//    args (varidic interface array) Represent runtime execution arguments
	// Returns:
	//    ( string execution output,
	//      errpr Any error that occurs fduring computation )
	Excute(s Session, args ...interface{}) (string, error)
}

// Interface that describes features for any printable structures
type Printable interface {
	// Print Executable sample type
	// Returns:
	//    string object descriptive string
	String() string
}

// Interface that describes  features for any structures that is ready to be externalized (read or written)
type Serializable interface {
	// Serialize class fields to the flow
	// Parameters:
	//    f (flows.Flow) Flow that is used to write data to the required selector
	// Returns:
	//    error Any error that occurs during the computation
	Serialize(f Flow) error
	// Deserialize class fields from the flow
	// Parameters:
	//    f (flows.Flow) Flow that is used to read data from the required selector
	// Returns:
	//    error Any error that occurs during the computation
	Deserialize(f Flow) error
}
