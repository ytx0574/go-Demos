package utils_demo9

import "fmt"

var ConstString string = "我是一句常量串"

func init () {
	fmt.Printf("demo9_package2 j_utils util2 init()\n")
}

func Print_string ()  {
	fmt.Print("utils_demo9 utils 就是打印一句话\n")
}

func print_string ()  {
	fmt.Print("utils_demo9 utils2 就是打印一句话\n")
}