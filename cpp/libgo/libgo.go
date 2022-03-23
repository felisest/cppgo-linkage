package main

import "C"

import (
	"fmt"
)

//export PrintString
func PrintString(str string) {

	fmt.Println(str)
}

//export ReverseSlice
func ReverseSlice(buffIn []byte) {

	r := len(buffIn) - 1
	for i := range buffIn {
		t := buffIn[r]
		buffIn[r] = buffIn[i]
		buffIn[i] = t

		r--
		if r-i <= 1 {
			break
		}
	}
}

func main() {}
