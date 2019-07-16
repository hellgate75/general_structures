## package types // import "github.com/hellgate75/general_structures/types"


### CONSTANTS

const (
	SAMPLE_SIZE int = 5
)

### FUNCTIONS

#### func CorrectInput(input string) string
    Corrects input string with space trim and lowering the case
#####     Parameters:
       input (string) input string
#####     Returns:
       string Represent the corrected string

#### func CreateFile(file string) (*os.File, error)
    Create a New file overwriting
#####     Parameters:
       file (string) input file name
#####     Returns:
       error Any error that can occur during the computation

#### func CreateFileAndUse(file string, consumer func(*os.File) (interface{}, error)) (interface{}, error)
    Create a New file and consume content
#####     Parameters:
       file (string) input file name
       consumer (unc(*os.File) (interface{}, error)) function that transform file in a final structure
#####     Returns:
       (interface{} outcome of the file content computation,
       error Any error that can occur during the computation)

#### func CreateFileIfNotExists(file string) error
    Create a New file after cheking existance of same
#####     Parameters:
       file (string) input file name
#####     Returns:
       error Any error that can occur during the computation

#### func CreateFolderIfNotExists(folder string) error
    Create a New file after cheking existance of same
#####     Parameters:
       folder (string) input folder name
#####     Returns:
       error Any error that can occur during the computation

#### func DecodeBytes(encodedByteArray []byte) []byte
    Decode Byte Array from internal format
#####     Parameters:
       encodedByteArray ([]byte) byte array to be decoded
#####     Returns:
       byte[] Output decoded byte array

#### func DeleteIfExists(file string) error
    Delete a New file if existsts in the FileSystem
#####     Parameters:
       file (string) input file name
#####     Returns:
       error Any error that can occur during the computation

#### func EncodeBytes(decodedByteArray []byte) []byte
    Encode Byte Array in internal format
#####     Parameters:
       decodedByteArray ([]byte) byte array to be encoded
#####     Returns:
       byte[] Output encoded byte array

#### func FileExists(file string) bool
    Verify existance of a file
#####     Parameters:
       file (string) input file name
#####     Returns:
       bool File existance feedback

#### func InitLogger()
    Initialize Package logger

#### func IntToString(n int) string
    Convert an integer to string
#####     Parameters:
       n (input) input integer
#####     Returns:
       string Represent the converted integer to string format

#### func MakeFolderIfNotExists(folder string) error
    Make a new folder if not existsts in the FileSystem
    Parameters:
       folder (string) input folder name
    Returns:
       error Any error that can occur during the computation

#### func StringToInt(s string) (int, error)
    Convert a string to integer
#####     Parameters:
       s (string) input string
#####     Returns:
       ( int output converted integer,
       error Any error that can occur during the computation)


### TYPES

##### type ArrayArgumentFunc func(arg0 ...RowElement) Iterator
     Function used to flatten mapping arrays or plurals in general to the list of elements inot the stream
##### var (
##### 	//Function that map one or more List(s) to exploded contained elements, used in types.Stream.FlatMap()
##### 	LIST_TO_SINGULAR_FUNC ArrayArgumentFunc = __listToSingular
##### 	//Function that map one or more Collection(s) to exploded contained elements, used in types.Stream.FlatMap()
##### 	COLLECTION_TO_SINGULAR_FUNC ArrayArgumentFunc = __collectionToSingular
##### 	//Function that map one or more Iterator(s) to exploded contained elements, used in types.Stream.FlatMap()
##### 	ITERATOR_TO_SINGULAR_FUNC ArrayArgumentFunc = __iteratorToSingular
##### 	//Function that map one or more Array of Arrays(s) to exploded contained elements, used in types.Stream.FlatMap()
##### 	ARRAYOFARRAY_TO_SINGULAR ArrayArgumentFunc = __arrayOfArrayToSingular
##### 	//Function that map one or more Array of Elements to exploded contained elements, used in types.Stream.FlatMap()
##### 	ARRAY_TO_SINGULAR_FUNC ArrayArgumentFunc = __arrayToSingular
##### )

##### type ArrayNav interface {
##### 	BaseArrayNav
##### 	Get() common.Type
##### }
    Generic Type Array Navigator Interface

##### func NewArrayNav(arr []common.Type) ArrayNav
    Create New Generic Type Array Navigator
##### 	Parameters:
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

