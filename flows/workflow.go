package flows

import (
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
	"reflect"
	"sync"
	"time"
)

var WORKFLOW_WAIT_TIMEOUT time.Duration = 300 * time.Millisecond

// Describes a feature, that is a set of basic commands, eecutable code
// It has an index feature used for be sorted for the runtime run
// It defines the Run (execution feature) method and implements flows.Printable
type Feature struct {
	Code        string       `json:"code" yaml:"code" xml:"code"`
	Name        string       `json:"name" yaml:"name" xml:"name"`
	Description string       `json:"description" yaml:"description" xml:"description"`
	ExecuteList []Executable `json:"list" yaml:"list" xml:"list"`
	Index       int64        `json:"index" yaml:"index" xml:"index"`
	__mutex     sync.RWMutex
}

func (ft Feature) String() string {
	var list string
	var lenght int = len(ft.ExecuteList)
	for i, v := range ft.ExecuteList {
		list += v.String()
		if i < lenght-1 {
			list += ", "
		}
	}
	return fmt.Sprintf("Feature{Code: '%s',Name: '%s', Description: '%s', Index: %d, Running: [%s]}", ft.Code, ft.Name, ft.Description, ft.Index, list)
}

// Run the features and report execution status / errors
// Parameters:
//    session (*flows.Session) Execution Session pointer
//    args (flows.Arguments) Per Executable Code argumnets array
// Returns:
//    error Any arror that occurs during computation
func (ft Feature) Run(session *Session, args Arguments) error {
	var err error = nil
	defer func() {
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				err = itf.(error)
			} else {
				err = errors.New(fmt.Sprintf("%v", itf))
			}
		}
	}()
	eSession := session.Clone()
	eSession.currentFeature = &ft
	var instancesCounter int = 0
	increaseFeature := func() {
		ft.__mutex.Lock()
		instancesCounter++
		ft.__mutex.Unlock()
	}
	decreaseFeature := func() {
		ft.__mutex.Lock()
		instancesCounter--
		ft.__mutex.Unlock()
	}
	var errList []error = make([]error, 0)
	for _, execute := range ft.ExecuteList {
		increaseFeature()
		go func(ex Executable) {
			defer func() {
				itf := recover()
				if itf != nil {
					if errs.IsError(itf) {
						err = itf.(error)
					} else {
						err = errors.New(fmt.Sprintf("%v", itf))
					}
				}
				decreaseFeature()
			}()
			var code string = "UNDEFINED"
			if reflect.TypeOf(execute).String() == "*flows.Command" {
				cmd := execute.(*Command)
				sepF := ""
				if "" != cmd.Path {
					sepF = fmt.Sprintf("%s%s", cmd.Path, fileSep)
				}
				code = fmt.Sprintf("%s%s-%d", sepF, cmd.Cmd, cmd.Index)
			} else if reflect.TypeOf(execute).String() == "*flows.Action" {
				act := execute.(*Action)
				code = fmt.Sprintf("%s-%d", act.Name, act.Index)
			}
			var arguments []interface{} = make([]interface{}, 0)
			if argsX, ok := args[code]; ok {
				arguments = append(arguments, argsX...)
			}
			out, errX := ex.Excute(eSession, arguments...)
			errList = append(errList, errX)
			if len(out) > 0 {
				if logger != nil {
					logger.Debug(fmt.Sprintf("Executable <%s> output : %s", ex.String(), out))
				}
			}
		}(execute)
	}
	for instancesCounter > 0 {
		time.Sleep(WORKFLOW_WAIT_TIMEOUT)
	}
	if len(errList) > 0 {
		err = errors.New(fmt.Sprintf("Errors occured during execution: %s", encodeErrors(errList...)))
	}
	return err
}

// Describes a step, that is a set of features
// It has an index feature used for be sorted for the runtime run
// It defines the Run (execution feature) method and implements flows.Printable
type Step struct {
	Code        string    `json:"code" yaml:"code" xml:"code"`
	Name        string    `json:"name" yaml:"name" xml:"name"`
	Description string    `json:"description" yaml:"description" xml:"description"`
	Commands    []Feature `json:"commands" yaml:"commands" xml:"commands"`
	SubSteps    []*Step   `json:"substeps" yaml:"substeps" xml:"substeps"`
	Index       int64     `json:"index" yaml:"index" xml:"index"`
	__mutex     sync.RWMutex
}

func (step Step) String() string {
	var list1 string
	var lenght1 int = len(step.Commands)
	for i, v := range step.Commands {
		list1 += v.String()
		if i < lenght1-1 {
			list1 += ", "
		}
	}
	var list2 string
	var lenght2 int = len(step.SubSteps)
	for i, v := range step.SubSteps {
		list2 += v.String()
		if i < lenght2-1 {
			list2 += ", "
		}
	}
	return fmt.Sprintf("Step{Code: '%s',Name: '%s', Description: '%s', Index: %d, Commands: [%s], SubSteps: [%s]}", step.Code, step.Name, step.Description, step.Index, list1, list2)
}

