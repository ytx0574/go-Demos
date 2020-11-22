package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

//进制转换
func main () {
	//------其他进制转十进制
	//二进制转十进制   从右往左开始, 依次求每一位乘以的2的(位数-1)次方, 然后求和
	//var two_number int = 1111  //15

	var two_convert_ten = 1 * math.Pow(2, 0) + 1 * math.Pow(2,1) + 1 * math.Pow(2, 2) + math.Pow(2, 3)
	fmt.Printf("二进制转十进制 1111 -> %v\n", two_convert_ten)

	//八进制转十进制   从右往左开始, 依次获得每一位乘以8的(位数-1次方), 然后求和
	//var eight_number = 0712   //458
	var eight_convert_ten = 2 * math.Pow(8, 0) + 1 * math.Pow(8, 1) + 7 * math.Pow(8, 2) + 0 * math.Pow(8, 3)
	fmt.Printf("八进制转十进制  0712 -> %v\n", eight_convert_ten)

	//十六进制转十进制, 从右往左开始, 依次求每一位乘以16的(位数-1)次方, 然后求和
	//var sixteen int = 0x0D3C10  //867344
	var sixteen_convert_ten = 0 * math.Pow(16, 0) + 1 * math.Pow(16,1) + 12 * math.Pow(16,2) + 3 * math.Pow(16, 3) + 13 * math.Pow(16,4)
	fmt.Printf("十六进制转十进制 0x0D3C10 -> %v\n", sixteen_convert_ten)


	//-------十进制转其他进制
	//十进制转二进制  一直除以2, 然后取余. 知道无法除尽, 倒着取每一位的余数. 其它进制同理 ( 转2除2 转8除8 转16除16)
	var origin_number int = 61
	num := origin_number
	var str_ten_convert_two string = ""
	for ; num > 0; num /= 2 {
		lsb := num % 2
		str_ten_convert_two = strconv.Itoa(lsb) + str_ten_convert_two
	}
	fmt.Printf("十进制转二进制: %v -> %v\n", origin_number, str_ten_convert_two)

	num = origin_number
	var str_ten_convert_eight string = "0"
	for ; num > 0; num /= 8 {
		lsb := num % 8
		str_ten_convert_eight = str_ten_convert_eight + strconv.Itoa(lsb)
	}
	fmt.Printf("十进制转八进制: %v -> %v\n", origin_number, str_ten_convert_eight)

	num = origin_number
	var str_ten_convert_sixteen string = "0x"
	for ; num > 0; num /= 16 {
		lsb := num % 16
		str_ten_convert_sixteen = str_ten_convert_sixteen + fmt.Sprintf("%x", lsb)
	}
	fmt.Printf("十进制转十六进制: %v -> %v\n", origin_number, str_ten_convert_sixteen)


	//------ 八进制 十六进制转二进制
	//八进制的每一位转为3位数的二进制即可   不够3位二进制, 则自动前面补全0
	var octal_number = 437
	var octal_number_string string = fmt.Sprintf("%v", octal_number)
	var octal_convert_binary_string = ""
	for _, s := range strings.Split(octal_number_string, "")  { //拆分每1位数字, 转为二进制
		num, _ := strconv.Atoi(s)
		sub_string := ""
		for ; num > 0; num /= 2 {
			lsb := num % 2
			sub_string = fmt.Sprintf("%d", lsb) + sub_string
		}

		for i := 0; i < 3 - len(sub_string); i++ {  //不足3位, 补全
			sub_string = "0" + sub_string
		}

		octal_convert_binary_string = octal_convert_binary_string + sub_string
	}
	fmt.Printf("八进制转二进制: %v -> %v\n", octal_number_string, octal_convert_binary_string)

	//------- 十六进制转二进制
	//十六进制的每一位转4位数的二进制, 不足4位的, 自动补全0
	var sixteen_number int = 0xA2CE9
	var sixteen_number_string = fmt.Sprintf("%x", sixteen_number)
	var sixteen_convert_binary_string string = ""
	for _, s := range strings.Split(sixteen_number_string, "")  {
		num, _ :=  strconv.ParseInt(s, 16, 64)
		sub_string := ""
		for ; num > 0; num /= 2 {
			lsb := num % 2
			sub_string = fmt.Sprintf("%d", lsb) + sub_string
		}

		for i := 0; i < 4 - len(sub_string); i++ {  //不足3位, 补全
			sub_string = "0" + sub_string
		}
		sixteen_convert_binary_string = sixteen_convert_binary_string + sub_string
	}
	fmt.Printf("十六进制转二进制: %v -> %v\n", sixteen_number_string, sixteen_convert_binary_string)


	//二进制转八进制  从右开始, 每3位一组, 转为八进制(转十进制"因为3位二进制最大为111 = 7, 刚好在八进制范围内")   (从上面的八转二反向推导)
	octal_convert_binary_string += "0"
	var splits int = len(octal_convert_binary_string) / 3
	var octal_convert_binary_string_len = len(octal_convert_binary_string)
	var binary_convert_eight_string string = ""
	if len(octal_convert_binary_string) % 3 != 0 {
		splits++
	}

	end := octal_convert_binary_string_len
	begin := end - 3
	//注意:go的切片两个值是开始坐标和结束坐标  不是常规的location和length
	for i := splits - 1; i >= 0 ; i-- {
		if begin < 0 {
			begin = 0
		}

		println(octal_convert_binary_string)
		s := octal_convert_binary_string[begin: end]
		println(s)
		var number int64 = 0
		//把简短的二进制转为10进制  为毛这里是转十进制不会出错 因为长度为3的二进制最大为111, 刚好就是8. 所以把每一位拼接就得到了8进制数
		for i, ss := range strings.Split(s, "") {
			num, _ := strconv.ParseInt(ss, 10, 64)
			num = num * int64(math.Pow(2, float64(len(s) - i - 1)))   //从右往左, 每一位的数字乘以2的(位数- 1)次方
			number += num
		}
		println(number)

		binary_convert_eight_string = fmt.Sprintf("%v", number) + binary_convert_eight_string

		end = end - 3
		begin = end - 3
	}
	fmt.Printf("二进制转八进制:%v -> %v\n", octal_convert_binary_string, binary_convert_eight_string)

	binary_convert_octal_string := binary_convert_other_base(octal_convert_binary_string, 8)
	fmt.Printf("二进制转八进制:%v -> %v\n", octal_convert_binary_string, binary_convert_octal_string)

	binary_convert_hex_string := binary_convert_other_base(octal_convert_binary_string, 16)
	fmt.Printf("二进制转十六进制:%v -> %v\n", octal_convert_binary_string, binary_convert_hex_string)
}

