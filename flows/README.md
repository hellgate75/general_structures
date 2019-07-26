## package flows // import "github.com/hellgate75/general_structures/flows"


### VARIABLES

##### var WORKFLOW_WAIT_TIMEOUT time.Duration = 300 * time.Millisecond

### FUNCTIONS

#### func InitLogger()
    Initialize package logger if not started

#### func NewAction(name string, function func(...interface{}) error, args []interface{}, index int64) Action
    Create a new flows.Action instance
#####     Parameters:
	    name (string) Action name
	    function (func(...interface{}) error) Action function
	    args ([]interface{}) Command arguments
	    index (int64) Execution order
#####     Returns:
	    flows.Action Action instance

#### func NewCommand(path string, cmd string, args []interface{}, index int64) Command
    Create a new flows.Command instance 
#####     Parameters:
	    path (string) Command path (if required)
	    cmd (string) Command text
	    args ([]interface{}) Command arguments
	    index (int64) Execution order
#####     Returns:
	    flows.Command Command instance

#### func NewEmptyContext() Context
    Creates New Empty Context 
#####     Returns:
    	flows.Context Context instance

#### func NewActionExecutable(name string, function func(...interface{}) error, args []interface{}, index int64) Executable
    Create a new flows.Action Action instance
##### 	 Parameters:
	    name (string) Action name
	    function (func(...interface{}) error) Action function
	    args ([]interface{}) Command arguments
	    index (int64) Execution order
#####     Returns:
	    flows.Action Action Executable instance

#### func NewCommandExecutable(path string, cmd string, args []interface{}, index int64) Executable
    Create a new flows.Executable Command instance 
#####     Parameters:
	    path (string) Command path (if required)
	    cmd (string) Command text
	    args ([]interface{}) Command arguments
	    index (int64) Execution order
#####     Returns:
	    flows.Executable Command Executable instance

#### func NewFeature(code string, description string, name string, index int64, executableList []Executable) Feature
    Create New Feature (flows.Feature) Describes features of flow.Feature
    component.
##### 	 Parameters:
	    code (string) It is the code of the workflow
	    description (string) Description of the workflow
	    name (string) Label for the worflow
	    index (int64) Order index, for the execution
	    executableList ([]flows.Executable) executable components list contained in the Feature
#####     Returns:
	    flows.Feature Feature instance

##### func (ft Feature) Run(session *Session, args Arguments) error
    Run the features and report execution status / errors Parameters:

    session (*flows.Session) Execution Session pointer
    args (flows.Arguments) Per Executable Code argumnets array

#### func NewByteArrayFlow(bytes []byte, buffer int) Flow
    Creates and Opens READ/WRITE a bytes.ByteArray based flows.Flow
##### 	 Parameters:
	    bytes (bytes []byte) Initial buffer content
	    buffer (int) read buffer size
#####     Returns:
	    flows.Flow Flow component instance

#### func NewFileFlow(file string, buffer int) (Flow, error)
    Creates and Opens READ/WRITE a file based flows.Flow
##### 	 Parameters:
	    file (string) File name
	    buffer (int) read buffer size
#####     Returns:
	    (flows.Flow Flow component instance,
	    error Any error that occurs during computation)

#### func NewFlow() Flow
    Creates and Opens READ/WRITE a new empty flows.Flow with 64k buffer
##### 	 Returns:
	    flows.Flow Flow component instance

#### func NewFlowBuff(buffer int) Flow
    Creates and Opens READ/WRITE a new empty flows.Flow with specified buffer
#####     Parameters:
	    buffer (int) read buffer size
#####     Returns:
	    flows.Flow Flow component instance

#### func NewPhase(code string, description string, name string, index int64, steps []Step) Phase
    Create New Phase (flows.Phase) Describes features of flow.Phase component.
#####     Parameters:
		    code (string) It is the code of the workflow
		    description (string) Description of the workflow
		    name (string) Label for the worflow
		    index (int64) Order index, for the execution
		    steps ([]flows.Step) step list contained in the Phase
#####     Returns:
		    flows.Phase Phase instance

#### func NewEmptySession(user *UserInfo, context *Context) Session
    Creates New Empty Session
##### 	 Parameters:
	    user (*flows.UserInfo) User information component instance pointer
	    context (*flows.Context) Session Context instance pointer
#####     Returns:
	    flows.Session Session instance

