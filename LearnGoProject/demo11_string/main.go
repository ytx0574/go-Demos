package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//字符串操作
//go中字符串由一个byte数组组成, 所以两者可以相互转换
func main() {
	//字符串遍历
	str1 := "我是John, 你还好吗?"
	for i, s := range []rune(str1) {
		fmt.Printf("i = %v, s = %q\n", i, s)
	}

	str2 := "12343"
	num1, _ := strconv.Atoi(str2)
	str3 := "12341lk"
	num2, _ := strconv.Atoi(str3)
	fmt.Printf("num1 = %v, num2 = %v\n", num1, num2)

	fmt.Printf("int 111 convert string %v\n", strconv.Itoa(111))
	num3, _ := strconv.ParseInt("111", 16, 64)
	fmt.Printf("num3 = %v\n", num3)

	//字符串转byte
	var bytes = []byte(str1)
	fmt.Printf("string convert bytes \n%%v:%v, \n%%s:%s  \n%%#v:%#v \n%%+v:%+v\n", bytes, bytes, bytes, bytes)

	//十进制转2 8 16进制
	fmt.Printf("十进制转16 %v\n", strconv.FormatInt(11, 16))
	fmt.Printf("十进制转2 %v\n", strconv.FormatInt(11, 2))
	fmt.Printf("十进制转8 %v\n", strconv.FormatInt(11, 8))


	str4 := "   我刚三 伏   天 系   是你 2时代风范来了时代风我统jljoijo223dfs   发    "
	//判断字符串包含  前缀 后缀
	fmt.Printf("strings.Contains 包含 %v\n", strings.Contains(str4, "三"))
	fmt.Printf("strings.ContainsAny 包含子串任意字符 %v\n", strings.ContainsAny(str4, "接送机哦日"))
	fmt.Printf("strings.HasPrefix %v\n", strings.HasPrefix(str4, "  "))
	fmt.Printf("strings.HasPrefix %v\n", strings.HasSuffix(str4, "dfs   发    "))

	fmt.Printf("strings.Count 子串个数%v\n", strings.Count(str4, "我"))
	fmt.Printf("strings.Index 第一次出现子串的位置, 没有为-1: %v\n", strings.Index(str4, "风"))
	fmt.Printf("strings.IndexAny 第一次出现子串的位置, 没有为-1: %v\n", strings.IndexAny(str4, "建瓯市"))
	fmt.Printf("string.LastIndex 最后一次出现子串的位置, 没有为-1: %v\n", strings.LastIndex(str4, "三"))

	fmt.Printf("不区分大小写比较:%v\n", strings.EqualFold("jjj", "JJj"))
	fmt.Printf("转大写:%v\n", strings.ToUpper("我师傅as2os"))
	fmt.Printf("转小写:%v\n", strings.ToLower("我师傅asAA2os"))
	fmt.Printf("转小写:%v\n", strings.ToTitle("swosf22of"))
	fmt.Printf("转大写:%v\n", strings.ToUpperSpecial(unicode.TurkishCase, "2fsow松紧藕粉"))

	// toUppper和toTitle大部分情况相同, 只是在处理某些unicode编码字符时有所不同
	str_special := "ǳ ǵǵǳǳǳ hello world！"
	fmt.Println(strings.Title(str_special))   // ǲ Ǵǵǳǳǳ Hello World！
	fmt.Println(strings.ToTitle(str_special)) // ǲ ǴǴǲǲǲ HELLO WORLD！
	fmt.Println(strings.ToUpper(str_special)) // Ǳ ǴǴǱǱǱ HELLO WORLD！


	//重复字符串
	fmt.Printf("strings.repeat  重复组合字符串%v\n", strings.Repeat("1234 ", 11))
	//替换字符串内容.  最后一个参数为替换的数量.  为-1就是替换全部. 参考ReplaceAll
	fmt.Printf("strings.Replace 替换字符串%v\n", strings.Replace(str4, "j", "[Replaced]", -1))
	//替换所有字符串内容
	fmt.Printf("strings.ReplaceAll 替换字符串%v\n", strings.ReplaceAll(str4, "j", "[replaceall]"))
	//分割字符串
	fmt.Printf("string.Split   		 		  按字符分割字符串:%v\n", strings.Split(str4, "时代风"))
	fmt.Printf("string.SplitAfter 按字符分割字符串, 参与分割的字符会被保留:%v\n", strings.SplitAfter(str4, "时代风"))
	fmt.Printf("strings.SplitN  按字符串分割字符串. 指定分割段数量%v\n", strings.SplitN(str4, "", 5))
	fmt.Printf("strings.Fields 按照空格分割字符串%v\n", strings.Fields("str4sss 11"))

	//去除字符串两边的字符
	fmt.Printf("strings.TrimSpace 移除字符串左右两边的空格: %v\n",strings.TrimSpace(str4))
	fmt.Printf("strings.Trim 移除字符串左右两边的空格: %v\n",strings.Trim(str4, " "))
	fmt.Printf("strings.TrimLeft 移除字符串右边的空格: %v\n",strings.TrimRight(str4, " "))
	fmt.Printf("strings.TrimPrefix 移除字符串相同的前缀: %v\n",strings.TrimPrefix(str4, "   我"))

	//拼接字符串数组为一个字符串
	fmt.Printf("strings.Join 拼接字符串数组: %v\n", strings.Join([]string{"11", "22", "333"}, "<>"))

	//strings.Reader 此类型不懂 后续再来学
	fmt.Printf("%v\n", strings.NewReader("wosoojojofs"))
}