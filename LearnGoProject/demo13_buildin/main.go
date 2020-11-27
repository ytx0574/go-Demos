package main

import "fmt"

//go 内建函数使用
func main() {
	//new 内建函数   用来分配内存, 主要用来分配值类型  int float struct array  返回的是指针类型
	//make 内建函数  用来分配内存, 主要用来分配引用类型. channel map slice  后面再说
	//append  追加slice, 如果slice的空间够, 直接在现有的slice中追加. 如果空间不够, 直接重新make一片新的空间来装载数据
	//copy 复制字符串
	//delete  删除map的key


	//go new和make的区别

	num1 := 100
	num2 := new(int) //num2是指针类型
	fmt.Printf("num1 = %v, num1类型为%T, num2 = %v, num2的值为%v, num2类型为%T\n", num1, num1, num2,  *num2, num2)

	*num2 = 2
	fmt.Printf("num1 = %v, num1类型为%T, num2 = %v, num2的值为%v, num2类型为%T\n", num1, num1, num2,  *num2, num2)
}