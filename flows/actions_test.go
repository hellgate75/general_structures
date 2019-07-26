package flows

import (
	"fmt"
	"runtime"
	"testing"
)

func __createTestUserInfo() UserInfo {
	return UserInfo{
		ID:       "0001",
		Name:     "Fabrizio",
		Surname:  "Torelli",
		UserName: "TORELFABR",
		Role:     Role(5),
	}
}

func __createTestContext() Context {
	return NewEmptyContext()
}

func __createTestSession(userInfo *UserInfo, ctx *Context) Session {
	return Session{
		user:    userInfo,
		context: ctx,
	}
}

func TestNewCommandExecutable(t *testing.T) {
	userInfo := __createTestUserInfo()
	ctx := __createTestContext()
	session := __createTestSession(&userInfo, &ctx)
	command := "ls"
	args := make([]interface{}, 0)
	args = append(args, "-l")
	if runtime.GOOS == "windows" {
		command = "cmd"
		args = make([]interface{}, 0)
		args = append(args, "\"dir /S\"")
	} else {
		return
	}
	var exec Executable = NewCommandExecutable("", command, args, int64(0))
	output, err := exec.Excute(session, []string{})
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error, message : %s", err.Error()))
	}
	if output == "" {
		t.Fatal("Unexpected empty output message!!")
	}
}

func TestNewCommand(t *testing.T) {
	userInfo := __createTestUserInfo()
	ctx := __createTestContext()
	session := __createTestSession(&userInfo, &ctx)
	command := "ls"
	args := make([]interface{}, 0)
	args = append(args, "-l")
	if runtime.GOOS == "windows" {
		command = "cmd"
		args = make([]interface{}, 0)
		args = append(args, "\"dir /S\"")
	} else {
		return
	}
	var exec Command = NewCommand("", command, args, int64(0))
	output, err := exec.Excute(session, []string{})
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error, message : %s", err.Error()))
	}
	if output == "" {
		t.Fatal("Unexpected empty output message!!")
	}
	expected := "Command{Path: '', Cmd: 'cmd', Index: 0, Args: [\"dir /S\"]}"
	current := exec.String()
	if expected != current {
		t.Fatal(fmt.Sprintf("Error in Command String(), Expected: <%s> but Given: <%s>", expected, current))
	}
}

func TestActionExecutable(t *testing.T) {
	userInfo := __createTestUserInfo()
	ctx := __createTestContext()
	session := __createTestSession(&userInfo, &ctx)
	var exec Executable = NewActionExecutable("Func1", func(...interface{}) error { return nil }, make([]interface{}, 0), int64(4))
	output, err := exec.Excute(session, []string{})
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error, message : %s", err.Error()))
	}
	if output == "" {
		t.Fatal("Unexpected empty output message!!")
	}
}

func TestAction(t *testing.T) {
	userInfo := __createTestUserInfo()
	ctx := __createTestContext()
	session := __createTestSession(&userInfo, &ctx)
	var exec Action = NewAction("Func1", func(...interface{}) error { return nil }, make([]interface{}, 0), int64(4))
	output, err := exec.Excute(session, []string{})
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error, message : %s", err.Error()))
	}
	if output == "" {
		t.Fatal("Unexpected empty output message!!")
	}
	expected := "Action{Name: 'Func1', Func: <function>, Index: 4, Args: []}"
	current := exec.String()
	if expected != current {
		t.Fatal(fmt.Sprintf("Error in Action String(), Expected: <%s> but Given: <%s>", expected, current))
	}
}
