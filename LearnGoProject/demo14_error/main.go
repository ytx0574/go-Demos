package main

import (
	"errors"
	"fmt"
	"runtime/debug"
	"time"
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
	}

	value = 1 / num
	fmt.Printf("value = %v\n", value)
	return
}

type INT int

func (this INT)Error() string {
return ""
}

func main() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("获取拦截到崩溃, 无限重启我自己 %v\n", err)
			time.Sleep(time.Second * 3)
			main()
		}
	}()

	fmt.Printf("test执行之前\n")
	test(1)
	fmt.Printf("test执行之前\n")


	//todo go中nil的特殊使用  有时会有条件 (nil != nil) 成立;
	//todo 此处虽然in指向了aa, aa的值为nil, 但是这是in并不是真的nil, 它指向的类型为*int, 值为nil
	var aa *int
	var in interface{}

	in = aa

	if in != nil {
		fmt.Printf("注意:  此时in指向的值虽然是nil, 但指向对象类型为*int, 所以它不是真的nil  \"in = %+v, aa = %s\"\n", in, aa)
	}

	var INT1 *INT
	var in1 error

	in1 = INT1

	if in1 != nil {
		fmt.Printf("此处一样不等于nil\n")
	}

	t := TEST(Test)
	t.Test("1", 1)
}

type TEST func(a string, b int)

func (t TEST)Test(a string, b int) {
	fmt.Printf("TEST Type func \n")
	t(a, b)
}

func Test(a string, b int) {
	fmt.Printf("Test func \n")
	panic("我崩溃了....")
}