##### type BiArgumentsFunc func(arg0 RowElement, arg1 RowElement) RowElement
      Function used to accumulate elements from a Collections or Streams to a singlw element (plurals to single)
      
##### type BoolArrayNav interface {
##### 	BaseArrayNav
##### 	Get() bool
##### }
    Boolean Array Navigator Interface

##### func NewBoolArrayNav(arr []bool) BoolArrayNav
    Create New Boolean Array Navigator
#####      Parameters:
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
#####     Get current Array length
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

##### type Collection interface {
##### 	// Retrieve Component Element Type
##### 	// Returns:
##### 	//   reflect.Type Component Type
##### 	GetType() reflect.Type
##### 	// Returns the array iterator providing elements in the Collection
##### 	// Returns:
##### 	//    Iterator Array iterator component that contains all list elements
##### 	Iterator() Iterator
##### 	// Verify if the Collection is Empty
##### 	// Returns:
##### 	//    bool State that desribes if the Collection is Empty
##### 	Empty() bool
##### 	// Add a new element at the end of the Collection
##### 	// Parameters:
##### 	//    item (types.RowElement) Element to have been added into the list
##### 	// Returns:
##### 	//    bool Added success indicator
##### 	Add(item RowElement) bool
##### 	// Add all elements in the Collection
##### 	// Parameters:
##### 	//    items (types.RowArray) Elements to have been added into the Collection
##### 	// Returns:
##### 	//    bool Added success indicator
##### 	AddAll(items RowArray) bool
##### 	// Add all elements in the Collection
##### 	// Parameters:
##### 	//    collection (types.Collection) Colelction of elements to have been added into the Collection
##### 	// Returns:
##### 	//    bool Added success indicator
##### 	AddCollection(collection Collection) bool
##### 	// Remove an element from the Collection
##### 	// Parameters:
##### 	//    item (types.RowElement) Element to have been removed from the Collection
##### 	// Returns:
##### 	//    bool Added success indicator
##### 	Remove(item RowElement) bool
##### 	// Remove all occurancies of an element from the Collection
##### 	// Parameters:
##### 	//    item (types.RowElement) Element to have been removed from the Collection
##### 	// Returns:
##### 	//    int64 Number of elements removed from the Collection
##### 	RemoveAll(item RowElement) int64
##### 	// Returns:
##### 	//    types.RowElement First Element of the Collection
##### 	First() RowElement
##### 	// Get Last Element of the Collection
##### 	// Returns:
##### 	//    types.RowElement Last Element of the Collection
##### 	Last() RowElement
##### 	// Clear All Elements of the Collection and reset the collection as Empty
##### 	Clear()
##### 	// Print resume of the Collection
##### 	// Returns:
##### 	//    string Collection descriptive list
##### 	String() string
##### 	// Return a Stream within Collection elements
##### 	// Returns:
##### 	//    types.Stream Stream containing List elements
##### 	Stream() Stream
##### }
    Collection Component containing unlimited size of elements and providing multiple features

##### func CloneCollection(coll Collection) Collection
    Clone a collection, creating a new instamce same to origin one
#####     Parameters:
       coll (types.Collection) Origin Collection component instance
#####     Returns:
       types.Collection Collection component instance

##### func IteratorToCollection(iterator Iterator) Collection
    Create a collection, using an Iterator component instance
#####     Parameters:
       iterator (types.Iterator) Origin Iterator component instance
#####     Returns:
       types.Collection Collection component instance

##### func NewCollection(t reflect.Type) Collection
    Create New Collection component 
#####     Parameters:
       t (reflect.Type) Component Type
#####     Returns:
       types.Collection Collection component instance

##### func NewCollectionWithArray(t reflect.Type, arr RowArray) Collection
    Create New Collection component filled with elements from an array
#####     Parameters:
       t (reflect.Type) Component Type
       arr (types.RowArray) Array of generic type elements
#####     Returns:
       types.Collection Collection component instance

##### type CollectionElement struct {
##### 	Element RowElement
##### 	Next    *CollectionElement
##### }
    Element that Contains information about data and next node, it's used to maintain a hierarchy and sequence into the Collections components
    
##### func ArrayToCollectionElement(arr RowArray) (reflect.Type, *CollectionElement)
	Tranforms an array to types.CollectionElement, used for filling internally Collections and derivates
##### 	 Parameters:
	     arr (type.RowArray) Array of data
##### 	 Returns:
	    (reflect.Type List item Type,
	    *types.CollectionElement Pointer to Root of new Data Sequence)

