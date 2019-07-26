package flows

import (
	"errors"
	"fmt"
	errs "github.com/hellgate75/general_utils/errors"
	"io"
	"os"
	"reflect"
)

//Flow Element provides Reading/Writing features for Flow Elements
type Flow interface {
	//Import ReadWriter interface
	io.ReadWriter
	//Import Closer interface
	io.Closer
	//Import Seeker interface
	io.Seeker
	//Import ReadFrom interface
	io.ReaderFrom
	//Import WriterTo interface
	io.WriterTo
	// Open a file in READ/WRITE and create a new file in case
	// Parameters:
	//    name (string) File name
	// Returns:
	//    error Any error that occurs during computation
	Open(name string) error
	// Reset the Selector and the Flow, if a file is open it will be closed and reponened
	// Returns:
	//    error Any error that occurs during computation
	Reset() error
	// Clear the Selector and the Flow, any reference will be closed
	// Returns:
	//    error Any error that occurs during computation
	Clear() error
	// Return the selector name or the empty string in case there is not selector
	// Returns:
	//    string  Selector name
	Selector() string
	// Return the byte array in the internal buffer
	// Returns:
	//    []byte  Byte array in the buffer
	Bytes() []byte
	// Return the atring made of byte array in the internal buffer
	// Returns:
	//    string  String made of the byte array in the buffer
	String() string
	// Return Size of the file or internal buffer
	// Returns:
	//    int64 Size of the file or the internal buffer
	Size() int64
}

type __flowStruct struct {
	buffer     []byte
	pos        int
	file       *os.File
	bufferSize int
}

func (flow *__flowStruct) String() string {
	return fmt.Sprintf("%s", flow.buffer)
}

func (flow *__flowStruct) Bytes() []byte {
	return flow.buffer
}

func (flow *__flowStruct) Clear() error {
	flow.file = nil
	flow.pos = 0
	flow.buffer = make([]byte, 0)
	return nil
}

func (flow *__flowStruct) Reset() error {
	if flow.file != nil {
		name := flow.Selector()
		err := flow.Close()
		flow.pos = 0
		flow.buffer = make([]byte, 0)
		if err != nil {
			return err
		}
		return flow.Open(name)
	}
	return flow.Clear()
}

func (flow *__flowStruct) Selector() string {
	defer func() {
		itf := recover()
		if itf != nil {
			var err error = nil
			if errs.IsError(itf) {
				errN := itf.(error)
				err = errors.New(fmt.Sprintf("Error during Flow::Close(), error is <%s>", errN.Error()))
			} else {
				err = errors.New(fmt.Sprintf("Error during Flow::Close(), error is <%v>", itf))
			}
			if logger != nil {
				logger.Error(err)
			}
		}
	}()
	if flow.file != nil {
		return flow.file.Name()
	}
	return ""
}

func (flow *__flowStruct) Open(name string) error {
	var err error = nil
	defer func() {
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				errN := itf.(error)
				err = errors.New(fmt.Sprintf("Error during Flow::Close(), error is <%s>", errN.Error()))
			} else {
				err = errors.New(fmt.Sprintf("Error during Flow::Close(), error is <%v>", itf))
			}
			if logger != nil {
				logger.Error(err)
			}
		}
	}()
	if _, err := os.Stat(name); err != nil {
		err = nil
		flow.file, err = os.Create(name)
		if err != nil {
			return err
		}
	} else {
		flow.file, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	}
	return err
}

