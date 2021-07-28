package main

/*
	#cgo CFLAGS: -I.
	#cgo LDFLAGS: -L. -ltest
	#include "test.h"
*/
import "C"
import (
	"fmt"
)

func main() {

	//todo 需确保lib和可执行在同一个工作目录
	//todo 需编译再运行 go build -o test && ./test
	fmt.Println(C.sum(1, 100))
}