#### func NewStep(code string, description string, name string, index int64, commands []Feature, subSteps []*Step) Step
    Create New Step (flows.Step) Describes features of flow.Step component.
#####     Parameters:
	    code (string) It is the code of the workflow
	    description (string) Description of the workflow
	    name (string) Label for the worflow
	    index (int64) Order index, for the execution
	    commands ([]flows.Feature) execution features list contained in the Step
	    subSteps ([]*flows.Step) sub step pointers list contained in the Step
#####     Returns:
	    flows.Step Step instance

##### func NewWorkflow(code string, description string, name string, phases []Phase) Workflow
    Create New Workflow (flows.Workflow) Describes features of flow.Workflow
    component.
##### 	 Parameters:
	    code (string) It is the code of the workflow
	    description (string) Description of the workflow
	    name (string) Label for the worflow
	    phases ([]flows.Phase) phase list contained in the Workflow
#####     Returns:
	    flows.Workflow Workflow instance

#### func NewWorkflowController(code string, description string, name string, phases []Phase) WorkflowController
    Create New Workflow Controller (flows.WorkflowController) Describes features
    of contained flow.Workflow component, and the execution plan.
#####     Parameters:
	    code (string) It is the code of the workflow
	    description (string) Description of the workflow
	    name (string) Label for the worflow
	    phases ([]flows.Phase) phase list contained in the Workflow
#####     Returns:
	    flows.WorkflowController Workflow controller instance



### TYPES

##### type Action struct {
##### 	Name  string
##### 	Func  func(...interface{}) error
##### 	Args  []interface{}
##### 	Index int64
##### }

##### func (act *Action) Excute(s Session, args ...interface{}) (string, error)
	 Execution method
##### 	 Parameters:
	    s (flows.Session) Is the current execution session
	    args (varidic interface array) Represent runtime execution arguments
##### 	 Returns:
	    ( string execution output,
	      errpr Any error that occurs fduring computation )

##### func (act *Action) String() string
	 Return the String made of byte array in the internal buffer
##### 	 Returns:
	    string  String made of the byte array in the buffer


##### type Arguments map\[string\][]interface{}
    Describes map of Executable code and related arguments It defines the Run
    (execution feature) method and implements flows.Printable

##### type Command struct {
##### 	Path  string
##### 	Cmd   string
##### 	Args  []interface{}
##### 	Index int64
##### }
    Structure that represent the

##### func (cmd *Command) Excute(s Session, args ...interface{}) (string, error)
	 Execution method
##### 	 Parameters:
	    s (flows.Session) Is the current execution session
	    args (varidic interface array) Represent runtime execution arguments
##### 	 Returns:
	    ( string execution output,
	      errpr Any error that occurs fduring computation )

##### func (cmd *Command) String() string
	 Return the String made of byte array in the internal buffer
##### 	 Returns:
	    string  String made of the byte array in the buffer

##### type Context struct {
##### 	// Map that contains workflow wnvironment variables or facts
##### 	Environment map[string]interface{}
##### 	// Context cache for the workflow in order to save teporary data
##### 	Cache map[interface{}]interface{}
##### }
    Context represents the ContextScope for the Executable run

##### func (ctx *Context) AddCacheEntry(name interface{}, value interface{})
    Insert a new entry in the cache
##### 	 Parameters:
	    name (string) New cache entry name
	    value (interface{}) New cache entry value

##### func (ctx *Context) AddEnvEntry(name string, value interface{})
    Insert a new entry in the environment 
#####     Parameters:
	    name (string) New Environmrnt entry name
	    value (interface) New Environmrnt entry value

##### func (ctx *Context) GetCacheEntry(name interface{}) (interface{}, error)
	 Get a cache element 
##### 	 Parameters:
	    name (interface{}) Cache entry name
#####     Returns:
    	(interface{} Chache entry value or nil in case key is not present,
	    error Any arror that occurs during computation)

##### func (ctx *Context) GetCacheKeys() ([]interface{}, error)
    Get an cache keys 
#####     Returns:
	    ([]interface{} Cache Keys,
	    error Any arror that occurs during computation)

##### func (ctx *Context) GetEnvEntry(name string) (interface{}, error)
    Get an environment element
##### 	 Parameters:
	    name (string) Environmrnt entry name
#####     Returns:
	    (interface{} Environment element or nil in case name doesn't exists,
	    error Any arror that occurs during computation)

