package main

import (
	"fmt"
	"time"
)

func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan<- "hello" + fmt.Sprintf("%d", i)
	}

	label_no_value:
	for {
		select {
			//case intChan <- 11:   //此处也可写入数据, 这里写入会一直读写 死循环
			//	fmt.Printf("发送11到intChan\n")
			case i := <-intChan:
				fmt.Printf("获取intChan的值:%v\n", i)
			case v := <-stringChan:
				fmt.Printf("获取stringChan的值:%v\n", v)
			default:
				fmt.Printf("取不到值了\n")
				//return   //此处使用return 直接终止执行
				break label_no_value
		}
	}




	go func() {
		defer func() {
			err := recover()
			if err != nil {
				fmt.Println(err)
			}
		}()

		// recover可以捕获goroutine的异常
		var myMap map[int]string
		myMap[1] = "11"

		//此处的死锁, recover是捕获不到的.
		for {
			_, ok := <-intChan
			if !ok {
				break
			}

		}
	}()
	time.Sleep(time.Second)
	fmt.Printf("执行结束\n")
}



