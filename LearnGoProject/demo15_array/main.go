package main

import (
	"fmt"
	"math/rand"
	"time"
)

// go 数组简单使用, 数组是值类型, 声明指定了长度之后, 使用时不可超过这个数量
// 数组元素只能位相同类型, 声明后, 默认元素值 值类型为0 bool类型为false 字符串为""
// 数组声明一旦指定长度, 那么声明和赋值长度必须一致 参考下面...语法
// 数组的长度也参与了数组定义, [1]int 和[2]int 是两个不同的数组类型
func main() {
	var arr [10]int = [10]int{13}
	arr[01] = 11
	arr[02] = 12
	arr[07] = 13
	//arr[09] = 11 //不可像类似oc的写法. 此处0开头表示八进制 9不在八进制的显示范围
	//arr的首地址就是第一个元素的地址, 此处int是64位8个字节.
	//arr地址:0xc0000b6000, arr[0]地址:0xc0000b6000 + 8 -> arr[1]地址:0xc0000b6008 + 8-> arr[1]地址:0xc0000b6010
	//arr[0]和arr[1]的地址相差8个字节
	fmt.Printf("arr = %v  arr地址:%p, arr[0]地址:%p, arr[1]地址:%p, arr[1]地址:%p\n", arr, &arr, &arr[0], &arr[1], &arr[2])

	arr1 := [3]int {0:1, 1:2, 2:3}
	//arr1 = arr //此处错误, arr1已指定大小, 和arr长度不匹配
	arr2 := [3]int{}
	arr2 = arr1 //此处可行, 两者都是int类型数组 长度为3
	fmt.Printf("arr1 = %v\n", arr1)
	fmt.Printf("arr2 = %v\n", arr2)

	//数组遍历
	for i, value := range arr2 {
		fmt.Printf("arr2[%v] = %v\n", i, value)
	}

	//数组声明
	//...只能用在后面,  简写手动指定长度. 一旦赋值后, 长度也会固定. 使用后前面如果声明了长度, 那么必须和后面赋值长度一致
	var arr3 [6] int = [...]int{1, 2, 3, 4, 6, 8}
	//前面声明不指定长度, 依靠后面声明来指定长度
	arr4 := [...]int{1, 2, 3, 4, 6, 8}
	//前面声明指定长度, 后面指定的长度也要保持一致
	var arr5 [6] int = [6]int{1, 2, 3, 4, 6, 8}
	//前面声明的长度和后面...声明的实际长度不一致, 编译不通过
	//var arr6 [3] int = [...]int{1, 2, 3, 4, 6, 8}
	arr3[0]  = 1
	arr3[1] = 11
	fmt.Printf("arr3 = %v, arr3 = %v, arr5 = %v\n", arr3, arr4, arr5)


	//数组练习题
	// 随机生成5个数
	// 得到随机数后反转打印

	 arrNum := [5]int{}
	 //rand.Seed(time.Now().UnixNano())
	 rand.Seed(int64(time.Now().Nanosecond()))
	 for i := range arrNum {
		 arrNum[i] = rand.Intn(100)
	 }
	 fmt.Printf("arrNum =%v\n", arrNum)

	 arrNumLen := len(arrNum)
	 cycle := arrNumLen / 2
	 if arrNumLen % 2 == 0 {
		 cycle++
	 }

	 for i := 0; i < cycle; i++ {
		 if i <= arrNumLen / 2 {
			 pairIndex := arrNumLen - 1 - i
			 //交换i和pairIndex的值
			 arrNum[i] = arrNum[i] ^ arrNum[pairIndex]
			 arrNum[pairIndex] = arrNum[i] ^ arrNum[pairIndex]
			 arrNum[i] = arrNum[pairIndex] ^ arrNum[i]

			 //按位异或分析
			 //a = 3; b = 5互换
			 //3 ^ 5 = 0011 ^ 0101 =  0110 = 6
			 //6 ^ 5 = 0110 ^ 0101 =  0011 = 3
			 //6 ^ 3 = 0110 ^ 0011 = 0101 = 5
		 }
	 }
	fmt.Printf("arrNum = %v\n", arrNum)

	arr6 := [][]int{{1}, {2}}
	for _, value := range arr6 {
		fmt.Printf("%v\n", value)
	}

	arr7 := [2][1][2]int{ {{1, 2}}, {{1, 2}}}  //三维数组
	fmt.Printf("arr6 type:%T, arr7 type:%T\n", arr6, arr7)
	fmt.Printf("arr7 %v\n", arr7)
}