// Run the step and sub steps and report execution status / errors
// Parameters:
//    session (*flows.Session) Execution Session pointer
//    args (flows.Arguments) Per Executable Code argumnets array
// Returns:
//    error Any arror that occurs during computation
func (step Step) Run(session *Session, args Arguments) error {
	var err error = nil
	defer func() {
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				err = itf.(error)
			} else {
				err = errors.New(fmt.Sprintf("%v", itf))
			}
		}
	}()
	eSession := session.Clone()
	eSession.currentStep = &step
	var errList []error = make([]error, 0)
	var instancesCounter int = 0
	increaseStep := func() {
		step.__mutex.Lock()
		instancesCounter++
		step.__mutex.Unlock()
	}
	decreaseStep := func() {
		step.__mutex.Lock()
		instancesCounter--
		step.__mutex.Unlock()
	}
	for _, cmd := range step.Commands {
		increaseStep()
		go func(c Feature) {
			defer func() {
				recover()
				decreaseStep()
			}()
			errX := c.Run(&eSession, args)
			errList = append(errList, errX)
		}(cmd)
	}
	for _, step := range step.SubSteps {
		increaseStep()
		go func(st Step) {
			defer func() {
				recover()
				decreaseStep()
			}()
			errX := st.Run(&eSession, args)
			errList = append(errList, errX)
		}(*step)
	}
	for instancesCounter > 0 {
		time.Sleep(WORKFLOW_WAIT_TIMEOUT)
	}
	if len(errList) > 0 {
		err = errors.New(fmt.Sprintf("Errors occured during execution: %s", encodeErrors(errList...)))
	}
	return err
}

// Describes a phase, that is a set of steps
// It has an index feature used for be sorted for the runtime run
// It defines the Run (execution feature) method and implements flows.Printable
type Phase struct {
	Code        string `json:"code" yaml:"code" xml:"code"`
	Name        string `json:"name" yaml:"name" xml:"name"`
	Description string `json:"description" yaml:"description" xml:"description"`
	Steps       []Step `json:"steps" yaml:"steps" xml:"steps"`
	Index       int64  `json:"index" yaml:"index" xml:"index"`
	__mutex     sync.RWMutex
}

func (ph Phase) String() string {
	var list string
	var lenght int = len(ph.Steps)
	for i, v := range ph.Steps {
		list += v.String()
		if i < lenght-1 {
			list += ", "
		}
	}
	return fmt.Sprintf("Phase{Code: '%s', Name: '%s', Description: '%s', Index: %d, Steps: [%s]}", ph.Code, ph.Name, ph.Description, ph.Index, list)
}

// Run the phase steps and report execution status / errors
// Parameters:
//    session (*flows.Session) Execution Session pointer
//    args (flows.Arguments) Per Executable Code argumnets array
// Returns:
//    error Any arror that occurs during computation
func (ph Phase) Run(session *Session, args Arguments) error {
	var err error = nil
	defer func() {
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				err = itf.(error)
			} else {
				err = errors.New(fmt.Sprintf("%v", itf))
			}
		}
	}()
	eSession := session.Clone()
	eSession.currentPhase = &ph
	var errList []error = make([]error, 0)
	var instancesCounter int = 0
	increasePhs := func() {
		ph.__mutex.Lock()
		instancesCounter++
		ph.__mutex.Unlock()
	}
	decreasePhs := func() {
		ph.__mutex.Lock()
		instancesCounter--
		ph.__mutex.Unlock()
	}
	for _, step := range ph.Steps {
		increasePhs()
		go func(st Step) {
			defer func() {
				recover()
				decreasePhs()
			}()
			errX := st.Run(&eSession, args)
			errList = append(errList, errX)
		}(step)
	}
	for instancesCounter > 0 {
		time.Sleep(WORKFLOW_WAIT_TIMEOUT)
	}
	if len(errList) > 0 {
		err = errors.New(fmt.Sprintf("Errors occured during execution: %s", encodeErrors(errList...)))
	}
	return err
}

//Interface that describes An interface controller
// It defines the Run (execution feature) method and implements flows.Printable
type WorkflowController interface {
	Printable
	// Run the workflow Phases and report execution status / errors
	// Parameters:
	//    context (*flows.Context) Execution Context pointer
	//    info (flows.UserInfo) User Descriptor
	//    args (flows.Arguments) Per Executable Code argumnets array
	Run(context *Context, info UserInfo, args Arguments) error
}

// Describes a workflow, that is a set of phases
type Workflow struct {
	Code        string  `json:"code" yaml:"code" xml:"code"`
	Name        string  `json:"name" yaml:"name" xml:"name"`
	Description string  `json:"description" yaml:"description" xml:"description"`
	Phases      []Phase `json:"phases" yaml:"phases" xml:"phases"`
	__mutex     sync.RWMutex
}

func (wf Workflow) String() string {
	var list string
	var lenght int = len(wf.Phases)
	for i, v := range wf.Phases {
		list += v.String()
		if i < lenght-1 {
			list += ", "
		}
	}
	return fmt.Sprintf("Workflow{Code: '%s',Name: '%s', Description: '%s', Phases: [%s]}", wf.Code, wf.Name, wf.Description, list)
}