##### func (ctx *Context) GetEnvKeys() ([]string, error)
    Get an environment keys 
#####     Returns:
	    ([]string Environment Keys,
	    error Any arror that occurs during computation)

##### type Executable interface {
##### 	Printable
##### 	// Execution method
##### 	// Parameters:
##### 	//    s (flows.Session) Is the current execution session
##### 	//    args (varidic interface array) Represent runtime execution arguments
##### 	// Returns:
##### 	//    ( string execution output,
##### 	//      errpr Any error that occurs fduring computation )
##### 	Excute(s Session, args ...interface{}) (string, error)
##### }
    Interface that describes features for any executable structures

##### type Feature struct {
##### 	Code        string       `json:"code" yaml:"code" xml:"code"`
##### 	Name        string       `json:"name" yaml:"name" xml:"name"`
##### 	Description string       `json:"description" yaml:"description" xml:"description"`
##### 	ExecuteList []Executable `json:"list" yaml:"list" xml:"list"`
##### 	Index       int64        `json:"index" yaml:"index" xml:"index"`
##### 	// Has unexported fields.
##### }
    Describes a feature, that is a set of basic commands, eecutable code It has
    an index feature used for be sorted for the runtime run It defines the Run
    (execution feature) method and implements flows.Printable

##### func (ft Feature) String() string
	 Get Component Representing String

##### type Flow interface {
##### 	//Import ReadWriter interface
##### 	io.ReadWriter
##### 	//Import Closer interface
##### 	io.Closer
##### 	//Import Seeker interface
##### 	io.Seeker
##### 	//Import ReadFrom interface
##### 	io.ReaderFrom
##### 	//Import WriterTo interface
##### 	io.WriterTo
##### 	// Open a file in READ/WRITE and create a new file in case
##### 	// Parameters:
##### 	//    name (string) File name
##### 	// Returns:
##### 	//    error Any error that occurs during computation
##### 	Open(name string) error
##### 	// Reset the Selector and the Flow, if a file is open it will be closed and reponened
##### 	// Returns:
##### 	//    error Any error that occurs during computation
##### 	Reset() error
##### 	// Clear the Selector and the Flow, any reference will be closed
##### 	// Returns:
##### 	//    error Any error that occurs during computation
##### 	Clear() error
##### 	// Return the selector name or the empty string in case there is not selector
##### 	// Returns:
##### 	//    string  Selector name
##### 	Selector() string
##### 	// Return the byte array in the internal buffer
##### 	// Returns:
##### 	//    []byte  Byte array in the buffer
##### 	Bytes() []byte
##### 	// Return the String made of byte array in the internal buffer
##### 	// Returns:
##### 	//    string  String made of the byte array in the buffer
##### 	String() string
##### 	// Return Size of the file or internal buffer
##### 	// Returns:
##### 	//    int64 Size of the file or the internal buffer
##### 	Size() int64
##### }
    Flow Element provides Reading/Writing features for Flow Elements

##### type Phase struct {
##### 	Code        string `json:"code" yaml:"code" xml:"code"`
##### 	Name        string `json:"name" yaml:"name" xml:"name"`
##### 	Description string `json:"description" yaml:"description" xml:"description"`
##### 	Steps       []Step `json:"steps" yaml:"steps" xml:"steps"`
##### 	Index       int64  `json:"index" yaml:"index" xml:"index"`
##### 	// Has unexported fields.
##### }
    Describes a phase, that is a set of steps It has an index feature used for
    be sorted for the runtime run It defines the Run (execution feature) method
    and implements flows.Printable

##### func (ph Phase) Run(session *Session, args Arguments) error
    Run the phase steps and report execution status / errors
##### 	 Parameters:
	    session (*flows.Session) Execution Session pointer
	    args (flows.Arguments) Per Executable Code argumnets array

##### func (ph Phase) String() string
	 Get Component Representing String

##### type Printable interface {
##### 	// Print Executable sample type
##### 	// Returns:
##### 	//    string object descriptive string
##### 	String() string
##### }
    Interface that describes features for any printable structures

##### type Role int
    Role type, that will be used for creating role codes in the flows.UserInfo
    component

