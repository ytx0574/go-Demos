package main

import (
	"C"
	"flag"
	"fmt"
	"os"
)

/**命令行的参数的解析*/
func main() {

	//-flag
	//-flag=x
	//-flag x  // 只有非bool类型的flag可以  此处bool类型的字段, 需要使用=

	fmt.Printf("Args:%v\n", os.Args)

	var (
		host string
		port int
		debug bool
	)

	flag.StringVar(&host, "host", "127.0.0.1", "host, 默认为localhost")
	flag.IntVar(&port, "port", 5900, "默认端口")
	flag.BoolVar(&debug, "d", false, "是否开启debug模式, 默认为false")
	//var debug1 *bool = flag.Bool("d", false, "是否开启debug模式d, 默认为false")
	// flag.Bool 和 flag.BoolVar 不可同时用, 两者都是指针返回的 已设置一个变量接收后, 不可再次设置变量

	flag.Parse()

	flag.PrintDefaults()

	fmt.Printf("host = %v, port = %d\n", host, port)
	fmt.Printf("debug = %t\n", debug)
	//fmt.Printf("debug1 = %t\n", *debug1)

	//时间格式解析 遵循time.ParseDuration()的格式
	//time.ParseDuration()
	//flag.DurationVar()
}
