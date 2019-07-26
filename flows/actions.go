package flows

import (
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
	"os/exec"
)

// Structure that represent the
type Command struct {
	Path  string
	Cmd   string
	Args  []interface{}
	Index int64
}

func (cmd *Command) String() string {
	var list string
	var lenght int = len(cmd.Args)
	for i, v := range cmd.Args {
		list += fmt.Sprintf("%s", v)
		if i < lenght-1 {
			list += ", "
		}
	}
	return fmt.Sprintf("Command{Path: '%s', Cmd: '%s', Index: %d, Args: [%s]}", cmd.Path, cmd.Cmd, cmd.Index, list)
}

func (cmd *Command) Excute(s Session, args ...interface{}) (string, error) {
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
	var argItfs []interface{}
	argItfs = append(argItfs, cmd.Args...)
	argItfs = append(argItfs, args...)

	var argStrs []string
	for _, v := range argItfs {
		argStrs = append(argStrs, fmt.Sprintf("%v", v))
	}

	var cmdStr string = cmd.Cmd
	if cmd.Path != "" {
		cmdStr = fmt.Sprintf("%s%s%s", cmd.Path, fileSep, cmd.Cmd)
	}
	var cmdExec *exec.Cmd = exec.Command(cmdStr, argStrs...)
	if err != nil {
		var message string = fmt.Sprintf("Error executing command, error is '%s'", err.Error())
		if logger != nil {
			logger.ErrorS(message)
		}
		bytes, errE := cmdExec.CombinedOutput()
		if errE != nil {
			bytes = []byte("<no output>")
		}
		return fmt.Sprintf("%s", bytes), errors.New(message)

	}
	bytes, errE := cmdExec.CombinedOutput()
	if errE != nil {
		if logger != nil {
			logger.ErrorS(fmt.Sprintf("Error retrieving command output, error is '%s'", errE.Error()))
		}
		return "<no output>", errE

	} else {
		return fmt.Sprintf("%s", bytes), err
	}
	return "", err
}

// Create a new flows.Executable Command instance
// Parameters:
//    path (string) Command path (if required)
//    cmd (string) Command text
//    args ([]interface{}) Command arguments
//    index (int64) Execution order
// Returns:
//    flows.Executable Command Executable instance
func NewCommandExecutable(path string, cmd string, args []interface{}, index int64) Executable {
	return &Command{
		Path:  path,
		Cmd:   cmd,
		Args:  args,
		Index: index,
	}
}

// Create a new flows.Command instance
// Parameters:
//    path (string) Command path (if required)
//    cmd (string) Command text
//    args ([]interface{}) Command arguments
//    index (int64) Execution order
// Returns:
//    flows.Command Command instance
func NewCommand(path string, cmd string, args []interface{}, index int64) Command {
	return Command{
		Path:  path,
		Cmd:   cmd,
		Args:  args,
		Index: index,
	}
}

type Action struct {
	Name  string
	Func  func(...interface{}) error
	Args  []interface{}
	Index int64
}

func (act *Action) String() string {
	var list string
	var lenght int = len(act.Args)
	for i, v := range act.Args {
		list += fmt.Sprintf("%s", v)
		if i < lenght-1 {
			list += ", "
		}
	}
	return fmt.Sprintf("Action{Name: '%s', Func: <function>, Index: %d, Args: [%s]}", act.Name, act.Index, list)
}

func (act *Action) Excute(s Session, args ...interface{}) (string, error) {
	var output string
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
	var argItfs []interface{}
	argItfs = append(argItfs, act.Args...)
	argItfs = append(argItfs, args...)
	//	for _, v := range act.Args {
	//		argItfs = append(argItfs, v)
	//	}
	//	for _, v := range args {
	//		argItfs = append(argItfs, v)
	//	}
	errF := act.Func(argItfs...)
	if errF != nil {
		output = fmt.Sprintf("Function %s, result: error, message: %s", act.Name, errF.Error())
		err = errF
	} else {
		output = fmt.Sprintf("Function %s, result: success", act.Name)
	}
	return output, err
}

// Create a new flows.Action Action instance
// Parameters:
//    name (string) Action name
//    function (func(...interface{}) error) Action function
//    args ([]interface{}) Command arguments
//    index (int64) Execution order
// Returns:
//    flows.Action Action Executable instance
func NewActionExecutable(name string, function func(...interface{}) error, args []interface{}, index int64) Executable {
	return &Action{
		Name:  name,
		Func:  function,
		Args:  args,
		Index: index,
	}
}

// Create a new flows.Action instance
// Parameters:
//    name (string) Action name
//    function (func(...interface{}) error) Action function
//    args ([]interface{}) Command arguments
//    index (int64) Execution order
// Returns:
//    flows.Action Action instance
func NewAction(name string, function func(...interface{}) error, args []interface{}, index int64) Action {
	return Action{
		Name:  name,
		Func:  function,
		Args:  args,
		Index: index,
	}
}