func (flow *__flowStruct) Seek(offset int64, whence int) (int64, error) {
	var out int64 = int64(0)
	var err error = nil
	defer func() {
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				errN := itf.(error)
				err = errors.New(fmt.Sprintf("Error during Flow::Seek(), error is <%s>", errN.Error()))
			} else {
				err = errors.New(fmt.Sprintf("Error during Flow::Seek(), error is <%v>", itf))
			}
			if logger != nil {
				logger.Error(err)
			}
		}
	}()
	if flow.file != nil {
		out, err = flow.file.Seek(offset, whence)
	} else {
		origPos := flow.pos
		if whence == 0 {
			flow.pos = int(offset)
		} else if whence == 1 {
			flow.pos += int(offset)

		} else {
			flow.pos = len(flow.buffer) - int(offset)
		}

		if flow.pos < 0 {
			flow.pos = origPos
			out = int64(0)
			err = errors.New("Worng seek, going before beginning of the file")
		} else if flow.pos >= len(flow.buffer) {
			flow.pos = origPos
			out = int64(0)
			err = errors.New("Worng seek, going over to the end of the file")
		} else {
			out = int64(origPos - flow.pos)
			if out < 0 {
				out = 0 - out
			}
			err = nil
		}
	}
	return out, err
}

func (flow *__flowStruct) Close() error {
	var err error = nil
	defer func() {
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				errN := itf.(error)
				err = errors.New(fmt.Sprintf("Error during Flow::Close(), error is <%s>", errN.Error()))
			} else {
				err = errors.New(fmt.Sprintf("Error during Flow::Close(), error is <%v>", itf))
			}
			if logger != nil {
				logger.Error(err)
			}
		}
	}()
	if flow.file != nil {
		err = flow.file.Close()
	} else {
		err = errors.New(fmt.Sprintf("File not open!!"))
	}
	return err
}

func (flow *__flowStruct) Write(p []byte) (int, error) {
	var out int = int(0)
	var err error = nil
	defer func() {
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				errN := itf.(error)
				err = errors.New(fmt.Sprintf("Error during Flow::Write(), error is <%s>", errN.Error()))
			} else {
				err = errors.New(fmt.Sprintf("Error during Flow::Write(), error is <%v>", itf))
			}
			if logger != nil {
				logger.Error(err)
			}
		}
	}()
	if flow.file != nil {
		n, errN := flow.file.Write(p)
		out = n
		err = errN
		if out > 0 && err == nil {
			err = flow.file.Sync()
		}
	} else {
		var cc int = 0
		for _, b := range p {
			flow.buffer = append(flow.buffer, b)
			cc++
		}
		out = cc
		if cc < len(p) {
			err = errors.New(fmt.Sprintf("Buffer not written completely -> %d < %d", cc, len(p)))

		}
	}
	return out, err

}

func (flow *__flowStruct) Read(p []byte) (int, error) {
	var out int = int(0)
	var err error = nil
	defer func() {
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				errN := itf.(error)
				err = errors.New(fmt.Sprintf("Error during Flow::Read(), error is <%s>", errN.Error()))
			} else {
				err = errors.New(fmt.Sprintf("Error during Flow::Read(), error is <%v>", itf))
			}
			if logger != nil {
				logger.Error(err)
			} else {
				fmt.Println(err.Error())
			}
		}
	}()
	if flow.file != nil {
		n, errN := flow.file.Read(p)
		out = n
		err = errN
	} else {
		length1 := len(p)
		length2 := len(flow.buffer)
		length := length1 + flow.pos
		if length > length2-flow.pos {
			length = length2 - flow.pos
		}
		if length <= 0 {
			return 0, errors.New(fmt.Sprintf("Buffer not available -> length %d ", length))
		}
		var leg int = flow.pos
		if length == 0 {
			//			fmt.Println("Nothing read!!")
			return 0, errors.New("Nothing read!!")
		} else if length > flow.pos+len(p) {
			//			fmt.Println("Resize to buffer size!!")
			length = flow.pos + len(p)
		} else if length < flow.pos+len(p) {
			//			fmt.Println("Cutting buffer!!")
			length = flow.pos + len(p)
		}
		var val []byte = make([]byte, length-flow.pos)
		for i := flow.pos; i < length; i++ {
			if i >= len(flow.buffer) {
				break
			}
			val[i-flow.pos] = flow.buffer[i]
			leg++
		}
		tVal := reflect.ValueOf(val)
		x := reflect.ValueOf(&p)
		reflect.Copy(x.Elem(), tVal)

		flow.pos = leg
		out = length
		if out == 0 {
			err = errors.New(fmt.Sprintf("Buffer not read completely -> %d == %d", out, 0))

		}
	}
	return out, err
}

