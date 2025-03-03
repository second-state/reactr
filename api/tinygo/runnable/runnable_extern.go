//go:build tinygo.wasm

package runnable

// #include <reactr.h>
import "C"
import (
	"reflect"
	"unsafe"
)

var runnable_ Runnable
var ident_ int32

func ident() int32 {
	return ident_
}

//export run_e
func run_e(rawdata uintptr, size int32, ident int32) {
	var input []byte
	ident_ = ident

	inputHeader := (*reflect.SliceHeader)(unsafe.Pointer(&input))
	inputHeader.Data = rawdata
	inputHeader.Len = uintptr(size)
	inputHeader.Cap = uintptr(size)

	result, err := runnable_.Run(input)

	if err != nil {
		returnError(err, ident)
		return
	}

	resPtr, resLen := rawSlicePointer(result)

	C.return_result(resPtr, resLen, ident)
}

func returnError(err error, ident int32) {
	code := int32(500)

	if err == nil {
		C.return_error(code, unsafe.Pointer(uintptr(0)), 0, ident)
		return
	}

	switch e := err.(type) {
	case RunErr:
		code = int32(e.Code)
	}

	errPtr, errLen := rawSlicePointer([]byte(err.Error()))

	C.return_error(code, errPtr, errLen, ident)
}
