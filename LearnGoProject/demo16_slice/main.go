package main

import "fmt"

//slice 数组的进阶, slice为引用类型
func main() {
	//输出中文字符串 二进制(byte数组)

	//slice三种初始化
	//1.从数组初始化
	arr := [3]int {1, 5, 6}
	slice1 := arr[1:]
	fmt.Printf("arr = %v, arr type:%T, slice1 = %v, slice1 type:%T\n", arr, arr, slice1, slice1)
	fmt.Printf("arr[0] 地址:%p, slice1[0]地址:%p\n", &arr[0], &slice1[0])

	//var slice = arr[startIndex:endIndex] //从startIndex开始, 到endInex - 1
	//var slice = arr[:endIndex] //从头开始到endIndex -1
	//var slice = arr[startIndex:] //从startIndex到结束

	//slice为引用类型, 下面使用num_one指针指向slice[0]的地址.
	//通过num_one赋值, slice1[0]的值发生变化, slice1引用的arr中的值, 那么arr[0]也会变了
	var num_one *int = &slice1[0]
	*num_one = 11 //得到第一个元素的指针, 修改指向的值
	var slice1_ptr *[]int = &slice1
	(*slice1_ptr)[0] = 12 //得到slice1的指针, 通过slice1_ptr修改第一个元素的值
	fmt.Printf("arr = %v, arr type:%T, slice1 = %v, slice1 type:%T\n", arr, arr, slice1, slice1)

	//arr的地址指向第一个元素的地址, slice1的地址不是
	//arr地址:0xc00001a160, slice1地址:0xc00000c060
	fmt.Printf("arr地址:%p, slice1地址:%p\n", &arr, &slice1)

	//2.使用make 初始化, 默认值为对应的类型默认值
	//make参数 (1类型, 2长度, 3cap容量)
	//slice内部其实也是一个结构体.  分别有三个参数. arr用于存放数据. len长度. cap容量
	slice2 := make([]int, 3) // make([]int, 3, 3)   len必填. cap可不填, 默认和len一样
	slice3 := make([]int, 2, 3)
	slice3[0] = 11
	//slice3[1] = 12 //超过len. 无效
	//slice3[2] = 13 //炒股len, 无效
	//slice3[4] = 22 //
	fmt.Printf("slice2 = %v, slice3 = %v\n", slice2, slice3)

	//slice4 := make([]int, 10, 3) //len > cap. 编译不通过

	//3.直接指定具体数组, 和make原理一样
	slice4 := []int{7, 8, 9}
	fmt.Printf("slice3 = %v\n", slice4)


	//切片遍历
	for i := 0; i < len(slice4); i++ {
		fmt.Printf("i = %v, value = %v\n", i, slice4[i])
	}
	for i, value := range slice4 {
		fmt.Printf("for range: i = %v, value = %v\n", i, value)
	}


	//append 内置函数, 对切片进行动态追加
	//append 如原有的切片容量够用, 会重新切片容纳新元素, 否则就分配一个新的基本数组, 再重重新切片
	slice5 := append(slice3, 1)
	//slice6 := append(slice3, 1, 2)
	fmt.Printf("slice5 = %v\n", slice5)
	//slice5中追击的元素满足现有容量3, 两者对应的前两个元素地址一样
	//slice6中追加的元素超过了3, 重新分配一个新的数组进行切片. 两者前两个元素对应的地址就不在一样
	fmt.Printf("slice3[0]地址:%p, slice5[0]地址:%p\n", &slice3[0], &slice5[0])
	fmt.Printf("slice3[01]地址:%p, slice5[01]地址:%p\n", &slice3[01], &slice5[01])

	fmt.Printf("slice3 = %v, slice5 = %v\n", slice3, slice5)
	fmt.Printf("slice3地址:%p, slice5地址:%p,\n", &slice3, &slice5)

	//copy  从一个切片复制到另一个切片 copy(dest, src) 返回长度短的那个切片
	// 复制时, 不够的舍弃. 多余的也舍弃
	num_copy := copy(slice5, slice3)
	fmt.Printf("num_copy:%v, slice3 = %v, slice5 = %v\n", num_copy, slice3, slice5)
	fmt.Printf("slice3[0]地址:%p, slice5[0]地址:%p\n", &slice3[0], &slice5[0])
	fmt.Printf("slice3[01]地址:%p, slice5[01]地址:%p\n", &slice3[01], &slice5[01])

	slice7 := []int {5, 6, 7, 8}
	num_copy2 := copy(slice5, slice7)
	fmt.Printf("num_copy:%v, slice7 = %v, slice5 = %v\n", num_copy2, slice7, slice5)
	fmt.Printf("slice7[01]地址:%p, slice5[01]地址:%p\n", &slice7[01], &slice5[01])

	//fmt.Printf("slice6 = %v %T\n", slice9, slice9)
}