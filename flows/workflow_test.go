package flows

import (
	"fmt"
	"testing"
)

var __testOut1, __testOut2 string

func __createTestArguments() Arguments {
	args := Arguments(make(map[string][]interface{}))
	args["FUNC1-1"] = make([]interface{}, 0)
	args["FUNC1-1"] = append(args["FUNC1-1"], "Sample")
	args["FUNC1-1"] = append(args["FUNC1-1"], "text")
	args["FUNC1-1"] = append(args["FUNC1-1"], "here")
	args["FUNC2-2"] = make([]interface{}, 0)
	args["FUNC2-2"] = append(args["FUNC2-2"], "Another")
	args["FUNC2-2"] = append(args["FUNC2-2"], "sample")
	args["FUNC2-2"] = append(args["FUNC2-2"], "text")
	args["FUNC2-2"] = append(args["FUNC2-2"], "here")
	return args
}

func __createTestWorkflowFunction1() Executable {
	return NewActionExecutable("FUNC1", func(args ...interface{}) error {
		var out string = ""
		out = "FUNC1 -> "
		for _, v := range args {
			out = fmt.Sprintf("%s%v%s", out, v, " ")
		}
		if len(out) > 1 {
			out = out[:len(out)-1]
		}
		//		fmt.Println(out)
		__testOut1 = out
		return nil
	}, make([]interface{}, 0), int64(1))
}

func __createTestWorkflowFunction2() Executable {
	return NewActionExecutable("FUNC2", func(args ...interface{}) error {
		var out string = ""
		out = "FUNC2 -> "
		for _, v := range args {
			out = fmt.Sprintf("%s%v%s", out, v, " ")
		}
		if len(out) > 1 {
			out = out[:len(out)-1]
		}
		//		fmt.Println(out)
		__testOut2 = out
		return nil
	}, make([]interface{}, 0), int64(2))
}

func __createTesExecutables() []Executable {
	var out []Executable = make([]Executable, 0)
	out = append(out, __createTestWorkflowFunction1())
	return out
}

func __createTesExecutables2() []Executable {
	var out []Executable = make([]Executable, 0)
	out = append(out, __createTestWorkflowFunction2())
	return out
}

func __createTestCommands() []Feature {
	var out []Feature = make([]Feature, 0)
	out = append(out, Feature{
		Code:        "FEATURE1",
		Name:        "Feature 1",
		Description: "Test Feature, used to test flows.Feature component",
		ExecuteList: __createTesExecutables(),
		Index:       int64(1),
	})
	return out
}

func __createTestCommands2() []Feature {
	var out []Feature = make([]Feature, 0)
	out = append(out, Feature{
		Code:        "FEATURE2",
		Name:        "Feature 2",
		Description: "Test Feature, used to test flows.Feature component as Steps' Features neated hierarchy",
		ExecuteList: __createTesExecutables2(),
		Index:       int64(2),
	})
	return out
}

func __createTestSubSteps() []*Step {
	var out []*Step = make([]*Step, 0)
	out = append(out, &Step{
		Code:        "STEP2",
		Name:        "Step 2",
		Description: "Test Sub Step, used to test flows.Step component as Steps neated hierarchy",
		Index:       int64(2),
		Commands:    __createTestCommands2(),
		SubSteps:    make([]*Step, 0),
	})
	return out
}

func __createTestSteps() []Step {
	var out []Step = make([]Step, 0)
	out = append(out, Step{
		Code:        "STEP1",
		Name:        "Step 1",
		Description: "Test Step, used to test flows.Step component",
		Index:       int64(1),
		Commands:    __createTestCommands(),
		SubSteps:    __createTestSubSteps(),
	})
	return out
}

func __createTestPhases() []Phase {
	var out []Phase = make([]Phase, 0)
	out = append(out, Phase{
		Code:        "PHASE1",
		Name:        "Phase 1",
		Description: "Test Phase, used to test flows.Phase component",
		Index:       int64(1),
		Steps:       __createTestSteps(),
	})
	return out
}

func TestNewFeature(t *testing.T) {
	__testOut1 = ""
	__testOut2 = ""
	feature := NewFeature("FEATURE1", "Test Feature, used to test flows.Feature component", "Feature 1", int64(1), __createTesExecutables())
	expectedText := "Feature{Code: 'FEATURE1',Name: 'Feature 1', Description: 'Test Feature, used to test flows.Feature component', Index: 1, Running: [Action{Name: 'FUNC1', Func: <function>, Index: 1, Args: []}]}"
	currentText := feature.String()
	if expectedText != currentText {
		t.Fatal(fmt.Sprintf("Wrong value, feature String() doesn't match, Expected: <%s> but Given: <%s>", expectedText, currentText))
	}
	context := NewEmptyContext()

	info := __createTestUserInfo()
	session := Session{
		user:    &info,
		context: &context,
	}

	feature.Run(&session, __createTestArguments())

	expectedOut1 := "FUNC1 -> Sample text here"

	if expectedOut1 != __testOut1 {
		t.Fatal(fmt.Sprintf("Wrong value, feature func1 out doesn't match, Expected: <%s> but Given: <%s>", expectedOut1, __testOut1))
	}
}

