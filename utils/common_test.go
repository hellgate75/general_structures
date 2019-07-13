package utils

import (
	"fmt"
	"os"
	"testing"
)

func GetCurrentPath() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return pwd
}

func TestFileExists(t *testing.T) {
	exists := FileExists(fmt.Sprintf("%s%c%s", GetCurrentPath(), os.PathSeparator, "common_test.go"))
	if !exists {
		t.Fatal("Specified source file exists, so in case of FALSE there is an error!!")
	}
	exists = FileExists(fmt.Sprintf("%s%c%s", GetCurrentPath(), os.PathSeparator, "_common_test.go"))
	if exists {
		t.Fatal("Specified source file doesn't exist, so in case of TRUE there is an error!!")
	}
}

var sampleFolder string = fmt.Sprintf("%s%c%s%c%s%c%s", GetCurrentPath(), os.PathSeparator, "..", os.PathSeparator, "test", os.PathSeparator, "resources")
var sample1File string = fmt.Sprintf("%s%c%s", sampleFolder, os.PathSeparator, "test_create1.txt")
var sample2File string = fmt.Sprintf("%s%c%s", sampleFolder, os.PathSeparator, "test_create2.txt")
var sample3File string = fmt.Sprintf("%s%c%s", sampleFolder, os.PathSeparator, "test_create3.txt")

func TestCreateFileIfNotExists(t *testing.T) {
	err := CreateFolderIfNotExists(sampleFolder)
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	defer func() {
		os.RemoveAll(sampleFolder)
	}()
	err = CreateFileIfNotExists(sample1File)
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if !FileExists(sample1File) {
		t.Fatal(fmt.Sprintf("File '%s' has not been created", sample1File))
	}
	defer func(file string) {
		os.Remove(file)
	}(sample1File)
	err = CreateFileIfNotExists(fmt.Sprintf("%s%s", "..", sample1File))
	if err == nil {
		t.Fatal("Not arisen expected error !!")
	}
}

func TestCreateFile(t *testing.T) {
	errF := CreateFolderIfNotExists(sampleFolder)
	if errF != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", errF.Error()))
	}
	defer func() {
		os.RemoveAll(sampleFolder)
	}()
	file, err := CreateFile(sample2File)
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if !FileExists(sample2File) {
		t.Fatal(fmt.Sprintf("File '%s' has not been created", sample2File))
	}
	defer func(fileX string) {
		file.Close()
		os.Remove(fileX)
	}(sample2File)
	_, err = CreateFile(fmt.Sprintf("%s%s", "..", sample2File))
	if err == nil {
		t.Fatal("Not arisen expected error !!")
	}
}

func TestCreateFileAndUse(t *testing.T) {
	errF := CreateFolderIfNotExists(sampleFolder)
	if errF != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", errF.Error()))
	}
	defer func() {
		os.RemoveAll(sampleFolder)
	}()
	var consumer func(*os.File) (interface{}, error) = func(f *os.File) (interface{}, error) {
		f.Close()
		return nil, nil
	}
	_, err := CreateFileAndUse(sample3File, consumer)
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if !FileExists(sample3File) {
		t.Fatal(fmt.Sprintf("File '%s' has not been created", sample3File))
	}
	defer func(fileX string) {
		os.Remove(fileX)
	}(sample3File)
	_, err = CreateFileAndUse(fmt.Sprintf("%s%s", "..", sample3File), consumer)
	if err == nil {
		t.Fatal("Not arisen expected error !!")
	}
}

func TestDeleteIfExists(t *testing.T) {
	errF := CreateFolderIfNotExists(sampleFolder)
	if errF != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", errF.Error()))
	}
	defer func() {
		os.RemoveAll(sampleFolder)
	}()
	file, err := os.Create(sample1File)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unable to create file %s - error : %s", sample1File, err.Error()))
	}
	file.Close()
	err = DeleteIfExists(sample1File)
	if err != nil {
		t.Fatal(fmt.Sprintf("Unable to delete file %s - error : %s", sample1File, err.Error()))
	}
	if FileExists(sample1File) {
		t.Fatal(fmt.Sprintf("File '%s' has not been deleted", sample1File))
		defer func(file string) {
			os.Remove(file)
		}(sample1File)
	}
}

func __matchByteArrays(a []byte, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestEncodeBytes(t *testing.T) {
	var expected []byte = []byte{96, 81, 95, 96}
	var input string = "test"
	var out []byte = EncodeBytes([]byte(input))
	//fmt.Println(out)
	if !__matchByteArrays(expected, out) {
		t.Fatal(fmt.Sprintf("Wrong result in EncodeBytes, Expected <%v> but Given <%v>", expected, out))
	}
}

func TestDecodeBytes(t *testing.T) {
	var input []byte = []byte{96, 81, 95, 96}
	var expected []byte = []byte("test")
	var out []byte = DecodeBytes(input)
	//fmt.Println(out)
	if !__matchByteArrays(expected, out) {
		t.Fatal(fmt.Sprintf("Wrong result in DecodeBytes, Expected <%v> but Given <%v>", expected, out))
	}
}

func TestCorrectInput(t *testing.T) {
	var expected string = "test"
	var input string = "  Test  "
	output := CorrectInput(input)
	if expected != output {
		t.Fatal(fmt.Sprintf("Wrong result in CorrectInput, Expected <%v> but Given <%v>", expected, output))
	}
}

func TestStringToInt(t *testing.T) {
	var expected int = 12
	var input string = "12"
	output, err := StringToInt(input)
	if err != nil {
		t.Fatal(fmt.Sprintf("Arisen unexpected error : %s", err.Error()))
	}
	if expected != output {
		t.Fatal(fmt.Sprintf("Wrong result in CorrectInput, Expected <%d> but Given <%d>", expected, output))
	}
}
func TestIntToString(t *testing.T) {
	var input int = 12
	var expected string = "12"
	output := IntToString(input)
	if expected != output {
		t.Fatal(fmt.Sprintf("Wrong result in CorrectInput, Expected <%s> but Given <%s>", expected, output))
	}
}
