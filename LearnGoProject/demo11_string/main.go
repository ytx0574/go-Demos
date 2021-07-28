package main

import (
	"fmt"
	"io/ioutil"
	"os"
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

	test_string()

	//此处专门找了一下\r \n的区别.  参考链接:https://www.jianshu.com/p/23804b0b03c8
	// \r 回车. 光标出现在行首. 意味着继续输入会覆盖本行之前的内容
	// \n 换行. 直接另起一行显示
	//在unix系统中行尾\n换行.  而windows中行尾变成了 \r\n. 所以我们会经常看到macos编辑的文本, 到了windows上面, 变成了一整行
	//macos早期的版本 行尾是\r. (10.6以前)   后面的版本都是行尾都是\n
	//在golang中, \r前面的字符都使用\r后面的字符覆盖(后面字符长度不足前面的, 前面后半部分会得到保留).  \r不带换行效果
	//使用oc测试时, 发现oc对此的处理不一样. \r直接和\n的处理是一样的, 都是换行  \r\n会换一行
	//要比较两者\r和\n, 只能判断其对应的ASCII码. 在golang下面能看到明显的区别.  在oc中只能通过打印看出
	//golang的字符串输出和写入到文件肉眼可见的内容是不一样的.(参考下面的str5) golang自行处理的输出显示. 写入文件时, 操作系统也会处理对字符串的内容处理
	fmt.Println("aaa\rbbb")
	fmt.Println("aaa\nbbb")
	fmt.Println("aaa\r\nbbb")
	fmt.Println("aaa\n\rbbb")
	fmt.Println("aaa\r1\nbbb")
	fmt.Println("aaa\n1\rbbb")

	str5 := "\r\raaa[\r\r]bbb\r\r"
	str6 := "\n\naaa[\n\n]bbb\n\n"
	str7 := "\r\raaa[\r\r]bbb\r\rsjoi333jo\rOOOO\r\r1111\r2222\r\r\r7"
	fmt.Println(str7)


	f, _ := os.Create("/Users/xxx/Desktop/go_aaaa1.txt")
	f.Write([]byte(str5))
	defer f.Close()

	f2, _ := os.Create("/Users/xxx/Desktop/go_aaaa2.txt")
	f2.Write([]byte(str6))
	defer f2.Close()

	fmt.Println("------")
	//此处str5 = ]bbb\n
	fmt.Printf("%q\n", str5)  //输出原值, "\r\raaa[\r\r]bbb\r\r"
	fmt.Printf("%v\n", str5)  //输出格式化后的字符串 "]bbb "
	fmt.Printf("str5 = %v\n", str5) //输出"]bbb = "  //因为最后的\r\r后面没有字符. 使用到前面的]bbb. 又因为"str5 = "比"]bbb"长, 所以最终输出 "]bbb = "
	fmt.Println("------")

	//此处注意. string其实就是byte数组  所有这里可以直接拷贝一个string到byte数组中
	var slice []byte = make([]byte, 5)
	copy(slice, "我2")
	fmt.Printf("slice = %v, slice = %b\n", slice, slice)

	//使用反引号初始化一大段字符串.   内部严格遵守``的字符串格式. 自行在里面添加\n是无效的, 内部所有的符号都会原样输出
	str8 := `
用t保存全局的itabTable地址，然后使用t.find函数查找，这么做是为了防止在查找
	过程中itabTable被替换导致错误
	如果未找到，再尝试加锁查找。原因是第一步查找时可能有另一个协程并发写入，从而导致Find函数未找到但实际数据是
	存在的。这时通过加锁防止itabTable被写入，然后在itabTable中查找
	如果扔为找到，此时根据接口类型和数据类型生成一个新的itab插入itabTable中。如果插入失败，则panic 注意这里添加时，申请的内存大
	小为len(inter.mhdr)-1，前面我们知道fun数组大小为1，所以这里再申请内存时只需再申请len(inter.mhdr)-1即可。`
	fmt.Printf("str8 = %v\n", str8)

	CompareTwoString()
}

func test_string()  {
	str1 := "2"
	str2 := "2114"
	fmt.Printf("str1地址:%p, str2地址:%p\n", &str1, &str2)

	str1_bytes := []byte(str1)
	fmt.Printf("%v\n", str1_bytes)

	if str1 == str2 {
		fmt.Printf("str1 == str2\n")
	}
}

func CompareTwoString () {

	sep := "\n"
	allBytes, err := ioutil.ReadFile("/Users/Johnson/Desktop/localize_All.txt")
	nowBytes, err := ioutil.ReadFile("/Users/Johnson/Desktop/localize_Now.txt")
	if err != nil {
		fmt.Printf("read file err:%v\n", err)
	}

	slice := make([]string, 0)
	nowSlice := strings.Split(string(nowBytes), sep)
	nowMap := make(map[string]interface{})

	for _, v := range nowSlice {
		nowMap[v] = ""
	}

	for _, v := range strings.Split(string(allBytes), sep) {
		if _, ok := nowMap[v]; ok == false {
			slice = append(slice, v)
		}
	}

	fmt.Println(slice)
	ioutil.WriteFile("/Users/Johnson/Desktop/localize_Diff.txt", []byte(strings.Join(slice, "\n")), os.ModePerm)
}