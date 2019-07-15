## package types // import "github.com/hellgate75/general_structures/types"


### FUNCTIONS

#### func CorrectInput(input string) string
    Corrects input string with space trim and lowering the case 
#####    Parameters:
         input (string) input string
#####     Returns:
         string Represent the corrected string

#### func CreateFile(file string) (*os.File, error)
    Create a New file overwriting 
#####    Parameters:
        file (string) input file name
#####    Returns:
        error Any error that can occur during the computation

#### func CreateFileAndUse(file string, consumer func(*os.File) (interface{}, error)) (interface{}, error)
    Create a New file and consume content 
#####    Parameters:
        file (string) input file name
        consumer (unc(*os.File) (interface{}, error)) function that transform file in a final structure
#####    Returns:
    (interface{} outcome of the file content computation,
    error Any error that can occur during the computation)

#### func CreateFileIfNotExists(file string) error
    Create a New file after cheking existance of same 
#####    Parameters:
        file (string) input file name
#####    Returns:
        error Any error that can occur during the computation

#### func DecodeBytes(encodedByteArray []byte) []byte
    Decode Byte Array from internal format 
#####    Parameters:
        encodedByteArray ([]byte) byte array to be decoded
#####    Returns:
        byte[] Output decoded byte array

#### func DeleteIfExists(file string) error
    Delete a New file if existsts in the FileSystem 
#####    Parameters:
        file (string) input file name
#####    Returns:
        error Any error that can occur during the computation

#### func EncodeBytes(decodedByteArray []byte) []byte
    Encode Byte Array in internal format 
#####    Parameters:
        decodedByteArray ([]byte) byte array to be encoded
#####    Returns:
        byte[] Output encoded byte array

#### func FileExists(file string) bool
    Verify existance of a file 
#####    Parameters:
        file (string) input file name
#####    Returns:
        bool File existance feedback

#### func InitLogger()
    Initialize Package logger

#### func IntToString(n int) string
    Convert an integer to string 
#####    Parameters:
        n (input) input integer
#####    Returns:
        string Represent the converted integer to string format

#### func MakeFolderIfNotExists(folder string) error
    Make a new folder if not existsts in the FileSystem 
#####    Parameters:
        folder (string) input folder name
#####    Returns:
        error Any error that can occur during the computation

#### func StringToInt(s string) (int, error)
    Convert a string to integer 
#####    Parameters:
        s (string) input string
#####    Returns:
    ( int output converted integer,
        error Any error that can occur during the computation)


### TYPES

##### type ArrayNav interface {
#####	BaseArrayNav
#####	Get() common.Type
##### }
    Generic Type Array Navigator Interface

##### func NewArrayNav(arr []common.Type) ArrayNav
    Create New Generic Type Array Navigator 
#####     Parameters:
        arr ([]common.Type) input Array to manage
#####     Returns:
        ArrayNav Array Navigator feature for the specified type

##### type BaseArrayNav interface {
##### 	Prev() bool
##### 	Next() bool
##### 	Len() int
##### 	Position() int
##### }
    Base Array Navigator Interface

##### type BoolArrayNav interface {
##### 	BaseArrayNav
##### 	Get() bool
##### }
    Boolean Array Navigator Interface

##### func NewBoolArrayNav(arr []bool) BoolArrayNav
    Create New Boolean Array Navigator 
#####     Parameters:
        arr ([]bool) input Array to manage
#####     Returns:
        BoolArrayNav Array Navigator feature for the specified type

##### type BoolNavAttr struct {
##### 	// Has unexported fields.
##### }
    Boolean Array Navigator structure

##### func (nav *BoolNavAttr) Get() bool
    Get current Element in the Array 
#####     Returns:
        bool Current Element or false in case of error

##### func (nav *BoolNavAttr) Len() int
    Get current Array length 
#####     Returns:
        int Length of the Array

##### func (nav *BoolNavAttr) Next() bool
    Move next Element in the Array 
#####     Returns:
        bool Next command success state

##### func (nav *BoolNavAttr) Position() int
    Get current position in the Array 
#####     Returns:
        int Position of cursor in the Array

##### func (nav *BoolNavAttr) Prev() bool
    Move previous Element in the Array 
#####     Returns:
        bool Prev command success state

##### type FloatArrayNav interface {
##### 	BaseArrayNav
##### 	Get() float64
##### }
    Float Array Navigator Interface
##### func NewFloatArrayNav(arr []float64) FloatArrayNav
    Create New Float Array Navigator 
#####     Parameters:
        arr ([]float64) input Array to manage
#####     Returns:
        FloatArrayNav Array Navigator feature for the specified type

##### type FloatNavAttr struct {
##### 	// Has unexported fields.
##### }
    Float Array Navigator structure

##### func (nav *FloatNavAttr) Get() float64
    Get current Element in the Array 
#####     Returns:
        float64 Current Element or 0.0 in case of error

##### func (nav *FloatNavAttr) Len() int
    Get current Array length 
#####     Returns:
        int Length of the Array

##### func (nav *FloatNavAttr) Next() bool
    Move next Element in the Array 
#####     Returns:
        bool Next command success state

##### func (nav *FloatNavAttr) Position() int
    Get current position in the Array 
#####     Returns:
        int Position of cursor in the Array

##### func (nav *FloatNavAttr) Prev() bool
    Move previous Element in the Array 
#####     Returns:
        bool Prev command success state

##### type IntArrayNav interface {
##### 	BaseArrayNav
##### 	Get() int
##### }
    Integer Array Navigator Interface

##### func NewIntArrayNav(arr []int) IntArrayNav
    Create New Integer Array Navigator 
#####     Parameters:
        arr ([]int) input Array to manage
#####     Returns:
        IntArrayNav Array Navigator feature for the specified type

##### type IntNavAttr struct {
##### 	// Has unexported fields.
##### }
    Integer Array Navigator structure

##### func (nav *IntNavAttr) Get() int
    Get current Element in the Array 
#####     Returns:
        int Current Element or 0 in case of error

##### func (nav *IntNavAttr) Len() int
    Get current Array length 
#####     Returns:
        int Length of the Array

##### func (nav *IntNavAttr) Next() bool
    Move next Element in the Array 
#####     Returns:
        bool Next command success state

##### func (nav *IntNavAttr) Position() int
    Get current position in the Array 
#####     Returns:
        int Position of cursor in the Array

##### func (nav *IntNavAttr) Prev() bool
    Move previous Element in the Array 
#####     Returns:
        bool Prev command success state

##### type NavAttr struct {
##### 	// Has unexported fields.
##### }
    Generic Array Navigator structure

##### func (nav *NavAttr) Get() common.Type
    Get current Element in the Array 
#####     Returns:
        common.Type Current Element or nil in case of error

##### func (nav *NavAttr) Len() int
    Get current Array length 
#####     Returns:
        int Length of the Array

##### func (nav *NavAttr) Next() bool
    Move next Element in the Array 
#####     Returns:
        bool Next command success state

##### func (nav *NavAttr) Position() int
    Get current position in the Array 
#####     Returns:
        int Position of cursor in the Array

##### func (nav *NavAttr) Prev() bool
    Move previous Element in the Array 
#####     Returns:
        bool Prev command success state

##### func (nav *NavAttr) Print()
    Print current Element in the Array

