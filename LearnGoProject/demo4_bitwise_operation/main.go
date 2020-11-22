package main

import (
	"fmt"
	"strconv"
)

//原码 反码 补码   位运算概念
func main () {
	/*
	概念:
	1. 原码/反码/补码的展示形式基于数值的位数的0 1展示(非二进制), 比如int8 int16等对应的位数(按实际情况选择位数)
	   演示:一般使用8位来做展示, 第一位代表符号位(0为整数, 1为负数)
	1. 正数的原码/反码/补码是一样的
	2. 负数的反码=除符号位外所有的其他位取反, 负数的补码=反码+1
	3. 0的反码和补码也是一样的, 都是0
	4. 计算机运行的时候都是以补码的形式运行的, 用补码计算, 得到的结果也是补码, 还需转为原码;
	   如 1 - 1 = = 1 + (-1);
	   -1的反码 = 1111 1110
	   -1的补码 = 1111 1111
	   1 - 1 = 0000 0001 + 1111 1111 = 1000 0000  因为第一位是符号位, 所以得到了-0
	*/

	/*
	位运算符:
	1. 按位与 &  两位全为1, 则为1, 否则为0
	2. 按位或 | 两位只要有一个为1, 则为1, 否则为0
	3. 按位异或 ^ 两位一个为0一个为1, 则为1, 否则为0
	4. 左移运算符 << 符号位不变, 低位补0, 高位溢出 (从右边开始向左补)
		如-3 << 3 = (-3原码 1000 0011 反码1111 1100 补码"1111 1101") << 3 = 补码"1110 1000" 反码1110 0111 原码1001 1000 = -24
	5. 右移运算符 >> 符号位不变, 低位溢出, 并用符号位补溢出的高位(从符号位开始向右补)
		如-3 >> 3 = (-3原码 1000 0011 反码1111 1100 补码"1111 1101") >> 3 = 补码"1111 1111" 反码1111 1110 原码1000 0001 = -1
	*/

	//基于上面的概念, 来分析以下范例:
	num1 := 2 & 3
	//2 = 0000 0010
	//3 = 0000 0011
	//& = 0000 0010   = 2
	fmt.Printf("num1 := 2 & 3, num1 = %v\n", num1)

	num2 := 2 | 3
	//2 = 0000 0010
	//3 = 0000 0011
	//| = 0000 0011   = 3
	fmt.Printf("num2 := 2 | 3, num2 = %v\n", num2)

	num3 := 2 ^ -3  //负数的计算, 要转为补码, 得到的结果也是补码, 还需要转为原码
	//2  = 0000 0010
	//-3 = 1111 1101 (补码参与运算)  原码 1000 0011 -> 反码 1111 1100 ->  补码1111 1101
	//^  = 1111 1111 (补码) -> 反码 1111 1110 -> 原码 1000 0001   = -1
	fmt.Printf("num3 := 2 ^ -3, num3 = %v\n", num3)


	num4 := -2 << 3
	// -2 原码1000 0010 -> 反码 1111 1101 -> 补码 1111 1110
	// 1111 1110   -> 左移3位得到的补码 1111 0000  反码 1110 1111  原码 1001 0000 - -16
	fmt.Printf("num4 := -2 << 3, num4 = %v\n", num4)
	num5 := 2 << 3
	//0000 0010 -> 左移3位  0001 0000   = 16
	fmt.Printf("num5 := 2 << 3, num5 = %v\n", num5)
	num6 := -3 >> 3
	//-3原码 1000 0011 反码 1111 1100 补码 1111 1101
	//1111 1101 右移3位得到补码 1111 1111 反码 1111 1110  原码 1000 0001 = -1
	fmt.Printf("num6 := -3 >> 3, num6 = %v\n", num6)
	//fmt.Printf("%b  %v %v\n", 0b11111111 >> 3, 127 << 2, -1 << 2)
	num7 := 3 >> 3
	//0000 0011     0000 0000   = 0
	fmt.Printf("num5 := 2 >> 3, num7 = %v\n", num7)


	fmt.Printf("3 | 5 | 8 = %v\n", 3 | 5 | 8)
	fmt.Printf("15 | 3 | 5 = %v\n", 15 | 3 | 5)
	fmt.Printf("15 | 5 = %v\n", 15 | 5)
	fmt.Printf("15 | 15 = %v\n", 15 | 15)



	//按位与最佳实践----------
	//http://fingerchou.com/2018/07/15/a-best-practice-of-bitwise-operation/
	//问：假设 S由a, b, c三个数值组成, S的大小为2个字节(Byte)。已知S的值。求解a, b, c分别是多少？
	//其中a = 4 Bit, b = 4 Bit, c = 8 Bit。(a >= 0, b >= 0, c >= 0)

	//分析 两个字节, 为16位
	s := 10086
	a := s >> 12  //右移12位, 得到 a的4bit
	b := s >> 8 & 0b1111 //右移8位得到前8位, 与8位0b00001111, 再次舍弃前面4位
	c := s & 0b11111111 //16位s与16位0b00000000111111111, 舍弃前8位
	s = a << 12 + b << 8 + c
	fmt.Printf("a = %v, b = %v, c = %v\n", a, b, c)


	//获取一个数的高12位和第12位  go中中文3个字节, 也就是24位  (此处不区分CPU的大小端)
	//https://www.oschina.net/question/254766_129663
	str := "我"
	str_bytes :=  []byte(str)
	len_str_bytes := len(str_bytes)
	str_ := string((str_bytes))
	//此处注意点, str_bytes数据类型为[]byte. 这个是切片, 不是数组
	fmt.Printf("str_bytes为:%b, str_bytes类型:%T, len_str_bytes:%v, str_bytes还原为:%v\n", str_bytes, str_bytes, len_str_bytes, str_)
	//str_bytes 111001101000100010010001
	num_str := 0b111001101000100010010001  //得到str的二进制表示方式
	//num_str =  111001101000100010010001
	aa := num_str >> 12  //获取str的高12位, 向右移12位, 高位补0, 低位溢出. 原来的高位到了低位, 得到高12
	bb := num_str & 0b000000000000111111111111 //利用按位与的特性, 前12位与上0, 直接为0. 舍弃前12位, 得到低12
	fmt.Printf("aa = %v, bb = %v, aa二进制:%b, bb二进制:%b\n", aa, bb, aa, bb)
	num_str_new := aa << 12 + bb  //根据拆分的高低12位得到原来的二进制数

	str_num_binary := fmt.Sprintf("%b", num_str_new) //把二进制数转为二进制字符串
	str_bytes_new := [3]byte{}
	for i := 0; i < len(str_num_binary) / 8; i ++ {
		s := str_num_binary[i * 8 : i * 8 + 8]
		//num, _ := strconv.ParseInt(s, 2, 8)  //字符串转数字  此处转不能转8位, 有符号int8为-127 - 127, 而11111111为255. 所以要选>int8的类型
		num, _ := strconv.ParseUint(s, 2, 8) //无符号uint8 刚好够表示数值
		s_byte := byte(num)  //数字转byte
		fmt.Printf("s = %q, num = %q, s_byte = %q\n", s, num, s_byte)
		str_bytes_new[i] = s_byte
	}
	//注意点  [3]byte 和[]byte是两个不同的类型, 所有下面要把[3]byte转为[]byte
	// []byte转[3]byte, [3]byte和[]byte 一个是数组 一个是切片
	str_new := string(str_bytes_new[:])

	fmt.Printf("num_str = %b, num_str类型:%T, str_num_binary = %q\n", num_str, num_str, str_num_binary)
	fmt.Printf("str_bytes_new = %b, str_bytes_new类型:%T, str_bytes_new长度:%v,str_new:%v\n", str_bytes_new, str_bytes_new, len(str_bytes_new), str_new)


	//按位或最佳实践----------
	//类似oc的枚举多参数设置 topLeft | toRight | bottomLeft
	//参与或运算的数, 或上最終的值都等于最终值
	//a = 2; b = 3
	// 2 | 3  = 0010 | 0101 = 0111 = 7
	// 2 | 7 =  0010 | 0111 = 1111 = 7


	//按位异或最佳实践----------
	//交换两个值
	//a = 3; b = 5互换
	//3 ^ 5 = 0011 ^ 0101 =  0110 = 6
	//6 ^ 5 = 0110 ^ 0101 =  0011 = 3
	//6 ^ 3 = 0110 ^ 0011 = 0101 = 5
}