##### type ComparatorFunc func(first RowElement, second RowElement) int
    Comparator functin that returns 1, 0, -1 in case first item suites bigger,
    same or lower position in the list

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

##### type Iterator interface {
##### 	// Assign Component Element Type
##### 	// Parameters:
##### 	//   t (reflect.Type) Component Type
##### 	Type(t reflect.Type)
##### 	// Retrieve Component Element Type
##### 	// Returns:
##### 	//   reflect.Type Component Type
##### 	GetType() reflect.Type
##### 	// Retrieve if iterator has a next element to merge
##### 	// Returns:
##### 	//   bool Iterator next
##### 	HasNext() bool
##### 	// Retrieve if iterator has a next element
##### 	// Returns:
##### 	//   types.RowElement next element if present or nil elsewise
##### 	Next() RowElement
##### }
    Collection Component containing unlimited size of elements and providing
    mulitple features

##### func NewIterator(t reflect.Type, rootElement *CollectionElement) Iterator
    Create New Iterator based on the Root of a Collection Elements hierarchy
#####     Parameters:
       t (reflect.Type) Component Type
       rootElement (*CollectionElement) Pointer to Collection Elements hierarchy root item
#####     Returns:
       types.Iterator Iterator component instance

##### type List interface {
##### 	Collection
##### 	// Sort list by Equals() bool method of the interface component or majority comparator
##### 	// Parameters:
##### 	//    compare (types.ComparatorFunc) Function used to compare elements in the List
##### 	// Returns:
##### 	//    bool Sort success indicator
##### 	Sort(compare ComparatorFunc) bool
##### 	// Verify if an element is inlcuded in the list using equality comparator
##### 	// Parameters:
##### 	//    item (types.RowElement) Element to have been seeked for into the list
##### 	// Returns:
##### 	//    bool Contains success indicator
##### 	Contains(item RowElement) bool
##### 	// Verify if an element is inlcuded in the list using MatcherFunc
##### 	// Parameters:
##### 	//    item (types.RowElement) Element to have been seeked for into the list
##### 	//    matcher (types.MatcherFunc) Function that equals
##### 	// Returns:
##### 	//    bool Contains success indicator
##### 	ContainsAs(item RowElement, matcher MatcherFunc) bool
##### 	// Get an element based on his index, or panic in case index is out of range
##### 	// Parameters:
##### 	//    index (int64) Element index in the List
##### 	// Returns:
##### 	//    types.RowElement Element found in the List
##### 	Get(index int64) RowElement
##### 	// Set an element based on his index, or panic in case index is out of range
##### 	// Parameters:
##### 	//    index (int64) Element index in the List
##### 	//    item (types.RowElement) Element to have been setted up into the list
##### 	// Returns:
##### 	//    bool Operation susccess status
##### 	Set(index int64, item RowElement) bool
##### 	// Add an element after specified index, or panic in case index is out of range
##### 	// Parameters:
##### 	//    index (int64) Element index in the List
##### 	//    item (types.RowElement) Element to have been setted up into the list
##### 	// Returns:
##### 	//    bool Operation susccess status
##### 	AddAfter(index int64, item RowElement) bool
##### 	// Add an element after specified index, or panic in case index is out of range
##### 	// Parameters:
##### 	//    index (int64) Element index in the List
##### 	//    items (types.RowArray) Element list to have been setted up into the list
##### 	// Returns:
##### 	//    int64 Number of elements Added into the List
##### 	AddAllAfter(index int64, items RowArray) int64
##### 	// Add an element after specified index, or panic in case index is out of range
##### 	// Parameters:
##### 	//    index (int64) Element index in the List
##### 	//    coll (types.Collection) Collection of Elements to have been setted up into the list
##### 	// Returns:
##### 	//    int64 Number of elements Added into the List
##### 	AddCollectionAfter(index int64, coll Collection) int64
##### 	// Remove Element at specified index, or panic in case index is out of range
##### 	// Parameters:
##### 	//    index (int64) Element index in the List
##### 	// Returns:
##### 	//    bool Operation susccess status
##### 	RemoveAt(index int64) bool
##### 	// Returns the size of the List
##### 	// Returns:
##### 	//    int64 Number of elements in the List
##### 	Size() int64
##### 	// Returns a sub list of the current List
##### 	// Parameters:
##### 	//    start (int64) First Element index in the new Sub List
##### 	//    end (int64) Last Element index in the new Sub List (excluded)
##### 	// Returns:
##### 	//    List New List containing elements between start (included) and end (excluded)
##### 	SubList(start int64, end int64) List
##### }
    List Component containing unlimited size of elements and providing mulitple
    features