/**
二进制 转八进制或十进制
规则: 把二进制从右往左网按3和4位拆分, 然后把拆分的二进制转为10进制, 再依次拼接得到的数字
为什么按3或4位拆分?  因为3位最大为111=7 而4位最大为15, 刚好在对应的进制内
注意点: 转16进制, 大于9的部分要转为对应的16进制, 如 a, b, c, d等
反向操作也可以得到十六进制转二进制 或 八进制转二进制
*/
func binary_convert_other_base(str string, base int) string {
	if base != 8 && base != 16 {
		fmt.Printf("base参数不正确, base只能为8 或 16")
		return ""
	}

	fmt.Printf("将要转换的二进制字符串:%v\n", str)
	var split_length int = 3
	if base == 16 {
		split_length = 4
	}

	var length = len(str)
	var splits int = length / split_length
	var binary_convert_string string = ""
	if length % split_length != 0 {
		splits++
	}

	end := length
	begin := end - split_length
	//注意:go的切片两个值是开始坐标和结束坐标  不是常规的location和length
	for i := splits - 1; i >= 0 ; i-- {
		if begin < 0 {
			begin = 0
		}

		s := str[begin: end]
		fmt.Printf("拆分的串:%v\n", s)
		var number int64 = 0
		//把简短的二进制转为10进制  为毛这里是转十进制不会出错 因为长度为3的二进制最大为111, 刚好就是8. 所以把每一位拼接就得到了8进制数
		for i, ss := range strings.Split(s, "") {
			num, _ := strconv.ParseInt(ss, 10, 64)
			num = num * int64(math.Pow(2, float64(len(s) - i - 1)))   //从右往左, 每一位的数字乘以2的(位数- 1)次方
			number += num
		}
		fmt.Printf("拆分的串:%v, 转为的十进制数:%v\n", s, number)

		if base == 16 {
			binary_convert_string = fmt.Sprintf("%x", number) + binary_convert_string
		}else {
			binary_convert_string = fmt.Sprintf("%x", number) + binary_convert_string
		}

		end = end - split_length
		begin = end - split_length
	}
	return binary_convert_string
}