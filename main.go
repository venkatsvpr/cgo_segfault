package main

// #cgo CFLAGS: -g -Wall
// #include <stdlib.h>
// #include "greeter.h"
import "C"
import (
	"fmt"
	"unsafe"
)

type Ctxt struct {
	Context uintptr
}

type traceback struct {
	Context    uintptr
	SigContext uintptr
	Buf        *uintptr
	Max        uintptr
}

type symbolizer struct {
	PC     uintptr // program counter to fetch information for
	File   *byte   // file name (NUL terminated)
	Lineno uintptr // line number
	Func   *byte   // function name (NUL terminated)
	Entry  uintptr // function entry point
	More   uintptr // set non-zero if more info for this PC
	Data   uintptr // unused by runtime, available for function
}

func init() {
	/*
		c := unsafe.Pointer(&Ctxt{})
		t := unsafe.Pointer(&traceback{})
		s := unsafe.Pointer(&symbolizer{})
		runtime.SetCgoTraceback(0, nil, nil, nil)
	*/
}

func main() {
	// panic(" some panic")
	name := C.CString("Gopher")
	defer C.free(unsafe.Pointer(name))

	year := C.int(2018)

	ptr := C.malloc(C.sizeof_char * 1024)
	defer C.free(unsafe.Pointer(ptr))

	size := C.greet(name, year, (*C.char)(ptr))

	b := C.GoBytes(ptr, size)
	fmt.Println(string(b))

	fmt.Println(" Completed normally")
}
