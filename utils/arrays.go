package utils

import (
	"fmt"
	"github.com/hellgate75/general_utils/common"
)

//Generic Array Navigator structure
type NavAttr struct {
	arr_sel []common.Type
	index   int
}

//Integer Array Navigator structure
type IntNavAttr struct {
	arr_sel []int
	index   int
}

//Float Array Navigator structure
type FloatNavAttr struct {
	arr_sel []float64
	index   int
}

//Boolean Array Navigator structure
type BoolNavAttr struct {
	arr_sel []bool
	index   int
}

//Move next Element in the Array
// Returns:
//    bool Next command success state
func (nav *NavAttr) Next() bool {
	flag := nav.index < len(nav.arr_sel)-1
	if flag {
		nav.index++
	}
	return flag
}

//Move next Element in the Array
// Returns:
//    bool Next command success state
func (nav *IntNavAttr) Next() bool {
	flag := nav.index < len(nav.arr_sel)-1
	if flag {
		nav.index++
	}
	return flag
}

//Move next Element in the Array
// Returns:
//    bool Next command success state
func (nav *FloatNavAttr) Next() bool {
	flag := nav.index < len(nav.arr_sel)-1
	if flag {
		nav.index++
	}
	return flag
}

//Move next Element in the Array
// Returns:
//    bool Next command success state
func (nav *BoolNavAttr) Next() bool {
	flag := nav.index < len(nav.arr_sel)-1
	if flag {
		nav.index++
	}
	return flag
}

//Move previous Element in the Array
// Returns:
//    bool Prev command success state
func (nav *NavAttr) Prev() bool {
	flag := nav.index >= 1
	if flag {
		nav.index--
	}
	return flag
}

//Move previous Element in the Array
// Returns:
//    bool Prev command success state
func (nav *IntNavAttr) Prev() bool {
	flag := nav.index >= 1
	if flag {
		nav.index--
	}
	return flag
}

//Move previous Element in the Array
// Returns:
//    bool Prev command success state
func (nav *FloatNavAttr) Prev() bool {
	flag := nav.index >= 1
	if flag {
		nav.index--
	}
	return flag
}

//Move previous Element in the Array
// Returns:
//    bool Prev command success state
func (nav *BoolNavAttr) Prev() bool {
	flag := nav.index >= 1
	if flag {
		nav.index--
	}
	return flag
}

//Get current Element in the Array
// Returns:
//    common.Type Current Element or nil in case of error
func (nav *NavAttr) Get() common.Type {
	if nav.index >= 0 {
		return nav.arr_sel[nav.index]
	}
	return nil
}

//Get current Element in the Array
// Returns:
//    int Current Element or 0 in case of error
func (nav *IntNavAttr) Get() int {
	if nav.index >= 0 {
		return nav.arr_sel[nav.index]
	}
	return 0
}

//Get current Element in the Array
// Returns:
//    float64 Current Element or 0.0 in case of error
func (nav *FloatNavAttr) Get() float64 {
	if nav.index >= 0 {
		return nav.arr_sel[nav.index]
	}
	return 0.0
}

//Get current Element in the Array
// Returns:
//    bool Current Element or false in case of error
func (nav *BoolNavAttr) Get() bool {
	if nav.index >= 0 {
		return nav.arr_sel[nav.index]
	}
	return false
}

//Print current Element in the Array
func (nav *NavAttr) Print() {
	fmt.Println("NavAttr")
}

//Get current Array length
// Returns:
//    int Length of the Array
func (nav *NavAttr) Len() int {
	return len(nav.arr_sel)
}

//Get current Array length
// Returns:
//    int Length of the Array
func (nav *IntNavAttr) Len() int {
	return len(nav.arr_sel)
}

//Get current Array length
// Returns:
//    int Length of the Array
func (nav *FloatNavAttr) Len() int {
	return len(nav.arr_sel)
}

//Get current Array length
// Returns:
//    int Length of the Array
func (nav *BoolNavAttr) Len() int {
	return len(nav.arr_sel)
}

//Get current position in the Array
// Returns:
//    int Position of cursor in the Array
func (nav *NavAttr) Position() int {
	return nav.index
}

//Get current position in the Array
// Returns:
//    int Position of cursor in the Array
func (nav *IntNavAttr) Position() int {
	return nav.index
}

//Get current position in the Array
// Returns:
//    int Position of cursor in the Array
func (nav *FloatNavAttr) Position() int {
	return nav.index
}

//Get current position in the Array
// Returns:
//    int Position of cursor in the Array
func (nav *BoolNavAttr) Position() int {
	return nav.index
}

// Base Array Navigator Interface
type BaseArrayNav interface {
	Prev() bool
	Next() bool
	Len() int
	Position() int
}

// Generic Type Array Navigator Interface
type ArrayNav interface {
	BaseArrayNav
	Get() common.Type
}

// Integer Array Navigator Interface
type IntArrayNav interface {
	BaseArrayNav
	Get() int
}

// Float Array Navigator Interface
type FloatArrayNav interface {
	BaseArrayNav
	Get() float64
}

// Boolean Array Navigator Interface
type BoolArrayNav interface {
	BaseArrayNav
	Get() bool
}

// Create New Generic Type Array Navigator
// Parameters:
//   arr ([]common.Type) input Array to manage
// Returns:
//   ArrayNav Array Navigator feature for the specified type
func NewArrayNav(arr []common.Type) ArrayNav {
	return &NavAttr{
		arr_sel: arr,
		index:   -1,
	}
}

// Create New Integer Array Navigator
// Parameters:
//   arr ([]int) input Array to manage
// Returns:
//   IntArrayNav Array Navigator feature for the specified type
func NewIntArrayNav(arr []int) IntArrayNav {
	return &IntNavAttr{
		arr_sel: arr,
		index:   -1,
	}
}

// Create New Float Array Navigator
// Parameters:
//   arr ([]float64) input Array to manage
// Returns:
//   FloatArrayNav Array Navigator feature for the specified type
func NewFloatArrayNav(arr []float64) FloatArrayNav {
	return &FloatNavAttr{
		arr_sel: arr,
		index:   -1,
	}
}

// Create New Boolean Array Navigator
// Parameters:
//   arr ([]bool) input Array to manage
// Returns:
//   BoolArrayNav Array Navigator feature for the specified type
func NewBoolArrayNav(arr []bool) BoolArrayNav {
	return &BoolNavAttr{
		arr_sel: arr,
		index:   -1,
	}
}