##### func CollectionAsList(coll Collection) List
    Create a List using a collection instance
#####     Parameters:
       coll (types.Collection) Origin Collection component instance
#####     Returns:
       types.List List component instance

##### func IteratorAsList(iterator Iterator) List
    Create a List, using an Iterator component instance
#####     Parameters:
       iterator (types.Iterator) Origin Iterator component instance
#####     Returns:
       types.List List component instance

##### func NewList(t reflect.Type) List
    Create New List component
#####     Parameters:
       t (reflect.Type) Component Type
#####     Returns:
       types.List List component instance

##### func NewListWithArray(t reflect.Type, arr RowArray) List
    Create New List component filled with elements from an array
#####     Parameters:
       t (reflect.Type) Component Type
       arr (types.RowArray) Array of generic type elements
#####     Returns:
       types.List List component instance

##### type MapperFunc func(arg0 RowElement) RowElement
    Function used to map en element to another into Collections or Streams

##### type MatcherFunc func(first RowElement, second RowElement) bool
    Returns if an element is equal to another

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

##### type RowArray []RowElement
    Generic Array

##### type RowElement interface{}
    Generic Item

##### type Stream interface {
##### 	// Define parallel stream computation, accordingly to max number of CPUs
##### 	// Returns:
##### 	//    types.Stream Parallel Stream
##### 	Parallel() Stream
##### 	// Varify if current is a Parallel Stream
##### 	// Returns:
##### 	//    bool Parallel Stream Status
##### 	IsParallel() bool
##### 	// Map elements to other ones, maintaining the same number of elements in the Stream
##### 	// Parameters:
##### 	//    mapper (types.MapperFunc) Function that map one element to another one
##### 	// Returns:
##### 	//    types.Stream New Elements Stream
##### 	Map(mapper MapperFunc) Stream
##### 	// Map elements to other flat ones (when elements are Collections, Arrays, Or Plurals generally of elements), maintaining the same number of elements in the Stream
##### 	// Parameters:
##### 	//    mapper (types.ArrayArgumentFunc) Function that map a plural element to a singular one
##### 	// Returns:
##### 	//    types.Stream New Elements Stream
##### 	// Raises:
##### 	//    panic in case of not plural element list
##### 	FlatMap(mapper ArrayArgumentFunc) Stream
##### 	// Reduce elements to other flat ones (when elements are Collections, Arrays, Or Plurals generally of elements), maintaining the same number of elements in the Stream
##### 	// Parameters:
##### 	//    mapper (types.ArrayArgumentFunc) Function that map a plural element to a singular one
##### 	// Returns:
##### 	//    types.Stream New Elements Stream
##### 	Reduce(accumulationElement RowElement, reducer BiArgumentsFunc) RowElement
##### 	// Return a new List within stream content
##### 	// Returns:
##### 	//    types.List New Elements List instance
##### 	ToList() List
##### 	// Return a new Collection within stream content
##### 	// Returns:
##### 	//    types.Collection New Elements Collection instance
##### 	ToCollection() Collection
##### 	// Return a new Iterator within stream content
##### 	// Returns:
##### 	//    types.Iterator New Elements Iterator instance
##### 	Iterator() Iterator
##### 	// Return elements count
##### 	// Returns:
##### 	//    *big.Int Count of stream elements
##### 	Count() *big.Int
##### 	// Return elements average
##### 	// Returns:
##### 	//    *big.Float Average of stream elements
##### 	Average() *big.Float
}
    Describe Stream, dynamic approach to the Collections, within implemented
    dynamic filters, mapping (re-map types and values in the Collections), flat
    mapping (flatten arrays into the list as arrays content), reductions
    (reducing values), summarize methods

##### func ArrayAsParallelStream(arr RowArray) Stream
    Retrieve A Stream related to provided array, with parallelized computation
#####     Returns:
       types.Stream New Parallel Data Stream

##### func ArrayAsStream(arr RowArray) Stream
    Retrieve A Stream related to provided array
#####     Returns:
       types.Stream New Data Stream

##### func IteratorAsStream(iter Iterator) Stream
    Retrieve A Stream related to provided Iterator data
#####     Returns:
       types.Stream New Data Stream