##### type Serializable interface {
##### 	// Serialize class fields to the flow
##### 	// Parameters:
##### 	//    f (flows.Flow) Flow that is used to write data to the required selector
##### 	// Returns:
##### 	//    error Any error that occurs during the computation
##### 	Serialize(f Flow) error
##### 	// Deserialize class fields from the flow
##### 	// Parameters:
##### 	//    f (flows.Flow) Flow that is used to read data from the required selector
##### 	// Returns:
##### 	//    error Any error that occurs during the computation
##### 	Deserialize(f Flow) error
##### }
    Interface that describes features for any structures that is ready to be
    externalized (read or written)

##### type Session struct {
##### 	// Has unexported fields.
##### }
    Session represents Jobs Execution reference component

##### func (session *Session) Clone() Session
    Get Session Clone
##### 	 Returns:
	    flows.Session Session cloned instance

##### func (session *Session) GerContext() *Context
    Get Current Running Phase
##### 	 Returns:
	    *flows.Context Context pointer

##### func (session *Session) GerCurrentFeature() *Feature
    Get Current Running Feature
##### 	 Returns:
	    *flows.Feature Current running feature pointer

##### func (session *Session) GerCurrentPhase() *Phase
    Get Current Running Phase
##### 	 Returns:
	    *flows.Phase Current running phase pointer

##### func (session *Session) GerCurrentStep() *Step
    Get Current Running Step
##### 	 Returns:
    	*flows.Step Current running step pointer

##### func (session *Session) GerUserInfo() *UserInfo
    Get User Information
##### 	 Returns:
    *flows.UserInfo User descriptor pointer

##### type Step struct {
##### 	Code        string    `json:"code" yaml:"code" xml:"code"`
##### 	Name        string    `json:"name" yaml:"name" xml:"name"`
##### 	Description string    `json:"description" yaml:"description" xml:"description"`
##### 	Commands    []Feature `json:"commands" yaml:"commands" xml:"commands"`
##### 	SubSteps    []*Step   `json:"substeps" yaml:"substeps" xml:"substeps"`
##### 	Index       int64     `json:"index" yaml:"index" xml:"index"`
##### 	// Has unexported fields.
##### }
    Describes a step, that is a set of features It has an index feature used for
    be sorted for the runtime run It defines the Run (execution feature) method
    and implements flows.Printable

##### func (step Step) Run(session *Session, args Arguments) error
    Run the step and sub steps and report execution status / errors
##### 	 Parameters:
	    session (*flows.Session) Execution Session pointer
	    args (flows.Arguments) Per Executable Code argumnets array
#####     Returns:
	    error Any error that occurs during computation

##### func (step Step) String() string
	 Get Component Representing String

##### type UserInfo struct {
##### 	// User identity
##### 	ID string `json:"identity" yaml:"identity" xml:"identity"`
##### 	// Username
##### 	UserName string `json:"userName" yaml:"userName" xml:"username"`
##### 	// User name
##### 	Name string `json:"name" yaml:"name" xml:"name"`
##### 	// User Surname
##### 	Surname string `json:"surname" yaml:"surname" xml:"surname"`
##### 	// User Role
##### 	Role Role `json:"role" yaml:"role" xml:"role"`
##### }
    User information

##### type Workflow struct {
##### 	Code        string  `json:"code" yaml:"code" xml:"code"`
##### 	Name        string  `json:"name" yaml:"name" xml:"name"`
##### 	Description string  `json:"description" yaml:"description" xml:"description"`
##### 	Phases      []Phase `json:"phases" yaml:"phases" xml:"phases"`
##### 	// Has unexported fields.
##### }
    Describes a workflow, that is a set of phases

##### func (wf Workflow) Run(context *Context, info UserInfo, args Arguments) error
    Run the workflow Phases and report execution status / errors
##### 	 Parameters:
	    context (*flows.Context) Execution Context pointer
	    info (flows.UserInfo) User Descriptor
	    args (flows.Arguments) Per Executable Code argumnets array
#####     Returns:
	    error Any error that occurs during computation

##### func (wf Workflow) String() string
	 Get Component Representing String

##### type WorkflowController interface {
##### 	Printable
##### 	// Run the workflow Phases and report execution status / errors
##### 	// Parameters:
##### 	//    context (*flows.Context) Execution Context pointer
##### 	//    info (flows.UserInfo) User Descriptor
##### 	//    args (flows.Arguments) Per Executable Code argumnets array
##### 	Run(context *Context, info UserInfo, args Arguments) error
##### }
    Interface that describes An interface controller It defines the Run
    (execution feature) method and implements flows.Printable