func TestNewStep(t *testing.T) {
	__testOut1 = ""
	__testOut2 = ""
	step := NewStep("STEP1", "Test Step, used to test flows.Step component", "Step 1", int64(1), __createTestCommands(), __createTestSubSteps())
	expectedText := "Step{Code: 'STEP1',Name: 'Step 1', Description: 'Test Step, used to test flows.Step component', Index: 1, Commands: [Feature{Code: 'FEATURE1',Name: 'Feature 1', Description: 'Test Feature, used to test flows.Feature component', Index: 1, Running: [Action{Name: 'FUNC1', Func: <function>, Index: 1, Args: []}]}], SubSteps: [Step{Code: 'STEP2',Name: 'Step 2', Description: 'Test Sub Step, used to test flows.Step component as Steps neated hierarchy', Index: 2, Commands: [Feature{Code: 'FEATURE2',Name: 'Feature 2', Description: 'Test Feature, used to test flows.Feature component as Steps' Features neated hierarchy', Index: 2, Running: [Action{Name: 'FUNC2', Func: <function>, Index: 2, Args: []}]}], SubSteps: []}]}"
	currentText := step.String()
	if expectedText != currentText {
		t.Fatal(fmt.Sprintf("Wrong value, step String() doesn't match, Expected: <%s> but Given: <%s>", expectedText, currentText))
	}
	context := NewEmptyContext()

	info := __createTestUserInfo()
	session := Session{
		user:    &info,
		context: &context,
	}

	step.Run(&session, __createTestArguments())

	expectedOut1 := "FUNC1 -> Sample text here"

	if expectedOut1 != __testOut1 {
		t.Fatal(fmt.Sprintf("Wrong value, step func1 out doesn't match, Expected: <%s> but Given: <%s>", expectedOut1, __testOut1))
	}

	expectedOut2 := "FUNC2 -> Another sample text here"

	if expectedOut2 != __testOut2 {
		t.Fatal(fmt.Sprintf("Wrong value, step func2 out doesn't match, Expected: <%s> but Given: <%s>", expectedOut2, __testOut2))
	}
}

func TestNewPhase(t *testing.T) {
	__testOut1 = ""
	__testOut2 = ""
	phase := NewPhase("PHASE1", "Test Phase, used to test flows.Phase component", "Phase 1", int64(1), __createTestSteps())
	expectedText := "Phase{Code: 'PHASE1', Name: 'Phase 1', Description: 'Test Phase, used to test flows.Phase component', Index: 1, Steps: [Step{Code: 'STEP1',Name: 'Step 1', Description: 'Test Step, used to test flows.Step component', Index: 1, Commands: [Feature{Code: 'FEATURE1',Name: 'Feature 1', Description: 'Test Feature, used to test flows.Feature component', Index: 1, Running: [Action{Name: 'FUNC1', Func: <function>, Index: 1, Args: []}]}], SubSteps: [Step{Code: 'STEP2',Name: 'Step 2', Description: 'Test Sub Step, used to test flows.Step component as Steps neated hierarchy', Index: 2, Commands: [Feature{Code: 'FEATURE2',Name: 'Feature 2', Description: 'Test Feature, used to test flows.Feature component as Steps' Features neated hierarchy', Index: 2, Running: [Action{Name: 'FUNC2', Func: <function>, Index: 2, Args: []}]}], SubSteps: []}]}]}"
	currentText := phase.String()
	if expectedText != currentText {
		t.Fatal(fmt.Sprintf("Wrong value, phase String() doesn't match, Expected: <%s> but Given: <%s>", expectedText, currentText))
	}
	context := NewEmptyContext()

	info := __createTestUserInfo()
	session := Session{
		user:    &info,
		context: &context,
	}

	phase.Run(&session, __createTestArguments())

	expectedOut1 := "FUNC1 -> Sample text here"

	if expectedOut1 != __testOut1 {
		t.Fatal(fmt.Sprintf("Wrong value, phase func1 out doesn't match, Expected: <%s> but Given: <%s>", expectedOut1, __testOut1))
	}

	expectedOut2 := "FUNC2 -> Another sample text here"

	if expectedOut2 != __testOut2 {
		t.Fatal(fmt.Sprintf("Wrong value, phase func2 out doesn't match, Expected: <%s> but Given: <%s>", expectedOut2, __testOut2))
	}
}

