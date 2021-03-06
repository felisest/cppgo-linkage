package main

/*
#cgo CFLAGS: -I./lib
#cgo LDFLAGS: -L./lib/build_release -lcpp_go_static -Wl,-rpath=./lib/build_release -lstdc++

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

	//Buffer by reference
	buff2 := []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H'}
	C.reverse_vector(unsafe.Pointer(&buff2[0]), C.int(len(buff2)))
	fmt.Println(string(buff2))

	num := 7
	s := C.lsqrt(C.int32_t(num))
	fmt.Printf("sqrt of %d = %d\n", num, s)
}

