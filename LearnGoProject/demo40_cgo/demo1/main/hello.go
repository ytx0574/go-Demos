package main

import "C"
import "fmt"

//export SayHello2
func SayHello2(s *C.char) {
	fmt.Println("go实现的C函数")
	fmt.Println(C.GoString(s))
}