func TestNewWorkflowController(t *testing.T) {
	__testOut1 = ""
	__testOut2 = ""
	workflow := NewWorkflowController("WKFL1", "Test Workflow used for test features of flows.Workflow component", "Test Workflow", __createTestPhases())
	expectedText := "Workflow{Code: 'WKFL1',Name: 'Test Workflow', Description: 'Test Workflow used for test features of flows.Workflow component', Phases: [Phase{Code: 'PHASE1', Name: 'Phase 1', Description: 'Test Phase, used to test flows.Phase component', Index: 1, Steps: [Step{Code: 'STEP1',Name: 'Step 1', Description: 'Test Step, used to test flows.Step component', Index: 1, Commands: [Feature{Code: 'FEATURE1',Name: 'Feature 1', Description: 'Test Feature, used to test flows.Feature component', Index: 1, Running: [Action{Name: 'FUNC1', Func: <function>, Index: 1, Args: []}]}], SubSteps: [Step{Code: 'STEP2',Name: 'Step 2', Description: 'Test Sub Step, used to test flows.Step component as Steps neated hierarchy', Index: 2, Commands: [Feature{Code: 'FEATURE2',Name: 'Feature 2', Description: 'Test Feature, used to test flows.Feature component as Steps' Features neated hierarchy', Index: 2, Running: [Action{Name: 'FUNC2', Func: <function>, Index: 2, Args: []}]}], SubSteps: []}]}]}]}"
	currentText := workflow.String()
	if expectedText != currentText {
		t.Fatal(fmt.Sprintf("Wrong value, workflow String() doesn't match, Expected: <%s> but Given: <%s>", expectedText, currentText))
	}
	context := NewEmptyContext()

	workflow.Run(&context, __createTestUserInfo(), __createTestArguments())

	expectedOut1 := "FUNC1 -> Sample text here"

	if expectedOut1 != __testOut1 {
		t.Fatal(fmt.Sprintf("Wrong value, workflow func1 out doesn't match, Expected: <%s> but Given: <%s>", expectedOut1, __testOut1))
	}

	expectedOut2 := "FUNC2 -> Another sample text here"

	if expectedOut2 != __testOut2 {
		t.Fatal(fmt.Sprintf("Wrong value, workflow func2 out doesn't match, Expected: <%s> but Given: <%s>", expectedOut2, __testOut2))
	}
}

func TestNewWorkflow(t *testing.T) {
	__testOut1 = ""
	__testOut2 = ""
	workflow := NewWorkflow("WKFL1", "Test Workflow used for test features of flows.Workflow component", "Test Workflow", __createTestPhases())
	expectedText := "Workflow{Code: 'WKFL1',Name: 'Test Workflow', Description: 'Test Workflow used for test features of flows.Workflow component', Phases: [Phase{Code: 'PHASE1', Name: 'Phase 1', Description: 'Test Phase, used to test flows.Phase component', Index: 1, Steps: [Step{Code: 'STEP1',Name: 'Step 1', Description: 'Test Step, used to test flows.Step component', Index: 1, Commands: [Feature{Code: 'FEATURE1',Name: 'Feature 1', Description: 'Test Feature, used to test flows.Feature component', Index: 1, Running: [Action{Name: 'FUNC1', Func: <function>, Index: 1, Args: []}]}], SubSteps: [Step{Code: 'STEP2',Name: 'Step 2', Description: 'Test Sub Step, used to test flows.Step component as Steps neated hierarchy', Index: 2, Commands: [Feature{Code: 'FEATURE2',Name: 'Feature 2', Description: 'Test Feature, used to test flows.Feature component as Steps' Features neated hierarchy', Index: 2, Running: [Action{Name: 'FUNC2', Func: <function>, Index: 2, Args: []}]}], SubSteps: []}]}]}]}"
	currentText := workflow.String()
	if expectedText != currentText {
		t.Fatal(fmt.Sprintf("Wrong value, workflow String() doesn't match, Expected: <%s> but Given: <%s>", expectedText, currentText))
	}
	context := NewEmptyContext()

	workflow.Run(&context, __createTestUserInfo(), __createTestArguments())

	expectedOut1 := "FUNC1 -> Sample text here"

	if expectedOut1 != __testOut1 {
		t.Fatal(fmt.Sprintf("Wrong value, workflow func1 out doesn't match, Expected: <%s> but Given: <%s>", expectedOut1, __testOut1))
	}

	expectedOut2 := "FUNC2 -> Another sample text here"

	if expectedOut2 != __testOut2 {
		t.Fatal(fmt.Sprintf("Wrong value, workflow func2 out doesn't match, Expected: <%s> but Given: <%s>", expectedOut2, __testOut2))
	}
}
