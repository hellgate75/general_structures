package types

import (
	"github.com/hellgate75/general_utils/log"
	"os"
	"strconv"
	"strings"
)

var logger log.Logger

//Initialize package logger if not started
func InitLogger() {
	currentLogger, err := log.New("types")
	if err != nil {
		panic(err.Error())
	}
	logger = currentLogger
}

//Verify existance of a file
// Parameters:
//   file (string) input file name
// Returns:
//   bool File existance feedback
func FileExists(file string) bool {
	_, err := os.Stat(file)
	return err == nil
}

//Create a New file after cheking existance of same
// Parameters:
//   file (string) input file name
// Returns:
//   error Any error that can occur during the computation
func CreateFileIfNotExists(file string) error {
	if FileExists(file) {
		return nil
	}
	emptyFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		f.Close()
	}(emptyFile)
	return nil
}

//Create a New file after cheking existance of same
// Parameters:
//   folder (string) input folder name
// Returns:
//   error Any error that can occur during the computation
func CreateFolderIfNotExists(folder string) error {
	if FileExists(folder) {
		return nil
	}
	err := os.MkdirAll(folder, 0666)
	if err != nil {
		return err
	}
	return nil
}

//Create a New file overwriting
// Parameters:
//   file (string) input file name
// Returns:
//   error Any error that can occur during the computation
func CreateFile(file string) (*os.File, error) {
	if FileExists(file) {
		return os.Open(file)
	}
	emptyFile, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	return emptyFile, nil
}

//Create a New file and consume content
// Parameters:
//   file (string) input file name
//   consumer (unc(*os.File) (interface{}, error)) function that transform file in a final structure
// Returns:
//   (interface{} outcome of the file content computation,
//   error Any error that can occur during the computation)
func CreateFileAndUse(file string, consumer func(*os.File) (interface{}, error)) (interface{}, error) {
	var emptyFile *os.File
	var err error
	if !FileExists(file) {
		emptyFile, err = os.Create(file)
	} else {
		emptyFile, err = os.Open(file)
	}
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		f.Close()
	}(emptyFile)
	return consumer(emptyFile)
}

//Delete a New file if existsts in the FileSystem
// Parameters:
//   file (string) input file name
// Returns:
//   error Any error that can occur during the computation
func DeleteIfExists(file string) error {
	_, err := os.Stat(file)
	if err == nil {
		return os.Remove(file)
	}
	return err
}

//Make a new folder if not existsts in the FileSystem
// Parameters:
//   folder (string) input folder name
// Returns:
//   error Any error that can occur during the computation
func MakeFolderIfNotExists(folder string) error {
	if _, err := os.Stat(folder); err != nil {
		err := os.MkdirAll(folder, 0777)
		return err
	}
	return nil
}

//Encode Byte Array in internal format
// Parameters:
//   decodedByteArray ([]byte) byte array to be encoded
// Returns:
//   byte[] Output encoded byte array
func EncodeBytes(decodedByteArray []byte) []byte {
	newBytes := make([]byte, 0)
	for _, byteElem := range decodedByteArray {
		newBytes = append(newBytes, byteElem-20)
	}
	return newBytes
	//return decodedByteArray
}

//Decode Byte Array from internal format
// Parameters:
//   encodedByteArray ([]byte) byte array to be decoded
// Returns:
//   byte[] Output decoded byte array
func DecodeBytes(encodedByteArray []byte) []byte {
	newBytes := make([]byte, 0)
	for _, byteElem := range encodedByteArray {
		newBytes = append(newBytes, byteElem+20)
	}
	return newBytes
}

//Corrects input string with space trim and lowering the case
// Parameters:
//   input (string) input string
// Returns:
//   string Represent the corrected string
func CorrectInput(input string) string {
	return strings.TrimSpace(strings.ToLower(input))
}

//Convert a string to integer
// Parameters:
//   s (string) input string
// Returns:
//   ( int output converted integer,
//   error Any error that can occur during the computation)
func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

//Convert an integer to string
// Parameters:
//   n (input) input integer
// Returns:
//   string Represent the converted integer to string format
func IntToString(n int) string {
	return strconv.Itoa(n)
}
