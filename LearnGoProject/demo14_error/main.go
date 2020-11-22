package main

import (
	"errors"
	"fmt"
	"runtime/debug"
)

//go 错误处理  go不支持传统的 try...catch...finally
//go的处理方式为. defer recover 用于捕获异常. panic用于抛出异常

func test(num int) (value int) {
	//哪个函数有异常, 就写在哪个里面.  如果把这个捕获写到main, 则无效. 因为main并没有异常.
	defer func() {
		err := recover() //捕获错误
		if err != nil {
			debug.PrintStack() //打印调用栈信息
			//fmt.Printf("%s\n", runtime.ReadTrace())
			fmt.Printf("捕获异常信息:\"%v\", 继续执行后面代码\n", err)
		}
	}()

	if num == 0 {
		 err := errors.New("除数不能为0")
		 panic(err)  //当除数为0时, 主动抛出异常.  即使这里不处理, 系统遇到除数为0也会抛出异常
		 panic("fdsf: ss")  //panic可返回任意类型
	}

	value = 1 / num
	fmt.Printf("value = %v\n", value)
	return
}

func main() {

	fmt.Printf("test执行之前\n")
	test(0)
	fmt.Printf("test执行之前\n")
}