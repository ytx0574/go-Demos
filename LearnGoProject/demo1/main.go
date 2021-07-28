package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)
/*
基本数据类型及转换
*/

type TestValue struct {
	id int
	name string
}

var mapTestValue map[int]*TestValue = make(map[int]*TestValue, 10)

func main() {
	fmt.Printf("hello world\n")

	var a1 int = 2222
	a2 := 1111
	var a3 int8 = 127
	var a4 int64 = 222

	var b1 float32 = 11
	var b2  float64 = 111
	var c1 bool = true


	fmt.Printf("a1=%d, a2=%d a3=%d, a4= %d\n", a1, a2, a3, a4)
	fmt.Printf("b1=%f, b2=%f b2类型=%T\n", b1, b2, b2)
	fmt.Print(c1)

	var char byte = 'A'
	fmt.Printf("原始值:%v, unicode:%c 整数值:%d 二进制:%b, 八进制:%o\n 十六进制:%x, 科学计数法:%E\n", char, char, char, char, char, char, char)
	var  char1 int8 = '$'
	fmt.Printf("char1 类型:%T, 值:%c \n", char1, char1)
	var  char2 int16 = '哎'
	fmt.Printf("char1 类型:%T, 值:%c %v \n", char2, char2, char2)
	//中文占三个字节. 上面不报错, 是因为哎的uinicode在int16内
	var char3 int32 = '龥'
	fmt.Printf("char3 类型:%T, 值:%c uinicode:%v  uinicode:%d 指针:%p\n", char3, char3, char3, char3, &char3)
	var char4 = &char3 //指向char3的地址
	fmt.Printf("char4=%p\n", char4)

	if (*char4 == char3) { //所以两者相等
		fmt.Print("char3 == char4\n")
	}

	var f3 float64 = -999911111111111111111111546564564654654564564654646546545645645646546456456456456456.12
	var f4 float32 = 999999999999999999999999999999999999
	fmt.Printf("f3 = %f 科学计数:%E\n", f3, f3)
	fmt.Printf("f4 = %f 科学计数:%E\n", f4, f4)

	//类型转换
	str := "string1"
	fmt.Printf("%s\n", str)
	var num1, error1 = strconv.ParseInt(str, 0 , 16)
	fmt.Printf("str转整数:%d, %v\n", num1, error1)

	str2 := "996489489489489849999.999"
	var num2, _ = strconv.ParseFloat(str2, 64)
	//默认宽度9
	fmt.Printf("str2 转float %f\n", num2)
	//精度短了之后 会四舍五入
	fmt.Printf("str2 转float %.2f\n", num2)
	fmt.Printf("str2 转float %1f\n", num2)

	var num3, error3 = strconv.ParseInt(str2, 10, 64)
	fmt.Printf("num3 小数转整数:%v, error:%v\n", num3, error3)
	var num4, error4 = parseCents(str2)
	fmt.Printf("num4 小数转整数:%v, error:%v\n", num4, error4)

	//字符串转整数
	var num5, _ = strconv.Atoi(str2)
	fmt.Printf("num5 %v\n", num5)
	//整数转字符串
	str3 := strconv.Itoa(a1)
	fmt.Printf("整数转string str3=%v\n", str3)


	//基本数据类型转string   1. fmt.Sprintf   2.strconv
	//string转基本数据类型   1. strconv
	str4:= fmt.Sprintf("%d", a1)
	str5:= fmt.Sprintf("%f", b2)
	str6:= fmt.Sprintf("%v", c1)
	fmt.Printf("str4=%v, str5=%s, str6= %q\n", str4, str5, str6)

	str7 := `//基本数据类型转string   1. fmt.Sprintf   2.strconv
	//string转基本数据类型   1. strconv
	str4:= fmt.Sprintf("%d", a1)
	str5:= fmt.Sprintf("%f", b2)
	str6:= fmt.Sprintf("%v", c1)
	fmt.Printf("str4=%v, str5=%s, str6= %q\n", str4, str5, str6)`
	fmt.Printf("str7 `反斜包含的多行串`%v\n", str7)

	if2 := 23
	fmt.Printf("%v\n", if2)

	var ptr1 *int  //赋值后才能指向的地址
	fmt.Printf("ptr1:%v, ptr1:%p\n", ptr1, ptr1)
	ptr1 = &num5
	fmt.Printf("ptr1:%v, ptr1:%p\n", ptr1, ptr1)

	slice := make([]TestValue, 10)
	for i := 0; i < 10; i ++ {
		slice[i] = TestValue{
			id: i,
			name: strconv.Itoa(i),
		}
	}

	for i, v := range  slice {
		//todo slice的遍历其他语言不一样. 这里遍历时, 内部逐步把下标和值赋值一份到i,v 所以i, v在内存属于共享状态
		fmt.Printf("i=%v, id=%v, name=%v  %p = %p\n", i, v.id, v.name, &v, &i)
		mapTestValue[v.id] = &v  //此处不可用

		//var new_v = v
		//mapTestValue[v.id] = &new_v
	}

	fmt.Println(mapTestValue)

	fmt.Printf("%v %v\n", math.Pow(9, 10), math.Pow(10, 9))
}

func parseCents(s string) (float64, error) {
	n := strings.SplitN(s, ".", 3)
	if len(n) != 2 {
		err := fmt.Errorf("format error: %s", s)
		return 0, err
	}
	d, err := strconv.ParseInt(n[0], 10, 56)
	if err != nil {
		return 0, err
	}
	c, err := strconv.ParseFloat(n[1], 64)
	clen := len(n[1])
	for i := 0; i < clen; i++ {
		c /= 10
	}
	if err != nil {
		return 0, err
	}
	if d < 0 {
		c = -c
	}

	return float64(d) + c, nil


	////常量每一行都写itoa或只第一行写iota, 常量都是依次递增 常量得访问限制还是根据定义位置和首字母得大小写来控制
	//const (
	//	aaa = iota //0
	//	bbb  //1
	//	ccc, ddd = iota, iota // 2 2  在同一行, 值一样
	//)
	//fmt.Printf("aaa:%d, bbb:%d, ccc:%d\n", aaa, bbb, ccc)
	//
	////complex128 和 complex64  代表 64位和32位得复数
	////复数由两个浮点数组成, 我们把形如z=a+bi（a,b均为实数）的数称为复数，其中a称为实部，b称为虚部，i称为虚数单位。
	////当z的虚部等于零时，常称z为实数；当z的虚部不等于零时，实部等于零时，常称z为纯虚数。
	////复数域是实数域的代数闭包，即任何复系数多项式在复数域中总有根。 复数是由意大利米兰学者卡当在十六世纪首次引入，经过达朗贝尔、棣莫弗、欧拉、高斯等人的工作，此概念逐渐为数学家所接受。
	////复数运算法则: https://baike.baidu.com/item/%E5%A4%8D%E6%95%B0%E8%BF%90%E7%AE%97%E6%B3%95%E5%88%99/2568041?fr=aladdin
	//complex := complex(1, 2)
	//fmt.Printf("得到得复数:%f, 实部:%f, 虚部:%f\n", complex, real(complex), imag(complex))
}