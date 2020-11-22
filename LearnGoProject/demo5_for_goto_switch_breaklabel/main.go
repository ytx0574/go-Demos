package main

import "fmt"

func aaaa()  {
	
}
//switch  默认就带break
//fallthrough  满足条件后继续执行下一个条件, 只执行下一个条件的原因(默认也是带了break, 所以会出去)
//for 支持label标签, 用于跳出指定循环  声明标签在for循环前面, 如(label1:),  使用时使用(break label1)即可  此时的break就可跳出指定循环
//goto  直接跳转到指定标签的语句后面接着执行  声明可以写在任意地方,如(gotolabel1:), 使用(goto gotolabel1), 如果循环里面, 也会立即执行后面的, 进而跳出循环
func main () {

	//for true {
	//	var sex string
	//	fmt.Printf("请输入性别:\n")
	//	fmt.Scanf("%v", &sex)
	//	fmt.Printf("你输入的字符串\"%v\"\n", sex)
	//	switch sex {
	//	case "男":
	//		fmt.Printf("这是一个男人\n")
	//		fallthrough
	//	case "女":
	//		fmt.Printf("你是一个女的\n")
	//		fallthrough
	//	case "人妖", "变性人":
	//		fmt.Printf("您是一个人妖或者变性人\n")
	//	default:
	//		fmt.Printf("不知道你的什么性别\n")
	//	}
	//}


	//break 配合label使用
	lebel1:
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("i = %v, j = %v\n", i, j)
			if i == 1 {
				break lebel1
			}
		}
	}

	//golang  goto语句的使用
	var i int = 0
	for ; i < 10; i++ {
		if i == 3 {
			fmt.Printf("i == 3 时, 执行goto\n")
			goto gotolabel  //此处使用goto也是直接跳出这个循环, 直接执行下面的打印c
			goto gotolabel2 //因为gotolabel2的执行还是接着for循环, 所以循环还会继续
		}
		gotolabel2:
		fmt.Printf("i = %v\n", i)
	}

	fmt.Printf("a\n")
	gotolabel:
	fmt.Printf("c\n")


	string1 := "wo shi yi ju hua !"
	for i := 0; i < len(string1); i++ {
		fmt.Printf("%c\n", string1[i])
	}

	//中文字符串不可采用默认的遍历形式, 乱码(中文字符是三个字节)
	string2 := "我是一句话 !"
	string2_len := len(string2)
	for i := 0; i < len(string2); i++ {
		fmt.Printf("%c, string2_len = %v\n", string2[i], string2_len)
	}

	//遍历字符串, 下标按实际字符的字节来计算的
	for i, str := range string2 {
		fmt.Printf("i = %v, str = %c\n", i, str)
	}

	//rune == Int32
	rune_string2 := []int32(string2)
	rune_string2_len := len(rune_string2)
	fmt.Printf("rune_string2 = %c, rune_string2_len = %v\n", rune_string2, rune_string2_len)
	for index := range rune_string2{
		fmt.Printf("index = %v  value:%c\n", index, rune_string2[index])
	}

}