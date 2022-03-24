package main

/*
#cgo CFLAGS: -I./lib
#cgo LDFLAGS: -L./lib/build_release -lcpp_go_shared -Wl,-rpath=./lib/build_release

#include "lib/simple_lib.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {

	c_str := C.CString("Hello from C library!")
	C.print_hello(c_str)
	C.free(unsafe.Pointer(c_str))

	buff := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}
	c_buff := C.CBytes(buff)
	C.reverse_vector(c_buff, C.int(len(buff)))
	fmt.Println(C.GoStringN((*C.char)(c_buff), C.int(len(buff)+1)))
	C.free(unsafe.Pointer(c_buff))
}