func (flow *__flowStruct) ReadFrom(r io.Reader) (int64, error) {
	var out int64 = int64(0)
	var err error = nil
	defer func() {
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				errN := itf.(error)
				err = errors.New(fmt.Sprintf("Error during Flow::ReadFrom(), error is <%s>", errN.Error()))
			} else {
				err = errors.New(fmt.Sprintf("Error during Flow::ReadFrom(), error is <%v>", itf))
			}
			if logger != nil {
				logger.Error(err)
			}
		}
	}()
	var bytes []byte = make([]byte, flow.bufferSize)
	n, err := r.Read(bytes)
	for n > 0 && err != nil {
		fmt.Println(n)
		for i := 0; i < n; i++ {
			flow.buffer = append(flow.buffer, bytes[i])
		}
		bytes = make([]byte, flow.bufferSize)
		n, _ = r.Read(bytes)
	}
	if err == nil {
		out += int64(n)
	}
	return out, err
}

func (flow *__flowStruct) WriteTo(w io.Writer) (int64, error) {
	var out int64 = int64(0)
	var n int = 0
	var err error = nil
	defer func() {
		itf := recover()
		if itf != nil {
			if errs.IsError(itf) {
				errN := itf.(error)
				err = errors.New(fmt.Sprintf("Error during Flow::WriteTo(), error is <%s>", errN.Error()))
			} else {
				err = errors.New(fmt.Sprintf("Error during Flow::WriteTo(), error is <%v>", itf))
			}
			if logger != nil {
				logger.Error(err)
			}
		}
	}()
	n, err = w.Write(flow.buffer)
	out = int64(n)
	return out, err
}

func (flow *__flowStruct) Size() int64 {
	if flow.file != nil {
		info, err := flow.file.Stat()
		if err != nil {
			return int64(0)
		}
		return info.Size()
	} else {
		return int64(len(flow.buffer))
	}
}

// Creates and Opens READ/WRITE a file based flows.Flow
// Parameters:
//    file (string) File name
//    buffer (int) read buffer size
// Returns:
//    (flows.Flow Flow component instance,
//    error Any error that occurs during computation)
func NewFileFlow(file string, buffer int) (Flow, error) {
	strct := &__flowStruct{
		buffer:     make([]byte, 0),
		file:       nil,
		pos:        0,
		bufferSize: buffer,
	}
	err := strct.Open(file)
	return strct, err
}

// Creates and Opens READ/WRITE a bytes.ByteArray based flows.Flow
// Parameters:
//    bytes (bytes []byte) Initial buffer content
//    buffer (int) read buffer size
// Returns:
//    flows.Flow Flow component instance
func NewByteArrayFlow(bytes []byte, buffer int) Flow {
	return &__flowStruct{
		buffer:     bytes,
		file:       nil,
		pos:        0,
		bufferSize: buffer,
	}
}

// Creates and Opens READ/WRITE a new empty flows.Flow with 64k buffer
// Returns:
//    flows.Flow Flow component instance
func NewFlow() Flow {
	return &__flowStruct{
		buffer:     make([]byte, 0),
		file:       nil,
		pos:        0,
		bufferSize: 65536,
	}
}

// Creates and Opens READ/WRITE a new empty flows.Flow with specified buffer
// Parameters:
//    buffer (int) read buffer size
// Returns:
//    flows.Flow Flow component instance
func NewFlowBuff(buffer int) Flow {
	return &__flowStruct{
		buffer:     make([]byte, 0),
		file:       nil,
		pos:        0,
		bufferSize: buffer,
	}
}
