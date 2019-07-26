package flows

import (
	"fmt"
	"github.com/hellgate75/general_utils/streams"
	"os"
	"testing"
)

func TestNewFileFlow(t *testing.T) {
	path := streams.GetCurrentPath()
	file := fmt.Sprintf("%s%s%s%s%s%s%s", path, fileSep, "..", fileSep, "test", fileSep, "text.txt")
	defer func(file string) {
		os.Remove(file)
	}(file)
	var fileFlow Flow
	var errFlow error
	fileFlow, errFlow = NewFileFlow(file, 1024)
	if errFlow != nil {
		t.Fatal(fmt.Sprintf("Unexpected error creating new file flow, message: %s", errFlow.Error()))
	}
	var expectedBytes []byte = []byte("This is a sample text!!")
	num, err := fileFlow.Write(expectedBytes)
	expectedN := len(expectedBytes)
	if num < expectedN {
		t.Fatal(fmt.Sprintf("Error writing flow, written data insufficient: %d < %d", num, expectedN))
	}
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error writing data into flow, message: %s", err.Error()))
	}
	err = fileFlow.Close()
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error closing the flow, message: %s", err.Error()))
	}
	err = fileFlow.Open(file)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error opening the flow, message: %s", err.Error()))
	}
	err = fileFlow.Reset()
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error resetting the flow, message: %s", err.Error()))
	}
	var p []byte = make([]byte, 23)
	num, err = fileFlow.Read(p)
	expectedN = len(expectedBytes)
	if num < expectedN {
		t.Fatal(fmt.Sprintf("Error reading flow, read data insufficient %d < %d", num, expectedN))
	}
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error reading data from the flow, message: %s", err.Error()))
	}

}

func TestNewBufferFlow(t *testing.T) {
	var expectedBytes []byte = []byte("This is a sample text!!")
	bufferFlow := NewByteArrayFlow(expectedBytes, 12)
	var buffer []byte = make([]byte, 23)
	num, errN := bufferFlow.Read(buffer)
	expectedN := len(expectedBytes)
	if num < expectedN {
		t.Fatal(fmt.Sprintf("Error reading flow, read data insufficient: %d < %d", num, expectedN))
	}
	if errN != nil {
		t.Fatal(fmt.Sprintf("Unexpected error reading data into flow, message: %s", errN.Error()))
	}

	bufferFlow = NewByteArrayFlow(make([]byte, 0), 12)
	num, errN = bufferFlow.Write(expectedBytes)
	expectedN = len(expectedBytes)
	if num < expectedN {
		t.Fatal(fmt.Sprintf("Error writing flow, write data insufficient: %d < %d", num, expectedN))
	}
	if errN != nil {
		t.Fatal(fmt.Sprintf("Unexpected error writing data into flow, message: %s", errN.Error()))
	}
	text1 := fmt.Sprintf("%s", bufferFlow.Bytes())
	text2 := bufferFlow.String()
	expectedText := fmt.Sprintf("%s", expectedBytes)
	if text1 != expectedText {
		t.Fatal(fmt.Sprintf("Error in Flow Buffer, Expected: <%s> but Given: <%s>", expectedText, text1))
	}
	if text2 != expectedText {
		t.Fatal(fmt.Sprintf("Error in Flow Buffer, Expected: <%s> but Given: <%s>", expectedText, text2))
	}
}

func TestNewFlowAndByBuffer(t *testing.T) {
	var expectedBytes []byte = []byte("This is a sample text!! ")
	buffFlow := NewFlowBuff(24)
	for i := 0; i < 10; i++ {
		array := expectedBytes
		num, errN := buffFlow.Write(array)
		expectedN := len(array)
		if errN != nil {
			t.Fatal(fmt.Sprintf("Unexpected error writing data into flow Write # %d, message: %s", i, errN.Error()))
		}
		if num < expectedN {
			t.Fatal(fmt.Sprintf("Error writing flow Write # %d, write data insufficient: %d < %d", i, num, expectedN))
		}
	}
	for i := 0; i < 10; i++ {
		var buffer []byte = make([]byte, 24)
		num, errN := buffFlow.Read(buffer)
		expectedN := len(expectedBytes)
		if num < expectedN {
			t.Fatal(fmt.Sprintf("Error reading flow Read # %d, write data insufficient: %d < %d", i, num, expectedN))
		}
		if errN != nil {
			t.Fatal(fmt.Sprintf("Unexpected error reading data into flow  Read # %d, i, message: %s", i, errN.Error()))
		}
		text1 := fmt.Sprintf("%s", buffer)
		expectedText := fmt.Sprintf("%s", expectedBytes)
		if text1 != expectedText {
			t.Fatal(fmt.Sprintf("Error in Flow Buffer Read # %d, Expected: <%s> but Given: <%s>", i, expectedText, text1))
		}
	}
	simpleFlow := NewFlow()
	for i := 0; i < 20; i++ {
		simpleFlow.Write(expectedBytes)
	}
	var expectedSize int64 = int64(480)
	if simpleFlow.Size() != expectedSize {
		t.Fatal(fmt.Sprintf("Error in expected Size, Expected: <%d> but Given: <%d>", expectedSize, simpleFlow.Size()))
	}
	simpleFlow.Seek(10, 0)
	simpleFlow.Seek(14, 1)
	bufferN := make([]byte, 24)
	simpleFlow.Read(bufferN)
	currentText := fmt.Sprintf("%s", bufferN)
	expectedText := fmt.Sprintf("%s", expectedBytes)
	if currentText != expectedText {
		t.Fatal(fmt.Sprintf("Wrong Next Seek Read value , Expected: <%s> but Given: <%s>", expectedText, currentText))
	}
}

func TestNewFlowReadFrom(t *testing.T) {
	expectedBytes := []byte("This is a sample text!! ")
	simpleFlow := NewFlow()
	for i := 0; i < 20; i++ {
		simpleFlow.Write(expectedBytes)
	}
	simpleFlow2 := NewFlow()
	nVal, err := simpleFlow2.ReadFrom(simpleFlow)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error reading from another flow, message: %s", err.Error()))
	}
	var expectedSize int64 = int64(480)
	if int64(nVal) < expectedSize {
		t.Fatal(fmt.Sprintf("Error in expected Transferred bytes, Expected: <%d> but Given: <%d>", expectedSize, nVal))
	}
	if simpleFlow.Size() != expectedSize {
		t.Fatal(fmt.Sprintf("Error in expected Size, Expected: <%d> but Given: <%d>", expectedSize, simpleFlow.Size()))
	}

}

func TestNewFlowWriteTo(t *testing.T) {
	expectedBytes := []byte("This is a sample text!! ")
	simpleFlow := NewFlow()
	for i := 0; i < 20; i++ {
		simpleFlow.Write(expectedBytes)
	}
	simpleFlow2 := NewFlow()
	nVal, err := simpleFlow.WriteTo(simpleFlow2)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unexpected error reading from another flow, message: %s", err.Error()))
	}
	var expectedSize int64 = int64(480)
	if int64(nVal) < expectedSize {
		t.Fatal(fmt.Sprintf("Error in expected Transferred bytes, Expected: <%d> but Given: <%d>", expectedSize, nVal))
	}
	if simpleFlow.Size() != expectedSize {
		t.Fatal(fmt.Sprintf("Error in expected Size, Expected: <%d> but Given: <%d>", expectedSize, simpleFlow.Size()))
	}

}