// Run the workflow Phases and report execution status / errors
// Parameters:
//    context (*flows.Context) Execution Context pointer
//    info (flows.UserInfo) User Descriptor
//    args (flows.Arguments) Per Executable Code argumnets array
func (wf Workflow) Run(context *Context, info UserInfo, args Arguments) error {
	var err error = nil
	defer func() {
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				err = itf.(error)
			} else {
				err = errors.New(fmt.Sprintf("%v", itf))
			}
		}
	}()
	session := Session{
		user:    &info,
		context: context,
	}
	var errList []error = make([]error, 0)
	var instancesCounter int = 0
	increaseWkf := func() {
		wf.__mutex.Lock()
		instancesCounter++
		wf.__mutex.Unlock()
	}
	decreaseWkf := func() {
		wf.__mutex.Lock()
		instancesCounter--
		wf.__mutex.Unlock()
	}
	for _, phase := range wf.Phases {
		increaseWkf()
		go func(ph Phase) {
			defer func() {
				recover()
				decreaseWkf()
			}()
			errX := ph.Run(&session, args)
			errList = append(errList, errX)
		}(phase)
	}
	for instancesCounter > 0 {
		time.Sleep(WORKFLOW_WAIT_TIMEOUT)
	}
	if len(errList) > 0 {
		err = errors.New(fmt.Sprintf("Errors occured during execution: %s", encodeErrors(errList...)))
	}
	return err
}

//Describes map of Executable code and related arguments
// It defines the Run (execution feature) method and implements flows.Printable
type Arguments map[string][]interface{}

func encodeErrors(errArr ...error) string {
	var out string = ""
	var length int = len(errArr)
	for i, v := range errArr {
		out += v.Error()
		if i < length-1 {
			out += ", "
		}
	}
	return out
}

// Create New Workflow Controller (flows.WorkflowController)
// Describes features of contained flow.Workflow component, and the execution plan.
// Parameters:
//    code (string) It is the code of the workflow
//    description (string) Description of the workflow
//    name (string) Label for the worflow
//    phases ([]flows.Phase) phase list contained in the Workflow
// Returns:
//    flows.WorkflowController Workflow controller instance
func NewWorkflowController(code string, description string, name string, phases []Phase) WorkflowController {
	return Workflow{
		Code:        code,
		Description: description,
		Name:        name,
		Phases:      phases,
		__mutex:     sync.RWMutex{},
	}
}

// Create New Workflow (flows.Workflow)
// Describes features of flow.Workflow component.
// Parameters:
//    code (string) It is the code of the workflow
//    description (string) Description of the workflow
//    name (string) Label for the worflow
//    phases ([]flows.Phase) phase list contained in the Workflow
// Returns:
//    flows.Workflow Workflow instance
func NewWorkflow(code string, description string, name string, phases []Phase) Workflow {
	return Workflow{
		Code:        code,
		Description: description,
		Name:        name,
		Phases:      phases,
		__mutex:     sync.RWMutex{},
	}
}

// Create New Phase (flows.Phase)
// Describes features of flow.Phase component.
// Parameters:
//    code (string) It is the code of the workflow
//    description (string) Description of the workflow
//    name (string) Label for the worflow
//    index (int64) Order index, for the execution
//    steps ([]flows.Step) step list contained in the Phase
// Returns:
//    flows.Phase Phase instance
func NewPhase(code string, description string, name string, index int64, steps []Step) Phase {
	return Phase{
		Code:        code,
		Description: description,
		Name:        name,
		Index:       index,
		Steps:       steps,
		__mutex:     sync.RWMutex{},
	}
}

// Create New Step (flows.Step)
// Describes features of flow.Step component.
// Parameters:
//    code (string) It is the code of the workflow
//    description (string) Description of the workflow
//    name (string) Label for the worflow
//    index (int64) Order index, for the execution
//    commands ([]flows.Feature) execution features list contained in the Step
//    subSteps ([]*flows.Step) sub step pointers list contained in the Step
// Returns:
//    flows.Step Step instance
func NewStep(code string, description string, name string, index int64, commands []Feature, subSteps []*Step) Step {
	return Step{
		Code:        code,
		Description: description,
		Name:        name,
		Index:       index,
		Commands:    commands,
		SubSteps:    subSteps,
		__mutex:     sync.RWMutex{},
	}
}

// Create New Feature (flows.Feature)
// Describes features of flow.Feature component.
// Parameters:
//    code (string) It is the code of the workflow
//    description (string) Description of the workflow
//    name (string) Label for the worflow
//    index (int64) Order index, for the execution
//    executableList ([]flows.Executable) executable components list contained in the Feature
// Returns:
//    flows.Feature Feature instance
func NewFeature(code string, description string, name string, index int64, executableList []Executable) Feature {
	return Feature{
		Code:        code,
		Description: description,
		Name:        name,
		Index:       index,
		ExecuteList: executableList,
		__mutex:     sync.RWMutex{},
	}
}
