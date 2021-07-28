package main

import (
	"fmt"
)

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
	//slice3[2] = 13 //炒股len, 无效[
	//slice3[4] = 22 //
	fmt.Printf("slice2 = %v, slice3 = %v\n", slice2, slice3)

	//slice4 := make([]int, 10, 3) //len > cap. 编译不通过

	//3.直接指定具体数组, 和make原理一样
	slice4 := []int{7, 8, 9}
	fmt.Printf("slice3 = %v\n", slice4)

	//slice取值 取值从左边的值的坐标开始, 到右边的值的坐标减一.
	//比如下面的, 左右两者都可以为3. 不可超过数组长度, 否则运行时错误, 下标越界
	//:左边的值不可小于右边, 否则编译失败
	var slice = []int{1, 2, 3}
	slice_new := slice[3:3] //  slice[3:3] = slice[3:] = slice[:3]
	fmt.Println("slice_new = ", slice_new)

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

	//修改字符串内容  因为string是值类型, 所以转为byte切片, 再进行修改 也可以用到前面使用到的string包里面的函数操作字符串
	str := "zhe shi yi zhi zhu!"
	fmt.Printf("%c, %T\n", str[3], str[3])
	slice8 := []byte(str) // str转byte数组
	slice8[0] = 'Z'
	str1 := string(slice8)
	str2 := str[:]  //str切片, 返回值还是string
	fmt.Printf("slice8 = %v\nstr1 = %v\nstr2 = %v\n", slice8, str1, str2)
	fmt.Printf("slice8 type:%T, str2 type:%T\n", slice8, str2)

	//int32 == rune  3个字节, 可表示中文
	str3 := string([]int32{1, 2, 3, 33333, 44334, 22333})

	//常规的char. byte == uint8. 表示单个字符
	str4 := string([]uint8{100, 77, 180})
	fmt.Printf("str3 = %v str4 = %v\n", str3, str4)

	//使用切片获取斐波那契数列
	slice9 := fbn(11)
	fmt.Printf("斐波那契数列为: %v\n", slice9)

	slice10 := append([]string{}, "11")
	fmt.Printf("slice10:%v\n", slice10)
	var slice11 []string

	slice12 := append(slice11, "123")
	fmt.Printf("slice12:%v\n", slice12)

	//切片在未初始化时, 它的值是nil. 初始化后才地址  此处使用%v, 两次打印都是空数组, 看不出区别
	fmt.Printf("slice11:%v slice11地址:%p slice11的值的地址:%p\n", slice11, &slice11, slice11)
	//slice11 = []string{"11"}
	slice11 = make([]string, 0)
	fmt.Printf("slice11:%v slice11地址:%p slice11的值的地址:%p\n", slice11, &slice11, slice11)


	str5 := "abc 1"
	str6 := "abc 1c"
	str7 := "abc 1我"

	str5_bytes := []byte(str5)
	str6_bytes := []byte(str6)
	str7_bytes := []uint8(str7)
	str7_bytes2 := []int32(str7)

	//关于字符在二进制的编码显示  参考链接:https://www.cnblogs.com/yaowen/p/10572455.html
	//[1100001 1100010 1100011 100000 110001]
	//[1100001 1100010 1100011 100000 110001 1100011]                       1100011 = 99 = c
	//[1100001 1100010 1100011 100000 110001 11100110 10001000 10010001]    后面的三个字符为 = 我
	//										 111001101000100010010001       = 我 (二进制表示)
	//										 1110xxxx 10xxxxxx 10xxxxxx     = UTF-8 标记中文3个字节的编码模板
	//										 11100110 10001000 10010001     = 使用 我 的ASCII码110001000010001, 从右向左进行套用模板得到 我 的二进制表示
	//										 0110 001000 010001             = 我 的ASCII码拆分
	//[1100001 1100010 1100011 100000 110001 110001000010001]               110001000010001 = 0x6211 =  25105 = 我 (ASCII)

	//UTF8编码模板如下
	//
	//1字节 0xxxxxxx
	//2字节 110xxxxx 10xxxxxx
	//3字节 1110xxxx 10xxxxxx 10xxxxxx
	//4字节 11110xxx 10xxxxxx 10xxxxxx 10xxxxxx
	//5字节 111110xx 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx
	//6字节 1111110x 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx 10xxxxxx

	//此处内部已经规定好, 常规的字符用byte(uint8)表示 中文因为占3个字节, 所以用rune(int32)表示 其他的类型都不行
	//带中文的字符串, 常规的字符还是用byte表示. 中文用3个字节表示
	//int32表示字符串时, 汉字转为对应的unicode码
	//
	fmt.Printf("%q\n%q\n%q\n%q\n", str5_bytes, str6_bytes, str7_bytes, str7_bytes2)
	fmt.Printf("%b\n%b\n%b\n%b\n", str5_bytes, str6_bytes, str7_bytes, str7_bytes2)

	//11100101 10011011 10111101  国 byte 转二进制
	//101011011111101   国 int32(ASCII码) 转二进制
	//0101 011011 111101 使用下面模板套用, 生成二进制, 高位不足补0
	//1110xxxx 10xxxxxx 10xxxxxx   UTF8中文模板
	//11100101 10011011 10111101   套用上面的模板, 使用 国 的 ASCII生成的二进制套用 得到 国 的二进制表示
	fmt.Printf("国 bytes binary:%b, 国 rune binary:%b\n", []byte("国"), []rune("国"))



	fruit := make([]string, 5, 10)
	fruit[0] = "11"
	fruit[1] = "22"
	fruit[2] = "33"
	fruit[3] = "44"
	fruit[4] = "55"
	myFruit := fruit[1:3:8]  //第三个参数代表切片的容量. 此处容量为8-1= 7, 不可大于原始切片或数组的容量 (比如设置为11, 就会panic)
	fmt.Printf("myFriut:  %v\n", myFruit)
	fmt.Printf("myFruit len:%v, cap:%v\n", len(myFruit), cap(myFruit))


	//todo 追加 直接append
	myFruit = append(myFruit, "77")
	fmt.Printf("myFruit len:%v, cap:%v\n", len(myFruit), cap(myFruit))

	//todo 插入头部 直接append
	myFruit = append([]string{"1", "2", "3"}, myFruit...)
	fmt.Printf("头部插入: myFruit:  %v\n", myFruit)
	fmt.Printf("myFruit len:%v, cap:%v\n", len(myFruit), cap(myFruit))

	//todo 插入中间
	var ii = "666"
	var index = 1
	myFruit = append(myFruit, "")
	copy(myFruit[index + 1:], myFruit[index:])
	myFruit[index] = ii
	fmt.Printf("index:%v插入%v myFruit:  %v\n", index, ii, myFruit)
	fmt.Printf("myFruit len:%v, cap:%v\n", len(myFruit), cap(myFruit))

	//todo 插入切片
	index = 5
	var iiiSlice = []string{"777", "888"}
	myFruit = append(myFruit, iiiSlice...)  //先扩展切片的容量个数
	copy(myFruit[index + len(iiiSlice):], myFruit[index:]) //把index后面的数据移动到index+len(切片)的位置
	//copy(myFruit[index : index + len(iiiSlice)], iiiSlice) //把插入的切片放入index的位置
	copy(myFruit[index :], iiiSlice)
	fmt.Printf("index:%v插入%v myFruit:  %v\n", index, iiiSlice, myFruit)
	fmt.Printf("myFruit len:%v, cap:%v\n", len(myFruit), cap(myFruit))

	//todo 删除元素  直接生成切片即可
	//todo 删除中间元素
	myFruit = append(myFruit[:index], myFruit[index + 2:]...)
	fmt.Printf("index:%v删除两个元素 myFruit:  %v\n", index, myFruit)
	fmt.Printf("myFruit len:%v, cap:%v\n", len(myFruit), cap(myFruit))

	index = 1
	//todo 删除index = 1, 后面的两个元素
	//todo 使用copy把要删除的片段用后面的数据覆盖, 然后再切片抛弃最后的删除切片的长度
	myFruit = myFruit[: index + copy(myFruit[index:], myFruit[index + 2:])]
	fmt.Printf("index:%v删除两个元素 myFruit:  %v\n", index, myFruit)
	fmt.Printf("myFruit len:%v, cap:%v\n", len(myFruit), cap(myFruit))


	//todo 利用切片[:0]可以很轻易的实现make(xx, 0, 11)的效果
	//todo 在已知长度的情况下, 避免对slice进行扩容, 可以使用[:0]来实现和原始数据容量一样的切片, 进而避免在处理过程中扩容
	//todo 比如下面我要筛选出MyFruit小于33的数据, 那么新得到的切片肯定小于原始数据的切片
	NewFruit := myFruit[:0]
	fmt.Printf("len:%v, cap:%v\n", len(NewFruit), cap(NewFruit))
	fmt.Printf("myFruit len:%v, cap:%v\n", len(myFruit), cap(myFruit))

}

func fbn (n int) []int {
	slice := make([]int, n)
	if n >= 1 {
		slice[0] = 1
	}
	if n >= 2 {
		slice[1] = 1
	}
	for i := 2; i < n; i++ {
		slice[i] = slice[i - 1] + slice[i - 2]
	}
	return slice
}